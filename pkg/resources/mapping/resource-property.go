package mapping

import (
	"github.com/apibrew/apibrew/pkg/abs"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/util"
	"google.golang.org/protobuf/types/known/structpb"
	"strings"
)

func ResourcePropertyToRecord(property *model.ResourceProperty, resource *model.Resource) abs.RecordLike {
	properties := make(map[string]*structpb.Value)

	if property.Title != nil {
		properties["title"] = structpb.NewStringValue(*property.Title)
	}
	if property.Description != nil {
		properties["description"] = structpb.NewStringValue(*property.Description)
	}
	properties["type"] = structpb.NewStringValue(property.Type.String())

	if property.Type == model.ResourceProperty_LIST || property.Type == model.ResourceProperty_MAP {
		properties["item"] = structpb.NewStructValue(&structpb.Struct{Fields: ResourcePropertyToRecord(property.Item, resource).GetProperties()})
	}

	if property.Type == model.ResourceProperty_STRUCT {
		properties["typeRef"] = structpb.NewStringValue(*property.TypeRef)
	}

	if property.Type == model.ResourceProperty_ENUM {
		properties["enumValues"] = structpb.NewListValue(&structpb.ListValue{Values: util.ArrayMap(property.EnumValues, func(v string) *structpb.Value {
			return structpb.NewStringValue(v)
		})})
	}

	properties["primary"] = structpb.NewBoolValue(property.Primary)
	properties["required"] = structpb.NewBoolValue(property.Required)
	properties["length"] = structpb.NewNumberValue(float64(property.Length))
	properties["unique"] = structpb.NewBoolValue(property.Unique)
	properties["immutable"] = structpb.NewBoolValue(property.Immutable)
	properties["virtual"] = structpb.NewBoolValue(property.Virtual)

	if property.Reference != nil {
		referenceNamespace := property.Reference.Namespace
		if referenceNamespace == "" {
			referenceNamespace = resource.Namespace
		}

		if property.BackReference != nil {
			properties["backReference"] = structpb.NewStringValue(property.BackReference.Property)
		}

		properties["reference"] = structpb.NewStringValue(referenceNamespace + "/" + property.Reference.Resource)
	}

	properties["defaultValue"] = property.DefaultValue
	properties["exampleValue"] = property.ExampleValue

	properties["annotations"], _ = structpb.NewValue(convertMap(property.Annotations, func(v string) interface{} {
		return v
	}))

	MapSpecialColumnsToRecord(property, &properties)

	return &model.Record{
		Properties: properties,
	}
}

func ResourcePropertyFromRecord(propertyName string, record abs.RecordLike) *model.ResourceProperty {
	if record == nil {
		return nil
	}

	var reference *model.Reference
	var backReference *model.BackReference

	if record.GetProperties()["reference"] != nil {
		reference = &model.Reference{}
		if record.GetProperties()["reference"].GetStructValue() != nil {
			var referenceProperties = record.GetProperties()["reference"].GetStructValue().GetFields()
			reference.Resource = referenceProperties["resource"].GetStructValue().GetFields()["name"].GetStringValue()
			if referenceProperties["resource"].GetStructValue().GetFields()["namespace"] != nil && referenceProperties["resource"].GetStructValue().GetFields()["namespace"].GetStructValue() != nil {
				reference.Namespace = referenceProperties["resource"].GetStructValue().GetFields()["namespace"].GetStructValue().GetFields()["name"].GetStringValue()
			}

			if record.GetProperties()["backReference"] != nil && record.GetProperties()["backReference"].GetStringValue() != "" {
				backReference = &model.BackReference{
					Property: record.GetProperties()["backReference"].GetStringValue(),
				}
			}
		} else {
			var referenceParts = strings.Split(record.GetProperties()["reference"].GetStringValue(), "/")

			if len(referenceParts) == 1 {
				reference.Resource = referenceParts[0]
			} else if len(referenceParts) == 2 {
				reference.Resource = referenceParts[1]
				reference.Namespace = referenceParts[0]
			} else {
				panic("Invalid reference format")
			}

			// reference.Cascade //todo implement it

			if record.GetProperties()["backReference"] != nil && record.GetProperties()["backReference"].GetStringValue() != "" {
				backReference = &model.BackReference{
					Property: record.GetProperties()["backReference"].GetStringValue(),
				}
			}
		}
	}

	var resourceProperty = &model.ResourceProperty{
		Name:          propertyName,
		Type:          model.ResourceProperty_Type(model.ResourceProperty_Type_value[strings.ToUpper(record.GetProperties()["type"].GetStringValue())]),
		Required:      record.GetProperties()["required"].GetBoolValue(),
		Length:        uint32(record.GetProperties()["length"].GetNumberValue()),
		Unique:        record.GetProperties()["unique"].GetBoolValue(),
		Primary:       record.GetProperties()["primary"].GetBoolValue(),
		Immutable:     record.GetProperties()["immutable"].GetBoolValue(),
		DefaultValue:  record.GetProperties()["defaultValue"],
		ExampleValue:  record.GetProperties()["exampleValue"],
		Reference:     reference,
		BackReference: backReference,
		Annotations: convertMap(record.GetProperties()["annotations"].GetStructValue().AsMap(), func(v interface{}) string {
			return v.(string)
		}),
	}

	if record.GetProperties()["virtual"] != nil {
		resourceProperty.Virtual = record.GetProperties()["virtual"].GetBoolValue()
	}

	if record.GetProperties()["title"] != nil {
		resourceProperty.Title = new(string)
		*resourceProperty.Title = record.GetProperties()["title"].GetStringValue()
	}

	if record.GetProperties()["description"] != nil {
		resourceProperty.Description = new(string)
		*resourceProperty.Description = record.GetProperties()["description"].GetStringValue()
	}

	if resourceProperty.Type == model.ResourceProperty_LIST || resourceProperty.Type == model.ResourceProperty_MAP {
		resourceProperty.Item = ResourcePropertyFromRecord("", &model.Record{
			Properties: record.GetProperties()["item"].GetStructValue().GetFields(),
		})
	}

	if resourceProperty.Type == model.ResourceProperty_STRUCT {
		resourceProperty.TypeRef = new(string)
		*resourceProperty.TypeRef = record.GetProperties()["typeRef"].GetStringValue()
	}

	if resourceProperty.Type == model.ResourceProperty_ENUM {
		resourceProperty.EnumValues = util.ArrayMap(record.GetProperties()["enumValues"].GetListValue().GetValues(), func(v *structpb.Value) string {
			return v.GetStringValue()
		})
	}

	var properties = record.GetProperties()
	MapSpecialColumnsFromRecord(resourceProperty, &properties)

	return resourceProperty
}
