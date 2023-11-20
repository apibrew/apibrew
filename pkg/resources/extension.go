package resources

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resources/special"
	sub_types "github.com/apibrew/apibrew/pkg/resources/sub-types"
	"github.com/apibrew/apibrew/pkg/service/annotations"
	"github.com/apibrew/apibrew/pkg/util"
)

var ExtensionResource = &model.Resource{
	Name:      "Extension",
	Namespace: "system",
	SourceConfig: &model.ResourceSourceConfig{
		DataSource: "system",
		Entity:     "extension",
	},
	Types: []*model.ResourceSubType{
		sub_types.BooleanExpression,
		sub_types.PairExpression,
		sub_types.RegexMatchExpression,
		sub_types.Expression,
		special.AuditDataSubType,
		{
			Name: "FunctionCall",
			Annotations: map[string]string{
				annotations.CommonType: annotations.Enabled,
			},
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
		},
		{
			Name: "HttpCall",
			Annotations: map[string]string{
				annotations.CommonType: annotations.Enabled,
			},
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
		},
		{
			Name: "ChannelCall",
			Annotations: map[string]string{
				annotations.CommonType: annotations.Enabled,
			},
			Properties: map[string]*model.ResourceProperty{
				"channelKey": {
					Type:     model.ResourceProperty_STRING,
					Required: true,
				},
			},
		},
		{
			Name: "ExternalCall",
			Annotations: map[string]string{
				annotations.CommonType: annotations.Enabled,
			},
			Properties: map[string]*model.ResourceProperty{
				"functionCall": {
					Type:    model.ResourceProperty_STRUCT,
					TypeRef: util.Pointer("FunctionCall"),
				},
				"httpCall": {
					Type:    model.ResourceProperty_STRUCT,
					TypeRef: util.Pointer("HttpCall"),
				},
				"channelCall": {
					Type:    model.ResourceProperty_STRUCT,
					TypeRef: util.Pointer("ChannelCall"),
				},
			},
		},
		{
			Name: "EventSelector",
			Annotations: map[string]string{
				annotations.CommonType: annotations.Enabled,
			},
			Properties: map[string]*model.ResourceProperty{
				"actions": {
					Type: model.ResourceProperty_LIST,
					Item: &model.ResourceProperty{
						Type: model.ResourceProperty_ENUM,
						EnumValues: []string{
							"CREATE",
							"UPDATE",
							"DELETE",
							"GET",
							"LIST",
							"OPERATE",
						},
						Annotations: map[string]string{
							annotations.TypeName: "EventAction",
						},
					},
				},
				"recordSelector": {
					Type:    model.ResourceProperty_STRUCT,
					TypeRef: util.Pointer(sub_types.BooleanExpression.Name),
				},
				"namespaces": {
					Type: model.ResourceProperty_LIST,
					Item: &model.ResourceProperty{Type: model.ResourceProperty_STRING},
				},
				"resources": {
					Type: model.ResourceProperty_LIST,
					Item: &model.ResourceProperty{Type: model.ResourceProperty_STRING},
				},
				"ids": {
					Type: model.ResourceProperty_LIST,
					Item: &model.ResourceProperty{Type: model.ResourceProperty_STRING},
				},
				"annotations": special.AnnotationsProperty,
			},
		},
		{
			Name: "RecordSearchParams",
			Annotations: map[string]string{
				annotations.CommonType: annotations.Enabled,
			},
			Properties: map[string]*model.ResourceProperty{
				"query": {
					Type:    model.ResourceProperty_STRUCT,
					TypeRef: &sub_types.BooleanExpression.Name,
				},
				"limit": {
					Type: model.ResourceProperty_INT32,
				},
				"offset": {
					Type: model.ResourceProperty_INT32,
				},
				"resolveReferences": {
					Type: model.ResourceProperty_LIST,
					Item: &model.ResourceProperty{
						Type: model.ResourceProperty_STRING,
					},
				},
			},
		},
		{
			Name: "Event",
			Annotations: map[string]string{
				annotations.CommonType: annotations.Enabled,
			},
			Properties: map[string]*model.ResourceProperty{
				"id": {
					Type:      model.ResourceProperty_STRING,
					Required:  true,
					Immutable: true,
				},
				"action": {
					Type: model.ResourceProperty_ENUM,
					EnumValues: []string{
						"CREATE",
						"UPDATE",
						"DELETE",
						"GET",
						"LIST",
						"OPERATE",
					},
					Required: true,
				},
				"recordSearchParams": {
					Type:    model.ResourceProperty_STRUCT,
					TypeRef: util.Pointer("RecordSearchParams"),
				},
				"actionSummary": {
					Type: model.ResourceProperty_STRING,
				},
				"actionDescription": {
					Type: model.ResourceProperty_STRING,
				},
				"resource": {
					Type: model.ResourceProperty_REFERENCE,
					Reference: &model.Reference{
						Resource: "Resource",
					},
				},
				"records": {
					Type: model.ResourceProperty_LIST,
					Item: &model.ResourceProperty{
						Type: model.ResourceProperty_REFERENCE,
						Reference: &model.Reference{
							Resource: "Record",
						},
					},
				},
				"finalizes": {
					Type: model.ResourceProperty_BOOL,
				},
				"sync": {
					Type: model.ResourceProperty_BOOL,
				},
				"time": {
					Type: model.ResourceProperty_TIMESTAMP,
				},
				"total": {
					Type: model.ResourceProperty_INT64,
				},
				"actionName": {
					Type: model.ResourceProperty_STRING,
				},
				"input": {
					Type: model.ResourceProperty_OBJECT,
				},
				"output": {
					Type: model.ResourceProperty_OBJECT,
				},
				"annotations": special.AnnotationsProperty,
				"error": {
					Type:    model.ResourceProperty_STRUCT,
					TypeRef: util.Pointer("Error"),
				},
			},
		},
		{
			Name: "ErrorField",
			Annotations: map[string]string{
				annotations.CommonType: annotations.Enabled,
			},
			Properties: map[string]*model.ResourceProperty{
				"recordId": {
					Type: model.ResourceProperty_STRING,
				},
				"property": {
					Type: model.ResourceProperty_STRING,
				},
				"message": {
					Type: model.ResourceProperty_STRING,
				},
				"value": {
					Type: model.ResourceProperty_OBJECT,
				},
			},
		},
		{
			Name: "Error",
			Annotations: map[string]string{
				annotations.CommonType: annotations.Enabled,
			},
			Properties: map[string]*model.ResourceProperty{
				"code": {
					Type: model.ResourceProperty_ENUM,
					EnumValues: []string{
						"UNKNOWN_ERROR",
						"RECORD_NOT_FOUND",
						"UNABLE_TO_LOCATE_PRIMARY_KEY",
						"INTERNAL_ERROR",
						"PROPERTY_NOT_FOUND",
						"RECORD_VALIDATION_ERROR",
						"RESOURCE_VALIDATION_ERROR",
						"AUTHENTICATION_FAILED",
						"ALREADY_EXISTS",
						"ACCESS_DENIED",
						"BACKEND_ERROR",
						"UNIQUE_VIOLATION",
						"REFERENCE_VIOLATION",
						"RESOURCE_NOT_FOUND",
						"UNSUPPORTED_OPERATION",
						"EXTERNAL_BACKEND_COMMUNICATION_ERROR",
						"EXTERNAL_BACKEND_ERROR",
						"RATE_LIMIT_ERROR",
					},
				},
				"message": {
					Type: model.ResourceProperty_STRING,
				},
				"fields": {
					Type: model.ResourceProperty_LIST,
					Item: &model.ResourceProperty{
						Type:    model.ResourceProperty_STRUCT,
						TypeRef: util.Pointer("ErrorField"),
					},
				},
			},
		},
	},
	Properties: map[string]*model.ResourceProperty{
		"id":        special.IdProperty,
		"version":   special.VersionProperty,
		"auditData": special.AuditProperty,
		"name": {
			Length:   256,
			Type:     model.ResourceProperty_STRING,
			Unique:   true,
			Required: true,
			Annotations: map[string]string{
				annotations.IsHclLabel: annotations.Enabled,
			},
		},
		"description": {
			Length:   1024,
			Type:     model.ResourceProperty_STRING,
			Required: false,
		},
		"selector": {
			Type:     model.ResourceProperty_STRUCT,
			Required: false,
			TypeRef:  util.Pointer("EventSelector"),
		},
		"order": {
			Type:     model.ResourceProperty_INT32,
			Required: true,
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
			TypeRef:  util.Pointer("ExternalCall"),
		},
		"annotations": special.AnnotationsProperty,
	},
	Annotations: map[string]string{
		annotations.EnableAudit:  "true",
		annotations.OpenApiGroup: OpenApiInternal,
	},
}
