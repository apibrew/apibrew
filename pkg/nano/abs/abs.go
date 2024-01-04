package abs

import (
	"github.com/apibrew/apibrew/pkg/service"
	backend_event_handler "github.com/apibrew/apibrew/pkg/service/backend-event-handler"
)

type GlobalObject interface {
	Define(name string, value interface{})
	Get(name string) interface{}
}

type CodeExecutorService interface {
	GetContainer() service.Container
	GetBackendEventHandler() backend_event_handler.BackendEventHandler
	GetGlobalObject() GlobalObject
}

type CodeExecutionContext interface {
	AddHandlerId(id string)
}
