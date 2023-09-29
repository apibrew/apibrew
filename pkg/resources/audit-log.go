package resources

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resources/special"
	"github.com/apibrew/apibrew/pkg/service/annotations"
)

var AuditLogResource = &model.Resource{
	Name:      "AuditLog",
	Namespace: "system",
	SourceConfig: &model.ResourceSourceConfig{
		DataSource: "system",
		Entity:     "audit_log",
	},
	Immutable: true,
	Properties: []*model.ResourceProperty{
		special.IdProperty,
		special.VersionProperty,
		{
			Name:     "namespace",
			Length:   256,
			Type:     model.ResourceProperty_STRING,
			Required: true,
		},
		{
			Name:     "resource",
			Length:   256,
			Type:     model.ResourceProperty_STRING,
			Required: true,
		},
		{
			Name:     "recordId",
			Length:   256,
			Type:     model.ResourceProperty_STRING,
			Required: true,
			Annotations: map[string]string{
				annotations.SourceDef: "record_id",
			},
		},
		{
			Name:     "time",
			Type:     model.ResourceProperty_TIMESTAMP,
			Required: true,
		},
		{
			Name:     "username",
			Type:     model.ResourceProperty_STRING,
			Required: true,
		},
		{
			Name:     "operation",
			Type:     model.ResourceProperty_ENUM,
			Required: true,
			EnumValues: []string{
				"CREATE",
				"UPDATE",
				"DELETE",
			},
		},
		special.AnnotationsProperty,
	},
	Annotations: map[string]string{
		annotations.OpenApiGroup:     OpenApiInternal,
		annotations.BypassExtensions: annotations.Enabled,
	},
}
