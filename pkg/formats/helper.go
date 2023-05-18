package formats

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/service/annotations"
)

func FixResourceForApply(resource *model.Resource) *model.Resource {
	resource.Id = ""
	resource.AuditData = nil
	resource.Version = 0

	var newProperties []*model.ResourceProperty
	for _, property := range resource.Properties {
		property.Id = nil

		// if property has special annotation, remove it
		if !annotations.IsEnabled(property, annotations.SpecialProperty) {
			newProperties = append(newProperties, property)
		}
	}

	resource.Properties = newProperties

	return resource
}
