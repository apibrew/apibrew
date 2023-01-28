package main

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/tislib/data-handler/server/stub"
)
import "google.golang.org/grpc"
import "google.golang.org/grpc/credentials/insecure"

var authenticationServiceClient stub.AuthenticationServiceClient
var dataSourceServiceClient stub.DataSourceServiceClient
var resourceServiceClient stub.ResourceServiceClient
var recordServiceClient stub.RecordServiceClient
var authToken string

func initClient(ctx context.Context) {
	configServer := locateConfigServer()
	conn, err := grpc.Dial(configServer.Host, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal(err)
		return
	}

	authenticationServiceClient = stub.NewAuthenticationServiceClient(conn)
	dataSourceServiceClient = stub.NewDataSourceServiceClient(conn)
	resourceServiceClient = stub.NewResourceServiceClient(conn)
	recordServiceClient = stub.NewRecordServiceClient(conn)

	if configServer.Authentication.Token != "" {
		authToken = configServer.Authentication.Token
	} else {
		authResp, err := authenticationServiceClient.Authenticate(ctx, &stub.AuthenticationRequest{
			Username: configServer.Authentication.Username,
			Password: configServer.Authentication.Password,
			Term:     0,
		})

		if err != nil {
			log.Fatal(err)
		}

		authToken = authResp.Token.Content
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

	log.Fatal("could not find server with name: " + server)

	return ConfigServer{}
}
