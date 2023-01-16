package mapping

import (
	"data-handler/model"
	"data-handler/service/system"
	"google.golang.org/protobuf/types/known/structpb"
)

func ResourceToRecord(Resource *model.Resource) *model.Record {
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

func ResourceFromRecord(record *model.Record) *model.Resource {
	if record == nil {
		return nil
	}

	var resource = &model.Resource{
		Id:        record.Id,
		DataType:  record.DataType,
		AuditData: record.AuditData,
		Version:   record.Version,
		Name:      record.Properties.AsMap()["name"].(string),
		Workspace: record.Properties.AsMap()["workspace"].(string),
		SourceConfig: &model.ResourceSourceConfig{
			DataSource: record.Properties.AsMap()["dataSource"].(string),
			Mapping:    record.Properties.AsMap()["mapping"].(string),
		},
		Flags: &model.ResourceFlags{
			ReadOnlyRecords:    record.Properties.AsMap()["readOnlyRecords"].(bool),
			UniqueRecord:       record.Properties.AsMap()["uniqueRecord"].(bool),
			KeepHistory:        record.Properties.AsMap()["keepHistory"].(bool),
			AutoCreated:        record.Properties.AsMap()["autoCreated"].(bool),
			DisableMigration:   record.Properties.AsMap()["disableMigration"].(bool),
			DisableAudit:       record.Properties.AsMap()["disableAudit"].(bool),
			DoPrimaryKeyLookup: record.Properties.AsMap()["doPrimaryKeyLookup"].(bool),
		},
	}

	return resource
}
