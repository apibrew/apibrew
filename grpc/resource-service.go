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
			return &stub.CreateResourceResponse{
				Resources: nil,
				Error:     toProtoError(err),
			}, nil
		}

		result = append(result, res)
	}

	return &stub.CreateResourceResponse{
		Resources: result,
		Error:     nil,
	}, nil
}

func (r resourceGrpcService) Update(ctx context.Context, request *stub.UpdateResourceRequest) (*stub.UpdateResourceResponse, error) {
	for _, resource := range request.Resources {
		err := r.resourceService.Update(ctx, resource, request.DoMigration, request.ForceMigration)

		if err != nil {
			return &stub.UpdateResourceResponse{
				Resources: nil,
				Error:     toProtoError(err),
			}, nil
		}
	}

	return &stub.UpdateResourceResponse{
		Resources: request.Resources,
		Error:     nil,
	}, nil
}

func (r resourceGrpcService) Delete(ctx context.Context, request *stub.DeleteResourceRequest) (*stub.DeleteResourceResponse, error) {
	err := r.resourceService.Delete(ctx, request.Workspace, request.Ids, request.DoMigration, request.ForceMigration)

	return &stub.DeleteResourceResponse{
		Error: toProtoError(err),
	}, nil
}

func (r resourceGrpcService) List(ctx context.Context, request *stub.ListResourceRequest) (*stub.ListResourceResponse, error) {
	resources, err := r.resourceService.List(ctx)

	return &stub.ListResourceResponse{
		Resources: resources,
		Error:     toProtoError(err),
	}, nil
}

func (r resourceGrpcService) Get(ctx context.Context, request *stub.GetResourceRequest) (*stub.GetResourceResponse, error) {
	resource, err := r.resourceService.Get(ctx, request.Workspace, request.Id)

	return &stub.GetResourceResponse{
		Resource: resource,
		Error:    toProtoError(err),
	}, nil
}

func NewResourceServiceServer(service service.ResourceService) stub.ResourceServiceServer {
	return &resourceGrpcService{resourceService: service}
}
