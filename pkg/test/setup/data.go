package setup

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resource_model"
	"github.com/apibrew/apibrew/pkg/service/annotations"
)

var SystemDataSource = &resource_model.DataSource{
	Backend: "POSTGRESQL",
	Name:    "system",
	Options: map[string]string{
		"username":       "dh_test",
		"password":       "dh_test",
		"host":           "127.0.0.1",
		"port":           "5432",
		"db_name":        "dh_test",
		"default_schema": "public",
	},
}

var DefaultDataSource = &resource_model.DataSource{
	Backend: "POSTGRESQL",
	Name:    "default",
	Options: map[string]string{
		"username":       "dh_test",
		"password":       "dh_test",
		"host":           "127.0.0.1",
		"port":           "5432",
		"db_name":        "dh_test",
		"default_schema": "public",
	},
}

var DhTest = &resource_model.DataSource{
	Backend: "POSTGRESQL",
	Name:    "dh-test",
	Options: map[string]string{
		"username":       "dh_test",
		"password":       "dh_test",
		"host":           "127.0.0.1",
		"port":           "5432",
		"db_name":        "dh_test",
		"default_schema": "public",
	},
}

var DhTestWrongPassword = &resource_model.DataSource{
	Backend: "POSTGRESQL",
	Name:    "data-source-1-wrong",
	Options: map[string]string{
		"username":       "dh_test",
		"password":       "dh_test_wrong",
		"host":           "127.0.0.1",
		"port":           "5432",
		"db_name":        "dh_test",
		"default_schema": "public",
	},
}

var DataSourceDhTest = &resource_model.DataSource{
	Backend: "POSTGRESQL",
	Name:    "data-source-test",
	Options: map[string]string{
		"username":       "dh_test",
		"password":       "dh_test",
		"host":           "127.0.0.1",
		"port":           "5432",
		"db_name":        "dh_test",
		"default_schema": "public",
	},
}

var DataSource1 = &resource_model.DataSource{
	Backend: "POSTGRESQL",
	Name:    "data-source-1",
	Options: map[string]string{
		"username":       "dh_test",
		"password":       "dh_test",
		"host":           "127.0.0.1",
		"port":           "5432",
		"db_name":        "dh_test",
		"default_schema": "public",
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

				Required: false,
			},
			{
				Name: "int32",
				Type: model.ResourceProperty_INT32,

				Required: true,
			},

			{
				Name: "int64",
				Type: model.ResourceProperty_INT64,

				Required: true,
			},

			{
				Name: "float",
				Type: model.ResourceProperty_FLOAT32,

				Required: true,
			},

			{
				Name: "double",
				Type: model.ResourceProperty_FLOAT64,

				Required: true,
			},

			{
				Name: "text",
				Type: model.ResourceProperty_STRING,

				Length:   255,
				Required: true,
			},

			{
				Name: "string",
				Type: model.ResourceProperty_STRING,

				Required: true,
				Length:   255,
			},
			{
				Name: "uuid",
				Type: model.ResourceProperty_UUID,

				Required: true,
			},

			{
				Name: "date",
				Type: model.ResourceProperty_DATE,

				Required: true,
			},

			{
				Name: "time",
				Type: model.ResourceProperty_TIME,

				Required: true,
			},

			{
				Name: "timestamp",
				Type: model.ResourceProperty_TIMESTAMP,

				Required: true,
			},

			{
				Name: "bool",
				Type: model.ResourceProperty_BOOL,

				Required: true,
			},

			{
				Name: "object",
				Type: model.ResourceProperty_OBJECT,

				Required: true,
				Annotations: map[string]string{
					annotations.HclBlock: "object",
				},
			},

			{
				Name: "bytes",
				Type: model.ResourceProperty_BYTES,

				Required: false,
			},
		},
		Annotations: map[string]string{
			annotations.EnableAudit: "true",
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

			Length:   255,
			Required: true,
		},
		{
			Name: "description",
			Type: model.ResourceProperty_STRING,

			Length:   255,
			Required: false,
		},
	},
}
