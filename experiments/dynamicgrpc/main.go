package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/dynamicpb"
	"log"
	"net"

	"github.com/ktr0731/grpcdynamic"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	service := grpcdynamic.NewService("api.Example")
	service.RegisterUnaryMethod("Unary", new(dynamicpb.Message), new(dynamicpb.Message), func(ctx context.Context, in interface{}) (interface{}, error) {
		req := in.(*proto.Message)
		desc := descriptorpb.MethodDescriptorProto{
			Name:            nil,
			InputType:       nil,
			OutputType:      nil,
			Options:         nil,
			ClientStreaming: nil,
			ServerStreaming: nil,
		}

		message := dynamicpb.NewMessage()
		return &res{Message: fmt.Sprintf("hi, %s", req)}, nil
	})
	srv := grpcdynamic.NewServer([]*grpcdynamic.Service{service})
	reflection.Register(srv)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	done := make(chan struct{})
	go func() {
		if err := startServer(ctx, srv); err != nil {
			log.Fatal(err)
		}
		close(done)
	}()

	//res, err := callUnaryMethod(service.FullMethodName("Unary"))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(res.Message)
	//
	//cancel()
	<-done
}

func startServer(ctx context.Context, srv *grpc.Server) error {
	l, err := net.Listen("tcp", ":50051")
	if err != nil {
		return err
	}

	go func() {
		<-ctx.Done()
		srv.Stop()
	}()

	if err := srv.Serve(l); err != nil {
		return err
	}
	return nil
}
