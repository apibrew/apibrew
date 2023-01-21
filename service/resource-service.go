package service

import (
	"context"
	"data-handler/logging"
	"data-handler/model"
	"data-handler/service/backend"
	"data-handler/service/errors"
	"data-handler/service/mapping"
	"data-handler/service/system"
	"fmt"
	"github.com/jellydator/ttlcache/v3"
	log "github.com/sirupsen/logrus"
	"time"
)

type ResourceService interface {
	MigrateResource(resource *model.Resource)
	Init(data *model.InitData)
	CheckResourceExists(ctx context.Context, namespace, name string) (bool, errors.ServiceError)
	GetResourceByName(ctx context.Context, namespace, resource string) (*model.Resource, errors.ServiceError)
	GetSystemResourceByName(resourceName string) (*model.Resource, errors.ServiceError)
	InjectBackendProviderService(backendProviderService BackendProviderService)
	Create(ctx context.Context, resource *model.Resource, doMigration bool, forceMigration bool) (*model.Resource, errors.ServiceError)
	Update(ctx context.Context, resource *model.Resource, doMigration bool, forceMigration bool) errors.ServiceError
	Delete(ctx context.Context, ids []string, doMigration bool, forceMigration bool) errors.ServiceError
	List(ctx context.Context) ([]*model.Resource, errors.ServiceError)
	Get(ctx context.Context, id string) (*model.Resource, errors.ServiceError)
}

type resourceService struct {
	cache                  *ttlcache.Cache[string, *model.Resource]
	disableCache           bool
	backendProviderService BackendProviderService
}

func (r *resourceService) InjectBackendProviderService(backendProviderService BackendProviderService) {
	r.backendProviderService = backendProviderService
}

func (r *resourceService) Update(ctx context.Context, resource *model.Resource, doMigration bool, forceMigration bool) errors.ServiceError {
	resource.DataType = model.DataType_USER

	if err := validateResource(resource); err != nil {
		return err
	}

	if err := r.mustModifyResource(resource); err != nil {
		return err
	}

	result, _, err := r.backendProviderService.GetSystemBackend(ctx).AddRecords(ctx, backend.BulkRecordsParams{
		Resource:       system.ResourceResource,
		Records:        []*model.Record{mapping.ResourceToRecord(resource)},
		CheckVersion:   false,
		IgnoreIfExists: false,
	})

	if err != nil && err.Code() == model.ErrorCode_UNIQUE_VIOLATION {
		return errors.AlreadyExistsError.WithMessage(fmt.Sprintf("Resource is already exiss: " + resource.Name))
	}

	if err != nil {
		return err
	}

	resource.Id = result[0].Id

	propertyRecords := mapping.MapToRecord(resource.Properties, func(property *model.ResourceProperty) *model.Record {
		return mapping.ResourcePropertyToRecord(property, resource)
	})

	_, err = r.backendProviderService.GetSystemBackend(ctx).UpdateRecords(ctx, backend.BulkRecordsParams{
		Resource:       system.ResourcePropertyResource,
		Records:        propertyRecords,
		CheckVersion:   false,
		IgnoreIfExists: false,
	})

	if err != nil {
		return err
	}

	//todo add references

	if doMigration {
		bck, err := r.backendProviderService.GetBackendByDataSourceId(ctx, resource.SourceConfig.DataSource)

		if err != nil {
			return err
		}

		err = bck.UpgradeResource(ctx, nil, resource, forceMigration)

		if err != nil {
			return err
		}
	}

	return nil
}

func (r *resourceService) Create(ctx context.Context, resource *model.Resource, doMigration bool, forceMigration bool) (*model.Resource, errors.ServiceError) {
	resource.DataType = model.DataType_USER

	if err := validateResource(resource); err != nil {
		return nil, err
	}

	if err := r.mustModifyResource(resource); err != nil {
		return nil, err
	}

	result, _, err := r.backendProviderService.GetSystemBackend(ctx).AddRecords(ctx, backend.BulkRecordsParams{
		Resource:       system.ResourceResource,
		Records:        []*model.Record{mapping.ResourceToRecord(resource)},
		CheckVersion:   false,
		IgnoreIfExists: false,
	})

	if err != nil && err.Code() == model.ErrorCode_UNIQUE_VIOLATION {
		return nil, errors.AlreadyExistsError.WithMessage(fmt.Sprintf("Resource is already exiss: " + resource.Name))
	}

	if err != nil {
		return nil, err
	}

	resource.Id = result[0].Id

	propertyRecords := mapping.MapToRecord(resource.Properties, func(property *model.ResourceProperty) *model.Record {
		return mapping.ResourcePropertyToRecord(property, resource)
	})

	_, _, err = r.backendProviderService.GetSystemBackend(ctx).AddRecords(ctx, backend.BulkRecordsParams{
		Resource:       system.ResourcePropertyResource,
		Records:        propertyRecords,
		CheckVersion:   false,
		IgnoreIfExists: false,
	})

	if err != nil {
		_ = r.Delete(ctx, []string{resource.Id}, false, false)
		return nil, err
	}

	//todo add references

	if doMigration {
		bck, err := r.backendProviderService.GetBackendByDataSourceId(ctx, resource.SourceConfig.DataSource)

		if err != nil {
			return nil, err
		}

		err = bck.UpgradeResource(ctx, nil, resource, forceMigration)

		if err != nil {
			return nil, err
		}
	}

	return resource, nil
}

func validateResource(resource *model.Resource) errors.ServiceError {
	if resource.SourceConfig == nil {
		return errors.RecordValidationError.WithDetails("resource source-config is null")
	}

	if resource.Flags == nil {
		return errors.RecordValidationError.WithDetails("resource flags has not initialized")
	}

	if resource.SourceConfig == nil {
		return errors.RecordValidationError.WithDetails("resource SourceConfig has not initialized")
	}

	return nil
}

func (r *resourceService) GetResourceByName(ctx context.Context, namespace string, resourceName string) (*model.Resource, errors.ServiceError) {
	logger := log.WithFields(logging.CtxFields(ctx))

	logger.Debugf("Begin resource-service GetResourceByName: %s / %s", namespace, resourceName)
	defer logger.Debug("End resource-service GetResourceByName")

	if namespace == "system" {
		logger.Debugf("Call GetSystemResourceByName: %s", resourceName)

		return r.GetSystemResourceByName(resourceName)
	}

	if namespace == "" {
		namespace = "default"
	}

	if !r.disableCache {
		if r.cache.Get(namespace+"-"+resourceName) != nil {
			return r.cache.Get(namespace + "-" + resourceName).Value(), nil
		}
	}

	logger.Debugf("Call backend GetResourceByName: %s", resourceName)

	queryMap := make(map[string]interface{})

	queryMap["name"] = resourceName
	queryMap["namespace"] = namespace

	logger.Debug("Call PrepareQuery: ", queryMap)
	query, err := PrepareQuery(system.ResourceResource, queryMap)
	logger.Debug("Result record-service: ", query)

	if err != nil {
		return nil, err
	}

	records, _, err := r.backendProviderService.GetSystemBackend(ctx).ListRecords(ctx, backend.ListRecordParams{
		Resource: system.ResourceResource,
		Query:    query,
		Limit:    1,
	})

	if err != nil {
		return nil, err
	}

	if len(records) == 0 {
		return nil, errors.ResourceNotFoundError.WithDetails(fmt.Sprintf("%s/%s", namespace, resourceName))
	}

	resource := mapping.ResourceFromRecord(records[0])

	err = r.loadResource(ctx, resource)

	if err != nil {
		logger.Error("Error GetResourceByName: %s", resourceName)
		return nil, err
	}

	if !r.disableCache {
		r.cache.Set(namespace+"-"+resourceName, resource, ttlcache.DefaultTTL)
	}

	return resource, nil
}

func (r *resourceService) GetSystemResourceByName(resourceName string) (*model.Resource, errors.ServiceError) {
	if resourceName == system.UserResource.Name {
		return system.UserResource, nil
	} else if resourceName == system.DataSourceResource.Name {
		return system.DataSourceResource, nil
	} else if resourceName == system.NamespaceResource.Name {
		return system.NamespaceResource, nil
	}

	return nil, errors.ResourceNotFoundError.WithDetails(fmt.Sprintf("system/%s", resourceName))
}

func (r *resourceService) CheckResourceExists(ctx context.Context, namespace, name string) (bool, errors.ServiceError) {
	if r.cache.Get(name) != nil {
		return true, nil
	}

	resource, err := r.GetResourceByName(ctx, namespace, name)

	if err != nil {
		return false, err
	}

	r.cache.Set(namespace+"-"+name, resource, ttlcache.DefaultTTL)

	return true, nil
}

func (r *resourceService) Init(data *model.InitData) {
	r.disableCache = data.Config.DisableCache

	r.MigrateResource(system.ResourceResource)
	r.MigrateResource(system.ResourcePropertyResource)
	r.MigrateResource(system.ResourceReferenceResource)
}

func (r *resourceService) MigrateResource(resource *model.Resource) {
	err := r.backendProviderService.GetSystemBackend(context.TODO()).UpgradeResource(context.TODO(), nil, resource, true)

	if err != nil {
		panic(err)
	}
}

func (r resourceService) Delete(ctx context.Context, ids []string, doMigration bool, forceMigration bool) errors.ServiceError {
	for _, resourceId := range ids {
		resource, err := r.Get(ctx, resourceId)

		if err != nil {
			return err
		}

		if err := r.mustModifyResource(resource); err != nil {
			return err
		}

		err = r.backendProviderService.GetSystemBackend(ctx).DeleteRecords(ctx, system.ResourceResource, []string{resourceId})

		if err != nil {
			return err
		}

		//@todo cascade delete resource properties and references from backend

		r.cache.Delete(resourceId)

		if doMigration {
			bck, err := r.backendProviderService.GetBackendByDataSourceId(ctx, resource.SourceConfig.DataSource)

			if err != nil {
				return err
			}
			err = bck.DowngradeResource(ctx, resource, forceMigration)

			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (r resourceService) mustModifyResource(resource *model.Resource) errors.ServiceError {
	if resource.Namespace == "system" || resource.DataType == model.DataType_SYSTEM {
		return errors.LogicalError.WithMessage("actions on system resource is not allowed")
	}

	if resource.DataType == model.DataType_STATIC {
		return errors.LogicalError.WithMessage("actions on static resource is not allowed")
	}
	return nil
}

func (r resourceService) List(ctx context.Context) ([]*model.Resource, errors.ServiceError) {
	list, _, err := r.backendProviderService.GetSystemBackend(ctx).ListRecords(ctx, backend.ListRecordParams{
		Resource: system.ResourceResource,
	})

	resources := mapping.MapFromRecord(list, mapping.ResourceFromRecord)

	if err != nil {
		return nil, err
	}

	var result []*model.Resource

	for _, item := range resources {
		if item.DataType == model.DataType_SYSTEM {
			continue
		}

		if item.DataType == model.DataType_STATIC {
			continue
		}

		err = r.loadResource(ctx, item)

		if err != nil {
			return nil, err
		}

		result = append(result, item)
	}

	return result, nil
}

func (r resourceService) Get(ctx context.Context, id string) (*model.Resource, errors.ServiceError) {
	record, err := r.backendProviderService.GetSystemBackend(ctx).GetRecord(ctx, system.ResourceResource, id)

	if err != nil {
		return nil, err
	}

	resource := mapping.ResourceFromRecord(record)

	err = r.loadResource(ctx, resource)

	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (r resourceService) loadResource(ctx context.Context, resource *model.Resource) errors.ServiceError {
	if resource.Properties == nil {
		queryMap := make(map[string]interface{})

		queryMap["resource"] = resource.Id

		query, err := PrepareQuery(system.ResourcePropertyResource, queryMap)

		if err != nil {
			return err
		}

		list, _, err := r.backendProviderService.GetSystemBackend(ctx).ListRecords(ctx, backend.ListRecordParams{
			Resource: system.ResourcePropertyResource,
			Query:    query,
		})

		if err != nil {
			return err
		}

		resource.Properties = mapping.MapFromRecord(list, mapping.ResourcePropertyFromRecord)
	}

	return nil
}

func NewResourceService() ResourceService {
	return &resourceService{
		cache: ttlcache.New[string, *model.Resource](
			ttlcache.WithTTL[string, *model.Resource](1 * time.Minute),
		),
	}
}
