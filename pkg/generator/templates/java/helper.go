package java

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/util"
	"github.com/gosimple/slug"
	"strings"
)

func propertyName(property *model.ResourceProperty) string {
	if isJavaReservedKeyword(property.Name) {
		return "$" + property.Name
	}

	return property.Name
}

func getRestPath(resource *model.Resource) string {
	if resource.Namespace == "" || resource.Namespace == "default" {
		return slug.Make(resource.Name)
	} else {
		return slug.Make(resource.Namespace + "/" + resource.Name)
	}
}

func getJavaPropertyAnnotations(resource *model.Resource, property *model.ResourceProperty) string {
	switch property.Type {
	case model.ResourceProperty_TIME:
		return "@JsonFormat(shape = JsonFormat.Shape.STRING, timezone = \"UTC\")"
	case model.ResourceProperty_DATE:
		return "@JsonFormat(shape = JsonFormat.Shape.STRING, timezone = \"UTC\")"
	case model.ResourceProperty_TIMESTAMP:
		return "@JsonFormat(shape = JsonFormat.Shape.STRING, timezone = \"UTC\")"
	}

	return ""
}

func getJavaType(resource *model.Resource, property *model.ResourceProperty, nonPrimitive bool) string {
	var required = property.Required && !nonPrimitive

	switch property.Type {
	case model.ResourceProperty_STRING:
		return "String"
	case model.ResourceProperty_INT32:
		if required {
			return "int"
		} else {
			return "Integer"
		}
	case model.ResourceProperty_INT64:
		if required {
			return "long"
		} else {
			return "Long"
		}
	case model.ResourceProperty_FLOAT32:
		if required {
			return "float"
		} else {
			return "Float"
		}
	case model.ResourceProperty_FLOAT64:
		if required {
			return "double"
		} else {
			return "Double"
		}
	case model.ResourceProperty_BOOL:
		if required {
			return "boolean"
		} else {
			return "Boolean"
		}
	case model.ResourceProperty_REFERENCE:
		return javaClassName(property.Reference.Resource)
	case model.ResourceProperty_ENUM:
		return javaClassName(resource.Name) + "." + javaClassName(property.Name)
	case model.ResourceProperty_OBJECT:
		return "Object"
	case model.ResourceProperty_LIST:
		return "java.util.List<" + getJavaType(resource, property.Item, true) + ">"
	case model.ResourceProperty_MAP:
		return "java.util.Map<String, " + getJavaType(resource, property.Item, true) + ">"
	case model.ResourceProperty_TIME:
		return "java.time.Instant"
	case model.ResourceProperty_DATE:
		return "java.time.LocalDate"
	case model.ResourceProperty_TIMESTAMP:
		return "java.time.Instant"
	case model.ResourceProperty_UUID:
		return "java.util.UUID"
	case model.ResourceProperty_STRUCT:
		return javaClassName(resource.Name) + "." + javaClassName(*property.TypeRef)
	}

	panic("Unknown type: " + property.Type.String())
}

func javaClassName(resourceName string) string {
	return util.Capitalize(resourceName)
}

func getAllEnums(resource *model.Resource) []*model.ResourceProperty {
	var enums []*model.ResourceProperty
	var addedEnum = make(map[string]bool)
	util.ResourceWalkProperties(resource, func(path string, prop *model.ResourceProperty) {
		var name = prop.Name
		if prop.Type == model.ResourceProperty_ENUM {
			var enumName = name

			if addedEnum[enumName] {
				return
			}

			enums = append(enums, &model.ResourceProperty{
				Name:       enumName,
				EnumValues: prop.EnumValues,
			})
			addedEnum[enumName] = true
		}
	})
	return enums
}

func enumName(enumValue string) string {
	return strings.ToUpper(util.ToSnakeCase(enumValue))
}
