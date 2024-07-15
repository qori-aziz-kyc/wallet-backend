package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func InitConfig() error {
	viper.SetConfigType("yaml")
	viper.SetConfigName("env")
	viper.AddConfigPath("./")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("error reading config file: %v", err))
	}

	return nil
}
