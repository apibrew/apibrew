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
	GetSystemResourceByName(resourceName string) (*model.Resource, errors.ServiceError)
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

	if err := r.checkSystemResource(ctx, resource.Workspace, resource.Id); err != nil {
		return err
	}

	return r.GetBackend().UpdateResource(ctx, resource, doMigration, forceMigration)
}

func (r *resourceService) checkSystemResource(ctx context.Context, workspace, id string) errors.ServiceError {
	if !security.IsSystemContext(ctx) {
		if workspace == "system" {
			return errors.LogicalError.WithMessage("you cannot access system workspace resources")
		}

		res, err := r.GetBackend().GetResource(ctx, workspace, id)

		if err != nil {
			return err
		}

		if res.Type == model.DataType_SYSTEM {
			return errors.LogicalError.WithMessage("you cannot access system workspace resources")
		}

		if res.Type == model.DataType_STATIC {
			return errors.LogicalError.WithMessage("static resources are not editable")
		}
	}
	return nil
}

func (r *resourceService) Create(ctx context.Context, resource *model.Resource, doMigration bool, forceMigration bool) (*model.Resource, errors.ServiceError) {
	resource.Type = model.DataType_USER

	err := validateResource(resource)

	if err != nil {
		return nil, err
	}

	if !security.IsSystemContext(ctx) {
		if resource.Workspace == "system" {
			return nil, errors.LogicalError.WithMessage("you cannot update system workspace resources")
		}
	}

	return r.GetBackend().AddResource(backend.AddResourceParams{
		Resource:       resource,
		IgnoreIfExists: false,
		Migrate:        doMigration,
		ForceMigrate:   forceMigration,
	})
}

func validateResource(resource *model.Resource) errors.ServiceError {
	if resource.SourceConfig == nil {
		return errors.RecordValidationError.WithDetails("resource source-config is null")
	}

	return nil
}

func (r *resourceService) GetResourceByName(ctx context.Context, workspace string, resourceName string) (*model.Resource, errors.ServiceError) {
	if security.IsSystemContext(ctx) && (workspace == system.WorkspaceResource.Name || workspace == "") {
		resource, err := r.GetSystemResourceByName(resourceName)
		if err != nil {
			return resource, err
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

func (r *resourceService) GetSystemResourceByName(resourceName string) (*model.Resource, errors.ServiceError) {
	if resourceName == system.UserResource.Name {
		return system.UserResource, nil
	} else if resourceName == system.DataSourceResource.Name {
		return system.DataSourceResource, nil
	} else if resourceName == system.WorkspaceResource.Name {
		return system.WorkspaceResource, nil
	}
	return nil, errors.NotFoundError
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
	for _, id := range ids {
		if err := r.checkSystemResource(ctx, workspace, id); err != nil {
			return err
		}
	}

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
	list, err := r.GetBackend().ListResources(ctx)

	if security.IsSystemContext(ctx) {
		return list, err
	}

	if err != nil {
		return nil, err
	}

	var result []*model.Resource

	for _, item := range list {
		if item.Type == model.DataType_SYSTEM {
			continue
		}

		if item.Type == model.DataType_STATIC {
			continue
		}

		result = append(result, item)
	}

	return result, nil
}

func (r resourceService) Get(ctx context.Context, workspace, id string) (*model.Resource, errors.ServiceError) {
	return r.GetBackend().GetResource(ctx, workspace, id)
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
