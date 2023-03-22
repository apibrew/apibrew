package mapping

import (
	"github.com/tislib/data-handler/pkg/model"
	"google.golang.org/protobuf/types/known/structpb"
)

func ExtensionToRecord(extension *model.Extension) *model.Record {
	properties := make(map[string]*structpb.Value)

	properties["name"] = structpb.NewStringValue(extension.Name)
	properties["description"] = structpb.NewStringValue(extension.Description)
	properties["namespace"] = structpb.NewStringValue(extension.Namespace)
	properties["resource"] = structpb.NewStringValue(extension.Resource)

	return &model.Record{
		Id:         extension.Id,
		Properties: properties,
		AuditData:  extension.AuditData,
		Version:    extension.Version,
	}
}

func ExtensionFromRecord(record *model.Record) *model.Extension {
	if record == nil {
		return nil
	}

	result := &model.Extension{
		Id:        record.Id,
		Name:      record.Properties["name"].GetStringValue(),
		AuditData: record.AuditData,
		Version:   record.Version,
	}

	if record.Properties["description"] != nil {
		result.Description = record.Properties["description"].GetStringValue()
	}

	return result
}
