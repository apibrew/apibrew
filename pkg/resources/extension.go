package resources

import (
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/service/annotations"
)

var ExtensionResource = &model.Resource{
	Name:      "extension",
	Namespace: "system",
	SourceConfig: &model.ResourceSourceConfig{
		DataSource: "system",
		Entity:     "extension",
	},
	Properties: []*model.ResourceProperty{
		IdProperty,
		VersionProperty,
		AuditProperties[0],
		AuditProperties[1],
		AuditProperties[2],
		AuditProperties[3],
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
			Name:     "namespace",
			Mapping:  "namespace",
			Length:   64,
			Primary:  false,
			Type:     model.ResourceProperty_STRING,
			Required: true,
		},
		{
			Name:     "resource",
			Mapping:  "resource",
			Length:   64,
			Primary:  false,
			Type:     model.ResourceProperty_STRING,
			Required: true,
		},
		{
			Name:    "before",
			Mapping: "before",
			Type:    model.ResourceProperty_OBJECT,
		},
		{
			Name:    "after",
			Mapping: "after",
			Type:    model.ResourceProperty_OBJECT,
		},
		{
			Name:    "instead",
			Mapping: "instead",
			Type:    model.ResourceProperty_OBJECT,
		},
	},
	SecurityContext: securityContextDisallowAll,
}
