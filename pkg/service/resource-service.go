package service

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/tislib/data-handler/pkg/abs"
	"github.com/tislib/data-handler/pkg/errors"
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/resources"
	"github.com/tislib/data-handler/pkg/resources/mapping"
	"github.com/tislib/data-handler/pkg/service/annotations"
	"github.com/tislib/data-handler/pkg/types"
	"github.com/tislib/data-handler/pkg/util"
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

	for _, record := range records {
		util.DeNormalizeRecord(resources.ResourceResource, record)
	}

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

	for _, record := range propertyRecordList {
		util.DeNormalizeRecord(resources.ResourcePropertyResource, record)
	}

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
		var bck abs.Backend
		bck, err = r.backendProviderService.GetBackendByDataSourceName(ctx, resource.SourceConfig.DataSource)

		if err != nil {
			return err
		}

		if forceMigration {
			// prepare local migration plan for backend
			existingResource, err = bck.PrepareResourceFromEntity(ctx, existingResource.SourceConfig.Catalog, existingResource.SourceConfig.Entity)

			if !errors.RecordNotFoundError.Is(err) && err != nil {
				return err
			}

			plan, err = r.resourceMigrationService.PreparePlan(ctx, existingResource, resource)
			if err != nil {
				return err
			}
		}

		err = bck.UpgradeResource(ctx, abs.UpgradeResourceParams{
			MigrationPlan:  plan,
			ForceMigration: forceMigration,
			Schema:         &r.schema,
		})

		if err != nil {
			return err
		}

		if annotations.IsEnabled(resource, annotations.KeepHistory) {
			err = bck.UpgradeResource(ctx, abs.UpgradeResourceParams{
				MigrationPlan:  util.HistoryPlan(plan),
				ForceMigration: forceMigration,
				Schema:         &r.schema,
			})

			if err != nil {
				return err
			}
		}
	}

	r.mustReloadResources(context.TODO())

	return nil
}

func (r *resourceService) ApplyPlan(ctx context.Context, plan *model.ResourceMigrationPlan) errors.ServiceError {
	var currentPropertyMap = util.GetNamedMap(plan.CurrentResource.Properties)
	var existingPropertyMap = util.GetNamedMap(plan.ExistingResource.Properties)

	for _, step := range plan.Steps {
		switch sk := step.Kind.(type) {
		case *model.ResourceMigrationStep_UpdateResource:
			resourceRecords := []*model.Record{mapping.ResourceToRecord(plan.CurrentResource)}

			for _, record := range resourceRecords {
				util.PrepareUpdateForRecord(ctx, record)
				util.NormalizeRecord(resources.ResourceResource, record)
			}

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
			propertyCreateRecord := mapping.ResourcePropertyToRecord(currentPropertyMap[sk.CreateProperty.Property], plan.CurrentResource)

			util.InitRecord(ctx, propertyCreateRecord)
			util.NormalizeRecord(resources.ResourcePropertyResource, propertyCreateRecord)

			_, _, err := r.backendProviderService.GetSystemBackend(ctx).AddRecords(ctx, abs.BulkRecordsParams{
				Resource:       resources.ResourcePropertyResource,
				Records:        []*model.Record{propertyCreateRecord},
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
			err := r.backendProviderService.GetSystemBackend(ctx).DeleteRecords(ctx, resources.ResourcePropertyResource, []string{*existingPropertyMap[sk.DeleteProperty.ExistingProperty].Id})

			if err != nil {
				return err
			}
		case *model.ResourceMigrationStep_UpdateProperty:
			propertyRecord := mapping.ResourcePropertyToRecord(currentPropertyMap[sk.UpdateProperty.Property], plan.CurrentResource)
			propertyRecord.Id = *existingPropertyMap[sk.UpdateProperty.ExistingProperty].Id
			util.PrepareUpdateForRecord(ctx, propertyRecord)
			util.NormalizeRecord(resources.ResourcePropertyResource, propertyRecord)

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

	if err != nil {
		return err
	}

	var propertyNameIdMap = make(map[string]string)
	for _, prop := range propertyRecordList {
		propertyNameIdMap[prop.Properties["name"].GetStringValue()] = prop.Properties["id"].GetStringValue()
	}

	for _, propRecord := range propertyRecords {
		propRecord.Id = propertyNameIdMap[propRecord.Properties["name"].GetStringValue()]
		util.PrepareUpdateForRecord(ctx, propRecord)
		util.NormalizeRecord(resources.ResourcePropertyResource, propRecord)
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
	if resource.Namespace == "" {
		resource.Namespace = "default"
	}

	if resource.SourceConfig == nil {
		resource.SourceConfig = &model.ResourceSourceConfig{
			DataSource: "default",
			Entity:     util.ToDashCase(resource.Name),
		}
	}

	if err := validateResource(resource); err != nil {
		return nil, err
	}

	if err := r.mustModifyResource(resource); err != nil {
		return nil, err
	}

	if !annotations.IsEnabled(resource, annotations.NormalizedResource) {
		util.NormalizeResource(resource)
	}

	systemBackend := r.backendProviderService.GetSystemBackend(ctx)

	txk, err := systemBackend.BeginTransaction(ctx, false)

	if err != nil {
		return nil, err
	}

	txCtx := context.WithValue(ctx, abs.TransactionContextKey, txk)

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

	resourceRecord := mapping.ResourceToRecord(resource)
	util.InitRecord(ctx, resourceRecord)
	util.NormalizeRecord(resources.ResourceResource, resourceRecord)

	result, _, err := systemBackend.AddRecords(txCtx, abs.BulkRecordsParams{
		Resource:       resources.ResourceResource,
		Records:        []*model.Record{resourceRecord},
		CheckVersion:   false,
		IgnoreIfExists: false,
		Schema:         r.GetSchema(),
	})

	if err != nil && err.Code() == model.ErrorCode_UNIQUE_VIOLATION {
		return nil, errors.AlreadyExistsError.WithMessage(fmt.Sprintf("resource is already exiss: " + resource.Name))
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
			record := mapping.ResourcePropertyToRecord(property, resource)

			util.InitRecord(ctx, record)
			util.NormalizeRecord(resources.ResourcePropertyResource, record)

			return record
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
		var bck abs.Backend
		var plan *model.ResourceMigrationPlan
		var existingResource *model.Resource

		bck, err = r.backendProviderService.GetBackendByDataSourceName(ctx, resource.SourceConfig.DataSource)
		if err != nil {
			return nil, err
		}

		if forceMigration {
			existingResource, err = bck.PrepareResourceFromEntity(ctx, resource.SourceConfig.Catalog, resource.SourceConfig.Entity)

			if errors.UnsupportedOperation.Is(err) {
				existingResource = resource
			} else {
				if !errors.RecordNotFoundError.Is(err) && err != nil {
					return nil, err
				}
			}

			plan, err = r.resourceMigrationService.PreparePlan(ctx, existingResource, resource)
			if err != nil {
				return nil, err
			}
		} else {
			plan, err = r.resourceMigrationService.PreparePlan(ctx, nil, resource)

			if err != nil {
				return nil, err
			}
		}

		err = bck.UpgradeResource(ctx, abs.UpgradeResourceParams{
			MigrationPlan:  plan,
			ForceMigration: forceMigration,
			Schema:         &r.schema,
		})

		if err != nil {
			return nil, err
		}

		if annotations.IsEnabled(resource, annotations.KeepHistory) {
			err = bck.UpgradeResource(ctx, abs.UpgradeResourceParams{
				MigrationPlan:  util.HistoryPlan(plan),
				ForceMigration: forceMigration,
				Schema:         &r.schema,
			})

			if err != nil {
				return nil, err
			}
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

		if prop.Type == model.ResourceProperty_STRING {
			if prop.Length <= 0 {
				errorFields = append(errorFields, &model.ErrorField{
					RecordId: resource.Id,
					Property: propertyPrefix + "Length",
					Message:  "Length should be positive number for string type",
					Value:    nil,
				})
			}
		}

		// check for additional fields
		if prop.DefaultValue != nil && prop.DefaultValue.AsInterface() != nil {
			propertyType := types.ByResourcePropertyType(prop.Type)
			err := propertyType.ValidatePackedValue(prop.DefaultValue)

			if err != nil {
				errorFields = append(errorFields, &model.ErrorField{
					RecordId: resource.Id,
					Property: propertyPrefix + "DefaultValue",
					Message:  err.Error(),
					Value:    prop.DefaultValue,
				})
			}
		}
		if prop.ExampleValue != nil && prop.ExampleValue.AsInterface() != nil {
			propertyType := types.ByResourcePropertyType(prop.Type)

			if prop.ExampleValue != nil {
				err := propertyType.ValidatePackedValue(prop.ExampleValue)

				if err != nil {
					errorFields = append(errorFields, &model.ErrorField{
						RecordId: resource.Id,
						Property: propertyPrefix + "ExampleValue",
						Message:  err.Error(),
						Value:    prop.ExampleValue,
					})
				}
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

	for _, step := range migrationPlan.Steps {
		jsonRes := protojson.Format(step)
		log.Tracef("Migration plan for %s/%s \n %s", resource.Namespace, resource.Name, jsonRes)
	}

	err = r.backendProviderService.GetSystemBackend(context.TODO()).UpgradeResource(context.TODO(), abs.UpgradeResourceParams{
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

			plan, err := r.resourceMigrationService.PreparePlan(ctx, resource, nil)

			if err != nil {
				return err
			}
			err = bck.UpgradeResource(ctx, abs.UpgradeResourceParams{
				ForceMigration: forceMigration,
				Schema:         r.GetSchema(),
				MigrationPlan:  plan,
			})

			if err != nil {
				return err
			}

			if annotations.IsEnabled(resource, annotations.KeepHistory) {
				err = bck.UpgradeResource(ctx, abs.UpgradeResourceParams{
					MigrationPlan:  util.HistoryPlan(plan),
					ForceMigration: forceMigration,
					Schema:         &r.schema,
				})

				if err != nil {
					return err
				}
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
	if resource.Namespace == "system" {
		return errors.LogicalError.WithMessage("actions on system resource is not allowed")
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
