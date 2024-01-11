package api

import (
	"context"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/formats/unstructured"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resource_model/extramappings"
	"github.com/apibrew/apibrew/pkg/service"
	"github.com/apibrew/apibrew/pkg/util"
)

type api struct {
	container service.Container
}

func (a api) Create(ctx context.Context, record unstructured.Unstructured) (unstructured.Unstructured, errors.ServiceError) {
	return a.Save(ctx, Create, record)
}

func (a api) checkType(record unstructured.Unstructured) errors.ServiceError {
	if _, ok := record["type"]; !ok {
		return errors.RecordValidationError.WithMessage("type field is required")
	}

	if _, ok := record["type"].(string); !ok {
		return errors.RecordValidationError.WithMessage("type field must be string")
	}

	return nil
}

func (a api) Update(ctx context.Context, record unstructured.Unstructured) (unstructured.Unstructured, errors.ServiceError) {
	return a.Save(ctx, Update, record)
}

func (a api) Apply(ctx context.Context, record unstructured.Unstructured) (unstructured.Unstructured, errors.ServiceError) {
	return a.Save(ctx, Apply, record)
}

func (a api) Save(ctx context.Context, saveMode SaveMode, recordObj unstructured.Unstructured) (unstructured.Unstructured, errors.ServiceError) {
	if err := a.checkType(recordObj); err != nil {
		return nil, err
	}

	var resourceIdentity = util.ParseType(recordObj["type"].(string))

	delete(recordObj, "type")

	var record *model.Record
	record, err2 := unstructured.ToRecord(recordObj)

	if err2 != nil {
		return nil, errors.RecordValidationError.WithMessage(err2.Error())
	}

	switch saveMode {
	case Create:
		result, err := a.container.GetRecordService().Create(ctx, service.RecordCreateParams{
			Namespace: resourceIdentity.Namespace,
			Resource:  resourceIdentity.Name,
			Records:   []*model.Record{record},
		})

		if err != nil {
			return nil, err
		}
		record = result[0]
	case Update:
		result, err := a.container.GetRecordService().Update(ctx, service.RecordUpdateParams{
			Namespace: resourceIdentity.Namespace,
			Resource:  resourceIdentity.Name,
			Records:   []*model.Record{record},
		})

		if err != nil {
			return nil, err
		}
		record = result[0]
	case Apply:
		result, err := a.container.GetRecordService().Apply(ctx, service.RecordUpdateParams{
			Namespace: resourceIdentity.Namespace,
			Resource:  resourceIdentity.Name,
			Records:   []*model.Record{record},
		})

		if err != nil {
			return nil, err
		}
		record = result[0]
	}

	processedRecordObj, err2 := unstructured.FromRecord(record)

	if err2 != nil {
		return nil, errors.RecordValidationError.WithMessage(err2.Error())
	}

	return processedRecordObj, nil
}

func (a api) Load(ctx context.Context, recordObj unstructured.Unstructured) (unstructured.Unstructured, errors.ServiceError) {
	if err := a.checkType(recordObj); err != nil {
		return nil, err
	}

	var resourceIdentity = util.ParseType(recordObj["type"].(string))

	delete(recordObj, "type")

	properties, err2 := unstructured.ToProperties(recordObj)

	if err2 != nil {
		return nil, errors.RecordValidationError.WithMessage(err2.Error())
	}

	record, err := a.container.GetRecordService().Load(ctx, resourceIdentity.Namespace, resourceIdentity.Name, properties)

	if err != nil {
		return nil, err
	}

	processedRecordObj, err2 := unstructured.FromRecord(record)

	if err2 != nil {
		return nil, errors.RecordValidationError.WithMessage(err2.Error())
	}

	return processedRecordObj, nil
}

func (a api) Delete(ctx context.Context, recordObj unstructured.Unstructured) errors.ServiceError {
	if err := a.checkType(recordObj); err != nil {
		return err
	}

	var resourceIdentity = util.ParseType(recordObj["type"].(string))

	delete(recordObj, "type")

	if recordObj["id"] == nil {
		var err errors.ServiceError
		recordObj, err = a.Load(ctx, recordObj)

		if err != nil {
			return err
		}
	}

	return a.container.GetRecordService().Delete(ctx, service.RecordDeleteParams{
		Namespace: resourceIdentity.Namespace,
		Resource:  resourceIdentity.Name,
		Ids:       []string{recordObj["id"].(string)},
	})
}

func (a api) List(ctx context.Context, params ListParams) ([]unstructured.Unstructured, uint32, errors.ServiceError) {
	var resourceIdentity = util.ParseType(params.Type)

	var query *model.BooleanExpression

	if params.Query != nil {
		query = extramappings.BooleanExpressionToProto(*params.Query)
	}

	var records []*model.Record

	var aggregation *model.Aggregation

	if params.Aggregation != nil {
		aggregation = &model.Aggregation{
			Items:    []*model.AggregationItem{},
			Grouping: []*model.GroupingItem{},
		}

		for _, item := range params.Aggregation.Items {
			aggregation.Items = append(aggregation.Items, &model.AggregationItem{
				Name:      item.Name,
				Algorithm: model.AggregationItem_Algorithm(model.AggregationItem_Algorithm_value[string(item.Algorithm)]),
				Property:  item.Property,
			})
		}

		for _, item := range params.Aggregation.Grouping {
			aggregation.Grouping = append(aggregation.Grouping, &model.GroupingItem{
				Property: item.Property,
			})
		}
	}

	records, count, err := a.container.GetRecordService().List(ctx, service.RecordListParams{
		Namespace:         resourceIdentity.Namespace,
		Resource:          resourceIdentity.Name,
		Query:             query,
		Limit:             params.Limit,
		Offset:            params.Offset,
		UseHistory:        params.UseHistory,
		ResolveReferences: params.ResolveReferences,
		Filters:           params.Filters,
		Aggregation:       aggregation,
	})

	if err != nil {
		return nil, 0, err
	}

	var result []unstructured.Unstructured

	for _, record := range records {
		recordObj, err2 := unstructured.FromRecord(record)

		if err2 != nil {
			return nil, 0, errors.RecordValidationError.WithMessage(err2.Error())
		}

		result = append(result, recordObj)
	}

	return result, count, nil
}

func NewInterface(container service.Container) Interface {
	return &api{container: container}
}
