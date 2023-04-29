package utils

import (
	"fmt"

	"github.com/spf13/viper"
)

var Config *viper.Viper

func init() {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("./")
	err := v.ReadInConfig()
	if err != nil {
		fmt.Println("Error when loading config file \n", err)
	}
	Config = v
}
