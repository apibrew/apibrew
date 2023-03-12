package resources

import (
	"github.com/tislib/data-handler/pkg/model"
)

var ExtensionResource = &model.Resource{
	Name:      "extension",
	Namespace: "system",
	SourceConfig: &model.ResourceSourceConfig{
		DataSource: "system",
		Entity:     "extension",
	},
	Properties: []*model.ResourceProperty{
		{
			Name: "name",

			Mapping:  "name",
			Length:   64,
			Primary:  false,
			Type:     model.ResourcePropertyType_TYPE_STRING,
			Unique:   true,
			Required: true,
		},
		{
			Name: "description",

			Mapping:  "description",
			Length:   64,
			Primary:  false,
			Type:     model.ResourcePropertyType_TYPE_STRING,
			Required: false,
		},
		{
			Name: "namespace",

			Mapping:  "namespace",
			Length:   64,
			Primary:  false,
			Type:     model.ResourcePropertyType_TYPE_STRING,
			Required: true,
		},
		{
			Name: "resource",

			Mapping:  "resource",
			Length:   64,
			Primary:  false,
			Type:     model.ResourcePropertyType_TYPE_STRING,
			Required: true,
		},
		{
			Name: "serverHost",

			Mapping:  "server_host",
			Length:   64,
			Primary:  false,
			Type:     model.ResourcePropertyType_TYPE_STRING,
			Required: true,
		},
		{
			Name: "serverPort",

			Mapping:  "server_port",
			Length:   64,
			Primary:  false,
			Type:     model.ResourcePropertyType_TYPE_INT32,
			Required: true,
		},
		{
			Name: "operations",

			Mapping: "operations",
			Type:    model.ResourcePropertyType_TYPE_OBJECT,
		},
	},
	SecurityContext: securityContextDisallowAll,
}
