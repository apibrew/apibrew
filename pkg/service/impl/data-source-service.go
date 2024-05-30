package impl

import (
	"context"
	"github.com/apibrew/apibrew/pkg/logging"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resources"
	"github.com/apibrew/apibrew/pkg/service"
	log "github.com/sirupsen/logrus"
)

type dataSourceService struct {
	resourceService        service.ResourceService
	recordService          service.RecordService
	ServiceName            string
	backendProviderService service.BackendProviderService
}

func (d *dataSourceService) ListEntities(ctx context.Context, id string) ([]*model.DataSourceCatalog, error) {
	logger := log.WithFields(logging.CtxFields(ctx))
	logger.WithField("req", id).Debug("Begin data-source ListEntities")
	defer logger.Debug("End data-source ListEntities")

	return d.backendProviderService.ListEntities(ctx, id)
}

func (d *dataSourceService) GetStatus(ctx context.Context, id string) (connectionAlreadyInitiated bool, testConnection bool, err error) {
	logger := log.WithFields(logging.CtxFields(ctx))
	logger.WithField("id", id).Debug("Begin data-source GetStatus")
	defer logger.Debug("End data-source GetStatus")

	return d.backendProviderService.GetStatus(ctx, id)
}

func (d *dataSourceService) PrepareResourceFromEntity(ctx context.Context, id string, catalog, entity string) (*model.Resource, error) {
	logger := log.WithFields(logging.CtxFields(ctx))
	logger.WithField("id", id).WithField("entity", entity).Debug("Begin data-source PrepareResourceFromEntity")
	defer logger.Debug("End data-source PrepareResourceFromEntity")

	dsRecord, err := d.recordService.Get(ctx, service.RecordGetParams{
		Namespace: resources.DataSourceResource.Namespace,
		Resource:  resources.DataSourceResource.Name,
		Id:        id,
	})

	if err != nil {
		return nil, err
	}

	resource, err := d.backendProviderService.PrepareResourceFromEntity(ctx, dsRecord.GetProperties()["name"].GetStringValue(), catalog, entity)

	if err != nil {
		return nil, err
	}

	resource.SourceConfig = &model.ResourceSourceConfig{
		DataSource: dsRecord.GetProperties()["name"].GetStringValue(),
		Catalog:    catalog,
		Entity:     entity,
	}

	return resource, nil
}

func (d *dataSourceService) Delete(ctx context.Context, ids []string) error {
	logger := log.WithFields(logging.CtxFields(ctx))
	logger.WithField("ids", ids).Debug("Begin data-source Delete_")
	defer logger.Debug("End data-source Delete_")

	return d.recordService.Delete(ctx, service.RecordDeleteParams{
		Namespace: resources.DataSourceResource.Namespace,
		Resource:  resources.DataSourceResource.Name,
		Ids:       ids,
	})
}

func (d *dataSourceService) Init(config *model.AppConfig) {

}

func NewDataSourceService(resourceService service.ResourceService, recordService service.RecordService, backendProviderService service.BackendProviderService) service.DataSourceService {
	return &dataSourceService{
		ServiceName:            "DataSourceService",
		resourceService:        resourceService,
		recordService:          recordService,
		backendProviderService: backendProviderService,
	}
}
