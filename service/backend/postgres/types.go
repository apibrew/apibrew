package postgres

import (
	"data-handler/stub/model"
	"github.com/google/uuid"
	"reflect"
	"strconv"
	"time"
)

func propertyPointer(propertyType model.ResourcePropertyType, required bool) interface{} {
	switch propertyType {
	case model.ResourcePropertyType_TYPE_INT32:
		if required {
			return new(int32)
		} else {
			return new(*int32)
		}
	case model.ResourcePropertyType_TYPE_TEXT:
		if required {
			return new(string)
		} else {
			return new(*string)
		}
	case model.ResourcePropertyType_TYPE_STRING:
		if required {
			return new(string)
		} else {
			return new(*string)
		}
	case model.ResourcePropertyType_TYPE_UUID:
		if required {
			return new(uuid.UUID)
		} else {
			return new(*uuid.UUID)
		}
	case model.ResourcePropertyType_TYPE_DATE:
		if required {
			return new(time.Time)
		} else {
			return new(*time.Time)
		}
	default:
		panic("unknown property type")
	}
}

func dereference(val interface{}) interface{} {
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

func stringifyProperty(value interface{}, propertyType model.ResourcePropertyType, required bool) string {
	switch propertyType {
	case model.ResourcePropertyType_TYPE_INT32:
		return strconv.Itoa(int(*value.(*int32)))
	case model.ResourcePropertyType_TYPE_STRING:
		return *value.(*string)
	case model.ResourcePropertyType_TYPE_UUID:
		return (*value.(*uuid.UUID)).String()
	case model.ResourcePropertyType_TYPE_DATE:
		return (*value.(*time.Time)).Format(time.RFC3339)
	case model.ResourcePropertyType_TYPE_TEXT:
		return *value.(*string)
	default:
		panic("unknown property type")
	}
}

func getPsqlTypeFromProperty(propertyType model.ResourcePropertyType, length uint32) string {
	switch propertyType {
	case model.ResourcePropertyType_TYPE_INT32:
		return "INT"
	case model.ResourcePropertyType_TYPE_UUID:
		return "UUID"
	case model.ResourcePropertyType_TYPE_STRING:
		return "VARCHAR(" + strconv.Itoa(int(length)) + ")"
	case model.ResourcePropertyType_TYPE_DATE:
		return "DATE"
	case model.ResourcePropertyType_TYPE_TEXT:
		return "TEXT"
	default:
		panic("unknown property type")
	}
}

func getPropertyTypeFromPsql(columnType string) model.ResourcePropertyType {
	switch columnType {
	case "bool":
		return model.ResourcePropertyType_TYPE_STRING
	case "bytea":
		return model.ResourcePropertyType_TYPE_BYTES
	case "char":
		return model.ResourcePropertyType_TYPE_STRING
	case "date":
		return model.ResourcePropertyType_TYPE_DATE
	case "time":
		return model.ResourcePropertyType_TYPE_TIME
	case "timestamp":
		return model.ResourcePropertyType_TYPE_TIMESTAMP
	case "timestampz":
		return model.ResourcePropertyType_TYPE_TIMESTAMP
	case "float4":
		return model.ResourcePropertyType_TYPE_FLOAT
	case "float8":
		return model.ResourcePropertyType_TYPE_FLOAT
	case "int2":
		return model.ResourcePropertyType_TYPE_INT32
	case "int4":
		return model.ResourcePropertyType_TYPE_INT32
	case "int8":
		return model.ResourcePropertyType_TYPE_INT64
	case "jsonb":
		return model.ResourcePropertyType_TYPE_OBJECT
	case "json":
		return model.ResourcePropertyType_TYPE_OBJECT
	case "numeric":
		return model.ResourcePropertyType_TYPE_NUMERIC
	case "varchar":
		return model.ResourcePropertyType_TYPE_STRING
	case "text":
		return model.ResourcePropertyType_TYPE_TEXT
	case "uuid":
		return model.ResourcePropertyType_TYPE_UUID
	default:
		return model.ResourcePropertyType_TYPE_TEXT
	}
}
