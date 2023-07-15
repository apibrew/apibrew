package setup

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resource_model"
	"github.com/apibrew/apibrew/pkg/service/annotations"
)

var SystemDataSource = &resource_model.DataSource{
	Backend: resource_model.DataSourceBackend_POSTGRESQL,
	Name:    "system",
	Options: map[string]string{
		"Username":      "dh_test",
		"Password":      "dh_test",
		"Host":          "127.0.0.1",
		"Port":          "5432",
		"DbName":        "dh_test",
		"DefaultSchema": "public",
	},
}

var DefaultDataSource = &resource_model.DataSource{
	Backend: resource_model.DataSourceBackend_POSTGRESQL,
	Name:    "default",
	Options: map[string]string{
		"Username":      "dh_test",
		"Password":      "dh_test",
		"Host":          "127.0.0.1",
		"Port":          "5432",
		"DbName":        "dh_test",
		"DefaultSchema": "public",
	},
}

var DhTest = &resource_model.DataSource{
	Backend: resource_model.DataSourceBackend_POSTGRESQL,
	Name:    "dh-test",
	Options: map[string]string{
		"Username":      "dh_test",
		"Password":      "dh_test",
		"Host":          "127.0.0.1",
		"Port":          "5432",
		"DbName":        "dh_test",
		"DefaultSchema": "public",
	},
}

var DhTestWrongPassword = &resource_model.DataSource{
	Backend: resource_model.DataSourceBackend_POSTGRESQL,
	Name:    "data-source-1-wrong",
	Options: map[string]string{
		"Username":      "dh_test",
		"Password":      "dh_test_wrong",
		"Host":          "127.0.0.1",
		"Port":          "5432",
		"DbName":        "dh_test",
		"DefaultSchema": "public",
	},
}

var DataSourceDhTest = &resource_model.DataSource{
	Backend: resource_model.DataSourceBackend_POSTGRESQL,
	Name:    "data-source-test",
	Options: map[string]string{
		"Username":      "dh_test",
		"Password":      "dh_test",
		"Host":          "127.0.0.1",
		"Port":          "5432",
		"DbName":        "dh_test",
		"DefaultSchema": "public",
	},
}

var DataSource1 = &resource_model.DataSource{
	Backend: resource_model.DataSourceBackend_POSTGRESQL,
	Name:    "data-source-1",
	Options: map[string]string{
		"Username":      "dh_test",
		"Password":      "dh_test",
		"Host":          "127.0.0.1",
		"Port":          "5432",
		"DbName":        "dh_test",
		"DefaultSchema": "public",
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
				Annotations: map[string]string{
					annotations.HclBlock: "object",
				},
			},

			{
				Name: "bytes",
				Type: model.ResourceProperty_BYTES,

				Mapping:  "bytes",
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
