package system

import "github.com/tislib/data-handler/model"

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
			SourceConfig: &model.ResourceProperty_Mapping{
				Mapping: &model.ResourcePropertyMappingConfig{
					Mapping:        "name",
					SourceDef:      "",
					AutoGeneration: 0,
				},
			},
			Primary:  false,
			Type:     model.ResourcePropertyType_TYPE_STRING,
			Length:   256,
			Required: true,
			Unique:   true,
		},
		{
			Name: "namespace",
			SourceConfig: &model.ResourceProperty_Mapping{
				Mapping: &model.ResourcePropertyMappingConfig{
					Mapping: "namespace",
				},
			},
			Type:     model.ResourcePropertyType_TYPE_STRING,
			Length:   256,
			Required: true,
		},
		{
			Name: "dataSource",
			SourceConfig: &model.ResourceProperty_Mapping{
				Mapping: &model.ResourcePropertyMappingConfig{
					Mapping: "source_data_source",
				},
			},
			Type:     model.ResourcePropertyType_TYPE_UUID,
			Length:   256,
			Required: true,
		},
		{
			Name: "entity",
			SourceConfig: &model.ResourceProperty_Mapping{
				Mapping: &model.ResourcePropertyMappingConfig{
					Mapping: "source_mapping",
				},
			},
			Type:     model.ResourcePropertyType_TYPE_STRING,
			Length:   256,
			Required: true,
		},
		{
			Name: "catalog",
			SourceConfig: &model.ResourceProperty_Mapping{
				Mapping: &model.ResourcePropertyMappingConfig{
					Mapping: "source_catalog",
				},
			},
			Type:     model.ResourcePropertyType_TYPE_STRING,
			Length:   256,
			Required: true,
		},
		{
			Name: "annotations",
			SourceConfig: &model.ResourceProperty_Mapping{
				Mapping: &model.ResourcePropertyMappingConfig{
					Mapping: "annotations",
				},
			},
			Type:     model.ResourcePropertyType_TYPE_OBJECT,
			Required: false,
		},
		{
			Name: "type",
			SourceConfig: &model.ResourceProperty_Mapping{
				Mapping: &model.ResourcePropertyMappingConfig{
					Mapping: "type",
				},
			},
			Type:     model.ResourcePropertyType_TYPE_INT32,
			Required: true,
		},
		securityContextProperty,
	},
	References: []*model.ResourceReference{
		{
			PropertyName:       "dataSource",
			ReferencedResource: DataSourceResource.Name,
			Cascade:            false,
		},
	},
	SecurityContext: securityContextDisallowAll,
}
