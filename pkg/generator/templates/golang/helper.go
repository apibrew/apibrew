package golang

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/service/annotations"
	"github.com/apibrew/apibrew/pkg/types"
	"github.com/apibrew/apibrew/pkg/util"
	"github.com/iancoleman/strcase"
	qt422016 "github.com/valyala/quicktemplate"
	"reflect"
	"strings"
)

func getImports(resource *model.Resource) []string {
	imports := []string{}
	util.ResourceWalkProperties(resource, func(path string, prop *model.ResourceProperty) {
		if prop.Type == model.ResourceProperty_UUID {
			imports = append(imports, "github.com/google/uuid")
		} else if prop.Type == model.ResourceProperty_TIMESTAMP || prop.Type == model.ResourceProperty_TIME || prop.Type == model.ResourceProperty_DATE {
			imports = append(imports, "time")
		} else if prop.Type == model.ResourceProperty_OBJECT {
			imports = append(imports, "github.com/apibrew/apibrew/pkg/formats/unstructured")
		}
	})

	return util.ArrayUnique(imports)
}

func getResourceSpecificImports(resource *model.Resource) []string {
	imports := []string{}

	if annotations.Get(resource, annotations.SelfContainedProperty) != "" {
		imports = append(imports, "encoding/json")
	}

	return util.ArrayUnique(imports)
}

func getAllSubTypes(resource *model.Resource) []*model.ResourceSubType {
	return resource.Types
}

func getAllEnums(resource *model.Resource) []*model.ResourceProperty {
	var enums []*model.ResourceProperty
	var addedEnum = make(map[string]bool)
	util.ResourceWalkProperties(resource, func(path string, prop *model.ResourceProperty) {
		var name = prop.Name
		if name == "" {
			pathParts := strings.Split(path, ".")
			name = pathParts[len(pathParts)-1]
			name = strings.ReplaceAll(name, "[]", "")
		}
		if prop.Type == model.ResourceProperty_ENUM {
			var enumName = GoName(resource.Name + "_" + name)

			if annotations.Get(prop, annotations.TypeName) != "" {
				enumName = annotations.Get(prop, annotations.TypeName)
			}

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

func isNullable(prop *model.ResourceProperty) bool {
	if annotations.IsEnabled(prop, annotations.AllowEmptyPrimitive) {
		return false
	}

	return prop.Type == model.ResourceProperty_REFERENCE || !prop.Required || annotations.IsEnabled(prop, annotations.SpecialProperty)
}

func isSelfNullable(prop *model.ResourceProperty) bool {
	return prop.Type == model.ResourceProperty_REFERENCE || prop.Type == model.ResourceProperty_LIST || prop.Type == model.ResourceProperty_MAP
}

func PropertyType(resource *model.Resource, prop *model.ResourceProperty) string {
	if !isNullable(prop) || isSelfNullable(prop) {
		return PropPureGoType(resource, prop, "")
	} else {
		return "*" + PropPureGoType(resource, prop, "")
	}
}

func HasPointer(prop *model.ResourceProperty, collectionItem bool) bool {
	if prop.Type == model.ResourceProperty_REFERENCE {
		return true
	}

	if collectionItem {
		return false
	}

	return isNullable(prop)
}

func NormalizePointer(prop *model.ResourceProperty, varName string, collectionItem bool, mustHavePointer bool) string {
	hasPointer := HasPointer(prop, collectionItem)

	if mustHavePointer && !hasPointer {
		return "&" + varName
	}

	if !mustHavePointer && hasPointer {
		return "*" + varName
	}

	return varName
}

func StreamNormalizePointer(qw422016 *qt422016.Writer, prop *model.ResourceProperty, varName string, collectionItem bool, mustHavePointer bool) {
	_, _ = qw422016.W().Write([]byte(NormalizePointer(prop, varName, collectionItem, mustHavePointer)))
}

func StreamPropertyType(qw422016 *qt422016.Writer, resource *model.Resource, prop *model.ResourceProperty) {
	_, _ = qw422016.W().Write([]byte(PropertyType(resource, prop)))
}

func PropPureGoType(resource *model.Resource, prop *model.ResourceProperty, actualName string) string {
	typeVal := GoTypeRaw(prop)

	if actualName == "" {
		actualName = prop.Name
	}

	if prop.Type == model.ResourceProperty_REFERENCE {
		typeVal = "*" + strcase.ToCamel(prop.Reference.Resource)
	} else if prop.Type == model.ResourceProperty_MAP {
		typeVal = "map[string]" + PropPureGoType(resource, prop.Item, prop.Name)
	} else if prop.Type == model.ResourceProperty_LIST {
		typeVal = "[]" + PropPureGoType(resource, prop.Item, prop.Name)
	} else if prop.Type == model.ResourceProperty_STRUCT {
		typeVal = strcase.ToCamel(resource.Name + "_" + *prop.TypeRef)
	} else if prop.Type == model.ResourceProperty_OBJECT {
		typeVal = "unstructured.Unstructured"
	} else if prop.Type == model.ResourceProperty_ENUM {
		typeVal = strcase.ToCamel(resource.Name + "_" + actualName)

		if annotations.Get(prop, annotations.TypeName) != "" {
			typeVal = annotations.Get(prop, annotations.TypeName)
		}
	}

	return typeVal
}

func StreamPropPureGoType(qw422016 *qt422016.Writer, resource *model.Resource, prop *model.ResourceProperty, actualName string) {
	_, _ = qw422016.W().Write([]byte(PropPureGoType(resource, prop, actualName)))
}

func StreamGoTypeRaw(qw422016 *qt422016.Writer, prop *model.ResourceProperty) {
	_, _ = qw422016.W().Write([]byte(GoTypeRaw(prop)))
}

func GoTypeRaw(prop *model.ResourceProperty) string {
	typ := types.ByResourcePropertyType(prop.Type)

	typeDef := reflect.TypeOf(typ.Default()).String()
	return typeDef
}

func GoName(name string) string {
	return util.Capitalize(util.SnakeCaseToCamelCase(name))
}

func StreamGoName(qw422016 *qt422016.Writer, name string) {
	_, _ = qw422016.W().Write([]byte(GoName(name)))
}

func GoVarName(name string) string {
	return util.SnakeCaseToCamelCase(name)
}

func StreamGoVarName(qw422016 *qt422016.Writer, name string) {
	_, _ = qw422016.W().Write([]byte(GoVarName(name)))
}

/*
_______________________
| required | collectionItem | mustHavePointer | result
| ---------|----------------|-----------------|--------
| true     | true           | true            | &
| true     | true           | false           |
| true     | false          | true            | &
| true     | false          | false           |
| false    | true           | true            | &
| false    | true           | false           |
| false    | false          | true            | &
| false    | false          | false           |
*/
