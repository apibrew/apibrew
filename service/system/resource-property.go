package system

import "data-handler/model"

var ResourcePropertyResource = &model.Resource{
	Name:      "resourceProperty",
	Workspace: "system",
	DataType:  model.DataType_SYSTEM,
	SourceConfig: &model.ResourceSourceConfig{
		DataSource: "system",
		Mapping:    "resource_property",
	},
	Flags: &model.ResourceFlags{},
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
		{
			Name: "sourceType",
			SourceConfig: &model.ResourceProperty_Mapping{
				Mapping: &model.ResourcePropertyMappingConfig{
					Mapping: "source_type",
				},
			},
			Type:     model.ResourcePropertyType_TYPE_INT32,
			Required: true,
		},
		{
			Name: "sourceMapping",
			SourceConfig: &model.ResourceProperty_Mapping{
				Mapping: &model.ResourcePropertyMappingConfig{
					Mapping: "source_mapping",
				},
			},
			Type:     model.ResourcePropertyType_TYPE_STRING,
			Length:   64,
			Required: true,
		},
		{
			Name: "sourceDef",
			SourceConfig: &model.ResourceProperty_Mapping{
				Mapping: &model.ResourcePropertyMappingConfig{
					Mapping: "source_def",
				},
			},
			Type:     model.ResourcePropertyType_TYPE_STRING,
			Length:   64,
			Required: true,
		},
		{
			Name: "sourcePrimary",
			SourceConfig: &model.ResourceProperty_Mapping{
				Mapping: &model.ResourcePropertyMappingConfig{
					Mapping: "source_primary",
				},
			},
			Type:     model.ResourcePropertyType_TYPE_BOOL,
			Required: true,
		},
		{
			Name: "sourceAutoGeneration",
			SourceConfig: &model.ResourceProperty_Mapping{
				Mapping: &model.ResourcePropertyMappingConfig{
					Mapping: "source_auto_generation",
				},
			},
			Type:     model.ResourcePropertyType_TYPE_INT32,
			Required: true,
		},
		{
			Name: "required",
			SourceConfig: &model.ResourceProperty_Mapping{
				Mapping: &model.ResourcePropertyMappingConfig{
					Mapping: "required",
				},
			},
			Type:     model.ResourcePropertyType_TYPE_BOOL,
			Required: true,
		},
		{
			Name: "unique",
			SourceConfig: &model.ResourceProperty_Mapping{
				Mapping: &model.ResourcePropertyMappingConfig{
					Mapping: "unique",
				},
			},
			Type:     model.ResourcePropertyType_TYPE_BOOL,
			Required: true,
		},
		{
			Name: "length",
			SourceConfig: &model.ResourceProperty_Mapping{
				Mapping: &model.ResourcePropertyMappingConfig{
					Mapping: "length",
				},
			},
			Type:     model.ResourcePropertyType_TYPE_INT32,
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
			Cascade:            false,
		},
	},
}
