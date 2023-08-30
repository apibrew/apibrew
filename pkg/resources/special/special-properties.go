package special

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/service/annotations"
)

var IdProperty = &model.ResourceProperty{
	Name:      "id",
	Type:      model.ResourceProperty_UUID,
	Required:  true,
	Immutable: true,
	Annotations: map[string]string{
		annotations.SpecialProperty: "true",
		annotations.PrimaryProperty: "true",
	},
}

var VersionProperty = &model.ResourceProperty{
	Name:     "version",
	Type:     model.ResourceProperty_INT32,
	Required: true,
	Annotations: map[string]string{
		annotations.SpecialProperty:     annotations.Enabled,
		annotations.AllowEmptyPrimitive: annotations.Enabled,
	},
}

var AuditPropertyCreatedBy = &model.ResourceProperty{
	Name:      "createdBy",
	Type:      model.ResourceProperty_STRING,
	Length:    256,
	Required:  true,
	Immutable: true,
	Annotations: map[string]string{
		annotations.SpecialProperty: "true",
	},
}

var AuditPropertyUpdatedBy = &model.ResourceProperty{
	Name:     "updatedBy",
	Type:     model.ResourceProperty_STRING,
	Length:   256,
	Required: false,
	Annotations: map[string]string{
		annotations.SpecialProperty: "true",
	},
}

var AuditPropertyCreatedOn = &model.ResourceProperty{
	Name:      "createdOn",
	Type:      model.ResourceProperty_TIMESTAMP,
	Required:  true,
	Immutable: true,
	Annotations: map[string]string{
		annotations.SpecialProperty: "true",
	},
}

var AuditPropertyUpdatedOn = &model.ResourceProperty{
	Name:     "updatedOn",
	Type:     model.ResourceProperty_TIMESTAMP,
	Required: false,
	Annotations: map[string]string{
		annotations.SpecialProperty: "true",
	},
}

var AuditProperties = []*model.ResourceProperty{
	AuditPropertyCreatedBy,
	AuditPropertyUpdatedBy,
	AuditPropertyCreatedOn,
	AuditPropertyUpdatedOn,
}

var AnnotationsProperty = &model.ResourceProperty{
	Name:      "annotations",
	Type:      model.ResourceProperty_MAP,
	Required:  false,
	Immutable: false,
	Item: &model.ResourceProperty{
		Type: model.ResourceProperty_STRING,
	},
	Annotations: map[string]string{
		annotations.SpecialProperty: "true",
	},
}

func IsIdProperty(property *model.ResourceProperty) bool {
	return property.Name == IdProperty.Name && property.Type == IdProperty.Type
}
