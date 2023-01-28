package system

import "data-handler/model"

var ResourceReferenceResource = &model.Resource{
	Name:      "resourceReference",
	Namespace: "system",
	DataType:  model.DataType_SYSTEM,
	SourceConfig: &model.ResourceSourceConfig{
		DataSource: "system",
		Entity:     "resource_reference",
	},
	Properties: []*model.ResourceProperty{
		{
			Name: "propertyName",
			SourceConfig: &model.ResourceProperty_Mapping{
				Mapping: &model.ResourcePropertyMappingConfig{
					Mapping:        "property_name",
					SourceDef:      "",
					AutoGeneration: 0,
				},
			},
			Primary:  false,
			Type:     model.ResourcePropertyType_TYPE_STRING,
			Length:   256,
			Required: true,
		},
		{
			Name: "referencedResource",
			SourceConfig: &model.ResourceProperty_Mapping{
				Mapping: &model.ResourcePropertyMappingConfig{
					Mapping:        "referenced_resource",
					SourceDef:      "",
					AutoGeneration: 0,
				},
			},
			Primary:  false,
			Type:     model.ResourcePropertyType_TYPE_STRING,
			Length:   256,
			Required: true,
		},
		{
			Name: "cascade",
			SourceConfig: &model.ResourceProperty_Mapping{
				Mapping: &model.ResourcePropertyMappingConfig{
					Mapping:        "cascade",
					SourceDef:      "",
					AutoGeneration: 0,
				},
			},
			Primary:  false,
			Type:     model.ResourcePropertyType_TYPE_BOOL,
			Required: true,
		},
		{
			Name: "resource",
			SourceConfig: &model.ResourceProperty_Mapping{
				Mapping: &model.ResourcePropertyMappingConfig{
					Mapping: "resource",
				},
			},
			Type:     model.ResourcePropertyType_TYPE_UUID,
			Required: true,
		},
	},
	References: []*model.ResourceReference{
		{
			PropertyName:       "resource",
			ReferencedResource: "resource",
			Cascade:            true,
		},
	},
}
