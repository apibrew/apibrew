package mapping

import (
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/system"
	"google.golang.org/protobuf/types/known/structpb"
)

func ExtensionToRecord(dataSource *model.Extension) *model.Record {
	properties := make(map[string]*structpb.Value)

	properties["name"] = structpb.NewStringValue(dataSource.Name)
	properties["description"] = structpb.NewStringValue(dataSource.Description)
	properties["namespace"] = structpb.NewStringValue(dataSource.Namespace)
	properties["resource"] = structpb.NewStringValue(dataSource.Resource)

	if dataSource.Server != nil {
		properties["serverHost"] = structpb.NewStringValue(dataSource.Server.Host)
		properties["serverPort"] = structpb.NewNumberValue(float64(dataSource.Server.Port))
	}

	return &model.Record{
		Id:         dataSource.Id,
		Resource:   system.DataSourceResource.Name,
		DataType:   dataSource.Type,
		Properties: properties,
		AuditData:  dataSource.AuditData,
		Version:    dataSource.Version,
	}
}

func ExtensionFromRecord(record *model.Record) *model.Extension {
	if record == nil {
		return nil
	}

	result := &model.Extension{
		Id:        record.Id,
		Name:      record.Properties["name"].GetStringValue(),
		Type:      record.DataType,
		Namespace: record.Properties["namespace"].GetStringValue(),
		Resource:  record.Properties["resource"].GetStringValue(),
		Server: &model.ExtensionServer{
			Host: record.Properties["serverHost"].GetStringValue(),
			Port: int32(record.Properties["serverPort"].GetNumberValue()),
		},
		AuditData: record.AuditData,
		Version:   record.Version,
	}

	if record.Properties["description"] != nil {
		result.Description = record.Properties["description"].GetStringValue()
	}

	return result
}
