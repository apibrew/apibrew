package lib

import (
	"data-handler/app"
	"data-handler/grpc/stub"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

var authenticationServiceClient stub.AuthenticationServiceClient
var dataSourceServiceClient stub.DataSourceServiceClient
var resourceServiceClient stub.ResourceServiceClient
var recordServiceClient stub.RecordServiceClient

func init() {
	application := new(app.App)

	var initData = prepareInitData()

	application.SetInitData(initData)

	application.Init()

	go application.Serve()
	time.Sleep(10 * time.Millisecond)

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(initData.Config.GrpcAddr, opts...)

	if err != nil {
		panic(err)
	}

	recordServiceClient = stub.NewRecordServiceClient(conn)
	authenticationServiceClient = stub.NewAuthenticationServiceClient(conn)
	resourceServiceClient = stub.NewResourceServiceClient(conn)
	dataSourceServiceClient = stub.NewDataSourceServiceClient(conn)
}
