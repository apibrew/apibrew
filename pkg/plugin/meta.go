package plugin

import (
	"github.com/tislib/data-handler/pkg/backend"
	"github.com/tislib/data-handler/pkg/service/handler"
)

const MetaDataKey = "MetaDataKeys"

type HandlerRegistration struct {
	handler  *handler.BaseHandler
	selector handler.EventSelector
}

type BackendRegistration struct {
	backend backend.Backend
	name    string
}

type MetaData struct {
	Handlers []HandlerRegistration
}
