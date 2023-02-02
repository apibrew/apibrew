package mapping

import (
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/resources"
	"google.golang.org/protobuf/types/known/structpb"
)

func ResourcePropertyToRecord(property *model.ResourceProperty, resource *model.Resource) *model.Record {
	properties := make(map[string]*structpb.Value)

	properties["name"] = structpb.NewStringValue(property.Name)
	properties["type"] = structpb.NewNumberValue(float64(property.Type.Number()))
	properties["resource"] = structpb.NewStringValue(resource.Id)
	properties["required"] = structpb.NewBoolValue(property.Required)
	properties["sourcePrimary"] = structpb.NewBoolValue(property.Primary)
	properties["length"] = structpb.NewNumberValue(float64(property.Length))
	properties["unique"] = structpb.NewBoolValue(property.Unique)

	sourceConfig := property.SourceConfig.(*model.ResourceProperty_Mapping)

	properties["sourceType"] = structpb.NewNumberValue(0)
	properties["sourceMapping"] = structpb.NewStringValue(sourceConfig.Mapping.Mapping)
	properties["sourceDef"] = structpb.NewStringValue(sourceConfig.Mapping.SourceDef)
	properties["sourceAutoGeneration"] = structpb.NewNumberValue(float64(sourceConfig.Mapping.AutoGeneration.Number()))
	properties["securityContext"] = SecurityContextToValue(resource.SecurityContext)

	return &model.Record{
		Resource:   resources.ResourcePropertyResource.Name,
		DataType:   model.DataType_SYSTEM,
		Properties: properties,
	}
}

func ResourcePropertyFromRecord(record *model.Record) *model.ResourceProperty {
	if record == nil {
		return nil
	}

	var resource = &model.ResourceProperty{
		Name: record.Properties["name"].GetStringValue(),
		Type: model.ResourcePropertyType(record.Properties["type"].GetNumberValue()),
		SourceConfig: &model.ResourceProperty_Mapping{
			Mapping: &model.ResourcePropertyMappingConfig{
				Mapping:        record.Properties["sourceMapping"].GetStringValue(),
				SourceDef:      record.Properties["sourceDef"].GetStringValue(),
				AutoGeneration: model.AutoGenerationType(record.Properties["sourceAutoGeneration"].GetNumberValue()),
			},
		},
		Primary:         record.Properties["sourcePrimary"].GetBoolValue(),
		Required:        record.Properties["required"].GetBoolValue(),
		Length:          uint32(record.Properties["length"].GetNumberValue()),
		Unique:          record.Properties["unique"].GetBoolValue(),
		SecurityContext: SecurityContextFromValue(record.Properties["securityContext"]),
	}

	return resource
}
