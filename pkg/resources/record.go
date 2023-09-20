package resources

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resources/special"
	"github.com/apibrew/apibrew/pkg/service/annotations"
)

var RecordResource = &model.Resource{
	Name:      "Record",
	Namespace: "system",
	Virtual:   true,
	Properties: []*model.ResourceProperty{
		special.IdProperty,
		{
			Name:     "properties",
			Type:     model.ResourceProperty_OBJECT,
			Required: true,
		},
		{
			Name: "packedProperties",
			Type: model.ResourceProperty_LIST,
			Item: &model.ResourceProperty{
				Type: model.ResourceProperty_OBJECT,
			},
		},
	},

	Annotations: map[string]string{
		//annotations.SelfContainedProperty: "properties",
		annotations.RestApiDisabled: annotations.Enabled,
	},
}
