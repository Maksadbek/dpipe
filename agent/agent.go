package agent

import (
	"sync"

	"log"

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
}

func New(conf *config.Config) *Agent {
	return &Agent{
		config: conf,
	}
}

// Init initializes inputs and outputs
func (a *Agent) Init() {
	inputs.All.Init(a.config.Inputs())
	outputs.All.Init(a.config.Outputs())
}

// Runs starts running all inputs
// and passes received data into outputs
func (a *Agent) Run() {
	wg := sync.WaitGroup{}

	go func() {
		for h := range a.gatherer.hotelsc {
			wg.Add(1)
			for _, name := range config.GetAllKeys(a.config.Outputs()) {
				output, ok := outputs.All[name]
				if ok {
					output.Write(h)
				} else {
					log.Printf("E! no output with name: '%s'", name)
				}
			}
			wg.Done()
		}
	}()

	for _, name := range config.GetAllKeys(a.config.Inputs()) {
		input, ok := inputs.All[name]
		if ok {
			input.Read(a.gatherer)
		} else {
			log.Printf("E! no input with name: '%s'", name)
		}
	}

	wg.Wait()
}
