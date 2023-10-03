package client

import (
	log "github.com/sirupsen/logrus"
)

func locateConfigServer(server string) ServerConfig {
	if server != "" {
		return locateServerByName(server)
	} else {
		return locateServerByName(config.DefaultServer)
	}
}

func locateServerByName(serverName string) ServerConfig {
	for _, item := range config.Servers {
		if item.Name == serverName {
			return item
		}
	}

	log.Fatal("could not find apbr-server with name: " + serverName)

	return ServerConfig{}
}
