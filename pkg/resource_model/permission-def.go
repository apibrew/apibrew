// Code generated by apbr generate. DO NOT EDIT.
// versions:
// 	apbr generate v1.2

//go:build !codeanalysis

package resource_model

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/util"
	"google.golang.org/protobuf/types/known/structpb"
)

var PermissionResource = &model.Resource{
	Name:        "Permission",
	Namespace:   "system",
	Title:       util.Pointer("Permission"),
	Description: util.Pointer("Permission is a resource that defines the access control rules for resources for users."),
	Types: []*model.ResourceSubType{
		{
			Name:        "AuditData",
			Title:       "Audit Data",
			Description: "Audit Data is a type that represents the audit data of a resource/record. ",
			Properties: []*model.ResourceProperty{
				{
					Name:         "createdBy",
					Type:         model.ResourceProperty_STRING,
					Length:       256,
					Immutable:    true,
					ExampleValue: structpb.NewStringValue("admin"),

					Annotations: map[string]string{
						"SpecialProperty": "true",
					},
				},
				{
					Name:         "updatedBy",
					Type:         model.ResourceProperty_STRING,
					Length:       256,
					ExampleValue: structpb.NewStringValue("admin"),

					Annotations: map[string]string{
						"SpecialProperty": "true",
					},
				},
				{
					Name:         "createdOn",
					Type:         model.ResourceProperty_TIMESTAMP,
					Immutable:    true,
					ExampleValue: structpb.NewStringValue("2024-05-03T21:31:07+04:00"),

					Annotations: map[string]string{
						"SpecialProperty": "true",
					},
				},
				{
					Name:         "updatedOn",
					Type:         model.ResourceProperty_TIMESTAMP,
					ExampleValue: structpb.NewStringValue("2024-05-03T21:31:07+04:00"),

					Annotations: map[string]string{
						"SpecialProperty": "true",
					},
				},
			},

			Annotations: map[string]string{
				"EnableAudit":  "true",
				"OpenApiGroup": "meta",
			},
		},
		{
			Name: "BooleanExpression",
			Properties: []*model.ResourceProperty{
				{
					Name: "and",
					Type: model.ResourceProperty_LIST,
					Item: &model.ResourceProperty{
						Name:    "",
						Type:    model.ResourceProperty_STRUCT,
						TypeRef: util.Pointer("BooleanExpression"),
					},
				},
				{
					Name: "or",
					Type: model.ResourceProperty_LIST,
					Item: &model.ResourceProperty{
						Name:    "",
						Type:    model.ResourceProperty_STRUCT,
						TypeRef: util.Pointer("BooleanExpression"),
					},
				},
				{
					Name:    "not",
					Type:    model.ResourceProperty_STRUCT,
					TypeRef: util.Pointer("BooleanExpression"),
				},
				{
					Name:    "equal",
					Type:    model.ResourceProperty_STRUCT,
					TypeRef: util.Pointer("PairExpression"),
				},
				{
					Name:    "lessThan",
					Type:    model.ResourceProperty_STRUCT,
					TypeRef: util.Pointer("PairExpression"),
				},
				{
					Name:    "greaterThan",
					Type:    model.ResourceProperty_STRUCT,
					TypeRef: util.Pointer("PairExpression"),
				},
				{
					Name:    "lessThanOrEqual",
					Type:    model.ResourceProperty_STRUCT,
					TypeRef: util.Pointer("PairExpression"),
				},
				{
					Name:    "greaterThanOrEqual",
					Type:    model.ResourceProperty_STRUCT,
					TypeRef: util.Pointer("PairExpression"),
				},
				{
					Name:    "in",
					Type:    model.ResourceProperty_STRUCT,
					TypeRef: util.Pointer("PairExpression"),
				},
				{
					Name:    "like",
					Type:    model.ResourceProperty_STRUCT,
					TypeRef: util.Pointer("PairExpression"),
				},
				{
					Name:    "ilike",
					Type:    model.ResourceProperty_STRUCT,
					TypeRef: util.Pointer("PairExpression"),
				},
				{
					Name:    "regex",
					Type:    model.ResourceProperty_STRUCT,
					TypeRef: util.Pointer("PairExpression"),
				},
				{
					Name:    "isNull",
					Type:    model.ResourceProperty_STRUCT,
					TypeRef: util.Pointer("Expression"),
				},
				{
					Name: "filters",
					Type: model.ResourceProperty_MAP,
					Item: &model.ResourceProperty{
						Name: "",
						Type: model.ResourceProperty_OBJECT,
					},
				},
			},

			Annotations: map[string]string{
				"EnableAudit":  "true",
				"OpenApiGroup": "meta",
			},
		},
		{
			Name: "PairExpression",
			Properties: []*model.ResourceProperty{
				{
					Name:    "left",
					Type:    model.ResourceProperty_STRUCT,
					TypeRef: util.Pointer("Expression"),
				},
				{
					Name:    "right",
					Type:    model.ResourceProperty_STRUCT,
					TypeRef: util.Pointer("Expression"),
				},
			},

			Annotations: map[string]string{
				"OpenApiGroup": "meta",
				"EnableAudit":  "true",
			},
		},
		{
			Name: "Expression",
			Properties: []*model.ResourceProperty{
				{
					Name: "property",
					Type: model.ResourceProperty_STRING,
				},
				{
					Name: "value",
					Type: model.ResourceProperty_OBJECT,
				},
			},

			Annotations: map[string]string{
				"EnableAudit":  "true",
				"OpenApiGroup": "meta",
			},
		},
	},
	Properties: []*model.ResourceProperty{
		{
			Name:         "id",
			Type:         model.ResourceProperty_UUID,
			Primary:      true,
			Required:     true,
			Immutable:    true,
			ExampleValue: structpb.NewStringValue("a39621a4-6d48-11ee-b962-0242ac120002"),

			Annotations: map[string]string{
				"SpecialProperty": "true",
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
		{
			Name:         "auditData",
			Type:         model.ResourceProperty_STRUCT,
			TypeRef:      util.Pointer("AuditData"),
			ExampleValue: structpb.NewStructValue(&structpb.Struct{Fields: map[string]*structpb.Value{"createdBy": structpb.NewStringValue("admin"), "updatedBy": structpb.NewStringValue("admin"), "createdOn": structpb.NewStringValue("2024-05-03T21:31:07+04:00"), "updatedOn": structpb.NewStringValue("2024-05-03T21:31:07+04:00")}}),

			Annotations: map[string]string{
				"SpecialProperty": "true",
			},
		},
		{
			Name:         "namespace",
			Type:         model.ResourceProperty_STRING,
			Length:       255,
			ExampleValue: structpb.NewStringValue("default"),
		},
		{
			Name:         "resource",
			Type:         model.ResourceProperty_STRING,
			Length:       255,
			ExampleValue: structpb.NewStringValue("Book"),
		},
		{
			Name:    "recordSelector",
			Type:    model.ResourceProperty_STRUCT,
			TypeRef: util.Pointer("BooleanExpression"),
		},
		{
			Name:         "operation",
			Type:         model.ResourceProperty_ENUM,
			Length:       255,
			Required:     true,
			DefaultValue: structpb.NewStringValue("FULL"),
			ExampleValue: structpb.NewStringValue("READ"),
			EnumValues:   []string{"READ", "CREATE", "UPDATE", "DELETE", "FULL"},
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
			Name:      "user",
			Type:      model.ResourceProperty_REFERENCE,
			Reference: &model.Reference{Resource: "User", Namespace: "system"},
		},
		{
			Name:      "role",
			Type:      model.ResourceProperty_REFERENCE,
			Reference: &model.Reference{Resource: "Role", Namespace: "system"},
		},
		{
			Name:       "permit",
			Type:       model.ResourceProperty_ENUM,
			Length:     255,
			Required:   true,
			EnumValues: []string{"ALLOW", "REJECT"},
		},
		{
			Name: "localFlags",
			Type: model.ResourceProperty_OBJECT,
		},
	},

	Annotations: map[string]string{
		"EnableAudit":  "true",
		"OpenApiGroup": "meta",
	},
}
