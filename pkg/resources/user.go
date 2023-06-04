package resources

import (
	"github.com/apibrew/apibrew/pkg/helper/proto"
	"github.com/apibrew/apibrew/pkg/model"
)

var helper = new(proto.ResourceHelper)

var UserResource = helper.ProtoToResource(new(model.User).ProtoReflect())
