package resources

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resources/special"
	"github.com/apibrew/apibrew/pkg/service/annotations"
)

var NamespaceResource = &model.Resource{
	Name:      "Namespace",
	Namespace: "system",
	SourceConfig: &model.ResourceSourceConfig{
		DataSource: "system",
		Entity:     "namespace",
	},
	Types: []*model.ResourceSubType{
		special.AuditDataSubType,
	},
	Properties: map[string]*model.ResourceProperty{
		"id":        special.IdProperty,
		"version":   special.VersionProperty,
		"auditData": special.AuditProperty,
		"name": {
			Type:      model.ResourceProperty_STRING,
			Length:    256,
			Required:  true,
			Unique:    true,
			Immutable: true,
			Annotations: map[string]string{
				annotations.IsHclLabel: annotations.Enabled,
			},
		},
		"description": {
			Type:     model.ResourceProperty_STRING,
			Length:   256,
			Required: false,
		},
		"details": {
			Type:     model.ResourceProperty_OBJECT,
			Required: false,
		},
	},
	Annotations: map[string]string{
		annotations.EnableAudit:  "true",
		annotations.OpenApiGroup: OpenApiInternal,
	},
}
