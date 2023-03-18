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
		IdProperty,
		VersionProperty,
		AuditProperties[0],
		AuditProperties[1],
		AuditProperties[2],
		AuditProperties[3],
		{
			Name:     "name",
			Mapping:  "name",
			Length:   64,
			Primary:  false,
			Unique:   true,
			Type:     model.ResourceProperty_STRING,
			Required: true,
		},
		{
			Name:     "description",
			Mapping:  "description",
			Length:   64,
			Primary:  false,
			Type:     model.ResourceProperty_STRING,
			Required: false,
		},
		{
			Name:     "backend",
			Mapping:  "backend",
			Primary:  false,
			Type:     model.ResourceProperty_INT32,
			Required: true,
		},

		{
			Name:     "options_postgres_username",
			Mapping:  "options_postgres_username",
			Type:     model.ResourceProperty_STRING,
			Length:   64,
			Required: false,
		},
		{
			Name: "options_postgres_password",

			Mapping:  "options_postgres_password",
			Type:     model.ResourceProperty_STRING,
			Length:   64,
			Required: false,
		},
		{
			Name: "options_postgres_host",

			Mapping:  "options_postgres_host",
			Type:     model.ResourceProperty_STRING,
			Length:   64,
			Required: false,
		},
		{
			Name: "options_postgres_port",

			Mapping:  "options_postgres_port",
			Type:     model.ResourceProperty_INT32,
			Required: false,
		},
		{
			Name: "options_postgres_db_name",

			Mapping:  "options_postgres_db_name",
			Type:     model.ResourceProperty_STRING,
			Length:   64,
			Required: false,
		},
		{
			Name: "options_postgres_default_schema",

			Mapping:  "options_postgres_default_schema",
			Type:     model.ResourceProperty_STRING,
			Length:   64,
			Required: false,
		},

		{
			Name:     "options_mysql_username",
			Mapping:  "options_mysql_username",
			Type:     model.ResourceProperty_STRING,
			Length:   64,
			Required: false,
		},
		{
			Name: "options_mysql_password",

			Mapping:  "options_mysql_password",
			Type:     model.ResourceProperty_STRING,
			Length:   64,
			Required: false,
		},
		{
			Name: "options_mysql_host",

			Mapping:  "options_mysql_host",
			Type:     model.ResourceProperty_STRING,
			Length:   64,
			Required: false,
		},
		{
			Name: "options_mysql_port",

			Mapping:  "options_mysql_port",
			Type:     model.ResourceProperty_INT32,
			Required: false,
		},
		{
			Name: "options_mysql_db_name",

			Mapping:  "options_mysql_db_name",
			Type:     model.ResourceProperty_STRING,
			Length:   64,
			Required: false,
		},
		{
			Name: "options_mysql_default_schema",

			Mapping:  "options_mysql_default_schema",
			Type:     model.ResourceProperty_STRING,
			Length:   64,
			Required: false,
		},
	},
	SecurityContext: securityContextDisallowAll,
}
