package mapping

import (
	"github.com/apibrew/apibrew/pkg/abs"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/util"
	"strings"
)

func ResourcePropertyToRecord(property *model.ResourceProperty, resource *model.Resource) map[string]interface{} {
	properties := make(map[string]interface{})

	if property.Title != nil {
		properties["title"] = *property.Title
	}
	if property.Description != nil {
		properties["description"] = *property.Description
	}
	properties["type"] = property.Type.String()

	if property.Type == model.ResourceProperty_LIST || property.Type == model.ResourceProperty_MAP {
		properties["item"] = ResourcePropertyToRecord(property.Item, resource)
	}

	if property.Type == model.ResourceProperty_STRUCT {
		properties["typeRef"] = *property.TypeRef
	}

	if property.Type == model.ResourceProperty_ENUM {
		properties["enumValues"] = util.ArrayMap(property.EnumValues, func(v string) interface{} {
			return v
		})
	}

	properties["primary"] = property.Primary
	properties["required"] = property.Required
	properties["length"] = int32(property.Length)
	properties["unique"] = property.Unique
	properties["immutable"] = property.Immutable
	properties["virtual"] = property.Virtual

	if property.Reference != nil {
		referenceNamespace := property.Reference.Namespace
		if referenceNamespace == "" {
			referenceNamespace = resource.Namespace
		}

		if property.BackReference != nil {
			properties["backReference"] = property.BackReference.Property
		}

		properties["reference"] = referenceNamespace + "/" + property.Reference.Resource
	}

	if property.DefaultValue != nil {
		properties["defaultValue"] = property.DefaultValue.AsInterface()
	}
	if property.ExampleValue != nil {
		properties["exampleValue"] = property.ExampleValue.AsInterface()
	}

	properties["annotations"] = convertMap(property.Annotations, func(v string) interface{} {
		return v
	})

	return properties
}

func ResourcePropertyFromRecord(propertyName string, record abs.RecordLike) *model.ResourceProperty {
	if record == nil {
		return nil
	}

	var reference *model.Reference
	var backReference *model.BackReference

	if record.HasProperty("reference") {
		reference = &model.Reference{}
		if record.GetStructProperty("reference").GetStructValue() != nil {
			var referenceProperties = record.GetStructProperty("reference").GetStructValue().GetFields()
			reference.Resource = referenceProperties["resource"].GetStructValue().GetFields()["name"].GetStringValue()
			if referenceProperties["resource"].GetStructValue().GetFields()["namespace"] != nil && referenceProperties["resource"].GetStructValue().GetFields()["namespace"].GetStructValue() != nil {
				reference.Namespace = referenceProperties["resource"].GetStructValue().GetFields()["namespace"].GetStructValue().GetFields()["name"].GetStringValue()
			}

			if record.HasProperty("backReference") && record.GetStructProperty("backReference").GetStringValue() != "" {
				backReference = &model.BackReference{
					Property: record.GetStructProperty("backReference").GetStringValue(),
				}
			}
		} else {
			var referenceParts = strings.Split(record.GetStructProperty("reference").GetStringValue(), "/")

			if len(referenceParts) == 1 {
				reference.Resource = referenceParts[0]
			} else if len(referenceParts) == 2 {
				reference.Resource = referenceParts[1]
				reference.Namespace = referenceParts[0]
			} else {
				panic("Invalid reference format")
			}

			// reference.Cascade //todo implement it

			if record.HasProperty("backReference") && record.GetStructProperty("backReference").GetStringValue() != "" {
				backReference = &model.BackReference{
					Property: record.GetStructProperty("backReference").GetStringValue(),
				}
			}
		}
	}

	var resourceProperty = &model.ResourceProperty{
		Name:          propertyName,
		Type:          model.ResourceProperty_Type(model.ResourceProperty_Type_value[strings.ToUpper(record.GetStructProperty("type").GetStringValue())]),
		Required:      record.GetStructProperty("required").GetBoolValue(),
		Length:        uint32(record.GetStructProperty("length").GetNumberValue()),
		Unique:        record.GetStructProperty("unique").GetBoolValue(),
		Primary:       record.GetStructProperty("primary").GetBoolValue(),
		Immutable:     record.GetStructProperty("immutable").GetBoolValue(),
		DefaultValue:  record.GetStructProperty("defaultValue"),
		ExampleValue:  record.GetStructProperty("exampleValue"),
		Reference:     reference,
		BackReference: backReference,
		Annotations: convertMap(record.GetStructProperty("annotations").GetStructValue().AsMap(), func(v interface{}) string {
			return v.(string)
		}),
	}

	if record.HasProperty("virtual") {
		resourceProperty.Virtual = record.GetStructProperty("virtual").GetBoolValue()
	}

	if record.HasProperty("title") {
		resourceProperty.Title = new(string)
		*resourceProperty.Title = record.GetStructProperty("title").GetStringValue()
	}

	if record.HasProperty("description") {
		resourceProperty.Description = new(string)
		*resourceProperty.Description = record.GetStructProperty("description").GetStringValue()
	}

	if resourceProperty.Type == model.ResourceProperty_LIST || resourceProperty.Type == model.ResourceProperty_MAP {
		resourceProperty.Item = ResourcePropertyFromRecord("", abs.NewRecordLikeWithStructProperties(record.GetStructProperty("item").GetStructValue().GetFields()))
	}

	if resourceProperty.Type == model.ResourceProperty_STRUCT {
		resourceProperty.TypeRef = new(string)
		*resourceProperty.TypeRef = record.GetStructProperty("typeRef").GetStringValue()
	}

	if resourceProperty.Type == model.ResourceProperty_ENUM {
		resourceProperty.EnumValues = util.ArrayMap(record.GetProperty("enumValues").([]interface{}), func(v interface{}) string {
			return v.(string)
		})
	}

	return resourceProperty
}
