package mapping

import (
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/util"
	"google.golang.org/protobuf/types/known/structpb"
)

func ResourcePropertyToRecord(property *model.ResourceProperty, resource *model.Resource) *model.Record {
	properties := make(map[string]*structpb.Value)

	properties["name"] = structpb.NewStringValue(property.Name)
	properties["type"] = structpb.NewNumberValue(float64(property.Type.Number()))
	properties["resource"] = util.StructKv("id", resource.Id)
	properties["required"] = structpb.NewBoolValue(property.Required)
	properties["sourcePrimary"] = structpb.NewBoolValue(property.Primary)
	properties["length"] = structpb.NewNumberValue(float64(property.Length))
	properties["unique"] = structpb.NewBoolValue(property.Unique)
	properties["immutable"] = structpb.NewBoolValue(property.Immutable)

	properties["mapping"] = structpb.NewStringValue(property.Mapping)
	properties["securityContext"] = SecurityContextToValue(resource.SecurityContext)

	if property.Reference != nil {
		properties["reference_resource"] = util.StructKv("name", property.Reference.ReferencedResource)
		properties["reference_cascade"] = structpb.NewBoolValue(property.Reference.Cascade)
	}

	return &model.Record{
		DataType:   model.DataType_SYSTEM,
		Properties: properties,
	}
}

func ResourcePropertyFromRecord(record *model.Record) *model.ResourceProperty {
	if record == nil {
		return nil
	}

	var reference = &model.Reference{}

	if record.Properties["reference_resource"] != nil {
		reference.ReferencedResource = record.Properties["reference_resource"].GetStructValue().GetFields()["name"].GetStringValue()
	}

	if record.Properties["reference_cascade"] != nil {
		reference.Cascade = record.Properties["reference_cascade"].GetBoolValue()
	}

	var resource = &model.ResourceProperty{
		Name:            record.Properties["name"].GetStringValue(),
		Type:            model.ResourcePropertyType(record.Properties["type"].GetNumberValue()),
		Mapping:         record.Properties["mapping"].GetStringValue(),
		Primary:         record.Properties["sourcePrimary"].GetBoolValue(),
		Required:        record.Properties["required"].GetBoolValue(),
		Length:          uint32(record.Properties["length"].GetNumberValue()),
		Unique:          record.Properties["unique"].GetBoolValue(),
		Immutable:       record.Properties["immutable"].GetBoolValue(),
		SecurityContext: SecurityContextFromValue(record.Properties["securityContext"]),
		Reference:       reference,
	}

	return resource
}
