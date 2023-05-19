package resources

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/service/annotations"
)

var ResourcePropertyResource = &model.Resource{
	Name:      "resourceProperty",
	Namespace: "system",
	SourceConfig: &model.ResourceSourceConfig{
		DataSource: "system",
		Entity:     "resource_property",
	},
	Properties: []*model.ResourceProperty{
		IdProperty,
		VersionProperty,
		AuditProperties[0],
		AuditProperties[1],
		AuditProperties[2],
		AuditProperties[3],
		{
			Name:     "name",
			Mapping:  "name",
			Primary:  false,
			Type:     model.ResourceProperty_STRING,
			Length:   256,
			Required: true,
		},
		{
			Name:     "type",
			Mapping:  "type",
			Type:     model.ResourceProperty_INT32,
			Required: true,
		},
		{
			Name:     "typeRef",
			Mapping:  "type_ref",
			Primary:  false,
			Type:     model.ResourceProperty_STRING,
			Length:   256,
			Required: false,
		},
		{
			Name:     "mapping",
			Mapping:  "mapping",
			Type:     model.ResourceProperty_STRING,
			Length:   64,
			Required: true,
		},
		{
			Name:     "primary",
			Mapping:  "primary",
			Type:     model.ResourceProperty_BOOL,
			Required: true,
		},
		{
			Name:     "required",
			Mapping:  "required",
			Type:     model.ResourceProperty_BOOL,
			Required: true,
		},
		{
			Name:     "unique",
			Mapping:  "unique",
			Type:     model.ResourceProperty_BOOL,
			Required: true,
		},
		{
			Name:     "immutable",
			Mapping:  "immutable",
			Type:     model.ResourceProperty_BOOL,
			Required: true,
		},
		{
			Name:     "length",
			Mapping:  "length",
			Type:     model.ResourceProperty_INT32,
			Required: true,
		},
		{
			Name:     "resource",
			Mapping:  "resource",
			Type:     model.ResourceProperty_REFERENCE,
			Required: true,
			Reference: &model.Reference{
				ReferencedResource: ResourceResource.Name,
				Cascade:            true,
			},
		},
		{
			Name:     "item",
			Mapping:  "item",
			Type:     model.ResourceProperty_OBJECT,
			Required: false,
		},
		{
			Name:     "properties",
			Mapping:  "properties",
			Type:     model.ResourceProperty_OBJECT,
			Required: false,
		},
		{
			Name:    "reference_resource",
			Mapping: "reference_resource",
			Type:    model.ResourceProperty_REFERENCE,
			Reference: &model.Reference{
				ReferencedResource: ResourceResource.Name,
				Cascade:            true,
			},
			Required: false,
		},
		{
			Name:     "reference_cascade",
			Mapping:  "reference_cascade",
			Type:     model.ResourceProperty_BOOL,
			Required: false,
		},
		securityContextProperty,
		{
			Name:     "defaultValue",
			Mapping:  "default_value",
			Type:     model.ResourceProperty_OBJECT,
			Required: false,
		},
		{
			Name:     "enumValues",
			Mapping:  "enum_values",
			Type:     model.ResourceProperty_OBJECT,
			Required: false,
		},
		{
			Name:     "exampleValue",
			Mapping:  "example_value",
			Type:     model.ResourceProperty_OBJECT,
			Required: false,
		},
		{
			Name:     "title",
			Mapping:  "title",
			Primary:  false,
			Type:     model.ResourceProperty_STRING,
			Length:   256,
			Required: false,
		},
		{
			Name:     "description",
			Mapping:  "description",
			Primary:  false,
			Type:     model.ResourceProperty_STRING,
			Length:   256,
			Required: false,
		},
		{
			Name:     "annotations",
			Mapping:  "annotations",
			Type:     model.ResourceProperty_OBJECT,
			Required: false,
		},
	},
	Indexes: []*model.ResourceIndex{
		{
			Properties: []*model.ResourceIndexProperty{
				{
					Name: "resource",
				}, {
					Name: "name",
				},
			},
			Unique: true,
		},
	},
	SecurityContext: securityContextDisallowAll,
	Annotations: map[string]string{
		annotations.EnableAudit: "true",
	},
}
