package types

import (
	"fmt"
	"github.com/tislib/data-handler/pkg/model"
	"google.golang.org/protobuf/types/known/structpb"
	"reflect"
	"time"
)

type PropertyType interface {
	Pointer(required bool) any
	String(val any) string
	IsEmpty(value any) bool
	ValidatePackedValue(value *structpb.Value) error
	Pack(value interface{}) (*structpb.Value, error)
	UnPack(value *structpb.Value) (interface{}, error)
	Default() any
	Equals(a, b interface{}) bool
}

func Dereference(val interface{}) interface{} {
	if val == nil {
		return nil
	}
	rfVal := reflect.ValueOf(val)

	return dereferenceReflect(rfVal)
}

func dereferenceReflect(rfVal reflect.Value) interface{} {
	if rfVal.Kind() == reflect.Pointer || rfVal.Kind() == reflect.Interface {
		if rfVal.IsNil() {
			return nil
		}
		return dereferenceReflect(rfVal.Elem())
	} else {
		return rfVal.Interface()
	}
}

func ByResourcePropertyType(resourcePropertyType model.ResourcePropertyType) PropertyType {
	switch resourcePropertyType {
	case model.ResourcePropertyType_TYPE_INT32:
		return Int32Type
	case model.ResourcePropertyType_TYPE_INT64:
		return Int64Type
	case model.ResourcePropertyType_TYPE_FLOAT32:
		return Float32Type
	case model.ResourcePropertyType_TYPE_FLOAT64:
		return float64Type{}
	case model.ResourcePropertyType_TYPE_STRING:
		return stringType{}
	case model.ResourcePropertyType_TYPE_UUID:
		return uuidType{}
	case model.ResourcePropertyType_TYPE_DATE:
		return dateType{}
	case model.ResourcePropertyType_TYPE_TIME:
		return timeType{}
	case model.ResourcePropertyType_TYPE_TIMESTAMP:
		return TimestampType
	case model.ResourcePropertyType_TYPE_BOOL:
		return boolType{}
	case model.ResourcePropertyType_TYPE_OBJECT:
		return objectType{}
	case model.ResourcePropertyType_TYPE_REFERENCE:
		return referenceType{}
	case model.ResourcePropertyType_TYPE_ENUM:
		return StringType
	case model.ResourcePropertyType_TYPE_MAP:
		return objectType{}
	case model.ResourcePropertyType_TYPE_LIST:
		return objectType{}
	case model.ResourcePropertyType_TYPE_BYTES:
		return bytesType{}
	default:
		panic("unknown property type: " + resourcePropertyType.String())
	}
}

func GetAllResourcePropertyTypes() []model.ResourcePropertyType {
	var types []model.ResourcePropertyType

	for key := range model.ResourcePropertyType_name {
		types = append(types, model.ResourcePropertyType(key))
	}

	return types
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

func ValidateDateTime(value interface{}) error {
	err := canCast[string]("string", value)

	if err != nil {
		return err
	}

	_, err = time.Parse(time.RFC3339, value.(string))

	return err
}

func ResourcePropertyTypeToJsonSchemaType(resourcePropertyType model.ResourcePropertyType) string {
	switch resourcePropertyType {
	case model.ResourcePropertyType_TYPE_STRING:
		return "string"
	case model.ResourcePropertyType_TYPE_INT64:
		return "number"
	case model.ResourcePropertyType_TYPE_INT32:
		return "number"
	case model.ResourcePropertyType_TYPE_FLOAT64:
		return "number"
	case model.ResourcePropertyType_TYPE_FLOAT32:
		return "number"
	case model.ResourcePropertyType_TYPE_TIMESTAMP:
		return "string"
	case model.ResourcePropertyType_TYPE_TIME:
		return "string"
	case model.ResourcePropertyType_TYPE_DATE:
		return "string"
	case model.ResourcePropertyType_TYPE_UUID:
		return "string"
	case model.ResourcePropertyType_TYPE_ENUM:
		return "string"
	case model.ResourcePropertyType_TYPE_BOOL:
		return "boolean"
	case model.ResourcePropertyType_TYPE_REFERENCE:
		return "object"
	case model.ResourcePropertyType_TYPE_OBJECT:
		return "object"
	case model.ResourcePropertyType_TYPE_MAP:
		return "object"
	case model.ResourcePropertyType_TYPE_LIST:
		return "object"
	default:
		panic("unknown property type: " + resourcePropertyType.String())
	}
}
