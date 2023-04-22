package resources

import (
	"github.com/tislib/apibrew/pkg/model"
)

var securityContextProperty = &model.ResourceProperty{
	Name:     "securityContext",
	Mapping:  "security_context",
	Type:     model.ResourceProperty_OBJECT,
	Required: false,
}

var securityContextDisallowAll = &model.SecurityContext{
	Constraints: []*model.SecurityConstraint{
		{
			Operation: model.OperationType_FULL,
			Permit:    model.PermitType_PERMIT_TYPE_REJECT,
		},
	},
}
