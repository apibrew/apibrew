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
var dataSourceClient stub.DataSourceClient
var resourceClient stub.ResourceClient

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

	var initData = prepareInitData()

	addr := fmt.Sprintf("%s:%d", initData.Config.Host, initData.Config.Port)

	application.SetInitData(initData)

	<-application.Init()

	grpcServer := grpc2.NewGrpcServer(application)
	grpcServer.Init(initData)

	container = application

	l, err := net.Listen("tcp", fmt.Sprintf("%s:%d", initData.Config.Host, initData.Config.Port))
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
	dataSourceClient = dhClient.GetDataSourceClient()
}
