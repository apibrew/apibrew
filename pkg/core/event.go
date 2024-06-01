package core

import (
	"github.com/apibrew/apibrew/pkg/abs"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/util"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Event_Action int32

const (
	Event_CREATE  Event_Action = 0
	Event_UPDATE  Event_Action = 1
	Event_DELETE  Event_Action = 2
	Event_GET     Event_Action = 3
	Event_LIST    Event_Action = 4
	Event_OPERATE Event_Action = 5 // for special cases
)

// Enum value maps for Event_Action.
var (
	Event_Action_name = map[int32]string{
		0: "CREATE",
		1: "UPDATE",
		2: "DELETE",
		3: "GET",
		4: "LIST",
		5: "OPERATE",
	}
	Event_Action_value = map[string]int32{
		"CREATE":  0,
		"UPDATE":  1,
		"DELETE":  2,
		"GET":     3,
		"LIST":    4,
		"OPERATE": 5,
	}
)

func (x Event_Action) Enum() *Event_Action {
	p := new(Event_Action)
	*p = x
	return p
}

func (x Event_Action) String() string {
	return Event_Action_name[int32(x)]
}

type EventSelector struct {
	Actions        []Event_Action           `protobuf:"varint,1,rep,packed,name=actions,proto3,enum=model.Event_Action" json:"actions,omitempty"`
	RecordSelector *model.BooleanExpression `protobuf:"bytes,2,opt,name=recordSelector,proto3" json:"recordSelector,omitempty"`
	Namespaces     []string                 `protobuf:"bytes,3,rep,name=namespaces,proto3" json:"namespaces,omitempty"`
	Resources      []string                 `protobuf:"bytes,4,rep,name=resources,proto3" json:"resources,omitempty"`
	Shallow        bool                     `protobuf:"varint,5,opt,name=shallow,proto3" json:"shallow,omitempty"`
	// star means all
	Ids []string `protobuf:"bytes,7,rep,name=ids,proto3" json:"ids,omitempty"`
	// star means all, empty means proceed
	Annotations map[string]string `protobuf:"bytes,103,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (e *EventSelector) ToProtoEventSelector() *model.EventSelector {
	return &model.EventSelector{
		Actions:        util.ArrayMap(e.Actions, func(i Event_Action) model.Event_Action { return model.Event_Action(i) }),
		RecordSelector: e.RecordSelector,
		Namespaces:     e.Namespaces,
		Resources:      e.Resources,
		Shallow:        e.Shallow,
		Ids:            e.Ids,
		Annotations:    e.Annotations,
	}
}

func FromProtoEventSelector(eventSelector *model.EventSelector) *EventSelector {
	return &EventSelector{
		Actions:        util.ArrayMap(eventSelector.Actions, func(i model.Event_Action) Event_Action { return Event_Action(i) }),
		RecordSelector: eventSelector.RecordSelector,
		Namespaces:     eventSelector.Namespaces,
		Resources:      eventSelector.Resources,
		Shallow:        eventSelector.Shallow,
		Ids:            eventSelector.Ids,
		Annotations:    eventSelector.Annotations,
	}
}

type Event struct {
	Id       string          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Action   Event_Action    `protobuf:"varint,2,opt,name=action,proto3,enum=model.Event_Action" json:"action,omitempty"`
	Resource *model.Resource `protobuf:"bytes,5,opt,name=resource,proto3" json:"resource,omitempty"`
	// CREATE,UPDATE - records are for incoming and outgoing records
	// GET - there will be only one record
	// LIST - result of the list operation
	Records []abs.RecordLike `protobuf:"bytes,6,rep,name=records,proto3" json:"records,omitempty"`
	// LIST - search params for the list operation
	RecordSearchParams *model.Event_RecordSearchParams `protobuf:"bytes,8,opt,name=recordSearchParams,proto3" json:"recordSearchParams,omitempty"`
	// If true, this will be last event on operation list
	Finalizes bool `protobuf:"varint,9,opt,name=finalizes,proto3" json:"finalizes,omitempty"`
	// If true, backend will wait for processing this event before sending next one on operation list
	Sync    bool                   `protobuf:"varint,10,opt,name=sync,proto3" json:"sync,omitempty"`
	Time    *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=time,proto3" json:"time,omitempty"`
	Total   uint64                 `protobuf:"varint,12,opt,name=total,proto3" json:"total,omitempty"`
	Shallow bool                   `protobuf:"varint,13,opt,name=shallow,proto3" json:"shallow,omitempty"`
	// Request annotations
	Annotations map[string]string `protobuf:"bytes,103,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Error       *model.Error      `protobuf:"bytes,104,opt,name=error,proto3" json:"error,omitempty"`
}

func (e *Event) ToProtoEvent() *model.Event {
	return &model.Event{
		Id:                 e.Id,
		Action:             model.Event_Action(e.Action),
		Resource:           e.Resource,
		Records:            abs.RecordLikeAsRecords(e.Records),
		RecordSearchParams: e.RecordSearchParams,
		Finalizes:          e.Finalizes,
		Sync:               e.Sync,
		Time:               e.Time,
		Total:              e.Total,
		Shallow:            e.Shallow,
		Annotations:        e.Annotations,
		Error:              e.Error,
	}
}

func FromProtoEvent(event *model.Event) *Event {
	return &Event{
		Id:                 event.Id,
		Action:             Event_Action(event.Action),
		Resource:           event.Resource,
		Records:            abs.RecordLikeAsRecords2(event.Records),
		RecordSearchParams: event.RecordSearchParams,
		Finalizes:          event.Finalizes,
		Sync:               event.Sync,
		Time:               event.Time,
		Total:              event.Total,
		Shallow:            event.Shallow,
		Annotations:        event.Annotations,
		Error:              event.Error,
	}
}
