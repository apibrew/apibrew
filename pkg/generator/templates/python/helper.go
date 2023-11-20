package python

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/util"
	"strings"
)

func propertyName(property *model.ResourceProperty) string {
	if isPythonReservedKeyword(propertyName) {
		return propertyName + "_"
	}

	return propertyName
}

func getRestPath(resource *model.Resource) string {
	return util.ResourceRestPath(resource)
}

func getPythonType(resource *model.Resource, property *model.ResourceProperty, nonPrimitive bool) string {
	switch property.Type {
	case model.ResourceProperty_STRING:
		return "str"
	case model.ResourceProperty_FLOAT64, model.ResourceProperty_FLOAT32, model.ResourceProperty_INT64, model.ResourceProperty_INT32:
		return "int"
	case model.ResourceProperty_BOOL:
		return "bool"
	case model.ResourceProperty_REFERENCE:
		return pythonClassName(property.Reference.Resource)
	case model.ResourceProperty_ENUM:
		return pythonClassName(propertyName)
	case model.ResourceProperty_OBJECT:
		return "dict"
	case model.ResourceProperty_LIST:
		return "list[" + getPythonType(resource, property.Item, true) + "]"
	case model.ResourceProperty_MAP:
		return "dict[str, " + getPythonType(resource, property.Item, true) + "]"
	case model.ResourceProperty_TIME:
		return "time"
	case model.ResourceProperty_DATE:
		return "datetime"
	case model.ResourceProperty_TIMESTAMP:
		return "datetime"
	case model.ResourceProperty_UUID:
		return "str"
	case model.ResourceProperty_BYTES:
		return "bytes"
	case model.ResourceProperty_STRUCT:
		return pythonClassName(*property.TypeRef)
	}

	panic("Unknown type: " + property.Type.String())
}

func pythonClassName(resourceName string) string {
	return util.Capitalize(util.SnakeCaseToCamelCase(resourceName))
}

func pythonVarName(resourceName string) string {
	return util.DeCapitalize(util.SnakeCaseToCamelCase(resourceName))
}

func hasInput(resource *model.Resource) bool {
	return len(resource.Types) > 0
}

func outputType(resource *model.Resource) string {
	if len(resource.Properties) > 0 {
		return getPythonType(resource, resource.Properties[0], false)
	} else {
		return "void"
	}
}

func getAllSubTypes(resource *model.Resource) []*model.ResourceSubType {
	var allTypes = resource.Types

	for _, resourceAction := range resourceActions {
		allTypes = append(allTypes, resourceAction.Types...)
	}

	for _, subType := range allTypes {
		subType.Name = pythonClassName(subType.Name)
	}

	return allTypes
}

func getAllEnums(resource *model.Resource) []*model.ResourceProperty {
	var enums []*model.ResourceProperty
	var addedEnum = make(map[string]bool)
	util.ResourceWalkProperties(resource, func(path string, prop *model.ResourceProperty) {
		var name = propName
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

func getAllReferencedResources(resource *model.Resource) []*model.Reference {
	var references []*model.Reference
	var addedReferencedResources = make(map[string]bool)

	addedReferencedResources[resource.Namespace+"."+resource.Name] = true

	util.ResourceWalkProperties(resource, func(path string, prop *model.ResourceProperty) {
		if prop.Type == model.ResourceProperty_REFERENCE {

			if addedReferencedResources[prop.Reference.Namespace+"."+prop.Reference.Resource] {
				return
			}

			references = append(references, prop.Reference)
			addedReferencedResources[prop.Reference.Namespace+"."+prop.Reference.Resource] = true
		}
	})

	return references
}

func fileName(name string) string {
	return util.ToSnakeCase(name)
}

func enumName(enumValue string) string {
	return strings.ToUpper(util.ToSnakeCase(enumValue))
}
