package sub_types

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/service/annotations"
	"google.golang.org/protobuf/types/known/structpb"
)

var SecurityConstraint = &model.ResourceSubType{
	Name: "SecurityConstraint",
	Properties: []*model.ResourceProperty{
		{
			Name:         "namespace",
			Type:         model.ResourceProperty_STRING,
			Required:     true,
			DefaultValue: structpb.NewStringValue("*"),
		},
		{
			Name:         "resource",
			Type:         model.ResourceProperty_STRING,
			Required:     true,
			DefaultValue: structpb.NewStringValue("*"),
		},
		{
			Name:         "property",
			Type:         model.ResourceProperty_STRING,
			Required:     true,
			DefaultValue: structpb.NewStringValue("*"),
		},
		{
			Name:         "propertyValue",
			Type:         model.ResourceProperty_STRING,
			Required:     false,
			DefaultValue: structpb.NewStringValue("*"),
		},
		{
			Name: "propertyMode",
			Type: model.ResourceProperty_ENUM,
			EnumValues: []string{
				"PROPERTY_MATCH_ONLY",
				"PROPERTY_MATCH_ANY",
			},
		},
		{
			Name: "operation",
			Type: model.ResourceProperty_ENUM,
			EnumValues: []string{
				"OPERATION_TYPE_READ",
				"OPERATION_TYPE_CREATE",
				"OPERATION_TYPE_UPDATE",
				"OPERATION_TYPE_DELETE",
				"FULL",
			},
		},
		{
			Name: "recordIds",
			Type: model.ResourceProperty_LIST,
			Item: &model.ResourceProperty{
				Type: model.ResourceProperty_STRING,
			},
		},
		{
			Name: "before",
			Type: model.ResourceProperty_TIMESTAMP,
		},
		{
			Name: "after",
			Type: model.ResourceProperty_TIMESTAMP,
		},
		{
			Name: "username",
			Type: model.ResourceProperty_STRING,
		},
		{
			Name: "role",
			Type: model.ResourceProperty_STRING,
		},
		{
			Name:     "permit",
			Required: true,
			Type:     model.ResourceProperty_ENUM,
			EnumValues: []string{
				"PERMIT_TYPE_ALLOW",
				"PERMIT_TYPE_REJECT",
			},
		},
		{
			Name: "localFlags",
			Type: model.ResourceProperty_OBJECT,
		},
	},
	Annotations: map[string]string{
		annotations.GenericType: "true",
	},
}
