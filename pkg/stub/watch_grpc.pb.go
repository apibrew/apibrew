// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: stub/watch.proto

package stub

import (
	context "context"
	model "github.com/tislib/apibrew/pkg/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// WatchClient is the client API for Watch service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WatchClient interface {
	// Sends a greeting
	Watch(ctx context.Context, in *WatchRequest, opts ...grpc.CallOption) (Watch_WatchClient, error)
}

type watchClient struct {
	cc grpc.ClientConnInterface
}

func NewWatchClient(cc grpc.ClientConnInterface) WatchClient {
	return &watchClient{cc}
}

func (c *watchClient) Watch(ctx context.Context, in *WatchRequest, opts ...grpc.CallOption) (Watch_WatchClient, error) {
	stream, err := c.cc.NewStream(ctx, &Watch_ServiceDesc.Streams[0], "/stub.Watch/Watch", opts...)
	if err != nil {
		return nil, err
	}
	x := &watchWatchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Watch_WatchClient interface {
	Recv() (*model.Event, error)
	grpc.ClientStream
}

type watchWatchClient struct {
	grpc.ClientStream
}

func (x *watchWatchClient) Recv() (*model.Event, error) {
	m := new(model.Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// WatchServer is the server API for Watch service.
// All implementations must embed UnimplementedWatchServer
// for forward compatibility
type WatchServer interface {
	// Sends a greeting
	Watch(*WatchRequest, Watch_WatchServer) error
	mustEmbedUnimplementedWatchServer()
}

// UnimplementedWatchServer must be embedded to have forward compatible implementations.
type UnimplementedWatchServer struct {
}

func (UnimplementedWatchServer) Watch(*WatchRequest, Watch_WatchServer) error {
	return status.Errorf(codes.Unimplemented, "method Watch not implemented")
}
func (UnimplementedWatchServer) mustEmbedUnimplementedWatchServer() {}

// UnsafeWatchServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WatchServer will
// result in compilation errors.
type UnsafeWatchServer interface {
	mustEmbedUnimplementedWatchServer()
}

func RegisterWatchServer(s grpc.ServiceRegistrar, srv WatchServer) {
	s.RegisterService(&Watch_ServiceDesc, srv)
}

func _Watch_Watch_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(WatchRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(WatchServer).Watch(m, &watchWatchServer{stream})
}

type Watch_WatchServer interface {
	Send(*model.Event) error
	grpc.ServerStream
}

type watchWatchServer struct {
	grpc.ServerStream
}

func (x *watchWatchServer) Send(m *model.Event) error {
	return x.ServerStream.SendMsg(m)
}

// Watch_ServiceDesc is the grpc.ServiceDesc for Watch service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Watch_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "stub.Watch",
	HandlerType: (*WatchServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Watch",
			Handler:       _Watch_Watch_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "stub/watch.proto",
}
