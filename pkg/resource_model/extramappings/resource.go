package extramappings

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resource_model"
	"github.com/apibrew/apibrew/pkg/resources/mapping"
)

func ResourceTo(resource *model.Resource) *resource_model.Resource {
	if resource == nil {
		return nil
	}
	resourceRec := mapping.ResourceToRecord(resource)
	return resource_model.ResourceMapperInstance.FromRecord(resourceRec)
}

func ResourceFrom(resource *resource_model.Resource) *model.Resource {
	if resource == nil {
		return nil
	}
	resourceRec := resource_model.ResourceMapperInstance.ToRecord(resource)
	return mapping.ResourceFromRecord(resourceRec)
}
