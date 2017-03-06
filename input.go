package dpipe

import "github.com/spf13/viper"

// Input interface must be implemented by
// all input sources
type Input interface {
	// Get receives gatherer
	// reades input data
	// iterates and write data into gatherer
	Read(Gatherer) error

	// LoadConf method must read
	// its config data from viper
	LoadConf(*viper.Viper)
}
