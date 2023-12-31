package docs

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resources"
	"github.com/apibrew/apibrew/pkg/util"
)

func (s *openApiBuilder) getResourceFQN(resource *model.Resource) string {
	if resource == resources.ResourceResource {
		return "resources"
	}

	if resource.Namespace == "default" {
		return util.PathSlug(resource.Name)
	} else {
		return util.PathSlug(resource.Namespace) + "-" + util.PathSlug(resource.Name)
	}
}
