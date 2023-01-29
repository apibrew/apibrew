package virtual

import (
	"context"
	"github.com/tislib/data-handler/pkg/backend"
	"github.com/tislib/data-handler/pkg/errors"
	"github.com/tislib/data-handler/pkg/model"
)

func (v virtualBackend) GetStatus(ctx context.Context) (connectionAlreadyInitiated bool, testConnection bool, err errors.ServiceError) {
	return true, true, nil
}

func (v virtualBackend) DestroyDataSource(ctx context.Context) {

}

func (v virtualBackend) AddRecords(ctx context.Context, params backend.BulkRecordsParams) ([]*model.Record, bool, errors.ServiceError) {
	if v.options.Mode == model.VirtualOptions_ERROR {
		return nil, false, errors.LogicalError.WithDetails("Virtual resource tried to be accessed through backend")
	}

	return nil, false, nil
}

func (v virtualBackend) UpdateRecords(ctx context.Context, params backend.BulkRecordsParams) ([]*model.Record, errors.ServiceError) {
	//TODO implement me
	panic("implement me")
}

func (v virtualBackend) GetRecord(ctx context.Context, resource *model.Resource, id string) (*model.Record, errors.ServiceError) {
	if v.options.Mode == model.VirtualOptions_ERROR {
		return nil, errors.LogicalError.WithDetails("Virtual resource tried to be accessed through backend")
	}

	return nil, nil
}

func (v virtualBackend) DeleteRecords(ctx context.Context, resource *model.Resource, list []string) errors.ServiceError {
	//TODO implement me
	panic("implement me")
}

func (v virtualBackend) ListRecords(ctx context.Context, params backend.ListRecordParams) ([]*model.Record, uint32, errors.ServiceError) {
	if v.options.Mode == model.VirtualOptions_ERROR {
		return nil, 0, errors.LogicalError.WithDetails("Virtual resource tried to be accessed through backend")
	}

	return nil, 0, nil
}

func (v virtualBackend) ListEntities(ctx context.Context) ([]string, errors.ServiceError) {
	//TODO implement me
	panic("implement me")
}

func (v virtualBackend) PrepareResourceFromEntity(ctx context.Context, catalog, entity string) (*model.Resource, errors.ServiceError) {
	//TODO implement me
	panic("implement me")
}

func (v virtualBackend) UpgradeResource(ctx context.Context, params backend.UpgradeResourceParams) errors.ServiceError {
	//TODO implement me
	panic("implement me")
}

func (v virtualBackend) DowngradeResource(ctx context.Context, resource *model.Resource, forceMigration bool) errors.ServiceError {
	//TODO implement me
	panic("implement me")
}

func (v virtualBackend) BeginTransaction(ctx context.Context, readOnly bool) (transactionKey string, serviceError errors.ServiceError) {
	//TODO implement me
	panic("implement me")
}

func (v virtualBackend) CommitTransaction(ctx context.Context) (serviceError errors.ServiceError) {
	//TODO implement me
	panic("implement me")
}

func (v virtualBackend) RollbackTransaction(ctx context.Context) (serviceError errors.ServiceError) {
	//TODO implement me
	panic("implement me")
}

func (v virtualBackend) IsTransactionAlive(ctx context.Context) (isAlive bool, serviceError errors.ServiceError) {
	//TODO implement me
	panic("implement me")
}
