package grpc

import (
	"context"
	"github.com/tislib/data-handler/pkg/abs"
	"github.com/tislib/data-handler/pkg/stub"
	"github.com/tislib/data-handler/pkg/util"
)

type ExtensionGrpcService interface {
	stub.ExtensionServiceServer
}

type ExtensionServiceServer struct {
	stub.ExtensionServiceServer
	service abs.ExtensionService
}

func (u *ExtensionServiceServer) Create(ctx context.Context, request *stub.CreateExtensionRequest) (*stub.CreateExtensionResponse, error) {
	Extensions, err := u.service.Create(ctx, request.Extensions)

	return &stub.CreateExtensionResponse{
		Extensions: Extensions,
	}, util.ToStatusError(err)
}

func (u *ExtensionServiceServer) Update(ctx context.Context, request *stub.UpdateExtensionRequest) (*stub.UpdateExtensionResponse, error) {
	Extensions, err := u.service.Update(ctx, request.Extensions)

	return &stub.UpdateExtensionResponse{
		Extensions: Extensions,
	}, util.ToStatusError(err)
}

func (u *ExtensionServiceServer) Delete(ctx context.Context, request *stub.DeleteExtensionRequest) (*stub.DeleteExtensionResponse, error) {
	err := u.service.Delete(ctx, request.Ids)

	return &stub.DeleteExtensionResponse{}, util.ToStatusError(err)
}

func (u *ExtensionServiceServer) Get(ctx context.Context, request *stub.GetExtensionRequest) (*stub.GetExtensionResponse, error) {
	Extension, err := u.service.Get(ctx, request.Id)

	return &stub.GetExtensionResponse{
		Extension: Extension,
	}, util.ToStatusError(err)
}

func (u *ExtensionServiceServer) List(ctx context.Context, request *stub.ListExtensionRequest) (*stub.ListExtensionResponse, error) {
	Extensions, err := u.service.List(ctx)

	return &stub.ListExtensionResponse{
		Content: Extensions,
	}, util.ToStatusError(err)
}

func NewExtensionServiceServer(service abs.ExtensionService) stub.ExtensionServiceServer {
	return &ExtensionServiceServer{service: service}
}
