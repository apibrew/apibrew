package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
	"github.com/tislib/data-handler/pkg/abs"
	"github.com/tislib/data-handler/pkg/errors"
	"github.com/tislib/data-handler/pkg/model"
	"google.golang.org/protobuf/proto"
	"time"
)

type redisBackend struct {
	dataSource *model.DataSource
	rdb        *redis.Client
}

func (r redisBackend) GetStatus(ctx context.Context) (connectionAlreadyInitiated bool, testConnection bool, err errors.ServiceError) {
	ping := r.rdb.Ping(ctx)

	log.Print(ping)

	return
}

func (r redisBackend) DestroyDataSource(ctx context.Context) {

}

func (r redisBackend) AddRecords(ctx context.Context, params abs.BulkRecordsParams) ([]*model.Record, []bool, errors.ServiceError) {
	var inserted []bool
	for _, record := range params.Records {
		data, err := proto.Marshal(record)

		if err != nil {
			log.Warn(err)

			return nil, nil, errors.InternalError.WithDetails(err.Error())
		}

		_, err = r.rdb.Set(ctx, r.getKey(params.Resource, record.Id), data, time.Hour*10000).Result()

		if err != nil {
			log.Warn(err)

			return nil, nil, errors.InternalError.WithDetails(err.Error())
		}
		inserted = append(inserted, true)
	}

	return params.Records, inserted, nil
}

func (r redisBackend) UpdateRecords(ctx context.Context, params abs.BulkRecordsParams) ([]*model.Record, errors.ServiceError) {
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

func (r redisBackend) GetRecord(ctx context.Context, resource *model.Resource, schema *abs.Schema, id string) (*model.Record, errors.ServiceError) {
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

func (r redisBackend) DeleteRecords(ctx context.Context, resource *model.Resource, list []string) errors.ServiceError {
	//TODO implement me
	panic("implement me")
}

func (r redisBackend) ListRecords(ctx context.Context, params abs.ListRecordParams) ([]*model.Record, uint32, errors.ServiceError) {
	return nil, 0, errors.UnsupportedOperation
}

func (r redisBackend) ListEntities(ctx context.Context) ([]*model.DataSourceCatalog, errors.ServiceError) {
	return nil, errors.UnsupportedOperation
}

func (r redisBackend) PrepareResourceFromEntity(ctx context.Context, catalog, entity string) (*model.Resource, errors.ServiceError) {
	return nil, errors.UnsupportedOperation
}

func (r redisBackend) UpgradeResource(ctx context.Context, params abs.UpgradeResourceParams) errors.ServiceError {
	return nil
}

func (r redisBackend) BeginTransaction(ctx context.Context, readOnly bool) (transactionKey string, serviceError errors.ServiceError) {
	return "", nil
}

func (r redisBackend) CommitTransaction(ctx context.Context) (serviceError errors.ServiceError) {
	return nil
}

func (r redisBackend) RollbackTransaction(ctx context.Context) (serviceError errors.ServiceError) {
	return nil
}

func (r redisBackend) IsTransactionAlive(ctx context.Context) (isAlive bool, serviceError errors.ServiceError) {
	return true, nil
}

func (r redisBackend) getKey(resource *model.Resource, recordId string) string {
	return resource.Namespace + "/" + resource.Name + "/" + recordId
}
