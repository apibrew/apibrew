package setup

import (
	"fmt"
	"github.com/apibrew/apibrew"
	"github.com/apibrew/apibrew/pkg/client"
	grpc2 "github.com/apibrew/apibrew/pkg/server/grpc"
	"github.com/apibrew/apibrew/pkg/service"
	"github.com/apibrew/apibrew/pkg/service/impl"
	"github.com/apibrew/apibrew/pkg/stub"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"time"
)

var AuthenticationClient stub.AuthenticationClient
var ResourceClient stub.ResourceClient
var RecordClient stub.RecordClient
var DataSourceClient stub.DataSourceClient

var container service.Container

var dhClient client.Client

func GetTestDhClient() client.Client {
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

	initWaitCh := application.Init()

	apibrew.RegisterModules(application)

	<-initWaitCh

	grpcServer := grpc2.NewGrpcServer(application)
	grpcServer.Init()

	container = application

	l, err := net.Listen("tcp", fmt.Sprintf("%s:%d", config.Host, config.Port))
	if err != nil {
		log.Fatal(err)
	}

	go grpcServer.Serve(l)

	time.Sleep(10 * time.Millisecond)

	dhClient, err = client.NewClientWithParams(client.Params{
		Addr:     addr,
		Insecure: true,
	})

	if err != nil {
		log.Fatal(err)
	}

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(addr, opts...)

	if err != nil {
		log.Fatal(err)
	}

	AuthenticationClient = stub.NewAuthenticationClient(conn)
	ResourceClient = stub.NewResourceClient(conn)
	RecordClient = stub.NewRecordClient(conn)
	DataSourceClient = stub.NewDataSourceClient(conn)
}
