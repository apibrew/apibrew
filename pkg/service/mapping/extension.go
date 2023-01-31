package mapping

import (
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/system"
	"github.com/tislib/data-handler/pkg/util"
	"google.golang.org/protobuf/types/known/structpb"
)

func ExtensionToRecord(extension *model.Extension) *model.Record {
	properties := make(map[string]*structpb.Value)

	properties["name"] = structpb.NewStringValue(extension.Name)
	properties["description"] = structpb.NewStringValue(extension.Description)
	properties["namespace"] = structpb.NewStringValue(extension.Namespace)
	properties["resource"] = structpb.NewStringValue(extension.Resource)

	if extension.Server != nil {
		properties["serverHost"] = structpb.NewStringValue(extension.Server.Host)
		properties["serverPort"] = structpb.NewNumberValue(float64(extension.Server.Port))
	}

	listVal, err := structpb.NewList(util.ArrayMap(extension.Operations, func(op *model.ExtensionOperation) interface{} {
		return map[string]interface{}{
			"operationType": op.OperationType.Number(),
			"step":          op.Step.Number(),
			"sync":          op.Sync,
		}
	}))

	if err != nil {
		panic(err)
	}

	properties["operations"] = structpb.NewListValue(listVal)

	return &model.Record{
		Id:         extension.Id,
		Resource:   system.DataSourceResource.Name,
		DataType:   extension.Type,
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

	if record.Properties["operations"] != nil && record.Properties["operations"].GetListValue() != nil {
		for _, value := range record.Properties["operations"].GetListValue().Values {
			fields := value.GetStructValue().GetFields()

			result.Operations = append(result.Operations, &model.ExtensionOperation{
				OperationType: model.ExtensionOperationType(fields["operationType"].GetNumberValue()),
				Step:          model.ExtensionOperationStep(fields["step"].GetNumberValue()),
				Sync:          fields["sync"].GetBoolValue(),
			})
		}
	}

	if record.Properties["description"] != nil {
		result.Description = record.Properties["description"].GetStringValue()
	}

	return result
}
