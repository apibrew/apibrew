package service

import (
	"context"
	"data-handler/stub"
)

type workSpaceService struct {
	stub.WorkSpaceServiceServer
}

func (receiver workSpaceService) Create(context.Context, *stub.CreateWorkspaceRequest) (*stub.CreateWorkspaceResponse, error) {
	return &stub.CreateWorkspaceResponse{}, nil
}

func NewWorkSpaceService(stub.ResourceServiceServer) stub.WorkSpaceServiceServer {
	return &workSpaceService{}
}
