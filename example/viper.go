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
	x := viper.New()
	y := viper.New()

	x.SetConfigFile("./config.default.yaml")
	x.AddConfigPath(".")
	err := x.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("%v\n", err))
	}

	var C config
	err = x.Unmarshal(&C)
	if err != nil {
		panic(fmt.Errorf("%v\n", err))
	}

	y.SetConfigFile("./config.yaml")
	y.AddConfigPath(".")
	err = y.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("%v\n", err))
	}

	err = y.Unmarshal(&C)
	if err != nil {
		panic(fmt.Errorf("%v\n", err))
	}

	fmt.Printf("weight=%d\n", C.module.weight)
}