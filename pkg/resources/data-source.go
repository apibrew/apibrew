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
	Types: []*model.ResourceSubType{
		special.AuditDataSubType,
	},
	Properties: []*model.ResourceProperty{
		special.IdProperty,
		special.VersionProperty,
		special.AuditProperty,
		{
			Name:     "name",
			Length:   64,
			Unique:   true,
			Type:     model.ResourceProperty_STRING,
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
			Annotations: map[string]string{
				annotations.AllowEmptyPrimitive: annotations.Enabled,
			},
		},
		{
			Name:     "backend",
			Type:     model.ResourceProperty_STRING,
			Required: true,
		},
		{
			Name:     "options",
			Required: true,
			Type:     model.ResourceProperty_MAP,
			Item: &model.ResourceProperty{
				Type: model.ResourceProperty_STRING,
			},
		},
	},
	Annotations: map[string]string{
		annotations.EnableAudit:  "true",
		annotations.OpenApiGroup: OpenApiInternal,
	},
}
