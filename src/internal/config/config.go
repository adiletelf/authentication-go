package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	ListenAddress string
	DB            struct {
		ConnectionString string
		DatabaseName     string
		CollectionName   string
	}
	AccessTokenMinuteLifespan string
	RefreshTokenHourLifespan  string
	ApiSecret                 string
}

var Cfg *Config

func New() (*Config, error) {
	if Cfg != nil {
		return Cfg, nil
	}

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
	Cfg = config
	return Cfg, err
}
