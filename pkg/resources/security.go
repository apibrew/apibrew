package resources

import (
	"github.com/tislib/data-handler/pkg/model"
)

var securityContextProperty = &model.ResourceProperty{
	Name: "securityContext",
	SourceConfig: &model.ResourceProperty_Mapping{
		Mapping: &model.ResourcePropertyMappingConfig{
			Mapping: "security_context",
		},
	},
	Type:     model.ResourcePropertyType_TYPE_OBJECT,
	Length:   256,
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
