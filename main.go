package main

import (
	"context"
	"data-handler/service"
	"data-handler/service/backend"
	"data-handler/service/backend/postgres"
	"data-handler/stub"
	"data-handler/stub/model"
	"data-handler/util"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	init := flag.String("init", "", "Initial Data for configuring system")

	flag.Parse()

	initData := &model.InitData{}

	err := util.Read(*init, initData)

	if err != nil {
		log.Fatalf("failed to load init data: %v", err)
	}

	authenticationService := service.NewAuthenticationService()
	dataSourceService := service.NewDataSourceService()
	resourceService := service.NewResourceService()
	recordService := service.NewRecordService()
	postgresResourceServiceBackend := postgres.NewPostgresResourceServiceBackend()
	//workSpaceService := service.NewWorkSpaceService(resourceService)

	initServices(
		dataSourceService,
		resourceService,
		recordService,
		authenticationService,
		postgresResourceServiceBackend,
		initData,
	)

	var port = 9009
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	stub.RegisterResourceServiceServer(grpcServer, resourceService)
	stub.RegisterAuthenticationServiceServer(grpcServer, authenticationService)
	stub.RegisterDataSourceServiceServer(grpcServer, dataSourceService)
	stub.RegisterRecordServiceServer(grpcServer, recordService)

	test(dataSourceService)

	// dead code
	if true {
		return
	}

	err = grpcServer.Serve(lis)

	if err != nil {
		panic(err)
	}
}

func test(dataSourceService service.DataSourceService) {
	var list []*model.DataSource

	for i := 0; i < 1000; i++ {
		list = append(list, &model.DataSource{
			Backend: model.DataSourceBackend_POSTGRESQL,
			Options: &model.DataSource_PostgresqlParams{
				PostgresqlParams: &model.PostgresqlOptions{
					Username:      "root",
					Password:      "52fa536f0c5b85f9d806633937f06446",
					Host:          "tiswork.tisserv.net",
					Port:          5432,
					DbName:        "market",
					DefaultSchema: "public",
				},
			},
		})
	}

	res, err := dataSourceService.Create(context.TODO(), &stub.CreateDataSourceRequest{
		Token:       "empty-token",
		DataSources: list,
	})

	log.Println(res, err)
}

func initServices(dataSourceService service.DataSourceService, resourceService service.ResourceService, recordService service.RecordService, authenticationService service.AuthenticationService, postgresResourceServiceBackend backend.ResourceServiceBackend, initData *model.InitData) {
	dataSourceService.InjectResourceService(resourceService)
	dataSourceService.InjectRecordService(recordService)
	resourceService.InjectDataSourceService(dataSourceService)
	resourceService.InjectAuthenticationService(authenticationService)
	resourceService.InjectPostgresResourceServiceBackend(postgresResourceServiceBackend)
	recordService.InjectPostgresResourceServiceBackend(postgresResourceServiceBackend)
	recordService.InjectDataSourceService(dataSourceService)

	dataSourceService.Init(initData)
	postgresResourceServiceBackend.Init(dataSourceService.GetSystemDataSourceBackend())
	resourceService.Init(initData)
	//workSpaceService.Init(initData)
	recordService.Init(initData)
}
