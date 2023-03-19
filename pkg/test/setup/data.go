package setup

import (
	"github.com/tislib/data-handler/pkg/model"
)

var SystemDataSource = &model.DataSource{
	Id:          "system",
	Backend:     model.DataSourceBackendType_POSTGRESQL,
	Name:        "system",
	Description: "system",
	Params: &model.DataSource_PostgresqlParams{
		PostgresqlParams: &model.PostgresqlParams{
			Username:      "root",
			Password:      "root",
			Host:          "127.0.0.1",
			Port:          5432,
			DbName:        "dh_system",
			DefaultSchema: "public",
		},
	},
}

var DhTest = &model.DataSource{
	Backend:     model.DataSourceBackendType_POSTGRESQL,
	Name:        "dh-test",
	Description: "dh-test",
	Params: &model.DataSource_PostgresqlParams{
		PostgresqlParams: &model.PostgresqlParams{
			Username:      "dh_test",
			Password:      "dh_test",
			Host:          "127.0.0.1",
			Port:          5432,
			DbName:        "dh_test",
			DefaultSchema: "public",
		},
	},
}

var DhTestWrongPassword = &model.DataSource{
	Backend:     model.DataSourceBackendType_POSTGRESQL,
	Name:        "data-source-1-wrong",
	Description: "data-source-1-wrong",
	Params: &model.DataSource_PostgresqlParams{
		PostgresqlParams: &model.PostgresqlParams{
			Username:      "dh_test_wrong_pass",
			Password:      "dh_test_wrong_pass",
			Host:          "127.0.0.1",
			Port:          5432,
			DbName:        "dh_test",
			DefaultSchema: "public",
		},
	},
}

var DataSourceDhTest = &model.DataSource{
	Backend:     model.DataSourceBackendType_POSTGRESQL,
	Name:        "data-source-test",
	Description: "data-source-test",
	Params: &model.DataSource_PostgresqlParams{
		PostgresqlParams: &model.PostgresqlParams{
			Username:      "dh_test",
			Password:      "dh_test",
			Host:          "127.0.0.1",
			Port:          5432,
			DbName:        "dh_test",
			DefaultSchema: "public",
		},
	},
}

var DataSource1 = &model.DataSource{
	Backend:     model.DataSourceBackendType_POSTGRESQL,
	Name:        "data-source-1",
	Description: "data-source-1",
	Version:     1,
	Params: &model.DataSource_PostgresqlParams{
		PostgresqlParams: &model.PostgresqlParams{
			Username:      "dh_test",
			Password:      "dh_test",
			Host:          "127.0.0.1",
			Port:          5432,
			DbName:        "dh_test",
			DefaultSchema: "public",
		},
	},
}

var RichResource1 = PrepareRichResource1()

func PrepareRichResource1() *model.Resource {
	return &model.Resource{
		Name:      "rich-test-3995",
		Namespace: "default",
		SourceConfig: &model.ResourceSourceConfig{
			DataSource: DhTest.Name,
			Entity:     "rich_test_3995",
		},
		Properties: []*model.ResourceProperty{
			{
				Name: "int32_o",
				Type: model.ResourceProperty_INT32,

				Mapping:  "int32_o",
				Required: false,
			},
			{
				Name: "int32",
				Type: model.ResourceProperty_INT32,

				Mapping:  "int32",
				Required: true,
			},

			{
				Name: "int64",
				Type: model.ResourceProperty_INT64,

				Mapping:  "int64",
				Required: true,
			},

			{
				Name: "float",
				Type: model.ResourceProperty_FLOAT32,

				Mapping:  "float",
				Required: true,
			},

			{
				Name: "double",
				Type: model.ResourceProperty_FLOAT64,

				Mapping:  "double",
				Required: true,
			},

			{
				Name: "text",
				Type: model.ResourceProperty_STRING,

				Mapping:  "text",
				Length:   255,
				Required: true,
			},

			{
				Name: "string",
				Type: model.ResourceProperty_STRING,

				Mapping:  "string",
				Required: true,
				Length:   255,
			},
			{
				Name: "uuid",
				Type: model.ResourceProperty_UUID,

				Mapping:  "uuid",
				Required: true,
			},

			{
				Name: "date",
				Type: model.ResourceProperty_DATE,

				Mapping:  "date",
				Required: true,
			},

			{
				Name: "time",
				Type: model.ResourceProperty_TIME,

				Mapping:  "time",
				Required: true,
			},

			{
				Name: "timestamp",
				Type: model.ResourceProperty_TIMESTAMP,

				Mapping:  "timestamp",
				Required: true,
			},

			{
				Name: "bool",
				Type: model.ResourceProperty_BOOL,

				Mapping:  "bool",
				Required: true,
			},

			{
				Name: "object",
				Type: model.ResourceProperty_OBJECT,

				Mapping:  "object",
				Required: true,
			},

			{
				Name: "bytes",
				Type: model.ResourceProperty_BYTES,

				Mapping:  "bytes",
				Required: true,
			},
		},
	}
}

var SimpleVirtualResource1 = &model.Resource{
	Name:      "virtualResource",
	Namespace: "default",
	Virtual:   true,
	Properties: []*model.ResourceProperty{
		{
			Name: "name",
			Type: model.ResourceProperty_STRING,

			Mapping:  "name",
			Length:   255,
			Required: true,
		},
		{
			Name: "description",
			Type: model.ResourceProperty_STRING,

			Mapping:  "description",
			Length:   255,
			Required: false,
		},
	},
}
