package docs

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resources"
	"github.com/gosimple/slug"
)

func (s *openApiBuilder) getResourceFQN(resource *model.Resource) string {
	if resource == resources.ResourceResource {
		return "resources"
	}

	if resource.Namespace == "default" {
		return slug.Make(resource.Name)
	} else {
		return slug.Make(resource.Namespace + "/" + resource.Name)
	}
}
