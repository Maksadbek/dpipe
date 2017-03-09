package urlfilter

import (
	"net/url"

	"github.com/maksadbek/dpipe/filters"
	"github.com/spf13/viper"
)

// URI filter checks given string
// for against RFC3986 specification: https://www.ietf.org/rfc/rfc3986.txt
// uses url package
type URLFilter struct{}

// URI does not need any configs
func (u *URLFilter) LoadConf(v *viper.Viper) error {
	return nil
}

// Validate casts interface into string
// and tries to parse URI, if err != nil, then
// given URI is not valid
func (u *URLFilter) Validate(v interface{}) bool {
	if value, ok := v.(string); ok {
		_, err := url.ParseRequestURI(value)
		if err != nil {
			return false
		}
		return true
	} else {
		return false
	}
}

func init() {
	filters.Add("url", &URLFilter{})
}
