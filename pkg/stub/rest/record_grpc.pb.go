// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: stub/rest/record.proto

package rest

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

// RecordClient is the client API for Record service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RecordClient interface {
	Create(ctx context.Context, in *CreateRecordRequest, opts ...grpc.CallOption) (*CreateRecordResponse, error)
	Update(ctx context.Context, in *UpdateRecordRequest, opts ...grpc.CallOption) (*UpdateRecordResponse, error)
	Apply(ctx context.Context, in *ApplyRecordRequest, opts ...grpc.CallOption) (*ApplyRecordResponse, error)
	Delete(ctx context.Context, in *DeleteRecordRequest, opts ...grpc.CallOption) (*DeleteRecordResponse, error)
}

type recordClient struct {
	cc grpc.ClientConnInterface
}

func NewRecordClient(cc grpc.ClientConnInterface) RecordClient {
	return &recordClient{cc}
}

func (c *recordClient) Create(ctx context.Context, in *CreateRecordRequest, opts ...grpc.CallOption) (*CreateRecordResponse, error) {
	out := new(CreateRecordResponse)
	err := c.cc.Invoke(ctx, "/rest.Record/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recordClient) Update(ctx context.Context, in *UpdateRecordRequest, opts ...grpc.CallOption) (*UpdateRecordResponse, error) {
	out := new(UpdateRecordResponse)
	err := c.cc.Invoke(ctx, "/rest.Record/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recordClient) Apply(ctx context.Context, in *ApplyRecordRequest, opts ...grpc.CallOption) (*ApplyRecordResponse, error) {
	out := new(ApplyRecordResponse)
	err := c.cc.Invoke(ctx, "/rest.Record/Apply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recordClient) Delete(ctx context.Context, in *DeleteRecordRequest, opts ...grpc.CallOption) (*DeleteRecordResponse, error) {
	out := new(DeleteRecordResponse)
	err := c.cc.Invoke(ctx, "/rest.Record/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RecordServer is the server API for Record service.
// All implementations must embed UnimplementedRecordServer
// for forward compatibility
type RecordServer interface {
	Create(context.Context, *CreateRecordRequest) (*CreateRecordResponse, error)
	Update(context.Context, *UpdateRecordRequest) (*UpdateRecordResponse, error)
	Apply(context.Context, *ApplyRecordRequest) (*ApplyRecordResponse, error)
	Delete(context.Context, *DeleteRecordRequest) (*DeleteRecordResponse, error)
	mustEmbedUnimplementedRecordServer()
}

// UnimplementedRecordServer must be embedded to have forward compatible implementations.
type UnimplementedRecordServer struct {
}

func (UnimplementedRecordServer) Create(context.Context, *CreateRecordRequest) (*CreateRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedRecordServer) Update(context.Context, *UpdateRecordRequest) (*UpdateRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedRecordServer) Apply(context.Context, *ApplyRecordRequest) (*ApplyRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Apply not implemented")
}
func (UnimplementedRecordServer) Delete(context.Context, *DeleteRecordRequest) (*DeleteRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedRecordServer) mustEmbedUnimplementedRecordServer() {}

// UnsafeRecordServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RecordServer will
// result in compilation errors.
type UnsafeRecordServer interface {
	mustEmbedUnimplementedRecordServer()
}

func RegisterRecordServer(s grpc.ServiceRegistrar, srv RecordServer) {
	s.RegisterService(&Record_ServiceDesc, srv)
}

func _Record_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecordServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rest.Record/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecordServer).Create(ctx, req.(*CreateRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Record_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecordServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rest.Record/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecordServer).Update(ctx, req.(*UpdateRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Record_Apply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplyRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecordServer).Apply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rest.Record/Apply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecordServer).Apply(ctx, req.(*ApplyRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Record_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecordServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rest.Record/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecordServer).Delete(ctx, req.(*DeleteRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Record_ServiceDesc is the grpc.ServiceDesc for Record service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Record_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rest.Record",
	HandlerType: (*RecordServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Record_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Record_Update_Handler,
		},
		{
			MethodName: "Apply",
			Handler:    _Record_Apply_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Record_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stub/rest/record.proto",
}
