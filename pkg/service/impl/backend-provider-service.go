package impl

import (
	"context"
	"github.com/apibrew/apibrew/pkg/abs"
	"github.com/apibrew/apibrew/pkg/backend/mongo"
	"github.com/apibrew/apibrew/pkg/backend/mysql"
	"github.com/apibrew/apibrew/pkg/backend/postgres"
	"github.com/apibrew/apibrew/pkg/backend/redis"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/logging"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resource_model"
	"github.com/apibrew/apibrew/pkg/resources"
	"github.com/apibrew/apibrew/pkg/service"
	backend_event_handler "github.com/apibrew/apibrew/pkg/service/backend-event-handler"
	backend_proxy "github.com/apibrew/apibrew/pkg/service/backend-proxy"
	"github.com/apibrew/apibrew/pkg/util"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

type backendProviderService struct {
	systemDataSource *resource_model.DataSource
	backendMap       map[string]abs.Backend
	backendIdMap     map[string]string
	schema           *abs.Schema
	eventHandler     backend_event_handler.BackendEventHandler
}

func (b *backendProviderService) SetSchema(schema *abs.Schema) {
	b.schema = schema
}

func (b *backendProviderService) DestroyBackend(ctx context.Context, dataSourceId string) errors.ServiceError {
	bck, err := b.GetBackendByDataSourceId(ctx, dataSourceId)

	if err != nil {
		return err
	}

	bck.DestroyDataSource(ctx)

	delete(b.backendMap, b.backendIdMap[dataSourceId])
	delete(b.backendMap, dataSourceId)
	delete(b.backendIdMap, dataSourceId)

	return nil
}

func (b *backendProviderService) GetBackendByDataSourceId(ctx context.Context, dataSourceId string) (abs.Backend, errors.ServiceError) {
	if b.backendMap[dataSourceId] != nil {
		return b.backendMap[dataSourceId], nil
	}

	if dataSourceId == b.systemDataSource.Id.String() {
		return b.GetSystemBackend(ctx), nil
	} else {
		systemCtx := util.WithSystemContext(context.TODO())
		record, err := b.GetSystemBackend(ctx).GetRecord(systemCtx, resources.DataSourceResource, dataSourceId, nil)
		DeNormalizeRecord(resources.DataSourceResource, record)

		if err != nil {
			return nil, err
		}

		return b.GetBackend(resource_model.DataSourceMapperInstance.FromRecord(record)), nil
	}
}

func (b *backendProviderService) GetBackendByDataSourceName(ctx context.Context, dataSourceName string) (abs.Backend, errors.ServiceError) {
	if b.backendMap[dataSourceName] != nil {
		return b.backendMap[dataSourceName], nil
	}

	logger := log.WithFields(logging.CtxFields(ctx))
	logger.WithField("dataSourceName", dataSourceName).Debug("Begin data-source GetDataSourceBackendById")
	defer logger.Debug("End data-source GetDataSourceBackendById")

	if dataSourceName == b.systemDataSource.Name {
		return b.GetSystemBackend(ctx), nil
	} else {
		systemCtx := util.WithSystemContext(context.TODO())
		query, err := util.PrepareQuery(resources.DataSourceResource, map[string]interface{}{
			"name": dataSourceName,
		})

		if err != nil {
			return nil, err
		}

		records, _, err := b.GetSystemBackend(ctx).ListRecords(systemCtx, resources.DataSourceResource, abs.ListRecordParams{
			Query: query,
			Limit: 1,
		}, nil)

		if err != nil {
			return nil, err
		}

		if len(records) == 0 {
			return nil, errors.RecordNotFoundError.WithMessage("Data source not found with name: " + dataSourceName)
		}

		var record = records[0]
		record.Id = records[0].Properties["id"].GetStringValue()

		return b.GetBackend(resource_model.DataSourceMapperInstance.FromRecord(record)), nil
	}
}

func (b *backendProviderService) GetSystemBackend(_ context.Context) abs.Backend {
	return b.GetBackend(b.systemDataSource)
}

func (b *backendProviderService) GetBackend(dataSource *resource_model.DataSource) abs.Backend {
	if b.backendMap[dataSource.Id.String()] != nil {
		return b.backendMap[dataSource.Id.String()]
	}

	constructor := b.GetBackendConstructor(dataSource.GetBackend())
	instance := constructor(dataSource)
	instance.SetSchema(b.schema)

	// apply proxy
	proxy := backend_proxy.NewBackendProxy(instance, b.eventHandler)

	instance = proxy

	b.backendMap[dataSource.Id.String()] = instance
	b.backendIdMap[dataSource.Id.String()] = dataSource.Name
	b.backendMap[dataSource.Name] = instance

	return instance
}

func (b *backendProviderService) GetBackendConstructor(backend resource_model.DataSourceBackend) abs.BackendConstructor {
	switch backend {
	case resource_model.DataSourceBackend_POSTGRESQL:
		return postgres.NewPostgresResourceServiceBackend
	case resource_model.DataSourceBackend_MYSQL:
		return mysql.NewMysqlResourceServiceBackend
	case resource_model.DataSourceBackend_MONGODB:
		return mongo.NewMongoResourceServiceBackend
	case resource_model.DataSourceBackend_REDIS:
		return redis.NewRedisResourceServiceBackend
	}

	panic("Not implemented backend: " + string(backend))
}

func (b *backendProviderService) Init(config *model.AppConfig) {
	b.systemDataSource = resource_model.DataSourceMapperInstance.FromRecord(config.SystemDataSource)

	id := uuid.New()
	b.systemDataSource.Id = &id
	b.systemDataSource.Name = "system"
}

func NewBackendProviderService(eventHandler backend_event_handler.BackendEventHandler) service.BackendProviderService {
	return &backendProviderService{
		backendMap:   make(map[string]abs.Backend),
		backendIdMap: make(map[string]string),
		eventHandler: eventHandler,
	}
}
