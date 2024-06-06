package abs

import (
	"github.com/apibrew/apibrew/pkg/model"
)

type ResourceIdentity struct {
	Namespace string
	Name      string
}

func (r ResourceIdentity) Type() string {
	if r.Namespace == "" || r.Namespace == "default" {
		return r.Name
	}
	return r.Namespace + "/" + r.Name
}

func GetType(resource *model.Resource) string {
	return ResourceIdentity{Namespace: resource.Namespace, Name: resource.Name}.Type()
}

type EntityMapper[Entity interface{}] interface {
	New() Entity
	ResourceIdentity() ResourceIdentity
	ToRecord(entity Entity) RecordLike
	FromRecord(record RecordLike) Entity
	ToProperties(entity Entity) map[string]interface{}
	FromProperties(properties map[string]interface{}) Entity
}
