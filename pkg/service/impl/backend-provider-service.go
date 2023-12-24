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
	"github.com/apibrew/apibrew/pkg/util"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

type backendProviderService struct {
	systemDataSource *resource_model.DataSource
	backendMap       map[string]abs.Backend
	backendIdMap     map[string]string
	backendNameMap   map[string]string
	schema           *abs.Schema
	eventHandler     backend_event_handler.BackendEventHandler
}

func (b *backendProviderService) SetSchema(schema *abs.Schema) {
	b.schema = schema
}

func (b *backendProviderService) DestroyBackend(ctx context.Context, dataSourceId string) errors.ServiceError {
	bck, err := b.getBackendByDataSourceId(ctx, dataSourceId)

	if err != nil {
		return err
	}

	bck.DestroyDataSource(ctx)

	delete(b.backendMap, b.backendIdMap[dataSourceId])
	delete(b.backendMap, dataSourceId)
	delete(b.backendIdMap, dataSourceId)
	delete(b.backendNameMap, b.backendIdMap[dataSourceId])

	return nil
}

func (b *backendProviderService) getBackendByDataSourceId(ctx context.Context, dataSourceId string) (abs.Backend, errors.ServiceError) {
	if b.backendMap[dataSourceId] != nil {
		return b.backendMap[dataSourceId], nil
	}

	if dataSourceId == b.systemDataSource.Id.String() {
		return b.getSystemBackend(ctx), nil
	} else {
		systemCtx := util.WithSystemContext(context.TODO())
		record, err := b.getSystemBackend(ctx).GetRecord(systemCtx, resources.DataSourceResource, dataSourceId, nil)

		if err != nil {
			return nil, err
		}

		return b.getBackend(resource_model.DataSourceMapperInstance.FromRecord(record)), nil
	}
}

func (b *backendProviderService) getBackendByResource(ctx context.Context, resource *model.Resource) (abs.Backend, errors.ServiceError) {
	if resource.Virtual {
		return nil, errors.LogicalError.WithMessage("Cannot get backend for virtual resource")
	}

	return b.getBackendByDataSourceName(ctx, resource.SourceConfig.DataSource)
}

func (b *backendProviderService) getBackendByDataSourceName(ctx context.Context, dataSourceName string) (abs.Backend, errors.ServiceError) {
	if b.backendMap[dataSourceName] != nil {
		return b.backendMap[dataSourceName], nil
	}

	logger := log.WithFields(logging.CtxFields(ctx))
	logger.WithField("dataSourceName", dataSourceName).Debug("Begin data-source GetDataSourceBackendById")
	defer logger.Debug("End data-source GetDataSourceBackendById")

	if dataSourceName == b.systemDataSource.Name {
		return b.getSystemBackend(ctx), nil
	} else {
		systemCtx := util.WithSystemContext(context.TODO())
		query, err := util.PrepareQuery(resources.DataSourceResource, map[string]string{
			"name": dataSourceName,
		})

		if err != nil {
			return nil, err
		}

		records, _, err := b.getSystemBackend(ctx).ListRecords(systemCtx, resources.DataSourceResource, abs.ListRecordParams{
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

		return b.getBackend(resource_model.DataSourceMapperInstance.FromRecord(record)), nil
	}
}

func (b *backendProviderService) getSystemBackend(_ context.Context) abs.Backend {
	return b.getBackend(b.systemDataSource)
}

func (b *backendProviderService) getBackend(dataSource *resource_model.DataSource) abs.Backend {
	if b.backendMap[dataSource.Id.String()] != nil {
		return b.backendMap[dataSource.Id.String()]
	}

	constructor := b.getBackendConstructor(dataSource.GetBackend())
	instance := constructor(dataSource)
	instance.SetSchema(b.schema)

	b.backendMap[dataSource.Id.String()] = instance
	b.backendIdMap[dataSource.Id.String()] = dataSource.Name
	b.backendNameMap[dataSource.Name] = dataSource.Id.String()
	b.backendMap[dataSource.Name] = instance

	return instance
}

func (b *backendProviderService) getBackendConstructor(backend resource_model.DataSourceBackend) abs.BackendConstructor {
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

	b.eventHandler.RegisterHandler(b.prepareActualHandler())
}

func (b *backendProviderService) prepareActualHandler() backend_event_handler.Handler {
	return backend_event_handler.Handler{
		Id:   "actualHandler",
		Name: "actualHandler",
		Fn:   b.actualHandlerFn,
		Selector: &model.EventSelector{
			Actions: []model.Event_Action{
				model.Event_CREATE,
				model.Event_UPDATE,
				model.Event_DELETE,
				model.Event_GET,
				model.Event_LIST,
				// OPERATE is not allowed to be handled by actualHandler
			},
		},
		Order:     backend_event_handler.NaturalOrder,
		Finalizes: false,
		Sync:      true,
		Responds:  true,
		Internal:  true,
	}
}

func (b *backendProviderService) actualHandlerFn(ctx context.Context, event *model.Event) (*model.Event, errors.ServiceError) {
	// if resource is virtual, do not handle it
	if event.Resource.Virtual {
		return event, nil
	}

	bck, err := b.getBackendByResource(ctx, event.Resource)

	if err != nil {
		return nil, err
	}

	switch event.Action {
	case model.Event_CREATE:
		result, err := bck.AddRecords(ctx, event.Resource, event.Records)

		event.Records = result

		return event, err
	case model.Event_UPDATE:
		result, err := bck.UpdateRecords(ctx, event.Resource, event.Records)

		event.Records = result

		return event, err
	case model.Event_GET:
		for i, record := range event.Records {
			result, err := bck.GetRecord(ctx, event.Resource, record.Properties["id"].GetStringValue(), event.RecordSearchParams.ResolveReferences)

			if err != nil {
				return nil, err
			}

			event.Records[i] = result
		}

		return event, err
	case model.Event_DELETE:
		err = bck.DeleteRecords(ctx, event.Resource, event.Records)
		return event, err
	case model.Event_LIST:
		result, total, err := bck.ListRecords(ctx, event.Resource, abs.ListRecordParams{
			Query:             event.RecordSearchParams.Query,
			Limit:             event.RecordSearchParams.Limit,
			Offset:            event.RecordSearchParams.Offset,
			ResolveReferences: event.RecordSearchParams.ResolveReferences,
			Aggregation:       event.RecordSearchParams.Aggregation,
		}, nil)

		event.Records = result
		event.Total = uint64(total)

		return event, err
	default:
		return nil, errors.InternalError.WithDetails("Unknown action: " + event.Action.String())
	}
}

func NewBackendProviderService(eventHandler backend_event_handler.BackendEventHandler) service.BackendProviderService {
	return &backendProviderService{
		backendMap:     make(map[string]abs.Backend),
		backendIdMap:   make(map[string]string),
		backendNameMap: make(map[string]string),
		eventHandler:   eventHandler,
	}
}
