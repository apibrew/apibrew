package service

import (
	"context"
	"data-handler/logging"
	"data-handler/model"
	"data-handler/service/backend"
	"data-handler/service/backend/postgres"
	"data-handler/service/errors"
	"data-handler/service/mapping"
	"data-handler/service/security"
	"data-handler/service/system"
	log "github.com/sirupsen/logrus"
)

type BackendProviderService interface {
	Init(data *model.InitData)
	GetSystemBackend(ctx context.Context) backend.Backend
	GetBackendByDataSourceId(ctx context.Context, dataSourceId string) (backend.Backend, errors.ServiceError)
	DestroyBackend(ctx context.Context, id string) error
}

type backendProviderService struct {
	systemDataSource *model.DataSource
	backendMap       map[string]backend.Backend
}

func (b *backendProviderService) DestroyBackend(ctx context.Context, dataSourceId string) error {
	bck, err := b.GetBackendByDataSourceId(ctx, dataSourceId)

	if err != nil {
		return err
	}

	bck.DestroyDataSource(ctx)

	delete(b.backendMap, dataSourceId)

	return nil
}

func (b *backendProviderService) GetBackendByDataSourceId(ctx context.Context, dataSourceId string) (backend.Backend, errors.ServiceError) {
	if b.backendMap[dataSourceId] != nil {
		return b.backendMap[dataSourceId], nil
	}

	logger := log.WithFields(logging.CtxFields(ctx))
	logger.WithField("dataSourceId", dataSourceId).Debug("Begin data-source GetDataSourceBackendById")
	defer logger.Debug("End data-source GetDataSourceBackendById")

	if dataSourceId == b.systemDataSource.Id {
		return b.GetSystemBackend(ctx), nil
	} else {
		systemCtx := security.WithSystemContext(context.TODO())
		record, err := b.GetSystemBackend(ctx).GetRecord(systemCtx, system.DataSourceResource, dataSourceId)

		if err != nil {
			return nil, err
		}

		return b.GetBackend(mapping.DataSourceFromRecord(record)), nil
	}
}

func (b *backendProviderService) GetSystemBackend(ctx context.Context) backend.Backend {
	return b.GetBackend(b.systemDataSource)
}

func (b *backendProviderService) GetBackend(dataSource *model.DataSource) backend.Backend {
	if b.backendMap[dataSource.Id] != nil {
		return b.backendMap[dataSource.Id]
	}

	constructor := b.GetBackendConstructor(dataSource.GetBackend())
	instance := constructor(dataSource.GetOptions())

	b.backendMap[dataSource.Id] = instance

	return instance
}

func (b *backendProviderService) GetBackendConstructor(backend model.DataSourceBackendType) backend.Constructor {
	switch backend {
	case model.DataSourceBackendType_POSTGRESQL:
		return postgres.NewPostgresResourceServiceBackend
	case model.DataSourceBackendType_MYSQL:
		return nil
	case model.DataSourceBackendType_MONGODB:
		return nil
	}

	panic("Not implemented backend: " + backend.String())
}

func (b *backendProviderService) Init(data *model.InitData) {
	b.systemDataSource = data.SystemDataSource
}

func NewBackendProviderService() BackendProviderService {
	return &backendProviderService{
		backendMap: make(map[string]backend.Backend),
	}
}

//
//func (p *postgresResourceServiceBackend) Init() {
//	err := p.withBackend(context.TODO(), p.systemBackend.GetDataSourceId(), false, func(tx *sql.Tx) errors.ServiceError {
//		return resourceSetupTables(tx)
//	})
//
//	if err != nil {
//		panic(err)
//	}
//}
