// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: stub/event-channel.proto

package stub

import (
	model "github.com/apibrew/apibrew/pkg/model"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type EventPollRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token      string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	ChannelKey string `protobuf:"bytes,2,opt,name=channelKey,proto3" json:"channelKey,omitempty"`
}

func (x *EventPollRequest) Reset() {
	*x = EventPollRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stub_event_channel_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventPollRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventPollRequest) ProtoMessage() {}

func (x *EventPollRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stub_event_channel_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventPollRequest.ProtoReflect.Descriptor instead.
func (*EventPollRequest) Descriptor() ([]byte, []int) {
	return file_stub_event_channel_proto_rawDescGZIP(), []int{0}
}

func (x *EventPollRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *EventPollRequest) GetChannelKey() string {
	if x != nil {
		return x.ChannelKey
	}
	return ""
}

type EventWriteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string       `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Event *model.Event `protobuf:"bytes,3,opt,name=event,proto3" json:"event,omitempty"`
}

func (x *EventWriteRequest) Reset() {
	*x = EventWriteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stub_event_channel_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventWriteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventWriteRequest) ProtoMessage() {}

func (x *EventWriteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stub_event_channel_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventWriteRequest.ProtoReflect.Descriptor instead.
func (*EventWriteRequest) Descriptor() ([]byte, []int) {
	return file_stub_event_channel_proto_rawDescGZIP(), []int{1}
}

func (x *EventWriteRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *EventWriteRequest) GetEvent() *model.Event {
	if x != nil {
		return x.Event
	}
	return nil
}

type EventWriteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token      string       `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	ChannelKey string       `protobuf:"bytes,2,opt,name=channelKey,proto3" json:"channelKey,omitempty"`
	Event      *model.Event `protobuf:"bytes,3,opt,name=event,proto3" json:"event,omitempty"`
}

func (x *EventWriteResponse) Reset() {
	*x = EventWriteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stub_event_channel_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventWriteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventWriteResponse) ProtoMessage() {}

func (x *EventWriteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stub_event_channel_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventWriteResponse.ProtoReflect.Descriptor instead.
func (*EventWriteResponse) Descriptor() ([]byte, []int) {
	return file_stub_event_channel_proto_rawDescGZIP(), []int{2}
}

func (x *EventWriteResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *EventWriteResponse) GetChannelKey() string {
	if x != nil {
		return x.ChannelKey
	}
	return ""
}

func (x *EventWriteResponse) GetEvent() *model.Event {
	if x != nil {
		return x.Event
	}
	return nil
}

var File_stub_event_channel_proto protoreflect.FileDescriptor

var file_stub_event_channel_proto_rawDesc = []byte{
	0x0a, 0x18, 0x73, 0x74, 0x75, 0x62, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2d, 0x63, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x73, 0x74, 0x75, 0x62,
	0x1a, 0x11, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x48, 0x0a, 0x10, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x50,
	0x6f, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4b, 0x65, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4b, 0x65, 0x79,
	0x22, 0x4d, 0x0a, 0x11, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x22, 0x0a, 0x05, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x22,
	0x6e, 0x0a, 0x12, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x63,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4b, 0x65, 0x79, 0x12, 0x22, 0x0a, 0x05, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x32,
	0x7e, 0x0a, 0x0c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12,
	0x30, 0x0a, 0x04, 0x50, 0x6f, 0x6c, 0x6c, 0x12, 0x16, 0x2e, 0x73, 0x74, 0x75, 0x62, 0x2e, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x50, 0x6f, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x30,
	0x01, 0x12, 0x3c, 0x0a, 0x05, 0x57, 0x72, 0x69, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x73, 0x74, 0x75,
	0x62, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x73, 0x74, 0x75, 0x62, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x70,
	0x69, 0x62, 0x72, 0x65, 0x77, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x72, 0x65, 0x77, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x73, 0x74, 0x75, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_stub_event_channel_proto_rawDescOnce sync.Once
	file_stub_event_channel_proto_rawDescData = file_stub_event_channel_proto_rawDesc
)

func file_stub_event_channel_proto_rawDescGZIP() []byte {
	file_stub_event_channel_proto_rawDescOnce.Do(func() {
		file_stub_event_channel_proto_rawDescData = protoimpl.X.CompressGZIP(file_stub_event_channel_proto_rawDescData)
	})
	return file_stub_event_channel_proto_rawDescData
}

var file_stub_event_channel_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_stub_event_channel_proto_goTypes = []interface{}{
	(*EventPollRequest)(nil),   // 0: stub.EventPollRequest
	(*EventWriteRequest)(nil),  // 1: stub.EventWriteRequest
	(*EventWriteResponse)(nil), // 2: stub.EventWriteResponse
	(*model.Event)(nil),        // 3: model.Event
}
var file_stub_event_channel_proto_depIdxs = []int32{
	3, // 0: stub.EventWriteRequest.event:type_name -> model.Event
	3, // 1: stub.EventWriteResponse.event:type_name -> model.Event
	0, // 2: stub.EventChannel.Poll:input_type -> stub.EventPollRequest
	1, // 3: stub.EventChannel.Write:input_type -> stub.EventWriteRequest
	3, // 4: stub.EventChannel.Poll:output_type -> model.Event
	2, // 5: stub.EventChannel.Write:output_type -> stub.EventWriteResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_stub_event_channel_proto_init() }
func file_stub_event_channel_proto_init() {
	if File_stub_event_channel_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_stub_event_channel_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventPollRequest); i {
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
		file_stub_event_channel_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventWriteRequest); i {
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
		file_stub_event_channel_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventWriteResponse); i {
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
			RawDescriptor: file_stub_event_channel_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_stub_event_channel_proto_goTypes,
		DependencyIndexes: file_stub_event_channel_proto_depIdxs,
		MessageInfos:      file_stub_event_channel_proto_msgTypes,
	}.Build()
	File_stub_event_channel_proto = out.File
	file_stub_event_channel_proto_rawDesc = nil
	file_stub_event_channel_proto_goTypes = nil
	file_stub_event_channel_proto_depIdxs = nil
}
