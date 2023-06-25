package resources

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resources/special"
	"github.com/apibrew/apibrew/pkg/service/annotations"
)

var ResourceResource = &model.Resource{
	Name:      "resource",
	Namespace: "system",
	SourceConfig: &model.ResourceSourceConfig{
		DataSource: "system",
		Entity:     "resource",
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
			Primary:  false,
			Type:     model.ResourceProperty_STRING,
			Length:   256,
			Required: true,
			Unique:   false,
			Annotations: map[string]string{
				annotations.IsHclLabel: annotations.Enabled,
			},
		},
		{
			Name:     "namespace",
			Mapping:  "namespace",
			Type:     model.ResourceProperty_REFERENCE,
			Required: true,
			Reference: &model.Reference{
				Resource:  NamespaceResource.Name,
				Namespace: NamespaceResource.Namespace,
				Cascade:   false,
			},
		},
		{
			Name:     "virtual",
			Mapping:  "virtual",
			Primary:  false,
			Type:     model.ResourceProperty_BOOL,
			Required: true,
		},
		{
			Name:     "types",
			Mapping:  "types",
			Primary:  false,
			Type:     model.ResourceProperty_OBJECT,
			Required: false,
		},
		{
			Name:     "immutable",
			Mapping:  "immutable",
			Primary:  false,
			Type:     model.ResourceProperty_BOOL,
			Required: true,
		},
		{
			Name:     "abstract",
			Mapping:  "abstract",
			Primary:  false,
			Type:     model.ResourceProperty_BOOL,
			Required: true,
		},
		{
			Name:     "dataSource",
			Mapping:  "source_data_source",
			Type:     model.ResourceProperty_REFERENCE,
			Required: false,
			Reference: &model.Reference{
				Resource:  DataSourceResource.Name,
				Namespace: DataSourceResource.Namespace,
				Cascade:   false,
			},
		},
		{
			Name:     "entity",
			Mapping:  "mapping",
			Type:     model.ResourceProperty_STRING,
			Length:   256,
			Required: false,
		},
		{
			Name:     "catalog",
			Mapping:  "source_catalog",
			Type:     model.ResourceProperty_STRING,
			Length:   256,
			Required: false,
		},
		{
			Name:     "annotations",
			Mapping:  "annotations",
			Type:     model.ResourceProperty_OBJECT,
			Required: false,
		},
		{
			Name:     "indexes",
			Mapping:  "indexes",
			Type:     model.ResourceProperty_OBJECT,
			Required: false,
		},
		special.SecurityConstraintsProperty,
		{
			Name:     "title",
			Mapping:  "title",
			Primary:  false,
			Type:     model.ResourceProperty_STRING,
			Length:   256,
			Required: false,
		},
		{
			Name:     "description",
			Mapping:  "description",
			Primary:  false,
			Type:     model.ResourceProperty_STRING,
			Length:   256,
			Required: false,
		},
	},
	Indexes: []*model.ResourceIndex{
		{
			Properties: []*model.ResourceIndexProperty{
				{
					Name: "namespace",
				},
				{
					Name: "name",
				},
			},
			Unique: true,
		},
	},
	SecurityConstraints: special.SecurityContextDisallowAll,
	Annotations: map[string]string{
		annotations.EnableAudit: "true",
	},
}
