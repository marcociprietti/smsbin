package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var configuration *viper.Viper

func init() {
	configuration = loadConfiguration()
}

func loadConfiguration() *viper.Viper {
	config := *viper.New()

	config.SetConfigName("config")
	config.SetConfigType("yml")
	config.AddConfigPath("./config")

	_ = config.BindEnv("server.host", "SERVER_HOST")
	_ = config.BindEnv("server.port", "SERVER_PORT")

	config.SetDefault("server.host", "127.0.0.1")
	config.SetDefault("server.port", 8080)

	err := config.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("Fatal error loading configuration: %s", err))
	}

	return &config
}

func GetConfig() *viper.Viper {
	return configuration
}
