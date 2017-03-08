package outputs

import (
	"log"

	"github.com/maksadbek/dpipe"
	"github.com/spf13/viper"
)

type Outputs map[string]dpipe.Output

func (o Outputs) Add(name string, output dpipe.Output) {
	if _, ok := o[name]; !ok {
		o[name] = output
	} else {
		log.Printf("W! outputs: dublicate ouput name: '%s'", name)
	}
}

// Init initializes all inputs
// by passing configurations
func (i Outputs) Init(conf *viper.Viper) {
	for name, output := range All {
		output.LoadConf(conf.Sub(name))
	}
}

var All = Outputs{}
