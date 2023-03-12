package resources

import (
	"github.com/tislib/data-handler/pkg/model"
)

var DataSourceResource = &model.Resource{
	Name:      "data-source",
	Namespace: "system",
	SourceConfig: &model.ResourceSourceConfig{
		DataSource: "system",
		Entity:     "data_source",
	},
	Properties: []*model.ResourceProperty{
		{
			Name:     "name",
			Mapping:  "name",
			Length:   64,
			Primary:  false,
			Unique:   true,
			Type:     model.ResourcePropertyType_TYPE_STRING,
			Required: true,
		},
		{
			Name:     "description",
			Mapping:  "description",
			Length:   64,
			Primary:  false,
			Type:     model.ResourcePropertyType_TYPE_STRING,
			Required: false,
		},
		{
			Name:     "backend",
			Mapping:  "backend",
			Primary:  false,
			Type:     model.ResourcePropertyType_TYPE_INT32,
			Required: true,
		},
		{
			Name:     "options_postgres_username",
			Mapping:  "options_postgres_username",
			Type:     model.ResourcePropertyType_TYPE_STRING,
			Length:   64,
			Required: false,
		},
		{
			Name: "options_postgres_password",

			Mapping:  "options_postgres_password",
			Type:     model.ResourcePropertyType_TYPE_STRING,
			Length:   64,
			Required: false,
		},
		{
			Name: "options_postgres_host",

			Mapping:  "options_postgres_host",
			Type:     model.ResourcePropertyType_TYPE_STRING,
			Length:   64,
			Required: false,
		},
		{
			Name: "options_postgres_port",

			Mapping:  "options_postgres_port",
			Type:     model.ResourcePropertyType_TYPE_INT32,
			Required: false,
		},
		{
			Name: "options_postgres_db_name",

			Mapping:  "options_postgres_db_name",
			Type:     model.ResourcePropertyType_TYPE_STRING,
			Length:   64,
			Required: false,
		},
		{
			Name: "options_postgres_default_schema",

			Mapping:  "options_postgres_default_schema",
			Type:     model.ResourcePropertyType_TYPE_STRING,
			Length:   64,
			Required: false,
		},
	},
	SecurityContext: securityContextDisallowAll,
}
