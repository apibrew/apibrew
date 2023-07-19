package resources

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resources/special"
	"github.com/apibrew/apibrew/pkg/service/annotations"
	"google.golang.org/protobuf/types/known/structpb"
)

var SecurityConstraintResource = &model.Resource{
	Name:      "SecurityConstraint",
	Namespace: "system",
	SourceConfig: &model.ResourceSourceConfig{
		DataSource: "system",
		Entity:     "security_constraint",
	},
	Properties: []*model.ResourceProperty{
		special.IdProperty,
		special.VersionProperty,
		special.AuditProperties[0],
		special.AuditProperties[1],
		special.AuditProperties[2],
		special.AuditProperties[3],
		{
			Name:    "namespace",
			Mapping: "namespace",
			Type:    model.ResourceProperty_REFERENCE,
			Reference: &model.Reference{
				Namespace: NamespaceResource.Namespace,
				Resource:  NamespaceResource.Name,
			},
		},
		{
			Name:    "resource",
			Mapping: "resource",
			Type:    model.ResourceProperty_REFERENCE,
			Reference: &model.Reference{
				Namespace: ResourceResource.Namespace,
				Resource:  ResourceResource.Name,
			},
		},
		{
			Name:     "property",
			Mapping:  "property",
			Type:     model.ResourceProperty_STRING,
			Length:   255,
			Required: false,
		},
		{
			Name:     "propertyValue",
			Mapping:  "property_value",
			Type:     model.ResourceProperty_STRING,
			Length:   255,
			Required: false,
		},
		{
			Name:    "propertyMode",
			Mapping: "property_mode",
			Type:    model.ResourceProperty_ENUM,
			Length:  255,
			EnumValues: []string{
				"PROPERTY_MATCH_ONLY",
				"PROPERTY_MATCH_ANY",
			},
		},
		{
			Name:         "operation",
			Mapping:      "operation",
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
			Name:    "recordIds",
			Mapping: "recordIds",
			Type:    model.ResourceProperty_LIST,
			Item: &model.ResourceProperty{
				Type: model.ResourceProperty_STRING,
			},
		},
		{
			Name:    "before",
			Mapping: "before",
			Type:    model.ResourceProperty_TIMESTAMP,
		},
		{
			Name:    "after",
			Mapping: "after",
			Type:    model.ResourceProperty_TIMESTAMP,
		},
		{
			Name:    "user",
			Mapping: "user",
			Type:    model.ResourceProperty_REFERENCE,
			Reference: &model.Reference{
				Namespace: UserResource.Namespace,
				Resource:  UserResource.Name,
			},
		},
		{
			Name:    "role",
			Mapping: "role",
			Type:    model.ResourceProperty_REFERENCE,
			Reference: &model.Reference{
				Namespace: RoleResource.Namespace,
				Resource:  RoleResource.Name,
			},
		},
		{
			Name:     "permit",
			Mapping:  "permit",
			Required: true,
			Length:   255,
			Type:     model.ResourceProperty_ENUM,
			EnumValues: []string{
				"ALLOW",
				"REJECT",
			},
		},
		{
			Name:    "localFlags",
			Mapping: "local_flags",
			Type:    model.ResourceProperty_OBJECT,
		},
	},
	Annotations: map[string]string{
		annotations.EnableAudit: "true",
	},
}
