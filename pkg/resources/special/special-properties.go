package special

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/service/annotations"
)

var IdProperty = &model.ResourceProperty{
	Name:      "id",
	Type:      model.ResourceProperty_UUID,
	Mapping:   "id",
	Required:  true,
	Immutable: true,
	Primary:   true,
	Annotations: map[string]string{
		annotations.SpecialProperty: "true",
	},
}

var VersionProperty = &model.ResourceProperty{
	Name:     "version",
	Type:     model.ResourceProperty_INT32,
	Mapping:  "version",
	Required: true,
	Primary:  false,
	Annotations: map[string]string{
		annotations.SpecialProperty: "true",
	},
}

var AuditPropertyCreatedBy = &model.ResourceProperty{
	Name:      "createdBy",
	Type:      model.ResourceProperty_STRING,
	Mapping:   "created_by",
	Length:    256,
	Required:  true,
	Immutable: true,
	Primary:   false,
	Annotations: map[string]string{
		annotations.SpecialProperty: "true",
	},
}

var AuditPropertyUpdatedBy = &model.ResourceProperty{
	Name:     "updatedBy",
	Type:     model.ResourceProperty_STRING,
	Mapping:  "updated_by",
	Length:   256,
	Required: false,
	Primary:  false,
	Annotations: map[string]string{
		annotations.SpecialProperty: "true",
	},
}

var AuditPropertyCreatedOn = &model.ResourceProperty{
	Name:      "createdOn",
	Type:      model.ResourceProperty_TIMESTAMP,
	Mapping:   "created_on",
	Required:  true,
	Immutable: true,
	Primary:   false,
	Annotations: map[string]string{
		annotations.SpecialProperty: "true",
	},
}

var AuditPropertyUpdatedOn = &model.ResourceProperty{
	Name:     "updatedOn",
	Type:     model.ResourceProperty_TIMESTAMP,
	Mapping:  "updated_on",
	Required: false,
	Primary:  false,
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
	Mapping:   "annotations",
	Required:  false,
	Immutable: false,
	Primary:   false,
	Item: &model.ResourceProperty{
		Type: model.ResourceProperty_STRING,
	},
	Annotations: map[string]string{
		annotations.SpecialProperty: "true",
	},
}
