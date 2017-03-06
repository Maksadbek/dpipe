package outputs

import (
	"log"

	"github.com/maksadbek/dpipe"
)

type Outputs map[string]dpipe.Output

func (o Outputs) Add(name string, output dpipe.Output) {
	if _, ok := o[name]; !ok {
		o[name] = output
	} else {
		log.Printf("W! outputs: dublicate ouput name: '%s'", name)
	}
}

var AllOutputs Outputs
