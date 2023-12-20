package nano

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/service"
	"github.com/apibrew/apibrew/pkg/util"
)

type resourceObject struct {
	container service.Container
	resource  *model.Resource
}

func resourceFn(container service.Container) func(args ...string) *resourceObject {
	resourceService := container.GetResourceService()
	return func(args ...string) *resourceObject {
		var resourceName string
		var namespace string

		if len(args) == 0 || len(args) > 2 {
			panic("resource function needs 1 or 2 parameters")
		}

		resourceName = args[0]
		if len(args) == 1 {
			namespace = "default"
		} else {
			namespace = args[1]
		}

		resource, err := resourceService.GetResourceByName(util.SystemContext, namespace, resourceName)

		if err != nil {
			panic(err)
		}

		return &resourceObject{resource: resource, container: container}
	}
}
