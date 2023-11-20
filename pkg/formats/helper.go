package formats

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/service/annotations"
)

func FixResourceForApply(resource *model.Resource) *model.Resource {
	resource.Id = ""
	resource.AuditData = nil
	resource.Version = 0

	for name, property := range resource.Properties {
		// if property has special annotation, remove it
		if annotations.IsEnabled(property, annotations.SpecialProperty) {
			delete(resource.Properties, name)
		}
	}
	return resource
}
