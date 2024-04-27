package abs

import (
	"github.com/apibrew/apibrew/pkg/model"
)

type PropertyWithPath struct {
	Path     string
	Property *model.ResourceProperty
}

type Schema struct {
	Resources                    []*model.Resource
	ResourceByNamespaceSlashName map[string]*model.Resource
	ResourceBySlug               map[string]*model.Resource
	ResourcePropertiesByType     map[string]map[model.ResourceProperty_Type][]PropertyWithPath
	ResourcePropertyPaths        map[string]map[string]bool
}
