package aggregators

import (
	"log"

	"github.com/maksadbek/dpipe"
)

var All Aggregators

type Aggregators map[string]dpipe.Aggregator

func (a Aggregators) Add(field string, aggregator dpipe.Aggregator) {
	if _, ok := a[field]; !ok {
		a[field] = aggregator
	} else {
		log.Printf("W! dublicate aggregator name: '%s'", field)
	}
}
