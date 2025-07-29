package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		log.Printf("Config not found, using ENV")
	}

	err = viper.Unmarshal(&config)
	return
}
