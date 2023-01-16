package mapping

import (
	"data-handler/model"
	"data-handler/service/system"
	"google.golang.org/protobuf/types/known/structpb"
)

/*
Flags: &model.ResourceFlags{
			ReadOnlyRecords:    record.Properties.AsMap()["readOnlyRecords"].(bool),
			UniqueRecord:       record.Properties.AsMap()["uniqueRecord"].(bool),
			KeepHistory:        record.Properties.AsMap()["keepHistory"].(bool),
			AutoCreated:        record.Properties.AsMap()["autoCreated"].(bool),
			DisableMigration:   record.Properties.AsMap()["disableMigration"].(bool),
			DisableAudit:       record.Properties.AsMap()["disableAudit"].(bool),
			DoPrimaryKeyLookup: record.Properties.AsMap()["doPrimaryKeyLookup"].(bool),
		},
*/

func ResourceToRecord(resource *model.Resource) *model.Record {
	properties := make(map[string]interface{})

	properties["name"] = resource.Name
	properties["namespace"] = resource.Namespace
	properties["dataSource"] = resource.SourceConfig.DataSource
	properties["mapping"] = resource.SourceConfig.Mapping
	properties["type"] = int32(resource.DataType.Number())
	properties["readOnlyRecords"] = resource.Flags.ReadOnlyRecords
	properties["uniqueRecord"] = resource.Flags.UniqueRecord
	properties["keepHistory"] = resource.Flags.KeepHistory
	properties["autoCreated"] = resource.Flags.AutoCreated
	properties["disableMigration"] = resource.Flags.DisableMigration
	properties["disableAudit"] = resource.Flags.DisableAudit
	properties["doPrimaryKeyLookup"] = resource.Flags.DoPrimaryKeyLookup

	structProperties, err := structpb.NewStruct(properties)

	if err != nil {
		panic(err)
	}

	return &model.Record{
		Id:         resource.Id,
		Resource:   system.ResourceResource.Name,
		DataType:   resource.DataType,
		Properties: structProperties,
		AuditData:  resource.AuditData,
		Version:    resource.Version,
	}
}

func ResourceFromRecord(record *model.Record) *model.Resource {
	if record == nil {
		return nil
	}

	var resource = &model.Resource{
		Id:        record.Id,
		DataType:  model.DataType(int32(record.Properties.AsMap()["type"].(float64))),
		AuditData: record.AuditData,
		Version:   record.Version,
		Name:      record.Properties.AsMap()["name"].(string),
		Namespace: record.Properties.AsMap()["namespace"].(string),
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
