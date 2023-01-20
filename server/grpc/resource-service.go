package grpc

import (
	"context"
	"data-handler/model"
	"data-handler/server/stub"
	"data-handler/service"
	"data-handler/service/errors"
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
			}, errors.ToStatusError(err)
		}

		result = append(result, res)
	}

	return &stub.CreateResourceResponse{
		Resources: result,
	}, nil
}

func (r resourceGrpcService) Update(ctx context.Context, request *stub.UpdateResourceRequest) (*stub.UpdateResourceResponse, error) {
	for _, resource := range request.Resources {
		err := r.resourceService.Update(ctx, resource, request.DoMigration, request.ForceMigration)

		if err != nil {
			return &stub.UpdateResourceResponse{
				Resources: nil,
			}, errors.ToStatusError(err)
		}
	}

	return &stub.UpdateResourceResponse{
		Resources: request.Resources,
	}, nil
}

func (r resourceGrpcService) Delete(ctx context.Context, request *stub.DeleteResourceRequest) (*stub.DeleteResourceResponse, error) {
	err := r.resourceService.Delete(ctx, request.Ids, request.DoMigration, request.ForceMigration)

	return &stub.DeleteResourceResponse{}, errors.ToStatusError(err)
}

func (r resourceGrpcService) List(ctx context.Context, request *stub.ListResourceRequest) (*stub.ListResourceResponse, error) {
	resources, err := r.resourceService.List(ctx)

	return &stub.ListResourceResponse{
		Resources: resources,
	}, errors.ToStatusError(err)
}

func (r resourceGrpcService) Get(ctx context.Context, request *stub.GetResourceRequest) (*stub.GetResourceResponse, error) {
	resource, err := r.resourceService.Get(ctx, request.Id)

	return &stub.GetResourceResponse{
		Resource: resource,
	}, errors.ToStatusError(err)
}

func (r resourceGrpcService) GetByName(ctx context.Context, request *stub.GetResourceByNameRequest) (*stub.GetResourceByNameResponse, error) {
	resource, err := r.resourceService.GetResourceByName(ctx, request.Namespace, request.Name)

	return &stub.GetResourceByNameResponse{
		Resource: resource,
	}, errors.ToStatusError(err)
}

func (r resourceGrpcService) GetSystemResource(ctx context.Context, request *stub.GetSystemResourceRequest) (*stub.GetSystemResourceResponse, error) {
	resource, err := r.resourceService.GetSystemResourceByName(request.GetName())

	return &stub.GetSystemResourceResponse{
		Resource: resource,
	}, errors.ToStatusError(err)
}

func NewResourceServiceServer(service service.ResourceService) stub.ResourceServiceServer {
	return &resourceGrpcService{resourceService: service}
}
