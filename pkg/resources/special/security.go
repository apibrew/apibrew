package special

import (
	"github.com/apibrew/apibrew/pkg/model"
	sub_types "github.com/apibrew/apibrew/pkg/resources/sub-types"
)

func StringPointer(str string) *string {
	var pointer = new(string)

	*pointer = str

	return pointer
}

var SecurityConstraintsProperty = &model.ResourceProperty{
	Name:    "securityConstraints",
	Mapping: "security_constraints",
	Type:    model.ResourceProperty_LIST,
	Item: &model.ResourceProperty{
		Type:    model.ResourceProperty_STRUCT,
		TypeRef: StringPointer(sub_types.SecurityConstraint.Name),
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
