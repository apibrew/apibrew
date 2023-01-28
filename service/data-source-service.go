package service

import (
	"context"
	"data-handler/logging"
	model "data-handler/model"
	"data-handler/service/errors"
	mapping "data-handler/service/mapping"
	"data-handler/service/params"
	"data-handler/service/security"
	"data-handler/service/system"
	log "github.com/sirupsen/logrus"
)

type DataSourceService interface {
	InjectResourceService(service ResourceService)
	Init(*model.InitData)
	InjectRecordService(service RecordService)
	ListEntities(ctx context.Context, id string) ([]string, errors.ServiceError)
	List(ctx context.Context) ([]*model.DataSource, errors.ServiceError)
	GetStatus(ctx context.Context, id string) (connectionAlreadyInitiated bool, testConnection bool, err errors.ServiceError)
	Create(ctx context.Context, sources []*model.DataSource) ([]*model.DataSource, errors.ServiceError)
	Update(ctx context.Context, sources []*model.DataSource) ([]*model.DataSource, errors.ServiceError)
	PrepareResourceFromEntity(ctx context.Context, dataSourceId string, catalog, entity string) (*model.Resource, errors.ServiceError)
	Get(ctx context.Context, id string) (*model.DataSource, errors.ServiceError)
	Delete(ctx context.Context, ids []string) errors.ServiceError
	InjectBackendProviderService(service BackendProviderService)
}

type dataSourceService struct {
	resourceService        ResourceService
	recordService          RecordService
	ServiceName            string
	backendProviderService BackendProviderService
}

func (d *dataSourceService) InjectBackendProviderService(backendProviderService BackendProviderService) {
	d.backendProviderService = backendProviderService
}

func (d *dataSourceService) ListEntities(ctx context.Context, id string) ([]string, errors.ServiceError) {
	logger := log.WithFields(logging.CtxFields(ctx))
	logger.WithField("req", id).Debug("Begin data-source ListEntities")
	defer logger.Debug("End data-source ListEntities")

	bck, err := d.backendProviderService.GetBackendByDataSourceId(ctx, id)

	if err != nil {
		return nil, err
	}

	return bck.ListEntities(ctx)
}

func (d *dataSourceService) List(ctx context.Context) ([]*model.DataSource, errors.ServiceError) {
	logger := log.WithFields(logging.CtxFields(ctx))
	logger.Debug("Begin data-source List")
	defer logger.Debug("End data-source List")

	result, _, err := d.recordService.List(ctx, params.RecordListParams{
		Namespace: system.DataSourceResource.Namespace,
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

	bck, err := d.backendProviderService.GetBackendByDataSourceId(ctx, id)

	if err != nil {
		return
	}

	return bck.GetStatus(ctx)
}

func (d *dataSourceService) Create(ctx context.Context, dataSources []*model.DataSource) ([]*model.DataSource, errors.ServiceError) {
	logger := log.WithFields(logging.CtxFields(ctx))
	logger.WithField("dataSources", dataSources).Debug("Begin data-source Create")
	defer logger.Debug("End data-source Create")

	// insert records via resource service
	records := mapping.MapToRecord(dataSources, mapping.DataSourceToRecord)
	result, _, err := d.recordService.Create(ctx, params.RecordCreateParams{
		Namespace: system.DataSourceResource.Namespace,
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
	result, err := d.recordService.Update(ctx, params.RecordUpdateParams{
		Namespace: system.DataSourceResource.Namespace,
		Records:   records,
	})

	if err != nil {
		return nil, err
	}

	for _, item := range dataSources {
		_ = d.backendProviderService.DestroyBackend(ctx, item.Id) //@fixme
	}

	return mapping.MapFromRecord(result, mapping.DataSourceFromRecord), nil
}

func (d *dataSourceService) PrepareResourceFromEntity(ctx context.Context, id string, catalog, entity string) (*model.Resource, errors.ServiceError) {
	logger := log.WithFields(logging.CtxFields(ctx))
	logger.WithField("id", id).WithField("entity", entity).Debug("Begin data-source PrepareResourceFromEntity")
	defer logger.Debug("End data-source PrepareResourceFromEntity")

	bck, err := d.backendProviderService.GetBackendByDataSourceId(ctx, id)

	if err != nil {
		return nil, err
	}

	resource, err := bck.PrepareResourceFromEntity(ctx, catalog, entity)

	if err != nil {
		return nil, err
	}

	resource.SourceConfig = &model.ResourceSourceConfig{
		DataSource: id,
		Catalog:    catalog,
		Entity:     entity,
	}

	return resource, nil
}

func (d *dataSourceService) Get(ctx context.Context, id string) (*model.DataSource, errors.ServiceError) {
	logger := log.WithFields(logging.CtxFields(ctx))
	logger.WithField("id", id).Debug("Begin data-source Get")
	defer logger.Debug("End data-source Get")

	record, err := d.recordService.Get(ctx, params.RecordGetParams{
		Namespace: system.DataSourceResource.Namespace,
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

	return d.recordService.Delete(ctx, params.RecordDeleteParams{
		Namespace: system.DataSourceResource.Namespace,
		Resource:  system.DataSourceResource.Name,
		Ids:       ids,
	})
}

func (d *dataSourceService) InjectResourceService(service ResourceService) {
	d.resourceService = service
}

func (d *dataSourceService) InjectRecordService(service RecordService) {
	d.recordService = service
}

func (d *dataSourceService) Init(data *model.InitData) {
	d.backendProviderService.MigrateResource(system.DataSourceResource, nil)

	if len(data.InitDataSources) > 0 {
		_, _, err := d.recordService.Create(security.SystemContext, params.RecordCreateParams{
			Namespace:      system.DataSourceResource.Namespace,
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
