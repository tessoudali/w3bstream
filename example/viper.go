package main

import (
	"fmt"
	"github.com/spf13/viper"
)

type Module struct {
	path  	string
	weight	int
}

type config struct {
	name    string
	module  Module
}

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	var C config
	err = viper.Unmarshal(&C)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	weight := viper.Get("module.weight")
	fmt.Printf("weight=%d\n", weight)
}