package impl

import (
	"context"
	"fmt"
	"github.com/apibrew/apibrew/pkg/abs"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resource_model"
	"github.com/apibrew/apibrew/pkg/resources"
	"github.com/apibrew/apibrew/pkg/resources/mapping"
	"github.com/apibrew/apibrew/pkg/service"
	"github.com/apibrew/apibrew/pkg/service/annotations"
	"github.com/apibrew/apibrew/pkg/service/validate"
	"github.com/apibrew/apibrew/pkg/util"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/structpb"
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

func (r *resourceService) LocateLocalReferences(resource *model.Resource) []string {
	references := r.GetSchema().ResourcePropertiesByType[resource.Namespace+"/"+resource.Name][model.ResourceProperty_REFERENCE]
	var filteredReferences []string

	var existingPathMap = make(map[string]bool)

	for _, reference := range references {
		var existingReferenceToResolveMap = make(map[string]bool)
		if existingReferenceToResolveMap[reference.Path] {
			continue
		}
		existingReferenceToResolveMap[reference.Path] = true

		if !existingPathMap[reference.Path] {
			filteredReferences = append(filteredReferences, reference.Path)
		}

		existingPathMap[reference.Path] = true
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
			existingResource, err = r.backendProviderService.PrepareResourceFromEntity(ctx, "system", resource.SourceConfig.Catalog, resource.SourceConfig.Entity)

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
	ctx = annotations.SetWithContext(ctx, annotations.UseJoinTable, "true") // O(1)
	records, _, err := r.backendProviderService.ListRecords(ctx, resources.ResourceResource, abs.ListRecordParams{
		Limit: 1000000,
		ResolveReferences: []string{
			"dataSource",
			"namespace",
		},
	}, nil)

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

func (r *resourceService) Update(ctx context.Context, resource *model.Resource, doMigration bool, forceMigration bool) (err errors.ServiceError) {
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

	defer func() {
		if err != nil {
			log.Print(err)
		}
	}()

	if resource.Namespace == "" {
		resource.Namespace = existingResource.Namespace
	}

	if resource.SourceConfig == nil {
		resource.SourceConfig = existingResource.SourceConfig
	}

	// datasource change not allowed
	resource.SourceConfig.DataSource = existingResource.SourceConfig.DataSource

	if resource.SourceConfig.Entity == "" {
		resource.SourceConfig.Entity = existingResource.SourceConfig.Entity
	}

	util.NormalizeResource(resource)

	resource.Version = existingResource.Version
	resource.AuditData = existingResource.AuditData

	if err := validate.ValidateResource(resource); err != nil {
		return err
	}

	if err := r.mustModifyResource(resource); err != nil {
		return err
	}

	resourceRecords := []unstructured.Unstructured{mapping.ResourceToRecord(resource)}

	if err := r.authorizationService.CheckRecordAccess(ctx, service.CheckRecordAccessParams{
		Resource:  resources.ResourceResource,
		Records:   &resourceRecords,
		Operation: resource_model.PermissionOperation_UPDATE,
	}); err != nil {
		return err
	}

	defer func() {
		if err := r.reloadSchema(util.WithSystemContext(context.TODO())); err != nil {
			log.Fatal(err)
		}
	}()

	for _, record := range resourceRecords {
		PrepareUpdateForRecord(ctx, resources.ResourceResource, record)
	}

	if _, err = r.backendProviderService.UpdateRecords(ctx, resources.ResourceResource, resourceRecords); err != nil {
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

	if err := r.validateMigration(ctx, resource, plan); err != nil {
		return err
	}

	if !resource.Virtual && doMigration {
		if forceMigration {
			// prepare local migration plan for backend
			existingResource, err = r.backendProviderService.PrepareResourceFromEntity(ctx, "system", existingResource.SourceConfig.Catalog, existingResource.SourceConfig.Entity)

			if !errors.RecordNotFoundError.Is(err) && err != nil {
				return err
			}

			plan, err = r.resourceMigrationService.PreparePlan(ctx, existingResource, resource)
			if err != nil {
				return err
			}
		}

		err = r.backendProviderService.UpgradeResource(ctx, "system", abs.UpgradeResourceParams{
			MigrationPlan:  plan,
			ForceMigration: forceMigration,
		})

		if err != nil {
			return err
		}

		if annotations.IsEnabled(resource, annotations.KeepHistory) {
			err = r.backendProviderService.UpgradeResource(ctx, resource.SourceConfig.DataSource, abs.UpgradeResourceParams{
				MigrationPlan:  util.HistoryPlan(plan),
				ForceMigration: forceMigration,
			})

			if err != nil {
				return err
			}
		}
	}

	r.registerResourceToSchema(resource)

	return nil
}

func (r *resourceService) Create(ctx context.Context, resource *model.Resource, doMigration bool, forceMigration bool) (res *model.Resource, err errors.ServiceError) {
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

	if err := validate.ValidateResource(resource); err != nil {
		return nil, err
	}

	if err := r.mustModifyResource(resource); err != nil {
		return nil, err
	}

	util.NormalizeResource(resource)

	resourceRecord := mapping.ResourceToRecord(resource)

	if err := r.authorizationService.CheckRecordAccess(ctx, service.CheckRecordAccessParams{
		Resource:  resources.ResourceResource,
		Records:   &[]unstructured.Unstructured{resourceRecord},
		Operation: resource_model.PermissionOperation_CREATE,
	}); err != nil {
		return nil, err
	}

	var txCtx = ctx

	InitRecord(ctx, resources.ResourceResource, resourceRecord)

	result, err := r.backendProviderService.AddRecords(txCtx, resources.ResourceResource, []unstructured.Unstructured{resourceRecord})

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

	defer func() {
		if err != nil {
			log.Print("Reverting created records for resource: " + resource.Name + " with id: " + util.GetRecordId(result[0]) + " due to error: " + err.Error() + " ...")
			if err := r.backendProviderService.DeleteRecords(util.WithSystemContext(context.TODO()), resources.ResourceResource, result); err != nil {
				log.Error(err)
			}
		}
	}()

	// fetch inserted record

	txCtx = annotations.SetWithContext(txCtx, annotations.UseJoinTable, "true")

	insertedRecord, err := r.backendProviderService.GetRecord(txCtx, resources.ResourceResource, util.GetRecordId(result[0]), []string{
		"*",
	})

	if err != nil {
		return nil, err
	}

	insertedResource := mapping.ResourceFromRecord(insertedRecord)

	if !insertedResource.Virtual && insertedResource.SourceConfig.DataSource == "" {
		return nil, errors.ResourceValidationError.WithMessage("DataSource not found with name: " + resource.SourceConfig.DataSource)
	}

	resource.Id = util.GetRecordId(result[0])

	if !resource.Virtual && doMigration {
		var plan *model.ResourceMigrationPlan
		var existingResource *model.Resource

		if forceMigration {
			existingResource, err = r.backendProviderService.PrepareResourceFromEntity(ctx, resource.SourceConfig.DataSource, resource.SourceConfig.Catalog, resource.SourceConfig.Entity)

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

		err = r.backendProviderService.UpgradeResource(ctx, resource.SourceConfig.DataSource, abs.UpgradeResourceParams{
			MigrationPlan:  plan,
			ForceMigration: forceMigration,
		})

		if err != nil {
			return nil, err
		}

		if annotations.IsEnabled(resource, annotations.KeepHistory) {
			err = r.backendProviderService.UpgradeResource(ctx, resource.SourceConfig.DataSource, abs.UpgradeResourceParams{
				MigrationPlan:  util.HistoryPlan(plan),
				ForceMigration: forceMigration,
			})

			if err != nil {
				return nil, err
			}
		}
	}

	r.schema.Resources = append(r.schema.Resources, insertedResource)
	r.registerResourceToSchema(insertedResource)

	return resource, nil
}

func (r *resourceService) GetResourceByName(ctx context.Context, namespace string, resourceName string) (*model.Resource, errors.ServiceError) {
	if namespace == "" {
		namespace = "default"
	}

	for _, item := range r.schema.Resources {
		if item.Namespace == namespace && item.Name == resourceName {

			if err := r.authorizationService.CheckRecordAccess(ctx, service.CheckRecordAccessParams{
				Resource:  resources.ResourceResource,
				Operation: resource_model.PermissionOperation_READ,
			}); err != nil {
				if err := r.authorizationService.CheckRecordAccess(ctx, service.CheckRecordAccessParams{
					Resource:  item,
					Operation: resource_model.PermissionOperation_READ,
				}); err != nil {
					return nil, err
				}
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
	r.migrateResource(resources.ResourceActionResource)
	r.migrateResource(resources.UserResource)
	r.migrateResource(resources.RoleResource)
	r.migrateResource(resources.PermissionResource)
	r.migrateResource(resources.ExtensionResource)
	r.migrateResource(resources.AuditLogResource)

	if err := r.reloadSchema(util.WithSystemContext(context.TODO())); err != nil {
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
		r.registerResourceToSchema(resource)
	}
}

func (r *resourceService) registerResourceToSchema(resource *model.Resource) {

	r.schema.ResourceByNamespaceSlashName[resource.Namespace+"/"+resource.Name] = resource
	r.schema.ResourcePropertiesByType[resource.Namespace+"/"+resource.Name] = r.mapPropertiesByType(resource)

	r.schema.ResourceBySlug[util.ResourceRestPath(resource)] = resource

	log.Debugf("Registered resource to schema: %s/%s", resource.Namespace, resource.Name)
}

func (r *resourceService) migrateResource(resource *model.Resource) {
	if resource.Annotations == nil {
		resource.Annotations = make(map[string]string)
	}

	if resource.AuditData == nil {
		resource.AuditData = &model.AuditData{}
	}

	preparedResource, err := r.backendProviderService.PrepareResourceFromEntity(context.TODO(), "system", resource.SourceConfig.Catalog, resource.SourceConfig.Entity)

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

	err = r.backendProviderService.UpgradeResource(context.TODO(), "system", abs.UpgradeResourceParams{
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
			for _, item := range r.schema.Resources {
				log.Print(item.Id, item.Namespace, item.Name)
			}
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
			Resource: resources.ResourceResource,
			Records: &[]unstructured.Unstructured{{
				Properties: map[string]*structpb.Value{
					"id": structpb.NewStringValue(resourceId),
				},
			}},
			Operation: resource_model.PermissionOperation_DELETE,
		}); err != nil {
			return err
		}

		err = r.backendProviderService.DeleteRecords(ctx, resources.ResourceResource, []unstructured.Unstructured{util.IdRecord(resourceId)})

		if err != nil {
			return err
		}

		if !resource.Virtual && doMigration {
			plan, err := r.resourceMigrationService.PreparePlan(ctx, resource, nil)

			if err != nil {
				return err
			}

			err = r.backendProviderService.UpgradeResource(ctx, resource.SourceConfig.DataSource, abs.UpgradeResourceParams{
				MigrationPlan:  plan,
				ForceMigration: forceMigration,
			})

			if err != nil {
				return err
			}

			if annotations.IsEnabled(resource, annotations.KeepHistory) {
				err = r.backendProviderService.UpgradeResource(ctx, resource.SourceConfig.DataSource, abs.UpgradeResourceParams{
					MigrationPlan:  util.HistoryPlan(plan),
					ForceMigration: forceMigration,
				})

				if err != nil {
					return err
				}
			}
		}
	}

	r.mustReloadResources()

	return nil
}

func (r *resourceService) mustReloadResources() {
	if err := r.reloadSchema(util.WithSystemContext(context.TODO())); err != nil {
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
		Operation: resource_model.PermissionOperation_READ,
	}); err == nil {
		return r.schema.Resources, nil
	}

	var filteredResources []*model.Resource
	for _, resource := range r.schema.Resources {
		if err := r.authorizationService.CheckRecordAccess(ctx, service.CheckRecordAccessParams{
			Resource:  resource,
			Operation: resource_model.PermissionOperation_READ,
		}); err == nil {
			filteredResources = append(filteredResources, resource)
		}
	}

	return filteredResources, nil
}

func (r *resourceService) Get(ctx context.Context, id string) (*model.Resource, errors.ServiceError) {
	for _, item := range r.schema.Resources {
		if item.Id != "" && item.Id == id {

			if err := r.authorizationService.CheckRecordAccess(ctx, service.CheckRecordAccessParams{
				Resource: resources.ResourceResource,
				Records: &[]unstructured.Unstructured{
					{
						Properties: map[string]*structpb.Value{
							"id": structpb.NewStringValue(id),
						},
					},
				},
				Operation: resource_model.PermissionOperation_READ,
			}); err != nil {
				if err := r.authorizationService.CheckRecordAccess(ctx, service.CheckRecordAccessParams{
					Resource:  item,
					Operation: resource_model.PermissionOperation_READ,
				}); err != nil {
					return nil, err
				}
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

func (r *resourceService) validateMigration(ctx context.Context, resource *model.Resource, plan *model.ResourceMigrationPlan) errors.ServiceError {
	log.Print(plan)

	return nil
}

func NewResourceService(backendProviderService service.BackendProviderService, resourceMigrationService service.ResourceMigrationService, authorizationService service.AuthorizationService) service.ResourceService {
	srv := &resourceService{
		backendProviderService:   backendProviderService,
		resourceMigrationService: resourceMigrationService,
		authorizationService:     authorizationService,
	}

	backendProviderService.SetSchema(&srv.schema)

	return srv
}
