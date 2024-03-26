package validate

import (
	"fmt"
	"github.com/apibrew/apibrew/pkg/abs"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/types"
	"github.com/apibrew/apibrew/pkg/util"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/structpb"
	"strconv"
	"strings"
)

func Records(resource abs.ResourceLike, list []unstructured.Unstructured, isUpdate bool) errors.ServiceError {
	var fieldErrors []*model.ErrorField

	var resourcePropertyExists = make(map[string]bool)

	for _, property := range resource.GetProperties() {
		resourcePropertyExists[property.Name] = true
	}

	for _, record := range list {
		for _, property := range resource.GetProperties() {

			packedVal, exists := record.Properties[property.Name]

			if !exists && property.DefaultValue != nil && property.DefaultValue.AsInterface() != nil {
				packedVal = property.DefaultValue
				exists = true
			}

			propertyType := types.ByResourcePropertyType(property.Type)

			if packedVal != nil {
				fieldErrors = append(fieldErrors, Value(resource, property, util.GetRecordId(record), property.Name, packedVal)...)
			}

			var val interface{}
			var err error

			if packedVal == nil {
				val = nil
			} else {
				val, err = propertyType.UnPack(packedVal)

				if err != nil {
					fieldErrors = append(fieldErrors, &model.ErrorField{
						RecordId: util.GetRecordId(record),
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
						RecordId: util.GetRecordId(record),
						Property: property.Name,
						Message:  "required",
						Value:    record.Properties[property.Name],
					})
				}

				if !property.Primary && property.Required && (exists || !isUpdate) {
					fieldErrors = append(fieldErrors, &model.ErrorField{
						RecordId: util.GetRecordId(record),
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
					RecordId: util.GetRecordId(record),
					Property: key,
					Message:  "there are no such property",
				})
			}
		}
	}

	if len(fieldErrors) == 0 {
		return nil
	}

	log.Debug("Record validation errors: ", fieldErrors)

	return errors.RecordValidationError.WithErrorFields(fieldErrors)
}

func Value(resource abs.ResourceLike, property *model.ResourceProperty, recordId string, propertyPath string, value *structpb.Value) []*model.ErrorField {
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
				errorFields = append(errorFields, Value(resource, property.Item, recordId, propertyPath+"."+property.Name+"["+strconv.Itoa(i)+"].", item)...)
			}
			return errorFields
		} else {
			return []*model.ErrorField{{
				RecordId: recordId,
				Property: propertyPath,
				Message:  "value is not list",
				Value:    value,
			}}
		}
	case model.ResourceProperty_MAP:
		if listValue, ok := value.Kind.(*structpb.Value_StructValue); ok {
			var errorFields []*model.ErrorField

			for key, item := range listValue.StructValue.Fields {
				errorFields = append(errorFields, Value(resource, property.Item, recordId, propertyPath+"["+key+"].", item)...)
			}

			return errorFields
		} else {
			return []*model.ErrorField{{
				RecordId: recordId,
				Property: propertyPath,
				Message:  "value is not map",
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
			if enumValue == strings.ToUpper(valStr) {
				return nil
			}
		}

		return []*model.ErrorField{{
			RecordId: recordId,
			Property: propertyPath,
			Message:  fmt.Sprintf("value must be one of enum values: [%s]", strings.Join(property.EnumValues, "|")),
			Value:    value,
		}}

	case model.ResourceProperty_STRUCT:
		var errorFields []*model.ErrorField

		if structValue, ok := value.Kind.(*structpb.Value_StructValue); ok {
			var properties []*model.ResourceProperty

			// locating type
			typeDef := util.LocateArrayElement(resource.GetTypes(), func(elem *model.ResourceSubType) bool {
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

				for _, Item := range properties {
					subType := types.ByResourcePropertyType(Item.Type)
					packedVal, exists := structValue.StructValue.Fields[Item.Name]

					if !exists && Item.DefaultValue != nil && Item.DefaultValue.AsInterface() != nil {
						packedVal = Item.DefaultValue
					}

					var val interface{}
					var err error

					if packedVal == nil {
						val = nil
					} else if _, ok := packedVal.Kind.(*structpb.Value_NullValue); ok {
						return nil
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

					if subType.IsEmpty(val) && Item.Required {
						errorFields = append(errorFields, &model.ErrorField{
							RecordId: recordId,
							Property: propertyPath + Item.Name,
							Message:  "required field is empty",
							Value:    value,
						})
						continue
					}

					errorFields = append(errorFields, Value(resource, Item, recordId, propertyPath+Item.Name, structValue.StructValue.Fields[Item.Name])...)
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
							Message:  "there is no such property",
							Value:    value,
						})
					}
				}
			}
		} else {
			errorFields = append(errorFields, &model.ErrorField{
				RecordId: recordId,
				Property: propertyPath,
				Message:  "value is not struct",
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
	return canCast[float64](typeName, val)
}
