package system

import "data-handler/model"

var UserResource = &model.Resource{
	Name:      "user",
	Namespace: "system",
	DataType:  model.DataType_SYSTEM,
	SourceConfig: &model.ResourceSourceConfig{
		DataSource: "system",
		Entity:     "user",
	},
	Flags: &model.ResourceFlags{},
	Properties: []*model.ResourceProperty{
		{
			Name: "username",
			SourceConfig: &model.ResourceProperty_Mapping{
				Mapping: &model.ResourcePropertyMappingConfig{
					Mapping:        "username",
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
			Name: "password",
			SourceConfig: &model.ResourceProperty_Mapping{
				Mapping: &model.ResourcePropertyMappingConfig{
					Mapping: "password",
				},
			},
			Type:     model.ResourcePropertyType_TYPE_PASSWORD,
			Length:   256,
			Required: true,
		},
		{
			Name: "scopes",
			SourceConfig: &model.ResourceProperty_Mapping{
				Mapping: &model.ResourcePropertyMappingConfig{
					Mapping: "scopes",
				},
			},
			Type:     model.ResourcePropertyType_TYPE_TEXT,
			Required: false,
		},
		{
			Name: "details",
			SourceConfig: &model.ResourceProperty_Mapping{
				Mapping: &model.ResourcePropertyMappingConfig{
					Mapping: "details",
				},
			},
			Type: model.ResourcePropertyType_TYPE_OBJECT,
		},
	},
}
