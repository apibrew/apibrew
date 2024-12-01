package main

import (
	_ "embed"
	"github.com/apibrew/apibrew/module"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/server/rest"
	"github.com/apibrew/apibrew/pkg/service/impl"
	"github.com/apibrew/apibrew/pkg/util"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	log "github.com/sirupsen/logrus"
)

//go:embed aws.json
var configContent string

func main() {
	var err error
	appConfig := &model.AppConfig{}

	err = util.ReadJsonContent(configContent, appConfig)

	if err != nil {
		log.Fatalf("failed to load init data: %v", err)
	}

	log.SetLevel(log.DebugLevel)
	log.SetReportCaller(true)
	log.SetFormatter(&log.TextFormatter{})

	app := new(impl.App)

	app.SetConfig(appConfig)

	initSig := app.Init()

	module.RegisterModules(app)

	restServer := rest.NewServer(app, appConfig)
	restServer.Init()
	handler := restServer.GetHandler()

	<-initSig

	lambda.Start(httpadapter.New(handler).ProxyWithContext)
}
