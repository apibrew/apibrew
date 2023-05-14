package resources

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/service/annotations"
)

var NamespaceResource = &model.Resource{
	Name:      "namespace",
	Namespace: "system",
	SourceConfig: &model.ResourceSourceConfig{
		DataSource: "system",
		Entity:     "namespace",
	},
	Properties: []*model.ResourceProperty{
		IdProperty,
		VersionProperty,
		AuditProperties[0],
		AuditProperties[1],
		AuditProperties[2],
		AuditProperties[3],
		{
			Name:      "name",
			Mapping:   "name",
			Primary:   false,
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
			Name: "description",

			Mapping:  "description",
			Primary:  false,
			Type:     model.ResourceProperty_STRING,
			Length:   256,
			Required: false,
		},
		{
			Name: "details",

			Mapping:  "details",
			Primary:  false,
			Type:     model.ResourceProperty_OBJECT,
			Required: false,
		},
		securityContextProperty,
	},
	SecurityContext: securityContextDisallowAll,
	Annotations: map[string]string{
		annotations.EnableAudit: "true",
	},
}
