// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: stub/extension.proto

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

const (
	Extension_List_FullMethodName   = "/stub.Extension/List"
	Extension_Get_FullMethodName    = "/stub.Extension/Get"
	Extension_Create_FullMethodName = "/stub.Extension/Create"
	Extension_Update_FullMethodName = "/stub.Extension/Update"
	Extension_Apply_FullMethodName  = "/stub.Extension/Apply"
	Extension_Delete_FullMethodName = "/stub.Extension/Delete"
)

// ExtensionClient is the client API for Extension service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExtensionClient interface {
	List(ctx context.Context, in *ListExtensionRequest, opts ...grpc.CallOption) (*ListExtensionResponse, error)
	Get(ctx context.Context, in *GetExtensionRequest, opts ...grpc.CallOption) (*GetExtensionResponse, error)
	Create(ctx context.Context, in *CreateExtensionRequest, opts ...grpc.CallOption) (*CreateExtensionResponse, error)
	Update(ctx context.Context, in *UpdateExtensionRequest, opts ...grpc.CallOption) (*UpdateExtensionResponse, error)
	Apply(ctx context.Context, in *ApplyExtensionRequest, opts ...grpc.CallOption) (*ApplyExtensionResponse, error)
	Delete(ctx context.Context, in *DeleteExtensionRequest, opts ...grpc.CallOption) (*DeleteExtensionResponse, error)
}

type extensionClient struct {
	cc grpc.ClientConnInterface
}

func NewExtensionClient(cc grpc.ClientConnInterface) ExtensionClient {
	return &extensionClient{cc}
}

func (c *extensionClient) List(ctx context.Context, in *ListExtensionRequest, opts ...grpc.CallOption) (*ListExtensionResponse, error) {
	out := new(ListExtensionResponse)
	err := c.cc.Invoke(ctx, Extension_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *extensionClient) Get(ctx context.Context, in *GetExtensionRequest, opts ...grpc.CallOption) (*GetExtensionResponse, error) {
	out := new(GetExtensionResponse)
	err := c.cc.Invoke(ctx, Extension_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *extensionClient) Create(ctx context.Context, in *CreateExtensionRequest, opts ...grpc.CallOption) (*CreateExtensionResponse, error) {
	out := new(CreateExtensionResponse)
	err := c.cc.Invoke(ctx, Extension_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *extensionClient) Update(ctx context.Context, in *UpdateExtensionRequest, opts ...grpc.CallOption) (*UpdateExtensionResponse, error) {
	out := new(UpdateExtensionResponse)
	err := c.cc.Invoke(ctx, Extension_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *extensionClient) Apply(ctx context.Context, in *ApplyExtensionRequest, opts ...grpc.CallOption) (*ApplyExtensionResponse, error) {
	out := new(ApplyExtensionResponse)
	err := c.cc.Invoke(ctx, Extension_Apply_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *extensionClient) Delete(ctx context.Context, in *DeleteExtensionRequest, opts ...grpc.CallOption) (*DeleteExtensionResponse, error) {
	out := new(DeleteExtensionResponse)
	err := c.cc.Invoke(ctx, Extension_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionServer is the server API for Extension service.
// All implementations must embed UnimplementedExtensionServer
// for forward compatibility
type ExtensionServer interface {
	List(context.Context, *ListExtensionRequest) (*ListExtensionResponse, error)
	Get(context.Context, *GetExtensionRequest) (*GetExtensionResponse, error)
	Create(context.Context, *CreateExtensionRequest) (*CreateExtensionResponse, error)
	Update(context.Context, *UpdateExtensionRequest) (*UpdateExtensionResponse, error)
	Apply(context.Context, *ApplyExtensionRequest) (*ApplyExtensionResponse, error)
	Delete(context.Context, *DeleteExtensionRequest) (*DeleteExtensionResponse, error)
	mustEmbedUnimplementedExtensionServer()
}

// UnimplementedExtensionServer must be embedded to have forward compatible implementations.
type UnimplementedExtensionServer struct {
}

func (UnimplementedExtensionServer) List(context.Context, *ListExtensionRequest) (*ListExtensionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedExtensionServer) Get(context.Context, *GetExtensionRequest) (*GetExtensionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedExtensionServer) Create(context.Context, *CreateExtensionRequest) (*CreateExtensionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedExtensionServer) Update(context.Context, *UpdateExtensionRequest) (*UpdateExtensionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedExtensionServer) Apply(context.Context, *ApplyExtensionRequest) (*ApplyExtensionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Apply not implemented")
}
func (UnimplementedExtensionServer) Delete(context.Context, *DeleteExtensionRequest) (*DeleteExtensionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedExtensionServer) mustEmbedUnimplementedExtensionServer() {}

// UnsafeExtensionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExtensionServer will
// result in compilation errors.
type UnsafeExtensionServer interface {
	mustEmbedUnimplementedExtensionServer()
}

func RegisterExtensionServer(s grpc.ServiceRegistrar, srv ExtensionServer) {
	s.RegisterService(&Extension_ServiceDesc, srv)
}

func _Extension_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListExtensionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExtensionServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Extension_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExtensionServer).List(ctx, req.(*ListExtensionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Extension_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetExtensionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExtensionServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Extension_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExtensionServer).Get(ctx, req.(*GetExtensionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Extension_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateExtensionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExtensionServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Extension_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExtensionServer).Create(ctx, req.(*CreateExtensionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Extension_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateExtensionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExtensionServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Extension_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExtensionServer).Update(ctx, req.(*UpdateExtensionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Extension_Apply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplyExtensionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExtensionServer).Apply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Extension_Apply_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExtensionServer).Apply(ctx, req.(*ApplyExtensionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Extension_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteExtensionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExtensionServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Extension_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExtensionServer).Delete(ctx, req.(*DeleteExtensionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Extension_ServiceDesc is the grpc.ServiceDesc for Extension service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Extension_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "stub.Extension",
	HandlerType: (*ExtensionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Extension_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Extension_Get_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Extension_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Extension_Update_Handler,
		},
		{
			MethodName: "Apply",
			Handler:    _Extension_Apply_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Extension_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stub/extension.proto",
}
