// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: model/batch.proto

package model

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type BatchHeader_BatchMode int32

const (
	BatchHeader_CREATE BatchHeader_BatchMode = 0
	BatchHeader_UPDATE BatchHeader_BatchMode = 1
	BatchHeader_DELETE BatchHeader_BatchMode = 2
)

// Enum value maps for BatchHeader_BatchMode.
var (
	BatchHeader_BatchMode_name = map[int32]string{
		0: "CREATE",
		1: "UPDATE",
		2: "DELETE",
	}
	BatchHeader_BatchMode_value = map[string]int32{
		"CREATE": 0,
		"UPDATE": 1,
		"DELETE": 2,
	}
)

func (x BatchHeader_BatchMode) Enum() *BatchHeader_BatchMode {
	p := new(BatchHeader_BatchMode)
	*p = x
	return p
}

func (x BatchHeader_BatchMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BatchHeader_BatchMode) Descriptor() protoreflect.EnumDescriptor {
	return file_model_batch_proto_enumTypes[0].Descriptor()
}

func (BatchHeader_BatchMode) Type() protoreflect.EnumType {
	return &file_model_batch_proto_enumTypes[0]
}

func (x BatchHeader_BatchMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BatchHeader_BatchMode.Descriptor instead.
func (BatchHeader_BatchMode) EnumDescriptor() ([]byte, []int) {
	return file_model_batch_proto_rawDescGZIP(), []int{0, 0}
}

type BatchHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mode        BatchHeader_BatchMode `protobuf:"varint,1,opt,name=mode,proto3,enum=model.BatchHeader_BatchMode" json:"mode,omitempty"`
	Annotations map[string]string     `protobuf:"bytes,103,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *BatchHeader) Reset() {
	*x = BatchHeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_batch_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchHeader) ProtoMessage() {}

func (x *BatchHeader) ProtoReflect() protoreflect.Message {
	mi := &file_model_batch_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchHeader.ProtoReflect.Descriptor instead.
func (*BatchHeader) Descriptor() ([]byte, []int) {
	return file_model_batch_proto_rawDescGZIP(), []int{0}
}

func (x *BatchHeader) GetMode() BatchHeader_BatchMode {
	if x != nil {
		return x.Mode
	}
	return BatchHeader_CREATE
}

func (x *BatchHeader) GetAnnotations() map[string]string {
	if x != nil {
		return x.Annotations
	}
	return nil
}

type BatchRecordsPart struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace string            `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Resource  string            `protobuf:"bytes,2,opt,name=resource,proto3" json:"resource,omitempty"`
	Values    []*structpb.Value `protobuf:"bytes,3,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *BatchRecordsPart) Reset() {
	*x = BatchRecordsPart{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_batch_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchRecordsPart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchRecordsPart) ProtoMessage() {}

func (x *BatchRecordsPart) ProtoReflect() protoreflect.Message {
	mi := &file_model_batch_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchRecordsPart.ProtoReflect.Descriptor instead.
func (*BatchRecordsPart) Descriptor() ([]byte, []int) {
	return file_model_batch_proto_rawDescGZIP(), []int{1}
}

func (x *BatchRecordsPart) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *BatchRecordsPart) GetResource() string {
	if x != nil {
		return x.Resource
	}
	return ""
}

func (x *BatchRecordsPart) GetValues() []*structpb.Value {
	if x != nil {
		return x.Values
	}
	return nil
}

type Batch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header       *BatchHeader        `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Resources    []*Resource         `protobuf:"bytes,2,rep,name=resources,proto3" json:"resources,omitempty"`
	BatchRecords []*BatchRecordsPart `protobuf:"bytes,3,rep,name=batchRecords,proto3" json:"batchRecords,omitempty"`
}

func (x *Batch) Reset() {
	*x = Batch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_batch_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Batch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Batch) ProtoMessage() {}

func (x *Batch) ProtoReflect() protoreflect.Message {
	mi := &file_model_batch_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Batch.ProtoReflect.Descriptor instead.
func (*Batch) Descriptor() ([]byte, []int) {
	return file_model_batch_proto_rawDescGZIP(), []int{2}
}

func (x *Batch) GetHeader() *BatchHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *Batch) GetResources() []*Resource {
	if x != nil {
		return x.Resources
	}
	return nil
}

func (x *Batch) GetBatchRecords() []*BatchRecordsPart {
	if x != nil {
		return x.BatchRecords
	}
	return nil
}

var File_model_batch_proto protoreflect.FileDescriptor

var file_model_batch_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x12, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xf7, 0x01, 0x0a, 0x0b, 0x42, 0x61, 0x74, 0x63, 0x68, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x12, 0x30, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x48, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x04,
	0x6d, 0x6f, 0x64, 0x65, 0x12, 0x45, 0x0a, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x67, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x41, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x3e, 0x0a, 0x10, 0x41,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x2f, 0x0a, 0x09, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x52, 0x45, 0x41,
	0x54, 0x45, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x10, 0x01,
	0x12, 0x0a, 0x0a, 0x06, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x02, 0x22, 0x7c, 0x0a, 0x10,
	0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x50, 0x61, 0x72, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x2e, 0x0a, 0x06, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x9f, 0x01, 0x0a, 0x05, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x12, 0x2a, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x42, 0x61, 0x74,
	0x63, 0x68, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x12, 0x2d, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x52, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12,
	0x3b, 0x0a, 0x0c, 0x62, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x50, 0x61, 0x72, 0x74, 0x52, 0x0c,
	0x62, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x42, 0x26, 0x5a, 0x24,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x72,
	0x65, 0x77, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x72, 0x65, 0x77, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_model_batch_proto_rawDescOnce sync.Once
	file_model_batch_proto_rawDescData = file_model_batch_proto_rawDesc
)

func file_model_batch_proto_rawDescGZIP() []byte {
	file_model_batch_proto_rawDescOnce.Do(func() {
		file_model_batch_proto_rawDescData = protoimpl.X.CompressGZIP(file_model_batch_proto_rawDescData)
	})
	return file_model_batch_proto_rawDescData
}

var file_model_batch_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_model_batch_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_model_batch_proto_goTypes = []interface{}{
	(BatchHeader_BatchMode)(0), // 0: model.BatchHeader.BatchMode
	(*BatchHeader)(nil),        // 1: model.BatchHeader
	(*BatchRecordsPart)(nil),   // 2: model.BatchRecordsPart
	(*Batch)(nil),              // 3: model.Batch
	nil,                        // 4: model.BatchHeader.AnnotationsEntry
	(*structpb.Value)(nil),     // 5: google.protobuf.Value
	(*Resource)(nil),           // 6: model.Resource
}
var file_model_batch_proto_depIdxs = []int32{
	0, // 0: model.BatchHeader.mode:type_name -> model.BatchHeader.BatchMode
	4, // 1: model.BatchHeader.annotations:type_name -> model.BatchHeader.AnnotationsEntry
	5, // 2: model.BatchRecordsPart.values:type_name -> google.protobuf.Value
	1, // 3: model.Batch.header:type_name -> model.BatchHeader
	6, // 4: model.Batch.resources:type_name -> model.Resource
	2, // 5: model.Batch.batchRecords:type_name -> model.BatchRecordsPart
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_model_batch_proto_init() }
func file_model_batch_proto_init() {
	if File_model_batch_proto != nil {
		return
	}
	file_model_record_proto_init()
	file_model_resource_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_model_batch_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchHeader); i {
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
		file_model_batch_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchRecordsPart); i {
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
		file_model_batch_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Batch); i {
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
			RawDescriptor: file_model_batch_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_model_batch_proto_goTypes,
		DependencyIndexes: file_model_batch_proto_depIdxs,
		EnumInfos:         file_model_batch_proto_enumTypes,
		MessageInfos:      file_model_batch_proto_msgTypes,
	}.Build()
	File_model_batch_proto = out.File
	file_model_batch_proto_rawDesc = nil
	file_model_batch_proto_goTypes = nil
	file_model_batch_proto_depIdxs = nil
}
