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
	Properties: []*model.ResourceProperty{
		special.IdProperty,
		special.VersionProperty,
		special.AuditProperty,
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
