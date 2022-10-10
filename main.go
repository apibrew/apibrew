package main

import (
	"data-handler/service"
	"data-handler/stub"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	flag.Parse()
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
