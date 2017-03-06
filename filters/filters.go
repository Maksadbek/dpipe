package filters

import (
	"log"

	"github.com/maksadbek/dpipe"
)

var All Filters

type Filters map[string]dpipe.Filter

func (f Filters) Add(field string, filter dpipe.Filter) {
	if _, ok := f[field]; !ok {
		f[field] = filter
	} else {
		log.Printf("W! filters: dublicate filter for field: '%s'", field)
	}
}
