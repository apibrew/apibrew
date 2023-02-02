// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: stub/extension.proto

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

type ListExtensionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *ListExtensionRequest) Reset() {
	*x = ListExtensionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stub_extension_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListExtensionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListExtensionRequest) ProtoMessage() {}

func (x *ListExtensionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stub_extension_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListExtensionRequest.ProtoReflect.Descriptor instead.
func (*ListExtensionRequest) Descriptor() ([]byte, []int) {
	return file_stub_extension_proto_rawDescGZIP(), []int{0}
}

func (x *ListExtensionRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type ListExtensionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content []*model.RemoteExtension `protobuf:"bytes,1,rep,name=content,proto3" json:"content,omitempty"`
}

func (x *ListExtensionResponse) Reset() {
	*x = ListExtensionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stub_extension_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListExtensionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListExtensionResponse) ProtoMessage() {}

func (x *ListExtensionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stub_extension_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListExtensionResponse.ProtoReflect.Descriptor instead.
func (*ListExtensionResponse) Descriptor() ([]byte, []int) {
	return file_stub_extension_proto_rawDescGZIP(), []int{1}
}

func (x *ListExtensionResponse) GetContent() []*model.RemoteExtension {
	if x != nil {
		return x.Content
	}
	return nil
}

type GetExtensionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Id    string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetExtensionRequest) Reset() {
	*x = GetExtensionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stub_extension_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetExtensionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetExtensionRequest) ProtoMessage() {}

func (x *GetExtensionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stub_extension_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetExtensionRequest.ProtoReflect.Descriptor instead.
func (*GetExtensionRequest) Descriptor() ([]byte, []int) {
	return file_stub_extension_proto_rawDescGZIP(), []int{2}
}

func (x *GetExtensionRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *GetExtensionRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetExtensionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Extension *model.RemoteExtension `protobuf:"bytes,1,opt,name=extension,proto3" json:"extension,omitempty"`
}

func (x *GetExtensionResponse) Reset() {
	*x = GetExtensionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stub_extension_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetExtensionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetExtensionResponse) ProtoMessage() {}

func (x *GetExtensionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stub_extension_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetExtensionResponse.ProtoReflect.Descriptor instead.
func (*GetExtensionResponse) Descriptor() ([]byte, []int) {
	return file_stub_extension_proto_rawDescGZIP(), []int{3}
}

func (x *GetExtensionResponse) GetExtension() *model.RemoteExtension {
	if x != nil {
		return x.Extension
	}
	return nil
}

type CreateExtensionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token      string                   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Extensions []*model.RemoteExtension `protobuf:"bytes,2,rep,name=extensions,proto3" json:"extensions,omitempty"`
}

func (x *CreateExtensionRequest) Reset() {
	*x = CreateExtensionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stub_extension_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateExtensionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateExtensionRequest) ProtoMessage() {}

func (x *CreateExtensionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stub_extension_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateExtensionRequest.ProtoReflect.Descriptor instead.
func (*CreateExtensionRequest) Descriptor() ([]byte, []int) {
	return file_stub_extension_proto_rawDescGZIP(), []int{4}
}

func (x *CreateExtensionRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *CreateExtensionRequest) GetExtensions() []*model.RemoteExtension {
	if x != nil {
		return x.Extensions
	}
	return nil
}

type CreateExtensionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Extensions []*model.RemoteExtension `protobuf:"bytes,1,rep,name=extensions,proto3" json:"extensions,omitempty"`
}

func (x *CreateExtensionResponse) Reset() {
	*x = CreateExtensionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stub_extension_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateExtensionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateExtensionResponse) ProtoMessage() {}

func (x *CreateExtensionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stub_extension_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateExtensionResponse.ProtoReflect.Descriptor instead.
func (*CreateExtensionResponse) Descriptor() ([]byte, []int) {
	return file_stub_extension_proto_rawDescGZIP(), []int{5}
}

func (x *CreateExtensionResponse) GetExtensions() []*model.RemoteExtension {
	if x != nil {
		return x.Extensions
	}
	return nil
}

type UpdateExtensionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token      string                   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Extensions []*model.RemoteExtension `protobuf:"bytes,2,rep,name=extensions,proto3" json:"extensions,omitempty"`
}

func (x *UpdateExtensionRequest) Reset() {
	*x = UpdateExtensionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stub_extension_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateExtensionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateExtensionRequest) ProtoMessage() {}

func (x *UpdateExtensionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stub_extension_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateExtensionRequest.ProtoReflect.Descriptor instead.
func (*UpdateExtensionRequest) Descriptor() ([]byte, []int) {
	return file_stub_extension_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateExtensionRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *UpdateExtensionRequest) GetExtensions() []*model.RemoteExtension {
	if x != nil {
		return x.Extensions
	}
	return nil
}

type UpdateExtensionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Extensions []*model.RemoteExtension `protobuf:"bytes,1,rep,name=extensions,proto3" json:"extensions,omitempty"`
}

func (x *UpdateExtensionResponse) Reset() {
	*x = UpdateExtensionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stub_extension_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateExtensionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateExtensionResponse) ProtoMessage() {}

func (x *UpdateExtensionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stub_extension_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateExtensionResponse.ProtoReflect.Descriptor instead.
func (*UpdateExtensionResponse) Descriptor() ([]byte, []int) {
	return file_stub_extension_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateExtensionResponse) GetExtensions() []*model.RemoteExtension {
	if x != nil {
		return x.Extensions
	}
	return nil
}

type DeleteExtensionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Ids   []string `protobuf:"bytes,2,rep,name=ids,proto3" json:"ids,omitempty"`
}

func (x *DeleteExtensionRequest) Reset() {
	*x = DeleteExtensionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stub_extension_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteExtensionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteExtensionRequest) ProtoMessage() {}

func (x *DeleteExtensionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stub_extension_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteExtensionRequest.ProtoReflect.Descriptor instead.
func (*DeleteExtensionRequest) Descriptor() ([]byte, []int) {
	return file_stub_extension_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteExtensionRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *DeleteExtensionRequest) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

type DeleteExtensionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteExtensionResponse) Reset() {
	*x = DeleteExtensionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stub_extension_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteExtensionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteExtensionResponse) ProtoMessage() {}

func (x *DeleteExtensionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stub_extension_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteExtensionResponse.ProtoReflect.Descriptor instead.
func (*DeleteExtensionResponse) Descriptor() ([]byte, []int) {
	return file_stub_extension_proto_rawDescGZIP(), []int{9}
}

var File_stub_extension_proto protoreflect.FileDescriptor

var file_stub_extension_proto_rawDesc = []byte{
	0x0a, 0x14, 0x73, 0x74, 0x75, 0x62, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x73, 0x74, 0x75, 0x62, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2c, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x22, 0x49, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x45, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x3b, 0x0a,
	0x13, 0x47, 0x65, 0x74, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4c, 0x0a, 0x14, 0x47, 0x65,
	0x74, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x34, 0x0a, 0x09, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x65,
	0x6d, 0x6f, 0x74, 0x65, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x66, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x36, 0x0a, 0x0a, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x45, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x22, 0x51, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x0a, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x45, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x22, 0x66, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x36, 0x0a, 0x0a, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x0a, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x51, 0x0a, 0x17, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x0a, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x0a, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x40,
	0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x10,
	0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x69, 0x64, 0x73,
	0x22, 0x19, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xf0, 0x02, 0x0a, 0x10,
	0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x41, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1a, 0x2e, 0x73, 0x74, 0x75, 0x62, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x73, 0x74, 0x75, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x19, 0x2e, 0x73, 0x74, 0x75,
	0x62, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x73, 0x74, 0x75, 0x62, 0x2e, 0x47, 0x65, 0x74,
	0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x2e,
	0x73, 0x74, 0x75, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x73, 0x74,
	0x75, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x06,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x2e, 0x73, 0x74, 0x75, 0x62, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x73, 0x74, 0x75, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12,
	0x1c, 0x2e, 0x73, 0x74, 0x75, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e,
	0x73, 0x74, 0x75, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x29,
	0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x69, 0x73,
	0x6c, 0x69, 0x62, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2d, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x74, 0x75, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_stub_extension_proto_rawDescOnce sync.Once
	file_stub_extension_proto_rawDescData = file_stub_extension_proto_rawDesc
)

func file_stub_extension_proto_rawDescGZIP() []byte {
	file_stub_extension_proto_rawDescOnce.Do(func() {
		file_stub_extension_proto_rawDescData = protoimpl.X.CompressGZIP(file_stub_extension_proto_rawDescData)
	})
	return file_stub_extension_proto_rawDescData
}

var file_stub_extension_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_stub_extension_proto_goTypes = []interface{}{
	(*ListExtensionRequest)(nil),    // 0: stub.ListExtensionRequest
	(*ListExtensionResponse)(nil),   // 1: stub.ListExtensionResponse
	(*GetExtensionRequest)(nil),     // 2: stub.GetExtensionRequest
	(*GetExtensionResponse)(nil),    // 3: stub.GetExtensionResponse
	(*CreateExtensionRequest)(nil),  // 4: stub.CreateExtensionRequest
	(*CreateExtensionResponse)(nil), // 5: stub.CreateExtensionResponse
	(*UpdateExtensionRequest)(nil),  // 6: stub.UpdateExtensionRequest
	(*UpdateExtensionResponse)(nil), // 7: stub.UpdateExtensionResponse
	(*DeleteExtensionRequest)(nil),  // 8: stub.DeleteExtensionRequest
	(*DeleteExtensionResponse)(nil), // 9: stub.DeleteExtensionResponse
	(*model.RemoteExtension)(nil),   // 10: model.RemoteExtension
}
var file_stub_extension_proto_depIdxs = []int32{
	10, // 0: stub.ListExtensionResponse.content:type_name -> model.RemoteExtension
	10, // 1: stub.GetExtensionResponse.extension:type_name -> model.RemoteExtension
	10, // 2: stub.CreateExtensionRequest.extensions:type_name -> model.RemoteExtension
	10, // 3: stub.CreateExtensionResponse.extensions:type_name -> model.RemoteExtension
	10, // 4: stub.UpdateExtensionRequest.extensions:type_name -> model.RemoteExtension
	10, // 5: stub.UpdateExtensionResponse.extensions:type_name -> model.RemoteExtension
	0,  // 6: stub.ExtensionService.List:input_type -> stub.ListExtensionRequest
	2,  // 7: stub.ExtensionService.Get:input_type -> stub.GetExtensionRequest
	4,  // 8: stub.ExtensionService.Create:input_type -> stub.CreateExtensionRequest
	6,  // 9: stub.ExtensionService.Update:input_type -> stub.UpdateExtensionRequest
	8,  // 10: stub.ExtensionService.Delete:input_type -> stub.DeleteExtensionRequest
	1,  // 11: stub.ExtensionService.List:output_type -> stub.ListExtensionResponse
	3,  // 12: stub.ExtensionService.Get:output_type -> stub.GetExtensionResponse
	5,  // 13: stub.ExtensionService.Create:output_type -> stub.CreateExtensionResponse
	7,  // 14: stub.ExtensionService.Update:output_type -> stub.UpdateExtensionResponse
	9,  // 15: stub.ExtensionService.Delete:output_type -> stub.DeleteExtensionResponse
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_stub_extension_proto_init() }
func file_stub_extension_proto_init() {
	if File_stub_extension_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_stub_extension_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListExtensionRequest); i {
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
		file_stub_extension_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListExtensionResponse); i {
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
		file_stub_extension_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetExtensionRequest); i {
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
		file_stub_extension_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetExtensionResponse); i {
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
		file_stub_extension_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateExtensionRequest); i {
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
		file_stub_extension_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateExtensionResponse); i {
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
		file_stub_extension_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateExtensionRequest); i {
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
		file_stub_extension_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateExtensionResponse); i {
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
		file_stub_extension_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteExtensionRequest); i {
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
		file_stub_extension_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteExtensionResponse); i {
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
			RawDescriptor: file_stub_extension_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_stub_extension_proto_goTypes,
		DependencyIndexes: file_stub_extension_proto_depIdxs,
		MessageInfos:      file_stub_extension_proto_msgTypes,
	}.Build()
	File_stub_extension_proto = out.File
	file_stub_extension_proto_rawDesc = nil
	file_stub_extension_proto_goTypes = nil
	file_stub_extension_proto_depIdxs = nil
}