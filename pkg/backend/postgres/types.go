package postgres

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/types"
	"strconv"
)

func (p postgreSqlBackendOptions) TypeModifier(propertyType model.ResourceProperty_Type) types.PropertyType {
	return types.ByResourcePropertyType(propertyType)
}

func (p postgreSqlBackendOptions) GetSqlTypeFromProperty(propertyType model.ResourceProperty_Type, length uint32) string {
	if propertyType == model.ResourceProperty_STRING && length == 0 {
		length = 255
	}

	switch propertyType {
	case model.ResourceProperty_INT32:
		return "INT"
	case model.ResourceProperty_UUID:
		return "UUID"
	case model.ResourceProperty_STRING:
		return "VARCHAR(" + strconv.Itoa(int(length)) + ")"
	case model.ResourceProperty_DATE:
		return "DATE"
	case model.ResourceProperty_INT64:
		return "INT8"
	case model.ResourceProperty_FLOAT32:
		return "FLOAT"
	case model.ResourceProperty_FLOAT64:
		return "DOUBLE PRECISION"
	case model.ResourceProperty_TIME:
		return "TIME"
	case model.ResourceProperty_TIMESTAMP:
		return "TIMESTAMP"
	case model.ResourceProperty_BOOL:
		return "BOOL"
	case model.ResourceProperty_OBJECT:
		return "JSONB"
	case model.ResourceProperty_BYTES:
		return "BYTEA"
	case model.ResourceProperty_ENUM:
		return "VARCHAR(64)"
	case model.ResourceProperty_REFERENCE:
		return "UUID"
	case model.ResourceProperty_MAP:
		return "JSONB"
	case model.ResourceProperty_LIST:
		return "JSONB"
	case model.ResourceProperty_STRUCT:
		return "JSONB"
	default:
		panic("unknown property type: " + propertyType.String() + " for postgres type conversion")
	}
}

func (p postgreSqlBackendOptions) GetPropertyTypeFromPsql(columnType string) model.ResourceProperty_Type {
	switch columnType {
	case "bool":
		return model.ResourceProperty_BOOL
	case "bytea":
		return model.ResourceProperty_BYTES
	case "char":
		return model.ResourceProperty_STRING
	case "date":
		return model.ResourceProperty_DATE
	case "time":
		return model.ResourceProperty_TIME
	case "timestamp":
		return model.ResourceProperty_TIMESTAMP
	case "timestampz":
		return model.ResourceProperty_TIMESTAMP
	case "float4":
		return model.ResourceProperty_FLOAT32
	case "float8":
		return model.ResourceProperty_FLOAT64
	case "int2":
		return model.ResourceProperty_INT32
	case "int4":
		return model.ResourceProperty_INT32
	case "int8":
		return model.ResourceProperty_INT64
	case "jsonb":
		return model.ResourceProperty_OBJECT
	case "json":
		return model.ResourceProperty_OBJECT
	case "varchar":
		return model.ResourceProperty_STRING
	case "text":
		return model.ResourceProperty_STRING
	case "uuid":
		return model.ResourceProperty_UUID
	default:
		return model.ResourceProperty_STRING
	}
}
