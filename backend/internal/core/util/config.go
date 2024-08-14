package util

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Environment string `mapstructure:"ENV"`
	Database    struct {
		Host      string `mapstructure:"HOST"`
		User      string `mapstructure:"USER"`
		Password  string `mapstructure:"PASSWORD"`
		Name      string `mapstructure:"NAME"`
		Port      string `mapstructure:"PORT"`
		EnableLog bool   `mapstructure:"ENABLELOG"`
	} `mapstructure:"DATABASE"`
	Server struct {
		Host string `mapstructure:"HOST"`
		Port string `mapstructure:"PORT"`
	} `mapstructure:"SERVER"`
	CronJob struct {
		LotteryDomain string            `mapstructure:"DOMAIN"`
		Limit         int16             `mapstructure:"LIMIT"`
		Delay         int16             `mapstructure:"DELAY"`
		Schedule      map[string]string `mapstructure:"SCHEDULE"`
		SkipTimeCheck bool              `mapstructure:"SKIP"`
	} `mapstructure:"JOB"`
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
