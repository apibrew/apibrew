package abs

import (
	"github.com/apibrew/apibrew/pkg/model"
	"google.golang.org/protobuf/types/known/structpb"
)

type EntityMapper[Entity interface{}] interface {
	New() Entity
	ToRecord(entity Entity) *model.Record
	FromRecord(record *model.Record) Entity
	ToProperties(entity Entity) map[string]*structpb.Value
	FromProperties(properties map[string]*structpb.Value) Entity
}
