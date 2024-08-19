package config

import (
	"github.com/spf13/viper"
	"log"
)

func ReadConfig() {
	viper.AddConfigPath("config")
	viper.SetConfigName("application")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("error while reading config file %s", err)
	}
}
