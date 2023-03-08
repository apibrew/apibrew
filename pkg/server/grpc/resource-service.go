package grpc

import (
	"context"
	"github.com/tislib/data-handler/pkg/abs"
	"github.com/tislib/data-handler/pkg/errors"
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/server/util"
	"github.com/tislib/data-handler/pkg/service/annotations"
	"github.com/tislib/data-handler/pkg/stub"
)

type resourceGrpcService struct {
	stub.ResourceServiceServer
	resourceService abs.ResourceService
}

func (r resourceGrpcService) Create(ctx context.Context, request *stub.CreateResourceRequest) (*stub.CreateResourceResponse, error) {
	var result []*model.Resource

	for _, resource := range request.Resources {
		res, err := r.resourceService.Create(annotations.WithContext(ctx, request), resource, request.DoMigration, request.ForceMigration)

		if err != nil {
			return &stub.CreateResourceResponse{
				Resources: nil,
			}, util.ToStatusError(err)
		}

		result = append(result, res)
	}

	return &stub.CreateResourceResponse{
		Resources: result,
	}, nil
}

func (r resourceGrpcService) Update(ctx context.Context, request *stub.UpdateResourceRequest) (*stub.UpdateResourceResponse, error) {
	for _, resource := range request.Resources {
		err := r.resourceService.Update(annotations.WithContext(ctx, request), resource, request.DoMigration, request.ForceMigration)

		if err != nil {
			return &stub.UpdateResourceResponse{
				Resources: nil,
			}, util.ToStatusError(err)
		}
	}

	return &stub.UpdateResourceResponse{
		Resources: request.Resources,
	}, nil
}

func (r resourceGrpcService) Delete(ctx context.Context, request *stub.DeleteResourceRequest) (*stub.DeleteResourceResponse, error) {
	err := r.resourceService.Delete(annotations.WithContext(ctx, request), request.Ids, request.DoMigration, request.ForceMigration)

	return &stub.DeleteResourceResponse{}, util.ToStatusError(err)
}

func (r resourceGrpcService) List(ctx context.Context, request *stub.ListResourceRequest) (*stub.ListResourceResponse, error) {
	resources := r.resourceService.List(annotations.WithContext(ctx, request))

	return &stub.ListResourceResponse{
		Resources: resources,
	}, nil
}

func (r resourceGrpcService) Get(ctx context.Context, request *stub.GetResourceRequest) (*stub.GetResourceResponse, error) {
	resource := r.resourceService.Get(annotations.WithContext(ctx, request), request.Id)

	var err errors.ServiceError
	if resource == nil {
		err = errors.ResourceNotFoundError
	}

	return &stub.GetResourceResponse{
		Resource: resource,
	}, util.ToStatusError(err)
}

func (r resourceGrpcService) GetByName(ctx context.Context, request *stub.GetResourceByNameRequest) (*stub.GetResourceByNameResponse, error) {
	resource := r.resourceService.GetResourceByName(annotations.WithContext(ctx, request), request.Namespace, request.Name)

	var err errors.ServiceError
	if resource == nil {
		err = errors.ResourceNotFoundError
	}

	return &stub.GetResourceByNameResponse{
		Resource: resource,
	}, util.ToStatusError(err)
}

func (r resourceGrpcService) GetSystemResource(ctx context.Context, request *stub.GetSystemResourceRequest) (*stub.GetSystemResourceResponse, error) {
	resource := r.resourceService.GetSystemResourceByName(annotations.WithContext(ctx, request), request.GetName())

	var err errors.ServiceError
	if resource == nil {
		err = errors.ResourceNotFoundError
	}

	return &stub.GetSystemResourceResponse{
		Resource: resource,
	}, util.ToStatusError(err)
}

func (r resourceGrpcService) PrepareResourceMigrationPlan(ctx context.Context, request *stub.PrepareResourceMigrationPlanRequest) (*stub.PrepareResourceMigrationPlanResponse, error) {
	plans, err := r.resourceService.PrepareResourceMigrationPlan(annotations.WithContext(ctx, request), request.Resources, request.PrepareFromDataSource)

	return &stub.PrepareResourceMigrationPlanResponse{
		Plans: plans,
	}, util.ToStatusError(err)
}

func NewResourceServiceServer(resourceService abs.ResourceService) stub.ResourceServiceServer {
	return &resourceGrpcService{resourceService: resourceService}
}
