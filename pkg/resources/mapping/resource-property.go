package mapping

import (
	"github.com/apibrew/apibrew/pkg/formats/unstructured"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/util"
	"google.golang.org/protobuf/types/known/structpb"
	"strings"
)

func ResourcePropertyToRecord(property *model.ResourceProperty, resource *model.Resource) unstructured.Unstructured {
	properties := make(map[string]*structpb.Value)

	if property.Title != nil {
		properties["title"] = structpb.NewStringValue(*property.Title)
	}
	if property.Description != nil {
		properties["description"] = structpb.NewStringValue(*property.Description)
	}
	properties["type"] = structpb.NewStringValue(property.Type.String())

	if property.Type == model.ResourceProperty_LIST || property.Type == model.ResourceProperty_MAP {
		properties["item"] = structpb.NewStructValue(&structpb.Struct{Fields: ResourcePropertyToRecord(property.Item, resource)})
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

func ResourcePropertyFromRecord(propertyName string, record unstructured.Unstructured) *model.ResourceProperty {
	if record == nil {
		return nil
	}

	var reference *model.Reference
	var backReference *model.BackReference

	if record["reference"] != nil {
		reference = &model.Reference{}
		if record["reference"] != nil {
			var referenceProperties = record["reference"].(unstructured.Unstructured)
			reference.Resource = referenceProperties["resource"].(unstructured.Unstructured)["name"].(string)
			if referenceProperties["resource"].(unstructured.Unstructured)["namespace"] != nil {
				reference.Namespace = referenceProperties["resource"].(unstructured.Unstructured)["namespace"].(unstructured.Unstructured)["name"].(string)
			}

			if record["backReference"] != nil && record["backReference"].(string) != "" {
				backReference = &model.BackReference{
					Property: record["backReference"].(string),
				}
			}
		} else {
			var referenceParts = strings.Split(record["reference"].(string), "/")

			if len(referenceParts) == 1 {
				reference.Resource = referenceParts[0]
			} else if len(referenceParts) == 2 {
				reference.Resource = referenceParts[1]
				reference.Namespace = referenceParts[0]
			} else {
				panic("Invalid reference format")
			}

			// reference.Cascade //todo implement it

			if record["backReference"] != nil && record["backReference"].GetStringValue() != "" {
				backReference = &model.BackReference{
					Property: record["backReference"].GetStringValue(),
				}
			}
		}
	}

	var resourceProperty = &model.ResourceProperty{
		Name:          propertyName,
		Type:          model.ResourceProperty_Type(model.ResourceProperty_Type_value[strings.ToUpper(record["type"].GetStringValue())]),
		Required:      record["required"].GetBoolValue(),
		Length:        uint32(record["length"].GetNumberValue()),
		Unique:        record["unique"].GetBoolValue(),
		Primary:       record["primary"].GetBoolValue(),
		Immutable:     record["immutable"].GetBoolValue(),
		DefaultValue:  record["defaultValue"],
		ExampleValue:  record["exampleValue"],
		Reference:     reference,
		BackReference: backReference,
		Annotations: convertMap(record["annotations"].AsMap(), func(v interface{}) string {
			return v.(string)
		}),
	}

	if record["virtual"] != nil {
		resourceProperty.Virtual = record["virtual"].GetBoolValue()
	}

	if record["title"] != nil {
		resourceProperty.Title = new(string)
		*resourceProperty.Title = record["title"].GetStringValue()
	}

	if record["description"] != nil {
		resourceProperty.Description = new(string)
		*resourceProperty.Description = record["description"].GetStringValue()
	}

	if resourceProperty.Type == model.ResourceProperty_LIST || resourceProperty.Type == model.ResourceProperty_MAP {
		resourceProperty.Item = ResourcePropertyFromRecord("", &model.Record{
			Properties: record["item"].GetFields(),
		})
	}

	if resourceProperty.Type == model.ResourceProperty_STRUCT {
		resourceProperty.TypeRef = new(string)
		*resourceProperty.TypeRef = record["typeRef"].GetStringValue()
	}

	if resourceProperty.Type == model.ResourceProperty_ENUM {
		resourceProperty.EnumValues = util.ArrayMap(record["enumValues"].GetListValue().GetValues(), func(v *structpb.Value) string {
			return v.GetStringValue()
		})
	}

	MapSpecialColumnsFromRecord(resourceProperty, &record)

	return resourceProperty
}
