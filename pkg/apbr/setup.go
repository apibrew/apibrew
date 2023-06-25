package apbr

import (
	"github.com/apibrew/apibrew/pkg/apbr/flags"
	"github.com/apibrew/apibrew/pkg/client"
	log "github.com/sirupsen/logrus"
)

var selectorFlags = flags.NewSelectorFlags(GetDhClient)
var overrideFlags = flags.NewOverrideFlags()

var dhClient client.DhClient

func GetDhClient() client.DhClient {
	if dhClient == nil {
		prepareDhClient()
	}

	return dhClient
}

func prepareDhClient() {
	configServer := locateConfigServer()

	var err error
	if err != nil {
		log.Fatal(err)
		return
	}

	dhClient, err = client.NewDhClient(client.DhClientParams{
		Addr:     configServer.Host,
		Insecure: true,
	})

	if configServer.Authentication.Token != "" {
		dhClient.AuthenticateWithToken(configServer.Authentication.Token)
	} else {
		err = dhClient.AuthenticateWithUsernameAndPassword(configServer.Authentication.Username, configServer.Authentication.Password)

		if err != nil {
			log.Fatal(err)
		}
	}

	if err != nil {
		log.Fatal(err)
	}
}

func locateConfigServer() ConfigServer {
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

	log.Fatal("could not find apbr-server with name: " + server)

	return ConfigServer{}
}
