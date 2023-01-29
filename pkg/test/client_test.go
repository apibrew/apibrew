package test

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/tislib/data-handler/pkg/app"
	grpc2 "github.com/tislib/data-handler/pkg/server/grpc"
	"github.com/tislib/data-handler/pkg/stub"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"time"
)

var authenticationServiceClient stub.AuthenticationServiceClient
var dataSourceServiceClient stub.DataSourceServiceClient
var resourceServiceClient stub.ResourceServiceClient
var recordServiceClient stub.RecordServiceClient
var userServiceClient stub.UserServiceClient

func init() {
	application := new(app.App)

	var initData = prepareInitData()

	addr := fmt.Sprintf("%s:%d", initData.Config.Host, initData.Config.Port)

	application.SetInitData(initData)

	application.Init()

	grpcServer := grpc2.NewGrpcServer(application)
	grpcServer.Init(initData)

	l, err := net.Listen("tcp", fmt.Sprintf("%s:%d", initData.Config.Host, initData.Config.Port))
	if err != nil {
		log.Fatal(err)
	}

	go grpcServer.Serve(l)

	time.Sleep(10 * time.Millisecond)

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(addr, opts...)

	if err != nil {
		panic(err)
	}

	recordServiceClient = stub.NewRecordServiceClient(conn)
	authenticationServiceClient = stub.NewAuthenticationServiceClient(conn)
	resourceServiceClient = stub.NewResourceServiceClient(conn)
	dataSourceServiceClient = stub.NewDataSourceServiceClient(conn)
	userServiceClient = stub.NewUserServiceClient(conn)
}
