package resources

import (
	"github.com/tislib/data-handler/pkg/model"
)

var ResourceResource = &model.Resource{
	Name:      "resource",
	Namespace: "system",
	DataType:  model.DataType_SYSTEM,
	SourceConfig: &model.ResourceSourceConfig{
		DataSource: "system",
		Entity:     "resource",
	},
	Properties: []*model.ResourceProperty{
		{
			Name: "name",

			Mapping:  "name",
			Primary:  false,
			Type:     model.ResourcePropertyType_TYPE_STRING,
			Length:   256,
			Required: true,
			Unique:   true,
		},
		{
			Name: "virtual",

			Mapping:  "virtual",
			Primary:  false,
			Type:     model.ResourcePropertyType_TYPE_BOOL,
			Length:   256,
			Required: false,
		},
		{
			Name: "namespace",

			Mapping:  "namespace",
			Type:     model.ResourcePropertyType_TYPE_STRING,
			Length:   256,
			Required: true,
		},
		{
			Name: "dataSource",

			Mapping:  "source_data_source",
			Type:     model.ResourcePropertyType_TYPE_REFERENCE,
			Length:   256,
			Required: false,
			Reference: &model.Reference{
				ReferencedResource: DataSourceResource.Name,
				Cascade:            false,
			},
		},
		{
			Name: "entity",

			Mapping:  "source_mapping",
			Type:     model.ResourcePropertyType_TYPE_STRING,
			Length:   256,
			Required: false,
		},
		{
			Name: "catalog",

			Mapping:  "source_catalog",
			Type:     model.ResourcePropertyType_TYPE_STRING,
			Length:   256,
			Required: false,
		},
		{
			Name: "annotations",

			Mapping:  "annotations",
			Type:     model.ResourcePropertyType_TYPE_OBJECT,
			Required: false,
		},
		{
			Name: "type",

			Mapping:  "type",
			Type:     model.ResourcePropertyType_TYPE_INT32,
			Required: true,
		},
		securityContextProperty,
	},
	SecurityContext: securityContextDisallowAll,
}
