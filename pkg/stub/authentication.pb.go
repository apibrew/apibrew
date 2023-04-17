// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: stub/authentication.proto

package stub

import (
	model "github.com/tislib/data-handler/pkg/model"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// AuthenticationRequest is the request for authentication
type AuthenticationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	// Token Lifespan
	Term model.TokenTerm `protobuf:"varint,3,opt,name=term,proto3,enum=model.TokenTerm" json:"term,omitempty"`
}

func (x *AuthenticationRequest) Reset() {
	*x = AuthenticationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stub_authentication_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticationRequest) ProtoMessage() {}

func (x *AuthenticationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stub_authentication_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticationRequest.ProtoReflect.Descriptor instead.
func (*AuthenticationRequest) Descriptor() ([]byte, []int) {
	return file_stub_authentication_proto_rawDescGZIP(), []int{0}
}

func (x *AuthenticationRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *AuthenticationRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *AuthenticationRequest) GetTerm() model.TokenTerm {
	if x != nil {
		return x.Term
	}
	return model.TokenTerm(0)
}

type AuthenticationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token *model.Token `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *AuthenticationResponse) Reset() {
	*x = AuthenticationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stub_authentication_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticationResponse) ProtoMessage() {}

func (x *AuthenticationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stub_authentication_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticationResponse.ProtoReflect.Descriptor instead.
func (*AuthenticationResponse) Descriptor() ([]byte, []int) {
	return file_stub_authentication_proto_rawDescGZIP(), []int{1}
}

func (x *AuthenticationResponse) GetToken() *model.Token {
	if x != nil {
		return x.Token
	}
	return nil
}

type RenewTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string          `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Term  model.TokenTerm `protobuf:"varint,2,opt,name=term,proto3,enum=model.TokenTerm" json:"term,omitempty"`
}

func (x *RenewTokenRequest) Reset() {
	*x = RenewTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stub_authentication_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RenewTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RenewTokenRequest) ProtoMessage() {}

func (x *RenewTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stub_authentication_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RenewTokenRequest.ProtoReflect.Descriptor instead.
func (*RenewTokenRequest) Descriptor() ([]byte, []int) {
	return file_stub_authentication_proto_rawDescGZIP(), []int{2}
}

func (x *RenewTokenRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *RenewTokenRequest) GetTerm() model.TokenTerm {
	if x != nil {
		return x.Term
	}
	return model.TokenTerm(0)
}

type RenewTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token *model.Token `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *RenewTokenResponse) Reset() {
	*x = RenewTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stub_authentication_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RenewTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RenewTokenResponse) ProtoMessage() {}

func (x *RenewTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stub_authentication_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RenewTokenResponse.ProtoReflect.Descriptor instead.
func (*RenewTokenResponse) Descriptor() ([]byte, []int) {
	return file_stub_authentication_proto_rawDescGZIP(), []int{3}
}

func (x *RenewTokenResponse) GetToken() *model.Token {
	if x != nil {
		return x.Token
	}
	return nil
}

var File_stub_authentication_proto protoreflect.FileDescriptor

var file_stub_authentication_proto_rawDesc = []byte{
	0x0a, 0x19, 0x73, 0x74, 0x75, 0x62, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x73, 0x74, 0x75,
	0x62, 0x1a, 0x11, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x2f, 0x6f,
	0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x75, 0x0a, 0x15, 0x41,
	0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x24, 0x0a, 0x04,
	0x74, 0x65, 0x72, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x65, 0x72, 0x6d, 0x52, 0x04, 0x74, 0x65,
	0x72, 0x6d, 0x22, 0x3c, 0x0a, 0x16, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0x4f, 0x0a, 0x11, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x24, 0x0a, 0x04, 0x74,
	0x65, 0x72, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x65, 0x72, 0x6d, 0x52, 0x04, 0x74, 0x65, 0x72,
	0x6d, 0x22, 0x38, 0x0a, 0x12, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0xd1, 0x03, 0x0a, 0x0e,
	0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0xff,
	0x01, 0x0a, 0x0c, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12,
	0x1b, 0x2e, 0x73, 0x74, 0x75, 0x62, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73,
	0x74, 0x75, 0x62, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xb3, 0x01, 0xba, 0x47, 0x8f,
	0x01, 0x0a, 0x0e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x25, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x20, 0x77, 0x69, 0x74, 0x68, 0x20, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x2f,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x1a, 0x3a, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e,
	0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x77, 0x69, 0x74, 0x68, 0x20, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x2f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x20,
	0x61, 0x6e, 0x64, 0x20, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x6e, 0x65, 0x77, 0x20, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x2a, 0x1a, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x3a, 0x01, 0x2a, 0x22, 0x15, 0x2f, 0x61, 0x75, 0x74, 0x68,
	0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0xbc, 0x01, 0x0a, 0x0a, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x17, 0x2e, 0x73, 0x74, 0x75, 0x62, 0x2e, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x73, 0x74, 0x75, 0x62, 0x2e,
	0x52, 0x65, 0x6e, 0x65, 0x77, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x7b, 0xba, 0x47, 0x58, 0x0a, 0x0e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0b, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x20, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x1f, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x20, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x20, 0x77, 0x69, 0x74, 0x68, 0x20, 0x65, 0x78, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x20,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2a, 0x18, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x3a, 0x01, 0x2a, 0x1a, 0x15, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x65,
	0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x42,
	0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x69,
	0x73, 0x6c, 0x69, 0x62, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2d, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65,
	0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x74, 0x75, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_stub_authentication_proto_rawDescOnce sync.Once
	file_stub_authentication_proto_rawDescData = file_stub_authentication_proto_rawDesc
)

func file_stub_authentication_proto_rawDescGZIP() []byte {
	file_stub_authentication_proto_rawDescOnce.Do(func() {
		file_stub_authentication_proto_rawDescData = protoimpl.X.CompressGZIP(file_stub_authentication_proto_rawDescData)
	})
	return file_stub_authentication_proto_rawDescData
}

var file_stub_authentication_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_stub_authentication_proto_goTypes = []interface{}{
	(*AuthenticationRequest)(nil),  // 0: stub.AuthenticationRequest
	(*AuthenticationResponse)(nil), // 1: stub.AuthenticationResponse
	(*RenewTokenRequest)(nil),      // 2: stub.RenewTokenRequest
	(*RenewTokenResponse)(nil),     // 3: stub.RenewTokenResponse
	(model.TokenTerm)(0),           // 4: model.TokenTerm
	(*model.Token)(nil),            // 5: model.Token
}
var file_stub_authentication_proto_depIdxs = []int32{
	4, // 0: stub.AuthenticationRequest.term:type_name -> model.TokenTerm
	5, // 1: stub.AuthenticationResponse.token:type_name -> model.Token
	4, // 2: stub.RenewTokenRequest.term:type_name -> model.TokenTerm
	5, // 3: stub.RenewTokenResponse.token:type_name -> model.Token
	0, // 4: stub.Authentication.Authenticate:input_type -> stub.AuthenticationRequest
	2, // 5: stub.Authentication.RenewToken:input_type -> stub.RenewTokenRequest
	1, // 6: stub.Authentication.Authenticate:output_type -> stub.AuthenticationResponse
	3, // 7: stub.Authentication.RenewToken:output_type -> stub.RenewTokenResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_stub_authentication_proto_init() }
func file_stub_authentication_proto_init() {
	if File_stub_authentication_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_stub_authentication_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticationRequest); i {
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
		file_stub_authentication_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticationResponse); i {
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
		file_stub_authentication_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RenewTokenRequest); i {
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
		file_stub_authentication_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RenewTokenResponse); i {
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
			RawDescriptor: file_stub_authentication_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_stub_authentication_proto_goTypes,
		DependencyIndexes: file_stub_authentication_proto_depIdxs,
		MessageInfos:      file_stub_authentication_proto_msgTypes,
	}.Build()
	File_stub_authentication_proto = out.File
	file_stub_authentication_proto_rawDesc = nil
	file_stub_authentication_proto_goTypes = nil
	file_stub_authentication_proto_depIdxs = nil
}
