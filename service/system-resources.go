package service

import "data-handler/stub/model"

var dataSourceResource = &model.Resource{
	Name:      "data-source",
	Workspace: "system",
	Type:      model.DataType_SYSTEM,
	SourceConfig: &model.ResourceSourceConfig{
		DataSource: "system",
		Mapping:    "data_source",
	},
	Properties: []*model.ResourceProperty{
		{
			Name: "backend",
			SourceConfig: &model.ResourceProperty_Mapping{
				Mapping: &model.ResourcePropertyMappingConfig{
					Mapping:        "backend",
					SourceDef:      "",
					AutoGeneration: 0,
				},
			},
			Primary:  false,
			Type:     model.ResourcePropertyType_TYPE_INT32,
			Required: true,
		},
		{
			Name: "options_postgres_username",
			SourceConfig: &model.ResourceProperty_Mapping{
				Mapping: &model.ResourcePropertyMappingConfig{
					Mapping: "options_postgres_username",
				},
			},
			Type:     model.ResourcePropertyType_TYPE_STRING,
			Length:   64,
			Required: false,
		},
		{
			Name: "options_postgres_password",
			SourceConfig: &model.ResourceProperty_Mapping{
				Mapping: &model.ResourcePropertyMappingConfig{
					Mapping: "options_postgres_password",
				},
			},
			Type:     model.ResourcePropertyType_TYPE_STRING,
			Length:   64,
			Required: false,
		},
		{
			Name: "options_postgres_host",
			SourceConfig: &model.ResourceProperty_Mapping{
				Mapping: &model.ResourcePropertyMappingConfig{
					Mapping: "options_postgres_host"},
			},
			Type:     model.ResourcePropertyType_TYPE_STRING,
			Length:   64,
			Required: false,
		},
		{
			Name: "options_postgres_port",
			SourceConfig: &model.ResourceProperty_Mapping{
				Mapping: &model.ResourcePropertyMappingConfig{
					Mapping: "options_postgres_port",
				},
			},
			Type:     model.ResourcePropertyType_TYPE_INT32,
			Required: false,
		},
		{
			Name: "options_postgres_db_name",
			SourceConfig: &model.ResourceProperty_Mapping{
				Mapping: &model.ResourcePropertyMappingConfig{
					Mapping: "options_postgres_db_name",
				},
			},
			Type:     model.ResourcePropertyType_TYPE_STRING,
			Length:   64,
			Required: false,
		},
		{
			Name: "options_postgres_default_schema",
			SourceConfig: &model.ResourceProperty_Mapping{
				Mapping: &model.ResourcePropertyMappingConfig{
					Mapping: "options_postgres_default_schema",
				},
			},
			Type:     model.ResourcePropertyType_TYPE_STRING,
			Length:   64,
			Required: false,
		},
	},
}
