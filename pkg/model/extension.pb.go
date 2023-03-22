// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: model/extension.proto

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

type Extension struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string             `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string             `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string             `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Namespace   string             `protobuf:"bytes,5,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Resource    string             `protobuf:"bytes,6,opt,name=resource,proto3" json:"resource,omitempty"`
	Before      *Extension_Before  `protobuf:"bytes,7,opt,name=before,proto3" json:"before,omitempty"`
	Instead     *Extension_Instead `protobuf:"bytes,8,opt,name=instead,proto3" json:"instead,omitempty"`
	After       *Extension_After   `protobuf:"bytes,9,opt,name=after,proto3" json:"after,omitempty"`
	AuditData   *AuditData         `protobuf:"bytes,101,opt,name=auditData,proto3" json:"auditData,omitempty"`
	Version     uint32             `protobuf:"varint,102,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *Extension) Reset() {
	*x = Extension{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_extension_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Extension) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Extension) ProtoMessage() {}

func (x *Extension) ProtoReflect() protoreflect.Message {
	mi := &file_model_extension_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Extension.ProtoReflect.Descriptor instead.
func (*Extension) Descriptor() ([]byte, []int) {
	return file_model_extension_proto_rawDescGZIP(), []int{0}
}

func (x *Extension) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Extension) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Extension) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Extension) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *Extension) GetResource() string {
	if x != nil {
		return x.Resource
	}
	return ""
}

func (x *Extension) GetBefore() *Extension_Before {
	if x != nil {
		return x.Before
	}
	return nil
}

func (x *Extension) GetInstead() *Extension_Instead {
	if x != nil {
		return x.Instead
	}
	return nil
}

func (x *Extension) GetAfter() *Extension_After {
	if x != nil {
		return x.After
	}
	return nil
}

func (x *Extension) GetAuditData() *AuditData {
	if x != nil {
		return x.AuditData
	}
	return nil
}

func (x *Extension) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

type Extension_After struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	All    *ExternalCall `protobuf:"bytes,1,opt,name=all,proto3" json:"all,omitempty"`
	Create *ExternalCall `protobuf:"bytes,2,opt,name=create,proto3" json:"create,omitempty"`
	Update *ExternalCall `protobuf:"bytes,3,opt,name=update,proto3" json:"update,omitempty"`
	Delete *ExternalCall `protobuf:"bytes,4,opt,name=delete,proto3" json:"delete,omitempty"`
	Get    *ExternalCall `protobuf:"bytes,5,opt,name=get,proto3" json:"get,omitempty"`
	List   *ExternalCall `protobuf:"bytes,6,opt,name=list,proto3" json:"list,omitempty"`
	Sync   bool          `protobuf:"varint,7,opt,name=sync,proto3" json:"sync,omitempty"` // if true, it will wait
}

func (x *Extension_After) Reset() {
	*x = Extension_After{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_extension_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Extension_After) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Extension_After) ProtoMessage() {}

func (x *Extension_After) ProtoReflect() protoreflect.Message {
	mi := &file_model_extension_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Extension_After.ProtoReflect.Descriptor instead.
func (*Extension_After) Descriptor() ([]byte, []int) {
	return file_model_extension_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Extension_After) GetAll() *ExternalCall {
	if x != nil {
		return x.All
	}
	return nil
}

func (x *Extension_After) GetCreate() *ExternalCall {
	if x != nil {
		return x.Create
	}
	return nil
}

func (x *Extension_After) GetUpdate() *ExternalCall {
	if x != nil {
		return x.Update
	}
	return nil
}

func (x *Extension_After) GetDelete() *ExternalCall {
	if x != nil {
		return x.Delete
	}
	return nil
}

func (x *Extension_After) GetGet() *ExternalCall {
	if x != nil {
		return x.Get
	}
	return nil
}

func (x *Extension_After) GetList() *ExternalCall {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *Extension_After) GetSync() bool {
	if x != nil {
		return x.Sync
	}
	return false
}

type Extension_Before struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	All    *ExternalCall `protobuf:"bytes,1,opt,name=all,proto3" json:"all,omitempty"`
	Create *ExternalCall `protobuf:"bytes,2,opt,name=create,proto3" json:"create,omitempty"`
	Update *ExternalCall `protobuf:"bytes,3,opt,name=update,proto3" json:"update,omitempty"`
	Delete *ExternalCall `protobuf:"bytes,4,opt,name=delete,proto3" json:"delete,omitempty"`
	Get    *ExternalCall `protobuf:"bytes,5,opt,name=get,proto3" json:"get,omitempty"`
	List   *ExternalCall `protobuf:"bytes,6,opt,name=list,proto3" json:"list,omitempty"`
	Sync   bool          `protobuf:"varint,7,opt,name=sync,proto3" json:"sync,omitempty"` // if true, it will wait
}

func (x *Extension_Before) Reset() {
	*x = Extension_Before{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_extension_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Extension_Before) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Extension_Before) ProtoMessage() {}

func (x *Extension_Before) ProtoReflect() protoreflect.Message {
	mi := &file_model_extension_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Extension_Before.ProtoReflect.Descriptor instead.
func (*Extension_Before) Descriptor() ([]byte, []int) {
	return file_model_extension_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Extension_Before) GetAll() *ExternalCall {
	if x != nil {
		return x.All
	}
	return nil
}

func (x *Extension_Before) GetCreate() *ExternalCall {
	if x != nil {
		return x.Create
	}
	return nil
}

func (x *Extension_Before) GetUpdate() *ExternalCall {
	if x != nil {
		return x.Update
	}
	return nil
}

func (x *Extension_Before) GetDelete() *ExternalCall {
	if x != nil {
		return x.Delete
	}
	return nil
}

func (x *Extension_Before) GetGet() *ExternalCall {
	if x != nil {
		return x.Get
	}
	return nil
}

func (x *Extension_Before) GetList() *ExternalCall {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *Extension_Before) GetSync() bool {
	if x != nil {
		return x.Sync
	}
	return false
}

type Extension_Instead struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	All    *ExternalCall `protobuf:"bytes,1,opt,name=all,proto3" json:"all,omitempty"`
	Create *ExternalCall `protobuf:"bytes,2,opt,name=create,proto3" json:"create,omitempty"`
	Update *ExternalCall `protobuf:"bytes,3,opt,name=update,proto3" json:"update,omitempty"`
	Delete *ExternalCall `protobuf:"bytes,4,opt,name=delete,proto3" json:"delete,omitempty"`
	Get    *ExternalCall `protobuf:"bytes,5,opt,name=get,proto3" json:"get,omitempty"`
	List   *ExternalCall `protobuf:"bytes,6,opt,name=list,proto3" json:"list,omitempty"`
}

func (x *Extension_Instead) Reset() {
	*x = Extension_Instead{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_extension_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Extension_Instead) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Extension_Instead) ProtoMessage() {}

func (x *Extension_Instead) ProtoReflect() protoreflect.Message {
	mi := &file_model_extension_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Extension_Instead.ProtoReflect.Descriptor instead.
func (*Extension_Instead) Descriptor() ([]byte, []int) {
	return file_model_extension_proto_rawDescGZIP(), []int{0, 2}
}

func (x *Extension_Instead) GetAll() *ExternalCall {
	if x != nil {
		return x.All
	}
	return nil
}

func (x *Extension_Instead) GetCreate() *ExternalCall {
	if x != nil {
		return x.Create
	}
	return nil
}

func (x *Extension_Instead) GetUpdate() *ExternalCall {
	if x != nil {
		return x.Update
	}
	return nil
}

func (x *Extension_Instead) GetDelete() *ExternalCall {
	if x != nil {
		return x.Delete
	}
	return nil
}

func (x *Extension_Instead) GetGet() *ExternalCall {
	if x != nil {
		return x.Get
	}
	return nil
}

func (x *Extension_Instead) GetList() *ExternalCall {
	if x != nil {
		return x.List
	}
	return nil
}

var File_model_extension_proto protoreflect.FileDescriptor

var file_model_extension_proto_rawDesc = []byte{
	0x0a, 0x15, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x11,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x12, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x65, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xab, 0x09, 0x0a, 0x09,
	0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x2f, 0x0a, 0x06, 0x62, 0x65, 0x66,
	0x6f, 0x72, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x42, 0x65, 0x66, 0x6f,
	0x72, 0x65, 0x52, 0x06, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x12, 0x32, 0x0a, 0x07, 0x69, 0x6e,
	0x73, 0x74, 0x65, 0x61, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x49, 0x6e,
	0x73, 0x74, 0x65, 0x61, 0x64, 0x52, 0x07, 0x69, 0x6e, 0x73, 0x74, 0x65, 0x61, 0x64, 0x12, 0x2c,
	0x0a, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2e,
	0x41, 0x66, 0x74, 0x65, 0x72, 0x52, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x09,
	0x61, 0x75, 0x64, 0x69, 0x74, 0x44, 0x61, 0x74, 0x61, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x74, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x09, 0x61, 0x75, 0x64, 0x69, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x66, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x1a, 0x99, 0x02, 0x0a, 0x05, 0x41, 0x66, 0x74, 0x65, 0x72,
	0x12, 0x25, 0x0a, 0x03, 0x61, 0x6c, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x43, 0x61,
	0x6c, 0x6c, 0x52, 0x03, 0x61, 0x6c, 0x6c, 0x12, 0x2b, 0x0a, 0x06, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x06, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x12, 0x2b, 0x0a, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x2b, 0x0a, 0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x25,
	0x0a, 0x03, 0x67, 0x65, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x43, 0x61, 0x6c, 0x6c,
	0x52, 0x03, 0x67, 0x65, 0x74, 0x12, 0x27, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x73, 0x79, 0x6e, 0x63, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x73, 0x79,
	0x6e, 0x63, 0x1a, 0x9a, 0x02, 0x0a, 0x06, 0x42, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x12, 0x25, 0x0a,
	0x03, 0x61, 0x6c, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x43, 0x61, 0x6c, 0x6c, 0x52,
	0x03, 0x61, 0x6c, 0x6c, 0x12, 0x2b, 0x0a, 0x06, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x06, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x12, 0x2b, 0x0a, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x2b,
	0x0a, 0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x43,
	0x61, 0x6c, 0x6c, 0x52, 0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x25, 0x0a, 0x03, 0x67,
	0x65, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x03, 0x67,
	0x65, 0x74, 0x12, 0x27, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73,
	0x79, 0x6e, 0x63, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x73, 0x79, 0x6e, 0x63, 0x1a,
	0x87, 0x02, 0x0a, 0x07, 0x49, 0x6e, 0x73, 0x74, 0x65, 0x61, 0x64, 0x12, 0x25, 0x0a, 0x03, 0x61,
	0x6c, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x03, 0x61,
	0x6c, 0x6c, 0x12, 0x2b, 0x0a, 0x06, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x06, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12,
	0x2b, 0x0a, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x43, 0x61, 0x6c, 0x6c, 0x52, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x2b, 0x0a, 0x06,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x43, 0x61, 0x6c,
	0x6c, 0x52, 0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x25, 0x0a, 0x03, 0x67, 0x65, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45,
	0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x03, 0x67, 0x65, 0x74,
	0x12, 0x27, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x43,
	0x61, 0x6c, 0x6c, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x69, 0x73, 0x6c, 0x69, 0x62, 0x2f, 0x64,
	0x61, 0x74, 0x61, 0x2d, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_model_extension_proto_rawDescOnce sync.Once
	file_model_extension_proto_rawDescData = file_model_extension_proto_rawDesc
)

func file_model_extension_proto_rawDescGZIP() []byte {
	file_model_extension_proto_rawDescOnce.Do(func() {
		file_model_extension_proto_rawDescData = protoimpl.X.CompressGZIP(file_model_extension_proto_rawDescData)
	})
	return file_model_extension_proto_rawDescData
}

var file_model_extension_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_model_extension_proto_goTypes = []interface{}{
	(*Extension)(nil),         // 0: model.Extension
	(*Extension_After)(nil),   // 1: model.Extension.After
	(*Extension_Before)(nil),  // 2: model.Extension.Before
	(*Extension_Instead)(nil), // 3: model.Extension.Instead
	(*AuditData)(nil),         // 4: model.AuditData
	(*ExternalCall)(nil),      // 5: model.ExternalCall
}
var file_model_extension_proto_depIdxs = []int32{
	2,  // 0: model.Extension.before:type_name -> model.Extension.Before
	3,  // 1: model.Extension.instead:type_name -> model.Extension.Instead
	1,  // 2: model.Extension.after:type_name -> model.Extension.After
	4,  // 3: model.Extension.auditData:type_name -> model.AuditData
	5,  // 4: model.Extension.After.all:type_name -> model.ExternalCall
	5,  // 5: model.Extension.After.create:type_name -> model.ExternalCall
	5,  // 6: model.Extension.After.update:type_name -> model.ExternalCall
	5,  // 7: model.Extension.After.delete:type_name -> model.ExternalCall
	5,  // 8: model.Extension.After.get:type_name -> model.ExternalCall
	5,  // 9: model.Extension.After.list:type_name -> model.ExternalCall
	5,  // 10: model.Extension.Before.all:type_name -> model.ExternalCall
	5,  // 11: model.Extension.Before.create:type_name -> model.ExternalCall
	5,  // 12: model.Extension.Before.update:type_name -> model.ExternalCall
	5,  // 13: model.Extension.Before.delete:type_name -> model.ExternalCall
	5,  // 14: model.Extension.Before.get:type_name -> model.ExternalCall
	5,  // 15: model.Extension.Before.list:type_name -> model.ExternalCall
	5,  // 16: model.Extension.Instead.all:type_name -> model.ExternalCall
	5,  // 17: model.Extension.Instead.create:type_name -> model.ExternalCall
	5,  // 18: model.Extension.Instead.update:type_name -> model.ExternalCall
	5,  // 19: model.Extension.Instead.delete:type_name -> model.ExternalCall
	5,  // 20: model.Extension.Instead.get:type_name -> model.ExternalCall
	5,  // 21: model.Extension.Instead.list:type_name -> model.ExternalCall
	22, // [22:22] is the sub-list for method output_type
	22, // [22:22] is the sub-list for method input_type
	22, // [22:22] is the sub-list for extension type_name
	22, // [22:22] is the sub-list for extension extendee
	0,  // [0:22] is the sub-list for field type_name
}

func init() { file_model_extension_proto_init() }
func file_model_extension_proto_init() {
	if File_model_extension_proto != nil {
		return
	}
	file_model_audit_proto_init()
	file_model_common_proto_init()
	file_model_external_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_model_extension_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Extension); i {
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
		file_model_extension_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Extension_After); i {
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
		file_model_extension_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Extension_Before); i {
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
		file_model_extension_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Extension_Instead); i {
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
			RawDescriptor: file_model_extension_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_model_extension_proto_goTypes,
		DependencyIndexes: file_model_extension_proto_depIdxs,
		MessageInfos:      file_model_extension_proto_msgTypes,
	}.Build()
	File_model_extension_proto = out.File
	file_model_extension_proto_rawDesc = nil
	file_model_extension_proto_goTypes = nil
	file_model_extension_proto_depIdxs = nil
}
