package formats

import (
	"github.com/apibrew/apibrew/pkg/resource_model"
	"github.com/apibrew/apibrew/pkg/service/annotations"
)

func FixResourceForApply(resource *resource_model.Resource) *resource_model.Resource {
	resource.Id = nil
	resource.AuditData = nil
	resource.Version = 0

	for name, property := range resource.Properties {
		// if property has special annotation, remove it
		if !annotations.IsEnabled(property, annotations.SpecialProperty) {
			delete(resource.Properties, name)
		}
	}

	return resource
}
