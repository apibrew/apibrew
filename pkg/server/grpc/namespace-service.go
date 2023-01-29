package grpc

import (
	"context"
	"github.com/tislib/data-handler/pkg/server/util"
	"github.com/tislib/data-handler/pkg/service"
	"github.com/tislib/data-handler/pkg/stub"
)

type NamespaceGrpcService interface {
	stub.NamespaceServiceServer
}

type NamespaceServiceServer struct {
	stub.NamespaceServiceServer
	service service.NamespaceService
}

func (u *NamespaceServiceServer) Create(ctx context.Context, request *stub.CreateNamespaceRequest) (*stub.CreateNamespaceResponse, error) {
	Namespaces, err := u.service.Create(ctx, request.Namespaces)

	return &stub.CreateNamespaceResponse{
		Namespaces: Namespaces,
	}, util.ToStatusError(err)
}

func (u *NamespaceServiceServer) Update(ctx context.Context, request *stub.UpdateNamespaceRequest) (*stub.UpdateNamespaceResponse, error) {
	Namespaces, err := u.service.Update(ctx, request.Namespaces)

	return &stub.UpdateNamespaceResponse{
		Namespaces: Namespaces,
	}, util.ToStatusError(err)
}

func (u *NamespaceServiceServer) Delete(ctx context.Context, request *stub.DeleteNamespaceRequest) (*stub.DeleteNamespaceResponse, error) {
	err := u.service.Delete(ctx, request.Ids)

	return &stub.DeleteNamespaceResponse{}, util.ToStatusError(err)
}

func (u *NamespaceServiceServer) Get(ctx context.Context, request *stub.GetNamespaceRequest) (*stub.GetNamespaceResponse, error) {
	Namespace, err := u.service.Get(ctx, request.Id)

	return &stub.GetNamespaceResponse{
		Namespace: Namespace,
	}, util.ToStatusError(err)
}

func (u *NamespaceServiceServer) List(ctx context.Context, request *stub.ListNamespaceRequest) (*stub.ListNamespaceResponse, error) {
	Namespaces, err := u.service.List(ctx)

	return &stub.ListNamespaceResponse{
		Content: Namespaces,
	}, util.ToStatusError(err)
}

func NewNamespaceServiceServer(service service.NamespaceService) stub.NamespaceServiceServer {
	return &NamespaceServiceServer{service: service}
}
