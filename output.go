package dpipe

import "github.com/spf13/viper"

type Output interface {
	Write(Hotel) error

	LoadConf(*viper.Viper) error
}
