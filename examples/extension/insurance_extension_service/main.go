package main

import (
	"context"
	"github.com/tislib/data-handler/pkg/ext"
	"github.com/tislib/data-handler/pkg/model"
	"google.golang.org/grpc"
	"net"
)

type recordExtensionService struct {
	ext.RecordExtensionServiceServer
}

func (r *recordExtensionService) List(ctx context.Context, req *ext.ListRecordRequest) (*ext.ListRecordResponse, error) {
	return &ext.ListRecordResponse{
		Total: 1,
		Content: []*model.Record{
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

	ext.RegisterRecordExtensionServiceServer(grpcServer, &recordExtensionService{})

	l, err := net.Listen("tcp", "0.0.0.0:40234")
	if err != nil {
		panic(err)
	}

	panic(grpcServer.Serve(l))
}
