package service

import (
	"context"
	"fmt"
	"github.com/jellydator/ttlcache/v3"
	log "github.com/sirupsen/logrus"
	"github.com/tislib/data-handler/pkg/abs"
	"github.com/tislib/data-handler/pkg/errors"
	"github.com/tislib/data-handler/pkg/logging"
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/resources"
	mapping2 "github.com/tislib/data-handler/pkg/resources/mapping"
	"github.com/tislib/data-handler/pkg/server/util"
	"strconv"
	"strings"
	"time"
)

type resourceService struct {
	cache                  *ttlcache.Cache[string, *model.Resource]
	disableCache           bool
	backendProviderService abs.BackendProviderService
}

func (r *resourceService) Update(ctx context.Context, resource *model.Resource, doMigration bool, forceMigration bool) errors.ServiceError {
	resource.DataType = model.DataType_USER

	if err := validateResource(resource); err != nil {
		return err
	}

	if err := r.mustModifyResource(resource); err != nil {
		return err
	}

	resourceRecords := []*model.Record{mapping2.ResourceToRecord(resource)}

	if err := checkAccess(ctx, checkAccessParams{
		Resource:  resource,
		Records:   &resourceRecords,
		Operation: model.OperationType_OPERATION_TYPE_UPDATE,
	}); err != nil {
		return err
	}

	result, _, err := r.backendProviderService.GetSystemBackend(ctx).AddRecords(ctx, abs.BulkRecordsParams{
		Resource:       resources.ResourceResource,
		Records:        resourceRecords,
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

	propertyRecords := mapping2.MapToRecord(resource.Properties, func(property *model.ResourceProperty) *model.Record {
		return mapping2.ResourcePropertyToRecord(property, resource)
	})

	_, err = r.backendProviderService.GetSystemBackend(ctx).UpdateRecords(ctx, abs.BulkRecordsParams{
		Resource:       resources.ResourcePropertyResource,
		Records:        propertyRecords,
		CheckVersion:   false,
		IgnoreIfExists: false,
	})

	if err != nil {
		return err
	}

	//todo add references

	if !resource.Virtual && doMigration {
		bck, err := r.backendProviderService.GetBackendByDataSourceId(ctx, resource.SourceConfig.DataSource)

		if err != nil {
			return err
		}

		err = bck.UpgradeResource(ctx, abs.UpgradeResourceParams{
			Resource:       resource,
			ForceMigration: forceMigration,
		})

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

	systemBackend := r.backendProviderService.GetSystemBackend(ctx)

	txk, err := systemBackend.BeginTransaction(ctx, false)

	if err != nil {
		return nil, err
	}

	txCtx := context.WithValue(ctx, "transactionKey", txk)

	var success = false

	defer func() {
		if success {
			err = systemBackend.CommitTransaction(txCtx)

			if err != nil {
				log.Print(err)
			}
		} else {
			err = systemBackend.RollbackTransaction(txCtx)

			if err != nil {
				log.Print(err)
			}
		}
	}()

	result, _, err := systemBackend.AddRecords(txCtx, abs.BulkRecordsParams{
		Resource:       resources.ResourceResource,
		Records:        []*model.Record{mapping2.ResourceToRecord(resource)},
		CheckVersion:   false,
		IgnoreIfExists: false,
	})

	if err != nil && err.Code() == model.ErrorCode_UNIQUE_VIOLATION {
		return nil, errors.AlreadyExistsError.WithMessage(fmt.Sprintf("Resource is already exiss: " + resource.Name))
	}

	if err != nil && err.Code() == model.ErrorCode_RECORD_VALIDATION_ERROR {
		return nil, errors.ResourceValidationError.WithMessage(err.Error()).WithDetails(err.GetDetails()).WithErrorFields(util.GetErrorFields(err))
	}

	if err != nil && err.Code() == model.ErrorCode_REFERENCE_VIOLATION {
		return nil, errors.ResourceValidationError.WithMessage(err.Error()).WithDetails(err.GetDetails()).WithErrorFields(util.GetErrorFields(err))
	}

	if err != nil {
		return nil, err
	}

	resource.Id = result[0].Id

	if len(resource.Properties) > 0 {

		propertyRecords := mapping2.MapToRecord(resource.Properties, func(property *model.ResourceProperty) *model.Record {
			return mapping2.ResourcePropertyToRecord(property, resource)
		})

		_, _, err = systemBackend.AddRecords(txCtx, abs.BulkRecordsParams{
			Resource:       resources.ResourcePropertyResource,
			Records:        propertyRecords,
			CheckVersion:   false,
			IgnoreIfExists: false,
		})

		if err != nil {
			return nil, err
		}
	}

	if len(resource.References) > 0 {
		referenceRecords := mapping2.MapToRecord(resource.References, func(property *model.ResourceReference) *model.Record {
			return mapping2.ResourceReferenceToRecord(property, resource)
		})

		_, _, err = systemBackend.AddRecords(txCtx, abs.BulkRecordsParams{
			Resource:       resources.ResourceReferenceResource,
			Records:        referenceRecords,
			CheckVersion:   false,
			IgnoreIfExists: false,
		})

		if err != nil {
			return nil, err
		}
	}

	if !resource.Virtual && doMigration {
		bck, err := r.backendProviderService.GetBackendByDataSourceId(ctx, resource.SourceConfig.DataSource)

		if err != nil {
			return nil, err
		}

		err = bck.UpgradeResource(ctx, abs.UpgradeResourceParams{
			Resource:       resource,
			ForceMigration: forceMigration,
		})

		if err != nil {
			return nil, err
		}
	}

	success = true

	return resource, nil
}

func validateResource(resource *model.Resource) errors.ServiceError {
	var errorFields []*model.ErrorField

	if resource.Name == "" {
		errorFields = append(errorFields, &model.ErrorField{
			RecordId: resource.Id,
			Property: "Name",
			Message:  "should not be empty",
			Value:    nil,
		})
	}

	if !resource.Virtual {
		if resource.SourceConfig == nil {
			errorFields = append(errorFields, &model.ErrorField{
				RecordId: resource.Id,
				Property: "SourceConfig",
				Message:  "should not be nil",
				Value:    nil,
			})
		} else {
			if resource.SourceConfig.DataSource == "" {
				errorFields = append(errorFields, &model.ErrorField{
					RecordId: resource.Id,
					Property: "SourceConfig.DataSource",
					Message:  "should not be blank",
					Value:    nil,
				})
			}

			if resource.SourceConfig.Entity == "" {
				errorFields = append(errorFields, &model.ErrorField{
					RecordId: resource.Id,
					Property: "SourceConfig.Entity",
					Message:  "should not be blank",
					Value:    nil,
				})
			}
		}
	}

	for i, prop := range resource.Properties {
		propertyPrefix := "Properties[" + strconv.Itoa(i) + "]."

		if prop.Name == "" {
			errorFields = append(errorFields, &model.ErrorField{
				RecordId: resource.Id,
				Property: propertyPrefix + "Name",
				Message:  "should not be blank",
				Value:    nil,
			})
		}

		if prop.SourceConfig == nil {
			errorFields = append(errorFields, &model.ErrorField{
				RecordId: resource.Id,
				Property: propertyPrefix + "SourceConfig",
				Message:  "should not be nil",
				Value:    nil,
			})
			continue
		}

		if _, ok := prop.SourceConfig.(*model.ResourceProperty_Computed); ok {
			errorFields = append(errorFields, &model.ErrorField{
				RecordId: resource.Id,
				Property: propertyPrefix + "SourceConfig",
				Message:  "computed property source type is not supported",
				Value:    nil,
			})
			continue
		}

		mp := prop.SourceConfig.(*model.ResourceProperty_Mapping)

		if mp.Mapping == nil {
			errorFields = append(errorFields, &model.ErrorField{
				RecordId: resource.Id,
				Property: propertyPrefix + "SourceConfig.Mapping",
				Message:  "Mapping not be nil",
				Value:    nil,
			})
		} else {
			if mp.Mapping.Mapping == "" {
				errorFields = append(errorFields, &model.ErrorField{
					RecordId: resource.Id,
					Property: propertyPrefix + "SourceConfig.Mapping",
					Message:  "Mapping should not be blank",
					Value:    nil,
				})
			}
		}

		if prop.Type == model.ResourcePropertyType_TYPE_STRING {
			if prop.Length <= 0 {
				errorFields = append(errorFields, &model.ErrorField{
					RecordId: resource.Id,
					Property: propertyPrefix + "Length",
					Message:  "Length should be positive number for string type",
					Value:    nil,
				})
			}
		}
	}

	if len(errorFields) > 0 {
		var details []string

		for _, errorField := range errorFields {
			details = append(details, fmt.Sprintf("%s: %s", errorField.Property, errorField.Message))
		}

		return errors.ResourceValidationError.WithDetails(strings.Join(details, ";")).WithErrorFields(errorFields)
	}

	return nil
}

func (r *resourceService) GetResourceByName(ctx context.Context, namespace string, resourceName string) (*model.Resource, errors.ServiceError) {
	logger := log.WithFields(logging.CtxFields(ctx))

	logger.Debugf("Begin resource-service GetResourceByName: %s / %s", namespace, resourceName)
	defer logger.Debug("End resource-service GetResourceByName")

	if namespace == "system" {
		logger.Debugf("Call GetSystemResourceByName: %s", resourceName)

		return r.GetSystemResourceByName(ctx, resourceName)
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
	query, err := PrepareQuery(resources.ResourceResource, queryMap)
	logger.Debug("Result record-service: ", query)

	if err != nil {
		return nil, err
	}

	records, _, err := r.backendProviderService.GetSystemBackend(ctx).ListRecords(ctx, abs.ListRecordParams{
		Resource: resources.ResourceResource,
		Query:    query,
		Limit:    1,
	})

	if err != nil {
		return nil, err
	}

	if len(records) == 0 {
		return nil, errors.ResourceNotFoundError.WithDetails(fmt.Sprintf("%s/%s", namespace, resourceName))
	}

	resource := mapping2.ResourceFromRecord(records[0])

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

func (r *resourceService) GetSystemResourceByName(ctx context.Context, resourceName string) (*model.Resource, errors.ServiceError) {
	if resourceName == resources.UserResource.Name {
		return resources.UserResource, nil
	} else if resourceName == resources.DataSourceResource.Name {
		return resources.DataSourceResource, nil
	} else if resourceName == resources.NamespaceResource.Name {
		return resources.NamespaceResource, nil
	} else if resourceName == resources.ExtensionResource.Name {
		return resources.ExtensionResource, nil
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

	var referenceMap = make(map[string]abs.ReferenceMapEntry)

	referenceMap[resources.ResourceResource.Name] = abs.ReferenceMapEntry{
		Catalog:  resources.ResourceResource.SourceConfig.Catalog,
		Entity:   resources.ResourceResource.SourceConfig.Entity,
		IdColumn: "id",
	}

	referenceMap[resources.DataSourceResource.Name] = abs.ReferenceMapEntry{
		Catalog:  resources.DataSourceResource.SourceConfig.Catalog,
		Entity:   resources.DataSourceResource.SourceConfig.Entity,
		IdColumn: "id",
	}

	r.backendProviderService.MigrateResource(resources.ResourceResource, referenceMap)
	r.backendProviderService.MigrateResource(resources.ResourcePropertyResource, referenceMap)
	r.backendProviderService.MigrateResource(resources.ResourceReferenceResource, referenceMap)
}

func (r *resourceService) Delete(ctx context.Context, ids []string, doMigration bool, forceMigration bool) errors.ServiceError {
	for _, resourceId := range ids {
		resource, err := r.Get(ctx, resourceId)

		if err != nil {
			return err
		}

		if err := r.mustModifyResource(resource); err != nil {
			return err
		}

		if err = r.loadResource(ctx, resource); err != nil {
			return err
		}

		if err != nil {
			return err
		}

		err = r.backendProviderService.GetSystemBackend(ctx).DeleteRecords(ctx, resources.ResourceResource, []string{resourceId})

		if err != nil {
			return err
		}

		r.cache.Delete(resourceId)

		if !resource.Virtual && doMigration {
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

func (r *resourceService) mustModifyResource(resource *model.Resource) errors.ServiceError {
	if resource.Namespace == "system" || resource.DataType == model.DataType_SYSTEM {
		return errors.LogicalError.WithMessage("actions on system resource is not allowed")
	}

	if resource.DataType == model.DataType_STATIC {
		return errors.LogicalError.WithMessage("actions on static resource is not allowed")
	}
	return nil
}

func (r *resourceService) List(ctx context.Context) ([]*model.Resource, errors.ServiceError) {
	list, _, err := r.backendProviderService.GetSystemBackend(ctx).ListRecords(ctx, abs.ListRecordParams{
		Resource: resources.ResourceResource,
		Limit:    1000000,
	})

	resources := mapping2.MapFromRecord(list, mapping2.ResourceFromRecord)

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

func (r *resourceService) Get(ctx context.Context, id string) (*model.Resource, errors.ServiceError) {
	record, err := r.backendProviderService.GetSystemBackend(ctx).GetRecord(ctx, resources.ResourceResource, id)

	if err != nil {
		return nil, err
	}

	resource := mapping2.ResourceFromRecord(record)

	err = r.loadResource(ctx, resource)

	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (r *resourceService) loadResource(ctx context.Context, resource *model.Resource) errors.ServiceError {
	if resource.Properties == nil {
		queryMap := make(map[string]interface{})

		queryMap["resource"] = resource.Id

		query, err := PrepareQuery(resources.ResourcePropertyResource, queryMap)

		if err != nil {
			return err
		}

		list, _, err := r.backendProviderService.GetSystemBackend(ctx).ListRecords(ctx, abs.ListRecordParams{
			Resource: resources.ResourcePropertyResource,
			Query:    query,
		})

		if err != nil {
			return err
		}

		resource.Properties = mapping2.MapFromRecord(list, mapping2.ResourcePropertyFromRecord)
	}

	return nil
}

func NewResourceService(backendProviderService abs.BackendProviderService) abs.ResourceService {
	return &resourceService{
		backendProviderService: backendProviderService,
		cache: ttlcache.New[string, *model.Resource](
			ttlcache.WithTTL[string, *model.Resource](1 * time.Minute),
		),
	}
}