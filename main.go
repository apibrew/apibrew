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

	var port = 9009
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	stub.RegisterWorkSpaceServiceServer(grpcServer, service.NewWorkSpaceService())
	grpcServer.Serve(lis)
}
