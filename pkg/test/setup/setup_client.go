package setup

import (
	"fmt"
	"github.com/apibrew/apibrew/pkg/client"
	grpc2 "github.com/apibrew/apibrew/pkg/server/grpc"
	"github.com/apibrew/apibrew/pkg/service"
	"github.com/apibrew/apibrew/pkg/service/impl"
	"github.com/apibrew/apibrew/pkg/stub"
	log "github.com/sirupsen/logrus"
	"net"
	"time"
)

var authenticationClient stub.AuthenticationClient
var resourceClient stub.ResourceClient
var recordClient stub.RecordClient
var dataSourceClient stub.DataSourceClient

var container service.Container

var dhClient client.DhClient

func GetTestDhClient() client.DhClient {
	return dhClient
}

func GetContainer() service.Container {
	return container
}

func initDb() {

}

func initClient() {
	initDb()

	log.SetLevel(log.DebugLevel)
	log.SetReportCaller(false)

	application := new(impl.App)

	var config = prepareInitData()

	addr := fmt.Sprintf("%s:%d", config.Host, config.Port)

	application.SetConfig(config)

	<-application.Init()

	grpcServer := grpc2.NewGrpcServer(application)
	grpcServer.Init(config)

	container = application

	l, err := net.Listen("tcp", fmt.Sprintf("%s:%d", config.Host, config.Port))
	if err != nil {
		log.Fatal(err)
	}

	go grpcServer.Serve(l)

	time.Sleep(10 * time.Millisecond)

	dhClient, err = client.NewDhClient(client.DhClientParams{
		Addr:     addr,
		Insecure: true,
	})

	if err != nil {
		log.Fatal(err)
	}

	authenticationClient = dhClient.GetAuthenticationClient()
	resourceClient = dhClient.GetResourceClient()
	recordClient = dhClient.GetRecordClient()
	dataSourceClient = dhClient.GetDataSourceClient()
}
