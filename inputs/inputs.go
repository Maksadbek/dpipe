package inputs

import (
	"log"

	"github.com/maksadbek/dpipe"
	"github.com/spf13/viper"
)

type Inputs map[string]dpipe.Input

// Add can be used to register input implementation
// with its name.
// The input is not added if such name already exists.
func (i Inputs) Add(name string, input dpipe.Input) {
	if _, ok := i[name]; !ok {
		i[name] = input
	} else {
		log.Printf("W! inputs: dublicate input name: '%s'", name)
	}
}

// Init initializes all inputs
// by passing configurations
func (i Inputs) Init(conf *viper.Viper) {
	for name, input := range All {
		input.LoadConf(conf.Sub(name))
	}
}

// AllInputs is the map that keeps input name
// mapped to input implementation
var All Inputs
