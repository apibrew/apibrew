package service

import (
	"context"
	"data-handler/logging"
	model "data-handler/model"
	"data-handler/service/backend"
	"data-handler/service/backend/postgres"
	"data-handler/service/errors"
	mapping "data-handler/service/mapping"
	"data-handler/service/params"
	"data-handler/service/security"
	"data-handler/service/system"
	log "github.com/sirupsen/logrus"
)

type DataSourceService interface {
	GetDataSourceBackend(ctx context.Context, dataSource *model.DataSource) backend.DataSourceBackend
	GetSystemDataSourceBackend(ctx context.Context) backend.DataSourceBackend
	GetDataSourceBackendById(ctx context.Context, dataSourceId string) (backend.DataSourceBackend, errors.ServiceError)
	InjectResourceService(service ResourceService)
	InjectInitData(data *model.InitData)
	Init(*model.InitData)
	InjectRecordService(service RecordService)
	InjectAuthenticationService(service AuthenticationService)
	ListEntities(ctx context.Context, id string) ([]string, errors.ServiceError)
	List(ctx context.Context) ([]*model.DataSource, errors.ServiceError)
	GetStatus(ctx context.Context, id string) (connectionAlreadyInitiated bool, testConnection bool, err errors.ServiceError)
	Create(ctx context.Context, sources []*model.DataSource) ([]*model.DataSource, errors.ServiceError)
	Update(ctx context.Context, sources []*model.DataSource) ([]*model.DataSource, errors.ServiceError)
	PrepareResourceFromEntity(ctx context.Context, id string, entity string) (*model.Resource, errors.ServiceError)
	Get(ctx context.Context, id string) (*model.DataSource, errors.ServiceError)
	Delete(ctx context.Context, ids []string) errors.ServiceError
	InjectBackendProviderService(service BackendProviderService)
}

type dataSourceService struct {
	resourceService        ResourceService
	recordService          RecordService
	systemDataSource       *model.DataSource
	authenticationService  AuthenticationService
	ServiceName            string
	backendProviderService BackendProviderService
}

func (d *dataSourceService) InjectBackendProviderService(backendProviderService BackendProviderService) {
	d.backendProviderService = backendProviderService
}

func (d *dataSourceService) InjectAuthenticationService(service AuthenticationService) {
	d.authenticationService = service
}

func (d *dataSourceService) ListEntities(ctx context.Context, id string) ([]string, errors.ServiceError) {
	logger := log.WithFields(logging.CtxFields(ctx))
	logger.WithField("req", id).Debug("Begin data-source ListEntities")
	defer logger.Debug("End data-source ListEntities")

	return d.backendProviderService.GetSystemBackend(ctx).ListEntities(ctx, id)
}

func (d *dataSourceService) List(ctx context.Context) ([]*model.DataSource, errors.ServiceError) {
	logger := log.WithFields(logging.CtxFields(ctx))
	logger.Debug("Begin data-source List")
	defer logger.Debug("End data-source List")

	systemCtx := security.WithSystemContext(ctx)
	result, _, err := d.recordService.List(systemCtx, params.RecordListParams{
		Workspace: system.DataSourceResource.Workspace,
		Resource:  system.DataSourceResource.Name,
	})

	if err != nil {
		return nil, err
	}

	return mapping.MapFromRecord(result, mapping.DataSourceFromRecord), nil
}

func (d *dataSourceService) GetStatus(ctx context.Context, id string) (connectionAlreadyInitiated bool, testConnection bool, err errors.ServiceError) {
	logger := log.WithFields(logging.CtxFields(ctx))
	logger.WithField("id", id).Debug("Begin data-source GetStatus")
	defer logger.Debug("End data-source GetStatus")

	return d.backendProviderService.GetSystemBackend(ctx).GetStatus(ctx, id)
}

func (d *dataSourceService) Create(ctx context.Context, dataSources []*model.DataSource) ([]*model.DataSource, errors.ServiceError) {
	logger := log.WithFields(logging.CtxFields(ctx))
	logger.WithField("dataSources", dataSources).Debug("Begin data-source Create")
	defer logger.Debug("End data-source Create")

	// insert records via resource service
	records := mapping.MapToRecord(dataSources, mapping.DataSourceToRecord)
	systemCtx := security.WithSystemContext(ctx)
	result, _, err := d.recordService.Create(systemCtx, params.RecordCreateParams{
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
	logger := log.WithFields(logging.CtxFields(ctx))
	logger.WithField("dataSources", dataSources).Debug("Begin data-source Update")
	defer logger.Debug("End data-source Update")

	// insert records via resource service
	records := mapping.MapToRecord(dataSources, mapping.DataSourceToRecord)
	systemCtx := security.WithSystemContext(ctx)
	result, err := d.recordService.Update(systemCtx, params.RecordUpdateParams{
		Workspace: system.DataSourceResource.Workspace,
		Records:   records,
	})

	if err != nil {
		return nil, err
	}

	for _, item := range dataSources {
		d.backendProviderService.GetSystemBackend(ctx).DestroyDataSource(ctx, item.Id)
	}

	return mapping.MapFromRecord(result, mapping.DataSourceFromRecord), nil
}

func (d *dataSourceService) PrepareResourceFromEntity(ctx context.Context, id string, entity string) (*model.Resource, errors.ServiceError) {
	logger := log.WithFields(logging.CtxFields(ctx))
	logger.WithField("id", id).WithField("entity", entity).Debug("Begin data-source PrepareResourceFromEntity")
	defer logger.Debug("End data-source PrepareResourceFromEntity")

	resource, err := d.backendProviderService.GetSystemBackend(ctx).PrepareResourceFromEntity(ctx, id, entity)

	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (d *dataSourceService) Get(ctx context.Context, id string) (*model.DataSource, errors.ServiceError) {
	logger := log.WithFields(logging.CtxFields(ctx))
	logger.WithField("id", id).Debug("Begin data-source Get")
	defer logger.Debug("End data-source Get")

	systemCtx := security.WithSystemContext(ctx)
	record, err := d.recordService.Get(systemCtx, params.RecordGetParams{
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
	logger := log.WithFields(logging.CtxFields(ctx))
	logger.WithField("ids", ids).Debug("Begin data-source Delete")
	defer logger.Debug("End data-source Delete")

	systemCtx := security.WithSystemContext(ctx)

	return d.recordService.Delete(systemCtx, params.RecordDeleteParams{
		Workspace: system.DataSourceResource.Workspace,
		Resource:  system.DataSourceResource.Name,
		Ids:       ids,
	})
}

func (d *dataSourceService) GetDataSourceBackendById(ctx context.Context, dataSourceId string) (backend.DataSourceBackend, errors.ServiceError) {
	logger := log.WithFields(logging.CtxFields(ctx))
	logger.WithField("dataSourceId", dataSourceId).Debug("Begin data-source GetDataSourceBackendById")
	defer logger.Debug("End data-source GetDataSourceBackendById")

	if dataSourceId == d.systemDataSource.Id {
		return d.GetSystemDataSourceBackend(ctx), nil
	}

	systemCtx := security.WithSystemContext(context.TODO())
	record, err := d.recordService.GetRecord(systemCtx, system.DataSourceResource.Workspace, system.DataSourceResource.Name, dataSourceId)

	if err != nil {
		return nil, err
	}

	dataSource := mapping.DataSourceFromRecord(record)

	return d.GetDataSourceBackend(ctx, dataSource), nil
}

func (d *dataSourceService) GetDataSourceBackend(ctx context.Context, dataSource *model.DataSource) backend.DataSourceBackend {
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

func (d *dataSourceService) GetSystemDataSourceBackend(ctx context.Context) backend.DataSourceBackend {
	return d.GetDataSourceBackend(ctx, d.systemDataSource)
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
		_, _, err := d.recordService.Create(security.SystemContext, params.RecordCreateParams{
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
