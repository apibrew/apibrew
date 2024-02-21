package api

import (
	"context"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/formats/unstructured"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resource_model"
	"github.com/apibrew/apibrew/pkg/resource_model/extramappings"
	"github.com/apibrew/apibrew/pkg/resources"
	"github.com/apibrew/apibrew/pkg/resources/mapping"
	"github.com/apibrew/apibrew/pkg/service"
	"github.com/apibrew/apibrew/pkg/service/validate"
	"github.com/apibrew/apibrew/pkg/util"
	"strings"
)

type api struct {
	container service.Container
}

func (a api) Create(ctx context.Context, record unstructured.Unstructured) (unstructured.Unstructured, errors.ServiceError) {
	return a.save(ctx, Create, record)
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
	return a.save(ctx, Update, record)
}

func (a api) Apply(ctx context.Context, record unstructured.Unstructured) (unstructured.Unstructured, errors.ServiceError) {
	return a.save(ctx, Apply, record)
}

func (a api) save(ctx context.Context, saveMode SaveMode, recordObj unstructured.Unstructured) (unstructured.Unstructured, errors.ServiceError) {
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

	if resourceIdentity.Namespace == resources.ResourceResource.Namespace && resourceIdentity.Name == resources.ResourceResource.Name {
		return a.saveResource(ctx, saveMode, recordObj)
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

func (a api) Load(ctx context.Context, recordObj unstructured.Unstructured, params LoadParams) (unstructured.Unstructured, errors.ServiceError) {
	if err := a.checkType(recordObj); err != nil {
		return nil, err
	}

	var resourceIdentity = util.ParseType(recordObj["type"].(string))

	delete(recordObj, "type")

	properties, err2 := unstructured.ToProperties(recordObj)

	if err2 != nil {
		return nil, errors.RecordValidationError.WithMessage(err2.Error())
	}

	if resourceIdentity.Namespace == resources.ResourceResource.Namespace && resourceIdentity.Name == resources.ResourceResource.Name {
		return nil, errors.InternalError.WithMessage("Resource load is not supported")
	}

	record, err := a.container.GetRecordService().Load(ctx, resourceIdentity.Namespace, resourceIdentity.Name, properties, service.RecordLoadParams{
		UseHistory:        params.UseHistory,
		ResolveReferences: params.ResolveReferences,
	})

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

	if resourceIdentity.Namespace == resources.ResourceResource.Namespace && resourceIdentity.Name == resources.ResourceResource.Name {
		return a.deleteResource(ctx, recordObj)
	}

	if recordObj["id"] == nil {
		var err errors.ServiceError
		recordObj, err = a.Load(ctx, recordObj, LoadParams{})

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

func (a api) List(ctx context.Context, params ListParams) (RecordListResult, errors.ServiceError) {
	var resourceIdentity = util.ParseType(params.Type)

	var query *model.BooleanExpression

	if params.Query != nil {
		query = extramappings.BooleanExpressionToProto(*params.Query)
	}

	var aggregation *model.Aggregation
	var sorting *model.Sorting

	if resourceIdentity.Namespace == resources.ResourceResource.Namespace && resourceIdentity.Name == resources.ResourceResource.Name {
		return a.listResource(ctx, params)
	}

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

	if len(params.Sorting) > 0 {
		sorting = &model.Sorting{}

		for _, item := range params.Sorting {
			sorting.Items = append(sorting.Items, &model.SortItem{
				Property:  item.Property,
				Direction: model.SortItem_Direction(model.SortItem_Direction_value[strings.ToUpper(string(item.Direction))]),
			})
		}
	}

	records, total, err := a.container.GetRecordService().List(ctx, service.RecordListParams{
		Namespace:         resourceIdentity.Namespace,
		Resource:          resourceIdentity.Name,
		Query:             query,
		Limit:             params.Limit,
		Offset:            params.Offset,
		UseHistory:        params.UseHistory,
		ResolveReferences: params.ResolveReferences,
		Filters:           params.Filters,
		Aggregation:       aggregation,
		Sorting:           sorting,
	})

	if err != nil {
		return RecordListResult{}, err
	}

	var result = make([]unstructured.Unstructured, 0)

	for _, record := range records {
		recordObj, err2 := unstructured.FromRecord(record)

		if err2 != nil {
			return RecordListResult{}, errors.RecordValidationError.WithMessage(err2.Error())
		}

		result = append(result, recordObj)
	}

	return RecordListResult{
		Total:   total,
		Content: result,
	}, nil
}

func (a api) saveResource(ctx context.Context, saveMode SaveMode, body unstructured.Unstructured) (unstructured.Unstructured, errors.ServiceError) {
	record, err := unstructured.ToRecord(body)

	if err != nil {
		return nil, errors.ResourceValidationError.WithMessage(err.Error())
	}

	if err = validate.Records(resources.ResourceResource, []*model.Record{record}, false); err != nil {
		return nil, errors.ResourceValidationError.WithMessage(err.Error())
	}

	resourceModel := resource_model.ResourceMapperInstance.FromRecord(record)

	resource := extramappings.ResourceFrom(resourceModel)

	switch saveMode {
	case Create:
		result, err := a.container.GetResourceService().Create(ctx, resource, true, false)

		if err != nil {
			return nil, err
		}
		record = mapping.ResourceToRecord(result)
	case Update:
		err := a.container.GetResourceService().Update(ctx, resource, true, false)

		if err != nil {
			return nil, err
		}
	case Apply:
		result, err := a.container.GetResourceService().GetResourceByName(ctx, resource.Namespace, resource.Name)

		if !errors.ResourceNotFoundError.Is(err) && err != nil {
			return nil, err
		}

		if errors.ResourceNotFoundError.Is(err) || result == nil { // create
			result, err := a.container.GetResourceService().Create(ctx, resource, true, false)

			if err != nil {
				return nil, err
			}
			record = mapping.ResourceToRecord(result)
		} else {
			err := a.container.GetResourceService().Update(ctx, resource, true, false)

			if err != nil {
				return nil, err
			}
		}
	default:
		return nil, errors.InternalError.WithMessage("Unknown save mode")
	}

	processedRecordObj, err2 := unstructured.FromRecord(record)

	if err2 != nil {
		return nil, errors.RecordValidationError.WithMessage(err2.Error())
	}

	return processedRecordObj, nil
}

func (a api) GetResourceByType(ctx context.Context, typeName string) (*resource_model.Resource, errors.ServiceError) {
	var namespace = "default"
	var resourceName string

	var parts = strings.Split(typeName, "/")

	if len(parts) == 2 {
		namespace = parts[0]
		resourceName = parts[1]
	} else if len(parts) == 1 {
		resourceName = parts[0]
	} else {
		return nil, errors.ResourceValidationError.WithMessage("Invalid resource type")
	}

	resource, err := a.container.GetResourceService().GetResourceByName(ctx, namespace, resourceName)

	if err != nil {
		return nil, err
	}

	return extramappings.ResourceTo(resource), nil
}

func (a api) deleteResource(ctx context.Context, obj unstructured.Unstructured) errors.ServiceError {
	if obj["id"] == nil {
		return errors.RecordValidationError.WithMessage("id field is required")
	}
	var id, ok = obj["id"].(string)

	if !ok {
		return errors.RecordValidationError.WithMessage("id field must be string")
	}

	return a.container.GetResourceService().Delete(ctx, []string{id}, true, false)
}

func (a api) listResource(ctx context.Context, params ListParams) (RecordListResult, errors.ServiceError) {
	list, err := a.container.GetResourceService().List(ctx)

	if err != nil {
		return RecordListResult{}, err
	}

	var result = make([]unstructured.Unstructured, 0)

	for _, resource := range list {
		recordObj, err2 := unstructured.FromRecord(mapping.ResourceToRecord(resource))

		if err2 != nil {
			return RecordListResult{}, errors.RecordValidationError.WithMessage(err2.Error())
		}

		result = append(result, recordObj)
	}

	return RecordListResult{
		Total:   uint32(len(list)),
		Content: result,
	}, nil
}

func NewInterface(container service.Container) Interface {
	return &api{container: container}
}
