package dpipe

import "github.com/spf13/viper"

// Filter checks value of field
// and returns True/False if it passes validation
type Filter interface {
	Validate(interface{}) bool
	LoadConf(*viper.Viper) error
}
