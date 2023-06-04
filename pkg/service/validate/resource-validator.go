package validate

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/util"
	"strconv"
)

func ValidateResourceProperties(resource *model.Resource, path string, depth int, properties []*model.ResourceProperty, wrapped bool) []*model.ErrorField {
	var errorFields []*model.ErrorField
	for i, prop := range properties {
		propertyPrefix := prop.Name + "."

		if path != "" {
			propertyPrefix = path + propertyPrefix
		}

		if !wrapped && prop.Name == "" {
			errorFields = append(errorFields, &model.ErrorField{
				Property: propertyPrefix + "Name{index:" + strconv.Itoa(i) + "}",
				Message:  "should not be blank",
				Value:    nil,
			})
		}

		if prop.Type == model.ResourceProperty_ENUM {
			if prop.EnumValues == nil || len(prop.EnumValues) == 0 {
				errorFields = append(errorFields, &model.ErrorField{
					Property: propertyPrefix + "EnumValues",
					Message:  "EnumValues should not be empty for enum type",
					Value:    nil,
				})
			}
		}

		if prop.TypeRef != nil && prop.Type != model.ResourceProperty_STRUCT {
			errorFields = append(errorFields, &model.ErrorField{
				Property: propertyPrefix + "TypeRef",
				Message:  "TypeRef should be empty for non-struct type",
				Value:    nil,
			})
		}

		if prop.Type == model.ResourceProperty_REFERENCE {
			if prop.Reference == nil {
				errorFields = append(errorFields, &model.ErrorField{
					Property: propertyPrefix + "Reference",
					Message:  "Reference should not be empty for reference type",
					Value:    nil,
				})
			} else if prop.Reference.ReferencedResource == "" {
				errorFields = append(errorFields, &model.ErrorField{
					Property: propertyPrefix + "Reference.ReferencedResource",
					Message:  "Reference.ReferencedResource should not be empty for reference type",
					Value:    nil,
				})
			}
		}

		if prop.Type == model.ResourceProperty_LIST {
			if prop.Item == nil {
				errorFields = append(errorFields, &model.ErrorField{
					Property: propertyPrefix + "Item",
					Message:  "Item should not be empty for list type",
					Value:    nil,
				})
			} else {
				errorFields = append(errorFields, ValidateResourceProperties(resource, propertyPrefix+"Item", depth+1, []*model.ResourceProperty{prop.Item}, true)...)
			}
		}

		if prop.Type == model.ResourceProperty_MAP {
			if prop.Item == nil {
				errorFields = append(errorFields, &model.ErrorField{
					Property: propertyPrefix + "item",
					Message:  "Item should not be empty for map type",
					Value:    nil,
				})
			} else {
				errorFields = append(errorFields, ValidateResourceProperties(resource, propertyPrefix+"Item", depth+1, []*model.ResourceProperty{prop.Item}, true)...)
			}
		}

		if prop.Type == model.ResourceProperty_STRUCT {
			if prop.TypeRef == nil && prop.Properties == nil {
				errorFields = append(errorFields, &model.ErrorField{
					Property: propertyPrefix + "Properties",
					Message:  "Properties or TypeRef should not be empty for struct type",
					Value:    nil,
				})
			}

			if prop.TypeRef != nil {
				// locate type
				typeRef := util.LocateArrayElement(resource.Types, func(elem *model.ResourceSubType) bool {
					return elem.Name == *prop.TypeRef
				})

				if typeRef == nil {
					errorFields = append(errorFields, &model.ErrorField{
						Property: propertyPrefix + "TypeRef",
						Message:  "TypeRef should reference an existing type",
						Value:    nil,
					})
				}
			}

			if prop.Properties != nil {
				errorFields = append(errorFields, ValidateResourceProperties(resource, propertyPrefix+"Properties", depth+1, prop.Properties, false)...)
			}
		}

		// check for additional fields
		if prop.DefaultValue != nil && prop.DefaultValue.AsInterface() != nil {
			errorFields = append(errorFields, PropertyPackedValue(resource, prop, resource.Id, "", prop.DefaultValue)...)
		}
		if prop.ExampleValue != nil && prop.ExampleValue.AsInterface() != nil {
			errorFields = append(errorFields, PropertyPackedValue(resource, prop, resource.Id, "", prop.ExampleValue)...)
		}
	}
	return errorFields
}
