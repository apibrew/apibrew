package service

import (
	"github.com/tislib/data-handler/pkg/abs"
	"github.com/tislib/data-handler/pkg/errors"
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/service/handler"
	"github.com/tislib/data-handler/pkg/types"
)

type recordService struct {
	ServiceName            string
	resourceService        abs.ResourceService
	genericHandler         *handler.GenericHandler
	backendServiceProvider abs.BackendProviderService
}

func (r *recordService) PrepareQuery(resource *model.Resource, queryMap map[string]interface{}) (*model.BooleanExpression, errors.ServiceError) {
	return PrepareQuery(resource, queryMap)
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

			packedVal, exists := record.Properties[property.Name]
			propertyType := types.ByResourcePropertyType(property.Type)

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

			if property.Primary && isEmpty && isUpdate {
				fieldErrors = append(fieldErrors, &model.ErrorField{
					RecordId: record.Id,
					Property: property.Name,
					Message:  "required",
					Value:    record.Properties[property.Name],
				})
			}

			if !property.Primary && property.Required && isEmpty && (exists || !isUpdate) {
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

	return errors.RecordValidationError.WithErrorFields(fieldErrors)
}

func NewRecordService(resourceService abs.ResourceService, backendProviderService abs.BackendProviderService, genericHandler *handler.GenericHandler) abs.RecordService {
	return &recordService{
		ServiceName:            "RecordService",
		resourceService:        resourceService,
		backendServiceProvider: backendProviderService,
		genericHandler:         genericHandler,
	}
}
