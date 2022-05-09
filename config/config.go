package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
)

var c Config

func InitConfig() *Config {

	// 程序执行路径
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path)

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)
	err = viper.ReadInConfig()

	err = viper.Unmarshal(&c)
	if err != nil {
		fmt.Printf("unable to decode into struct, %v", err)
	}
	return &c
}

func GetConfig() *Config {
	return &c
}
