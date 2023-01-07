package service

import (
	"context"
	"data-handler/model"
	"data-handler/service/backend"
	"data-handler/service/backend/mongo"
	"data-handler/service/backend/mysql"
	"data-handler/service/backend/postgres"
)

type BackendProviderService interface {
	Init(data *model.InitData)
	GetSystemBackend(ctx context.Context) backend.ResourceServiceBackend
	GetBackend(ctx context.Context, backend model.DataSourceBackend) backend.ResourceServiceBackend
	InjectDataSourceService(service DataSourceService)
}

type backendProviderService struct {
	systemBackend   model.DataSourceBackend
	postgresBackend backend.ResourceServiceBackend
	mysqlBackend    backend.ResourceServiceBackend
	mongoBackend    backend.ResourceServiceBackend
}

func (b *backendProviderService) InjectDataSourceService(service DataSourceService) {
	b.postgresBackend.InjectDataSourceService(service)
}

func (b *backendProviderService) GetSystemBackend(ctx context.Context) backend.ResourceServiceBackend {
	return b.GetBackend(ctx, b.systemBackend)
}

func (b *backendProviderService) GetBackend(ctx context.Context, backend model.DataSourceBackend) backend.ResourceServiceBackend {
	switch backend {
	case model.DataSourceBackend_POSTGRESQL:
		return b.postgresBackend
	case model.DataSourceBackend_MYSQL:
		return b.mysqlBackend
	case model.DataSourceBackend_MONGODB:
		return b.mongoBackend
	}

	panic("Not implemented backend: " + backend.String())
}

func (b *backendProviderService) Init(data *model.InitData) {
	b.systemBackend = data.SystemDataSource.Backend

	b.postgresBackend.Init()
}

func NewBackendProviderService() BackendProviderService {
	return &backendProviderService{
		postgresBackend: postgres.NewPostgresResourceServiceBackend(),
		mysqlBackend:    mysql.NewMysqlResourceServiceBackend(),
		mongoBackend:    mongo.NewMongoResourceServiceBackend(),
	}
}
