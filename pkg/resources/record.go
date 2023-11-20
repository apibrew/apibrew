package resources

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resources/special"
	"github.com/apibrew/apibrew/pkg/service/annotations"
	"github.com/apibrew/apibrew/pkg/util"
)

var RecordResource = &model.Resource{
	Name:        "Record",
	Title:       util.Pointer("Generic Record"),
	Description: util.Pointer("A generic record resource. All Apis are extended from Generic Record resource"),
	Namespace:   "system",
	Virtual:     true,
	Properties: map[string]*model.ResourceProperty{
		"id": special.IdProperty,
		"properties": {
			Title: util.Pointer("Properties"),
			Description: util.Pointer(`The properties of the record. The schema of properties are defined in the resource definition. 
Here you will put the payload corresponding to the resource definition.
`),
			Type:     model.ResourceProperty_OBJECT,
			Required: true,
		},
		"packedProperties": {
			Type: model.ResourceProperty_LIST,
			Item: &model.ResourceProperty{
				Type: model.ResourceProperty_OBJECT,
			},
			Annotations: map[string]string{
				annotations.OpenApiHide: "true",
			},
		},
	},

	Annotations: map[string]string{
		//annotations.SelfContainedProperty: "properties",
		annotations.RestApiDisabled: annotations.Enabled,
	},
}
