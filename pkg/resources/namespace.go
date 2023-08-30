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
	Properties: []*model.ResourceProperty{
		special.IdProperty,
		special.VersionProperty,
		special.AuditProperties[0],
		special.AuditProperties[1],
		special.AuditProperties[2],
		special.AuditProperties[3],
		{
			Name:      "name",
			Type:      model.ResourceProperty_STRING,
			Length:    256,
			Required:  true,
			Unique:    true,
			Immutable: true,
			Annotations: map[string]string{
				annotations.IsHclLabel: annotations.Enabled,
			},
		},
		{
			Name:     "description",
			Type:     model.ResourceProperty_STRING,
			Length:   256,
			Required: false,
		},
		{
			Name:     "details",
			Type:     model.ResourceProperty_OBJECT,
			Required: false,
		},
	},
	Annotations: map[string]string{
		annotations.EnableAudit:  "true",
		annotations.OpenApiGroup: OpenApiInternal,
	},
}
