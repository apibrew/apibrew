package types

import (
	"data-handler/stub/model"
	"reflect"
)

type PropertyType interface {
	Pointer(required bool) any
	String(val any) string
	IsEmpty(value any) bool
	ValidateValue(value any) error
	Default() any
}

func Dereference(val interface{}) interface{} {
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
		return int32Type{}
	case model.ResourcePropertyType_TYPE_INT64:
		return int64Type{}
	case model.ResourcePropertyType_TYPE_FLOAT:
		return floatType{}
	case model.ResourcePropertyType_TYPE_DOUBLE:
		return doubleType{}
	case model.ResourcePropertyType_TYPE_NUMERIC:
		return numericType{}
	case model.ResourcePropertyType_TYPE_TEXT:
		return stringType{}
	case model.ResourcePropertyType_TYPE_STRING:
		return stringType{}
	case model.ResourcePropertyType_TYPE_UUID:
		return uuidType{}
	case model.ResourcePropertyType_TYPE_DATE:
		return dateType{}
	case model.ResourcePropertyType_TYPE_TIME:
		return timeType{}
	case model.ResourcePropertyType_TYPE_TIMESTAMP:
		return timestampType{}
	case model.ResourcePropertyType_TYPE_BOOL:
		return boolType{}
	case model.ResourcePropertyType_TYPE_OBJECT:
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
