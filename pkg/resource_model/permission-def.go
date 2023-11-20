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
			Properties: map[string]*model.ResourceProperty{
				"createdOn": {
					Type:         model.ResourceProperty_TIMESTAMP,
					Immutable:    true,
					ExampleValue: structpb.NewStringValue("2023-11-20T23:38:43+04:00"),

					Annotations: map[string]string{
						"SpecialProperty": "true",
					},
				},
				"updatedOn": {
					Type:         model.ResourceProperty_TIMESTAMP,
					ExampleValue: structpb.NewStringValue("2023-11-20T23:38:43+04:00"),

					Annotations: map[string]string{
						"SpecialProperty": "true",
					},
				},
				"createdBy": {
					Type:         model.ResourceProperty_STRING,
					Length:       256,
					Immutable:    true,
					ExampleValue: structpb.NewStringValue("admin"),

					Annotations: map[string]string{
						"SpecialProperty": "true",
					},
				},
				"updatedBy": {
					Type:         model.ResourceProperty_STRING,
					Length:       256,
					ExampleValue: structpb.NewStringValue("admin"),

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
			Properties: map[string]*model.ResourceProperty{
				"regexMatch": {
					Type: model.ResourceProperty_STRUCT,
				},
				"and": {
					Type: model.ResourceProperty_LIST,
				},
				"or": {
					Type: model.ResourceProperty_LIST,
				},
				"greaterThanOrEqual": {
					Type: model.ResourceProperty_STRUCT,
				},
				"greaterThan": {
					Type: model.ResourceProperty_STRUCT,
				},
				"lessThanOrEqual": {
					Type: model.ResourceProperty_STRUCT,
				},
				"in": {
					Type: model.ResourceProperty_STRUCT,
				},
				"isNull": {
					Type: model.ResourceProperty_STRUCT,
				},
				"not": {
					Type: model.ResourceProperty_STRUCT,
				},
				"equal": {
					Type: model.ResourceProperty_STRUCT,
				},
				"lessThan": {
					Type: model.ResourceProperty_STRUCT,
				},
			},

			Annotations: map[string]string{
				"EnableAudit":  "true",
				"OpenApiGroup": "meta",
			},
		},
		{
			Name: "PairExpression",
			Properties: map[string]*model.ResourceProperty{
				"left": {
					Type: model.ResourceProperty_STRUCT,
				},
				"right": {
					Type: model.ResourceProperty_STRUCT,
				},
			},

			Annotations: map[string]string{
				"EnableAudit":  "true",
				"OpenApiGroup": "meta",
			},
		},
		{
			Name: "RegexMatchExpression",
			Properties: map[string]*model.ResourceProperty{
				"pattern": {
					Type: model.ResourceProperty_STRING,
				},
				"expression": {
					Type: model.ResourceProperty_STRUCT,
				},
			},

			Annotations: map[string]string{
				"EnableAudit":  "true",
				"OpenApiGroup": "meta",
			},
		},
		{
			Name: "Expression",
			Properties: map[string]*model.ResourceProperty{
				"property": {
					Type: model.ResourceProperty_STRING,
				},
				"value": {
					Type: model.ResourceProperty_OBJECT,
				},
			},

			Annotations: map[string]string{
				"EnableAudit":  "true",
				"OpenApiGroup": "meta",
			},
		},
	},
	Properties: map[string]*model.ResourceProperty{
		"auditData": {
			Type:         model.ResourceProperty_STRUCT,
			ExampleValue: structpb.NewStructValue(&structpb.Struct{Fields: map[string]*structpb.Value{"createdBy": structpb.NewStringValue("admin"), "updatedBy": structpb.NewStringValue("admin"), "createdOn": structpb.NewStringValue("2023-11-20T23:38:43+04:00"), "updatedOn": structpb.NewStringValue("2023-11-20T23:38:43+04:00")}}),

			Annotations: map[string]string{
				"SpecialProperty": "true",
			},
		},
		"recordSelector": {
			Type: model.ResourceProperty_STRUCT,
		},
		"after": {
			Type: model.ResourceProperty_TIMESTAMP,
		},
		"user": {
			Type: model.ResourceProperty_REFERENCE,
		},
		"permit": {
			Type:     model.ResourceProperty_ENUM,
			Length:   255,
			Required: true,

			Annotations: map[string]string{
				"TypeName": "PermissionPermit",
			},
		},
		"id": {
			Type:         model.ResourceProperty_UUID,
			Required:     true,
			Immutable:    true,
			ExampleValue: structpb.NewStringValue("a39621a4-6d48-11ee-b962-0242ac120002"),

			Annotations: map[string]string{
				"SpecialProperty": "true",
				"PrimaryProperty": "true",
			},
		},
		"version": {
			Type:         model.ResourceProperty_INT32,
			Required:     true,
			DefaultValue: structpb.NewNumberValue(1),
			ExampleValue: structpb.NewNumberValue(1),

			Annotations: map[string]string{
				"SpecialProperty":     "true",
				"AllowEmptyPrimitive": "true",
			},
		},
		"namespace": {
			Type:         model.ResourceProperty_STRING,
			Length:       255,
			ExampleValue: structpb.NewStringValue("default"),
		},
		"resource": {
			Type:         model.ResourceProperty_STRING,
			Length:       255,
			ExampleValue: structpb.NewStringValue("Book"),
		},
		"operation": {
			Type:         model.ResourceProperty_ENUM,
			Length:       255,
			Required:     true,
			DefaultValue: structpb.NewStringValue("FULL"),
			ExampleValue: structpb.NewStringValue("READ"),

			Annotations: map[string]string{
				"TypeName": "PermissionOperation",
			},
		},
		"before": {
			Type: model.ResourceProperty_TIMESTAMP,
		},
		"role": {
			Type: model.ResourceProperty_REFERENCE,
		},
		"localFlags": {
			Type: model.ResourceProperty_OBJECT,
		},
	},

	Annotations: map[string]string{
		"EnableAudit":  "true",
		"OpenApiGroup": "meta",
	},
}