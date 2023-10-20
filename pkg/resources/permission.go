package resources

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resources/special"
	"github.com/apibrew/apibrew/pkg/service/annotations"
	"github.com/apibrew/apibrew/pkg/util"
	"google.golang.org/protobuf/types/known/structpb"
)

var PermissionResource = &model.Resource{
	Name:        "Permission",
	Title:       util.Pointer("Permission"),
	Description: util.Pointer("Permission is a resource that defines the access control rules for resources for users."),
	Namespace:   "system",
	SourceConfig: &model.ResourceSourceConfig{
		DataSource: "system",
		Entity:     "permission",
	},
	Types: []*model.ResourceSubType{
		special.AuditDataSubType,
	},
	Properties: []*model.ResourceProperty{
		special.IdProperty,
		special.VersionProperty,
		special.AuditProperty,
		{
			Name:         "namespace",
			Title:        util.Pointer("Namespace"),
			Description:  util.Pointer(`The namespace(name) of the resource. If given it will be used to match the resource by namespace.`),
			Type:         model.ResourceProperty_STRING,
			Length:       255,
			ExampleValue: structpb.NewStringValue("default"),
		},
		{
			Name:         "resource",
			Title:        util.Pointer("Resource"),
			Description:  util.Pointer(`The name of the resource. If given it will be used to match the resource by name.`),
			Type:         model.ResourceProperty_STRING,
			Length:       255,
			ExampleValue: structpb.NewStringValue("Book"),
		},
		{
			Name:  "property",
			Title: util.Pointer("Property"),
			Description: util.Pointer(`The name of the property.
property and propertyValue are used to match the resource by property value. If record matches property => propertyValue then the permission will be considered. If not, it will be ignored.
Besides that we also have propertyMode which indicate how to match the property value.
If propertyMode is PROPERTY_MATCH_ONLY then only the given property is allowed to be updated, if any other property is sent and not matching by any permission, it will cause an error.
Like for example, you want user to update only title property of resource and to not able to update any other property.
But PROPERTY_MATCH_ANY means that if any of the property is matching then the permission will be considered. It is more useful for owner matching or etc.
For example you want to allow records where owner is user own, etc.
`),
			Type:         model.ResourceProperty_STRING,
			Length:       255,
			Required:     false,
			ExampleValue: structpb.NewStringValue("author"),
		},
		{
			Name:         "propertyValue",
			Title:        util.Pointer("Property Value"),
			Description:  util.Pointer(`The value of the property. It is used by combination with property, please see the description of property.`),
			Type:         model.ResourceProperty_STRING,
			Length:       255,
			Required:     false,
			ExampleValue: structpb.NewStringValue("John Doe"),
		},
		{
			Name:        "propertyMode",
			Title:       util.Pointer("Property Mode"),
			Description: util.Pointer(`The mode of the property. It is used by combination with property and property value, please see the description of property.`),
			Type:        model.ResourceProperty_ENUM,
			Length:      255,
			EnumValues: []string{
				"PROPERTY_MATCH_ONLY",
				"PROPERTY_MATCH_ANY",
			},
		},
		{
			Name:         "operation",
			Title:        util.Pointer("Operation"),
			Description:  util.Pointer(`The operation of the permission. It is used to match the operation of the request. If given it will be used to match the operation of the request.`),
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
			ExampleValue: structpb.NewStringValue("READ"),
		},
		{
			Name:        "recordIds",
			Title:       util.Pointer("Record Ids"),
			Description: util.Pointer(`The ids of the records. It is used to match the record ids of the request. If you want to match only specific records, otherwise all records will be considered`),
			Type:        model.ResourceProperty_LIST,
			Item: &model.ResourceProperty{
				Type: model.ResourceProperty_STRING,
			},
		},
		{
			Name:        "before",
			Title:       util.Pointer("Before"),
			Description: util.Pointer(`The timestamp before which the permission is valid. If given it will be used to match the timestamp of the request.`),
			Type:        model.ResourceProperty_TIMESTAMP,
		},
		{
			Name:        "after",
			Title:       util.Pointer("After"),
			Description: util.Pointer(`The timestamp after which the permission is valid. If given it will be used to match the timestamp of the request.`),
			Type:        model.ResourceProperty_TIMESTAMP,
		},
		{
			Name:        "user",
			Title:       util.Pointer("User"),
			Description: util.Pointer(`The user who has the permission. If given it will be used to match the user of the request. It is ignored by default, because if permissions is set through User this property is overrides and auto-populated by system`),
			Type:        model.ResourceProperty_REFERENCE,
			Reference: &model.Reference{
				Namespace: UserResource.Namespace,
				Resource:  UserResource.Name,
			},
		},
		{
			Name:        "role",
			Title:       util.Pointer("Role"),
			Description: util.Pointer(`The role who has the permission. If given it will be used to match the role of the request. It is ignored by default, because if permissions is set through Role this property is overrides and auto-populated by system`),
			Type:        model.ResourceProperty_REFERENCE,
			Reference: &model.Reference{
				Namespace: RoleResource.Namespace,
				Resource:  RoleResource.Name,
			},
		},
		{
			Name:        "permit",
			Title:       util.Pointer("Permit"),
			Description: util.Pointer(`The permit of the permission. If permission is matched, this property is judging field to indicate that if operation is allowed or not`),
			Required:    true,
			Length:      255,
			Type:        model.ResourceProperty_ENUM,
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
