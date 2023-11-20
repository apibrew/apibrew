package resources

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resources/special"
	"github.com/apibrew/apibrew/pkg/service/annotations"
	"github.com/apibrew/apibrew/pkg/util"
)

var UserResource = &model.Resource{
	Name:        "User",
	Title:       util.Pointer("User"),
	Description: util.Pointer("User is a resource that defines the access control model. It is used to authenticate and authorize users."),
	Namespace:   "system",
	SourceConfig: &model.ResourceSourceConfig{
		DataSource: "system",
		Entity:     "user",
	},
	Types: []*model.ResourceSubType{
		special.AuditDataSubType,
	},
	Properties: []*model.ResourceProperty{
		special.IdProperty,
		special.VersionProperty,
		special.AuditProperty,
		{
			Name:     "username",
			Title:    util.Pointer("Username"),
			Type:     model.ResourceProperty_STRING,
			Length:   256,
			Required: true,
			Unique:   true,
			Annotations: map[string]string{
				annotations.IsHclLabel: annotations.Enabled,
			},
		},
		{
			Name:     "password",
			Title:    util.Pointer("Password"),
			Type:     model.ResourceProperty_STRING,
			Length:   256,
			Required: false,
		},
		{
			Name:  "roles",
			Title: util.Pointer("Roles"),
			Type:  model.ResourceProperty_LIST,
			Item: &model.ResourceProperty{
				Type: model.ResourceProperty_REFERENCE,
				Reference: &model.Reference{
					Namespace: RoleResource.Namespace,
					Resource:  RoleResource.Name,
				},
			},
		},
		{
			Name:        "permissions",
			Title:       util.Pointer("Permissions"),
			Description: util.Pointer(`The permissions of the user. It is used to define the access control rules for resources for users. When you set permissions it is automatically created though Permission Resource. No need to manage it manually`),
			Type:        model.ResourceProperty_LIST,
			Item: &model.ResourceProperty{
				Type: model.ResourceProperty_REFERENCE,
				Reference: &model.Reference{
					Namespace: "system",
					Resource:  "Permission",
				},
				BackReference: &model.BackReference{
					Property: "user",
				},
			},
			Required: false,
		},
		{
			Name:        "details",
			Title:       util.Pointer("Details"),
			Description: util.Pointer(`The details of the user. It is used to store additional information about the user.`),
			Type:        model.ResourceProperty_OBJECT,
		},
	},
	Annotations: map[string]string{
		annotations.EnableAudit:  annotations.Enabled,
		annotations.OpenApiGroup: OpenApiMeta,
	},
}
