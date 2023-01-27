package postgres

import (
	"data-handler/model"
	"strconv"
)

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
	case model.ResourcePropertyType_TYPE_INT64:
		return "INT8"
	case model.ResourcePropertyType_TYPE_FLOAT32:
		return "FLOAT"
	case model.ResourcePropertyType_TYPE_FLOAT64:
		return "DOUBLE PRECISION"
	case model.ResourcePropertyType_TYPE_TIME:
		return "TIME"
	case model.ResourcePropertyType_TYPE_TIMESTAMP:
		return "TIMESTAMP"
	case model.ResourcePropertyType_TYPE_BOOL:
		return "BOOL"
	case model.ResourcePropertyType_TYPE_OBJECT:
		return "JSONB"
	case model.ResourcePropertyType_TYPE_BYTES:
		return "BYTEA"

	default:
		panic("unknown property type")
	}
}

func getPropertyTypeFromPsql(columnType string) model.ResourcePropertyType {
	switch columnType {
	case "bool":
		return model.ResourcePropertyType_TYPE_BOOL
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
		return model.ResourcePropertyType_TYPE_FLOAT32
	case "float8":
		return model.ResourcePropertyType_TYPE_FLOAT64
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
	case "varchar":
		return model.ResourcePropertyType_TYPE_STRING
	case "text":
		return model.ResourcePropertyType_TYPE_STRING
	case "uuid":
		return model.ResourcePropertyType_TYPE_UUID
	default:
		return model.ResourcePropertyType_TYPE_STRING
	}
}
