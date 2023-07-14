package resources

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resources/special"
	"github.com/apibrew/apibrew/pkg/service/annotations"
)

var DataSourceResource = &model.Resource{
	Name:      "DataSource",
	Namespace: "system",
	SourceConfig: &model.ResourceSourceConfig{
		DataSource: "system",
		Entity:     "data_source",
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
			Unique:   true,
			Type:     model.ResourceProperty_STRING,
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
			Name:     "backend",
			Mapping:  "backend",
			Primary:  false,
			Type:     model.ResourceProperty_ENUM,
			Required: true,
			EnumValues: []string{
				"POSTGRESQL",
				"MYSQL",
				"MONGODB",
				"REDIS",
			},
		},
		{
			Name:     "options",
			Mapping:  "options",
			Primary:  false,
			Required: true,
			Type:     model.ResourceProperty_MAP,
			Item: &model.ResourceProperty{
				Type: model.ResourceProperty_STRING,
			},
		},
	},
	SecurityConstraints: special.SecurityContextDisallowAll,
	Annotations: map[string]string{
		annotations.EnableAudit: "true",
	},
}
