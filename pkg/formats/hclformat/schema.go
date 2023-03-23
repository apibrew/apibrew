package hclformat

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/resources"
	"github.com/tislib/data-handler/pkg/service/annotations"
	"google.golang.org/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
)

func prepareSchema([]*model.Resource) *hcl.BodySchema {

	var schema = &hcl.BodySchema{
		Attributes: []hcl.AttributeSchema{},
		Blocks: []hcl.BlockHeaderSchema{
			{
				Type: "record",
				LabelNames: []string{
					"resource",
				},
			},
		},
	}

	for _, resource := range resources.GetAllSystemResources() {
		schema.Blocks = append(schema.Blocks, prepareResourceBlockDefinition(resource))
	}

	return schema
}

func prepareSystemResourceSchema(systemResource proto.Message) *hcl.BodySchema {
	var attributes []hcl.AttributeSchema

	systemResource.ProtoReflect().Range(func(descriptor protoreflect.FieldDescriptor, value protoreflect.Value) bool {
		attributes = append(attributes, hcl.AttributeSchema{
			Name:     string(descriptor.Name()),
			Required: false,
		})

		return true
	})

	descriptor := systemResource.ProtoReflect().Descriptor()

	for index := 0; index < descriptor.Fields().Len(); index++ {
		field := descriptor.Fields().Get(index)

		attributes = append(attributes, hcl.AttributeSchema{
			Name:     string(field.Name()),
			Required: false,
		})
	}

	return &hcl.BodySchema{
		Attributes: attributes,
		Blocks:     nil,
	}
}

func prepareResourceBlockDefinition(resource *model.Resource) hcl.BlockHeaderSchema {
	bhs := hcl.BlockHeaderSchema{}

	for _, prop := range resource.Properties {
		if annotations.IsEnabled(prop, annotations.HclLabelProperty) {
			bhs.LabelNames = append(bhs.LabelNames, prop.Name)
		}
	}
	bhs.Type = resource.Name

	return bhs
}
