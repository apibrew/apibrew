package service

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
	"github.com/apibrew/apibrew/pkg/modelnew"
	"github.com/apibrew/apibrew/pkg/resources"
	backend_event_handler "github.com/apibrew/apibrew/pkg/service/backend-event-handler"
	backend_proxy "github.com/apibrew/apibrew/pkg/service/backend-proxy"
	"github.com/apibrew/apibrew/pkg/service/security"
	"github.com/apibrew/apibrew/pkg/util"
	log "github.com/sirupsen/logrus"
)

type backendProviderService struct {
	systemDataSource *modelnew.DataSource
	backendMap       map[string]abs.Backend
	backendIdMap     map[string]string
	schema           *abs.Schema
	eventHandler     backend_event_handler.BackendEventHandler
}

func (b *backendProviderService) SetSchema(schema *abs.Schema) {
	b.schema = schema
}

func (b *backendProviderService) DestroyBackend(ctx context.Context, dataSourceId string) error {
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

	if dataSourceId == b.systemDataSource.Id {
		return b.GetSystemBackend(ctx), nil
	} else {
		systemCtx := security.WithSystemContext(context.TODO())
		record, err := b.GetSystemBackend(ctx).GetRecord(systemCtx, resources.DataSourceResource, dataSourceId)
		util.DeNormalizeRecord(resources.DataSourceResource, record)

		if err != nil {
			return nil, err
		}

		dataSource := &modelnew.DataSource{}

		dataSource.FromRecord(record)

		return b.GetBackend(dataSource), nil
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
		systemCtx := security.WithSystemContext(context.TODO())
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

		dataSource := &modelnew.DataSource{}

		dataSource.FromRecord(record)

		return b.GetBackend(dataSource), nil
	}
}

func (b *backendProviderService) GetSystemBackend(_ context.Context) abs.Backend {
	return b.GetBackend(b.systemDataSource)
}

func (b *backendProviderService) GetBackend(dataSource *modelnew.DataSource) abs.Backend {
	if b.backendMap[dataSource.Id] != nil {
		return b.backendMap[dataSource.Id]
	}

	constructor := b.GetBackendConstructor(dataSource.GetBackend())
	instance := constructor(dataSource)
	instance.SetSchema(b.schema)

	// apply proxy
	proxy := backend_proxy.NewBackendProxy(instance, b.eventHandler)

	instance = proxy

	b.backendMap[dataSource.Id] = instance
	b.backendIdMap[dataSource.Id] = dataSource.Name
	b.backendMap[dataSource.Name] = instance

	return instance
}

func (b *backendProviderService) GetBackendConstructor(backend modelnew.DataSourceBackendType) abs.BackendConstructor {
	switch backend {
	case modelnew.DataSourceBackendType_POSTGRESQL:
		return postgres.NewPostgresResourceServiceBackend
	case modelnew.DataSourceBackendType_MYSQL:
		return mysql.NewMysqlResourceServiceBackend
	case modelnew.DataSourceBackendType_MONGODB:
		return mongo.NewMongoResourceServiceBackend
	case modelnew.DataSourceBackendType_REDIS:
		return redis.NewRedisResourceServiceBackend
	}

	panic("Not implemented backend: " + string(backend))
}

func (b *backendProviderService) Init(config *model.AppConfig) {
	b.systemDataSource = &modelnew.DataSource{}
	b.systemDataSource.FromRecord(config.SystemDataSource)

	b.systemDataSource.Id = "system"
	b.systemDataSource.Name = "system"
}

func NewBackendProviderService(eventHandler backend_event_handler.BackendEventHandler) abs.BackendProviderService {
	return &backendProviderService{
		backendMap:   make(map[string]abs.Backend),
		backendIdMap: make(map[string]string),
		eventHandler: eventHandler,
	}
}
