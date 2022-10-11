package main

import (
	"data-handler/service"
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
	//workSpaceService := service.NewWorkSpaceService(resourceService)

	dataSourceService.InjectResourceService(resourceService)
	resourceService.InjectDataSourceService(dataSourceService)
	resourceService.InjectAuthenticationService(authenticationService)

	dataSourceService.Init(initData)
	resourceService.Init(initData)
	//workSpaceService.Init(initData)

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

	err = grpcServer.Serve(lis)

	if err != nil {
		panic(err)
	}
}
