package resources

import (
	"github.com/tislib/data-handler/pkg/model"
)

var UserResource = &model.Resource{
	Name:      "user",
	Namespace: "system",
	DataType:  model.DataType_SYSTEM,
	SourceConfig: &model.ResourceSourceConfig{
		DataSource: "system",
		Entity:     "user",
	},
	Properties: []*model.ResourceProperty{
		{
			Name: "username",
			SourceConfig: &model.ResourceProperty_Mapping{
				Mapping: &model.ResourcePropertyMappingConfig{
					Mapping:        "username",
					SourceDef:      "",
					AutoGeneration: 0,
				},
			},
			Primary:  false,
			Type:     model.ResourcePropertyType_TYPE_STRING,
			Length:   256,
			Required: true,
			Unique:   true,
		},
		{
			Name: "password",
			SourceConfig: &model.ResourceProperty_Mapping{
				Mapping: &model.ResourcePropertyMappingConfig{
					Mapping: "password",
				},
			},
			Type:     model.ResourcePropertyType_TYPE_STRING,
			Length:   256,
			Required: true,
		},
		securityContextProperty,
		{
			Name: "details",
			SourceConfig: &model.ResourceProperty_Mapping{
				Mapping: &model.ResourcePropertyMappingConfig{
					Mapping: "details",
				},
			},
			Type: model.ResourcePropertyType_TYPE_OBJECT,
		},
	},
	SecurityContext: securityContextDisallowAll,
}
