// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: model/event.proto

package model

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
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
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Event_Action) Descriptor() protoreflect.EnumDescriptor {
	return file_model_event_proto_enumTypes[0].Descriptor()
}

func (Event_Action) Type() protoreflect.EnumType {
	return &file_model_event_proto_enumTypes[0]
}

func (x Event_Action) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Event_Action.Descriptor instead.
func (Event_Action) EnumDescriptor() ([]byte, []int) {
	return file_model_event_proto_rawDescGZIP(), []int{0, 0}
}

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Action            Event_Action `protobuf:"varint,2,opt,name=action,proto3,enum=model.Event_Action" json:"action,omitempty"`
	ActionSummary     string       `protobuf:"bytes,3,opt,name=actionSummary,proto3" json:"actionSummary,omitempty"`
	ActionDescription string       `protobuf:"bytes,4,opt,name=actionDescription,proto3" json:"actionDescription,omitempty"`
	Resource          *Resource    `protobuf:"bytes,5,opt,name=resource,proto3" json:"resource,omitempty"`
	// CREATE,UPDATE - records are for incoming and outgoing records
	// GET - there will be only one record
	// LIST - result of the list operation
	Records []*Record `protobuf:"bytes,6,rep,name=records,proto3" json:"records,omitempty"`
	// GET - there will be only one id, for getting record with id
	// DELETE - there will be multiple ids, for deleting multiple records
	Ids []string `protobuf:"bytes,7,rep,name=ids,proto3" json:"ids,omitempty"`
	// LIST - search params for the list operation
	RecordSearchParams *Event_RecordSearchParams `protobuf:"bytes,8,opt,name=recordSearchParams,proto3" json:"recordSearchParams,omitempty"`
	// If true, this will be last event on operation list
	Finalizes bool `protobuf:"varint,9,opt,name=finalizes,proto3" json:"finalizes,omitempty"`
	// If true, backend will wait for processing this event before sending next one on operation list
	Sync bool                   `protobuf:"varint,10,opt,name=sync,proto3" json:"sync,omitempty"`
	Time *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=time,proto3" json:"time,omitempty"`
	// Request annotations
	Annotations map[string]string `protobuf:"bytes,103,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_model_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_model_event_proto_rawDescGZIP(), []int{0}
}

func (x *Event) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Event) GetAction() Event_Action {
	if x != nil {
		return x.Action
	}
	return Event_CREATE
}

func (x *Event) GetActionSummary() string {
	if x != nil {
		return x.ActionSummary
	}
	return ""
}

func (x *Event) GetActionDescription() string {
	if x != nil {
		return x.ActionDescription
	}
	return ""
}

func (x *Event) GetResource() *Resource {
	if x != nil {
		return x.Resource
	}
	return nil
}

func (x *Event) GetRecords() []*Record {
	if x != nil {
		return x.Records
	}
	return nil
}

func (x *Event) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *Event) GetRecordSearchParams() *Event_RecordSearchParams {
	if x != nil {
		return x.RecordSearchParams
	}
	return nil
}

func (x *Event) GetFinalizes() bool {
	if x != nil {
		return x.Finalizes
	}
	return false
}

func (x *Event) GetSync() bool {
	if x != nil {
		return x.Sync
	}
	return false
}

func (x *Event) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *Event) GetAnnotations() map[string]string {
	if x != nil {
		return x.Annotations
	}
	return nil
}

// Events pass through selectors, if selector returns true, event will be processed
// Selector returns true if no selector fails.
// For example, if you passed empty selector, it will return true for all events.
type EventSelector struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Actions        []Event_Action     `protobuf:"varint,1,rep,packed,name=actions,proto3,enum=model.Event_Action" json:"actions,omitempty"`
	RecordSelector *BooleanExpression `protobuf:"bytes,2,opt,name=recordSelector,proto3" json:"recordSelector,omitempty"`
	Namespaces     []string           `protobuf:"bytes,3,rep,name=namespaces,proto3" json:"namespaces,omitempty"`
	Resources      []string           `protobuf:"bytes,4,rep,name=resources,proto3" json:"resources,omitempty"`
	// star means all
	Ids []string `protobuf:"bytes,7,rep,name=ids,proto3" json:"ids,omitempty"`
	// star means all, empty means proceed
	Annotations map[string]string `protobuf:"bytes,103,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *EventSelector) Reset() {
	*x = EventSelector{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_event_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventSelector) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventSelector) ProtoMessage() {}

func (x *EventSelector) ProtoReflect() protoreflect.Message {
	mi := &file_model_event_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventSelector.ProtoReflect.Descriptor instead.
func (*EventSelector) Descriptor() ([]byte, []int) {
	return file_model_event_proto_rawDescGZIP(), []int{1}
}

func (x *EventSelector) GetActions() []Event_Action {
	if x != nil {
		return x.Actions
	}
	return nil
}

func (x *EventSelector) GetRecordSelector() *BooleanExpression {
	if x != nil {
		return x.RecordSelector
	}
	return nil
}

func (x *EventSelector) GetNamespaces() []string {
	if x != nil {
		return x.Namespaces
	}
	return nil
}

func (x *EventSelector) GetResources() []string {
	if x != nil {
		return x.Resources
	}
	return nil
}

func (x *EventSelector) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *EventSelector) GetAnnotations() map[string]string {
	if x != nil {
		return x.Annotations
	}
	return nil
}

type Event_RecordSearchParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query             *BooleanExpression `protobuf:"bytes,4,opt,name=query,proto3" json:"query,omitempty"`
	Limit             uint32             `protobuf:"varint,5,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset            uint64             `protobuf:"varint,6,opt,name=offset,proto3" json:"offset,omitempty"`
	ResolveReferences []string           `protobuf:"bytes,8,rep,name=resolveReferences,proto3" json:"resolveReferences,omitempty"`
}

func (x *Event_RecordSearchParams) Reset() {
	*x = Event_RecordSearchParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_event_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event_RecordSearchParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event_RecordSearchParams) ProtoMessage() {}

func (x *Event_RecordSearchParams) ProtoReflect() protoreflect.Message {
	mi := &file_model_event_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event_RecordSearchParams.ProtoReflect.Descriptor instead.
func (*Event_RecordSearchParams) Descriptor() ([]byte, []int) {
	return file_model_event_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Event_RecordSearchParams) GetQuery() *BooleanExpression {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *Event_RecordSearchParams) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *Event_RecordSearchParams) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *Event_RecordSearchParams) GetResolveReferences() []string {
	if x != nil {
		return x.ResolveReferences
	}
	return nil
}

var File_model_event_proto protoreflect.FileDescriptor

var file_model_event_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x14, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x12, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa5, 0x06, 0x0a, 0x05, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x2b, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x13, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x24, 0x0a, 0x0d, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75,
	0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x2c, 0x0a, 0x11, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x11, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x12, 0x27, 0x0a, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x52, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73,
	0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x69, 0x64, 0x73, 0x12, 0x4f, 0x0a, 0x12, 0x72,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x12, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x1c, 0x0a, 0x09,
	0x66, 0x69, 0x6e, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x09, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x79,
	0x6e, 0x63, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x73, 0x79, 0x6e, 0x63, 0x12, 0x2e,
	0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x3f,
	0x0a, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x67, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x2e, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a,
	0xa0, 0x01, 0x0a, 0x12, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x2e, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x42, 0x6f,
	0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x12, 0x2c, 0x0a, 0x11, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x52,
	0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x11, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x73, 0x1a, 0x3e, 0x0a, 0x10, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0x4c, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0a, 0x0a, 0x06,
	0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x50, 0x44, 0x41,
	0x54, 0x45, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x02,
	0x12, 0x07, 0x0a, 0x03, 0x47, 0x45, 0x54, 0x10, 0x03, 0x12, 0x08, 0x0a, 0x04, 0x4c, 0x49, 0x53,
	0x54, 0x10, 0x04, 0x12, 0x0b, 0x0a, 0x07, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x45, 0x10, 0x05,
	0x22, 0xd9, 0x02, 0x0a, 0x0d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x12, 0x2d, 0x0a, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x40, 0x0a, 0x0e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x53, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x53, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03,
	0x69, 0x64, 0x73, 0x12, 0x47, 0x0a, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x67, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x41,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x3e, 0x0a, 0x10,
	0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x26, 0x5a, 0x24,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x72,
	0x65, 0x77, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x72, 0x65, 0x77, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_model_event_proto_rawDescOnce sync.Once
	file_model_event_proto_rawDescData = file_model_event_proto_rawDesc
)

func file_model_event_proto_rawDescGZIP() []byte {
	file_model_event_proto_rawDescOnce.Do(func() {
		file_model_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_model_event_proto_rawDescData)
	})
	return file_model_event_proto_rawDescData
}

var file_model_event_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_model_event_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_model_event_proto_goTypes = []interface{}{
	(Event_Action)(0),                // 0: model.Event.Action
	(*Event)(nil),                    // 1: model.Event
	(*EventSelector)(nil),            // 2: model.EventSelector
	(*Event_RecordSearchParams)(nil), // 3: model.Event.RecordSearchParams
	nil,                              // 4: model.Event.AnnotationsEntry
	nil,                              // 5: model.EventSelector.AnnotationsEntry
	(*Resource)(nil),                 // 6: model.Resource
	(*Record)(nil),                   // 7: model.Record
	(*timestamppb.Timestamp)(nil),    // 8: google.protobuf.Timestamp
	(*BooleanExpression)(nil),        // 9: model.BooleanExpression
}
var file_model_event_proto_depIdxs = []int32{
	0,  // 0: model.Event.action:type_name -> model.Event.Action
	6,  // 1: model.Event.resource:type_name -> model.Resource
	7,  // 2: model.Event.records:type_name -> model.Record
	3,  // 3: model.Event.recordSearchParams:type_name -> model.Event.RecordSearchParams
	8,  // 4: model.Event.time:type_name -> google.protobuf.Timestamp
	4,  // 5: model.Event.annotations:type_name -> model.Event.AnnotationsEntry
	0,  // 6: model.EventSelector.actions:type_name -> model.Event.Action
	9,  // 7: model.EventSelector.recordSelector:type_name -> model.BooleanExpression
	5,  // 8: model.EventSelector.annotations:type_name -> model.EventSelector.AnnotationsEntry
	9,  // 9: model.Event.RecordSearchParams.query:type_name -> model.BooleanExpression
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_model_event_proto_init() }
func file_model_event_proto_init() {
	if File_model_event_proto != nil {
		return
	}
	file_model_resource_proto_init()
	file_model_record_proto_init()
	file_model_query_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_model_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_model_event_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventSelector); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_model_event_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event_RecordSearchParams); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_model_event_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_model_event_proto_goTypes,
		DependencyIndexes: file_model_event_proto_depIdxs,
		EnumInfos:         file_model_event_proto_enumTypes,
		MessageInfos:      file_model_event_proto_msgTypes,
	}.Build()
	File_model_event_proto = out.File
	file_model_event_proto_rawDesc = nil
	file_model_event_proto_goTypes = nil
	file_model_event_proto_depIdxs = nil
}
