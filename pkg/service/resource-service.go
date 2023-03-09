package service

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/tislib/data-handler/pkg/abs"
	"github.com/tislib/data-handler/pkg/errors"
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/resources"
	mapping "github.com/tislib/data-handler/pkg/resources/mapping"
	util2 "github.com/tislib/data-handler/pkg/util"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/structpb"
	"strconv"
	"strings"
	"sync"
)

type resourceService struct {
	backendProviderService   abs.BackendProviderService
	schema                   abs.Schema
	resourceMigrationService abs.ResourceMigrationService
	mu                       sync.Mutex
}

func (r *resourceService) GetSchema() *abs.Schema {
	return &r.schema
}

func (r *resourceService) PrepareResourceMigrationPlan(ctx context.Context, resources []*model.Resource, prepareFromDataSource bool) ([]*model.ResourceMigrationPlan, errors.ServiceError) {
	var result []*model.ResourceMigrationPlan

	for _, resource := range resources {
		var existingResource *model.Resource
		var err errors.ServiceError

		if prepareFromDataSource {
			existingResource, err = r.backendProviderService.GetSystemBackend(ctx).PrepareResourceFromEntity(ctx, resource.SourceConfig.Catalog, resource.SourceConfig.Entity)

			if err != nil {
				if !errors.RecordNotFoundError.Is(err) {
					return nil, err
				}
			}
		} else {
			existingResource = r.Get(ctx, resource.Id)
		}

		plan, err := r.resourceMigrationService.PreparePlan(ctx, existingResource, resource)

		if err != nil {
			return nil, err
		}

		result = append(result, plan)
	}

	return result, nil
}

func (r *resourceService) ReloadSchema(ctx context.Context) errors.ServiceError {
	records, _, err := r.backendProviderService.GetSystemBackend(ctx).ListRecords(ctx, abs.ListRecordParams{
		Resource: resources.ResourceResource,
		Limit:    1000000,
		ResolveReferences: []string{
			"dataSource",
			"namespace",
		},
		Schema: &r.schema,
	})

	r.schema.Resources = mapping.MapFromRecord(records, mapping.ResourceFromRecord)

	if err != nil {
		return err
	}

	r.schema.ResourceByNamespaceSlashName = make(map[string]*model.Resource)

	r.schema.Resources = append(r.schema.Resources, resources.GetAllSystemResources()...)

	var resourceMap = make(map[string]*model.Resource)
	for _, resource := range r.schema.Resources {
		if resource.Id != "" {
			resourceMap[resource.Id] = resource
		}
		r.schema.ResourceByNamespaceSlashName[resource.Namespace+"/"+resource.Name] = resource
	}

	propertyRecordList, _, err := r.backendProviderService.GetSystemBackend(ctx).ListRecords(ctx, abs.ListRecordParams{
		Resource: resources.ResourcePropertyResource,
		ResolveReferences: []string{
			"resource",
			"reference_resource",
		},
		Schema: &r.schema,
		Limit:  1000000,
	})

	if err != nil {
		return err
	}

	for _, propRec := range propertyRecordList {
		property := mapping.ResourcePropertyFromRecord(propRec)
		propRes := propRec.Properties["resource"].GetStructValue()

		propResource := resourceMap[propRes.Fields["id"].GetStringValue()]

		if propResource == nil {
			panic("propResource is null")
		}

		propResource.Properties = append(propResource.Properties, property)
	}

	if log.GetLevel() == log.TraceLevel {
		// trace all resources
		for _, res := range r.schema.Resources {
			fmt.Println("========" + res.Namespace + "/" + res.Name + "=======")
			jsonRes := protojson.Format(res)
			fmt.Println(jsonRes)
			fmt.Println("================")
		}
	}

	return nil
}

func (r *resourceService) Update(ctx context.Context, resource *model.Resource, doMigration bool, forceMigration bool) errors.ServiceError {
	r.mu.Lock()
	defer func() {
		r.mu.Unlock()
	}()

	existingResource := r.Get(ctx, resource.Id)

	if existingResource == nil {
		return errors.ResourceNotFoundError
	}

	resource.DataType = model.DataType_USER

	if resource.Namespace == "" {
		resource.Namespace = "default"
	}

	if err := validateResource(resource); err != nil {
		return err
	}

	if err := r.mustModifyResource(resource); err != nil {
		return err
	}

	resourceRecords := []*model.Record{mapping.ResourceToRecord(resource)}

	if err := checkAccess(ctx, checkAccessParams{
		Resource:  resource,
		Records:   &resourceRecords,
		Operation: model.OperationType_OPERATION_TYPE_UPDATE,
	}); err != nil {
		return err
	}

	// running migration through existing resource
	//reload schema for comparing
	if err := r.ReloadSchema(ctx); err != nil {
		return err
	}

	// prepare local migration plan
	plan, err := r.resourceMigrationService.PreparePlan(ctx, existingResource, resource)
	if err != nil {
		return err
	}

	if err := r.ApplyPlan(ctx, plan); err != nil {
		return err
	}

	if !resource.Virtual && doMigration {
		bck, err := r.backendProviderService.GetBackendByDataSourceName(ctx, resource.SourceConfig.DataSource)

		if err != nil {
			return err
		}

		// prepare local migration plan for backend
		preparedResource, err := bck.PrepareResourceFromEntity(ctx, existingResource.SourceConfig.Catalog, existingResource.SourceConfig.Entity)

		if !errors.RecordNotFoundError.Is(err) && err != nil {
			return err
		}

		migrationPlan, err := r.resourceMigrationService.PreparePlan(ctx, preparedResource, resource)
		if err != nil {
			return err
		}

		err = bck.UpgradeResource(ctx, abs.UpgradeResourceParams{
			Resource:       resource,
			MigrationPlan:  migrationPlan,
			ForceMigration: forceMigration,
			Schema:         &r.schema,
		})

		if err != nil {
			return err
		}
	}

	r.mustReloadResources(context.TODO())

	return nil
}

func (r *resourceService) ApplyPlan(ctx context.Context, plan *model.ResourceMigrationPlan) errors.ServiceError {
	var currentPropertyMap = util2.GetNamedMap(plan.CurrentResource.Properties)
	var existingPropertyMap = util2.GetNamedMap(plan.ExistingResource.Properties)

	for _, step := range plan.Steps {
		switch sk := step.Kind.(type) {
		case *model.ResourceMigrationStep_UpdateResource:
			resourceRecords := []*model.Record{mapping.ResourceToRecord(plan.CurrentResource)}

			_, err := r.backendProviderService.GetSystemBackend(ctx).UpdateRecords(ctx, abs.BulkRecordsParams{
				Resource:       resources.ResourceResource,
				Records:        resourceRecords,
				CheckVersion:   false,
				IgnoreIfExists: false,
				Schema:         r.GetSchema(),
			})

			if err != nil {
				return err
			}
		case *model.ResourceMigrationStep_CreateResource:
		case *model.ResourceMigrationStep_DeleteResource:
		case *model.ResourceMigrationStep_CreateProperty:
			_, _, err := r.backendProviderService.GetSystemBackend(ctx).AddRecords(ctx, abs.BulkRecordsParams{
				Resource:       resources.ResourcePropertyResource,
				Records:        []*model.Record{mapping.ResourcePropertyToRecord(currentPropertyMap[sk.CreateProperty.Property], plan.CurrentResource)},
				CheckVersion:   false,
				IgnoreIfExists: false,
				Schema:         r.GetSchema(),
			})

			if err != nil {
				return err
			}

			if err != nil && err.Code() == model.ErrorCode_UNIQUE_VIOLATION {
				return errors.AlreadyExistsError.WithMessage(fmt.Sprintf("resource is already exiss: " + plan.CurrentResource.Name))
			}

			if err != nil {
				return err
			}
		case *model.ResourceMigrationStep_DeleteProperty:
			err := r.backendProviderService.GetSystemBackend(ctx).DeleteRecords(ctx, resources.ResourcePropertyResource, []string{existingPropertyMap[sk.DeleteProperty.ExistingProperty].Id})

			if err != nil {
				return err
			}
		case *model.ResourceMigrationStep_UpdateProperty:
			propertyRecord := mapping.ResourcePropertyToRecord(currentPropertyMap[sk.UpdateProperty.Property], plan.CurrentResource)
			propertyRecord.Id = existingPropertyMap[sk.UpdateProperty.ExistingProperty].Id

			_, err := r.backendProviderService.GetSystemBackend(ctx).UpdateRecords(ctx, abs.BulkRecordsParams{
				Resource:       resources.ResourcePropertyResource,
				Records:        []*model.Record{propertyRecord},
				CheckVersion:   false,
				IgnoreIfExists: false,
				Schema:         r.GetSchema(),
			})

			if err != nil {
				return err
			}
		}
	}

	resource := plan.CurrentResource

	propertyRecords := mapping.MapToRecord(resource.Properties, func(property *model.ResourceProperty) *model.Record {
		return mapping.ResourcePropertyToRecord(property, resource)
	})

	propertyRecordList, _, err := r.backendProviderService.GetSystemBackend(ctx).ListRecords(ctx, abs.ListRecordParams{
		Resource: resources.ResourcePropertyResource,
		Query:    newEqualExpression("resource", structpb.NewStringValue(resource.Id)),
		Schema:   &r.schema,
		Limit:    1000000,
	})

	var propertyNameIdMap = make(map[string]string)
	for _, prop := range propertyRecordList {
		propertyNameIdMap[prop.Properties["name"].GetStringValue()] = prop.Id
	}

	for _, prop := range propertyRecords {
		prop.Id = propertyNameIdMap[prop.Properties["name"].GetStringValue()]
	}

	_, err = r.backendProviderService.GetSystemBackend(ctx).UpdateRecords(ctx, abs.BulkRecordsParams{
		Resource:       resources.ResourcePropertyResource,
		Records:        propertyRecords,
		CheckVersion:   false,
		IgnoreIfExists: false,
		Schema:         r.GetSchema(),
	})

	if err != nil {
		return err
	}

	return nil
}

func (r *resourceService) Create(ctx context.Context, resource *model.Resource, doMigration bool, forceMigration bool) (*model.Resource, errors.ServiceError) {
	resource.DataType = model.DataType_USER

	if resource.Namespace == "" {
		resource.Namespace = "default"
	}

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

			r.mustReloadResources(context.TODO())
		} else {
			err = systemBackend.RollbackTransaction(txCtx)

			if err != nil {
				log.Print(err)
			}
		}
	}()

	result, _, err := systemBackend.AddRecords(txCtx, abs.BulkRecordsParams{
		Resource:       resources.ResourceResource,
		Records:        []*model.Record{mapping.ResourceToRecord(resource)},
		CheckVersion:   false,
		IgnoreIfExists: false,
		Schema:         r.GetSchema(),
	})

	if err != nil && err.Code() == model.ErrorCode_UNIQUE_VIOLATION {
		return nil, errors.AlreadyExistsError.WithMessage(fmt.Sprintf("resource is already exiss: " + resource.Name))
	}

	if err != nil && err.Code() == model.ErrorCode_RECORD_VALIDATION_ERROR {
		return nil, errors.ResourceValidationError.WithMessage(err.Error()).WithDetails(err.GetDetails()).WithErrorFields(util2.GetErrorFields(err))
	}

	if err != nil && err.Code() == model.ErrorCode_REFERENCE_VIOLATION {
		return nil, errors.ResourceValidationError.WithMessage(err.Error()).WithDetails(err.GetDetails()).WithErrorFields(util2.GetErrorFields(err))
	}

	if err != nil {
		return nil, err
	}

	// fetch inserted record

	insertedRecord, err := systemBackend.GetRecord(txCtx, resources.ResourceResource, r.GetSchema(), result[0].Id)

	if err != nil {
		return nil, err
	}

	insertedResource := mapping.ResourceFromRecord(insertedRecord)

	if !insertedResource.Virtual && insertedResource.SourceConfig.DataSource == "" {
		return nil, errors.ResourceValidationError.WithMessage("DataSource not found with name: " + resource.SourceConfig.DataSource)
	}

	resource.Id = result[0].Id

	if len(resource.Properties) > 0 {
		propertyRecords := mapping.MapToRecord(resource.Properties, func(property *model.ResourceProperty) *model.Record {
			return mapping.ResourcePropertyToRecord(property, resource)
		})

		_, _, err = systemBackend.AddRecords(txCtx, abs.BulkRecordsParams{
			Resource:       resources.ResourcePropertyResource,
			Records:        propertyRecords,
			CheckVersion:   false,
			IgnoreIfExists: false,
			Schema:         r.GetSchema(),
		})

		if err != nil {
			return nil, err
		}
	}

	if !resource.Virtual && doMigration {
		bck, err := r.backendProviderService.GetBackendByDataSourceName(ctx, resource.SourceConfig.DataSource)

		if err != nil {
			return nil, err
		}

		preparedResource, err := bck.PrepareResourceFromEntity(ctx, resource.SourceConfig.Catalog, resource.SourceConfig.Entity)

		if !errors.RecordNotFoundError.Is(err) && err != nil {
			return nil, err
		}

		migrationPlan, err := r.resourceMigrationService.PreparePlan(ctx, preparedResource, resource)
		if err != nil {
			return nil, err
		}

		err = bck.UpgradeResource(ctx, abs.UpgradeResourceParams{
			Resource:       resource,
			MigrationPlan:  migrationPlan,
			ForceMigration: forceMigration,
			Schema:         &r.schema,
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

		if prop.Mapping == "" {
			errorFields = append(errorFields, &model.ErrorField{
				RecordId: resource.Id,
				Property: propertyPrefix + "Mapping",
				Message:  "Mapping should not be blank",
				Value:    nil,
			})
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

func (r *resourceService) GetResourceByName(_ context.Context, namespace string, resourceName string) *model.Resource {
	if namespace == "" {
		namespace = "default"
	}

	for _, item := range r.schema.Resources {
		if item.Namespace == namespace && item.Name == resourceName {
			return item
		}
	}

	return nil
}

func (r *resourceService) GetSystemResourceByName(ctx context.Context, resourceName string) *model.Resource {
	return r.GetResourceByName(ctx, "system", resourceName)
}

func (r *resourceService) CheckResourceExists(ctx context.Context, namespace, name string) bool {
	return r.GetResourceByName(ctx, namespace, name) != nil
}

func (r *resourceService) Init(_ *model.InitData) {
	r.schema.Resources = append(r.schema.Resources, resources.GetAllSystemResources()...)

	r.schema.ResourceByNamespaceSlashName = make(map[string]*model.Resource)
	for _, resource := range r.schema.Resources {
		r.schema.ResourceByNamespaceSlashName[resource.Namespace+"/"+resource.Name] = resource
	}

	r.MigrateResource(resources.NamespaceResource, r.schema)
	r.MigrateResource(resources.DataSourceResource, r.schema)

	r.MigrateResource(resources.ResourceResource, r.schema)
	r.MigrateResource(resources.ResourcePropertyResource, r.schema)

	r.MigrateResource(resources.UserResource, r.schema)
	r.MigrateResource(resources.ExtensionResource, r.schema)

	if err := r.ReloadSchema(context.TODO()); err != nil {
		panic(err)
	}
}

func (r *resourceService) MigrateResource(resource *model.Resource, schema abs.Schema) {
	if resource.Annotations == nil {
		resource.Annotations = make(map[string]string)
	}

	preparedResource, err := r.backendProviderService.GetSystemBackend(context.TODO()).PrepareResourceFromEntity(context.TODO(), resource.SourceConfig.Catalog, resource.SourceConfig.Entity)

	if err != nil && !errors.RecordNotFoundError.Is(err) {
		panic(err)
	}

	migrationPlan, err := r.resourceMigrationService.PreparePlan(context.TODO(), preparedResource, resource)
	if err != nil {
		panic(err)
	}

	if len(migrationPlan.Steps) == 0 {
		return
	}

	fmt.Println("========" + resource.Namespace + "/" + resource.Name + "=======")
	for _, step := range migrationPlan.Steps {
		jsonRes := protojson.Format(step)
		fmt.Println(jsonRes)
	}
	fmt.Println("================")

	err = r.backendProviderService.GetSystemBackend(context.TODO()).UpgradeResource(context.TODO(), abs.UpgradeResourceParams{
		Resource:       resource,
		MigrationPlan:  migrationPlan,
		ForceMigration: true,
		Schema:         &schema,
	})

	if err != nil {
		panic(err)
	}
}

func (r *resourceService) Delete(ctx context.Context, ids []string, doMigration bool, forceMigration bool) errors.ServiceError {
	for _, resourceId := range ids {
		resource := r.Get(ctx, resourceId)

		if resource == nil {
			return errors.ResourceNotFoundError.WithDetails("Id: " + resourceId)
		}

		var err errors.ServiceError

		if err = r.mustModifyResource(resource); err != nil {
			return err
		}

		if err != nil {
			return err
		}

		err = r.backendProviderService.GetSystemBackend(ctx).DeleteRecords(ctx, resources.ResourceResource, []string{resourceId})

		if err != nil {
			return err
		}

		if !resource.Virtual && doMigration {
			bck, err := r.backendProviderService.GetBackendByDataSourceName(ctx, resource.SourceConfig.DataSource)

			if err != nil {
				return err
			}
			err = bck.DowngradeResource(ctx, resource, forceMigration)

			if err != nil {
				return err
			}
		}
	}

	r.mustReloadResources(context.TODO())

	return nil
}

func (r *resourceService) mustReloadResources(ctx context.Context) {
	if err := r.ReloadSchema(ctx); err != nil {
		panic(err)
	}
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

func (r *resourceService) List(_ context.Context) []*model.Resource {
	return r.schema.Resources
}

func (r *resourceService) Get(_ context.Context, id string) *model.Resource {
	for _, item := range r.schema.Resources {
		if item.Id != "" && item.Id == id {
			return item
		}
	}

	return nil
}

func NewResourceService(backendProviderService abs.BackendProviderService, resourceMigrationService abs.ResourceMigrationService) abs.ResourceService {
	return &resourceService{
		backendProviderService:   backendProviderService,
		resourceMigrationService: resourceMigrationService,
	}
}
