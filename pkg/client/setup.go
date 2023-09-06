package client

import (
	log "github.com/sirupsen/logrus"
)

func locateConfigServer(server string) ConfigServer {
	if server != "" {
		return locateServerByName(server)
	} else {
		return locateServerByName(config.DefaultServer)
	}
}

func locateServerByName(serverName string) ConfigServer {
	for _, item := range config.Servers {
		if item.Name == serverName {
			return item
		}
	}

	log.Fatal("could not find apbr-server with name: " + serverName)

	return ConfigServer{}
}
