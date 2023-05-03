package redis

import (
	"context"
	"github.com/apibrew/apibrew/pkg/abs"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/util"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
	"time"
)

type redisBackend struct {
	dataSource *model.DataSource
	rdb        *redis.Client
	schema     *abs.Schema
}

func (r *redisBackend) SetSchema(schema *abs.Schema) {
	r.schema = schema
}

func (r redisBackend) GetStatus(ctx context.Context) (connectionAlreadyInitiated bool, testConnection bool, err errors.ServiceError) {
	ping := r.rdb.Ping(ctx)

	log.Print(ping)

	return
}

func (r redisBackend) DestroyDataSource(ctx context.Context) {

}

func (r redisBackend) AddRecords(ctx context.Context, resource *model.Resource, records []*model.Record) ([]*model.Record, errors.ServiceError) {
	for _, record := range records {
		data, err := proto.Marshal(record)

		if err != nil {
			log.Warn(err)

			return nil, r.handleError(err)
		}

		_, err = r.rdb.Set(ctx, r.getKey(resource, record.Id), data, time.Hour*10000).Result()

		if err != nil {
			log.Warn(err)

			return nil, r.handleError(err)
		}
	}

	return records, nil
}

func (r redisBackend) UpdateRecords(ctx context.Context, resource *model.Resource, records []*model.Record) ([]*model.Record, errors.ServiceError) {
	for _, record := range records {
		data, err := proto.Marshal(record)

		if err != nil {
			log.Warn(err)

			return nil, r.handleError(err)
		}

		_, err = r.rdb.Set(ctx, r.getKey(resource, record.Id), data, time.Hour*10000).Result()

		if err != nil {
			log.Warn(err)

			return nil, r.handleError(err)
		}
	}

	return records, nil
}

func (r redisBackend) GetRecord(ctx context.Context, resource *model.Resource, id string) (*model.Record, errors.ServiceError) {
	recData, err := r.rdb.Get(ctx, r.getKey(resource, id)).Bytes()

	if err != nil {
		return nil, r.handleError(err)
	}

	var record = new(model.Record)

	err = proto.Unmarshal(recData, record)

	if err != nil {
		return nil, r.handleError(err)
	}

	return record, nil
}

func (r redisBackend) DeleteRecords(ctx context.Context, resource *model.Resource, list []string) errors.ServiceError {
	_, err := r.rdb.Del(ctx, util.ArrayMap(list, func(item string) string {
		return r.getKey(resource, item)
	})...).Result()

	if err != nil {
		log.Warn(err)

		return errors.InternalError.WithDetails(err.Error())
	}

	return nil
}

func (r redisBackend) ListRecords(ctx context.Context, resource *model.Resource, params abs.ListRecordParams, resultChan chan<- *model.Record) ([]*model.Record, uint32, errors.ServiceError) {
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
	return resource.SourceConfig.Catalog + "/" + resource.SourceConfig.Entity + "/" + recordId
}

func (r redisBackend) handleError(err error) errors.ServiceError {
	if redisErr, ok := err.(redis.Error); ok {
		if redisErr.Error() == "redis: nil" {
			return errors.RecordNotFoundError
		}
	}
	return errors.InternalError.WithDetails(err.Error())
}
