package extramappings

import (
	"github.com/apibrew/apibrew/pkg/formats/unstructured"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resource_model"
	"github.com/apibrew/apibrew/pkg/util"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func EventToProto(result *resource_model.Event) *model.Event {
	var event = new(model.Event)

	event.Id = result.Id
	event.Action = model.Event_Action(model.Event_Action_value[string(result.Action)])
	if result.Time != nil {
		event.Time = timestamppb.New(*result.Time)
	}
	if result.ActionDescription != nil {
		event.ActionDescription = *result.ActionDescription
	}
	if result.ActionSummary != nil {
		event.ActionSummary = *result.ActionSummary
	}
	event.Annotations = result.Annotations
	if result.Finalizes != nil {
		event.Finalizes = *result.Finalizes
	}
	if result.Sync != nil {
		event.Sync = *result.Sync
	}
	if result.RecordSearchParams != nil {
		var query *model.BooleanExpression

		if result.RecordSearchParams.Query != nil {
			query = BooleanExpressionToProto(*result.RecordSearchParams.Query)
		}

		event.RecordSearchParams = &model.Event_RecordSearchParams{
			Query:             query,
			Limit:             uint32(util.DePointer(result.RecordSearchParams.Limit, 0)),
			Offset:            uint64(util.DePointer(result.RecordSearchParams.Offset, 0)),
			ResolveReferences: result.RecordSearchParams.ResolveReferences,
		}
	}
	event.Records = util.ArrayMapX(result.Records, func(t *resource_model.Record) *model.Record {
		return &model.Record{
			Properties: resource_model.RecordMapperInstance.ToRecord(t).Properties["properties"].GetStructValue().Fields,
		}
	})

	event.Resource = ResourceFrom(result.Resource)

	if result.Error != nil {
		event.Error = ErrorToProto(*result.Error)
	}

	event.Total = uint64(util.DePointer(result.Total, 0))
	event.ActionName = util.DePointer(result.ActionName, "")
	if result.Input != nil {
		var err error
		event.Input, err = unstructured.ToValue(result.Input)
		if err != nil {
			panic(err)
		}
	}
	if result.Output != nil {
		var err error
		event.Output, err = unstructured.ToValue(result.Output)
		if err != nil {
			panic(err)
		}
	}

	return event
}

func EventFromProto(event *model.Event) *resource_model.Event {
	extensionEvent := new(resource_model.Event)
	extensionEvent.Id = event.Id
	extensionEvent.Action = resource_model.ExtensionAction(model.Event_Action_name[int32(event.Action)])
	if event.Time != nil {
		extensionEvent.Time = new(time.Time)
		*extensionEvent.Time = event.Time.AsTime()
	}
	extensionEvent.ActionDescription = &event.ActionDescription
	extensionEvent.ActionSummary = &event.ActionSummary
	extensionEvent.Annotations = event.Annotations
	extensionEvent.Finalizes = &event.Finalizes
	extensionEvent.Sync = &event.Sync
	if event.RecordSearchParams != nil {
		var query *resource_model.ExtensionBooleanExpression

		if event.RecordSearchParams.Query != nil {
			query = util.Pointer(BooleanExpressionFromProto(event.RecordSearchParams.Query))
		}

		extensionEvent.RecordSearchParams = &resource_model.RecordSearchParams{
			Query:             query,
			Limit:             util.Pointer(int32(event.RecordSearchParams.Limit)),
			Offset:            util.Pointer(int32(event.RecordSearchParams.Offset)),
			ResolveReferences: event.RecordSearchParams.ResolveReferences,
		}
	}
	extensionEvent.Records = util.ArrayMapX(event.Records, func(item *model.Record) *resource_model.Record {
		return resource_model.RecordMapperInstance.FromRecord(&model.Record{
			Properties: map[string]*structpb.Value{
				"id":         item.Properties["id"],
				"properties": structpb.NewStructValue(&structpb.Struct{Fields: item.Properties}),
			},
		})
	})

	extensionEvent.Resource = ResourceTo(event.Resource)

	if event.Error != nil {
		extensionEvent.Error = util.Pointer(ErrorFromProto(event.Error))
	}

	extensionEvent.Total = util.Pointer(int64(event.Total))
	extensionEvent.ActionName = util.Pointer(event.ActionName)
	if event.Input != nil {
		extensionEvent.Input = unstructured.FromValue(event.Input)
	}
	if event.Input != nil {
		extensionEvent.Output = unstructured.FromValue(event.Input)
	}

	return extensionEvent
}
