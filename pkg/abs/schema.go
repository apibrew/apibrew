package abs

import (
	"github.com/tislib/apibrew/pkg/model"
	"google.golang.org/protobuf/types/known/structpb"
)

type Schema struct {
	Resources                    []*model.Resource
	ResourceByNamespaceSlashName map[string]*model.Resource
}

type Entity[T any] interface {
	ToRecord() *model.Record
	FromRecord(record *model.Record)
	FromProperties(properties map[string]*structpb.Value)
	ToProperties() map[string]*structpb.Value
	GetResourceName() string
	GetNamespace() string
	Equals(other T) bool
	Same(other T) bool
}
