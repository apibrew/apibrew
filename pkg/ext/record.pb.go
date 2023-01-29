// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: ext/record.proto

package ext

import (
	model "github.com/tislib/data-handler/pkg/model"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ListRecordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace         string                   `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Resource          string                   `protobuf:"bytes,3,opt,name=resource,proto3" json:"resource,omitempty"`
	Query             *model.BooleanExpression `protobuf:"bytes,4,opt,name=query,proto3" json:"query,omitempty"`
	Limit             uint32                   `protobuf:"varint,5,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset            uint64                   `protobuf:"varint,6,opt,name=offset,proto3" json:"offset,omitempty"`
	UseHistory        bool                     `protobuf:"varint,7,opt,name=useHistory,proto3" json:"useHistory,omitempty"`
	ResolveReferences bool                     `protobuf:"varint,8,opt,name=resolveReferences,proto3" json:"resolveReferences,omitempty"`
	Annotations       map[string]string        `protobuf:"bytes,103,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ListRecordRequest) Reset() {
	*x = ListRecordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ext_record_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRecordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRecordRequest) ProtoMessage() {}

func (x *ListRecordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ext_record_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRecordRequest.ProtoReflect.Descriptor instead.
func (*ListRecordRequest) Descriptor() ([]byte, []int) {
	return file_ext_record_proto_rawDescGZIP(), []int{0}
}

func (x *ListRecordRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *ListRecordRequest) GetResource() string {
	if x != nil {
		return x.Resource
	}
	return ""
}

func (x *ListRecordRequest) GetQuery() *model.BooleanExpression {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *ListRecordRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListRecordRequest) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListRecordRequest) GetUseHistory() bool {
	if x != nil {
		return x.UseHistory
	}
	return false
}

func (x *ListRecordRequest) GetResolveReferences() bool {
	if x != nil {
		return x.ResolveReferences
	}
	return false
}

func (x *ListRecordRequest) GetAnnotations() map[string]string {
	if x != nil {
		return x.Annotations
	}
	return nil
}

type ListRecordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total   uint32          `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Content []*model.Record `protobuf:"bytes,2,rep,name=content,proto3" json:"content,omitempty"`
}

func (x *ListRecordResponse) Reset() {
	*x = ListRecordResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ext_record_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRecordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRecordResponse) ProtoMessage() {}

func (x *ListRecordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ext_record_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRecordResponse.ProtoReflect.Descriptor instead.
func (*ListRecordResponse) Descriptor() ([]byte, []int) {
	return file_ext_record_proto_rawDescGZIP(), []int{1}
}

func (x *ListRecordResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListRecordResponse) GetContent() []*model.Record {
	if x != nil {
		return x.Content
	}
	return nil
}

var File_ext_record_proto protoreflect.FileDescriptor

var file_ext_record_proto_rawDesc = []byte{
	0x0a, 0x10, 0x65, 0x78, 0x74, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x03, 0x65, 0x78, 0x74, 0x1a, 0x12, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x72,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x84, 0x03,
	0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x2e, 0x0a,
	0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x45, 0x78, 0x70, 0x72,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x75,
	0x73, 0x65, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0a, 0x75, 0x73, 0x65, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x2c, 0x0a, 0x11, 0x72,
	0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x52,
	0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x49, 0x0a, 0x0b, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x67, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27,
	0x2e, 0x65, 0x78, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x3e, 0x0a, 0x10, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0x53, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x12, 0x27, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x32, 0x53, 0x0a, 0x16, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x65, 0x78,
	0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x65, 0x78, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x28,
	0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x69, 0x73,
	0x6c, 0x69, 0x62, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2d, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x65, 0x78, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ext_record_proto_rawDescOnce sync.Once
	file_ext_record_proto_rawDescData = file_ext_record_proto_rawDesc
)

func file_ext_record_proto_rawDescGZIP() []byte {
	file_ext_record_proto_rawDescOnce.Do(func() {
		file_ext_record_proto_rawDescData = protoimpl.X.CompressGZIP(file_ext_record_proto_rawDescData)
	})
	return file_ext_record_proto_rawDescData
}

var file_ext_record_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_ext_record_proto_goTypes = []interface{}{
	(*ListRecordRequest)(nil),       // 0: ext.ListRecordRequest
	(*ListRecordResponse)(nil),      // 1: ext.ListRecordResponse
	nil,                             // 2: ext.ListRecordRequest.AnnotationsEntry
	(*model.BooleanExpression)(nil), // 3: model.BooleanExpression
	(*model.Record)(nil),            // 4: model.Record
}
var file_ext_record_proto_depIdxs = []int32{
	3, // 0: ext.ListRecordRequest.query:type_name -> model.BooleanExpression
	2, // 1: ext.ListRecordRequest.annotations:type_name -> ext.ListRecordRequest.AnnotationsEntry
	4, // 2: ext.ListRecordResponse.content:type_name -> model.Record
	0, // 3: ext.RecordExtensionService.List:input_type -> ext.ListRecordRequest
	1, // 4: ext.RecordExtensionService.List:output_type -> ext.ListRecordResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_ext_record_proto_init() }
func file_ext_record_proto_init() {
	if File_ext_record_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ext_record_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRecordRequest); i {
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
		file_ext_record_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRecordResponse); i {
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
			RawDescriptor: file_ext_record_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ext_record_proto_goTypes,
		DependencyIndexes: file_ext_record_proto_depIdxs,
		MessageInfos:      file_ext_record_proto_msgTypes,
	}.Build()
	File_ext_record_proto = out.File
	file_ext_record_proto_rawDesc = nil
	file_ext_record_proto_goTypes = nil
	file_ext_record_proto_depIdxs = nil
}
