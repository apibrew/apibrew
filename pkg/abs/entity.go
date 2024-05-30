package abs

import (
	"github.com/apibrew/apibrew/pkg/model"
	"google.golang.org/protobuf/types/known/structpb"
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
	ToProperties(entity Entity) map[string]*structpb.Value
	FromProperties(properties map[string]*structpb.Value) Entity
}
