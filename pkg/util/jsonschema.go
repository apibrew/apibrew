package util

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/getkin/kin-openapi/openapi3"
)

type PropertiesWithTitleAndDescription interface {
	GetTitle() string
	GetDescription() string
	GetProperties() []*model.ResourceProperty
}

func ResourceToJsonSchema(resource *model.Resource) *openapi3.Schema {
	return PropertiesWithTitleToJsonSchema(resource, resource)
}

func PropertiesWithTitleToJsonSchema(resource *model.Resource, elem PropertiesWithTitleAndDescription) *openapi3.Schema {
	var requiredItems []string

	recordSchema := &openapi3.Schema{
		Type:       "object",
		Properties: map[string]*openapi3.SchemaRef{},
	}

	for _, property := range elem.GetProperties() {
		recordSchema.Properties[property.Name] = ResourcePropertyTypeToJsonSchemaType(resource, property)

		if property.Required {
			requiredItems = append(requiredItems, property.Name)
		}
	}

	recordSchema.Required = requiredItems
	recordSchema.Title = elem.GetTitle()
	recordSchema.Description = elem.GetDescription()

	return recordSchema
}

func ResourcePropertyTypeToJsonSchemaType(resource *model.Resource, property *model.ResourceProperty) *openapi3.SchemaRef {
	var propSchemaRef = &openapi3.SchemaRef{}

	switch property.Type {
	case model.ResourceProperty_STRING:
		propSchemaRef.Value = &openapi3.Schema{
			Type: "string",
		}
	case model.ResourceProperty_INT64:
		propSchemaRef.Value = &openapi3.Schema{
			Type:   "number",
			Format: "int64",
		}
	case model.ResourceProperty_INT32:
		propSchemaRef.Value = &openapi3.Schema{
			Type:   "number",
			Format: "int32",
		}
	case model.ResourceProperty_FLOAT64:
		propSchemaRef.Value = &openapi3.Schema{
			Type:   "number",
			Format: "float",
		}
	case model.ResourceProperty_FLOAT32:
		propSchemaRef.Value = &openapi3.Schema{
			Type:   "number",
			Format: "double",
		}
	case model.ResourceProperty_TIMESTAMP:
		propSchemaRef.Value = &openapi3.Schema{
			Type:   "string",
			Format: "datetime",
		}
	case model.ResourceProperty_TIME:
		propSchemaRef.Value = &openapi3.Schema{
			Type:   "string",
			Format: "time",
		}
	case model.ResourceProperty_DATE:
		propSchemaRef.Value = &openapi3.Schema{
			Type:   "string",
			Format: "date",
		}
	case model.ResourceProperty_UUID:
		propSchemaRef.Value = &openapi3.Schema{
			Type:   "string",
			Format: "uuid",
		}
	case model.ResourceProperty_ENUM:
		propSchemaRef.Value = &openapi3.Schema{
			Type: "string",
		}
	case model.ResourceProperty_BOOL:
		propSchemaRef.Value = &openapi3.Schema{
			Type: "boolean",
		}
	case model.ResourceProperty_REFERENCE:
		var namespace = resource.Namespace
		if property.Reference.Namespace != "" {
			namespace = property.Reference.Namespace
		}

		var schemaName = Capitalize(namespace) + property.Reference.Resource

		if namespace == "default" || namespace == "" {
			schemaName = property.Reference.Resource
		}

		propSchemaRef.Ref = "#/components/schemas/" + schemaName
	case model.ResourceProperty_OBJECT:
		propSchemaRef.Value = &openapi3.Schema{
			Type: "object",
		}
	case model.ResourceProperty_MAP:
		propSchemaRef.Value = &openapi3.Schema{
			Type:                 "object",
			AdditionalProperties: ResourcePropertyTypeToJsonSchemaType(resource, property.Item),
		}
	case model.ResourceProperty_LIST:
		propSchemaRef.Value = &openapi3.Schema{
			Type:  "array",
			Items: ResourcePropertyTypeToJsonSchemaType(resource, property.Item),
		}
	case model.ResourceProperty_BYTES:
		propSchemaRef.Value = &openapi3.Schema{
			Type:   "string",
			Format: "base64",
		}
	case model.ResourceProperty_STRUCT:
		if property.TypeRef != nil {
			propSchemaRef.Ref = "#/components/schemas/" + ResourceJsonSchemaName(resource) + *property.TypeRef
		} else {
			propSchemaRef.Value = PropertiesWithTitleToJsonSchema(resource, property)
		}
	default:
		panic("unknown property type: " + property.String())
	}

	if propSchemaRef.Value != nil {
		if property.ExampleValue != nil {
			propSchemaRef.Value.Example = property.ExampleValue.AsInterface()
		}

		if property.DefaultValue != nil {
			propSchemaRef.Value.Default = property.DefaultValue.AsInterface()
		}

		if property.Unique {
			propSchemaRef.Value.UniqueItems = true
		}

		if property.Immutable {
			propSchemaRef.Value.Extensions = map[string]interface{}{
				"x-immutable": true,
			}
		}

		propSchemaRef.Value.Extensions = make(map[string]interface{})

		for key, value := range property.Annotations {
			propSchemaRef.Value.Extensions["x-"+key] = value
		}
	}

	return propSchemaRef
}

func ResourceJsonSchemaName(item *model.Resource) string {
	var schemaName = Capitalize(item.Namespace) + item.Name

	if item.Namespace == "default" {
		schemaName = item.Name
	}
	return schemaName
}
