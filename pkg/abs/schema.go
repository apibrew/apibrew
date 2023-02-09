package abs

import "github.com/tislib/data-handler/pkg/model"

type Schema struct {
	Resources                    []*model.Resource
	ResourceByNamespaceSlashName map[string]*model.Resource
}
