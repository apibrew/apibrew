package resources

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resources/special"
)

var RecordResource = &model.Resource{
	Name:      "Record",
	Namespace: "system",
	Virtual:   true,
	Properties: []*model.ResourceProperty{
		special.IdProperty,
		{
			Name: "properties",
			Type: model.ResourceProperty_MAP,
			Item: &model.ResourceProperty{
				Type: model.ResourceProperty_OBJECT,
			},
		},
		{
			Name: "packedProperties",
			Type: model.ResourceProperty_LIST,
			Item: &model.ResourceProperty{
				Type: model.ResourceProperty_OBJECT,
			},
		},
	},
}
