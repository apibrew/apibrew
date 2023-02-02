// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
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

type DataSourceBackendType int32

const (
	DataSourceBackendType_POSTGRESQL DataSourceBackendType = 0
	DataSourceBackendType_VIRTUAL    DataSourceBackendType = 1
	DataSourceBackendType_MYSQL      DataSourceBackendType = 2
	DataSourceBackendType_ORACLE     DataSourceBackendType = 3
	DataSourceBackendType_MONGODB    DataSourceBackendType = 4
	DataSourceBackendType_CUSTOM     DataSourceBackendType = 5
)

// Enum value maps for DataSourceBackendType.
var (
	DataSourceBackendType_name = map[int32]string{
		0: "POSTGRESQL",
		1: "VIRTUAL",
		2: "MYSQL",
		3: "ORACLE",
		4: "MONGODB",
		5: "CUSTOM",
	}
	DataSourceBackendType_value = map[string]int32{
		"POSTGRESQL": 0,
		"VIRTUAL":    1,
		"MYSQL":      2,
		"ORACLE":     3,
		"MONGODB":    4,
		"CUSTOM":     5,
	}
)

func (x DataSourceBackendType) Enum() *DataSourceBackendType {
	p := new(DataSourceBackendType)
	*p = x
	return p
}

func (x DataSourceBackendType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DataSourceBackendType) Descriptor() protoreflect.EnumDescriptor {
	return file_model_data_source_proto_enumTypes[0].Descriptor()
}

func (DataSourceBackendType) Type() protoreflect.EnumType {
	return &file_model_data_source_proto_enumTypes[0]
}

func (x DataSourceBackendType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DataSourceBackendType.Descriptor instead.
func (DataSourceBackendType) EnumDescriptor() ([]byte, []int) {
	return file_model_data_source_proto_rawDescGZIP(), []int{0}
}

type VirtualOptions_Mode int32

const (
	VirtualOptions_DISCARD VirtualOptions_Mode = 0
	VirtualOptions_ERROR   VirtualOptions_Mode = 1
)

// Enum value maps for VirtualOptions_Mode.
var (
	VirtualOptions_Mode_name = map[int32]string{
		0: "DISCARD",
		1: "ERROR",
	}
	VirtualOptions_Mode_value = map[string]int32{
		"DISCARD": 0,
		"ERROR":   1,
	}
)

func (x VirtualOptions_Mode) Enum() *VirtualOptions_Mode {
	p := new(VirtualOptions_Mode)
	*p = x
	return p
}

func (x VirtualOptions_Mode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (VirtualOptions_Mode) Descriptor() protoreflect.EnumDescriptor {
	return file_model_data_source_proto_enumTypes[1].Descriptor()
}

func (VirtualOptions_Mode) Type() protoreflect.EnumType {
	return &file_model_data_source_proto_enumTypes[1]
}

func (x VirtualOptions_Mode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use VirtualOptions_Mode.Descriptor instead.
func (VirtualOptions_Mode) EnumDescriptor() ([]byte, []int) {
	return file_model_data_source_proto_rawDescGZIP(), []int{1, 0}
}

type PostgresqlOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username      string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password      string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Host          string `protobuf:"bytes,3,opt,name=host,proto3" json:"host,omitempty"`
	Port          uint32 `protobuf:"varint,4,opt,name=port,proto3" json:"port,omitempty"`
	DbName        string `protobuf:"bytes,5,opt,name=dbName,proto3" json:"dbName,omitempty"`
	DefaultSchema string `protobuf:"bytes,6,opt,name=defaultSchema,proto3" json:"defaultSchema,omitempty"`
}

func (x *PostgresqlOptions) Reset() {
	*x = PostgresqlOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_data_source_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostgresqlOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostgresqlOptions) ProtoMessage() {}

func (x *PostgresqlOptions) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use PostgresqlOptions.ProtoReflect.Descriptor instead.
func (*PostgresqlOptions) Descriptor() ([]byte, []int) {
	return file_model_data_source_proto_rawDescGZIP(), []int{0}
}

func (x *PostgresqlOptions) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *PostgresqlOptions) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *PostgresqlOptions) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *PostgresqlOptions) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *PostgresqlOptions) GetDbName() string {
	if x != nil {
		return x.DbName
	}
	return ""
}

func (x *PostgresqlOptions) GetDefaultSchema() string {
	if x != nil {
		return x.DefaultSchema
	}
	return ""
}

type VirtualOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mode VirtualOptions_Mode `protobuf:"varint,1,opt,name=mode,proto3,enum=model.VirtualOptions_Mode" json:"mode,omitempty"`
}

func (x *VirtualOptions) Reset() {
	*x = VirtualOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_data_source_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VirtualOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VirtualOptions) ProtoMessage() {}

func (x *VirtualOptions) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use VirtualOptions.ProtoReflect.Descriptor instead.
func (*VirtualOptions) Descriptor() ([]byte, []int) {
	return file_model_data_source_proto_rawDescGZIP(), []int{1}
}

func (x *VirtualOptions) GetMode() VirtualOptions_Mode {
	if x != nil {
		return x.Mode
	}
	return VirtualOptions_DISCARD
}

type DataSource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                string                `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // id; read only
	Backend           DataSourceBackendType `protobuf:"varint,2,opt,name=backend,proto3,enum=model.DataSourceBackendType" json:"backend,omitempty"`
	CustomBackendName string                `protobuf:"bytes,3,opt,name=customBackendName,proto3" json:"customBackendName,omitempty"`
	Type              DataType              `protobuf:"varint,4,opt,name=type,proto3,enum=model.DataType" json:"type,omitempty"` // read only
	Name              string                `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Description       string                `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	// Types that are assignable to Options:
	//
	//	*DataSource_PostgresqlParams
	//	*DataSource_VirtualParams
	Options   isDataSource_Options `protobuf_oneof:"options"`
	AuditData *AuditData           `protobuf:"bytes,101,opt,name=auditData,proto3" json:"auditData,omitempty"` // read only
	Version   uint32               `protobuf:"varint,102,opt,name=version,proto3" json:"version,omitempty"`    // read only
}

func (x *DataSource) Reset() {
	*x = DataSource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_data_source_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataSource) ProtoMessage() {}

func (x *DataSource) ProtoReflect() protoreflect.Message {
	mi := &file_model_data_source_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataSource.ProtoReflect.Descriptor instead.
func (*DataSource) Descriptor() ([]byte, []int) {
	return file_model_data_source_proto_rawDescGZIP(), []int{2}
}

func (x *DataSource) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DataSource) GetBackend() DataSourceBackendType {
	if x != nil {
		return x.Backend
	}
	return DataSourceBackendType_POSTGRESQL
}

func (x *DataSource) GetCustomBackendName() string {
	if x != nil {
		return x.CustomBackendName
	}
	return ""
}

func (x *DataSource) GetType() DataType {
	if x != nil {
		return x.Type
	}
	return DataType_SYSTEM
}

func (x *DataSource) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DataSource) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (m *DataSource) GetOptions() isDataSource_Options {
	if m != nil {
		return m.Options
	}
	return nil
}

func (x *DataSource) GetPostgresqlParams() *PostgresqlOptions {
	if x, ok := x.GetOptions().(*DataSource_PostgresqlParams); ok {
		return x.PostgresqlParams
	}
	return nil
}

func (x *DataSource) GetVirtualParams() *VirtualOptions {
	if x, ok := x.GetOptions().(*DataSource_VirtualParams); ok {
		return x.VirtualParams
	}
	return nil
}

func (x *DataSource) GetAuditData() *AuditData {
	if x != nil {
		return x.AuditData
	}
	return nil
}

func (x *DataSource) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

type isDataSource_Options interface {
	isDataSource_Options()
}

type DataSource_PostgresqlParams struct {
	PostgresqlParams *PostgresqlOptions `protobuf:"bytes,7,opt,name=postgresqlParams,proto3,oneof"`
}

type DataSource_VirtualParams struct {
	VirtualParams *VirtualOptions `protobuf:"bytes,8,opt,name=virtualParams,proto3,oneof"`
}

func (*DataSource_PostgresqlParams) isDataSource_Options() {}

func (*DataSource_VirtualParams) isDataSource_Options() {}

var File_model_data_source_proto protoreflect.FileDescriptor

var file_model_data_source_proto_rawDesc = []byte{
	0x0a, 0x17, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2d, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x1a, 0x11, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb1, 0x01, 0x0a, 0x11, 0x50, 0x6f, 0x73, 0x74,
	0x67, 0x72, 0x65, 0x73, 0x71, 0x6c, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1a, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x64, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64,
	0x62, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74,
	0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x64, 0x65,
	0x66, 0x61, 0x75, 0x6c, 0x74, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x22, 0x60, 0x0a, 0x0e, 0x56,
	0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2e, 0x0a,
	0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x22, 0x1e, 0x0a,
	0x04, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x49, 0x53, 0x43, 0x41, 0x52, 0x44,
	0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x01, 0x22, 0xb9, 0x03,
	0x0a, 0x0a, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x36, 0x0a, 0x07,
	0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x62, 0x61, 0x63,
	0x6b, 0x65, 0x6e, 0x64, 0x12, 0x2c, 0x0a, 0x11, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x42, 0x61,
	0x63, 0x6b, 0x65, 0x6e, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x11, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x23, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x0f, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x46, 0x0a,
	0x10, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x71, 0x6c, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x50, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x71, 0x6c, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x48, 0x00, 0x52, 0x10, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x71, 0x6c, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x3d, 0x0a, 0x0d, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x48, 0x00, 0x52, 0x0d, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x12, 0x2e, 0x0a, 0x09, 0x61, 0x75, 0x64, 0x69, 0x74, 0x44, 0x61, 0x74,
	0x61, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x41, 0x75, 0x64, 0x69, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x09, 0x61, 0x75, 0x64, 0x69, 0x74,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x66, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x09,
	0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2a, 0x64, 0x0a, 0x15, 0x44, 0x61, 0x74,
	0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x4f, 0x53, 0x54, 0x47, 0x52, 0x45, 0x53, 0x51, 0x4c,
	0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x56, 0x49, 0x52, 0x54, 0x55, 0x41, 0x4c, 0x10, 0x01, 0x12,
	0x09, 0x0a, 0x05, 0x4d, 0x59, 0x53, 0x51, 0x4c, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x4f, 0x52,
	0x41, 0x43, 0x4c, 0x45, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x4d, 0x4f, 0x4e, 0x47, 0x4f, 0x44,
	0x42, 0x10, 0x04, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x10, 0x05, 0x42,
	0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x69,
	0x73, 0x6c, 0x69, 0x62, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2d, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65,
	0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
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

var file_model_data_source_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_model_data_source_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_model_data_source_proto_goTypes = []interface{}{
	(DataSourceBackendType)(0), // 0: model.DataSourceBackendType
	(VirtualOptions_Mode)(0),   // 1: model.VirtualOptions.Mode
	(*PostgresqlOptions)(nil),  // 2: model.PostgresqlOptions
	(*VirtualOptions)(nil),     // 3: model.VirtualOptions
	(*DataSource)(nil),         // 4: model.DataSource
	(DataType)(0),              // 5: model.DataType
	(*AuditData)(nil),          // 6: model.AuditData
}
var file_model_data_source_proto_depIdxs = []int32{
	1, // 0: model.VirtualOptions.mode:type_name -> model.VirtualOptions.Mode
	0, // 1: model.DataSource.backend:type_name -> model.DataSourceBackendType
	5, // 2: model.DataSource.type:type_name -> model.DataType
	2, // 3: model.DataSource.postgresqlParams:type_name -> model.PostgresqlOptions
	3, // 4: model.DataSource.virtualParams:type_name -> model.VirtualOptions
	6, // 5: model.DataSource.auditData:type_name -> model.AuditData
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_model_data_source_proto_init() }
func file_model_data_source_proto_init() {
	if File_model_data_source_proto != nil {
		return
	}
	file_model_audit_proto_init()
	file_model_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_model_data_source_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostgresqlOptions); i {
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
			switch v := v.(*VirtualOptions); i {
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
		file_model_data_source_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataSource); i {
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
	file_model_data_source_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*DataSource_PostgresqlParams)(nil),
		(*DataSource_VirtualParams)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_model_data_source_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_model_data_source_proto_goTypes,
		DependencyIndexes: file_model_data_source_proto_depIdxs,
		EnumInfos:         file_model_data_source_proto_enumTypes,
		MessageInfos:      file_model_data_source_proto_msgTypes,
	}.Build()
	File_model_data_source_proto = out.File
	file_model_data_source_proto_rawDesc = nil
	file_model_data_source_proto_goTypes = nil
	file_model_data_source_proto_depIdxs = nil
}
