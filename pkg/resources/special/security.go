package special

import (
	"github.com/apibrew/apibrew/pkg/model"
)

var SecurityContextProperty = &model.ResourceProperty{
	Name:     "securityContext",
	Mapping:  "security_context",
	Type:     model.ResourceProperty_OBJECT,
	Required: false,
}

var SecurityContextDisallowAll = &model.SecurityContext{
	Constraints: []*model.SecurityConstraint{
		{
			Operation: model.OperationType_FULL,
			Permit:    model.PermitType_PERMIT_TYPE_REJECT,
		},
	},
}
