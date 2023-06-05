// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: model/role.proto

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

type Role struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name            string           `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	SecurityContext *SecurityContext `protobuf:"bytes,5,opt,name=securityContext,proto3" json:"securityContext,omitempty"`
	Details         *structpb.Struct `protobuf:"bytes,6,opt,name=details,proto3,oneof" json:"details,omitempty"`
	AuditData       *AuditData       `protobuf:"bytes,101,opt,name=auditData,proto3" json:"auditData,omitempty"`
	Version         uint32           `protobuf:"varint,102,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *Role) Reset() {
	*x = Role{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_role_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Role) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Role) ProtoMessage() {}

func (x *Role) ProtoReflect() protoreflect.Message {
	mi := &file_model_role_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Role.ProtoReflect.Descriptor instead.
func (*Role) Descriptor() ([]byte, []int) {
	return file_model_role_proto_rawDescGZIP(), []int{0}
}

func (x *Role) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Role) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Role) GetSecurityContext() *SecurityContext {
	if x != nil {
		return x.SecurityContext
	}
	return nil
}

func (x *Role) GetDetails() *structpb.Struct {
	if x != nil {
		return x.Details
	}
	return nil
}

func (x *Role) GetAuditData() *AuditData {
	if x != nil {
		return x.AuditData
	}
	return nil
}

func (x *Role) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

var File_model_role_proto protoreflect.FileDescriptor

var file_model_role_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x61,
	0x75, 0x64, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x68, 0x63, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc0,
	0x02, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x06, 0xd0, 0x47, 0x01, 0xc8, 0x7d, 0x08, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xa0,
	0x7d, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x54, 0x0a, 0x0f, 0x73, 0x65, 0x63, 0x75,
	0x72, 0x69, 0x74, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69,
	0x74, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x42, 0x12, 0xc2, 0x47, 0x0f, 0x73, 0x65,
	0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x0f, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x36,
	0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x48, 0x00, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x88, 0x01, 0x01, 0x12, 0x2e, 0x0a, 0x09, 0x61, 0x75, 0x64, 0x69, 0x74, 0x44,
	0x61, 0x74, 0x61, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x09, 0x61, 0x75, 0x64,
	0x69, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x66, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x3a, 0x23, 0xda, 0x3e, 0x06, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0xe2, 0x3e, 0x06, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0xf2, 0x3e, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0xa2, 0x3f, 0x04, 0x72, 0x6f,
	0x6c, 0x65, 0xb0, 0x3f, 0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x61, 0x70, 0x69, 0x62, 0x72, 0x65, 0x77, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x72, 0x65, 0x77, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_model_role_proto_rawDescOnce sync.Once
	file_model_role_proto_rawDescData = file_model_role_proto_rawDesc
)

func file_model_role_proto_rawDescGZIP() []byte {
	file_model_role_proto_rawDescOnce.Do(func() {
		file_model_role_proto_rawDescData = protoimpl.X.CompressGZIP(file_model_role_proto_rawDescData)
	})
	return file_model_role_proto_rawDescData
}

var file_model_role_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_model_role_proto_goTypes = []interface{}{
	(*Role)(nil),            // 0: model.Role
	(*SecurityContext)(nil), // 1: model.SecurityContext
	(*structpb.Struct)(nil), // 2: google.protobuf.Struct
	(*AuditData)(nil),       // 3: model.AuditData
}
var file_model_role_proto_depIdxs = []int32{
	1, // 0: model.Role.securityContext:type_name -> model.SecurityContext
	2, // 1: model.Role.details:type_name -> google.protobuf.Struct
	3, // 2: model.Role.auditData:type_name -> model.AuditData
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_model_role_proto_init() }
func file_model_role_proto_init() {
	if File_model_role_proto != nil {
		return
	}
	file_model_audit_proto_init()
	file_model_common_proto_init()
	file_model_security_proto_init()
	file_model_hcl_proto_init()
	file_model_annotations_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_model_role_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Role); i {
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
	file_model_role_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_model_role_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_model_role_proto_goTypes,
		DependencyIndexes: file_model_role_proto_depIdxs,
		MessageInfos:      file_model_role_proto_msgTypes,
	}.Build()
	File_model_role_proto = out.File
	file_model_role_proto_rawDesc = nil
	file_model_role_proto_goTypes = nil
	file_model_role_proto_depIdxs = nil
}
