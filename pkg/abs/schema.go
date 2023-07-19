package abs

import (
	"github.com/apibrew/apibrew/pkg/model"
	"google.golang.org/protobuf/types/known/structpb"
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
