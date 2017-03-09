package encodingUTF8

import (
	"unicode/utf8"

	"github.com/maksadbek/dpipe/filters"
	"github.com/spf13/viper"
)

type EncodingUTF8 struct{}

func (e *EncodingUTF8) LoadConf(v *viper.Viper) error {
	return nil
}

func (e *EncodingUTF8) Validate(v interface{}) bool {
	if value, ok := v.(string); ok {
		return utf8.ValidString(value)
	} else {
		return false
	}
}

func init() {
	filters.Add("encodingUTF8", &EncodingUTF8{})
}
