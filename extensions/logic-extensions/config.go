package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	ApbrAddr         string
	EngineAddr       string
	Token            string
	RemoteEngineAddr string
	ExtensionName    string
}

func NewConfig() *Config {
	viper.AutomaticEnv()

	var config = new(Config)

	err := viper.Unmarshal(&config)

	if err != nil {
		log.Fatal(err)
	}

	if config.ApbrAddr == "" {
		config.ApbrAddr = "localhost:9009"
	}

	if config.EngineAddr == "" {
		config.EngineAddr = "localhost:9008"
	}

	if config.RemoteEngineAddr == "" {
		config.RemoteEngineAddr = "localhost:9008"
	}

	if config.ExtensionName == "" {
		config.ExtensionName = "logic-extensions"
	}

	return config
}
