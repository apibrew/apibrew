package resources

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resources/special"
	"github.com/apibrew/apibrew/pkg/service/annotations"
)

var DataSourceResource = &model.Resource{
	Name:      "data-source",
	Namespace: "system",
	SourceConfig: &model.ResourceSourceConfig{
		DataSource: "system",
		Entity:     "data_source",
	},
	Properties: []*model.ResourceProperty{
		special.IdProperty,
		special.VersionProperty,
		special.AuditProperties[0],
		special.AuditProperties[1],
		special.AuditProperties[2],
		special.AuditProperties[3],
		{
			Name:     "name",
			Mapping:  "name",
			Length:   64,
			Primary:  false,
			Unique:   true,
			Type:     model.ResourceProperty_STRING,
			Required: true,
			Annotations: map[string]string{
				annotations.IsHclLabel: annotations.Enabled,
			},
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
		{
			Name: "options_redis_addr",

			Mapping:  "options_redis_addr",
			Type:     model.ResourceProperty_STRING,
			Length:   256,
			Required: false,
		},
		{
			Name: "options_redis_password",

			Mapping:  "options_redis_password",
			Type:     model.ResourceProperty_STRING,
			Length:   256,
			Required: false,
		},
		{
			Name: "options_redis_db",

			Mapping:  "options_redis_db",
			Type:     model.ResourceProperty_INT32,
			Required: false,
		},
		{
			Name: "options_mongo_uri",

			Mapping:  "options_mongo_uri",
			Type:     model.ResourceProperty_STRING,
			Length:   256,
			Required: false,
		},
		{
			Name: "options_mongo_db_name",

			Mapping:  "options_mongo_db_name",
			Type:     model.ResourceProperty_STRING,
			Length:   256,
			Required: false,
		},
	},
	SecurityContext: special.SecurityContextDisallowAll,
	Annotations: map[string]string{
		annotations.EnableAudit: "true",
	},
}
