package mapping

import (
	"data-handler/service/system"
	"data-handler/stub/model"
	"google.golang.org/protobuf/types/known/structpb"
)

func WorkspaceToRecord(workspace *model.Workspace) *model.Record {
	properties := make(map[string]interface{})

	properties["name"] = workspace.Name
	properties["description"] = workspace.Description
	if workspace.Details != nil {
		properties["details"] = workspace.Details.AsMap()
	}

	structProperties, err := structpb.NewStruct(properties)

	if err != nil {
		panic(err)
	}

	return &model.Record{
		Id:         workspace.Id,
		Resource:   system.WorkspaceResource.Name,
		Type:       workspace.Type,
		Properties: structProperties,
		AuditData:  workspace.AuditData,
		Version:    workspace.Version,
	}
}

func WorkspaceFromRecord(record *model.Record) *model.Workspace {
	if record == nil {
		return nil
	}

	result := &model.Workspace{
		Id:        record.Id,
		Type:      record.Type,
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
