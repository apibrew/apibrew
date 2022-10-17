package service

import (
	"context"
	"data-handler/service/backend"
	"data-handler/stub"
	"data-handler/stub/model"
)

type ResourceService interface {
	stub.ResourceServiceServer
	InjectDataSourceService(service DataSourceService)
	InjectAuthenticationService(service AuthenticationService)

	InitResource(resource *model.Resource)
	Init(data *model.InitData)
	InjectPostgresResourceServiceBackend(serviceBackend backend.ResourceServiceBackend)
}

type resourceService struct {
	stub.ResourceServiceServer
	dataSourceService              DataSourceService
	authenticationService          AuthenticationService
	postgresResourceServiceBackend backend.ResourceServiceBackend
}

func (r *resourceService) InjectPostgresResourceServiceBackend(resourceServiceBackend backend.ResourceServiceBackend) {
	r.postgresResourceServiceBackend = resourceServiceBackend
}

func (r *resourceService) Init(data *model.InitData) {
}

func (r *resourceService) InjectDataSourceService(service DataSourceService) {
	r.dataSourceService = service
}

func (r *resourceService) InjectAuthenticationService(service AuthenticationService) {
	r.authenticationService = service
}

func (r *resourceService) InitResource(resource *model.Resource) {
	_, err := r.postgresResourceServiceBackend.AddResource(backend.AddResourceParams{
		Resource:             resource,
		IgnoreIfExists:       true,
		AllowSystemAndStatic: true,
		Migrate:              true,
		ForceMigrate:         false,
	})

	if err != nil {
		panic(err)
	}
}

func (r resourceService) Create(ctx context.Context, request *stub.CreateResourceRequest) (*stub.CreateResourceResponse, error) {
	// validate token
	err := r.authenticationService.validateToken(request.Token)

	if err != nil {
		return nil, err
	}

	var result []*model.Resource

	for _, resource := range request.Resources {

		res, err := r.postgresResourceServiceBackend.AddResource(backend.AddResourceParams{
			Resource:             resource,
			AllowSystemAndStatic: false,
			Migrate:              request.DoMigration,
			ForceMigrate:         request.ForceMigration,
		})

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

func (r resourceService) Update(ctx context.Context, request *stub.UpdateResourceRequest) (*stub.UpdateResourceResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r resourceService) Delete(ctx context.Context, request *stub.DeleteResourceRequest) (*stub.DeleteResourceResponse, error) {
	err := r.postgresResourceServiceBackend.DeleteResources(ctx, request.Ids, request.DoMigration, request.ForceMigration)

	if err != nil {
		return nil, err
	}

	return &stub.DeleteResourceResponse{}, nil
}

func (r resourceService) List(ctx context.Context, request *stub.ListResourceRequest) (*stub.ListResourceResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r resourceService) Get(ctx context.Context, request *stub.GetResourceRequest) (*stub.GetResourceResponse, error) {
	//TODO implement me
	panic("implement me")
}

func NewResourceService() ResourceService {
	return &resourceService{}
}
