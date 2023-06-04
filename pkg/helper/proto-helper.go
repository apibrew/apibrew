package helper

import (
	"github.com/tislib/data-handler/pkg/model"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

type ProtoParsedData struct {
	Resources []*model.Resource
}

type ProtoHelper struct {
	params ProtoHelperParams
}

func (p *ProtoHelper) ProtoFileToProtoParsedData(protoFileContent string) *ProtoParsedData {
	//TODO implement me
	panic("implement me")
}

func (p *ProtoHelper) ResourceToProtoFileContent(resource *model.Resource) string {
	panic("not implemented")
}

func (p *ProtoHelper) ResourceToProtoDescriptorMessage(resource *model.Resource) *descriptorpb.DescriptorProto {
	panic("not implemented")
}

func (p *ProtoHelper) ResourceToProtoDescriptorFile(resource *model.Resource) *descriptorpb.FileDescriptorProto {
	panic("not implemented")
}

func (p *ProtoHelper) ParseFileDescriptorProto(descriptor *descriptorpb.FileDescriptorProto) *ProtoParsedData {
	var data = new(ProtoParsedData)

	for _, mt := range descriptor.MessageType {
		data.Resources = append(data.Resources, p.DescriptorProtoToResource(mt))
	}

	return data
}

func (p *ProtoHelper) DescriptorProtoToResource(mt *descriptorpb.DescriptorProto) *model.Resource {
	resource := new(model.Resource)

	resource.Name = *mt.Name
	resource.Namespace = "default"

	for _, field := range mt.Field {
		resource.Properties = append(resource.Properties, p.ToProperty(field))
	}

	resource.SourceConfig = &model.ResourceSourceConfig{
		DataSource: proto.GetExtension(mt.GetOptions(), model.E_ResourceDataSource).(string),
		Catalog:    proto.GetExtension(mt.GetOptions(), model.E_ResourceCatalog).(string),
		Entity:     proto.GetExtension(mt.GetOptions(), model.E_ResourceEntity).(string),
	}

	return resource
}

func (p *ProtoHelper) ToProperty(field *descriptorpb.FieldDescriptorProto) *model.ResourceProperty {
	property := &model.ResourceProperty{
		Name: *field.Name,
		Type: protoFieldTypeToPropertyType(field.Type),
	}

	return property
}

func protoFieldTypeToPropertyType(protoType *descriptorpb.FieldDescriptorProto_Type) model.ResourcePropertyType {
	switch *protoType {
	case descriptorpb.FieldDescriptorProto_TYPE_INT32:
		return model.ResourcePropertyType_TYPE_INT32
	case descriptorpb.FieldDescriptorProto_TYPE_STRING:
		return model.ResourcePropertyType_TYPE_STRING
	default:
		panic("Unknown file")
	}
}

type ProtoHelperParams struct {
}

func NewProtoHelper(params ProtoHelperParams) *ProtoHelper {
	return &ProtoHelper{params: params}
}
