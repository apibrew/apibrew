package mysql

import (
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/types"
	"google.golang.org/protobuf/types/known/structpb"
	"strconv"
)

func (p mysqlBackendOptions) TypeModifier(propertyType model.ResourceProperty_Type) types.PropertyType {
	if propertyType == model.ResourceProperty_TIME {
		return types.CustomTypeFromType(types.ByResourcePropertyType(model.ResourceProperty_TIME), types.CustomType{
			CustomPack: func(value interface{}) (*structpb.Value, error) {
				return structpb.NewValue(string(value.([]uint8)))
			},
			CustomUnPack: func(value *structpb.Value) (interface{}, error) {
				return []uint8(value.GetStringValue()), nil
			},
			CustomPointer: func(required bool) any {
				if required {
					return new([]uint8)
				} else {
					return new(*[]uint8)
				}
			},
		})
	} else {
		return types.ByResourcePropertyType(propertyType)
	}
}

func (p mysqlBackendOptions) GetSqlTypeFromProperty(propertyType model.ResourceProperty_Type, length uint32) string {
	switch propertyType {
	case model.ResourceProperty_INT32:
		return "INT"
	case model.ResourceProperty_UUID:
		return "varchar(36)"
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
		return "TEXT"
	case model.ResourceProperty_BYTES:
		return "BLOB"
	case model.ResourceProperty_ENUM:
		return "VARCHAR(64)"
	case model.ResourceProperty_REFERENCE:
		return "varchar(36)"
	case model.ResourceProperty_MAP:
		return "TEXT"
	case model.ResourceProperty_LIST:
		return "TEXT"

	default:
		panic("unknown property type")
	}
}

func (p mysqlBackendOptions) GetPropertyTypeFromPsql(columnType string) model.ResourceProperty_Type {
	switch columnType {
	case "bool":
		return model.ResourceProperty_BOOL
	case "blob":
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
	case "float":
		return model.ResourceProperty_FLOAT32
	case "double":
		return model.ResourceProperty_FLOAT64
	case "tinyint":
		return model.ResourceProperty_INT32
	case "smallint":
		return model.ResourceProperty_INT32
	case "bigint":
		return model.ResourceProperty_INT32
	case "int":
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
