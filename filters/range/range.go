package rangefilter

import (
	"github.com/maksadbek/dpipe/filters"
	"github.com/spf13/viper"
)

type Range struct {
	min int
	max int
}

func (r *Range) LoadConf(v *viper.Viper) error {
	r.min = v.GetInt("min")
	r.max = v.GetInt("max")

	return nil
}

// Validate checks the v for range between min and max
func (r *Range) Validate(v interface{}) bool {
	if value, ok := v.(int); ok {
		if value < r.min {
			return false
		}

		if value > r.max {
			return false
		}
	} else {
		return false
	}

	return true
}

func init() {
	filters.Add("range", &Range{})
}
