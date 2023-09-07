package resources

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resources/special"
	"github.com/apibrew/apibrew/pkg/service/annotations"
	"google.golang.org/protobuf/types/known/structpb"
)

var PermissionResource = &model.Resource{
	Name:      "Permission",
	Namespace: "system",
	SourceConfig: &model.ResourceSourceConfig{
		DataSource: "system",
		Entity:     "permission",
	},
	Properties: []*model.ResourceProperty{
		special.IdProperty,
		special.VersionProperty,
		special.AuditProperties[0],
		special.AuditProperties[1],
		special.AuditProperties[2],
		special.AuditProperties[3],
		{
			Name:   "namespace",
			Type:   model.ResourceProperty_STRING,
			Length: 255,
		},
		{
			Name:   "resource",
			Type:   model.ResourceProperty_STRING,
			Length: 255,
		},
		{
			Name:     "property",
			Type:     model.ResourceProperty_STRING,
			Length:   255,
			Required: false,
		},
		{
			Name:     "propertyValue",
			Type:     model.ResourceProperty_STRING,
			Length:   255,
			Required: false,
		},
		{
			Name:   "propertyMode",
			Type:   model.ResourceProperty_ENUM,
			Length: 255,
			EnumValues: []string{
				"PROPERTY_MATCH_ONLY",
				"PROPERTY_MATCH_ANY",
			},
		},
		{
			Name:         "operation",
			Type:         model.ResourceProperty_ENUM,
			Length:       255,
			Required:     true,
			DefaultValue: structpb.NewStringValue("FULL"),
			EnumValues: []string{
				"READ",
				"CREATE",
				"UPDATE",
				"DELETE",
				"FULL",
			},
		},
		{
			Name: "recordIds",
			Type: model.ResourceProperty_LIST,
			Item: &model.ResourceProperty{
				Type: model.ResourceProperty_STRING,
			},
		},
		{
			Name: "before",
			Type: model.ResourceProperty_TIMESTAMP,
		},
		{
			Name: "after",
			Type: model.ResourceProperty_TIMESTAMP,
		},
		{
			Name: "user",
			Type: model.ResourceProperty_REFERENCE,
			Reference: &model.Reference{
				Namespace: UserResource.Namespace,
				Resource:  UserResource.Name,
			},
		},
		{
			Name: "role",
			Type: model.ResourceProperty_REFERENCE,
			Reference: &model.Reference{
				Namespace: RoleResource.Namespace,
				Resource:  RoleResource.Name,
			},
		},
		{
			Name:     "permit",
			Required: true,
			Length:   255,
			Type:     model.ResourceProperty_ENUM,
			EnumValues: []string{
				"ALLOW",
				"REJECT",
			},
		},
		{
			Name: "localFlags",
			Type: model.ResourceProperty_OBJECT,
		},
	},
	Annotations: map[string]string{
		annotations.EnableAudit:  annotations.Enabled,
		annotations.OpenApiGroup: OpenApiMeta,
	},
}
