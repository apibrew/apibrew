package grpc

import (
	"context"
	"github.com/apibrew/apibrew/pkg/abs"
	"github.com/apibrew/apibrew/pkg/stub"
	"github.com/apibrew/apibrew/pkg/util"
)

type NamespaceGrpcService interface {
	stub.NamespaceServer
}

type NamespaceServer struct {
	stub.NamespaceServer
	service abs.NamespaceService
}

func (u *NamespaceServer) Create(ctx context.Context, request *stub.CreateNamespaceRequest) (*stub.CreateNamespaceResponse, error) {
	Namespaces, err := u.service.Create(ctx, request.Namespaces)

	return &stub.CreateNamespaceResponse{
		Namespaces: Namespaces,
	}, util.ToStatusError(err)
}

func (u *NamespaceServer) Update(ctx context.Context, request *stub.UpdateNamespaceRequest) (*stub.UpdateNamespaceResponse, error) {
	Namespaces, err := u.service.Update(ctx, request.Namespaces)

	return &stub.UpdateNamespaceResponse{
		Namespaces: Namespaces,
	}, util.ToStatusError(err)
}

func (u *NamespaceServer) Delete(ctx context.Context, request *stub.DeleteNamespaceRequest) (*stub.DeleteNamespaceResponse, error) {
	err := u.service.Delete(ctx, request.Ids)

	return &stub.DeleteNamespaceResponse{}, util.ToStatusError(err)
}

func (u *NamespaceServer) Get(ctx context.Context, request *stub.GetNamespaceRequest) (*stub.GetNamespaceResponse, error) {
	Namespace, err := u.service.Get(ctx, request.Id)

	return &stub.GetNamespaceResponse{
		Namespace: Namespace,
	}, util.ToStatusError(err)
}

func (u *NamespaceServer) List(ctx context.Context, request *stub.ListNamespaceRequest) (*stub.ListNamespaceResponse, error) {
	Namespaces, err := u.service.List(ctx)

	return &stub.ListNamespaceResponse{
		Content: Namespaces,
	}, util.ToStatusError(err)
}

func NewNamespaceServer(service abs.NamespaceService) stub.NamespaceServer {
	return &NamespaceServer{service: service}
}
