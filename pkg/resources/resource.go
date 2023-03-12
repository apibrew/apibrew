package resources

import (
	"github.com/tislib/data-handler/pkg/model"
)

var ResourceResource = &model.Resource{
	Name:      "resource",
	Namespace: "system",
	SourceConfig: &model.ResourceSourceConfig{
		DataSource: "system",
		Entity:     "resource",
	},
	Properties: []*model.ResourceProperty{
		{
			Name:     "name",
			Mapping:  "name",
			Primary:  false,
			Type:     model.ResourcePropertyType_TYPE_STRING,
			Length:   256,
			Required: true,
			Unique:   false,
		},
		{
			Name:     "namespace",
			Mapping:  "namespace",
			Type:     model.ResourcePropertyType_TYPE_REFERENCE,
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
			Type:     model.ResourcePropertyType_TYPE_BOOL,
			Required: true,
		},
		{
			Name:     "immutable",
			Mapping:  "immutable",
			Primary:  false,
			Type:     model.ResourcePropertyType_TYPE_BOOL,
			Required: true,
		},
		{
			Name:     "abstract",
			Mapping:  "abstract",
			Primary:  false,
			Type:     model.ResourcePropertyType_TYPE_BOOL,
			Required: true,
		},
		{
			Name:     "dataSource",
			Mapping:  "source_data_source",
			Type:     model.ResourcePropertyType_TYPE_REFERENCE,
			Required: false,
			Reference: &model.Reference{
				ReferencedResource: DataSourceResource.Name,
				Cascade:            false,
			},
		},
		{
			Name:     "entity",
			Mapping:  "source_mapping",
			Type:     model.ResourcePropertyType_TYPE_STRING,
			Length:   256,
			Required: false,
		},
		{
			Name:     "catalog",
			Mapping:  "source_catalog",
			Type:     model.ResourcePropertyType_TYPE_STRING,
			Length:   256,
			Required: false,
		},
		{
			Name:     "annotations",
			Mapping:  "annotations",
			Type:     model.ResourcePropertyType_TYPE_OBJECT,
			Required: false,
		},
		{
			Name:     "indexes",
			Mapping:  "indexes",
			Type:     model.ResourcePropertyType_TYPE_OBJECT,
			Required: false,
		},
		securityContextProperty,
		{
			Name:     "title",
			Mapping:  "title",
			Primary:  false,
			Type:     model.ResourcePropertyType_TYPE_STRING,
			Length:   256,
			Required: false,
		},
		{
			Name:     "description",
			Mapping:  "description",
			Primary:  false,
			Type:     model.ResourcePropertyType_TYPE_STRING,
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
