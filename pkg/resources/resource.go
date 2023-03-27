package resources

import (
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/service/annotations"
)

var ResourceResource = &model.Resource{
	Name:      "resource",
	Namespace: "system",
	SourceConfig: &model.ResourceSourceConfig{
		DataSource: "system",
		Entity:     "resource",
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
				ReferencedResource: NamespaceResource.Name,
				Cascade:            false,
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
				ReferencedResource: DataSourceResource.Name,
				Cascade:            false,
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
		securityContextProperty,
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
	SecurityContext: securityContextDisallowAll,
}
