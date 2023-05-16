package resources

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/service/annotations"
)

var UserResource = &model.Resource{
	Name:      "user",
	Namespace: "system",
	SourceConfig: &model.ResourceSourceConfig{
		DataSource: "system",
		Entity:     "user",
	},
	Properties: []*model.ResourceProperty{
		IdProperty,
		VersionProperty,
		AuditProperties[0],
		AuditProperties[1],
		AuditProperties[2],
		AuditProperties[3],
		{
			Name:     "username",
			Mapping:  "username",
			Primary:  false,
			Type:     model.ResourceProperty_STRING,
			Length:   256,
			Required: true,
			Unique:   true,
			Annotations: map[string]string{
				annotations.IsHclLabel: annotations.Enabled,
			},
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
	Annotations: map[string]string{
		annotations.EnableAudit: "true",
	},
}
