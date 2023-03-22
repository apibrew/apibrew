package types

import (
	"fmt"
	"github.com/getkin/kin-openapi/openapi3"
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

func ByResourcePropertyType(resourcePropertyType model.ResourceProperty_Type) PropertyType {
	switch resourcePropertyType {
	case model.ResourceProperty_INT32:
		return Int32Type
	case model.ResourceProperty_INT64:
		return Int64Type
	case model.ResourceProperty_FLOAT32:
		return Float32Type
	case model.ResourceProperty_FLOAT64:
		return float64Type{}
	case model.ResourceProperty_STRING:
		return stringType{}
	case model.ResourceProperty_UUID:
		return uuidType{}
	case model.ResourceProperty_DATE:
		return dateType{}
	case model.ResourceProperty_TIME:
		return timeType{}
	case model.ResourceProperty_TIMESTAMP:
		return TimestampType
	case model.ResourceProperty_BOOL:
		return boolType{}
	case model.ResourceProperty_OBJECT:
		return objectType{}
	case model.ResourceProperty_REFERENCE:
		return referenceType{}
	case model.ResourceProperty_ENUM:
		return StringType
	case model.ResourceProperty_MAP:
		return objectType{}
	case model.ResourceProperty_LIST:
		return objectType{}
	case model.ResourceProperty_BYTES:
		return bytesType{}
	default:
		panic("unknown property type: " + resourcePropertyType.String())
	}
}

func GetAllResourcePropertyTypes() []model.ResourceProperty_Type {
	var types []model.ResourceProperty_Type

	for key := range model.ResourceProperty_Type_name {
		types = append(types, model.ResourceProperty_Type(key))
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

func ResourcePropertyTypeToJsonSchemaType(resourcePropertyType model.ResourceProperty_Type) *openapi3.Schema {
	switch resourcePropertyType {
	case model.ResourceProperty_STRING:
		return &openapi3.Schema{
			Type: "string",
		}
	case model.ResourceProperty_INT64:
		return &openapi3.Schema{
			Type:   "number",
			Format: "int64",
		}
	case model.ResourceProperty_INT32:
		return &openapi3.Schema{
			Type:   "number",
			Format: "int32",
		}
	case model.ResourceProperty_FLOAT64:
		return &openapi3.Schema{
			Type:   "number",
			Format: "float",
		}
	case model.ResourceProperty_FLOAT32:
		return &openapi3.Schema{
			Type:   "number",
			Format: "double",
		}
	case model.ResourceProperty_TIMESTAMP:
		return &openapi3.Schema{
			Type:   "string",
			Format: "datetime",
		}
	case model.ResourceProperty_TIME:
		return &openapi3.Schema{
			Type:   "string",
			Format: "time",
		}
	case model.ResourceProperty_DATE:
		return &openapi3.Schema{
			Type:   "string",
			Format: "date",
		}
	case model.ResourceProperty_UUID:
		return &openapi3.Schema{
			Type:   "string",
			Format: "uuid",
		}
	case model.ResourceProperty_ENUM:
		return &openapi3.Schema{
			Type: "string",
		}
	case model.ResourceProperty_BOOL:
		return &openapi3.Schema{
			Type: "boolean",
		}
	case model.ResourceProperty_REFERENCE:
		return &openapi3.Schema{
			Type: "object",
		}
	case model.ResourceProperty_OBJECT:
		return &openapi3.Schema{
			Type: "object",
		}
	case model.ResourceProperty_MAP:
		return &openapi3.Schema{
			Type: "object",
		}
	case model.ResourceProperty_LIST:
		return &openapi3.Schema{
			Type: "array",
		}
	case model.ResourceProperty_BYTES:
		return &openapi3.Schema{
			Type:   "string",
			Format: "base64",
		}
	default:
		panic("unknown property type: " + resourcePropertyType.String())
	}
}
