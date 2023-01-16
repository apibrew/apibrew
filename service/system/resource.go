package system

import "data-handler/model"

var ResourceResource = &model.Resource{
	Name:      "resource",
	Workspace: "system",
	DataType:  model.DataType_SYSTEM,
	SourceConfig: &model.ResourceSourceConfig{
		DataSource: "system",
		Mapping:    "resource",
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
			Unique:   true,
		},
		{
			Name: "workspace",
			SourceConfig: &model.ResourceProperty_Mapping{
				Mapping: &model.ResourcePropertyMappingConfig{
					Mapping: "workspace",
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
			Type:     model.ResourcePropertyType_TYPE_STRING,
			Length:   256,
			Required: true,
		},
		{
			Name: "mapping",
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
			Name: "readOnlyRecords",
			SourceConfig: &model.ResourceProperty_Mapping{
				Mapping: &model.ResourcePropertyMappingConfig{
					Mapping: "read_only_records",
				},
			},
			Type:     model.ResourcePropertyType_TYPE_BOOL,
			Required: true,
		},
		{
			Name: "uniqueRecord",
			SourceConfig: &model.ResourceProperty_Mapping{
				Mapping: &model.ResourcePropertyMappingConfig{
					Mapping: "unique_record",
				},
			},
			Type:     model.ResourcePropertyType_TYPE_BOOL,
			Required: true,
		},
		{
			Name: "keepHistory",
			SourceConfig: &model.ResourceProperty_Mapping{
				Mapping: &model.ResourcePropertyMappingConfig{
					Mapping: "keep_history",
				},
			},
			Type:     model.ResourcePropertyType_TYPE_BOOL,
			Required: true,
		},
		{
			Name: "autoCreated",
			SourceConfig: &model.ResourceProperty_Mapping{
				Mapping: &model.ResourcePropertyMappingConfig{
					Mapping: "auto_created",
				},
			},
			Type:     model.ResourcePropertyType_TYPE_BOOL,
			Required: true,
		},
		{
			Name: "disableMigration",
			SourceConfig: &model.ResourceProperty_Mapping{
				Mapping: &model.ResourcePropertyMappingConfig{
					Mapping: "disable_migration",
				},
			},
			Type:     model.ResourcePropertyType_TYPE_BOOL,
			Required: true,
		},
		{
			Name: "disableAudit",
			SourceConfig: &model.ResourceProperty_Mapping{
				Mapping: &model.ResourcePropertyMappingConfig{
					Mapping: "disable_audit",
				},
			},
			Type:     model.ResourcePropertyType_TYPE_BOOL,
			Required: true,
		},
		{
			Name: "doPrimaryKeyLookup",
			SourceConfig: &model.ResourceProperty_Mapping{
				Mapping: &model.ResourcePropertyMappingConfig{
					Mapping: "do_primary_key_lookup",
				},
			},
			Type:     model.ResourcePropertyType_TYPE_BOOL,
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
	},
}
