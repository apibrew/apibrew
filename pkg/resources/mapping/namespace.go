package mapping

import (
	"github.com/tislib/data-handler/pkg/model"
	"google.golang.org/protobuf/types/known/structpb"
)

func NamespaceToRecord(namespace *model.Namespace) *model.Record {
	properties := make(map[string]*structpb.Value)

	properties["name"] = structpb.NewStringValue(namespace.Name)
	properties["description"] = structpb.NewStringValue(namespace.Description)
	if namespace.Details != nil {
		properties["details"] = structpb.NewStructValue(namespace.Details)
	}

	properties["securityContext"] = SecurityContextToValue(namespace.SecurityContext)

	return &model.Record{
		Id:         namespace.Id,
		DataType:   namespace.Type,
		Properties: properties,
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

	if record.Properties["name"] != nil {
		result.Name = record.Properties["name"].GetStringValue()
	}

	if record.Properties["description"] != nil {
		result.Description = record.Properties["description"].GetStringValue()
	}

	if record.Properties["details"] != nil {
		result.Details = record.Properties["details"].GetStructValue()
	}

	result.SecurityContext = SecurityContextFromValue(record.Properties["securityContext"])

	return result
}
