package resources

import "github.com/tislib/data-handler/pkg/model"

func GetAllSystemResources() []*model.Resource {
	return []*model.Resource{
		NamespaceResource, UserResource, DataSourceResource, ExtensionResource, ResourceResource,
	}
}
