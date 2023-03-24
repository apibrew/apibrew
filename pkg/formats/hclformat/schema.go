package hclformat

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/resources"
	"github.com/tislib/data-handler/pkg/service/annotations"
	"github.com/tislib/data-handler/pkg/util"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func prepareSchema(list []*model.Resource) *hcl.BodySchema {
	var schema = &hcl.BodySchema{
		Attributes: []hcl.AttributeSchema{},
		Blocks: []hcl.BlockHeaderSchema{
			{
				Type: "record",
				LabelNames: []string{
					"namespace",
					"resource",
				},
			},
		},
	}

	for _, resource := range resources.GetAllSystemResources() {
		schema.Blocks = append(schema.Blocks, prepareSystemResourceSelfSchema(resources.GetSystemResourceType(resource).ProtoReflect().Descriptor()))
	}

	return schema
}

func prepareSystemResourceSelfSchema(descriptor protoreflect.MessageDescriptor) hcl.BlockHeaderSchema {
	blockSchema := hcl.BlockHeaderSchema{
		Type: util.ToSnakeCase(string(descriptor.Name())),
	}

	for index := 0; index < descriptor.Fields().Len(); index++ {
		field := descriptor.Fields().Get(index)

		hclLabel := proto.GetExtension(field.Options(), model.E_HclLabel)

		if hclLabel != "" {
			blockSchema.LabelNames = append(blockSchema.LabelNames, util.ToSnakeCase(hclLabel.(string)))
		}
	}

	return blockSchema
}

func prepareResourceRecordSchema(resource *model.Resource) *hcl.BodySchema {
	var attributes []hcl.AttributeSchema
	var blocks []hcl.BlockHeaderSchema

	for _, prop := range resource.Properties {
		isSpecial := annotations.IsEnabled(prop, annotations.SpecialProperty)
		blockProperty := annotations.Get(prop, annotations.HclBlockProperty)

		if blockProperty != "" {
			blocks = append(blocks, hcl.BlockHeaderSchema{
				Type: util.ToSnakeCase(blockProperty),
			})
		} else {
			attributes = append(attributes, hcl.AttributeSchema{
				Name:     util.ToSnakeCase(prop.Name),
				Required: prop.Required && !isSpecial,
			})
		}
	}

	return &hcl.BodySchema{
		Attributes: attributes,
		Blocks:     blocks,
	}
}

func prepareSystemResourceSchema(descriptor protoreflect.MessageDescriptor) *hcl.BodySchema {
	var attributes []hcl.AttributeSchema
	var blocks []hcl.BlockHeaderSchema

	for index := 0; index < descriptor.Fields().Len(); index++ {
		field := descriptor.Fields().Get(index)

		if field.Kind() == protoreflect.MessageKind {
			hclBlock := proto.GetExtension(field.Options(), model.E_HclBlock)

			block := prepareSystemResourceSelfSchema(field.Message())

			if hclBlock != "" {
				block.Type = hclBlock.(string)
			} else {
				block.Type = util.ToSnakeCase(string(field.Name()))
			}

			blocks = append(blocks, block)
		} else {
			attributes = append(attributes, hcl.AttributeSchema{
				Name:     util.ToSnakeCase(string(field.Name())),
				Required: false,
			})
		}

	}

	return &hcl.BodySchema{
		Attributes: attributes,
		Blocks:     blocks,
	}
}
