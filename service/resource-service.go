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
}

type resourceService struct {
	stub.ResourceServiceServer
	dataSourceService              DataSourceService
	postgresResourceServiceBackend backend.ResourceServiceBackend
	authenticationService          AuthenticationService
}

func (r *resourceService) Init(data *model.InitData) {
	b, err := r.dataSourceService.GetDataSourceBackend("system")

	if err != nil {
		panic(err)
	}

	r.postgresResourceServiceBackend.Init(b)
}

func (r *resourceService) InjectDataSourceService(service DataSourceService) {
	r.dataSourceService = service
}

func (r *resourceService) InjectAuthenticationService(service AuthenticationService) {
	r.authenticationService = service
}

func (r *resourceService) InitResource(resource *model.Resource) {
	b, err := r.dataSourceService.GetDataSourceBackend(resource.SourceConfig.DataSource)

	if err != nil {
		panic(err)
	}

	_, err = r.postgresResourceServiceBackend.AddResource(backend.AddResourceParams{
		Backend:              b,
		Resource:             resource,
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
		b, err := r.dataSourceService.GetDataSourceBackend(resource.SourceConfig.DataSource)

		if err != nil {
			panic(err)
		}

		res, err := r.postgresResourceServiceBackend.AddResource(backend.AddResourceParams{
			Backend:              b,
			Resource:             resource,
			AllowSystemAndStatic: false,
			Migrate:              request.DoMigration,
			ForceMigrate:         request.ForceMigration,
		})

		if err != nil {
			panic(err)
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
	//TODO implement me
	panic("implement me")
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
	return &resourceService{
		postgresResourceServiceBackend: backend.NewPostgresResourceServiceBackend(),
	}
}
