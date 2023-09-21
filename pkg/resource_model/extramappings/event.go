package extramappings

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resource_model"
	"github.com/apibrew/apibrew/pkg/resources/mapping"
	"github.com/apibrew/apibrew/pkg/util"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func EventToProto(result *resource_model.ExtensionEvent) *model.Event {
	var event = new(model.Event)

	event.Id = result.Id
	event.Action = model.Event_Action(model.Event_Action_value[string(result.Action)])
	event.Ids = result.Ids
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

	event.Resource = resourceFrom(result.Resource)

	return event
}

func EventFromProto(event *model.Event) *resource_model.ExtensionEvent {
	extensionEvent := new(resource_model.ExtensionEvent)
	extensionEvent.Id = event.Id
	extensionEvent.Action = resource_model.ExtensionAction(model.Event_Action_name[int32(event.Action)])
	extensionEvent.Ids = event.Ids
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

		extensionEvent.RecordSearchParams = &resource_model.ExtensionRecordSearchParams{
			Query:             query,
			Limit:             util.Pointer(int32(event.RecordSearchParams.Limit)),
			Offset:            util.Pointer(int32(event.RecordSearchParams.Offset)),
			ResolveReferences: event.RecordSearchParams.ResolveReferences,
		}
	}
	extensionEvent.Records = util.ArrayMapX(event.Records, func(item *model.Record) *resource_model.Record {
		return resource_model.RecordMapperInstance.FromRecord(&model.Record{
			Id: item.Id,
			Properties: map[string]*structpb.Value{
				"id":         item.Properties["id"],
				"properties": structpb.NewStructValue(&structpb.Struct{Fields: item.Properties}),
			},
		})
	})

	extensionEvent.Resource = resourceTo(event.Resource)

	return extensionEvent
}

func resourceTo(resource *model.Resource) *resource_model.Resource {
	if resource == nil {
		return nil
	}
	resourceRec := mapping.ResourceToRecord(resource)
	return resource_model.ResourceMapperInstance.FromRecord(resourceRec)
}

func resourceFrom(resource *resource_model.Resource) *model.Resource {
	if resource == nil {
		return nil
	}
	resourceRec := resource_model.ResourceMapperInstance.ToRecord(resource)
	return mapping.ResourceFromRecord(resourceRec)
}
