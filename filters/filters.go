package filters

import (
	"log"

	"github.com/maksadbek/dpipe"
	"github.com/spf13/viper"
)

var Filters = map[string]dpipe.Filter{}
var FieldFilters = map[string]string{}

func Add(name string, filter dpipe.Filter) {
	Filters[name] = filter
}

func Init(v *viper.Viper) {
	for name, f := range Filters {
		c := v.Sub(name)

		// check if the filter enabled
		var enabled bool
		if c.IsSet("enabled") {
			enabled = c.GetBool("enabled")
		}

		if !enabled {
			continue
		}

		field := c.GetString("field")
		err := f.LoadConf(c)
		if err != nil {
			log.Printf("E! failed to configure filter - '%s', error: %v", name, err)
		} else {
			FieldFilters[field] = name
		}
	}
}

// Validate ranges over all filters
// checks the fields with their names mapped to filters
// if any of them does not pass validation
// the false value is returned
func Validate(h dpipe.Hotel) bool {
	for field, filterName := range FieldFilters {
		if v := h.GetFieldValue(field); v != nil {
			if f, ok := Filters[filterName]; ok {
				if !f.Validate(v) {
					return false
				}
			}
		}
	}

	return true
}
