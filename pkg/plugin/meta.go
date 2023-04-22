package plugin

import (
	"github.com/tislib/apibrew/pkg/service/handler"
)

const MetaDataKey = "MetaDataKey"

type HandlerRegistration struct {
	Handler  *handler.BaseHandler
	Selector handler.EventSelector
}

type MetaData struct {
	Handlers []HandlerRegistration
}
