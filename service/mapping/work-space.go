package mapping

import (
	"data-handler/service/system"
	"data-handler/stub/model"
	"google.golang.org/protobuf/types/known/structpb"
)

func WorkSpaceToRecord(workSpace *model.WorkSpace) *model.Record {
	properties := make(map[string]interface{})

	properties["name"] = workSpace.Name
	properties["description"] = workSpace.Description
	properties["details"] = workSpace.Details

	structProperties, err := structpb.NewStruct(properties)

	if err != nil {
		panic(err)
	}

	return &model.Record{
		Id:         workSpace.Id,
		Resource:   system.WorkSpaceResource.Name,
		Type:       workSpace.Type,
		Properties: structProperties,
		AuditData:  workSpace.AuditData,
		Version:    workSpace.Version,
	}
}

func WorkSpaceFromRecord(record *model.Record) *model.WorkSpace {
	if record == nil {
		return nil
	}

	return &model.WorkSpace{
		Id:          record.Id,
		Type:        record.Type,
		Name:        record.Properties.AsMap()["name"].(string),
		Description: record.Properties.AsMap()["description"].(string),
		Details:     record.Properties.AsMap()["details"].(*structpb.Struct),
		AuditData:   record.AuditData,
		Version:     record.Version,
	}
}
