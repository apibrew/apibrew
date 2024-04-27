package validate

import (
	"fmt"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/util"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/structpb"
	"strconv"
	"strings"
)

func ValidateResource(resource *model.Resource) error {
	var errorFields []*model.ErrorField

	if resource.Name == "" {
		errorFields = append(errorFields, &model.ErrorField{
			RecordId: resource.Id,
			Property: "Name",
			Message:  "should not be empty",
			Value:    nil,
		})
	} else if !NamePattern.MatchString(resource.Name) {
		errorFields = append(errorFields, &model.ErrorField{
			RecordId: resource.Id,
			Property: "Name",
			Message:  "should match pattern " + NamePattern.String(),
			Value:    structpb.NewStringValue(resource.Name),
		})
	}

	if !resource.Virtual {
		if resource.SourceConfig == nil {
			errorFields = append(errorFields, &model.ErrorField{
				RecordId: resource.Id,
				Property: "SourceConfig",
				Message:  "should not be nil",
				Value:    nil,
			})
		} else {
			if resource.SourceConfig.DataSource == "" {
				errorFields = append(errorFields, &model.ErrorField{
					RecordId: resource.Id,
					Property: "SourceConfig.DataSource",
					Message:  "should not be blank",
					Value:    nil,
				})
			}

			if resource.SourceConfig.Entity == "" {
				errorFields = append(errorFields, &model.ErrorField{
					RecordId: resource.Id,
					Property: "SourceConfig.Entity",
					Message:  "should not be blank",
					Value:    nil,
				})
			}
		}
	}

	errorFields = append(errorFields, ValidateResourceProperties(resource, "", 0, resource.Properties, false)...)

	for _, subType := range resource.Types {
		if !NamePattern.MatchString(subType.Name) {
			errorFields = append(errorFields, &model.ErrorField{
				RecordId: resource.Id,
				Property: "subType.Name",
				Message:  "should match pattern " + NamePattern.String(),
				Value:    structpb.NewStringValue(subType.Name),
			})
		}
		errorFields = append(errorFields, ValidateResourceProperties(resource, subType.Name+".", 1, subType.Properties, false)...)
	}

	if len(errorFields) > 0 {
		var details []string

		for _, errorField := range errorFields {
			details = append(details, fmt.Sprintf("%s: %s", errorField.Property, errorField.Message))
		}

		return errors.ResourceValidationError.WithDetails(strings.Join(details, ";")).WithErrorFields(errorFields)
	}

	return nil
}

func ValidateResourceProperties(resource *model.Resource, path string, depth int, properties []*model.ResourceProperty, wrapped bool) []*model.ErrorField {
	var errorFields []*model.ErrorField
	for i, prop := range properties {
		if prop.Name == "type" {
			errorFields = append(errorFields, &model.ErrorField{
				Property: path + "Name{index:" + strconv.Itoa(i) + "}",
				Message:  "property name 'type' is reserved",
				Value:    nil,
			})
		}

		propertyPrefix := prop.Name + "."

		if path != "" {
			propertyPrefix = path + propertyPrefix
		}

		if prop.Name == "" && !wrapped {
			errorFields = append(errorFields, &model.ErrorField{
				Property: propertyPrefix + "Name{index:" + strconv.Itoa(i) + "}",
				Message:  "should not be blank",
				Value:    nil,
			})
		}

		if prop.Name != "" && !NamePattern.MatchString(prop.Name) {
			errorFields = append(errorFields, &model.ErrorField{
				Property: propertyPrefix + "Name{index:" + strconv.Itoa(i) + "}",
				Message:  "should match pattern " + NamePattern.String(),
				Value:    structpb.NewStringValue(prop.Name),
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

			for _, val := range prop.EnumValues {
				if strings.ToUpper(val) != val {
					errorFields = append(errorFields, &model.ErrorField{
						Property: propertyPrefix + "EnumValues",
						Message:  "Enum values should be uppercase: " + val,
						Value:    structpb.NewStringValue(val),
					})
				}
			}
		}

		if prop.Type == 0 && prop.TypeRef != nil {
			prop.Type = model.ResourceProperty_STRUCT
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
			} else if prop.Reference.Resource == "" {
				errorFields = append(errorFields, &model.ErrorField{
					Property: propertyPrefix + "Reference.Resource",
					Message:  "Reference.Resource should not be empty for reference type",
					Value:    nil,
				})
			} // fixme: else validate if referenced resource exists
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
			if prop.TypeRef == nil {
				errorFields = append(errorFields, &model.ErrorField{
					Property: propertyPrefix + "Properties",
					Message:  "TypeRef should not be empty for struct type",
					Value:    nil,
				})
			} else {
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
		}

		// check for additional fields
		if prop.DefaultValue != nil && prop.DefaultValue.AsInterface() != nil {
			errorFields = append(errorFields, Value(resource, prop, resource.Id, "", prop.DefaultValue)...)
		}
		if prop.ExampleValue != nil && prop.ExampleValue.AsInterface() != nil {
			errorFields = append(errorFields, Value(resource, prop, resource.Id, "", prop.ExampleValue)...)
		}
	}

	if len(errorFields) > 0 {
		log.Warnf("Resource %s has errors: %v", resource.Name, errorFields)
	}

	return errorFields
}
