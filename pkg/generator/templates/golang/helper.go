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
			imports = append(imports, "encoding/json")
		}
	})

	return util.ArrayUnique(imports)
}

func getAllSubTypes(resource *model.Resource) []*model.ResourceSubType {
	var subTypes []*model.ResourceSubType
	for _, subType := range resource.Types {
		subTypes = append(subTypes, subType)
	}
	return subTypes
}

func getAllEnums(resource *model.Resource) []*model.ResourceProperty {
	var enums []*model.ResourceProperty
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

			enums = append(enums, &model.ResourceProperty{
				Name:       enumName,
				EnumValues: prop.EnumValues,
			})
		}
	})
	return enums
}

func isPrimitive(prop *model.ResourceProperty) bool {
	switch prop.Type {
	case model.ResourceProperty_STRUCT, model.ResourceProperty_OBJECT, model.ResourceProperty_REFERENCE, model.ResourceProperty_MAP, model.ResourceProperty_LIST:
		return false
	default:
		return true
	}
}

func isNullable(prop *model.ResourceProperty) bool {
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

func StreamPropertyType(qw422016 *qt422016.Writer, resource *model.Resource, prop *model.ResourceProperty) {
	qw422016.W().Write([]byte(PropertyType(resource, prop)))
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
		if prop.TypeRef != nil {
			typeVal = strcase.ToCamel(resource.Name + "_" + *prop.TypeRef)
		} else {
			typeVal = GenerateInlineStructCode(resource, prop.Properties)
		}
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
	qw422016.W().Write([]byte(PropPureGoType(resource, prop, actualName)))
}

func StreamGoTypeRaw(qw422016 *qt422016.Writer, prop *model.ResourceProperty) {
	qw422016.W().Write([]byte(GoTypeRaw(prop)))
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
	qw422016.W().Write([]byte(GoName(name)))
}

func GoVarName(name string) string {
	return util.SnakeCaseToCamelCase(name)
}

func StreamGoVarName(qw422016 *qt422016.Writer, name string) {
	qw422016.W().Write([]byte(GoVarName(name)))
}
