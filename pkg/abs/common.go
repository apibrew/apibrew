package abs

import "github.com/apibrew/apibrew/pkg/model"

type ResourceLike interface {
	GetProperties() []*model.ResourceProperty
	GetTypes() []*model.ResourceSubType
}
