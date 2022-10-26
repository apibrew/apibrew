package service

import (
	"context"
	model "data-handler/model"
	"data-handler/service/backend"
	"data-handler/service/backend/postgres"
	"data-handler/service/errors"
	mapping "data-handler/service/mapping"
	"data-handler/service/security"
	"data-handler/service/system"
	log "github.com/sirupsen/logrus"
)

type DataSourceService interface {
	GetDataSourceBackend(dataSource *model.DataSource) backend.DataSourceBackend
	GetSystemDataSourceBackend() backend.DataSourceBackend
	GetDataSourceBackendById(dataSourceId string) (backend.DataSourceBackend, errors.ServiceError)
	InjectResourceService(service ResourceService)
	InjectInitData(data *model.InitData)
	Init(*model.InitData)
	InjectRecordService(service RecordService)
	InjectPostgresResourceServiceBackend(serviceBackend backend.ResourceServiceBackend)
	InjectAuthenticationService(service AuthenticationService)
	ListEntities(ctx context.Context, id string) ([]string, errors.ServiceError)
	List(ctx context.Context) ([]*model.DataSource, errors.ServiceError)
	GetStatus(ctx context.Context, id string) (connectionAlreadyInitiated bool, testConnection bool, err errors.ServiceError)
	Create(ctx context.Context, sources []*model.DataSource) ([]*model.DataSource, errors.ServiceError)
	Update(ctx context.Context, sources []*model.DataSource) ([]*model.DataSource, errors.ServiceError)
	PrepareResourceFromEntity(ctx context.Context, id string, entity string) (*model.Resource, errors.ServiceError)
	Get(ctx context.Context, id string) (*model.DataSource, errors.ServiceError)
	Delete(ctx context.Context, ids []string) errors.ServiceError
}

type dataSourceService struct {
	resourceService                ResourceService
	recordService                  RecordService
	systemDataSource               *model.DataSource
	postgresResourceServiceBackend backend.ResourceServiceBackend
	authenticationService          AuthenticationService
	ServiceName                    string
}

func (d *dataSourceService) InjectAuthenticationService(service AuthenticationService) {
	d.authenticationService = service
}

func (d *dataSourceService) ListEntities(ctx context.Context, id string) ([]string, errors.ServiceError) {
	return d.postgresResourceServiceBackend.ListEntities(ctx, id)
}

func (d *dataSourceService) List(ctx context.Context) ([]*model.DataSource, errors.ServiceError) {
	systemCtx := security.WithSystemContext(ctx)
	result, _, err := d.recordService.List(systemCtx, RecordListParams{
		Workspace: system.DataSourceResource.Workspace,
		Resource:  system.DataSourceResource.Name,
	})

	if err != nil {
		return nil, err
	}

	return mapping.MapFromRecord(result, mapping.DataSourceFromRecord), nil
}

func (d *dataSourceService) GetStatus(ctx context.Context, id string) (connectionAlreadyInitiated bool, testConnection bool, err errors.ServiceError) {
	return d.postgresResourceServiceBackend.GetStatus(id)
}

func (d *dataSourceService) Create(ctx context.Context, dataSources []*model.DataSource) ([]*model.DataSource, errors.ServiceError) {
	// insert records via resource service
	records := mapping.MapToRecord(dataSources, mapping.DataSourceToRecord)
	systemCtx := security.WithSystemContext(ctx)
	result, _, err := d.recordService.Create(systemCtx, RecordCreateParams{
		Workspace: system.DataSourceResource.Workspace,
		Resource:  system.DataSourceResource.Name,
		Records:   records,
	})

	if err != nil {
		return nil, err
	}

	return mapping.MapFromRecord(result, mapping.DataSourceFromRecord), nil
}

func (d *dataSourceService) Update(ctx context.Context, dataSources []*model.DataSource) ([]*model.DataSource, errors.ServiceError) {
	// insert records via resource service
	records := mapping.MapToRecord(dataSources, mapping.DataSourceToRecord)
	systemCtx := security.WithSystemContext(ctx)
	result, err := d.recordService.Update(systemCtx, RecordUpdateParams{
		Workspace: system.DataSourceResource.Workspace,
		Records:   records,
	})

	if err != nil {
		return nil, err
	}

	for _, item := range dataSources {
		d.postgresResourceServiceBackend.DestroyDataSource(item.Id)
	}

	return mapping.MapFromRecord(result, mapping.DataSourceFromRecord), nil
}

func (d *dataSourceService) PrepareResourceFromEntity(ctx context.Context, id string, entity string) (*model.Resource, errors.ServiceError) {
	resource, err := d.postgresResourceServiceBackend.PrepareResourceFromEntity(ctx, id, entity)

	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (d *dataSourceService) Get(ctx context.Context, id string) (*model.DataSource, errors.ServiceError) {
	systemCtx := security.WithSystemContext(ctx)
	record, err := d.recordService.Get(systemCtx, RecordGetParams{
		Workspace: system.DataSourceResource.Workspace,
		Resource:  system.DataSourceResource.Name,
		Id:        id,
	})

	if err != nil {
		return nil, err
	}

	return mapping.DataSourceFromRecord(record), nil
}

func (d *dataSourceService) Delete(ctx context.Context, ids []string) errors.ServiceError {
	systemCtx := security.WithSystemContext(ctx)

	return d.recordService.Delete(systemCtx, RecordDeleteParams{
		Workspace: system.DataSourceResource.Workspace,
		Resource:  system.DataSourceResource.Name,
		Ids:       ids,
	})
}

func (d *dataSourceService) GetDataSourceBackendById(dataSourceId string) (backend.DataSourceBackend, errors.ServiceError) {
	if dataSourceId == d.systemDataSource.Id {
		return d.GetSystemDataSourceBackend(), nil
	}

	systemCtx := security.WithSystemContext(context.TODO())
	record, err := d.recordService.GetRecord(systemCtx, system.DataSourceResource.Workspace, system.DataSourceResource.Name, dataSourceId)

	if err != nil {
		return nil, err
	}

	dataSource := mapping.DataSourceFromRecord(record)

	return d.GetDataSourceBackend(dataSource), nil
}

func (d *dataSourceService) GetDataSourceBackend(dataSource *model.DataSource) backend.DataSourceBackend {
	if dataSource == nil {
		panic("data-source is nil")
	}
	switch d.systemDataSource.Backend {
	case model.DataSourceBackend_POSTGRESQL:
		return postgres.NewPostgresDataSourceBackend(dataSource.Id, dataSource.Options.(*model.DataSource_PostgresqlParams).PostgresqlParams)
	case model.DataSourceBackend_MONGODB:
		panic("mongodb data-source not init")
	default:
		panic("unknown data-source type")
	}
}

func (d *dataSourceService) GetSystemDataSourceBackend() backend.DataSourceBackend {
	return d.GetDataSourceBackend(d.systemDataSource)
}

func (d *dataSourceService) InjectPostgresResourceServiceBackend(serviceBackend backend.ResourceServiceBackend) {
	d.postgresResourceServiceBackend = serviceBackend
}

func (d *dataSourceService) InjectResourceService(service ResourceService) {
	d.resourceService = service
}

func (d *dataSourceService) InjectRecordService(service RecordService) {
	d.recordService = service
}

func (d *dataSourceService) InjectInitData(data *model.InitData) {
	d.systemDataSource = data.SystemDataSource
}

func (d *dataSourceService) Init(data *model.InitData) {
	d.resourceService.InitResource(system.DataSourceResource)

	if len(data.InitDataSources) > 0 {
		_, _, err := d.recordService.Create(security.SystemContext, RecordCreateParams{
			Workspace:      system.DataSourceResource.Workspace,
			Records:        mapping.MapToRecord(data.InitDataSources, mapping.DataSourceToRecord),
			IgnoreIfExists: true,
		})

		if err != nil {
			log.Error(err)
		}
	}
}

func NewDataSourceService() DataSourceService {
	return &dataSourceService{
		ServiceName: "DataSourceService",
	}
}
