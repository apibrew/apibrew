// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: stub/namespace.proto

package stub

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// NamespaceServiceClient is the client API for NamespaceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NamespaceServiceClient interface {
	Create(ctx context.Context, in *CreateNamespaceRequest, opts ...grpc.CallOption) (*CreateNamespaceResponse, error)
	List(ctx context.Context, in *ListNamespaceRequest, opts ...grpc.CallOption) (*ListNamespaceResponse, error)
	Update(ctx context.Context, in *UpdateNamespaceRequest, opts ...grpc.CallOption) (*UpdateNamespaceResponse, error)
	Delete(ctx context.Context, in *DeleteNamespaceRequest, opts ...grpc.CallOption) (*DeleteNamespaceResponse, error)
	Get(ctx context.Context, in *GetNamespaceRequest, opts ...grpc.CallOption) (*GetNamespaceResponse, error)
}

type namespaceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNamespaceServiceClient(cc grpc.ClientConnInterface) NamespaceServiceClient {
	return &namespaceServiceClient{cc}
}

func (c *namespaceServiceClient) Create(ctx context.Context, in *CreateNamespaceRequest, opts ...grpc.CallOption) (*CreateNamespaceResponse, error) {
	out := new(CreateNamespaceResponse)
	err := c.cc.Invoke(ctx, "/stub.NamespaceService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *namespaceServiceClient) List(ctx context.Context, in *ListNamespaceRequest, opts ...grpc.CallOption) (*ListNamespaceResponse, error) {
	out := new(ListNamespaceResponse)
	err := c.cc.Invoke(ctx, "/stub.NamespaceService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *namespaceServiceClient) Update(ctx context.Context, in *UpdateNamespaceRequest, opts ...grpc.CallOption) (*UpdateNamespaceResponse, error) {
	out := new(UpdateNamespaceResponse)
	err := c.cc.Invoke(ctx, "/stub.NamespaceService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *namespaceServiceClient) Delete(ctx context.Context, in *DeleteNamespaceRequest, opts ...grpc.CallOption) (*DeleteNamespaceResponse, error) {
	out := new(DeleteNamespaceResponse)
	err := c.cc.Invoke(ctx, "/stub.NamespaceService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *namespaceServiceClient) Get(ctx context.Context, in *GetNamespaceRequest, opts ...grpc.CallOption) (*GetNamespaceResponse, error) {
	out := new(GetNamespaceResponse)
	err := c.cc.Invoke(ctx, "/stub.NamespaceService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NamespaceServiceServer is the server API for NamespaceService service.
// All implementations must embed UnimplementedNamespaceServiceServer
// for forward compatibility
type NamespaceServiceServer interface {
	Create(context.Context, *CreateNamespaceRequest) (*CreateNamespaceResponse, error)
	List(context.Context, *ListNamespaceRequest) (*ListNamespaceResponse, error)
	Update(context.Context, *UpdateNamespaceRequest) (*UpdateNamespaceResponse, error)
	Delete(context.Context, *DeleteNamespaceRequest) (*DeleteNamespaceResponse, error)
	Get(context.Context, *GetNamespaceRequest) (*GetNamespaceResponse, error)
	mustEmbedUnimplementedNamespaceServiceServer()
}

// UnimplementedNamespaceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNamespaceServiceServer struct {
}

func (UnimplementedNamespaceServiceServer) Create(context.Context, *CreateNamespaceRequest) (*CreateNamespaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedNamespaceServiceServer) List(context.Context, *ListNamespaceRequest) (*ListNamespaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedNamespaceServiceServer) Update(context.Context, *UpdateNamespaceRequest) (*UpdateNamespaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedNamespaceServiceServer) Delete(context.Context, *DeleteNamespaceRequest) (*DeleteNamespaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedNamespaceServiceServer) Get(context.Context, *GetNamespaceRequest) (*GetNamespaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedNamespaceServiceServer) mustEmbedUnimplementedNamespaceServiceServer() {}

// UnsafeNamespaceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NamespaceServiceServer will
// result in compilation errors.
type UnsafeNamespaceServiceServer interface {
	mustEmbedUnimplementedNamespaceServiceServer()
}

func RegisterNamespaceServiceServer(s grpc.ServiceRegistrar, srv NamespaceServiceServer) {
	s.RegisterService(&NamespaceService_ServiceDesc, srv)
}

func _NamespaceService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNamespaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NamespaceServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stub.NamespaceService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NamespaceServiceServer).Create(ctx, req.(*CreateNamespaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NamespaceService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNamespaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NamespaceServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stub.NamespaceService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NamespaceServiceServer).List(ctx, req.(*ListNamespaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NamespaceService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNamespaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NamespaceServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stub.NamespaceService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NamespaceServiceServer).Update(ctx, req.(*UpdateNamespaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NamespaceService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteNamespaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NamespaceServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stub.NamespaceService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NamespaceServiceServer).Delete(ctx, req.(*DeleteNamespaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NamespaceService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNamespaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NamespaceServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stub.NamespaceService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NamespaceServiceServer).Get(ctx, req.(*GetNamespaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NamespaceService_ServiceDesc is the grpc.ServiceDesc for NamespaceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NamespaceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "stub.NamespaceService",
	HandlerType: (*NamespaceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _NamespaceService_Create_Handler,
		},
		{
			MethodName: "List",
			Handler:    _NamespaceService_List_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _NamespaceService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _NamespaceService_Delete_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _NamespaceService_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stub/namespace.proto",
}