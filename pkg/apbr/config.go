package apbr

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type ConfigServerAuthentication struct {
	Username string
	Password string
	Token    string
}

type ConfigServer struct {
	Host           string
	Name           string
	Authentication ConfigServerAuthentication
}

type Config struct {
	Type          string
	Servers       []ConfigServer
	DefaultServer string
}

var config Config

func init() {
	viper.SetConfigName("config")      // name of config file (without extension)
	viper.SetConfigType("yaml")        // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("$HOME/.apbr") // call multiple times to add many search paths
	viper.AddConfigPath(".")           // optionally look for config in the working directory
	err := viper.ReadInConfig()        // Find and read the config file
	if err != nil {                    // Handle errors reading the config file
		log.Error("config file not found")
	}
	err = viper.Unmarshal(&config)

	if err != nil { // Handle errors reading the config file
		log.Error(err.Error())
	}
}
