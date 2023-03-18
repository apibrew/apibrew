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
		IdProperty,
		VersionProperty,
		AuditProperties[0],
		AuditProperties[1],
		AuditProperties[2],
		AuditProperties[3],
		{
			Name: "name",

			Mapping:  "name",
			Length:   64,
			Primary:  false,
			Type:     model.ResourceProperty_STRING,
			Unique:   true,
			Required: true,
		},
		{
			Name: "description",

			Mapping:  "description",
			Length:   64,
			Primary:  false,
			Type:     model.ResourceProperty_STRING,
			Required: false,
		},
		{
			Name: "namespace",

			Mapping:  "namespace",
			Length:   64,
			Primary:  false,
			Type:     model.ResourceProperty_STRING,
			Required: true,
		},
		{
			Name: "resource",

			Mapping:  "resource",
			Length:   64,
			Primary:  false,
			Type:     model.ResourceProperty_STRING,
			Required: true,
		},
		{
			Name: "serverHost",

			Mapping:  "server_host",
			Length:   64,
			Primary:  false,
			Type:     model.ResourceProperty_STRING,
			Required: true,
		},
		{
			Name: "serverPort",

			Mapping:  "server_port",
			Length:   64,
			Primary:  false,
			Type:     model.ResourceProperty_INT32,
			Required: true,
		},
		{
			Name: "operations",

			Mapping: "operations",
			Type:    model.ResourceProperty_OBJECT,
		},
	},
	SecurityContext: securityContextDisallowAll,
}
