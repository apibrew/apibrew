package resources

import (
	"github.com/apibrew/apibrew/pkg/model"
	"google.golang.org/protobuf/proto"
)

var resourceTypeMap = map[*model.Resource]proto.Message{
	NamespaceResource:  &model.Namespace{},
	UserResource:       &model.User{},
	DataSourceResource: &model.DataSource{},
	ExtensionResource:  &model.Extension{},
	ResourceResource:   &model.Resource{},
}

func GetAllSystemResources() []*model.Resource {
	return []*model.Resource{
		NamespaceResource, UserResource, DataSourceResource, ExtensionResource, ResourceResource,
		RoleResource,
	}
}

func GetSystemResourceType(resource *model.Resource) proto.Message {
	return resourceTypeMap[resource]
}
