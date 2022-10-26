package system

import "data-handler/model"

var WorkspaceResource = &model.Resource{
	Name:      "workspace",
	Workspace: "system",
	Type:      model.DataType_SYSTEM,
	SourceConfig: &model.ResourceSourceConfig{
		DataSource: "system",
		Mapping:    "workspace",
	},
	Flags: &model.ResourceFlags{},
	Properties: []*model.ResourceProperty{
		{
			Name: "name",
			SourceConfig: &model.ResourceProperty_Mapping{
				Mapping: &model.ResourcePropertyMappingConfig{
					Mapping: "name",
				},
			},
			Primary:  false,
			Type:     model.ResourcePropertyType_TYPE_STRING,
			Length:   256,
			Required: true,
			Unique:   true,
		},
		{
			Name: "description",
			SourceConfig: &model.ResourceProperty_Mapping{
				Mapping: &model.ResourcePropertyMappingConfig{
					Mapping: "description",
				},
			},
			Primary:  false,
			Type:     model.ResourcePropertyType_TYPE_STRING,
			Length:   256,
			Required: false,
		},
		{
			Name: "details",
			SourceConfig: &model.ResourceProperty_Mapping{
				Mapping: &model.ResourcePropertyMappingConfig{
					Mapping: "details",
				},
			},
			Primary:  false,
			Type:     model.ResourcePropertyType_TYPE_OBJECT,
			Required: false,
		},
	},
}
