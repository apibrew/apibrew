package resources

import (
	"github.com/apibrew/apibrew/pkg/helper/protohelper"
	"github.com/apibrew/apibrew/pkg/model"
)

var helper = new(protohelper.ResourceHelper)

var UserResource = helper.ProtoToResource(new(model.User).ProtoReflect())
var RoleResource = helper.ProtoToResource(new(model.Role).ProtoReflect())
