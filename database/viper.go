package database

import (
	"github.com/spf13/viper"
	"log"
)

func viperEnvVariable(key string) string {
	viper.SetConfigFile("../runconfig/.env")
	error := viper.ReadInConfig()

	if error != nil {
		log.Fatal("Error while reading config file", error)
		panic(error)
	}

	value, ok := viper.Get(key).(string)

	if !ok {
		log.Fatal("Invalid type assertion")
		panic(error)
	}
	return value
}
