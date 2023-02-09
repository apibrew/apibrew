package resources

import (
	"github.com/tislib/data-handler/pkg/model"
)

var NamespaceResource = &model.Resource{
	Name:      "namespace",
	Namespace: "system",
	DataType:  model.DataType_SYSTEM,
	SourceConfig: &model.ResourceSourceConfig{
		DataSource: "system",
		Entity:     "namespace",
	},
	Properties: []*model.ResourceProperty{
		{
			Name: "name",

			Mapping:  "name",
			Primary:  false,
			Type:     model.ResourcePropertyType_TYPE_STRING,
			Length:   256,
			Required: true,
			Unique:   true,
		},
		{
			Name: "description",

			Mapping:  "description",
			Primary:  false,
			Type:     model.ResourcePropertyType_TYPE_STRING,
			Length:   256,
			Required: false,
		},
		{
			Name: "details",

			Mapping:  "details",
			Primary:  false,
			Type:     model.ResourcePropertyType_TYPE_OBJECT,
			Required: false,
		},
		securityContextProperty,
	},
	SecurityContext: securityContextDisallowAll,
}
