package grpc_service

import (
	"context"
	"data-handler/grpc/stub"
	"data-handler/service"
)

type WorkspaceGrpcService interface {
	stub.WorkspaceServiceServer
}

type WorkspaceServiceServer struct {
	stub.WorkspaceServiceServer
	service service.WorkspaceService
}

func (u *WorkspaceServiceServer) Create(ctx context.Context, request *stub.CreateWorkspaceRequest) (*stub.CreateWorkspaceResponse, error) {
	Workspaces, err := u.service.Create(ctx, request.Workspaces)

	return &stub.CreateWorkspaceResponse{
		Workspaces: Workspaces,
		Error:      toProtoError(err),
	}, nil
}

func (u *WorkspaceServiceServer) Update(ctx context.Context, request *stub.UpdateWorkspaceRequest) (*stub.UpdateWorkspaceResponse, error) {
	Workspaces, err := u.service.Update(ctx, request.Workspaces)

	return &stub.UpdateWorkspaceResponse{
		Workspaces: Workspaces,
		Error:      toProtoError(err),
	}, err
}

func (u *WorkspaceServiceServer) Delete(ctx context.Context, request *stub.DeleteWorkspaceRequest) (*stub.DeleteWorkspaceResponse, error) {
	err := u.service.Delete(ctx, request.Ids)

	return &stub.DeleteWorkspaceResponse{
		Error: toProtoError(err),
	}, nil
}

func (u *WorkspaceServiceServer) Get(ctx context.Context, request *stub.GetWorkspaceRequest) (*stub.GetWorkspaceResponse, error) {
	Workspace, err := u.service.Get(ctx, request.Id)

	return &stub.GetWorkspaceResponse{
		Workspace: Workspace,
		Error:     toProtoError(err),
	}, nil
}

func (u *WorkspaceServiceServer) List(ctx context.Context, request *stub.ListWorkspaceRequest) (*stub.ListWorkspaceResponse, error) {
	Workspaces, err := u.service.List(ctx)

	return &stub.ListWorkspaceResponse{
		Content: Workspaces,
		Error:   toProtoError(err),
	}, err
}

func NewWorkspaceServiceServer(service service.WorkspaceService) stub.WorkspaceServiceServer {
	return &WorkspaceServiceServer{service: service}
}
