package backend_proxy

import (
	"context"
	"github.com/tislib/apibrew/pkg/abs"
	"github.com/tislib/apibrew/pkg/errors"
	"github.com/tislib/apibrew/pkg/model"
	backend_event_handler "github.com/tislib/apibrew/pkg/service/backend-event-handler"
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
	b.eventHandler.OnAddRecords(ctx, resource, records)

	return b.backend.AddRecords(ctx, resource, records)
}

func (b backendProxy) UpdateRecords(ctx context.Context, resource *model.Resource, records []*model.Record) ([]*model.Record, errors.ServiceError) {
	return b.backend.UpdateRecords(ctx, resource, records)
}

func (b backendProxy) GetRecord(ctx context.Context, resource *model.Resource, id string) (*model.Record, errors.ServiceError) {
	return b.backend.GetRecord(ctx, resource, id)
}

func (b backendProxy) DeleteRecords(ctx context.Context, resource *model.Resource, list []string) errors.ServiceError {
	return b.backend.DeleteRecords(ctx, resource, list)
}

func (b backendProxy) ListRecords(ctx context.Context, resource *model.Resource, params abs.ListRecordParams, resultChan chan<- *model.Record) ([]*model.Record, uint32, errors.ServiceError) {
	return b.backend.ListRecords(ctx, resource, params, resultChan)
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
