package grpc

import (
	"context"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/service"
	"github.com/apibrew/apibrew/pkg/service/annotations"
	"github.com/apibrew/apibrew/pkg/stub"
	"github.com/apibrew/apibrew/pkg/util"
)

type resourceGrpcService struct {
	stub.ResourceServer
	resourceService service.ResourceService
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
	resources, err := r.resourceService.List(annotations.WithContext(ctx, request))

	return &stub.ListResourceResponse{
		Resources: resources,
	}, util.ToStatusError(err)
}

func (r resourceGrpcService) Get(ctx context.Context, request *stub.GetResourceRequest) (*stub.GetResourceResponse, error) {
	resource, err := r.resourceService.Get(annotations.WithContext(ctx, request), request.Id)

	return &stub.GetResourceResponse{
		Resource: resource,
	}, util.ToStatusError(err)
}

func (r resourceGrpcService) GetByName(ctx context.Context, request *stub.GetResourceByNameRequest) (*stub.GetResourceByNameResponse, error) {
	resource, err := r.resourceService.GetResourceByName(annotations.WithContext(ctx, request), request.Namespace, request.Name)

	return &stub.GetResourceByNameResponse{
		Resource: resource,
	}, util.ToStatusError(err)
}

func (r resourceGrpcService) GetSystemResource(ctx context.Context, request *stub.GetSystemResourceRequest) (*stub.GetSystemResourceResponse, error) {
	resource, err := r.resourceService.GetSystemResourceByName(annotations.WithContext(ctx, request), request.GetName())

	return &stub.GetSystemResourceResponse{
		Resource: resource,
	}, util.ToStatusError(err)
}

func NewResourceServer(resourceService service.ResourceService) stub.ResourceServer {
	return &resourceGrpcService{resourceService: resourceService}
}
