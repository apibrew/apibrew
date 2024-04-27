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

type EntityMapper[Entity interface{}] interface {
	New() Entity
	ResourceIdentity() ResourceIdentity
	ToRecord(entity Entity) *model.Record
	FromRecord(record *model.Record) Entity
	ToProperties(entity Entity) map[string]*structpb.Value
	FromProperties(properties map[string]*structpb.Value) Entity
}
