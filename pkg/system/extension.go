package system

import (
	"github.com/tislib/data-handler/pkg/model"
)

var ExtensionResource = &model.Resource{
	Name:      "extension",
	Namespace: "system",
	DataType:  model.DataType_SYSTEM,
	SourceConfig: &model.ResourceSourceConfig{
		DataSource: "system",
		Entity:     "extension",
	},
	Properties: []*model.ResourceProperty{
		{
			Name: "name",
			SourceConfig: &model.ResourceProperty_Mapping{
				Mapping: &model.ResourcePropertyMappingConfig{
					Mapping: "name",
				},
			},
			Length:   64,
			Primary:  false,
			Type:     model.ResourcePropertyType_TYPE_STRING,
			Unique:   true,
			Required: true,
		},
		{
			Name: "description",
			SourceConfig: &model.ResourceProperty_Mapping{
				Mapping: &model.ResourcePropertyMappingConfig{
					Mapping: "description",
				},
			},
			Length:   64,
			Primary:  false,
			Type:     model.ResourcePropertyType_TYPE_STRING,
			Required: false,
		},
		{
			Name: "namespace",
			SourceConfig: &model.ResourceProperty_Mapping{
				Mapping: &model.ResourcePropertyMappingConfig{
					Mapping: "namespace",
				},
			},
			Length:   64,
			Primary:  false,
			Type:     model.ResourcePropertyType_TYPE_STRING,
			Required: true,
		},
		{
			Name: "resource",
			SourceConfig: &model.ResourceProperty_Mapping{
				Mapping: &model.ResourcePropertyMappingConfig{
					Mapping: "resource",
				},
			},
			Length:   64,
			Primary:  false,
			Type:     model.ResourcePropertyType_TYPE_STRING,
			Required: true,
		},
		{
			Name: "serverHost",
			SourceConfig: &model.ResourceProperty_Mapping{
				Mapping: &model.ResourcePropertyMappingConfig{
					Mapping: "server_host",
				},
			},
			Length:   64,
			Primary:  false,
			Type:     model.ResourcePropertyType_TYPE_STRING,
			Required: true,
		},
		{
			Name: "serverPort",
			SourceConfig: &model.ResourceProperty_Mapping{
				Mapping: &model.ResourcePropertyMappingConfig{
					Mapping: "server_port",
				},
			},
			Length:   64,
			Primary:  false,
			Type:     model.ResourcePropertyType_TYPE_INT32,
			Required: true,
		},
		{
			Name: "operations",
			SourceConfig: &model.ResourceProperty_Mapping{
				Mapping: &model.ResourcePropertyMappingConfig{
					Mapping: "operations",
				},
			},
			Type: model.ResourcePropertyType_TYPE_OBJECT,
		},
	},
	SecurityContext: securityContextDisallowAll,
}
