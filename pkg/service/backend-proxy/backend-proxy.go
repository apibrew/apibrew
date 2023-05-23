package backend_proxy

import (
	"context"
	"github.com/apibrew/apibrew/pkg/abs"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/model"
	backend_event_handler "github.com/apibrew/apibrew/pkg/service/backend-event-handler"
)

type BackendProxy interface {
	abs.Backend
}

type backendProxy struct {
	backend      abs.Backend
	eventHandler backend_event_handler.BackendEventHandler
}

func (b backendProxy) SetSchema(schema *abs.Schema) {
	b.backend.SetSchema(schema)
}

func (b backendProxy) GetStatus(ctx context.Context) (connectionAlreadyInitiated bool, testConnection bool, err errors.ServiceError) {
	return b.backend.GetStatus(ctx)
}

func (b backendProxy) DestroyDataSource(ctx context.Context) {
	b.backend.DestroyDataSource(ctx)
}

func (b backendProxy) AddRecords(ctx context.Context, resource *model.Resource, records []*model.Record) ([]*model.Record, errors.ServiceError) {
	endEvent, err := b.eventHandler.HandleInternalOperation(ctx, b.eventHandler.PrepareInternalEvent(ctx, &model.Event{
		Action:   model.Event_CREATE,
		Resource: resource,
		Records:  records,
	}),
		func(ctx context.Context, passedEvent *model.Event) (*model.Event, errors.ServiceError) {
			result, err := b.backend.AddRecords(ctx, resource, passedEvent.Records)

			passedEvent.Records = result

			return passedEvent, err
		})

	if err != nil {
		return nil, err
	}

	return endEvent.Records, nil
}

func (b backendProxy) UpdateRecords(ctx context.Context, resource *model.Resource, records []*model.Record) ([]*model.Record, errors.ServiceError) {
	endEvent, err := b.eventHandler.HandleInternalOperation(ctx, b.eventHandler.PrepareInternalEvent(ctx, &model.Event{
		Action:   model.Event_UPDATE,
		Resource: resource,
		Records:  records,
	}),
		func(ctx context.Context, passedEvent *model.Event) (*model.Event, errors.ServiceError) {
			result, err := b.backend.UpdateRecords(ctx, resource, passedEvent.Records)

			passedEvent.Records = result

			return passedEvent, err
		})

	if err != nil {
		return nil, err
	}

	return endEvent.Records, nil
}

func (b backendProxy) GetRecord(ctx context.Context, resource *model.Resource, id string) (*model.Record, errors.ServiceError) {
	endEvent, err := b.eventHandler.HandleInternalOperation(ctx, b.eventHandler.PrepareInternalEvent(ctx, &model.Event{
		Action:   model.Event_GET,
		Resource: resource,
		Ids:      []string{id},
	}),
		func(ctx context.Context, passedEvent *model.Event) (*model.Event, errors.ServiceError) {
			result, err := b.backend.GetRecord(ctx, resource, id)

			passedEvent.Records = []*model.Record{result}

			return passedEvent, err
		})

	if err != nil {
		return nil, err
	}

	if len(endEvent.Records) == 0 {
		return nil, nil
	}

	return endEvent.Records[0], nil
}

func (b backendProxy) DeleteRecords(ctx context.Context, resource *model.Resource, list []string) errors.ServiceError {
	_, err := b.eventHandler.HandleInternalOperation(ctx, b.eventHandler.PrepareInternalEvent(ctx, &model.Event{
		Action:   model.Event_DELETE,
		Resource: resource,
		Ids:      list,
	}),
		func(ctx context.Context, passedEvent *model.Event) (*model.Event, errors.ServiceError) {
			err := b.backend.DeleteRecords(ctx, resource, passedEvent.Ids)

			return passedEvent, err
		})

	return err
}

func (b backendProxy) ListRecords(ctx context.Context, resource *model.Resource, params abs.ListRecordParams, resultChan chan<- *model.Record) ([]*model.Record, uint32, errors.ServiceError) {
	var total uint32
	endEvent, err := b.eventHandler.HandleInternalOperation(ctx, b.eventHandler.PrepareInternalEvent(ctx, &model.Event{
		Action:   model.Event_LIST,
		Resource: resource,
		RecordSearchParams: &model.Event_RecordSearchParams{
			Query:             params.Query,
			Limit:             params.Limit,
			Offset:            params.Offset,
			ResolveReferences: params.ResolveReferences,
		},
	}),
		func(ctx context.Context, passedEvent *model.Event) (*model.Event, errors.ServiceError) {
			result, localTotal, err := b.backend.ListRecords(ctx, resource, params, resultChan)

			if localTotal != 0 {
				total = localTotal
			}

			passedEvent.Records = result

			return passedEvent, err
		})

	if err != nil {
		return nil, total, err
	}

	if total == 0 {
		total = uint32(len(endEvent.Records))
	}

	return endEvent.Records, total, nil
}

func (b backendProxy) ListEntities(ctx context.Context) ([]*model.DataSourceCatalog, errors.ServiceError) {
	return b.backend.ListEntities(ctx)
}

func (b backendProxy) PrepareResourceFromEntity(ctx context.Context, catalog, entity string) (*model.Resource, errors.ServiceError) {
	return b.backend.PrepareResourceFromEntity(ctx, catalog, entity)
}

func (b backendProxy) UpgradeResource(ctx context.Context, params abs.UpgradeResourceParams) errors.ServiceError {
	return b.backend.UpgradeResource(ctx, params)
}

func (b backendProxy) BeginTransaction(ctx context.Context, readOnly bool) (transactionKey string, serviceError errors.ServiceError) {
	return b.backend.BeginTransaction(ctx, readOnly)
}

func (b backendProxy) CommitTransaction(ctx context.Context) (serviceError errors.ServiceError) {
	return b.backend.CommitTransaction(ctx)
}

func (b backendProxy) RollbackTransaction(ctx context.Context) (serviceError errors.ServiceError) {
	return b.backend.RollbackTransaction(ctx)
}

func (b backendProxy) IsTransactionAlive(ctx context.Context) (isAlive bool, serviceError errors.ServiceError) {
	return b.backend.IsTransactionAlive(ctx)
}

func NewBackendProxy(backend abs.Backend, eventHandler backend_event_handler.BackendEventHandler) BackendProxy {
	return &backendProxy{backend: backend, eventHandler: eventHandler}
}
