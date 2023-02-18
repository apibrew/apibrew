package mapping

import (
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/util"
	"google.golang.org/protobuf/types/known/structpb"
)

func ExtensionToRecord(extension *model.RemoteExtension) *model.Record {
	properties := make(map[string]*structpb.Value)

	properties["name"] = structpb.NewStringValue(extension.Name)
	properties["description"] = structpb.NewStringValue(extension.Description)
	properties["namespace"] = structpb.NewStringValue(extension.Config.Namespace)
	properties["resource"] = structpb.NewStringValue(extension.Config.Resource)

	if extension.Server != nil {
		properties["serverHost"] = structpb.NewStringValue(extension.Server.Host)
		properties["serverPort"] = structpb.NewNumberValue(float64(extension.Server.Port))
	}

	listVal, err := structpb.NewList(util.ArrayMap(extension.Config.Operations, func(op *model.ExtensionOperation) interface{} {
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
		DataType:   extension.Type,
		Properties: properties,
		AuditData:  extension.AuditData,
		Version:    extension.Version,
	}
}

func ExtensionFromRecord(record *model.Record) *model.RemoteExtension {
	if record == nil {
		return nil
	}

	result := &model.RemoteExtension{
		Id:   record.Id,
		Name: record.Properties["name"].GetStringValue(),
		Type: record.DataType,
		Config: &model.ExtensionConfig{
			Namespace: record.Properties["namespace"].GetStringValue(),
			Resource:  record.Properties["resource"].GetStringValue(),
		},
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

			result.Config.Operations = append(result.Config.Operations, &model.ExtensionOperation{
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
