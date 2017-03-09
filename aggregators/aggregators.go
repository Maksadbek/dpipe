package aggregators

import (
	"log"

	"github.com/maksadbek/dpipe"
	"github.com/spf13/viper"
)

var Aggregators = map[string]dpipe.Aggregator{}
var FieldAggregations = map[string]string{}

func Add(name string, aggregator dpipe.Aggregator) {
	if _, ok := Aggregators[name]; !ok {
		Aggregators[name] = aggregator
	} else {
		log.Printf("W! dublicate aggregator name: '%s'", name)
	}
}

func Init(v *viper.Viper) {
	for name, _ := range Aggregators {
		c := v.Sub(name)

		// check if the aggrgator enabled
		// if aggregators is not enabled
		// remove it from global map
		enabled := c.GetBool("enabled")
		if !enabled {
			delete(Aggregators, name)
			continue
		}

		field := c.GetString("field")
		if field != "" {
			FieldAggregations[field] = name
		}
	}
}
