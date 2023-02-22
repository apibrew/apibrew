package test

import (
	"github.com/tislib/data-handler/pkg/model"
)

var systemDataSource = &model.DataSource{
	Id:          "system",
	Backend:     model.DataSourceBackendType_POSTGRESQL,
	Type:        model.DataType_SYSTEM,
	Name:        "system",
	Description: "system",
	Options: &model.DataSource_PostgresqlParams{
		PostgresqlParams: &model.PostgresqlOptions{
			Username:      "root",
			Password:      "root",
			Host:          "127.0.0.1",
			Port:          5432,
			DbName:        "dh_system",
			DefaultSchema: "public",
		},
	},
}

var dhTest = &model.DataSource{
	Backend:     model.DataSourceBackendType_POSTGRESQL,
	Name:        "dh-test",
	Description: "dh-test",
	Options: &model.DataSource_PostgresqlParams{
		PostgresqlParams: &model.PostgresqlOptions{
			Username:      "dh_test",
			Password:      "dh_test",
			Host:          "127.0.0.1",
			Port:          5432,
			DbName:        "dh_test",
			DefaultSchema: "public",
		},
	},
}

var dhTestWrongPassword = &model.DataSource{
	Backend:     model.DataSourceBackendType_POSTGRESQL,
	Name:        "data-source-1-wrong",
	Description: "data-source-1-wrong",
	Options: &model.DataSource_PostgresqlParams{
		PostgresqlParams: &model.PostgresqlOptions{
			Username:      "dh_test_wrong_pass",
			Password:      "dh_test_wrong_pass",
			Host:          "127.0.0.1",
			Port:          5432,
			DbName:        "dh_test",
			DefaultSchema: "public",
		},
	},
}

var dataSourceDhTest = &model.DataSource{
	Backend:     model.DataSourceBackendType_POSTGRESQL,
	Name:        "data-source-test",
	Description: "data-source-test",
	Options: &model.DataSource_PostgresqlParams{
		PostgresqlParams: &model.PostgresqlOptions{
			Username:      "dh_test",
			Password:      "dh_test",
			Host:          "127.0.0.1",
			Port:          5432,
			DbName:        "dh_test",
			DefaultSchema: "public",
		},
	},
}

var dataSource1 = &model.DataSource{
	Backend:     model.DataSourceBackendType_POSTGRESQL,
	Name:        "data-source-1",
	Description: "data-source-1",
	Options: &model.DataSource_PostgresqlParams{
		PostgresqlParams: &model.PostgresqlOptions{
			Username:      "dh_test",
			Password:      "dh_test",
			Host:          "127.0.0.1",
			Port:          5432,
			DbName:        "dh_test",
			DefaultSchema: "public",
		},
	},
}

var richResource1 = &model.Resource{
	Name:      "rich-test-3995",
	Namespace: "default",
	DataType:  2,
	SourceConfig: &model.ResourceSourceConfig{
		DataSource: dhTest.Name,
		Entity:     "rich_test_3995",
	},
	Properties: []*model.ResourceProperty{
		{
			Name: "int32_o",
			Type: model.ResourcePropertyType_TYPE_INT32,

			Mapping:  "int32_o",
			Required: false,
		},
		{
			Name: "int32",
			Type: model.ResourcePropertyType_TYPE_INT32,

			Mapping:  "int32",
			Required: true,
		},

		{
			Name: "int64",
			Type: model.ResourcePropertyType_TYPE_INT64,

			Mapping:  "int64",
			Required: true,
		},

		{
			Name: "float",
			Type: model.ResourcePropertyType_TYPE_FLOAT32,

			Mapping:  "float",
			Required: true,
		},

		{
			Name: "double",
			Type: model.ResourcePropertyType_TYPE_FLOAT64,

			Mapping:  "double",
			Required: true,
		},

		{
			Name: "text",
			Type: model.ResourcePropertyType_TYPE_STRING,

			Mapping:  "text",
			Length:   255,
			Required: true,
		},

		{
			Name: "string",
			Type: model.ResourcePropertyType_TYPE_STRING,

			Mapping:  "string",
			Required: true,
			Length:   255,
		},
		{
			Name: "uuid",
			Type: model.ResourcePropertyType_TYPE_UUID,

			Mapping:  "uuid",
			Required: true,
		},

		{
			Name: "date",
			Type: model.ResourcePropertyType_TYPE_DATE,

			Mapping:  "date",
			Required: true,
		},

		{
			Name: "time",
			Type: model.ResourcePropertyType_TYPE_TIME,

			Mapping:  "time",
			Required: true,
		},

		{
			Name: "timestamp",
			Type: model.ResourcePropertyType_TYPE_TIMESTAMP,

			Mapping:  "timestamp",
			Required: true,
		},

		{
			Name: "bool",
			Type: model.ResourcePropertyType_TYPE_BOOL,

			Mapping:  "bool",
			Required: true,
		},

		{
			Name: "object",
			Type: model.ResourcePropertyType_TYPE_OBJECT,

			Mapping:  "object",
			Required: true,
		},

		{
			Name: "bytes",
			Type: model.ResourcePropertyType_TYPE_BYTES,

			Mapping:  "bytes",
			Required: true,
		},
	},
}

var simpleVirtualResource1 = &model.Resource{
	Name:      "virtualResource",
	Namespace: "default",
	DataType:  2,
	Virtual:   true,
	Properties: []*model.ResourceProperty{
		{
			Name: "name",
			Type: model.ResourcePropertyType_TYPE_STRING,

			Mapping:  "name",
			Length:   255,
			Required: true,
		},
		{
			Name: "description",
			Type: model.ResourcePropertyType_TYPE_STRING,

			Mapping:  "description",
			Length:   255,
			Required: false,
		},
	},
}
