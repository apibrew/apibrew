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
		sub_types.RefValue,
		sub_types.RegexMatchExpression,
		sub_types.Expression,
		{
			Name: "FunctionCall",
			Annotations: map[string]string{
				annotations.CommonType: annotations.Enabled,
			},
			Properties: []*model.ResourceProperty{
				{
					Name:     "host",
					Type:     model.ResourceProperty_STRING,
					Required: true,
				},
				{
					Name:     "functionName",
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
			Properties: []*model.ResourceProperty{
				{
					Name:     "uri",
					Type:     model.ResourceProperty_STRING,
					Required: true,
				},
				{
					Name:     "method",
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
			Properties: []*model.ResourceProperty{
				{
					Name:     "channelKey",
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
			Properties: []*model.ResourceProperty{
				{
					Name:    "functionCall",
					Type:    model.ResourceProperty_STRUCT,
					TypeRef: util.Pointer("FunctionCall"),
				},
				{
					Name:    "httpCall",
					Type:    model.ResourceProperty_STRUCT,
					TypeRef: util.Pointer("HttpCall"),
				},
				{
					Name:    "channelCall",
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
			Properties: []*model.ResourceProperty{
				{
					Name: "actions",
					Type: model.ResourceProperty_LIST,
					Item: &model.ResourceProperty{
						Name: "action",
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
				{
					Name:    "recordSelector",
					Type:    model.ResourceProperty_STRUCT,
					TypeRef: util.Pointer(sub_types.BooleanExpression.Name),
				},
				{
					Name: "namespaces",
					Type: model.ResourceProperty_LIST,
					Item: &model.ResourceProperty{Type: model.ResourceProperty_STRING},
				},
				{
					Name: "resources",
					Type: model.ResourceProperty_LIST,
					Item: &model.ResourceProperty{Type: model.ResourceProperty_STRING},
				},
				{
					Name: "ids",
					Type: model.ResourceProperty_LIST,
					Item: &model.ResourceProperty{Type: model.ResourceProperty_STRING},
				},
				special.AnnotationsProperty,
			},
		},
		{
			Name: "RecordSearchParams",
			Annotations: map[string]string{
				annotations.CommonType: annotations.Enabled,
			},
			Properties: []*model.ResourceProperty{
				{
					Name:    "query",
					Type:    model.ResourceProperty_STRUCT,
					TypeRef: &sub_types.BooleanExpression.Name,
				},
				{
					Name: "limit",
					Type: model.ResourceProperty_INT32,
				},
				{
					Name: "offset",
					Type: model.ResourceProperty_INT32,
				},
				{
					Name: "resolveReferences",
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
			Properties: []*model.ResourceProperty{
				{
					Name:      "id",
					Type:      model.ResourceProperty_STRING,
					Required:  true,
					Immutable: true,
				},
				{
					Name: "action",
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
				{
					Name:    "recordSearchParams",
					Type:    model.ResourceProperty_STRUCT,
					TypeRef: util.Pointer("RecordSearchParams"),
				},
				{
					Name: "actionSummary",
					Type: model.ResourceProperty_STRING,
				},
				{
					Name: "actionDescription",
					Type: model.ResourceProperty_STRING,
				},
				{
					Name: "resource",
					Type: model.ResourceProperty_REFERENCE,
					Reference: &model.Reference{
						Resource: "Resource",
					},
				},
				{
					Name: "records",
					Type: model.ResourceProperty_LIST,
					Item: &model.ResourceProperty{
						Type: model.ResourceProperty_REFERENCE,
						Reference: &model.Reference{
							Resource: "Record",
						},
					},
				},
				{
					Name: "ids",
					Type: model.ResourceProperty_LIST,
					Item: &model.ResourceProperty{
						Type: model.ResourceProperty_STRING,
					},
				},
				{
					Name: "finalizes",
					Type: model.ResourceProperty_BOOL,
				},
				{
					Name: "sync",
					Type: model.ResourceProperty_BOOL,
				},
				{
					Name: "time",
					Type: model.ResourceProperty_TIMESTAMP,
				},
				{
					Name: "total",
					Type: model.ResourceProperty_INT64,
				},
				{
					Name: "actionName",
					Type: model.ResourceProperty_STRING,
				},
				{
					Name: "input",
					Type: model.ResourceProperty_OBJECT,
				},
				{
					Name: "output",
					Type: model.ResourceProperty_OBJECT,
				},
				special.AnnotationsProperty,
				{
					Name:    "error",
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
			Properties: []*model.ResourceProperty{
				{
					Name: "recordId",
					Type: model.ResourceProperty_STRING,
				},
				{
					Name: "property",
					Type: model.ResourceProperty_STRING,
				},
				{
					Name: "message",
					Type: model.ResourceProperty_STRING,
				},
				{
					Name: "value",
					Type: model.ResourceProperty_OBJECT,
				},
			},
		},
		{
			Name: "Error",
			Annotations: map[string]string{
				annotations.CommonType: annotations.Enabled,
			},
			Properties: []*model.ResourceProperty{
				{
					Name: "code",
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
				{
					Name: "message",
					Type: model.ResourceProperty_STRING,
				},
				{
					Name: "fields",
					Type: model.ResourceProperty_LIST,
					Item: &model.ResourceProperty{
						Type:    model.ResourceProperty_STRUCT,
						TypeRef: util.Pointer("ErrorField"),
					},
				},
			},
		},
	},
	Properties: []*model.ResourceProperty{
		special.IdProperty,
		special.VersionProperty,
		special.AuditProperties[0],
		special.AuditProperties[1],
		special.AuditProperties[2],
		special.AuditProperties[3],
		{
			Name:     "name",
			Length:   64,
			Type:     model.ResourceProperty_STRING,
			Unique:   true,
			Required: true,
			Annotations: map[string]string{
				annotations.IsHclLabel: annotations.Enabled,
			},
		},
		{
			Name:     "description",
			Length:   64,
			Type:     model.ResourceProperty_STRING,
			Required: false,
		},
		{
			Name:     "selector",
			Type:     model.ResourceProperty_STRUCT,
			Required: false,
			TypeRef:  util.Pointer("EventSelector"),
		},
		{
			Name:     "order",
			Type:     model.ResourceProperty_INT32,
			Required: true,
		},
		{
			Name:     "finalizes",
			Type:     model.ResourceProperty_BOOL,
			Required: true,
		},
		{
			Name:     "sync",
			Type:     model.ResourceProperty_BOOL,
			Required: true,
		},
		{
			Name:     "responds",
			Type:     model.ResourceProperty_BOOL,
			Required: true,
		},
		{
			Name:     "call",
			Type:     model.ResourceProperty_STRUCT,
			Required: true,
			TypeRef:  util.Pointer("ExternalCall"),
		},
		special.AnnotationsProperty,
	},
	Annotations: map[string]string{
		annotations.EnableAudit:  "true",
		annotations.OpenApiGroup: OpenApiInternal,
	},
}
