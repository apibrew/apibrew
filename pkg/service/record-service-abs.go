package service

import (
	"context"
	"github.com/tislib/data-handler/pkg/errors"
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/service/handler"
	"github.com/tislib/data-handler/pkg/service/params"
	"github.com/tislib/data-handler/pkg/types"
	"github.com/tislib/data-handler/pkg/util"
	"strings"
)

type RecordService interface {
	PrepareQuery(resource *model.Resource, queryMap map[string]interface{}) (*model.BooleanExpression, errors.ServiceError)
	GetRecord(ctx context.Context, namespace, resourceName, id string) (*model.Record, errors.ServiceError)
	FindBy(ctx context.Context, namespace, resourceName, propertyName string, value interface{}) (*model.Record, errors.ServiceError)

	Init(data *model.InitData)
	InjectBackendProviderService(backendProviderService BackendProviderService)
	InjectResourceService(service ResourceService)

	List(ctx context.Context, params params.RecordListParams) ([]*model.Record, uint32, errors.ServiceError)
	Create(ctx context.Context, params params.RecordCreateParams) ([]*model.Record, []bool, errors.ServiceError)
	Update(ctx context.Context, params params.RecordUpdateParams) ([]*model.Record, errors.ServiceError)
	Get(ctx context.Context, params params.RecordGetParams) (*model.Record, errors.ServiceError)
	Delete(ctx context.Context, params params.RecordDeleteParams) errors.ServiceError
	InjectGenericHandler(handler *handler.GenericHandler)
}

type recordService struct {
	ServiceName            string
	resourceService        ResourceService
	genericHandler         *handler.GenericHandler
	backendServiceProvider BackendProviderService
}

func (r *recordService) PrepareQuery(resource *model.Resource, queryMap map[string]interface{}) (*model.BooleanExpression, errors.ServiceError) {
	return PrepareQuery(resource, queryMap)
}

func (r *recordService) InjectBackendProviderService(backendProviderService BackendProviderService) {
	r.backendServiceProvider = backendProviderService
}

func (r *recordService) InjectGenericHandler(genericHandler *handler.GenericHandler) {
	r.genericHandler = genericHandler
}

func (r *recordService) InjectResourceService(service ResourceService) {
	r.resourceService = service
}

func (r *recordService) Init(data *model.InitData) {

}

func (r *recordService) validateRecords(resource *model.Resource, list []*model.Record, isUpdate bool) errors.ServiceError {
	var fieldErrors []*model.ErrorField

	var resourcePropertyExists = make(map[string]bool)

	for _, property := range resource.Properties {
		resourcePropertyExists[property.Name] = true
	}

	for _, record := range list {
		for _, property := range resource.Properties {
			propertyType := types.ByResourcePropertyType(property.Type)
			packedVal, exists := record.Properties[property.Name]

			if packedVal != nil {
				err := propertyType.ValidatePackedValue(packedVal)

				if err != nil {
					fieldErrors = append(fieldErrors, &model.ErrorField{
						RecordId: record.Id,
						Property: property.Name,
						Message:  err.Error(),
						Value:    record.Properties[property.Name],
					})
					continue
				}
			}

			var val interface{}
			var err error

			if packedVal == nil {
				val = nil
			} else {
				val, err = propertyType.UnPack(packedVal)

				if err != nil {
					fieldErrors = append(fieldErrors, &model.ErrorField{
						RecordId: record.Id,
						Property: property.Name,
						Message:  "wrong type: " + err.Error(),
						Value:    record.Properties[property.Name],
					})
					continue
				}
			}

			isEmpty := propertyType.IsEmpty(val)

			if property.Required && isEmpty && (exists || !isUpdate) {
				fieldErrors = append(fieldErrors, &model.ErrorField{
					RecordId: record.Id,
					Property: property.Name,
					Message:  "required",
					Value:    record.Properties[property.Name],
				})
			}
		}

		for key := range record.Properties {
			if !resourcePropertyExists[key] {
				fieldErrors = append(fieldErrors, &model.ErrorField{
					RecordId: record.Id,
					Property: key,
					Message:  "there are no such property",
				})
			}
		}
	}

	if len(fieldErrors) == 0 {
		return nil
	}

	return errors.RecordValidationError.WithDetails("Validation failed on some fields:" + strings.Join(util.ArrayMap[*model.ErrorField, string](fieldErrors, func(fieldError *model.ErrorField) string {
		return fieldError.Property + ":" + fieldError.Message
	}), ";")).WithErrorFields(fieldErrors)
}

func NewRecordService() RecordService {
	return &recordService{ServiceName: "RecordService"}
}
