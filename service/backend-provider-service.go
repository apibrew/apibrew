package service

import (
	"context"
	"data-handler/logging"
	"data-handler/model"
	"data-handler/service/backend"
	"data-handler/service/backend/mongo"
	"data-handler/service/backend/mysql"
	"data-handler/service/backend/postgres"
	"data-handler/service/errors"
	"data-handler/service/mapping"
	"data-handler/service/security"
	"data-handler/service/system"
	log "github.com/sirupsen/logrus"
)

type BackendProviderService interface {
	backend.DataSourceLocator
	Init(data *model.InitData)
	GetSystemBackend(ctx context.Context) backend.ResourceServiceBackend
	GetBackendByDataSourceId(ctx context.Context, dataSourceId string) (backend.ResourceServiceBackend, errors.ServiceError)
}

type backendProviderService struct {
	postgresBackend  backend.ResourceServiceBackend
	mysqlBackend     backend.ResourceServiceBackend
	mongoBackend     backend.ResourceServiceBackend
	systemDataSource *model.DataSource
}

func (d *backendProviderService) GetDataSourceBackendById(ctx context.Context, dataSourceId string) (backend.DataSourceConnectionDetails, errors.ServiceError) {
	logger := log.WithFields(logging.CtxFields(ctx))
	logger.WithField("dataSourceId", dataSourceId).Debug("Begin data-source GetDataSourceBackendById")
	defer logger.Debug("End data-source GetDataSourceBackendById")

	if dataSourceId == d.systemDataSource.Id {
		return d.GetSystemDataSourceBackend(ctx), nil
	}

	systemCtx := security.WithSystemContext(context.TODO())
	record, err := d.GetSystemBackend(systemCtx).GetRecord(systemCtx, system.DataSourceResource, dataSourceId)

	if err != nil {
		return nil, err
	}

	dataSource := mapping.DataSourceFromRecord(record)

	return d.GetDataSourceBackend(ctx, dataSource), nil
}

func (d *backendProviderService) GetDataSourceBackend(ctx context.Context, dataSource *model.DataSource) backend.DataSourceConnectionDetails {
	if dataSource == nil {
		panic("data-source is nil")
	}
	switch d.systemDataSource.Backend {
	case model.DataSourceBackendType_POSTGRESQL:
		return postgres.NewPostgresDataSourceBackend(dataSource.Id, dataSource.Options.(*model.DataSource_PostgresqlParams).PostgresqlParams)
	case model.DataSourceBackendType_MONGODB:
		panic("mongodb data-source not init")
	default:
		panic("unknown data-source type")
	}
}

func (d *backendProviderService) GetSystemDataSourceBackend(ctx context.Context) backend.DataSourceConnectionDetails {
	return d.GetDataSourceBackend(ctx, d.systemDataSource)
}

func (b *backendProviderService) GetSystemBackend(ctx context.Context) backend.ResourceServiceBackend {
	return b.GetBackend(ctx, b.systemDataSource.Backend)
}

func (b *backendProviderService) GetBackendByDataSourceId(ctx context.Context, dataSourceId string) (backend.ResourceServiceBackend, errors.ServiceError) {
	dsb, err := b.GetDataSourceBackendById(ctx, dataSourceId)

	if err != nil {
		return nil, err
	}

	return b.GetBackend(ctx, dsb.GetBackendType()), nil
}

func (b *backendProviderService) GetBackend(ctx context.Context, backend model.DataSourceBackendType) backend.ResourceServiceBackend {
	switch backend {
	case model.DataSourceBackendType_POSTGRESQL:
		return b.postgresBackend
	case model.DataSourceBackendType_MYSQL:
		return b.mysqlBackend
	case model.DataSourceBackendType_MONGODB:
		return b.mongoBackend
	}

	panic("Not implemented backend: " + backend.String())
}

func (b *backendProviderService) Init(data *model.InitData) {
	b.systemDataSource = data.SystemDataSource

	b.postgresBackend.InjectDataSourceLocator(b)

	b.postgresBackend.Init()
}

func NewBackendProviderService() BackendProviderService {
	return &backendProviderService{
		postgresBackend: postgres.NewPostgresResourceServiceBackend(),
		mysqlBackend:    mysql.NewMysqlResourceServiceBackend(),
		mongoBackend:    mongo.NewMongoResourceServiceBackend(),
	}
}
