package plugin

import (
	"github.com/tislib/data-handler/pkg/service/handler"
)

const MetaDataKey = "MetaDataKey"

type HandlerRegistration struct {
	Handler  *handler.BaseHandler
	Selector handler.EventSelector
}

type MetaData struct {
	Handlers []HandlerRegistration
}
