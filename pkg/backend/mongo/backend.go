package mongo

import (
	"context"
	"github.com/apibrew/apibrew/pkg/abs"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resource_model"
	"github.com/apibrew/apibrew/pkg/service/annotations"
	"github.com/apibrew/apibrew/pkg/util"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"google.golang.org/protobuf/types/known/structpb"
)

type mongoBackend struct {
	dataSource *resource_model.DataSource
	client     *mongo.Client
	dbName     string
	schema     *abs.Schema
}

func (r *mongoBackend) SetSchema(schema *abs.Schema) {
	r.schema = schema
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

func (r mongoBackend) AddRecords(ctx context.Context, resource *model.Resource, records []*model.Record) ([]*model.Record, errors.ServiceError) {
	var documents []interface{}
	for _, record := range records {
		documents = append(documents, r.recordToDocument(resource, record))
	}
	_, err := r.getCollection(resource).InsertMany(ctx, documents)

	if err != nil {
		return nil, r.handleError(err)
	}

	return records, nil
}

func (r mongoBackend) recordToDocument(resource *model.Resource, record *model.Record) bson.M {
	var data = bson.M{}

	for propName := range resource.Properties {
		val, exists := record.Properties[propName]

		if exists {
			data[propName] = val.AsInterface()
		}
	}

	return data
}

func (r mongoBackend) UpdateRecords(ctx context.Context, resource *model.Resource, records []*model.Record) ([]*model.Record, errors.ServiceError) {
	for _, record := range records {
		var filter = bson.M{}
		var update = bson.M{}
		var set = bson.M{}
		update["$set"] = set

		for propName, prop := range resource.Properties {
			if annotations.IsEnabled(prop, annotations.PrimaryProperty) {
				if record.Properties[propName] == nil {
					filter[propName] = nil
				} else {
					filter[propName] = record.Properties[propName].AsInterface()
				}
			}

			val, exists := record.Properties[propName]

			if exists {
				set[propName] = val.AsInterface()
			}
		}

		_, err := r.getCollection(resource).UpdateOne(ctx, filter, update)

		if err != nil {
			return nil, r.handleError(err)
		}
	}

	return records, nil
}

func (r mongoBackend) GetRecord(ctx context.Context, resource *model.Resource, id string, resolveReferences []string) (*model.Record, errors.ServiceError) {
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

	for propName := range resource.Properties {
		val, exists := (data)[propName]

		if exists {
			st, err := structpb.NewValue(val)

			if err != nil {
				return nil, r.handleError(err)
			}

			record.Properties[propName] = st
		}
	}
	return record, nil
}

func (r mongoBackend) DeleteRecords(ctx context.Context, resource *model.Resource, records []*model.Record) errors.ServiceError {
	var ids = util.ArrayMap(records, func(record *model.Record) string {
		return util.GetRecordId(resource, record)
	})
	for _, item := range ids {
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

func (r mongoBackend) ListRecords(ctx context.Context, resource *model.Resource, params abs.ListRecordParams, _ chan<- *model.Record) ([]*model.Record, uint32, errors.ServiceError) {
	var filter bson.M = nil

	if params.Query != nil {
		filter = r.expressionToMongoFilter(params.Query)
	}

	cursor, err := r.getCollection(resource).Find(ctx, filter)

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

		record, err := r.documentToRecord(resource, *data)

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
