package mongo

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/tislib/apibrew/pkg/abs"
	"github.com/tislib/apibrew/pkg/errors"
	"github.com/tislib/apibrew/pkg/model"
	"github.com/tislib/apibrew/pkg/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"google.golang.org/protobuf/types/known/structpb"
)

type mongoBackend struct {
	dataSource *model.DataSource
	client     *mongo.Client
	dbName     string
}

func (r mongoBackend) handleError(err error) errors.ServiceError {
	if mongo.ErrNoDocuments == err {
		return errors.RecordNotFoundError
	}
	return errors.InternalError.WithDetails(err.Error())
}

func (r mongoBackend) GetStatus(ctx context.Context) (connectionAlreadyInitiated bool, testConnection bool, err errors.ServiceError) {
	var rp = new(readpref.ReadPref)
	perr := r.client.Ping(ctx, rp)

	if perr != nil {
		log.Error(perr)
		err = r.handleError(err)
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
	var documents []interface{}
	for _, record := range params.Records {
		documents = append(documents, r.recordToDocument(params.Resource, record))
	}
	res, err := r.getCollection(params.Resource).InsertMany(ctx, documents)

	if err != nil {
		return nil, nil, r.handleError(err)
	}

	for range res.InsertedIDs {
		inserted = append(inserted, true)
	}

	return params.Records, inserted, nil
}

func (r mongoBackend) recordToDocument(resource *model.Resource, record *model.Record) bson.M {
	var data = bson.M{}

	for _, prop := range resource.Properties {
		val, exists := record.Properties[prop.Name]

		if exists {
			data[prop.Mapping] = val.AsInterface()
		}
	}

	return data
}

func (r mongoBackend) UpdateRecords(ctx context.Context, params abs.BulkRecordsParams) ([]*model.Record, errors.ServiceError) {
	for _, record := range params.Records {
		var filter = bson.M{}
		var update = bson.M{}
		var set = bson.M{}
		update["$set"] = set

		for _, prop := range params.Resource.Properties {
			if prop.Primary {
				if record.Properties[prop.Name] == nil {
					filter[prop.Mapping] = nil
				} else {
					filter[prop.Mapping] = record.Properties[prop.Name].AsInterface()
				}
			}

			val, exists := record.Properties[prop.Name]

			if exists {
				set[prop.Mapping] = val.AsInterface()
			}
		}

		_, err := r.getCollection(params.Resource).UpdateOne(ctx, filter, update)

		if err != nil {
			return nil, r.handleError(err)
		}
	}

	return params.Records, nil
}

func (r mongoBackend) GetRecord(ctx context.Context, resource *model.Resource, schema *abs.Schema, id string) (*model.Record, errors.ServiceError) {
	res := r.getCollection(resource).FindOne(ctx, bson.M{
		"id": id,
	})

	if res.Err() != nil {
		return nil, r.handleError(res.Err())
	}

	var data = new(map[string]interface{})

	err := res.Decode(data)

	if err != nil {
		return nil, r.handleError(err)
	}

	record, err := r.documentToRecord(resource, *data)

	if err != nil {
		return nil, r.handleError(err)
	}

	return record, nil
}

func (r mongoBackend) documentToRecord(resource *model.Resource, data map[string]interface{}) (*model.Record, errors.ServiceError) {
	var record = new(model.Record)
	record.Properties = make(map[string]*structpb.Value)

	for _, prop := range resource.Properties {
		val, exists := (data)[prop.Name]

		if exists {
			st, err := structpb.NewValue(val)

			if err != nil {
				return nil, r.handleError(err)
			}

			record.Properties[prop.Name] = st
		}
	}
	return record, nil
}

func (r mongoBackend) DeleteRecords(ctx context.Context, resource *model.Resource, list []string) errors.ServiceError {
	for _, item := range list {
		var filter = bson.M{
			"id": item,
		}

		_, err := r.getCollection(resource).DeleteOne(ctx, filter)

		if err != nil {
			return r.handleError(err)
		}
	}

	return nil
}

func (r mongoBackend) ListRecords(ctx context.Context, params abs.ListRecordParams) ([]*model.Record, uint32, errors.ServiceError) {
	var filter bson.M = nil

	if params.Query != nil {
		filter = r.expressionToMongoFilter(params.Query)
	}

	cursor, err := r.getCollection(params.Resource).Find(ctx, filter)

	if err != nil {
		return nil, 0, r.handleError(err)
	}

	var records []*model.Record
	for cursor.Next(ctx) {
		var data = new(map[string]interface{})

		err := cursor.Decode(data)

		if err != nil {
			return nil, 0, r.handleError(err)
		}

		record, err := r.documentToRecord(params.Resource, *data)

		if err != nil {
			return nil, 0, r.handleError(err)
		}

		records = append(records, record)
	}

	return records, uint32(len(records)), nil
}

func (r mongoBackend) expressionToMongoFilter(expression *model.BooleanExpression) bson.M {
	var filter = bson.M{}

	switch expr := expression.Expression.(type) {
	case *model.BooleanExpression_And:
		filter["$and"] = util.ArrayMap(expr.And.Expressions, r.expressionToMongoFilter)
	case *model.BooleanExpression_Or:
		filter["$or"] = util.ArrayMap(expr.Or.Expressions, r.expressionToMongoFilter)
	case *model.BooleanExpression_Not:
		filter["$not"] = r.expressionToMongoFilter(expr.Not)
	case *model.BooleanExpression_Equal:
		if propertyExpression, ok := expr.Equal.Left.Expression.(*model.Expression_Property); ok {
			if valueExpression, ok := expr.Equal.Right.Expression.(*model.Expression_Value); ok {
				filter[propertyExpression.Property] = valueExpression.Value.AsInterface()
			}
		}
	}

	return filter
}

func (r mongoBackend) ListEntities(ctx context.Context) ([]*model.DataSourceCatalog, errors.ServiceError) {
	return nil, errors.UnsupportedOperation
}

func (r mongoBackend) PrepareResourceFromEntity(ctx context.Context, catalog, entity string) (*model.Resource, errors.ServiceError) {
	return nil, errors.UnsupportedOperation
}

func (r mongoBackend) UpgradeResource(ctx context.Context, params abs.UpgradeResourceParams) errors.ServiceError {
	for _, step := range params.MigrationPlan.Steps {
		switch step.Kind.(type) {
		case *model.ResourceMigrationStep_CreateResource:
		case *model.ResourceMigrationStep_DeleteResource:
			err := r.getCollection(params.MigrationPlan.CurrentResource).Drop(ctx)

			if err != nil {
				return r.handleError(err)
			}
		}
	}
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

func (r mongoBackend) getCollection(resource *model.Resource) *mongo.Collection {
	return r.client.Database(r.dbName).Collection(resource.SourceConfig.Entity)
}
