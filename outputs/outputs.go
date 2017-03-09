package outputs

import (
	"log"

	"github.com/maksadbek/dpipe"
	"github.com/spf13/viper"
)

var Outputs = map[string]dpipe.Output{}

func Add(name string, output dpipe.Output) {
	if _, ok := Outputs[name]; !ok {
		Outputs[name] = output
	} else {
		log.Printf("W! outputs: dublicate ouput name: '%s'", name)
	}
}

// Init initializes all inputs
// by passing configurations
func Init(conf *viper.Viper) {
	for name, output := range Outputs {
		output.LoadConf(conf.Sub(name))
	}
}
