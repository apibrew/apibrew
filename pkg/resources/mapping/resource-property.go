package mapping

import (
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/util"
	"google.golang.org/protobuf/types/known/structpb"
)

func ResourcePropertyToRecord(property *model.ResourceProperty, resource *model.Resource) *model.Record {
	properties := make(map[string]*structpb.Value)

	properties["name"] = structpb.NewStringValue(property.Name)
	if property.Title != nil {
		properties["title"] = structpb.NewStringValue(*property.Title)
	}
	if property.Description != nil {
		properties["description"] = structpb.NewStringValue(*property.Description)
	}
	properties["type"] = structpb.NewNumberValue(float64(property.Type.Number()))
	properties["subType"] = structpb.NewNumberValue(float64(property.Type.Number()))
	properties["resource"] = util.StructKv("id", resource.Id)
	properties["required"] = structpb.NewBoolValue(property.Required)
	properties["primary"] = structpb.NewBoolValue(property.Primary)
	properties["length"] = structpb.NewNumberValue(float64(property.Length))
	properties["unique"] = structpb.NewBoolValue(property.Unique)
	properties["immutable"] = structpb.NewBoolValue(property.Immutable)

	properties["mapping"] = structpb.NewStringValue(property.Mapping)
	properties["securityContext"] = SecurityContextToValue(property.SecurityContext)

	if property.Reference != nil {
		properties["reference_resource"] = util.StructKv2("name", property.Reference.ReferencedResource, "namespace", map[string]interface{}{
			"name": resource.Namespace,
		})
		properties["reference_cascade"] = structpb.NewBoolValue(property.Reference.Cascade)
	}

	properties["defaultValue"] = property.DefaultValue
	properties["exampleValue"] = property.ExampleValue
	properties["enumValues"] = structpb.NewListValue(&structpb.ListValue{Values: property.EnumValues})

	properties["annotations"], _ = structpb.NewValue(convertMap(property.Annotations, func(v string) interface{} {
		return v
	}))

	if property.Id == nil {
		property.Id = new(string)
	}

	mapSpecialColumnsToRecord(property, &properties)

	return &model.Record{
		Id:         *property.Id,
		Properties: properties,
	}
}

func ResourcePropertyFromRecord(record *model.Record) *model.ResourceProperty {
	if record == nil {
		return nil
	}

	var reference = &model.Reference{}
	var hasReference bool

	if record.Properties["reference_resource"] != nil {
		reference.ReferencedResource = record.Properties["reference_resource"].GetStructValue().GetFields()["name"].GetStringValue()
		hasReference = true
	}

	if record.Properties["reference_cascade"] != nil {
		reference.Cascade = record.Properties["reference_cascade"].GetBoolValue()
		hasReference = true
	}

	if !hasReference {
		reference = nil
	}

	var resourceProperty = &model.ResourceProperty{
		Id:              &record.Id,
		Name:            record.Properties["name"].GetStringValue(),
		Type:            model.ResourceProperty_Type(record.Properties["type"].GetNumberValue()),
		Mapping:         record.Properties["mapping"].GetStringValue(),
		Primary:         record.Properties["primary"].GetBoolValue(),
		Required:        record.Properties["required"].GetBoolValue(),
		Length:          uint32(record.Properties["length"].GetNumberValue()),
		Unique:          record.Properties["unique"].GetBoolValue(),
		Immutable:       record.Properties["immutable"].GetBoolValue(),
		SecurityContext: SecurityContextFromValue(record.Properties["securityContext"]),
		DefaultValue:    record.Properties["defaultValue"],
		ExampleValue:    record.Properties["exampleValue"],
		Reference:       reference,
		Annotations: convertMap(record.Properties["annotations"].GetStructValue().AsMap(), func(v interface{}) string {
			return v.(string)
		}),
	}

	if record.Properties["title"] != nil {
		resourceProperty.Title = new(string)
		*resourceProperty.Title = record.Properties["title"].GetStringValue()
	}

	if record.Properties["description"] != nil {
		resourceProperty.Description = new(string)
		*resourceProperty.Description = record.Properties["description"].GetStringValue()
	}

	mapSpecialColumnsFromRecord(resourceProperty, &record.Properties)

	return resourceProperty
}
