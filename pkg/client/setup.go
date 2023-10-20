package client

import (
	log "github.com/sirupsen/logrus"
)

func LocateConfigServer(server string) ServerConfig {
	if server != "" {
		return LocateServerByName(server)
	} else {
		return LocateServerByName(config.DefaultServer)
	}
}

func LocateServerByName(serverName string) ServerConfig {
	for _, item := range config.Servers {
		if item.Name == serverName {
			return item
		}
	}

	log.Fatal("could not find apbr-server with name: " + serverName)

	return ServerConfig{}
}
