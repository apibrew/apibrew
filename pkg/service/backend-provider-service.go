package service

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/tislib/apibrew/pkg/abs"
	"github.com/tislib/apibrew/pkg/backend/mongo"
	"github.com/tislib/apibrew/pkg/backend/mysql"
	"github.com/tislib/apibrew/pkg/backend/postgres"
	"github.com/tislib/apibrew/pkg/backend/redis"
	"github.com/tislib/apibrew/pkg/errors"
	"github.com/tislib/apibrew/pkg/logging"
	"github.com/tislib/apibrew/pkg/model"
	"github.com/tislib/apibrew/pkg/resources"
	"github.com/tislib/apibrew/pkg/resources/mapping"
	backend_proxy "github.com/tislib/apibrew/pkg/service/backend-proxy"
	"github.com/tislib/apibrew/pkg/service/security"
	"github.com/tislib/apibrew/pkg/util"
)

type backendProviderService struct {
	systemDataSource *model.DataSource
	backendMap       map[string]abs.Backend
	backendIdMap     map[string]string
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
		record, err := b.GetSystemBackend(ctx).GetRecord(systemCtx, resources.DataSourceResource, &abs.Schema{}, dataSourceId)
		util.DeNormalizeRecord(resources.DataSourceResource, record)

		if err != nil {
			return nil, err
		}

		return b.GetBackend(mapping.DataSourceFromRecord(record)), nil
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

		records, _, err := b.GetSystemBackend(ctx).ListRecords(systemCtx, abs.ListRecordParams{
			Resource: resources.DataSourceResource,
			Query:    query,
			Limit:    1,
			Schema:   &abs.Schema{},
		})

		if err != nil {
			return nil, err
		}

		if len(records) == 0 {
			return nil, errors.RecordNotFoundError.WithMessage("Data source not found with name: " + dataSourceName)
		}

		var record = records[0]
		record.Id = records[0].Properties["id"].GetStringValue()

		return b.GetBackend(mapping.DataSourceFromRecord(record)), nil
	}
}

func (b *backendProviderService) GetSystemBackend(_ context.Context) abs.Backend {
	return b.GetBackend(b.systemDataSource)
}

func (b *backendProviderService) GetBackend(dataSource *model.DataSource) abs.Backend {
	if b.backendMap[dataSource.Id] != nil {
		return b.backendMap[dataSource.Id]
	}

	constructor := b.GetBackendConstructor(dataSource.GetBackend())
	instance := constructor(dataSource)

	// apply proxy
	proxy := backend_proxy.NewBackendProxy(instance)

	instance = proxy

	b.backendMap[dataSource.Id] = instance
	b.backendIdMap[dataSource.Id] = dataSource.Name
	b.backendMap[dataSource.Name] = instance

	return instance
}

func (b *backendProviderService) GetBackendConstructor(backend model.DataSourceBackendType) abs.BackendConstructor {
	switch backend {
	case model.DataSourceBackendType_POSTGRESQL:
		return postgres.NewPostgresResourceServiceBackend
	case model.DataSourceBackendType_MYSQL:
		return mysql.NewMysqlResourceServiceBackend
	case model.DataSourceBackendType_MONGODB:
		return mongo.NewMongoResourceServiceBackend
	case model.DataSourceBackendType_REDIS:
		return redis.NewRedisResourceServiceBackend
	}

	panic("Not implemented backend: " + backend.String())
}

func (b *backendProviderService) Init(data *model.InitData) {
	b.systemDataSource = data.SystemDataSource
	b.systemDataSource.Id = "system"
	b.systemDataSource.Name = "system"
}

func NewBackendProviderService() abs.BackendProviderService {
	return &backendProviderService{
		backendMap:   make(map[string]abs.Backend),
		backendIdMap: make(map[string]string),
	}
}
