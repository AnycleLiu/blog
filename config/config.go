package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

var config *viper.Viper

func init() {
	config = viper.New()

	config.SetConfigName("config")  // name of config file (without extension)
	config.SetConfigType("json")    // REQUIRED if the config file does not have the extension in the name
	config.AddConfigPath("config/") // optionally look for config in the working directory
	err := config.ReadInConfig()    // Find and read the config file
	if err != nil {                 // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	log.Println("config load success")
}

func GetConfig() *viper.Viper {
	return config
}
