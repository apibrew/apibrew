package service

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/tislib/data-handler/pkg/abs"
	"github.com/tislib/data-handler/pkg/errors"
	"github.com/tislib/data-handler/pkg/logging"
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/resources"
	mapping2 "github.com/tislib/data-handler/pkg/resources/mapping"
	"github.com/tislib/data-handler/pkg/service/security"
)

type dataSourceService struct {
	resourceService        abs.ResourceService
	recordService          abs.RecordService
	ServiceName            string
	backendProviderService abs.BackendProviderService
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

	result, _, err := d.recordService.List(ctx, abs.RecordListParams{
		Namespace: resources.DataSourceResource.Namespace,
		Resource:  resources.DataSourceResource.Name,
	})

	if err != nil {
		return nil, err
	}

	return mapping2.MapFromRecord(result, mapping2.DataSourceFromRecord), nil
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
	records := mapping2.MapToRecord(dataSources, mapping2.DataSourceToRecord)
	result, _, err := d.recordService.Create(ctx, abs.RecordCreateParams{
		Namespace: resources.DataSourceResource.Namespace,
		Resource:  resources.DataSourceResource.Name,
		Records:   records,
	})

	if err != nil {
		return nil, err
	}

	return mapping2.MapFromRecord(result, mapping2.DataSourceFromRecord), nil
}

func (d *dataSourceService) Update(ctx context.Context, dataSources []*model.DataSource) ([]*model.DataSource, errors.ServiceError) {
	logger := log.WithFields(logging.CtxFields(ctx))
	logger.WithField("dataSources", dataSources).Debug("Begin data-source Update")
	defer logger.Debug("End data-source Update")

	// insert records via resource service
	records := mapping2.MapToRecord(dataSources, mapping2.DataSourceToRecord)
	result, err := d.recordService.Update(ctx, abs.RecordUpdateParams{
		Namespace: resources.DataSourceResource.Namespace,
		Resource:  resources.DataSourceResource.Name,
		Records:   records,
	})

	if err != nil {
		return nil, err
	}

	for _, item := range dataSources {
		_ = d.backendProviderService.DestroyBackend(ctx, item.Id) //@fixme
	}

	return mapping2.MapFromRecord(result, mapping2.DataSourceFromRecord), nil
}

func (d *dataSourceService) PrepareResourceFromEntity(ctx context.Context, id string, catalog, entity string) (*model.Resource, errors.ServiceError) {
	logger := log.WithFields(logging.CtxFields(ctx))
	logger.WithField("id", id).WithField("entity", entity).Debug("Begin data-source PrepareResourceFromEntity")
	defer logger.Debug("End data-source PrepareResourceFromEntity")

	dsRecord, err := d.recordService.Get(ctx, abs.RecordGetParams{
		Namespace: resources.DataSourceResource.Namespace,
		Resource:  resources.DataSourceResource.Name,
		Id:        id,
	})

	if err != nil {
		return nil, err
	}

	bck, err := d.backendProviderService.GetBackendByDataSourceId(ctx, id)

	if err != nil {
		return nil, err
	}

	resource, err := bck.PrepareResourceFromEntity(ctx, catalog, entity)

	if err != nil {
		return nil, err
	}

	resource.SourceConfig = &model.ResourceSourceConfig{
		DataSource: dsRecord.Properties["name"].GetStringValue(),
		Catalog:    catalog,
		Entity:     entity,
	}

	return resource, nil
}

func (d *dataSourceService) Get(ctx context.Context, id string) (*model.DataSource, errors.ServiceError) {
	logger := log.WithFields(logging.CtxFields(ctx))
	logger.WithField("id", id).Debug("Begin data-source Get")
	defer logger.Debug("End data-source Get")

	record, err := d.recordService.Get(ctx, abs.RecordGetParams{
		Namespace: resources.DataSourceResource.Namespace,
		Resource:  resources.DataSourceResource.Name,
		Id:        id,
	})

	if err != nil {
		return nil, err
	}

	return mapping2.DataSourceFromRecord(record), nil
}

func (d *dataSourceService) Delete(ctx context.Context, ids []string) errors.ServiceError {
	logger := log.WithFields(logging.CtxFields(ctx))
	logger.WithField("ids", ids).Debug("Begin data-source Delete")
	defer logger.Debug("End data-source Delete")

	return d.recordService.Delete(ctx, abs.RecordDeleteParams{
		Namespace: resources.DataSourceResource.Namespace,
		Resource:  resources.DataSourceResource.Name,
		Ids:       ids,
	})
}

func (d *dataSourceService) Init(data *model.InitData) {
	if len(data.InitDataSources) > 0 {
		_, _, err := d.recordService.Create(security.SystemContext, abs.RecordCreateParams{
			Namespace:      resources.DataSourceResource.Namespace,
			Resource:       resources.DataSourceResource.Name,
			Records:        mapping2.MapToRecord(data.InitDataSources, mapping2.DataSourceToRecord),
			IgnoreIfExists: true,
		})

		if err != nil {
			log.Fatal(err)
		}
	}
}

func NewDataSourceService(resourceService abs.ResourceService, recordService abs.RecordService, backendProviderService abs.BackendProviderService) abs.DataSourceService {
	return &dataSourceService{
		ServiceName:            "DataSourceService",
		resourceService:        resourceService,
		recordService:          recordService,
		backendProviderService: backendProviderService,
	}
}
