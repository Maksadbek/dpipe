package config

import "github.com/spf13/viper"

const (
	configType = "toml"
)

type Config struct {
	Main *viper.Viper
}

// New creates config from given file path
func New(path string) (*Config, error) {
	v := viper.New()
	v.SetConfigType(configType)
	v.AddConfigPath(".")
	v.SetConfigName(path)

	err := v.ReadInConfig()

	if err != nil {
		return nil, err
	}

	return &Config{
		Main: v,
	}, nil
}

func (c *Config) Outputs() *viper.Viper {
	return c.Main.Sub("outputs")
}

func (c *Config) Inputs() *viper.Viper {
	return c.Main.Sub("inputs")
}

func (c *Config) Filters() *viper.Viper {
	return c.Main.Sub("filters")
}

func (c *Config) Aggregators() *viper.Viper {
	return c.Main.Sub("aggregators")
}

// GetAllKeys is the helper function to
// get all sub sections of given section.
func GetAllKeys(v *viper.Viper) []string {
	keys := []string{}
	for k, _ := range v.AllSettings() {
		keys = append(keys, k)
	}
	return keys
}
