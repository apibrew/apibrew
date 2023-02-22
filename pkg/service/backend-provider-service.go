package service

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/tislib/data-handler/pkg/abs"
	"github.com/tislib/data-handler/pkg/backend/postgres"
	"github.com/tislib/data-handler/pkg/errors"
	"github.com/tislib/data-handler/pkg/logging"
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/resources"
	"github.com/tislib/data-handler/pkg/resources/mapping"
	"github.com/tislib/data-handler/pkg/service/security"
)

type backendProviderService struct {
	systemDataSource *model.DataSource
	backendMap       map[string]abs.Backend
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

func (b *backendProviderService) GetBackendByDataSourceId(ctx context.Context, dataSourceId string) (abs.Backend, errors.ServiceError) {
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
		record, err := b.GetSystemBackend(ctx).GetRecord(systemCtx, resources.DataSourceResource, &abs.Schema{}, dataSourceId)

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
		query, err := PrepareQuery(resources.DataSourceResource, map[string]interface{}{
			"name": dataSourceName,
		})

		records, _, err := b.GetSystemBackend(ctx).ListRecords(systemCtx, abs.ListRecordParams{
			Resource: resources.DataSourceResource,
			Query:    query,
			Limit:    1,
			Schema:   &abs.Schema{},
		})

		if len(records) == 0 {
			return nil, errors.RecordNotFoundError.WithMessage("Data source not found with name: " + dataSourceName)
		}

		var record = records[0]

		if err != nil {
			return nil, err
		}

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
	instance := constructor(dataSource.GetOptions())

	b.backendMap[dataSource.Id] = instance
	b.backendMap[dataSource.Name] = instance

	return instance
}

func (b *backendProviderService) GetBackendConstructor(backend model.DataSourceBackendType) abs.BackendConstructor {
	switch backend {
	case model.DataSourceBackendType_POSTGRESQL:
		return postgres.NewPostgresResourceServiceBackend
	case model.DataSourceBackendType_MYSQL:
		return nil
	case model.DataSourceBackendType_MONGODB:
		return nil
	case model.DataSourceBackendType_CUSTOM:
		return nil
	}

	panic("Not implemented backend: " + backend.String())
}

func (b *backendProviderService) Init(data *model.InitData) {
	b.systemDataSource = data.SystemDataSource
}

func (b *backendProviderService) MigrateResource(resource *model.Resource, schema abs.Schema) {
	if resource.Annotations == nil {
		resource.Annotations = make(map[string]string)
	}

	err := b.GetSystemBackend(context.TODO()).UpgradeResource(context.TODO(), abs.UpgradeResourceParams{
		Resource:       resource,
		ForceMigration: true,
		Schema:         &schema,
	})

	if err != nil {
		panic(err)
	}
}

func NewBackendProviderService() abs.BackendProviderService {
	return &backendProviderService{
		backendMap: make(map[string]abs.Backend),
	}
}
