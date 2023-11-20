package resources

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resources/special"
	"github.com/apibrew/apibrew/pkg/service/annotations"
	"github.com/apibrew/apibrew/pkg/util"
)

var RoleResource = &model.Resource{
	Name:        "Role",
	Title:       util.Pointer("Role"),
	Description: util.Pointer("Role is a resource that defines the access control model. It is used to give permissions to users in a grouped way"),
	Namespace:   "system",
	SourceConfig: &model.ResourceSourceConfig{
		DataSource: "system",
		Entity:     "role",
	},
	Types: []*model.ResourceSubType{
		special.AuditDataSubType,
	},
	Properties: []*model.ResourceProperty{
		special.IdProperty,
		special.VersionProperty,
		special.AuditProperty,
		{
			Name:        "name",
			Title:       util.Pointer("Name"),
			Description: util.Pointer(`The name of the role`),
			Type:        model.ResourceProperty_STRING,
			Length:      256,
			Required:    true,
			Unique:      true,
			Annotations: map[string]string{
				annotations.IsHclLabel: annotations.Enabled,
			},
		},
		{
			Name:        "permissions",
			Description: util.Pointer(`The permissions of the role. It is used to define the access control rules for resources for roles. When you set permissions it is automatically created though Permission Resource. No need to manage it manually`),
			Type:        model.ResourceProperty_LIST,
			Item: &model.ResourceProperty{
				Type: model.ResourceProperty_REFERENCE,
				Reference: &model.Reference{
					Namespace: "system",
					Resource:  "Permission",
				},
				BackReference: &model.BackReference{
					Property: "role",
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
