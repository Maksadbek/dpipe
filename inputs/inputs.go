package inputs

import (
	"log"

	"github.com/maksadbek/dpipe"
)

type Inputs map[string]dpipe.Input

// Add can be used to register input implementation
// with its name.
// The input is not added if such name already exists.
func (o Inputs) Add(name string, input dpipe.Input) {
	if _, ok := o[name]; !ok {
		o[name] = input
	} else {
		log.Printf("W! inputs: dublicate input name: '%s'", name)
	}
}

// AllInputs is the map that keeps input name
// mapped to input implementation
var All Inputs
