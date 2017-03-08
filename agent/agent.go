package agent

import (
	"log"

	"github.com/maksadbek/dpipe"
	"github.com/maksadbek/dpipe/config"
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
	inputs.All.Init(a.config.Inputs())
	outputs.All.Init(a.config.Outputs())
}

func (a *Agent) CloseOutputs() {
	for _, name := range config.GetAllKeys(a.config.Outputs()) {
		if output, ok := outputs.All[name]; ok {
			output.Close()
		}
	}
}

// Runs starts running all inputs
// and passes received data into outputs
func (a *Agent) Run() {
	go func() {
		for {
			select {
			case h := <-a.gatherer.hotelsc:
				for _, name := range config.GetAllKeys(a.config.Outputs()) {
					output, ok := outputs.All[name]
					if ok {
						err := output.Write(h)
						if err != nil {
							log.Printf("E! failed to write data to output: '%s', error: %v", name, err)
						}
					} else {
						log.Printf("E! no registered output with name: '%s'", name)
						continue
					}
				}
			case <-a.done:
				return
			}
		}
	}()

	for _, name := range config.GetAllKeys(a.config.Inputs()) {
		input, ok := inputs.All[name]
		if ok {
			err := input.Read(a.gatherer)
			if err != nil {
				log.Printf("E! failed to read from input: '%s', error: %v", name, err)
			}
		} else {
			log.Printf("E! no registered input with name: '%s'", name)
			continue
		}
		println("done")
		a.done <- struct{}{}
	}

}
