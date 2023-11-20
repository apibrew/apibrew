package abs

import "github.com/apibrew/apibrew/pkg/model"

type ResourceLike interface {
	GetProperties() map[string]*model.ResourceProperty
	GetTypes() []*model.ResourceSubType
}
