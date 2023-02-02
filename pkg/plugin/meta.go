package plugin

import (
	"github.com/tislib/data-handler/pkg/abs"
	"github.com/tislib/data-handler/pkg/service/handler"
)

const MetaDataKey = "MetaDataKey"

type HandlerRegistration struct {
	Handler  *handler.BaseHandler
	Selector handler.EventSelector
}

type BackendRegistration struct {
	backend abs.Backend
	name    string
}

type MetaData struct {
	Handlers []HandlerRegistration
}
