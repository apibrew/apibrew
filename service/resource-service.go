package service

import (
	"context"
	"data-handler/service/backend"
	"data-handler/stub"
	"data-handler/stub/model"
	"github.com/jellydator/ttlcache/v3"
	"time"
)

type ResourceService interface {
	stub.ResourceServiceServer
	InjectDataSourceService(service DataSourceService)
	InjectAuthenticationService(service AuthenticationService)

	InitResource(resource *model.Resource)
	Init(data *model.InitData)
	InjectPostgresResourceServiceBackend(serviceBackend backend.ResourceServiceBackend)
	CheckResourceExists(name string) (bool, error)
	GetResourceByName(resource string) (*model.Resource, error)
}

type resourceService struct {
	stub.ResourceServiceServer
	dataSourceService              DataSourceService
	authenticationService          AuthenticationService
	postgresResourceServiceBackend backend.ResourceServiceBackend
	ServiceName                    string
	cache                          *ttlcache.Cache[string, *model.Resource]
	existsCache                    *ttlcache.Cache[string, bool]
}

func (r *resourceService) GetResourceByName(resourceName string) (*model.Resource, error) {
	if r.cache.Get(resourceName) != nil {
		return r.cache.Get(resourceName).Value(), nil
	}

	resource, err := r.postgresResourceServiceBackend.GetResourceByName(resourceName)

	if err != nil {
		return nil, err
	}

	r.cache.Set(resourceName, resource, ttlcache.DefaultTTL)

	return resource, nil
}

func (r *resourceService) CheckResourceExists(name string) (bool, error) {
	if r.cache.Get(name) != nil {
		return r.cache.Get(name).Value() != nil, nil
	}

	resource, err := r.postgresResourceServiceBackend.GetResourceByName(name)

	if err != nil {
		return false, err
	}

	r.cache.Set(name, resource, ttlcache.DefaultTTL)

	return resource != nil, nil
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
	_, err := r.GetBackend().AddResource(backend.AddResourceParams{
		Resource:       resource,
		IgnoreIfExists: true,
		Migrate:        true,
		ForceMigrate:   false,
	})

	if err != nil {
		panic(err)
	}
}

func (r resourceService) Create(ctx context.Context, request *stub.CreateResourceRequest) (*stub.CreateResourceResponse, error) {
	err := r.authenticationService.Check(CheckParams{
		Ctx:       ctx,
		Token:     request.Token,
		Service:   r.ServiceName,
		Method:    "Create",
		Resources: request.Resources,
	})

	if err != nil {
		return nil, err
	}

	// validate token
	err = r.authenticationService.validateToken(request.Token)

	if err != nil {
		return nil, err
	}

	var result []*model.Resource

	for _, resource := range request.Resources {

		resource.Type = model.DataType_USER

		res, err := r.GetBackend().AddResource(backend.AddResourceParams{
			Resource:     resource,
			Migrate:      request.DoMigration,
			ForceMigrate: request.ForceMigration,
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
	err := r.authenticationService.Check(CheckParams{
		Ctx:       ctx,
		Token:     request.Token,
		Service:   r.ServiceName,
		Method:    "Update",
		Resources: request.Resources,
	})

	if err != nil {
		return nil, err
	}

	for _, resource := range request.Resources {
		err = r.GetBackend().UpdateResource(ctx, resource, request.DoMigration, request.ForceMigration)

		if err != nil {
			return nil, err
		}
	}

	return &stub.UpdateResourceResponse{
		Resources: request.Resources,
		Error:     nil,
	}, nil
}

func (r resourceService) Delete(ctx context.Context, request *stub.DeleteResourceRequest) (*stub.DeleteResourceResponse, error) {
	err := r.authenticationService.Check(CheckParams{
		Ctx:       ctx,
		Token:     request.Token,
		Service:   r.ServiceName,
		Method:    "Delete",
		Resources: request.Ids,
	})

	if err != nil {
		return nil, err
	}

	err = r.GetBackend().DeleteResources(ctx, request.Ids, request.DoMigration, request.ForceMigration)

	if err != nil {
		return nil, err
	}

	return &stub.DeleteResourceResponse{}, nil
}

func (r resourceService) List(ctx context.Context, request *stub.ListResourceRequest) (*stub.ListResourceResponse, error) {
	err := r.authenticationService.Check(CheckParams{
		Ctx:     ctx,
		Token:   request.Token,
		Service: r.ServiceName,
		Method:  "List",
	})

	if err != nil {
		return nil, err
	}

	resources, err := r.GetBackend().ListResources(ctx)

	if err != nil {
		return nil, err
	}

	return &stub.ListResourceResponse{
		Resources: resources,
	}, nil
}

func (r resourceService) Get(ctx context.Context, request *stub.GetResourceRequest) (*stub.GetResourceResponse, error) {
	err := r.authenticationService.Check(CheckParams{
		Ctx:     ctx,
		Token:   request.Token,
		Service: r.ServiceName,
		Method:  "List",
	})

	if err != nil {
		return nil, err
	}

	resource, err := r.GetBackend().GetResourceByName(request.Name)

	if err != nil {
		return nil, err
	}

	return &stub.GetResourceResponse{
		Resource: resource,
		Error:    nil,
	}, nil
}

func (r resourceService) GetBackend() backend.ResourceServiceBackend {
	return r.postgresResourceServiceBackend
}

func NewResourceService() ResourceService {
	return &resourceService{
		ServiceName: "ResourceService",
		cache: ttlcache.New[string, *model.Resource](
			ttlcache.WithTTL[string, *model.Resource](1 * time.Minute),
		),
	}
}
