package resources

import (
	"github.com/tislib/data-handler/pkg/model"
)

var ResourceReferenceResource = &model.Resource{
	Name:      "resourceReference",
	Namespace: "system",
	DataType:  model.DataType_SYSTEM,
	SourceConfig: &model.ResourceSourceConfig{
		DataSource: "system",
		Entity:     "resource_reference",
	},
	Properties: []*model.ResourceProperty{
		{
			Name: "propertyName",

			Mapping:  "property_name",
			Primary:  false,
			Type:     model.ResourcePropertyType_TYPE_STRING,
			Length:   256,
			Required: true,
		},
		{
			Name: "referencedResource",

			Mapping:  "referenced_resource",
			Primary:  false,
			Type:     model.ResourcePropertyType_TYPE_STRING,
			Length:   256,
			Required: true,
		},
		{
			Name: "cascade",

			Mapping:  "cascade",
			Primary:  false,
			Type:     model.ResourcePropertyType_TYPE_BOOL,
			Required: true,
		},
		{
			Name: "resource",

			Mapping:  "resource",
			Type:     model.ResourcePropertyType_TYPE_REFERENCE,
			Required: true,
			Reference: &model.Reference{
				ReferencedResource: ResourceResource.Name,
				Cascade:            false,
			},
		},
	},
	SecurityContext: securityContextDisallowAll,
}
