package mapping

import (
	"github.com/apibrew/apibrew/pkg/model"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/structpb"
)

func ExtensionToRecord(extension *model.Extension) *model.Record {
	properties := make(map[string]*structpb.Value)

	properties["name"] = structpb.NewStringValue(extension.Name)
	properties["description"] = structpb.NewStringValue(extension.Description)

	properties["selector"] = EventSelectorToStruct(extension.Selector)

	properties["order"] = structpb.NewNumberValue(float64(extension.Order))
	properties["finalizes"] = structpb.NewBoolValue(extension.Finalizes)
	properties["sync"] = structpb.NewBoolValue(extension.Sync)
	properties["responds"] = structpb.NewBoolValue(extension.Responds)

	properties["call"] = ExternalCallToStruct(extension.Call)

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

	if call.FunctionCall != nil {
		return structpb.NewStructValue(&structpb.Struct{Fields: map[string]*structpb.Value{
			"kind":         structpb.NewStringValue("functionCall"),
			"functionName": structpb.NewStringValue(call.FunctionCall.FunctionName),
			"host":         structpb.NewStringValue(call.FunctionCall.Host),
		}})
	} else if call.HttpCall != nil {
		return structpb.NewStructValue(&structpb.Struct{Fields: map[string]*structpb.Value{
			"kind":   structpb.NewStringValue("httpCall"),
			"uri":    structpb.NewStringValue(call.HttpCall.Uri),
			"method": structpb.NewStringValue(call.HttpCall.Method),
		}})
	} else {
		return nil
	}
}

func ExternalCallFromStruct(value *structpb.Value) *model.ExternalCall {
	if value == nil {
		return nil
	}

	switch value.GetStructValue().GetFields()["kind"].GetStringValue() {
	case "functionCall":
		return &model.ExternalCall{FunctionCall: &model.FunctionCall{
			Host:         value.GetStructValue().GetFields()["host"].GetStringValue(),
			FunctionName: value.GetStructValue().GetFields()["functionName"].GetStringValue(),
		}}
	case "httpCall":
		return &model.ExternalCall{HttpCall: &model.HttpCall{
			Uri:    value.GetStructValue().GetFields()["uri"].GetStringValue(),
			Method: value.GetStructValue().GetFields()["method"].GetStringValue(),
		}}
	default:
		return nil
	}
}

func EventSelectorToStruct(selector *model.EventSelector) *structpb.Value {
	uData, err := protojson.Marshal(selector)

	if err != nil {
		panic(err)
	}

	userStruct := new(structpb.Struct)
	err = protojson.Unmarshal(uData, userStruct)

	if err != nil {
		panic(err)
	}

	return structpb.NewStructValue(userStruct)
}

func EventSelectorFromStruct(value *structpb.Value) *model.EventSelector {
	uData, err := protojson.Marshal(value.GetStructValue())

	if err != nil {
		panic(err)
	}

	userStruct := new(model.EventSelector)
	err = protojson.Unmarshal(uData, userStruct)

	if err != nil {
		panic(err)
	}

	return userStruct
}

func ExtensionFromRecord(record *model.Record) *model.Extension {
	if record == nil {
		return nil
	}

	result := &model.Extension{
		Id:        record.Id,
		Name:      record.Properties["name"].GetStringValue(),
		Call:      ExternalCallFromStruct(record.Properties["call"]),
		Selector:  EventSelectorFromStruct(record.Properties["selector"]),
		Order:     int32(record.Properties["order"].GetNumberValue()),
		Finalizes: record.Properties["finalizes"].GetBoolValue(),
		Sync:      record.Properties["sync"].GetBoolValue(),
		Responds:  record.Properties["responds"].GetBoolValue(),
	}

	if record.Properties["description"] != nil {
		result.Description = record.Properties["description"].GetStringValue()
	}

	MapSpecialColumnsFromRecord(result, &record.Properties)

	return result
}
