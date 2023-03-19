package mongo

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/tislib/data-handler/pkg/abs"
	"github.com/tislib/data-handler/pkg/errors"
	"github.com/tislib/data-handler/pkg/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"google.golang.org/protobuf/proto"
	"time"
)

type mongoBackend struct {
	dataSource *model.DataSource
	client     *mongo.Client
}

func (r mongoBackend) GetStatus(ctx context.Context) (connectionAlreadyInitiated bool, testConnection bool, err errors.ServiceError) {
	var rp = new(readpref.ReadPref)
	perr := r.client.Ping(ctx, rp)

	if perr != nil {
		log.Error(perr)
		err = errors.InternalError.WithDetails(perr.Error())
	}

	connectionAlreadyInitiated = true
	testConnection = true

	return
}

func (r mongoBackend) DestroyDataSource(ctx context.Context) {
	err := r.client.Disconnect(ctx)

	if err != nil {
		log.Error(err)
	}
}

func (r mongoBackend) AddRecords(ctx context.Context, params abs.BulkRecordsParams) ([]*model.Record, []bool, errors.ServiceError) {
	var inserted []bool
	for _, record := range params.Records {
		r.client.Database()
	}

	return params.Records, inserted, nil
}

func (r mongoBackend) UpdateRecords(ctx context.Context, params abs.BulkRecordsParams) ([]*model.Record, errors.ServiceError) {
	for _, record := range params.Records {
		data, err := proto.Marshal(record)

		if err != nil {
			log.Warn(err)

			return nil, errors.InternalError.WithDetails(err.Error())
		}

		_, err = r.rdb.Set(ctx, r.getKey(params.Resource, record.Id), data, time.Hour*10000).Result()

		if err != nil {
			log.Warn(err)

			return nil, errors.InternalError.WithDetails(err.Error())
		}
	}

	return params.Records, nil
}

func (r mongoBackend) GetRecord(ctx context.Context, resource *model.Resource, schema *abs.Schema, id string) (*model.Record, errors.ServiceError) {
	recData, err := r.rdb.Get(ctx, r.getKey(resource, id)).Bytes()

	if err != nil {
		return nil, errors.InternalError.WithDetails(err.Error())
	}

	var record = new(model.Record)

	err = proto.Unmarshal(recData, record)

	if err != nil {
		return nil, errors.InternalError.WithDetails(err.Error())
	}

	return record, nil
}

func (r mongoBackend) DeleteRecords(ctx context.Context, resource *model.Resource, list []string) errors.ServiceError {
	//TODO implement me
	panic("implement me")
}

func (r mongoBackend) ListRecords(ctx context.Context, params abs.ListRecordParams) ([]*model.Record, uint32, errors.ServiceError) {
	return nil, 0, errors.UnsupportedOperation
}

func (r mongoBackend) ListEntities(ctx context.Context) ([]*model.DataSourceCatalog, errors.ServiceError) {
	return nil, errors.UnsupportedOperation
}

func (r mongoBackend) PrepareResourceFromEntity(ctx context.Context, catalog, entity string) (*model.Resource, errors.ServiceError) {
	return nil, errors.UnsupportedOperation
}

func (r mongoBackend) UpgradeResource(ctx context.Context, params abs.UpgradeResourceParams) errors.ServiceError {
	return nil
}

func (r mongoBackend) BeginTransaction(ctx context.Context, readOnly bool) (transactionKey string, serviceError errors.ServiceError) {
	return "", nil
}

func (r mongoBackend) CommitTransaction(ctx context.Context) (serviceError errors.ServiceError) {
	return nil
}

func (r mongoBackend) RollbackTransaction(ctx context.Context) (serviceError errors.ServiceError) {
	return nil
}

func (r mongoBackend) IsTransactionAlive(ctx context.Context) (isAlive bool, serviceError errors.ServiceError) {
	return true, nil
}

func (r mongoBackend) getKey(resource *model.Resource, recordId string) string {
	return resource.Namespace + "/" + resource.Name + "/" + recordId
}
