package mapping

import (
	"data-handler/model"
	"data-handler/service/system"
	"google.golang.org/protobuf/types/known/structpb"
)

func ResourcePropertyToRecord(Resource *model.Resource) *model.Record {
	properties := make(map[string]interface{})

	properties["name"] = Resource.Name
	properties["workspace"] = Resource.Workspace
	properties["dataSource"] = Resource.SourceConfig.DataSource
	properties["mapping"] = Resource.SourceConfig.Mapping

	structProperties, err := structpb.NewStruct(properties)

	if err != nil {
		panic(err)
	}

	return &model.Record{
		Id:         Resource.Id,
		Resource:   system.ResourceResource.Name,
		DataType:   Resource.DataType,
		Properties: structProperties,
		AuditData:  Resource.AuditData,
		Version:    Resource.Version,
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
