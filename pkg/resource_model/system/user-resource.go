package system

import (
	"github.com/apibrew/apibrew/pkg/resource_model"
	"github.com/apibrew/apibrew/pkg/resources"
	"github.com/apibrew/apibrew/pkg/resources/mapping"
)

var UserResourceModel = resource_model.ResourceMapperInstance.FromRecord(mapping.ResourceToRecord(resources.UserResource))
var NamespaceResourceModel = resource_model.ResourceMapperInstance.FromRecord(mapping.ResourceToRecord(resources.NamespaceResource))
