package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Database struct {
		Host     string `mapstructure:"HOST"`
		User     string `mapstructure:"USER"`
		Password string `mapstructure:"PASSWORD"`
		Name     string `mapstructure:"NAME"`
		Port     string `mapstructure:"PORT"`
	} `mapstructure:"DATABASE"`
	Server struct {
		Host string `mapstructure:"HOST"`
		Port string `mapstructure:"PORT"`
	} `mapstructure:"SERVER"`
}

func LoadConfig() (config Config, err error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	if err = viper.Unmarshal(&config); err != nil {
		return
	}

	fmt.Printf("Running with config: %+v \n", config)

	return
}
