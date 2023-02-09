package resources

import (
	"github.com/tislib/data-handler/pkg/model"
)

var ResourcePropertyResource = &model.Resource{
	Name:      "resourceProperty",
	Namespace: "system",
	DataType:  model.DataType_SYSTEM,
	SourceConfig: &model.ResourceSourceConfig{
		DataSource: "system",
		Entity:     "resource_property",
	},
	Properties: []*model.ResourceProperty{
		{
			Name:     "name",
			Mapping:  "name",
			Primary:  false,
			Type:     model.ResourcePropertyType_TYPE_STRING,
			Length:   256,
			Required: true,
		},
		{
			Name:     "type",
			Mapping:  "type",
			Type:     model.ResourcePropertyType_TYPE_INT32,
			Required: true,
		},
		{
			Name:     "mapping",
			Mapping:  "source_mapping",
			Type:     model.ResourcePropertyType_TYPE_STRING,
			Length:   64,
			Required: true,
		},
		{
			Name:     "sourcePrimary",
			Mapping:  "source_primary",
			Type:     model.ResourcePropertyType_TYPE_BOOL,
			Required: true,
		},
		{
			Name:     "required",
			Mapping:  "required",
			Type:     model.ResourcePropertyType_TYPE_BOOL,
			Required: true,
		},
		{
			Name:     "unique",
			Mapping:  "unique",
			Type:     model.ResourcePropertyType_TYPE_BOOL,
			Required: true,
		},
		{
			Name:     "length",
			Mapping:  "length",
			Type:     model.ResourcePropertyType_TYPE_INT32,
			Required: true,
		},
		{
			Name:     "resource",
			Mapping:  "resource",
			Type:     model.ResourcePropertyType_TYPE_REFERENCE,
			Required: true,
			Reference: &model.Reference{
				ReferencedResource: ResourceResource.Name,
				Cascade:            true,
			},
		},
		{
			Name:     "subType",
			Mapping:  "sub_type",
			Type:     model.ResourcePropertyType_TYPE_INT32,
			Required: false,
		},
		{
			Name:    "reference_resource",
			Mapping: "reference_resource",
			Type:    model.ResourcePropertyType_TYPE_REFERENCE,
			Reference: &model.Reference{
				ReferencedResource: ResourceResource.Name,
				Cascade:            true,
			},
			Required: false,
		},
		{
			Name:     "reference_cascade",
			Mapping:  "reference_cascade",
			Type:     model.ResourcePropertyType_TYPE_BOOL,
			Required: false,
		},
		securityContextProperty,
	},
	SecurityContext: securityContextDisallowAll,
}
