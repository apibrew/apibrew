package impl

import (
	"context"
	"fmt"
	"github.com/apibrew/apibrew/pkg/abs"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resources"
	"github.com/apibrew/apibrew/pkg/resources/mapping"
	"github.com/apibrew/apibrew/pkg/service"
	"github.com/apibrew/apibrew/pkg/service/annotations"
	"github.com/apibrew/apibrew/pkg/service/validate"
	"github.com/apibrew/apibrew/pkg/util"
	"github.com/gosimple/slug"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/encoding/protojson"
	"strings"
	"sync"
)

type resourceService struct {
	backendProviderService   service.BackendProviderService
	schema                   abs.Schema
	resourceMigrationService service.ResourceMigrationService
	mu                       sync.Mutex
	authorizationService     service.AuthorizationService
}

func (r *resourceService) LocateResourceByReference(resource *model.Resource, reference *model.Reference) *model.Resource {
	if reference == nil {
		return nil
	}

	if reference.Namespace != "" {
		referencedResource, _ := r.GetResourceByName(util.WithSystemContext(context.TODO()), reference.Namespace, reference.Resource)

		return referencedResource
	} else {
		referencedResource, _ := r.GetResourceByName(util.WithSystemContext(context.TODO()), resource.Namespace, reference.Resource)

		return referencedResource
	}
}

func (r *resourceService) LocateReferences(resource *model.Resource, referencesToResolve []string) []string {
	references := r.GetSchema().ResourcePropertiesByType[resource.Namespace+"/"+resource.Name][model.ResourceProperty_REFERENCE]
	var filteredReferences []string

	var existingPathMap = make(map[string]bool)

	for _, reference := range references {
		var existingReferenceToResolveMap = make(map[string]bool)
		for _, checkRef := range referencesToResolve {
			if existingReferenceToResolveMap[checkRef] {
				continue
			}
			existingReferenceToResolveMap[checkRef] = true

			if strings.HasPrefix(checkRef, reference.Path) {
				if !existingPathMap[reference.Path] {
					filteredReferences = append(filteredReferences, reference.Path)
				}

				existingPathMap[reference.Path] = true

				// continue to check deep references
				if checkRef != reference.Path {
					referencedResource := r.LocateResourceByReference(resource, reference.Property.Reference)

					nextRefCheck := strings.Replace(checkRef, reference.Path+".", "", 1)

					subReferences := r.LocateReferences(referencedResource, []string{"$." + nextRefCheck})

					for _, subReference := range subReferences {
						if !existingPathMap[reference.Path+"."+subReference] {
							filteredReferences = append(filteredReferences, reference.Path+"."+subReference[2:])
						}
						existingPathMap[reference.Path+"."+subReference] = true
					}
				}
			}
		}
	}

	return filteredReferences
}

func (r *resourceService) GetSchema() *abs.Schema {
	return &r.schema
}

func (r *resourceService) PrepareResourceMigrationPlan(ctx context.Context, resources []*model.Resource, prepareFromDataSource bool) ([]*model.ResourceMigrationPlan, errors.ServiceError) {
	r.mu.Lock()
	defer func() {
		r.mu.Unlock()
	}()

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
			existingResource, err = r.Get(ctx, resource.Id)

			if err != nil {
				return nil, err
			}
		}

		plan, err := r.resourceMigrationService.PreparePlan(ctx, existingResource, resource)

		if err != nil {
			return nil, err
		}

		result = append(result, plan)
	}

	return result, nil
}

func (r *resourceService) reloadSchema(ctx context.Context) errors.ServiceError {
	ctx = annotations.SetWithContext(ctx, annotations.UseJoinTable, "true")
	records, _, err := r.backendProviderService.GetSystemBackend(ctx).ListRecords(ctx, resources.ResourceResource, abs.ListRecordParams{
		Limit: 1000000,
		ResolveReferences: []string{
			"dataSource",
			"namespace",
		},
	}, nil)

	for _, record := range records {
		DeNormalizeRecord(resources.ResourceResource, record)
	}

	r.schema.Resources = mapping.MapFromRecord(records, mapping.ResourceFromRecord)

	if err != nil {
		return err
	}

	r.schema.Resources = append(r.schema.Resources, resources.GetAllSystemResources()...)
	r.prepareSchemaMappings()

	if err != nil {
		return err
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

	for _, resource := range r.schema.Resources {
		r.schema.ResourcePropertiesByType[resource.Namespace+"/"+resource.Name] = r.mapPropertiesByType(resource)
	}

	return nil
}

func (r *resourceService) Update(ctx context.Context, resource *model.Resource, doMigration bool, forceMigration bool) errors.ServiceError {
	r.mu.Lock()
	defer func() {
		r.mu.Unlock()
	}()

	existingResource, err := r.Get(ctx, resource.Id)

	if err != nil {
		return err
	}

	if existingResource == nil {
		return errors.ResourceNotFoundError.WithDetails(fmt.Sprintf("%s/%s", resource.Namespace, resource.Name))
	}

	if resource.Namespace == "" {
		resource.Namespace = existingResource.Namespace
	}

	if resource.SourceConfig == nil {
		resource.SourceConfig = existingResource.SourceConfig
	}

	if resource.SourceConfig.DataSource == "" {
		resource.SourceConfig.DataSource = existingResource.SourceConfig.DataSource
	}

	if resource.SourceConfig.Entity == "" {
		resource.SourceConfig.Entity = existingResource.SourceConfig.Entity
	}

	if !annotations.IsEnabled(resource, annotations.NormalizedResource) {
		util.NormalizeResource(resource)
	}

	existingPropertiesNamedMap := util.GetNamedMap(existingResource.Properties)
	for _, prop := range resource.Properties {

		if prop.Mapping == "" {
			if existingPropertiesNamedMap[prop.Name] != nil {
				prop.Mapping = existingPropertiesNamedMap[prop.Name].Mapping
			} else {
				prop.Mapping = util.ToSnakeCase(prop.Name)
			}
		}

		if prop.Id == nil || *prop.Id == "" {
			if existingPropertiesNamedMap[prop.Name] != nil {
				prop.Id = existingPropertiesNamedMap[prop.Name].Id
			}
		}
	}

	resource.Version = existingResource.Version
	resource.AuditData = existingResource.AuditData

	if err := validateResource(resource); err != nil {
		return err
	}

	if err := r.mustModifyResource(resource); err != nil {
		return err
	}

	resourceRecords := []*model.Record{mapping.ResourceToRecord(resource)}

	if err := r.authorizationService.CheckRecordAccess(ctx, service.CheckRecordAccessParams{
		Resource:  resources.ResourceResource,
		Records:   &resourceRecords,
		Operation: model.OperationType_OPERATION_TYPE_UPDATE,
	}); err != nil {
		return err
	}

	// running migration through existing resource
	//reload schema for comparing
	if err := r.reloadSchema(ctx); err != nil {
		return err
	}

	// prepare local migration plan
	plan, err := r.resourceMigrationService.PreparePlan(ctx, existingResource, resource)
	if err != nil {
		return err
	}

	if plan.Steps == nil && !forceMigration {
		return nil
	}

	systemBackend := r.backendProviderService.GetSystemBackend(ctx)

	var txCtx = ctx
	var success = false

	if txCtx.Value(abs.TransactionContextKey) == nil {
		txk, err := systemBackend.BeginTransaction(ctx, false)

		if err != nil {
			return err
		}

		txCtx = context.WithValue(ctx, abs.TransactionContextKey, txk)

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

			r.mustReloadResources(context.TODO())
		}()
	}

	if err := r.applyPlan(txCtx, plan); err != nil {
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
		})

		if err != nil {
			return err
		}

		if annotations.IsEnabled(resource, annotations.KeepHistory) {
			err = bck.UpgradeResource(ctx, abs.UpgradeResourceParams{
				MigrationPlan:  util.HistoryPlan(plan),
				ForceMigration: forceMigration,
			})

			if err != nil {
				return err
			}
		}
	}

	success = true

	return nil
}

func (r *resourceService) applyPlan(ctx context.Context, plan *model.ResourceMigrationPlan) errors.ServiceError {
	resourceRecords := []*model.Record{mapping.ResourceToRecord(plan.CurrentResource)}

	for _, record := range resourceRecords {
		PrepareUpdateForRecord(ctx, resources.ResourceResource, record)
		NormalizeRecord(resources.ResourceResource, record)
	}

	_, err := r.backendProviderService.GetSystemBackend(ctx).UpdateRecords(ctx, resources.ResourceResource, resourceRecords)

	return err
}

func (r *resourceService) Create(ctx context.Context, resource *model.Resource, doMigration bool, forceMigration bool) (*model.Resource, errors.ServiceError) {
	r.mu.Lock()
	defer func() {
		r.mu.Unlock()
	}()

	if resource.Namespace == "" {
		resource.Namespace = "default"
	}

	if resource.SourceConfig == nil {
		resource.SourceConfig = &model.ResourceSourceConfig{}
	}

	if resource.SourceConfig.DataSource == "" {
		resource.SourceConfig.DataSource = "default"
	}

	if resource.SourceConfig.Entity == "" {
		resource.SourceConfig.Entity = util.ToSnakeCase(resource.Name)
	}

	for _, prop := range resource.Properties {
		if prop.Mapping == "" {
			prop.Mapping = util.ToSnakeCase(prop.Name)
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

	resourceRecord := mapping.ResourceToRecord(resource)

	if err := r.authorizationService.CheckRecordAccess(ctx, service.CheckRecordAccessParams{
		Resource:  resources.ResourceResource,
		Records:   &[]*model.Record{resourceRecord},
		Operation: model.OperationType_OPERATION_TYPE_CREATE,
	}); err != nil {
		return nil, err
	}

	systemBackend := r.backendProviderService.GetSystemBackend(ctx)

	var txCtx = ctx

	var success = false
	if txCtx.Value(abs.TransactionContextKey) == nil {
		txk, err := systemBackend.BeginTransaction(ctx, false)

		if err != nil {
			return nil, err
		}

		txCtx := context.WithValue(ctx, abs.TransactionContextKey, txk)

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
	}

	InitRecord(ctx, resources.ResourceResource, resourceRecord)
	NormalizeRecord(resources.ResourceResource, resourceRecord)

	result, err := systemBackend.AddRecords(txCtx, resources.ResourceResource, []*model.Record{resourceRecord})

	if err != nil && err.Code() == model.ErrorCode_UNIQUE_VIOLATION {
		return nil, errors.AlreadyExistsError.WithMessage(fmt.Sprintf("resource is already exists: " + resource.Name))
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

	txCtx = annotations.SetWithContext(txCtx, annotations.UseJoinTable, "true")

	insertedRecord, err := systemBackend.GetRecord(txCtx, resources.ResourceResource, result[0].Id)

	if err != nil {
		return nil, err
	}

	insertedResource := mapping.ResourceFromRecord(insertedRecord)

	if !insertedResource.Virtual && insertedResource.SourceConfig.DataSource == "" {
		return nil, errors.ResourceValidationError.WithMessage("DataSource not found with name: " + resource.SourceConfig.DataSource)
	}

	resource.Id = result[0].Id

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
		})

		if err != nil {
			return nil, err
		}

		if annotations.IsEnabled(resource, annotations.KeepHistory) {
			err = bck.UpgradeResource(ctx, abs.UpgradeResourceParams{
				MigrationPlan:  util.HistoryPlan(plan),
				ForceMigration: forceMigration,
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

	errorFields = append(errorFields, validate.ValidateResourceProperties(resource, "", 0, resource.Properties, false)...)

	for _, subType := range resource.Types {
		errorFields = append(errorFields, validate.ValidateResourceProperties(resource, subType.Name+".", 1, subType.Properties, false)...)
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
	if namespace == "" {
		namespace = "default"
	}

	for _, item := range r.schema.Resources {
		if item.Namespace == namespace && item.Name == resourceName {

			if err := r.authorizationService.CheckRecordAccess(ctx, service.CheckRecordAccessParams{
				Resource:  resources.ResourceResource,
				Operation: model.OperationType_OPERATION_TYPE_READ,
			}); err != nil {
				return nil, err
			}

			return item, nil
		}
	}

	return nil, errors.ResourceNotFoundError.WithDetails("Namespace: " + namespace + " resourceName: " + resourceName)
}

func (r *resourceService) GetSystemResourceByName(ctx context.Context, resourceName string) (*model.Resource, errors.ServiceError) {
	return r.GetResourceByName(ctx, "system", resourceName)
}

func (r *resourceService) Init(config *model.AppConfig) {
	r.schema.Resources = append(r.schema.Resources, resources.GetAllSystemResources()...)

	r.prepareSchemaMappings()

	r.migrateResource(resources.NamespaceResource)
	r.migrateResource(resources.DataSourceResource)

	r.migrateResource(resources.ResourceResource)
	r.migrateResource(resources.UserResource)
	r.migrateResource(resources.RoleResource)
	r.migrateResource(resources.SecurityConstraintResource)
	r.migrateResource(resources.ExtensionResource)

	if err := r.reloadSchema(context.TODO()); err != nil {
		panic(err)
	}

	for _, resource := range config.InitResources {
		if _, err := r.Create(util.WithSystemContext(context.TODO()), &model.Resource{
			Name:       resource.Name,
			Namespace:  resource.Namespace,
			Properties: resource.Properties,
			Types:      resource.Types,
			Virtual:    resource.Virtual,
		}, true, true); err != nil {
			panic(err)
		}
	}
}

func (r *resourceService) prepareSchemaMappings() {
	r.schema.ResourceByNamespaceSlashName = make(map[string]*model.Resource)
	r.schema.ResourceBySlug = make(map[string]*model.Resource)
	r.schema.ResourcePropertiesByType = make(map[string]map[model.ResourceProperty_Type][]abs.PropertyWithPath)
	for _, resource := range r.schema.Resources {
		r.schema.ResourceByNamespaceSlashName[resource.Namespace+"/"+resource.Name] = resource
		r.schema.ResourcePropertiesByType[resource.Namespace+"/"+resource.Name] = r.mapPropertiesByType(resource)
		r.schema.ResourceBySlug[slug.Make(resource.Namespace+"/"+resource.Name)] = resource
		if resource.Namespace == "default" {
			r.schema.ResourceBySlug[slug.Make(resource.Name)] = resource
		}
	}
}

func (r *resourceService) migrateResource(resource *model.Resource) {
	if resource.Annotations == nil {
		resource.Annotations = make(map[string]string)
	}

	if resource.AuditData == nil {
		resource.AuditData = &model.AuditData{}
	}

	preparedResource, err := r.backendProviderService.GetSystemBackend(context.TODO()).PrepareResourceFromEntity(context.TODO(), resource.SourceConfig.Catalog, resource.SourceConfig.Entity)

	if err != nil && !errors.RecordNotFoundError.Is(err) {
		log.Fatal(err)
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
	})

	if err != nil {
		panic(err)
	}
}

func (r *resourceService) Delete(ctx context.Context, ids []string, doMigration bool, forceMigration bool) errors.ServiceError {
	r.mu.Lock()
	defer func() {
		r.mu.Unlock()
	}()

	for _, resourceId := range ids {
		resource, err := r.Get(ctx, resourceId)

		if err != nil {
			return err
		}

		if resource == nil {
			return errors.ResourceNotFoundError.WithDetails("Id: " + resourceId)
		}

		if err = r.mustModifyResource(resource); err != nil {
			return err
		}

		if err != nil {
			return err
		}

		if err := r.authorizationService.CheckRecordAccess(ctx, service.CheckRecordAccessParams{
			Resource:  resources.ResourceResource,
			Records:   &[]*model.Record{{Id: resourceId}},
			Operation: model.OperationType_OPERATION_TYPE_DELETE,
		}); err != nil {
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
				MigrationPlan:  plan,
			})

			if err != nil {
				return err
			}

			if annotations.IsEnabled(resource, annotations.KeepHistory) {
				err = bck.UpgradeResource(ctx, abs.UpgradeResourceParams{
					MigrationPlan:  util.HistoryPlan(plan),
					ForceMigration: forceMigration,
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
	if err := r.reloadSchema(ctx); err != nil {
		panic(err)
	}
}

func (r *resourceService) mustModifyResource(resource *model.Resource) errors.ServiceError {
	if resource.Namespace == "system" {
		return errors.LogicalError.WithMessage("actions on system resource is not allowed")
	}

	return nil
}

func (r *resourceService) List(ctx context.Context) ([]*model.Resource, errors.ServiceError) {
	if err := r.authorizationService.CheckRecordAccess(ctx, service.CheckRecordAccessParams{
		Resource:  resources.ResourceResource,
		Operation: model.OperationType_OPERATION_TYPE_READ,
	}); err != nil {
		return nil, err
	}

	return r.schema.Resources, nil
}

func (r *resourceService) Get(ctx context.Context, id string) (*model.Resource, errors.ServiceError) {
	for _, item := range r.schema.Resources {
		if item.Id != "" && item.Id == id {

			if err := r.authorizationService.CheckRecordAccess(ctx, service.CheckRecordAccessParams{
				Resource:  resources.ResourceResource,
				Records:   &[]*model.Record{{Id: id}},
				Operation: model.OperationType_OPERATION_TYPE_READ,
			}); err != nil {
				return nil, err
			}

			return item, nil
		}
	}

	return nil, errors.ResourceNotFoundError.WithDetails("Id: " + id)
}

func (r *resourceService) mapPropertiesByType(resource *model.Resource) map[model.ResourceProperty_Type][]abs.PropertyWithPath {
	result := make(map[model.ResourceProperty_Type][]abs.PropertyWithPath)

	util.ResourceWalkProperties(resource, func(path string, property *model.ResourceProperty) {
		result[property.Type] = append(result[property.Type], abs.PropertyWithPath{
			Path:     path,
			Property: property,
		})
	})

	return result
}

func NewResourceService(backendProviderService service.BackendProviderService, resourceMigrationService service.ResourceMigrationService, authorizationService service.AuthorizationService) service.ResourceService {
	service := &resourceService{
		backendProviderService:   backendProviderService,
		resourceMigrationService: resourceMigrationService,
		authorizationService:     authorizationService,
	}

	backendProviderService.SetSchema(&service.schema)

	return service
}
