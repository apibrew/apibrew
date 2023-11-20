// Code generated by apbr generate. DO NOT EDIT.
// versions:
// 	apbr generate v1.2

//go:build !codeanalysis

package resource_model

import (
	"github.com/apibrew/apibrew/pkg/model"
	"google.golang.org/protobuf/types/known/structpb"
)

var ExtensionResource = &model.Resource{
	Name:      "Extension",
	Namespace: "system",
	Types: []*model.ResourceSubType{
		{
			Name: "BooleanExpression",
			Properties: map[string]*model.ResourceProperty{
				"and": {
					Type: model.ResourceProperty_LIST,
				},
				"or": {
					Type: model.ResourceProperty_LIST,
				},
				"greaterThanOrEqual": {
					Type: model.ResourceProperty_STRUCT,
				},
				"regexMatch": {
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
				"greaterThan": {
					Type: model.ResourceProperty_STRUCT,
				},
				"lessThanOrEqual": {
					Type: model.ResourceProperty_STRUCT,
				},
				"in": {
					Type: model.ResourceProperty_STRUCT,
				},
			},

			Annotations: map[string]string{
				"EnableAudit":  "true",
				"OpenApiGroup": "internal",
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
				"OpenApiGroup": "internal",
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
				"OpenApiGroup": "internal",
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
				"OpenApiGroup": "internal",
				"EnableAudit":  "true",
			},
		},
		{
			Name:        "AuditData",
			Title:       "Audit Data",
			Description: "Audit Data is a type that represents the audit data of a resource/record. ",
			Properties: map[string]*model.ResourceProperty{
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
			},

			Annotations: map[string]string{
				"EnableAudit":  "true",
				"OpenApiGroup": "internal",
			},
		},
		{
			Name: "FunctionCall",
			Properties: map[string]*model.ResourceProperty{
				"host": {
					Type:     model.ResourceProperty_STRING,
					Required: true,
				},
				"functionName": {
					Type:     model.ResourceProperty_STRING,
					Required: true,
				},
			},

			Annotations: map[string]string{
				"EnableAudit":  "true",
				"OpenApiGroup": "internal",
			},
		},
		{
			Name: "HttpCall",
			Properties: map[string]*model.ResourceProperty{
				"uri": {
					Type:     model.ResourceProperty_STRING,
					Required: true,
				},
				"method": {
					Type:     model.ResourceProperty_STRING,
					Required: true,
				},
			},

			Annotations: map[string]string{
				"EnableAudit":  "true",
				"OpenApiGroup": "internal",
			},
		},
		{
			Name: "ChannelCall",
			Properties: map[string]*model.ResourceProperty{
				"channelKey": {
					Type:     model.ResourceProperty_STRING,
					Required: true,
				},
			},

			Annotations: map[string]string{
				"EnableAudit":  "true",
				"OpenApiGroup": "internal",
			},
		},
		{
			Name: "ExternalCall",
			Properties: map[string]*model.ResourceProperty{
				"functionCall": {
					Type: model.ResourceProperty_STRUCT,
				},
				"httpCall": {
					Type: model.ResourceProperty_STRUCT,
				},
				"channelCall": {
					Type: model.ResourceProperty_STRUCT,
				},
			},

			Annotations: map[string]string{
				"EnableAudit":  "true",
				"OpenApiGroup": "internal",
			},
		},
		{
			Name: "EventSelector",
			Properties: map[string]*model.ResourceProperty{
				"resources": {
					Type: model.ResourceProperty_LIST,
				},
				"ids": {
					Type: model.ResourceProperty_LIST,
				},
				"annotations": {
					Type:         model.ResourceProperty_MAP,
					ExampleValue: structpb.NewStructValue(&structpb.Struct{Fields: map[string]*structpb.Value{"CheckVersion": structpb.NewStringValue("true"), "IgnoreIfExists": structpb.NewStringValue("true"), "CommonType": structpb.NewStringValue("testType")}}),

					Annotations: map[string]string{
						"SpecialProperty": "true",
					},
				},
				"actions": {
					Type: model.ResourceProperty_LIST,
				},
				"recordSelector": {
					Type: model.ResourceProperty_STRUCT,
				},
				"namespaces": {
					Type: model.ResourceProperty_LIST,
				},
			},

			Annotations: map[string]string{
				"EnableAudit":  "true",
				"OpenApiGroup": "internal",
			},
		},
		{
			Name: "RecordSearchParams",
			Properties: map[string]*model.ResourceProperty{
				"query": {
					Type: model.ResourceProperty_STRUCT,
				},
				"limit": {
					Type: model.ResourceProperty_INT32,
				},
				"offset": {
					Type: model.ResourceProperty_INT32,
				},
				"resolveReferences": {
					Type: model.ResourceProperty_LIST,
				},
			},

			Annotations: map[string]string{
				"EnableAudit":  "true",
				"OpenApiGroup": "internal",
			},
		},
		{
			Name: "Event",
			Properties: map[string]*model.ResourceProperty{
				"time": {
					Type: model.ResourceProperty_TIMESTAMP,
				},
				"actionName": {
					Type: model.ResourceProperty_STRING,
				},
				"resource": {
					Type: model.ResourceProperty_REFERENCE,
				},
				"records": {
					Type: model.ResourceProperty_LIST,
				},
				"finalizes": {
					Type: model.ResourceProperty_BOOL,
				},
				"total": {
					Type: model.ResourceProperty_INT64,
				},
				"id": {
					Type:      model.ResourceProperty_STRING,
					Required:  true,
					Immutable: true,
				},
				"action": {
					Type:     model.ResourceProperty_ENUM,
					Required: true,

					Annotations: map[string]string{
						"TypeName": "ExtensionAction",
					},
				},
				"recordSearchParams": {
					Type: model.ResourceProperty_STRUCT,
				},
				"actionSummary": {
					Type: model.ResourceProperty_STRING,
				},
				"error": {
					Type: model.ResourceProperty_STRUCT,
				},
				"annotations": {
					Type:         model.ResourceProperty_MAP,
					ExampleValue: structpb.NewStructValue(&structpb.Struct{Fields: map[string]*structpb.Value{"CheckVersion": structpb.NewStringValue("true"), "IgnoreIfExists": structpb.NewStringValue("true"), "CommonType": structpb.NewStringValue("testType")}}),

					Annotations: map[string]string{
						"SpecialProperty": "true",
					},
				},
				"actionDescription": {
					Type: model.ResourceProperty_STRING,
				},
				"sync": {
					Type: model.ResourceProperty_BOOL,
				},
				"input": {
					Type: model.ResourceProperty_OBJECT,
				},
				"output": {
					Type: model.ResourceProperty_OBJECT,
				},
			},

			Annotations: map[string]string{
				"EnableAudit":  "true",
				"OpenApiGroup": "internal",
			},
		},
		{
			Name: "ErrorField",
			Properties: map[string]*model.ResourceProperty{
				"message": {
					Type: model.ResourceProperty_STRING,
				},
				"value": {
					Type: model.ResourceProperty_OBJECT,
				},
				"recordId": {
					Type: model.ResourceProperty_STRING,
				},
				"property": {
					Type: model.ResourceProperty_STRING,
				},
			},

			Annotations: map[string]string{
				"EnableAudit":  "true",
				"OpenApiGroup": "internal",
			},
		},
		{
			Name: "Error",
			Properties: map[string]*model.ResourceProperty{
				"code": {
					Type: model.ResourceProperty_ENUM,

					Annotations: map[string]string{
						"TypeName": "ExtensionCode",
					},
				},
				"message": {
					Type: model.ResourceProperty_STRING,
				},
				"fields": {
					Type: model.ResourceProperty_LIST,
				},
			},

			Annotations: map[string]string{
				"EnableAudit":  "true",
				"OpenApiGroup": "internal",
			},
		},
	},
	Properties: map[string]*model.ResourceProperty{
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
		"name": {
			Type:     model.ResourceProperty_STRING,
			Length:   256,
			Required: true,
			Unique:   true,

			Annotations: map[string]string{
				"IsHclLabel": "true",
			},
		},
		"selector": {
			Type: model.ResourceProperty_STRUCT,
		},
		"finalizes": {
			Type:     model.ResourceProperty_BOOL,
			Required: true,
		},
		"sync": {
			Type:     model.ResourceProperty_BOOL,
			Required: true,
		},
		"responds": {
			Type:     model.ResourceProperty_BOOL,
			Required: true,
		},
		"call": {
			Type:     model.ResourceProperty_STRUCT,
			Required: true,
		},
		"annotations": {
			Type:         model.ResourceProperty_MAP,
			ExampleValue: structpb.NewStructValue(&structpb.Struct{Fields: map[string]*structpb.Value{"IgnoreIfExists": structpb.NewStringValue("true"), "CommonType": structpb.NewStringValue("testType"), "CheckVersion": structpb.NewStringValue("true")}}),

			Annotations: map[string]string{
				"SpecialProperty": "true",
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
		"auditData": {
			Type:         model.ResourceProperty_STRUCT,
			ExampleValue: structpb.NewStructValue(&structpb.Struct{Fields: map[string]*structpb.Value{"createdBy": structpb.NewStringValue("admin"), "updatedBy": structpb.NewStringValue("admin"), "createdOn": structpb.NewStringValue("2023-11-20T23:38:43+04:00"), "updatedOn": structpb.NewStringValue("2023-11-20T23:38:43+04:00")}}),

			Annotations: map[string]string{
				"SpecialProperty": "true",
			},
		},
		"description": {
			Type:   model.ResourceProperty_STRING,
			Length: 1024,
		},
		"order": {
			Type:     model.ResourceProperty_INT32,
			Required: true,
		},
	},

	Annotations: map[string]string{
		"EnableAudit":  "true",
		"OpenApiGroup": "internal",
	},
}