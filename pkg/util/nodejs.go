package util

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/tislib/data-handler/pkg/model"
)

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
