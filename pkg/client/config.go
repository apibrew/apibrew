package client

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
	"os"
)

type ConfigServerAuthentication struct {
	Username string
	Password string
	Token    string
}

type ServerConfig struct {
	Host           string
	Port           uint32 // GRPC and HTTP
	HttpPort       uint32 `yaml:"httpPort"` // If not specified, defaults to GRPC PORT
	Name           string
	Authentication *ConfigServerAuthentication
	Insecure       bool
}

type Config struct {
	Type          string         `yaml:"type"`
	Servers       []ServerConfig `yaml:"servers"`
	DefaultServer string         `yaml:"defaultServer"`
}

var configFileFound = false
var config = new(Config)

func GetConfig() *Config {
	return config
}

var homeDir = os.Getenv("HOME")

func AssureConfigFileExists() error {
	if configFileFound {
		return nil
	}

	err := os.MkdirAll(homeDir+"/.apbr", 0775)

	if err != nil {
		return err
	}

	_, err = os.Create(homeDir + "/.apbr/config")

	if err != nil {
		return err
	}

	return nil
}

func init() {
	LoadConfig()
}

func LoadConfig() {
	viper.SetConfigName("config")           // name of config file (without extension)
	viper.SetConfigType("yaml")             // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(homeDir + "/.apbr") // call multiple times to add many search paths
	viper.AddConfigPath(".apbr")            // call multiple times to add many search paths
	viper.AddConfigPath(".")                // optionally look for config in the working directory
	err := viper.ReadInConfig()             // Find and read the config file
	if err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			configFileFound = false
			return
		}
		// Handle errors reading the config file
		log.Error("config file not found")
	}
	err = viper.Unmarshal(&config)

	if err != nil { // Handle errors reading the config file
		log.Error(err.Error())
	}
}

func WriteConfig() error {
	configContent, err := yaml.Marshal(config)

	if err != nil {
		return err
	}

	return os.WriteFile(homeDir+"/.apbr/config", configContent, 0644)
}
