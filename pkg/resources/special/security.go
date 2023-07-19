package special

import (
	"github.com/apibrew/apibrew/pkg/model"
)

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
