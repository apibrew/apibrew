package main

import (
	"context"
	"github.com/apibrew/apibrew/pkg/ext"
	"github.com/apibrew/apibrew/pkg/model"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

type recordExtensionService struct {
	ext.RecordExtensionServer
}

func (r *recordExtensionService) BeforeList(ctx context.Context, req *ext.BeforeListRecordRequest) (*ext.BeforeListRecordResponse, error) {
	log.Print("BeforeList called")

	time.Sleep(1 * time.Second)

	return &ext.BeforeListRecordResponse{}, nil
}

func (r *recordExtensionService) List(ctx context.Context, req *ext.ListRecordRequest) (*ext.ListRecordResponse, error) {
	return &ext.ListRecordResponse{
		Total: 1,
		Records: []*model.Record{
			{
				Resource: "test-resource",
				Version:  76,
			},
		},
	}, nil
}

func main() {
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	ext.RegisterRecordExtensionServer(grpcServer, &recordExtensionService{})

	l, err := net.Listen("tcp", "0.0.0.0:40234")
	if err != nil {
		panic(err)
	}

	panic(grpcServer.Serve(l))
}
