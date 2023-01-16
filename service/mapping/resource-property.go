package mapping

import (
	"data-handler/model"
	"data-handler/service/system"
	"google.golang.org/protobuf/types/known/structpb"
)

func ResourcePropertyToRecord(property *model.ResourceProperty, resource *model.Resource) *model.Record {
	properties := make(map[string]interface{})

	properties["name"] = property.Name
	properties["type"] = int32(property.Type.Number())
	properties["resource"] = resource.Id
	properties["required"] = property.Required
	properties["sourcePrimary"] = property.Primary
	properties["length"] = property.Length
	properties["unique"] = property.Unique

	sourceConfig := property.SourceConfig.(*model.ResourceProperty_Mapping)

	properties["sourceType"] = 0
	properties["sourceMapping"] = sourceConfig.Mapping.Mapping
	properties["sourceDef"] = sourceConfig.Mapping.SourceDef
	properties["sourceAutoGeneration"] = int32(sourceConfig.Mapping.AutoGeneration.Number())

	structProperties, err := structpb.NewStruct(properties)

	if err != nil {
		panic(err)
	}

	return &model.Record{
		Resource:   system.ResourcePropertyResource.Name,
		DataType:   model.DataType_SYSTEM,
		Properties: structProperties,
	}
}

func ResourcePropertyFromRecord(record *model.Record) *model.ResourceProperty {
	if record == nil {
		return nil
	}

	var resource = &model.ResourceProperty{
		Name: record.Properties.AsMap()["name"].(string),
		Type: model.ResourcePropertyType(record.Properties.AsMap()["type"].(float64)),
		SourceConfig: &model.ResourceProperty_Mapping{
			Mapping: &model.ResourcePropertyMappingConfig{
				Mapping:        record.Properties.AsMap()["sourceMapping"].(string),
				SourceDef:      record.Properties.AsMap()["sourceDef"].(string),
				AutoGeneration: model.AutoGenerationType(record.Properties.AsMap()["sourceAutoGeneration"].(float64)),
			},
		},
		Primary:  record.Properties.AsMap()["sourcePrimary"].(bool),
		Required: record.Properties.AsMap()["required"].(bool),
		Length:   uint32(record.Properties.AsMap()["length"].(float64)),
		Unique:   record.Properties.AsMap()["unique"].(bool),
	}

	return resource
}
