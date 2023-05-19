package service

import (
	"fmt"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/types"
	"github.com/apibrew/apibrew/pkg/util"
	"google.golang.org/protobuf/types/known/structpb"
)

func validateRecords(resource *model.Resource, list []*model.Record, isUpdate bool) errors.ServiceError {
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
				err := validatePropertyPackedValue(resource, property, packedVal)

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

func validatePropertyPackedValue(resource *model.Resource, property *model.ResourceProperty, value *structpb.Value) error {
	if value == nil {
		return nil
	}

	if _, ok := value.Kind.(*structpb.Value_NullValue); ok {
		return nil
	}

	propertyType := types.ByResourcePropertyType(property.Type)

	switch property.Type {
	case model.ResourceProperty_BOOL:
		return canCast[bool]("bool", value.AsInterface())

	case model.ResourceProperty_DATE, model.ResourceProperty_TIME, model.ResourceProperty_TIMESTAMP, model.ResourceProperty_BYTES, model.ResourceProperty_UUID:
		// validation of string based types
		err := canCast[string]("string", value.AsInterface())

		if err != nil {
			return err
		}

		_, err = propertyType.UnPack(value)

		return err
	case model.ResourceProperty_STRING:
		return canCast[string]("string", value.AsInterface())
	case model.ResourceProperty_FLOAT32:
		return canCastNumber[float32]("float32", value.AsInterface())
	case model.ResourceProperty_FLOAT64:
		return canCastNumber[float64]("float64", value.AsInterface())
	case model.ResourceProperty_INT32:
		return canCastNumber[int32]("int32", value.AsInterface())
	case model.ResourceProperty_INT64:
		return canCastNumber[int64]("int64", value.AsInterface())
	case model.ResourceProperty_REFERENCE:
		return canCast[map[string]interface{}]("ReferenceType", value.AsInterface())
	case model.ResourceProperty_OBJECT:
		return nil
	case model.ResourceProperty_LIST:
		if listValue, ok := value.Kind.(*structpb.Value_ListValue); ok {
			for _, item := range listValue.ListValue.Values {
				err := validatePropertyPackedValue(resource, property.Item, item)

				if err != nil {
					return err
				}
			}
			return nil
		} else {
			return fmt.Errorf("value is not list: %v", value)
		}
	case model.ResourceProperty_MAP:
		if listValue, ok := value.Kind.(*structpb.Value_StructValue); ok {
			for _, item := range listValue.StructValue.Fields {
				err := validatePropertyPackedValue(resource, property.Item, item)

				if err != nil {
					return err
				}
			}
			return nil
		} else {
			return fmt.Errorf("value is not map: %v", value)
		}
	case model.ResourceProperty_ENUM:
		err := canCast[string]("enum", value.AsInterface())

		if err != nil {
			return err
		}

		valStr := value.GetStringValue()

		for _, enumValue := range property.EnumValues {
			if enumValue.GetStringValue() == valStr {
				return nil
			}
		}

		return fmt.Errorf("value must be one of enum values: %v", value)
	case model.ResourceProperty_STRUCT:
		if structValue, ok := value.Kind.(*structpb.Value_StructValue); ok {
			if property.TypeRef != nil {
				// locating type
				typeDef := util.LocateArrayElement(resource.Types, func(elem *model.ResourceSubType) bool {
					return elem.Name == *property.TypeRef
				})

				if typeDef == nil {
					return fmt.Errorf("type %s not found", *property.TypeRef)
				}

				for _, Item := range typeDef.Properties {
					subType := types.ByResourcePropertyType(Item.Type)
					packedVal := structValue.StructValue.Fields[Item.Name]

					var val interface{}
					var err error
					if packedVal == nil {
						val = nil
					} else {
						val, err = subType.UnPack(packedVal)

						if err != nil {
							return err
						}
					}

					if subType.IsEmpty(val) {
						if Item.Required {
							return fmt.Errorf("required field is empty: %v[%s]", property.Name, Item.Name)
						} else {
							continue
						}
					}

					err = validatePropertyPackedValue(resource, Item, structValue.StructValue.Fields[Item.Name])

					if err != nil {
						return err
					}
				}
			} else {
				for _, Item := range property.Properties {
					subType := types.ByResourcePropertyType(Item.Type)
					packedVal := structValue.StructValue.Fields[Item.Name]

					var val interface{}
					var err error
					if packedVal == nil {
						val = nil
					} else {
						val, err = subType.UnPack(packedVal)

						if err != nil {
							return err
						}
					}

					if subType.IsEmpty(val) {
						if Item.Required {
							return fmt.Errorf("required field is empty: %v[%s]", property.Name, Item.Name)
						} else {
							continue
						}
					}

					err = validatePropertyPackedValue(resource, Item, structValue.StructValue.Fields[Item.Name])

					if err != nil {
						return err
					}
				}

				for key := range structValue.StructValue.Fields {
					found := false
					for _, Item := range property.Properties {
						if Item.Name == key {
							found = true
							break
						}
					}

					if !found {
						return fmt.Errorf("there is no such property: %v", key)
					}
				}
			}

			return nil
		} else {
			return fmt.Errorf("value is not struct: %v", value)
		}
	default:
		panic("unknown property type: " + property.Type.String())
	}
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
