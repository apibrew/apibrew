package setup

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/tislib/data-handler/pkg/abs"
	"github.com/tislib/data-handler/pkg/client"
	grpc2 "github.com/tislib/data-handler/pkg/server/grpc"
	"github.com/tislib/data-handler/pkg/service"
	"github.com/tislib/data-handler/pkg/stub"
	"net"
	"time"
)

var authenticationClient stub.AuthenticationClient
var dataSourceClient stub.DataSourceClient
var resourceClient stub.ResourceClient

var container abs.Container

var dhClient client.DhClient

func GetTestDhClient() client.DhClient {
	return dhClient
}

func GetContainer() abs.Container {
	return container
}

func initClient() {
	log.SetLevel(log.DebugLevel)
	log.SetReportCaller(false)

	application := new(service.App)

	var initData = prepareInitData()

	addr := fmt.Sprintf("%s:%d", initData.Config.Host, initData.Config.Port)

	application.SetInitData(initData)

	application.Init()

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
