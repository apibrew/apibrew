package special

import (
	"github.com/apibrew/apibrew/pkg/model"
)

var SecurityConstraintsProperty = &model.ResourceProperty{
	Name:    "securityConstraints",
	Mapping: "security_constraints",
	Type:    model.ResourceProperty_LIST,
	Item: &model.ResourceProperty{
		Type: model.ResourceProperty_REFERENCE,
		Reference: &model.Reference{
			Namespace: "system",
			Resource:  "SecurityConstraint",
			Cascade:   false,
		},
	},
	Required: false,
}

var rootRoleName = new(string)

func init() {
	*rootRoleName = "root"
}

var SecurityContextDisallowAll = []*model.SecurityConstraint{
	{
		Operation: model.OperationType_FULL,
		Role:      rootRoleName,
		Permit:    model.PermitType_PERMIT_TYPE_ALLOW,
	},
}
