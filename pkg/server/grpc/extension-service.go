package grpc

import (
	"context"
	"github.com/tislib/apibrew/pkg/abs"
	"github.com/tislib/apibrew/pkg/stub"
	"github.com/tislib/apibrew/pkg/util"
)

type ExtensionGrpcService interface {
	stub.ExtensionServer
}

type ExtensionServer struct {
	stub.ExtensionServer
	service abs.ExtensionService
}

func (u *ExtensionServer) Create(ctx context.Context, request *stub.CreateExtensionRequest) (*stub.CreateExtensionResponse, error) {
	Extensions, err := u.service.Create(ctx, request.Extensions)

	return &stub.CreateExtensionResponse{
		Extensions: Extensions,
	}, util.ToStatusError(err)
}

func (u *ExtensionServer) Update(ctx context.Context, request *stub.UpdateExtensionRequest) (*stub.UpdateExtensionResponse, error) {
	Extensions, err := u.service.Update(ctx, request.Extensions)

	return &stub.UpdateExtensionResponse{
		Extensions: Extensions,
	}, util.ToStatusError(err)
}

func (u *ExtensionServer) Delete(ctx context.Context, request *stub.DeleteExtensionRequest) (*stub.DeleteExtensionResponse, error) {
	err := u.service.Delete(ctx, request.Ids)

	return &stub.DeleteExtensionResponse{}, util.ToStatusError(err)
}

func (u *ExtensionServer) Get(ctx context.Context, request *stub.GetExtensionRequest) (*stub.GetExtensionResponse, error) {
	Extension, err := u.service.Get(ctx, request.Id)

	return &stub.GetExtensionResponse{
		Extension: Extension,
	}, util.ToStatusError(err)
}

func (u *ExtensionServer) List(ctx context.Context, request *stub.ListExtensionRequest) (*stub.ListExtensionResponse, error) {
	Extensions, err := u.service.List(ctx)

	return &stub.ListExtensionResponse{
		Content: Extensions,
	}, util.ToStatusError(err)
}

func NewExtensionServer(service abs.ExtensionService) stub.ExtensionServer {
	return &ExtensionServer{service: service}
}
