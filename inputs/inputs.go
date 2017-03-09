package inputs

import (
	"log"

	"github.com/maksadbek/dpipe"
	"github.com/spf13/viper"
)

var Inputs = map[string]dpipe.Input{}

func RegisteredInputs() []string {
	ins := []string{}
	for name, _ := range Inputs {
		ins = append(ins, name)
	}

	return ins
}

// Add can be used to register input implementation
// with its name.
// The input is not added if such name already exists.
func Add(name string, input dpipe.Input) {
	if _, ok := Inputs[name]; !ok {
		Inputs[name] = input
	} else {
		log.Printf("W! inputs: dublicate input name: '%s'", name)
	}
}

// Init initializes all inputs
// by passing configurations
func Init(conf *viper.Viper) {
	for name, input := range Inputs {
		input.LoadConf(conf.Sub(name))
	}
}
