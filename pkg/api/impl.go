package api

import (
	"context"
	"github.com/apibrew/apibrew/pkg/abs"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/formats/unstructured"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resource_model"
	"github.com/apibrew/apibrew/pkg/resource_model/extramappings"
	"github.com/apibrew/apibrew/pkg/resources"
	"github.com/apibrew/apibrew/pkg/service"
	"github.com/apibrew/apibrew/pkg/service/validate"
	"github.com/apibrew/apibrew/pkg/util"
	"strings"
)

type api struct {
	recordService   InterfaceRecordService
	resourceService InterfaceResourceService
}

type InterfaceRecordService interface {
	Create(ctx context.Context, params service.RecordCreateParams) ([]abs.RecordLike, error)
	Update(ctx context.Context, params service.RecordUpdateParams) ([]abs.RecordLike, error)
	Apply(ctx context.Context, params service.RecordUpdateParams) ([]abs.RecordLike, error)
	Delete(ctx context.Context, params service.RecordDeleteParams) error
	Load(ctx context.Context, namespace string, name string, properties map[string]interface{}, listParams service.RecordLoadParams) (abs.RecordLike, error)
	List(ctx context.Context, params service.RecordListParams) ([]abs.RecordLike, uint32, error)
}

type InterfaceResourceService interface {
	GetResourceByName(ctx context.Context, namespace, resource string) (*model.Resource, error)
	Create(ctx context.Context, resource *model.Resource, doMigration bool, forceMigration bool) (*model.Resource, error)
	Update(ctx context.Context, resource *model.Resource, doMigration bool, forceMigration bool) error
	Delete(ctx context.Context, ids []string, doMigration bool, forceMigration bool) error
	List(ctx context.Context) ([]*model.Resource, error)
}

func (a api) Create(ctx context.Context, record unstructured.Unstructured) (unstructured.Unstructured, error) {
	return a.save(ctx, Create, record)
}

func (a api) checkType(record unstructured.Unstructured) error {
	if _, ok := record["type"]; !ok {
		return errors.RecordValidationError.WithMessage("type field is required")
	}

	if _, ok := record["type"].(string); !ok {
		return errors.RecordValidationError.WithMessage("type field must be string")
	}

	return nil
}

func (a api) Update(ctx context.Context, record unstructured.Unstructured) (unstructured.Unstructured, error) {
	return a.save(ctx, Update, record)
}

func (a api) Apply(ctx context.Context, record unstructured.Unstructured) (unstructured.Unstructured, error) {
	return a.save(ctx, Apply, record)
}

func (a api) save(ctx context.Context, saveMode SaveMode, recordObj unstructured.Unstructured) (unstructured.Unstructured, error) {
	if err := a.checkType(recordObj); err != nil {
		return nil, err
	}

	var resourceIdentity = util.ParseType(recordObj["type"].(string))

	var record abs.RecordLike
	record, err2 := unstructured.ToRecord(recordObj)

	if err2 != nil {
		return nil, errors.RecordValidationError.WithMessage(err2.Error())
	}

	if resourceIdentity.Name == "resource" || resourceIdentity.Namespace == resources.ResourceResource.Namespace && resourceIdentity.Name == resources.ResourceResource.Name {
		return a.saveResource(ctx, saveMode, recordObj)
	}

	switch saveMode {
	case Create:
		result, err := a.recordService.Create(ctx, service.RecordCreateParams{
			Namespace: resourceIdentity.Namespace,
			Resource:  resourceIdentity.Name,
			Records:   []abs.RecordLike{record},
		})

		if err != nil {
			return nil, err
		}
		record = result[0]
	case Update:
		result, err := a.recordService.Update(ctx, service.RecordUpdateParams{
			Namespace: resourceIdentity.Namespace,
			Resource:  resourceIdentity.Name,
			Records:   []abs.RecordLike{record},
		})

		if err != nil {
			return nil, err
		}
		record = result[0]
	case Apply:
		result, err := a.recordService.Apply(ctx, service.RecordUpdateParams{
			Namespace: resourceIdentity.Namespace,
			Resource:  resourceIdentity.Name,
			Records:   []abs.RecordLike{record},
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

func (a api) Load(ctx context.Context, recordObj unstructured.Unstructured, params LoadParams) (unstructured.Unstructured, error) {
	if err := a.checkType(recordObj); err != nil {
		return nil, err
	}

	var resourceIdentity = util.ParseType(recordObj["type"].(string))

	if resourceIdentity.Namespace == resources.ResourceResource.Namespace && resourceIdentity.Name == resources.ResourceResource.Name {
		return nil, errors.InternalError.WithMessage("Resource load is not supported")
	}

	record, err := a.recordService.Load(ctx, resourceIdentity.Namespace, resourceIdentity.Name, recordObj, service.RecordLoadParams{
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

	processedRecordObj["type"] = resourceIdentity.Type()

	return processedRecordObj, nil
}

func (a api) Delete(ctx context.Context, recordObj unstructured.Unstructured) error {
	if err := a.checkType(recordObj); err != nil {
		return err
	}

	var resourceIdentity = util.ParseType(recordObj["type"].(string))

	if resourceIdentity.Namespace == resources.ResourceResource.Namespace && resourceIdentity.Name == resources.ResourceResource.Name {
		return a.deleteResource(ctx, recordObj)
	}

	if recordObj["id"] == nil {
		var err error
		recordObj, err = a.Load(ctx, recordObj, LoadParams{})

		if err != nil {
			return err
		}
	}

	return a.recordService.Delete(ctx, service.RecordDeleteParams{
		Namespace: resourceIdentity.Namespace,
		Resource:  resourceIdentity.Name,
		Ids:       []string{recordObj["id"].(string)},
	})
}

func (a api) List(ctx context.Context, params ListParams) (RecordListResult, error) {
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

	records, total, err := a.recordService.List(ctx, service.RecordListParams{
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

		recordObj["type"] = resourceIdentity.Type()

		result = append(result, recordObj)
	}

	return RecordListResult{
		Total:   total,
		Content: result,
	}, nil
}

func (a api) saveResource(ctx context.Context, saveMode SaveMode, body unstructured.Unstructured) (unstructured.Unstructured, error) {
	record, err := unstructured.ToRecord(body)

	if err != nil {
		return nil, errors.ResourceValidationError.WithMessage(err.Error())
	}

	if err = validate.Records(resources.ResourceResource, []abs.RecordLike{record}, false); err != nil {
		return nil, errors.ResourceValidationError.WithMessage(err.Error())
	}

	resourceModel := resource_model.ResourceMapperInstance.FromRecord(record)

	resource := extramappings.ResourceFrom(resourceModel)

	switch saveMode {
	case Create:
		result, err := a.resourceService.Create(ctx, resource, true, false)

		if err != nil {
			return nil, err
		}
		resourceModel = extramappings.ResourceTo(result)
	case Update:
		err := a.resourceService.Update(ctx, resource, true, false)

		if err != nil {
			return nil, err
		}
	case Apply:
		result, err := a.resourceService.GetResourceByName(ctx, resource.Namespace, resource.Name)

		if !errors.ResourceNotFoundError.Is(err) && err != nil {
			return nil, err
		}

		if errors.ResourceNotFoundError.Is(err) || result == nil { // create
			result, err = a.resourceService.Create(ctx, resource, true, false)

			if err != nil {
				return nil, err
			}
		} else {
			resource.Id = result.Id
			err = a.resourceService.Update(ctx, resource, true, false)

			if err != nil {
				return nil, err
			}
		}
		resourceModel = extramappings.ResourceTo(result)
	default:
		return nil, errors.InternalError.WithMessage("Unknown save mode")
	}

	processedRecordObj := resource_model.ResourceMapperInstance.ToRecord(resourceModel).MapCopy()

	return processedRecordObj, nil
}

func (a api) GetResourceByType(ctx context.Context, typeName string) (*resource_model.Resource, error) {
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

	resource, err := a.resourceService.GetResourceByName(ctx, namespace, resourceName)

	if err != nil {
		return nil, err
	}

	return extramappings.ResourceTo(resource), nil
}

func (a api) deleteResource(ctx context.Context, obj unstructured.Unstructured) error {
	if obj["id"] == nil {
		return errors.RecordValidationError.WithMessage("id field is required")
	}
	var id, ok = obj["id"].(string)

	if !ok {
		return errors.RecordValidationError.WithMessage("id field must be string")
	}

	return a.resourceService.Delete(ctx, []string{id}, true, false)
}

func (a api) listResource(ctx context.Context, params ListParams) (RecordListResult, error) {
	list, err := a.resourceService.List(ctx)

	if err != nil {
		return RecordListResult{}, err
	}

	var result = make([]unstructured.Unstructured, 0)

	for _, resourceObj := range list {
		resourceUn := extramappings.ResourceTo(resourceObj)

		result = append(result, resource_model.ResourceMapperInstance.ToRecord(resourceUn).MapCopy())
	}

	return RecordListResult{
		Total:   uint32(len(list)),
		Content: result,
	}, nil
}

func NewInterface(container service.Container) Interface {
	return &api{recordService: container.GetRecordService(), resourceService: container.GetResourceService()}
}

func NewInterface2(resourceService InterfaceResourceService, recordService InterfaceRecordService) Interface {
	return &api{resourceService: resourceService, recordService: recordService}
}
