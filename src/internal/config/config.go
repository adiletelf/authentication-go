package config

import (
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	ListenAddress    string
	ConnectionString string
	DatabaseName     string
	CollectionName   string
}

func New() (*Config, error) {
	config := &Config{}
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(config)
	return config, err
}
