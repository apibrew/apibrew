package grpc_service

import (
	"context"
	"data-handler/grpc/stub"
	"data-handler/model"
	"data-handler/service"
)

type resourceGrpcService struct {
	stub.ResourceServiceServer
	resourceService service.ResourceService
}

func (r resourceGrpcService) Create(ctx context.Context, request *stub.CreateResourceRequest) (*stub.CreateResourceResponse, error) {
	var result []*model.Resource

	for _, resource := range request.Resources {
		res, err := r.resourceService.Create(ctx, resource, request.DoMigration, request.ForceMigration)

		if err != nil {
			return nil, err
		}

		result = append(result, res)
	}

	return &stub.CreateResourceResponse{
		Resources: result,
		Error:     nil,
	}, nil
}

func (r resourceGrpcService) Update(ctx context.Context, request *stub.UpdateResourceRequest) (*stub.UpdateResourceResponse, error) {
	var err error
	for _, resource := range request.Resources {
		err = r.resourceService.Update(ctx, resource, request.DoMigration, request.ForceMigration)

		if err != nil {
			return nil, err
		}
	}

	return &stub.UpdateResourceResponse{
		Resources: request.Resources,
		Error:     nil,
	}, nil
}

func (r resourceGrpcService) Delete(ctx context.Context, request *stub.DeleteResourceRequest) (*stub.DeleteResourceResponse, error) {
	var err error
	err = r.resourceService.Delete(ctx, request.Workspace, request.Ids, request.DoMigration, request.ForceMigration)

	if err != nil {
		return nil, err
	}

	return &stub.DeleteResourceResponse{}, nil
}

func (r resourceGrpcService) List(ctx context.Context, request *stub.ListResourceRequest) (*stub.ListResourceResponse, error) {
	resources, err := r.resourceService.List(ctx)

	if err != nil {
		return nil, err
	}

	return &stub.ListResourceResponse{
		Resources: resources,
	}, nil
}

func (r resourceGrpcService) Get(ctx context.Context, request *stub.GetResourceRequest) (*stub.GetResourceResponse, error) {
	resource, err := r.resourceService.Get(ctx, request.Workspace, request.Name)

	if err != nil {
		return nil, err
	}

	return &stub.GetResourceResponse{
		Resource: resource,
		Error:    nil,
	}, nil
}

func NewResourceServiceServer(service service.ResourceService) stub.ResourceServiceServer {
	return &resourceGrpcService{resourceService: service}
}
