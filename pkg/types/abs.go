package types

import (
	"github.com/tislib/apibrew/pkg/model"
	"google.golang.org/protobuf/types/known/structpb"
	"reflect"
)

type PropertyType interface {
	Pointer(required bool) any
	String(val any) string
	IsEmpty(value any) bool
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
		return Float64Type
	case model.ResourceProperty_STRING:
		return StringType
	case model.ResourceProperty_UUID:
		return UuidType
	case model.ResourceProperty_DATE:
		return DateType
	case model.ResourceProperty_TIME:
		return TimeType
	case model.ResourceProperty_TIMESTAMP:
		return TimestampType
	case model.ResourceProperty_BOOL:
		return BoolType
	case model.ResourceProperty_OBJECT:
		return ObjectType
	case model.ResourceProperty_REFERENCE:
		return ReferenceType
	case model.ResourceProperty_ENUM:
		return StringType
	case model.ResourceProperty_MAP:
		return MapType
	case model.ResourceProperty_LIST:
		return ListType
	case model.ResourceProperty_BYTES:
		return BytesType
	case model.ResourceProperty_STRUCT:
		return StructType
	default:
		panic("unknown property type: " + resourcePropertyType.String())
	}
}

func IsPrimitive(resourcePropertyType model.ResourceProperty_Type) bool {
	switch resourcePropertyType {
	case model.ResourceProperty_OBJECT, model.ResourceProperty_REFERENCE, model.ResourceProperty_MAP, model.ResourceProperty_LIST:
		return false
	default:
		return true
	}
}

func GetAllResourcePropertyTypes() []model.ResourceProperty_Type {
	var types []model.ResourceProperty_Type

	for key := range model.ResourceProperty_Type_name {
		types = append(types, model.ResourceProperty_Type(key))
	}

	return types
}
