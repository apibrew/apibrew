package special

import (
	"github.com/apibrew/apibrew/pkg/model"
)

var SecurityConstraintsProperty = &model.ResourceProperty{
	Name:    "securityConstraints",
	Mapping: "security_constraints",
	Type:    model.ResourceProperty_LIST,
	Item: &model.ResourceProperty{
		Type: model.ResourceProperty_OBJECT,
	},
	Required: false,
}

var SecurityContextDisallowAll = []*model.SecurityConstraint{
	{
		Operation:   model.OperationType_FULL,
		Role:        "root",
		Permit:      model.PermitType_PERMIT_TYPE_ALLOW,
		RequirePass: true,
	},
}
