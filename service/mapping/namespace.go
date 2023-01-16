package mapping

import (
	"data-handler/model"
	"data-handler/service/system"
	"google.golang.org/protobuf/types/known/structpb"
)

func NamespaceToRecord(namespace *model.Namespace) *model.Record {
	properties := make(map[string]interface{})

	properties["name"] = namespace.Name
	properties["description"] = namespace.Description
	if namespace.Details != nil {
		properties["details"] = namespace.Details.AsMap()
	}

	structProperties, err := structpb.NewStruct(properties)

	if err != nil {
		panic(err)
	}

	return &model.Record{
		Id:         namespace.Id,
		Resource:   system.NamespaceResource.Name,
		DataType:   namespace.Type,
		Properties: structProperties,
		AuditData:  namespace.AuditData,
		Version:    namespace.Version,
	}
}

func NamespaceFromRecord(record *model.Record) *model.Namespace {
	if record == nil {
		return nil
	}

	result := &model.Namespace{
		Id:        record.Id,
		Type:      record.DataType,
		AuditData: record.AuditData,
		Version:   record.Version,
	}

	if record.Properties.AsMap()["name"] != nil {
		result.Name = record.Properties.AsMap()["name"].(string)
	}

	if record.Properties.AsMap()["description"] != nil {
		result.Description = record.Properties.AsMap()["description"].(string)
	}

	if record.Properties.AsMap()["details"] != nil {
		result.Details = record.Properties.Fields["details"].GetStructValue()
	}

	return result
}
