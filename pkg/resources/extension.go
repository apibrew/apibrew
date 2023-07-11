package resources

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resources/special"
	sub_types "github.com/apibrew/apibrew/pkg/resources/sub-types"
	"github.com/apibrew/apibrew/pkg/service/annotations"
	"github.com/apibrew/apibrew/pkg/util"
)

var ExtensionResource = &model.Resource{
	Name:      "extension",
	Namespace: "system",
	SourceConfig: &model.ResourceSourceConfig{
		DataSource: "system",
		Entity:     "extension",
	},
	Types: []*model.ResourceSubType{
		{
			Name: "functionCall",
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
			Name: "httpCall",
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
			Name: "call",
			Properties: []*model.ResourceProperty{
				{
					Name:    "functionCall",
					Type:    model.ResourceProperty_STRUCT,
					TypeRef: util.StringPointer("functionCall"),
				},
				{
					Name:    "httpCall",
					Type:    model.ResourceProperty_STRUCT,
					TypeRef: util.StringPointer("httpCall"),
				},
			},
		},
		sub_types.BooleanExpression,
		{
			Name: "selector",
			Properties: []*model.ResourceProperty{
				{
					Name: "actions",
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
					},
				},
				{
					Name:    "recordSelector",
					Type:    model.ResourceProperty_STRUCT,
					TypeRef: util.StringPointer(sub_types.BooleanExpression.Name),
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
			Mapping:  "name",
			Length:   64,
			Primary:  false,
			Type:     model.ResourceProperty_STRING,
			Unique:   true,
			Required: true,
			Annotations: map[string]string{
				annotations.IsHclLabel: annotations.Enabled,
			},
		},
		{
			Name:     "description",
			Mapping:  "description",
			Length:   64,
			Primary:  false,
			Type:     model.ResourceProperty_STRING,
			Required: false,
		},
		{
			Name:     "selector",
			Mapping:  "selector",
			Type:     model.ResourceProperty_STRUCT,
			Required: false,
			TypeRef:  util.StringPointer("selector"),
		},
		{
			Name:     "order",
			Mapping:  "order",
			Primary:  false,
			Type:     model.ResourceProperty_INT32,
			Required: true,
		},
		{
			Name:     "finalizes",
			Mapping:  "finalizes",
			Primary:  false,
			Type:     model.ResourceProperty_BOOL,
			Required: true,
		},
		{
			Name:     "sync",
			Mapping:  "sync",
			Primary:  false,
			Type:     model.ResourceProperty_BOOL,
			Required: true,
		},
		{
			Name:     "responds",
			Mapping:  "responds",
			Primary:  false,
			Type:     model.ResourceProperty_BOOL,
			Required: true,
		},
		{
			Name:     "call",
			Mapping:  "call",
			Type:     model.ResourceProperty_STRUCT,
			Required: true,
			TypeRef:  util.StringPointer("call"),
		},
		special.AnnotationsProperty,
	},
	SecurityConstraints: special.SecurityContextDisallowAll,
	Annotations: map[string]string{
		annotations.EnableAudit: "true",
	},
}
