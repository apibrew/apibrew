// Code generated by apbr generate. DO NOT EDIT.
// versions:
// 	apbr generate v1.2

//go:build !codeanalysis

package testing

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/util"
	"google.golang.org/protobuf/types/known/structpb"
)

var TestCaseResource = &model.Resource{
	Name:        "TestCase",
	Namespace:   "testing",
	Title:       util.Pointer("Test Case"),
	Description: util.Pointer("Test Case is a test case"),
	Types: []*model.ResourceSubType{
		{
			Name: "TestCaseStep",
			Properties: []*model.ResourceProperty{
				{
					Name:       "operation",
					Type:       model.ResourceProperty_ENUM,
					Required:   true,
					EnumValues: []string{"CREATE", "UPDATE", "APPLY", "DELETE", "GET", "LIST", "NANO"},
				},
				{
					Name: "payload",
					Type: model.ResourceProperty_OBJECT,
				},
				{
					Name:   "name",
					Type:   model.ResourceProperty_STRING,
					Length: 255,
					Unique: true,
				},
			},

			Annotations: map[string]string{
				"NormalizedResource": "true",
			},
		},
		{
			Name: "TestCaseAssertion",
			Properties: []*model.ResourceProperty{
				{
					Name:   "name",
					Type:   model.ResourceProperty_STRING,
					Length: 255,
				},
				{
					Name:       "assertionType",
					Type:       model.ResourceProperty_ENUM,
					Required:   true,
					EnumValues: []string{"EQUAL", "NOT_EQUAL", "EXPECT_ERROR", "NANO"},
				},
				{
					Name: "left",
					Type: model.ResourceProperty_STRING,
				},
				{
					Name: "right",
					Type: model.ResourceProperty_OBJECT,
				},
				{
					Name: "script",
					Type: model.ResourceProperty_STRING,
				},
				{
					Name: "errorCode",
					Type: model.ResourceProperty_STRING,
				},
			},

			Annotations: map[string]string{
				"NormalizedResource": "true",
			},
		},
	},
	Properties: []*model.ResourceProperty{
		{
			Name:         "id",
			Type:         model.ResourceProperty_UUID,
			Required:     true,
			Immutable:    true,
			ExampleValue: structpb.NewStringValue("a39621a4-6d48-11ee-b962-0242ac120002"),

			Annotations: map[string]string{
				"PrimaryProperty": "true",
				"SpecialProperty": "true",
			},
		},
		{
			Name: "steps",
			Type: model.ResourceProperty_LIST,
			Item: &model.ResourceProperty{
				Name:    "",
				Type:    model.ResourceProperty_STRUCT,
				TypeRef: util.Pointer("TestCaseStep"),
			},
		},
		{
			Name: "assertions",
			Type: model.ResourceProperty_LIST,
			Item: &model.ResourceProperty{
				Name:    "",
				Type:    model.ResourceProperty_STRUCT,
				TypeRef: util.Pointer("TestCaseAssertion"),
			},
		},
		{
			Name:      "testSuite",
			Type:      model.ResourceProperty_REFERENCE,
			Required:  true,
			Reference: &model.Reference{Resource: "TestSuite", Namespace: "testing"},
		},
		{
			Name:     "name",
			Type:     model.ResourceProperty_STRING,
			Length:   255,
			Required: true,
		},
		{
			Name:   "description",
			Type:   model.ResourceProperty_STRING,
			Length: 64000,
		},
		{
			Name: "annotations",
			Type: model.ResourceProperty_MAP,
			Item: &model.ResourceProperty{
				Name: "",
				Type: model.ResourceProperty_STRING,
			},
		},
		{
			Name:         "version",
			Type:         model.ResourceProperty_INT32,
			Required:     true,
			DefaultValue: structpb.NewNumberValue(1),
			ExampleValue: structpb.NewNumberValue(1),

			Annotations: map[string]string{
				"SpecialProperty":     "true",
				"AllowEmptyPrimitive": "true",
			},
		},
	},
	Indexes: []*model.ResourceIndex{
		{
			Properties: []*model.ResourceIndexProperty{
				{
					Name:  "testSuite",
					Order: model.Order_ORDER_UNKNOWN,
				},
				{
					Name:  "name",
					Order: model.Order_ORDER_UNKNOWN,
				},
			},
			IndexType: model.ResourceIndexType_BTREE,
			Unique:    true,
		},
	},

	Annotations: map[string]string{
		"NormalizedResource": "true",
	},
}
