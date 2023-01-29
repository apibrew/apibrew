// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: model/token.proto

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

type TokenTerm int32

const (
	TokenTerm_SHORT     TokenTerm = 0 // 1 minute
	TokenTerm_MIDDLE    TokenTerm = 1 // 2 hours
	TokenTerm_LONG      TokenTerm = 2 // 2 days
	TokenTerm_VERY_LONG TokenTerm = 3 // 2 years
)

// Enum value maps for TokenTerm.
var (
	TokenTerm_name = map[int32]string{
		0: "SHORT",
		1: "MIDDLE",
		2: "LONG",
		3: "VERY_LONG",
	}
	TokenTerm_value = map[string]int32{
		"SHORT":     0,
		"MIDDLE":    1,
		"LONG":      2,
		"VERY_LONG": 3,
	}
)

func (x TokenTerm) Enum() *TokenTerm {
	p := new(TokenTerm)
	*p = x
	return p
}

func (x TokenTerm) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TokenTerm) Descriptor() protoreflect.EnumDescriptor {
	return file_model_token_proto_enumTypes[0].Descriptor()
}

func (TokenTerm) Type() protoreflect.EnumType {
	return &file_model_token_proto_enumTypes[0]
}

func (x TokenTerm) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TokenTerm.Descriptor instead.
func (TokenTerm) EnumDescriptor() ([]byte, []int) {
	return file_model_token_proto_rawDescGZIP(), []int{0}
}

type Token struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Term       TokenTerm              `protobuf:"varint,1,opt,name=term,proto3,enum=model.TokenTerm" json:"term,omitempty"` // issue term
	Content    string                 `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`                 // jwt token
	Expiration *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=expiration,proto3" json:"expiration,omitempty"`           // expiration time
}

func (x *Token) Reset() {
	*x = Token{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_token_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Token) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Token) ProtoMessage() {}

func (x *Token) ProtoReflect() protoreflect.Message {
	mi := &file_model_token_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Token.ProtoReflect.Descriptor instead.
func (*Token) Descriptor() ([]byte, []int) {
	return file_model_token_proto_rawDescGZIP(), []int{0}
}

func (x *Token) GetTerm() TokenTerm {
	if x != nil {
		return x.Term
	}
	return TokenTerm_SHORT
}

func (x *Token) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Token) GetExpiration() *timestamppb.Timestamp {
	if x != nil {
		return x.Expiration
	}
	return nil
}

var File_model_token_proto protoreflect.FileDescriptor

var file_model_token_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x83, 0x01, 0x0a, 0x05,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x24, 0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x54, 0x65, 0x72, 0x6d, 0x52, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x3a, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2a, 0x3b, 0x0a, 0x09, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x65, 0x72, 0x6d, 0x12, 0x09,
	0x0a, 0x05, 0x53, 0x48, 0x4f, 0x52, 0x54, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x49, 0x44,
	0x44, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x4c, 0x4f, 0x4e, 0x47, 0x10, 0x02, 0x12,
	0x0d, 0x0a, 0x09, 0x56, 0x45, 0x52, 0x59, 0x5f, 0x4c, 0x4f, 0x4e, 0x47, 0x10, 0x03, 0x42, 0x2a,
	0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x69, 0x73,
	0x6c, 0x69, 0x62, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2d, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_model_token_proto_rawDescOnce sync.Once
	file_model_token_proto_rawDescData = file_model_token_proto_rawDesc
)

func file_model_token_proto_rawDescGZIP() []byte {
	file_model_token_proto_rawDescOnce.Do(func() {
		file_model_token_proto_rawDescData = protoimpl.X.CompressGZIP(file_model_token_proto_rawDescData)
	})
	return file_model_token_proto_rawDescData
}

var file_model_token_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_model_token_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_model_token_proto_goTypes = []interface{}{
	(TokenTerm)(0),                // 0: model.TokenTerm
	(*Token)(nil),                 // 1: model.Token
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_model_token_proto_depIdxs = []int32{
	0, // 0: model.Token.term:type_name -> model.TokenTerm
	2, // 1: model.Token.expiration:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_model_token_proto_init() }
func file_model_token_proto_init() {
	if File_model_token_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_model_token_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Token); i {
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
			RawDescriptor: file_model_token_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_model_token_proto_goTypes,
		DependencyIndexes: file_model_token_proto_depIdxs,
		EnumInfos:         file_model_token_proto_enumTypes,
		MessageInfos:      file_model_token_proto_msgTypes,
	}.Build()
	File_model_token_proto = out.File
	file_model_token_proto_rawDesc = nil
	file_model_token_proto_goTypes = nil
	file_model_token_proto_depIdxs = nil
}
