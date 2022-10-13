package test

import (
	"data-handler/app"
	"data-handler/stub"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

type SimpleAppGrpcContainer struct {
	app.GrpcContainer

	authenticationService stub.AuthenticationServiceClient
	dataSourceService     stub.DataSourceServiceClient
	resourceService       stub.ResourceServiceClient
	recordService         stub.RecordServiceClient
}

func (receiver SimpleAppGrpcContainer) GetRecordService() stub.RecordServiceClient {
	return receiver.recordService
}
func (receiver SimpleAppGrpcContainer) GetAuthenticationService() stub.AuthenticationServiceClient {
	return receiver.authenticationService
}
func (receiver SimpleAppGrpcContainer) GetResourceService() stub.ResourceServiceClient {
	return receiver.resourceService
}
func (receiver SimpleAppGrpcContainer) GetDataSourceService() stub.DataSourceServiceClient {
	return receiver.dataSourceService
}

func withClient(fn func(container *SimpleAppGrpcContainer)) {
	withApp(func(application *app.App) {
		var opts []grpc.DialOption
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

		conn, err := grpc.Dial(application.Addr, opts...)

		defer conn.Close()

		container := &SimpleAppGrpcContainer{
			recordService:         stub.NewRecordServiceClient(conn),
			authenticationService: stub.NewAuthenticationServiceClient(conn),
			resourceService:       stub.NewResourceServiceClient(conn),
			dataSourceService:     stub.NewDataSourceServiceClient(conn),
		}

		fn(container)

		if err != nil {
			panic(err)
		}

	})
}

func withApp(exec func(application *app.App)) {
	application := new(app.App)

	application.Addr = "127.0.0.1:17912"

	application.SetInitData(prepareInitData())

	application.Init()

	go application.Serve()
	time.Sleep(10 * time.Millisecond)

	exec(application)

	defer application.Stop()
}
