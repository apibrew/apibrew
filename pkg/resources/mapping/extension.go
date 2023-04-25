package mapping

import (
	"github.com/tislib/apibrew/pkg/model"
	"google.golang.org/protobuf/types/known/structpb"
)

func ExtensionToRecord(extension *model.Extension) *model.Record {
	properties := make(map[string]*structpb.Value)

	properties["name"] = structpb.NewStringValue(extension.Name)
	properties["description"] = structpb.NewStringValue(extension.Description)
	properties["namespace"] = structpb.NewStringValue(extension.Namespace)
	properties["resource"] = structpb.NewStringValue(extension.Resource)

	if extension.After != nil {
		properties["after"] = structpb.NewStructValue(&structpb.Struct{Fields: map[string]*structpb.Value{
			"create": ExternalCallToStruct(extension.After.Create),
			"update": ExternalCallToStruct(extension.After.Update),
			"delete": ExternalCallToStruct(extension.After.Delete),
			"get":    ExternalCallToStruct(extension.After.Get),
			"list":   ExternalCallToStruct(extension.After.List),
			"all":    ExternalCallToStruct(extension.After.All),
			"sync":   structpb.NewBoolValue(extension.After.Sync),
		}})
	}

	if extension.Before != nil {
		properties["before"] = structpb.NewStructValue(&structpb.Struct{Fields: map[string]*structpb.Value{
			"create": ExternalCallToStruct(extension.Before.Create),
			"update": ExternalCallToStruct(extension.Before.Update),
			"delete": ExternalCallToStruct(extension.Before.Delete),
			"get":    ExternalCallToStruct(extension.Before.Get),
			"list":   ExternalCallToStruct(extension.Before.List),
			"all":    ExternalCallToStruct(extension.Before.All),
			"sync":   structpb.NewBoolValue(extension.Before.Sync),
		}})
	}

	if extension.Instead != nil {
		properties["instead"] = structpb.NewStructValue(&structpb.Struct{Fields: map[string]*structpb.Value{
			"create":   ExternalCallToStruct(extension.Instead.Create),
			"update":   ExternalCallToStruct(extension.Instead.Update),
			"delete":   ExternalCallToStruct(extension.Instead.Delete),
			"get":      ExternalCallToStruct(extension.Instead.Get),
			"list":     ExternalCallToStruct(extension.Instead.List),
			"all":      ExternalCallToStruct(extension.Instead.All),
			"finalize": structpb.NewBoolValue(extension.Instead.Finalize),
		}})
	}

	MapSpecialColumnsToRecord(extension, &properties)

	return &model.Record{
		Id:         extension.Id,
		Properties: properties,
	}
}

func ExternalCallToStruct(call *model.ExternalCall) *structpb.Value {
	if call == nil {
		return nil
	}

	switch callKind := call.Kind.(type) {
	case *model.ExternalCall_FunctionCall:
		return structpb.NewStructValue(&structpb.Struct{Fields: map[string]*structpb.Value{
			"kind":         structpb.NewStringValue("functionCall"),
			"functionName": structpb.NewStringValue(callKind.FunctionCall.FunctionName),
			"host":         structpb.NewStringValue(callKind.FunctionCall.Host),
		}})
	case *model.ExternalCall_HttpCall:
		return structpb.NewStructValue(&structpb.Struct{Fields: map[string]*structpb.Value{
			"kind":   structpb.NewStringValue("httpCall"),
			"uri":    structpb.NewStringValue(callKind.HttpCall.Uri),
			"method": structpb.NewStringValue(callKind.HttpCall.Method),
		}})
	default:
		return nil
	}
}

func ExternalCallFromStruct(value *structpb.Value) *model.ExternalCall {
	if value == nil {
		return nil
	}

	switch value.GetStructValue().GetFields()["kind"].GetStringValue() {
	case "functionCall":
		return &model.ExternalCall{Kind: &model.ExternalCall_FunctionCall{FunctionCall: &model.FunctionCall{
			Host:         value.GetStructValue().GetFields()["host"].GetStringValue(),
			FunctionName: value.GetStructValue().GetFields()["functionName"].GetStringValue(),
		}}}
	case "httpCall":
		return &model.ExternalCall{Kind: &model.ExternalCall_HttpCall{HttpCall: &model.HttpCall{
			Uri:    value.GetStructValue().GetFields()["uri"].GetStringValue(),
			Method: value.GetStructValue().GetFields()["method"].GetStringValue(),
		}}}
	default:
		return nil
	}
}

func ExtensionFromRecord(record *model.Record) *model.Extension {
	if record == nil {
		return nil
	}

	result := &model.Extension{
		Id:        record.Id,
		Name:      record.Properties["name"].GetStringValue(),
		Namespace: record.Properties["namespace"].GetStringValue(),
		Resource:  record.Properties["resource"].GetStringValue(),
	}

	if record.Properties["description"] != nil {
		result.Description = record.Properties["description"].GetStringValue()
	}

	if record.Properties["before"] != nil {
		val := record.Properties["before"].GetStructValue()
		result.Before = &model.Extension_Before{
			All:    ExternalCallFromStruct(val.Fields["all"]),
			Create: ExternalCallFromStruct(val.Fields["create"]),
			Update: ExternalCallFromStruct(val.Fields["update"]),
			Delete: ExternalCallFromStruct(val.Fields["delete"]),
			Get:    ExternalCallFromStruct(val.Fields["get"]),
			List:   ExternalCallFromStruct(val.Fields["list"]),
			Sync:   val.Fields["sync"].GetBoolValue(),
		}
	}

	if record.Properties["after"] != nil {
		val := record.Properties["after"].GetStructValue()
		result.After = &model.Extension_After{
			All:    ExternalCallFromStruct(val.Fields["all"]),
			Create: ExternalCallFromStruct(val.Fields["create"]),
			Update: ExternalCallFromStruct(val.Fields["update"]),
			Delete: ExternalCallFromStruct(val.Fields["delete"]),
			Get:    ExternalCallFromStruct(val.Fields["get"]),
			List:   ExternalCallFromStruct(val.Fields["list"]),
			Sync:   val.Fields["sync"].GetBoolValue(),
		}
	}

	if record.Properties["instead"] != nil {
		val := record.Properties["instead"].GetStructValue()
		result.Instead = &model.Extension_Instead{
			All:      ExternalCallFromStruct(val.Fields["all"]),
			Create:   ExternalCallFromStruct(val.Fields["create"]),
			Update:   ExternalCallFromStruct(val.Fields["update"]),
			Delete:   ExternalCallFromStruct(val.Fields["delete"]),
			Get:      ExternalCallFromStruct(val.Fields["get"]),
			List:     ExternalCallFromStruct(val.Fields["list"]),
			Finalize: val.Fields["finalize"].GetBoolValue(),
		}
	}

	MapSpecialColumnsFromRecord(result, &record.Properties)

	return result
}
