// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: model/data-source.proto

package model

import (
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

type DataSourceEntity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ReadOnly bool   `protobuf:"varint,2,opt,name=readOnly,proto3" json:"readOnly,omitempty"`
}

func (x *DataSourceEntity) Reset() {
	*x = DataSourceEntity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_data_source_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataSourceEntity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataSourceEntity) ProtoMessage() {}

func (x *DataSourceEntity) ProtoReflect() protoreflect.Message {
	mi := &file_model_data_source_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataSourceEntity.ProtoReflect.Descriptor instead.
func (*DataSourceEntity) Descriptor() ([]byte, []int) {
	return file_model_data_source_proto_rawDescGZIP(), []int{0}
}

func (x *DataSourceEntity) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DataSourceEntity) GetReadOnly() bool {
	if x != nil {
		return x.ReadOnly
	}
	return false
}

type DataSourceCatalog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string              `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Entities []*DataSourceEntity `protobuf:"bytes,2,rep,name=entities,proto3" json:"entities,omitempty"`
}

func (x *DataSourceCatalog) Reset() {
	*x = DataSourceCatalog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_data_source_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataSourceCatalog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataSourceCatalog) ProtoMessage() {}

func (x *DataSourceCatalog) ProtoReflect() protoreflect.Message {
	mi := &file_model_data_source_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataSourceCatalog.ProtoReflect.Descriptor instead.
func (*DataSourceCatalog) Descriptor() ([]byte, []int) {
	return file_model_data_source_proto_rawDescGZIP(), []int{1}
}

func (x *DataSourceCatalog) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DataSourceCatalog) GetEntities() []*DataSourceEntity {
	if x != nil {
		return x.Entities
	}
	return nil
}

var File_model_data_source_proto protoreflect.FileDescriptor

var file_model_data_source_proto_rawDesc = []byte{
	0x0a, 0x17, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2d, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x22, 0x42, 0x0a, 0x10, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x61, 0x64,
	0x4f, 0x6e, 0x6c, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x72, 0x65, 0x61, 0x64,
	0x4f, 0x6e, 0x6c, 0x79, 0x22, 0x5c, 0x0a, 0x11, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x33, 0x0a,
	0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x61, 0x70, 0x69, 0x62, 0x72, 0x65, 0x77, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x72, 0x65, 0x77,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_model_data_source_proto_rawDescOnce sync.Once
	file_model_data_source_proto_rawDescData = file_model_data_source_proto_rawDesc
)

func file_model_data_source_proto_rawDescGZIP() []byte {
	file_model_data_source_proto_rawDescOnce.Do(func() {
		file_model_data_source_proto_rawDescData = protoimpl.X.CompressGZIP(file_model_data_source_proto_rawDescData)
	})
	return file_model_data_source_proto_rawDescData
}

var file_model_data_source_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_model_data_source_proto_goTypes = []interface{}{
	(*DataSourceEntity)(nil),  // 0: model.DataSourceEntity
	(*DataSourceCatalog)(nil), // 1: model.DataSourceCatalog
}
var file_model_data_source_proto_depIdxs = []int32{
	0, // 0: model.DataSourceCatalog.entities:type_name -> model.DataSourceEntity
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_model_data_source_proto_init() }
func file_model_data_source_proto_init() {
	if File_model_data_source_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_model_data_source_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataSourceEntity); i {
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
		file_model_data_source_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataSourceCatalog); i {
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
			RawDescriptor: file_model_data_source_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_model_data_source_proto_goTypes,
		DependencyIndexes: file_model_data_source_proto_depIdxs,
		MessageInfos:      file_model_data_source_proto_msgTypes,
	}.Build()
	File_model_data_source_proto = out.File
	file_model_data_source_proto_rawDesc = nil
	file_model_data_source_proto_goTypes = nil
	file_model_data_source_proto_depIdxs = nil
}
