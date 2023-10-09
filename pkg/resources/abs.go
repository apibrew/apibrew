package resources

import (
	"github.com/apibrew/apibrew/pkg/model"
	"google.golang.org/protobuf/proto"
)

var resourceTypeMap = map[*model.Resource]proto.Message{
	ResourceResource: &model.Resource{},
}

func GetAllSystemResources() []*model.Resource {
	return []*model.Resource{
		NamespaceResource,
		UserResource,
		DataSourceResource,
		ExtensionResource,
		ResourceResource,
		RoleResource,
		PermissionResource,
		RecordResource,
		AuditLogResource,
		ResourceActionResource,
	}
}

func GetSystemResourceType(resource *model.Resource) proto.Message {
	return resourceTypeMap[resource]
}
