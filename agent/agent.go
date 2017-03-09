package agent

import (
	"log"
	"sync"
	"sync/atomic"

	"github.com/maksadbek/dpipe"
	"github.com/maksadbek/dpipe/aggregators"
	"github.com/maksadbek/dpipe/config"
	"github.com/maksadbek/dpipe/filters"
	"github.com/maksadbek/dpipe/inputs"
	"github.com/maksadbek/dpipe/outputs"
)

// Agent initializes all inputs
// runs them and passes its gatherer
// receives metrics
// passes them into outputs
type Agent struct {
	config   *config.Config
	gatherer *Gatherer
	done     chan struct{}

	// Stats is the statistics that Agent
	// creates after processing data from input
	// and sending it to output
	Stats struct {
		// failed to write count
		DataWrittenFailed uint32

		// Data that did not pass validation
		DataValidationFailed uint32

		// succeed writes
		DataWrittenOK uint32

		// All received data count
		DataReceived uint32

		// aggregation metrics
		DataAggregatedFailed uint32
		DataAggregatedOK     uint32
		AggregationErrors    uint32
	}
}

func New(conf *config.Config) *Agent {
	return &Agent{
		config: conf,
		done:   make(chan struct{}),
		gatherer: &Gatherer{
			hotelsc: make(chan dpipe.Hotel),
		},
	}
}

// Init initializes inputs and outputs
func (a *Agent) Init() {
	inputs.Init(a.config.Inputs())
	outputs.Init(a.config.Outputs())
	filters.Init(a.config.Filters())
	aggregators.Init(a.config.Aggregators())
}

// CloseOuputs call Close method on each outputs
func (a *Agent) CloseOutputs() {
	for _, name := range config.GetAllKeys(a.config.Outputs()) {
		if output, ok := outputs.Outputs[name]; ok {
			output.Close()
		}
	}
}

// Runs starts running all inputs
// and passes received data into outputs
// wait group is used in aggregating data
// if there is any aggragator enabled
// data is passed into aggregators
// when inputs done their job and sent signal to a.done channel
// aggregators start aggregating and writing data into outputs
// Run is not exited until all data is not aggregated
func (a *Agent) Run() {
	wg := sync.WaitGroup{}

	var aggrEnabled bool
	if len(aggregators.Aggregators) > 0 {
		wg.Add(1)
		aggrEnabled = true
	}

	go func() {
		for {
			select {
			case h := <-a.gatherer.hotelsc:
				atomic.AddUint32(&a.Stats.DataReceived, 1)

				// validate hotel data
				if !filters.Validate(h) {
					atomic.AddUint32(&a.Stats.DataValidationFailed, 1)
					log.Printf("E! invalid hotel data, skipping")
					continue
				}

				// check if there any aggregators enabled
				// add hotel data into aggregator
				if aggrEnabled {
					for _, aggregator := range aggregators.Aggregators {
						aggregator.Add(h)
					}
					continue
				}
				a.flushOutputs(h)
			case <-a.done:
				if aggrEnabled {
					for field, aggrName := range aggregators.FieldAggregations {
						aggregator := aggregators.Aggregators[aggrName]
						aggrHotels, err := aggregator.Do(field)
						if err != nil {
							atomic.AddUint32(&a.Stats.AggregationErrors, 1)
						}

						for _, h := range aggrHotels {
							a.flushOutputs(h)
							atomic.AddUint32(&a.Stats.DataAggregatedOK, 1)
						}
					}
					wg.Done()
				}
				return
			}
		}
	}()

	for _, name := range config.GetAllKeys(a.config.Inputs()) {
		input, ok := inputs.Inputs[name]
		if ok {
			err := input.Read(a.gatherer)
			if err != nil {
				log.Printf("E! failed to read from input: '%s', error: %v", name, err)
			}
		} else {
			log.Printf("E! no registered input with name: '%s'", name)
			continue
		}
		a.done <- struct{}{}
	}

	// wait untill aggregators do not stop their job
	wg.Wait()
}

// flushOutputs writes hotel data into outputs
func (a *Agent) flushOutputs(h dpipe.Hotel) {
	for _, name := range config.GetAllKeys(a.config.Outputs()) {
		output, ok := outputs.Outputs[name]
		if ok {
			err := output.Write(h)
			if err != nil {
				atomic.AddUint32(&a.Stats.DataWrittenFailed, 1)
				log.Printf("E! failed to write data to output: '%s', error: %v", name, err)
			} else {
				atomic.AddUint32(&a.Stats.DataWrittenOK, 1)
			}
		} else {
			log.Printf("E! no registered output with name: '%s'", name)
			continue
		}
	}
}
