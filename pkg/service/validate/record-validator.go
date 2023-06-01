package validate

import (
	"fmt"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/types"
	"github.com/apibrew/apibrew/pkg/util"
	"google.golang.org/protobuf/types/known/structpb"
	"strconv"
)

func ValidateRecords(resource *model.Resource, list []*model.Record, isUpdate bool) errors.ServiceError {
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
				fieldErrors = append(fieldErrors, ValidatePropertyPackedValue(resource, property, record.Id, "", packedVal)...)
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

			if isEmpty {
				if property.Primary && isUpdate {
					fieldErrors = append(fieldErrors, &model.ErrorField{
						RecordId: record.Id,
						Property: property.Name,
						Message:  "required",
						Value:    record.Properties[property.Name],
					})
				}

				if !property.Primary && property.Required && (exists || !isUpdate) {
					fieldErrors = append(fieldErrors, &model.ErrorField{
						RecordId: record.Id,
						Property: property.Name,
						Message:  "required",
						Value:    record.Properties[property.Name],
					})
				}
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

func ValidatePropertyPackedValue(resource *model.Resource, property *model.ResourceProperty, recordId string, propertyPath string, value *structpb.Value) []*model.ErrorField {
	if value == nil {
		return nil
	}

	if _, ok := value.Kind.(*structpb.Value_NullValue); ok {
		return nil
	}

	propertyType := types.ByResourcePropertyType(property.Type)

	var err error

	// validating simple fields:

	switch property.Type {
	case model.ResourceProperty_BOOL:
		err = canCast[bool]("bool", value.AsInterface())
	case model.ResourceProperty_DATE, model.ResourceProperty_TIME, model.ResourceProperty_TIMESTAMP, model.ResourceProperty_BYTES, model.ResourceProperty_UUID:
		// validation of string based types
		err = canCast[string]("string", value.AsInterface())

		if err != nil {
			break
		}

		_, err = propertyType.UnPack(value)
	case model.ResourceProperty_STRING:
		err = canCast[string]("string", value.AsInterface())
	case model.ResourceProperty_FLOAT32:
		err = canCastNumber[float32]("float32", value.AsInterface())
	case model.ResourceProperty_FLOAT64:
		err = canCastNumber[float64]("float64", value.AsInterface())
	case model.ResourceProperty_INT32:
		err = canCastNumber[int32]("int32", value.AsInterface())
	case model.ResourceProperty_INT64:
		err = canCastNumber[int64]("int64", value.AsInterface())
	case model.ResourceProperty_REFERENCE:
		err = canCast[map[string]interface{}]("ReferenceType", value.AsInterface())
	case model.ResourceProperty_OBJECT:
		return nil
	}

	if err != nil {
		return []*model.ErrorField{{
			RecordId: recordId,
			Property: propertyPath,
			Message:  err.Error(),
			Value:    value,
		}}
	}

	// validating complex fields:
	switch property.Type {
	case model.ResourceProperty_LIST:

		if listValue, ok := value.Kind.(*structpb.Value_ListValue); ok {
			var errorFields []*model.ErrorField

			for i, item := range listValue.ListValue.Values {
				errorFields = append(errorFields, ValidatePropertyPackedValue(resource, property.Item, recordId, propertyPath+property.Name+"["+strconv.Itoa(i)+"].", item)...)
			}
			return errorFields
		} else {
			return []*model.ErrorField{{
				RecordId: recordId,
				Property: propertyPath,
				Message:  fmt.Sprintf("value is not list: %v", value),
				Value:    value,
			}}
		}
	case model.ResourceProperty_MAP:
		if listValue, ok := value.Kind.(*structpb.Value_StructValue); ok {
			var errorFields []*model.ErrorField

			for key, item := range listValue.StructValue.Fields {
				errorFields = append(errorFields, ValidatePropertyPackedValue(resource, property.Item, recordId, propertyPath+property.Name+"["+key+"].", item)...)
			}

			return errorFields
		} else {
			return []*model.ErrorField{{
				RecordId: recordId,
				Property: propertyPath,
				Message:  fmt.Sprintf("value is not map: %v", value),
				Value:    value,
			}}
		}
	case model.ResourceProperty_ENUM:
		err = canCast[string]("enum", value.AsInterface())

		if err != nil {
			return []*model.ErrorField{{
				RecordId: recordId,
				Property: propertyPath,
				Message:  err.Error(),
				Value:    value,
			}}
		}

		valStr := value.GetStringValue()

		for _, enumValue := range property.EnumValues {
			if enumValue == valStr {
				return nil
			}
		}

		return []*model.ErrorField{{
			RecordId: recordId,
			Property: propertyPath,
			Message:  fmt.Sprintf("value must be one of enum values: %v", value),
			Value:    value,
		}}

	case model.ResourceProperty_STRUCT:
		var errorFields []*model.ErrorField

		if structValue, ok := value.Kind.(*structpb.Value_StructValue); ok {
			var properties = property.Properties
			if property.TypeRef != nil {
				// locating type
				typeDef := util.LocateArrayElement(resource.Types, func(elem *model.ResourceSubType) bool {
					return elem.Name == *property.TypeRef
				})

				if typeDef == nil {
					errorFields = append(errorFields, &model.ErrorField{
						RecordId: recordId,
						Property: propertyPath,
						Message:  fmt.Sprintf("type %s not found", *property.TypeRef),
						Value:    value,
					})
				} else {
					properties = typeDef.Properties
				}
			}

			for _, Item := range properties {
				subType := types.ByResourcePropertyType(Item.Type)
				packedVal := structValue.StructValue.Fields[Item.Name]

				var val interface{}
				var err error
				if packedVal == nil {
					val = nil
				} else {
					val, err = subType.UnPack(packedVal)

					if err != nil {
						errorFields = append(errorFields, &model.ErrorField{
							RecordId: recordId,
							Property: propertyPath + Item.Name,
							Message:  err.Error(),
							Value:    value,
						})
						continue
					}
				}

				if subType.IsEmpty(val) {
					if Item.Required {
						errorFields = append(errorFields, &model.ErrorField{
							RecordId: recordId,
							Property: propertyPath + Item.Name,
							Message:  fmt.Sprintf("required field is empty: %v[%s]", Item.Name, Item.Name),
							Value:    value,
						})
						continue
					}
				}

				errorFields = append(errorFields, ValidatePropertyPackedValue(resource, Item, recordId, propertyPath+Item.Name+".", structValue.StructValue.Fields[Item.Name])...)
			}

			for key := range structValue.StructValue.Fields {
				found := false
				for _, Item := range properties {
					if Item.Name == key {
						found = true
						break
					}
				}

				if !found {
					errorFields = append(errorFields, &model.ErrorField{
						RecordId: recordId,
						Property: propertyPath,
						Message:  fmt.Sprintf("there is no such property: %v", key),
						Value:    value,
					})
				}
			}
		} else {
			errorFields = append(errorFields, &model.ErrorField{
				RecordId: recordId,
				Property: propertyPath,
				Message:  fmt.Sprintf("value is not struct: %v", value),
				Value:    value,
			})
		}

		return errorFields
	}

	return nil
}

func canCast[T interface{}](typeName string, val interface{}) error {
	if _, ok := val.(T); ok {
		return nil
	} else {
		return fmt.Errorf("value is not %s: %v", typeName, val)
	}

}

type number interface {
	float64 | float32 | int64 | int32 | int | int8 | uint64 | uint32 | uint8
}

func canCastNumber[T number](typeName string, val interface{}) error {
	if val == T(0) {
		return nil
	}
	err := canCast[float64](typeName, val)

	if err != nil {
		return err
	}

	castedValue := float64(T(val.(float64)))
	if val.(float64)-castedValue > 0.000001 {
		return fmt.Errorf("value is not in type %s: %v", typeName, val)
	}

	return nil
}
