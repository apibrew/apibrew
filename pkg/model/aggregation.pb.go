// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: model/aggregation.proto

package model

import (
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

type AggregationItem_Algorithm int32

const (
	AggregationItem_COUNT AggregationItem_Algorithm = 0
	AggregationItem_SUM   AggregationItem_Algorithm = 1
	AggregationItem_AVG   AggregationItem_Algorithm = 2
	AggregationItem_MAX   AggregationItem_Algorithm = 3
	AggregationItem_MIN   AggregationItem_Algorithm = 4
)

// Enum value maps for AggregationItem_Algorithm.
var (
	AggregationItem_Algorithm_name = map[int32]string{
		0: "COUNT",
		1: "SUM",
		2: "AVG",
		3: "MAX",
		4: "MIN",
	}
	AggregationItem_Algorithm_value = map[string]int32{
		"COUNT": 0,
		"SUM":   1,
		"AVG":   2,
		"MAX":   3,
		"MIN":   4,
	}
)

func (x AggregationItem_Algorithm) Enum() *AggregationItem_Algorithm {
	p := new(AggregationItem_Algorithm)
	*p = x
	return p
}

func (x AggregationItem_Algorithm) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AggregationItem_Algorithm) Descriptor() protoreflect.EnumDescriptor {
	return file_model_aggregation_proto_enumTypes[0].Descriptor()
}

func (AggregationItem_Algorithm) Type() protoreflect.EnumType {
	return &file_model_aggregation_proto_enumTypes[0]
}

func (x AggregationItem_Algorithm) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AggregationItem_Algorithm.Descriptor instead.
func (AggregationItem_Algorithm) EnumDescriptor() ([]byte, []int) {
	return file_model_aggregation_proto_rawDescGZIP(), []int{0, 0}
}

type AggregationItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string                    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Algorithm AggregationItem_Algorithm `protobuf:"varint,2,opt,name=algorithm,proto3,enum=model.AggregationItem_Algorithm" json:"algorithm,omitempty"`
	Property  string                    `protobuf:"bytes,3,opt,name=property,proto3" json:"property,omitempty"`
}

func (x *AggregationItem) Reset() {
	*x = AggregationItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_aggregation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AggregationItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AggregationItem) ProtoMessage() {}

func (x *AggregationItem) ProtoReflect() protoreflect.Message {
	mi := &file_model_aggregation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AggregationItem.ProtoReflect.Descriptor instead.
func (*AggregationItem) Descriptor() ([]byte, []int) {
	return file_model_aggregation_proto_rawDescGZIP(), []int{0}
}

func (x *AggregationItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AggregationItem) GetAlgorithm() AggregationItem_Algorithm {
	if x != nil {
		return x.Algorithm
	}
	return AggregationItem_COUNT
}

func (x *AggregationItem) GetProperty() string {
	if x != nil {
		return x.Property
	}
	return ""
}

type GroupingItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Property string `protobuf:"bytes,1,opt,name=property,proto3" json:"property,omitempty"`
}

func (x *GroupingItem) Reset() {
	*x = GroupingItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_aggregation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupingItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupingItem) ProtoMessage() {}

func (x *GroupingItem) ProtoReflect() protoreflect.Message {
	mi := &file_model_aggregation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupingItem.ProtoReflect.Descriptor instead.
func (*GroupingItem) Descriptor() ([]byte, []int) {
	return file_model_aggregation_proto_rawDescGZIP(), []int{1}
}

func (x *GroupingItem) GetProperty() string {
	if x != nil {
		return x.Property
	}
	return ""
}

type Aggregation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items    []*AggregationItem `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	Grouping []*GroupingItem    `protobuf:"bytes,2,rep,name=grouping,proto3" json:"grouping,omitempty"`
}

func (x *Aggregation) Reset() {
	*x = Aggregation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_aggregation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Aggregation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Aggregation) ProtoMessage() {}

func (x *Aggregation) ProtoReflect() protoreflect.Message {
	mi := &file_model_aggregation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Aggregation.ProtoReflect.Descriptor instead.
func (*Aggregation) Descriptor() ([]byte, []int) {
	return file_model_aggregation_proto_rawDescGZIP(), []int{2}
}

func (x *Aggregation) GetItems() []*AggregationItem {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *Aggregation) GetGrouping() []*GroupingItem {
	if x != nil {
		return x.Grouping
	}
	return nil
}

var File_model_aggregation_proto protoreflect.FileDescriptor

var file_model_aggregation_proto_rawDesc = []byte{
	0x0a, 0x17, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbd,
	0x01, 0x0a, 0x0f, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x74,
	0x65, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3e, 0x0a, 0x09, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69,
	0x74, 0x68, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x74, 0x65,
	0x6d, 0x2e, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x52, 0x09, 0x61, 0x6c, 0x67,
	0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72,
	0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72,
	0x74, 0x79, 0x22, 0x3a, 0x0a, 0x09, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x12,
	0x09, 0x0a, 0x05, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x53, 0x55,
	0x4d, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x56, 0x47, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03,
	0x4d, 0x41, 0x58, 0x10, 0x03, 0x12, 0x07, 0x0a, 0x03, 0x4d, 0x49, 0x4e, 0x10, 0x04, 0x22, 0x2a,
	0x0a, 0x0c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x69, 0x6e, 0x67, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x22, 0x6c, 0x0a, 0x0b, 0x41, 0x67,
	0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x0a, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x74, 0x65, 0x6d,
	0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x2f, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x69, 0x6e, 0x67, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x08,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x69, 0x6e, 0x67, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x72, 0x65, 0x77, 0x2f, 0x61,
	0x70, 0x69, 0x62, 0x72, 0x65, 0x77, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_model_aggregation_proto_rawDescOnce sync.Once
	file_model_aggregation_proto_rawDescData = file_model_aggregation_proto_rawDesc
)

func file_model_aggregation_proto_rawDescGZIP() []byte {
	file_model_aggregation_proto_rawDescOnce.Do(func() {
		file_model_aggregation_proto_rawDescData = protoimpl.X.CompressGZIP(file_model_aggregation_proto_rawDescData)
	})
	return file_model_aggregation_proto_rawDescData
}

var file_model_aggregation_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_model_aggregation_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_model_aggregation_proto_goTypes = []interface{}{
	(AggregationItem_Algorithm)(0), // 0: model.AggregationItem.Algorithm
	(*AggregationItem)(nil),        // 1: model.AggregationItem
	(*GroupingItem)(nil),           // 2: model.GroupingItem
	(*Aggregation)(nil),            // 3: model.Aggregation
}
var file_model_aggregation_proto_depIdxs = []int32{
	0, // 0: model.AggregationItem.algorithm:type_name -> model.AggregationItem.Algorithm
	1, // 1: model.Aggregation.items:type_name -> model.AggregationItem
	2, // 2: model.Aggregation.grouping:type_name -> model.GroupingItem
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_model_aggregation_proto_init() }
func file_model_aggregation_proto_init() {
	if File_model_aggregation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_model_aggregation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AggregationItem); i {
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
		file_model_aggregation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupingItem); i {
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
		file_model_aggregation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Aggregation); i {
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
			RawDescriptor: file_model_aggregation_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_model_aggregation_proto_goTypes,
		DependencyIndexes: file_model_aggregation_proto_depIdxs,
		EnumInfos:         file_model_aggregation_proto_enumTypes,
		MessageInfos:      file_model_aggregation_proto_msgTypes,
	}.Build()
	File_model_aggregation_proto = out.File
	file_model_aggregation_proto_rawDesc = nil
	file_model_aggregation_proto_goTypes = nil
	file_model_aggregation_proto_depIdxs = nil
}
