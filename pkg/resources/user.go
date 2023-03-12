package resources

import (
	"github.com/tislib/data-handler/pkg/model"
)

var UserResource = &model.Resource{
	Name:      "user",
	Namespace: "system",
	SourceConfig: &model.ResourceSourceConfig{
		DataSource: "system",
		Entity:     "user",
	},
	Properties: []*model.ResourceProperty{
		{
			Name: "username",

			Mapping:  "username",
			Primary:  false,
			Type:     model.ResourceProperty_STRING,
			Length:   256,
			Required: true,
			Unique:   true,
		},
		{
			Name: "password",

			Mapping:  "password",
			Type:     model.ResourceProperty_STRING,
			Length:   256,
			Required: true,
		},
		securityContextProperty,
		{
			Name: "details",

			Mapping:         "details",
			Type:            model.ResourceProperty_OBJECT,
			SecurityContext: securityContextDisallowAll,
		},
	},
	SecurityContext: securityContextDisallowAll,
}
