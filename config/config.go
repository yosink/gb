package config

import (
	"log"

	"github.com/spf13/viper"
)

func Setup() {
	viper.SetConfigName("conf")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
}
