package resources

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resources/special"
	sub_types "github.com/apibrew/apibrew/pkg/resources/sub-types"
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
		sub_types.BooleanExpression,
		sub_types.PairExpression,
		sub_types.RegexMatchExpression,
		sub_types.Expression,
	},
	Properties: map[string]*model.ResourceProperty{
		"id":        special.IdProperty,
		"version":   special.VersionProperty,
		"auditData": special.AuditProperty,
		"namespace": {
			Title:        util.Pointer("Namespace"),
			Description:  util.Pointer(`The namespace(name) of the resource. If given it will be used to match the resource by namespace.`),
			Type:         model.ResourceProperty_STRING,
			Length:       255,
			ExampleValue: structpb.NewStringValue("default"),
		},
		"resource": {
			Title:        util.Pointer("Resource"),
			Description:  util.Pointer(`The name of the resource. If given it will be used to match the resource by name.`),
			Type:         model.ResourceProperty_STRING,
			Length:       255,
			ExampleValue: structpb.NewStringValue("Book"),
		},
		"recordSelector": {
			Type:    model.ResourceProperty_STRUCT,
			TypeRef: util.Pointer(sub_types.BooleanExpression.Name),
		},
		"operation": {
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
		"before": {
			Title:       util.Pointer("Before"),
			Description: util.Pointer(`The timestamp before which the permission is valid. If given it will be used to match the timestamp of the request.`),
			Type:        model.ResourceProperty_TIMESTAMP,
		},
		"after": {
			Title:       util.Pointer("After"),
			Description: util.Pointer(`The timestamp after which the permission is valid. If given it will be used to match the timestamp of the request.`),
			Type:        model.ResourceProperty_TIMESTAMP,
		},
		"user": {
			Title:       util.Pointer("User"),
			Description: util.Pointer(`The user who has the permission. If given it will be used to match the user of the request. It is ignored by default, because if permissions is set through User this property is overrides and auto-populated by system`),
			Type:        model.ResourceProperty_REFERENCE,
			Reference: &model.Reference{
				Namespace: UserResource.Namespace,
				Resource:  UserResource.Name,
			},
		},
		"role": {
			Title:       util.Pointer("Role"),
			Description: util.Pointer(`The role who has the permission. If given it will be used to match the role of the request. It is ignored by default, because if permissions is set through Role this property is overrides and auto-populated by system`),
			Type:        model.ResourceProperty_REFERENCE,
			Reference: &model.Reference{
				Namespace: RoleResource.Namespace,
				Resource:  RoleResource.Name,
			},
		},
		"permit": {
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
		"localFlags": {
			Type: model.ResourceProperty_OBJECT,
		},
	},
	Annotations: map[string]string{
		annotations.EnableAudit:  annotations.Enabled,
		annotations.OpenApiGroup: OpenApiMeta,
	},
}
