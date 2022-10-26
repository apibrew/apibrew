package service

import (
	"context"
	"data-handler/model"
	"data-handler/service/backend"
	"data-handler/service/errors"
	"data-handler/service/security"
	"data-handler/service/system"
	"github.com/jellydator/ttlcache/v3"
	"time"
)

type ResourceService interface {
	InitResource(resource *model.Resource)
	Init(data *model.InitData)
	CheckResourceExists(workspace, name string) (bool, errors.ServiceError)
	GetResourceByName(ctx context.Context, workspace, resource string) (*model.Resource, errors.ServiceError)
	InjectDataSourceService(service DataSourceService)
	InjectAuthenticationService(service AuthenticationService)
	InjectPostgresResourceServiceBackend(serviceBackend backend.ResourceServiceBackend)
	Create(ctx context.Context, resource *model.Resource, doMigration bool, forceMigration bool) (*model.Resource, errors.ServiceError)
	Update(ctx context.Context, resource *model.Resource, doMigration bool, forceMigration bool) errors.ServiceError
	Delete(ctx context.Context, workspace string, ids []string, doMigration bool, forceMigration bool) errors.ServiceError
	List(ctx context.Context) ([]*model.Resource, errors.ServiceError)
	Get(ctx context.Context, workspace, id string) (*model.Resource, errors.ServiceError)
}

type resourceService struct {
	dataSourceService              DataSourceService
	authenticationService          AuthenticationService
	postgresResourceServiceBackend backend.ResourceServiceBackend
	cache                          *ttlcache.Cache[string, *model.Resource]
	disableCache                   bool
}

func (r *resourceService) Update(ctx context.Context, resource *model.Resource, doMigration bool, forceMigration bool) errors.ServiceError {
	r.cache.Delete(resource.Workspace + "-" + resource.Name)
	return r.GetBackend().UpdateResource(ctx, resource, doMigration, forceMigration)
}

func (r *resourceService) Create(ctx context.Context, resource *model.Resource, doMigration bool, forceMigration bool) (*model.Resource, errors.ServiceError) {
	resource.Type = model.DataType_USER

	return r.GetBackend().AddResource(backend.AddResourceParams{
		Resource:       resource,
		IgnoreIfExists: false,
		Migrate:        doMigration,
		ForceMigrate:   forceMigration,
	})
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

func (r resourceService) Delete(ctx context.Context, workspace string, ids []string, doMigration bool, forceMigration bool) errors.ServiceError {
	err := r.GetBackend().DeleteResources(ctx, workspace, ids, doMigration, forceMigration)

	if err != nil {
		return err
	}

	for _, id := range ids {
		r.cache.Delete(id)
	}

	return nil
}

func (r resourceService) List(ctx context.Context) ([]*model.Resource, errors.ServiceError) {
	return r.GetBackend().ListResources(ctx)
}

func (r resourceService) Get(ctx context.Context, workspace, resourceName string) (*model.Resource, errors.ServiceError) {
	return r.GetBackend().GetResourceByName(ctx, workspace, resourceName)
}

func (r resourceService) GetBackend() backend.ResourceServiceBackend {
	return r.postgresResourceServiceBackend
}

func NewResourceService() ResourceService {
	return &resourceService{
		cache: ttlcache.New[string, *model.Resource](
			ttlcache.WithTTL[string, *model.Resource](1 * time.Minute),
		),
	}
}
