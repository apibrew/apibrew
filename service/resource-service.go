package service

import (
	"context"
	"data-handler/service/backend"
	"data-handler/service/errors"
	"data-handler/service/security"
	"data-handler/service/system"
	"data-handler/stub"
	"data-handler/stub/model"
	"github.com/jellydator/ttlcache/v3"
	"time"
)

type ResourceServiceInternal interface {
	InitResource(resource *model.Resource)
	Init(data *model.InitData)
	CheckResourceExists(workspace, name string) (bool, errors.ServiceError)
	GetResourceByName(ctx context.Context, workspace, resource string) (*model.Resource, errors.ServiceError)
}

type ResourceService interface {
	stub.ResourceServiceServer
	ResourceServiceInternal
	InjectDataSourceService(service DataSourceService)
	InjectAuthenticationService(service AuthenticationService)
	InjectPostgresResourceServiceBackend(serviceBackend backend.ResourceServiceBackend)
}

type resourceService struct {
	stub.ResourceServiceServer
	dataSourceService              DataSourceServiceInternal
	authenticationService          AuthenticationServiceInternal
	postgresResourceServiceBackend backend.ResourceServiceBackend
	ServiceName                    string
	cache                          *ttlcache.Cache[string, *model.Resource]
	disableCache                   bool
}

func (r *resourceService) GetResourceByName(ctx context.Context, workspace string, resourceName string) (*model.Resource, errors.ServiceError) {
	if security.IsSystemContext(ctx) && (workspace == system.WorkspaceResource.Name || workspace == "") {
		if resourceName == system.UserResource.Name {
			return system.UserResource, nil
		} else if resourceName == system.DataSourceResource.Name {
			return system.DataSourceResource, nil
		} else if resourceName == system.WorkspaceResource.Name {
			return system.WorkspaceResource, nil
		}
	}

	if workspace == "" {
		workspace = "default"
	}

	if !r.disableCache {
		if r.cache.Get(workspace+"-"+resourceName) != nil {
			return r.cache.Get(workspace + "-" + resourceName).Value(), nil
		}
	}

	resource, err := r.postgresResourceServiceBackend.GetResourceByName(ctx, workspace, resourceName)

	if err != nil {
		return nil, err
	}

	if !r.disableCache {
		r.cache.Set(workspace+"-"+resourceName, resource, ttlcache.DefaultTTL)
	}

	return resource, nil
}

func (r *resourceService) CheckResourceExists(workspace, name string) (bool, errors.ServiceError) {
	if r.cache.Get(name) != nil {
		return true, nil
	}

	resource, err := r.postgresResourceServiceBackend.GetResourceByName(nil, workspace, name)

	if err != nil {
		return false, err
	}

	r.cache.Set(workspace+"-"+name, resource, ttlcache.DefaultTTL)

	return true, nil
}

func (r *resourceService) InjectPostgresResourceServiceBackend(resourceServiceBackend backend.ResourceServiceBackend) {
	r.postgresResourceServiceBackend = resourceServiceBackend
}

func (r *resourceService) Init(data *model.InitData) {
	r.disableCache = data.Config.DisableCache
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
	var err error
	for _, resource := range request.Resources {
		err = r.GetBackend().UpdateResource(ctx, resource, request.DoMigration, request.ForceMigration)

		if err != nil {
			return nil, err
		}

		r.cache.Delete(resource.Name)
	}

	return &stub.UpdateResourceResponse{
		Resources: request.Resources,
		Error:     nil,
	}, nil
}

func (r resourceService) Delete(ctx context.Context, request *stub.DeleteResourceRequest) (*stub.DeleteResourceResponse, error) {
	var err error
	err = r.GetBackend().DeleteResources(ctx, request.Workspace, request.Ids, request.DoMigration, request.ForceMigration)

	if err != nil {
		return nil, err
	}

	for _, id := range request.Ids {
		r.cache.Delete(id)
	}

	return &stub.DeleteResourceResponse{}, nil
}

func (r resourceService) List(ctx context.Context, request *stub.ListResourceRequest) (*stub.ListResourceResponse, error) {
	resources, err := r.GetBackend().ListResources(ctx)

	if err != nil {
		return nil, err
	}

	return &stub.ListResourceResponse{
		Resources: resources,
	}, nil
}

func (r resourceService) Get(ctx context.Context, request *stub.GetResourceRequest) (*stub.GetResourceResponse, error) {
	resource, err := r.GetBackend().GetResourceByName(ctx, request.Workspace, request.Name)

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
