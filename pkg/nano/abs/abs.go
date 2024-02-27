package abs

import (
	"context"
	"github.com/apibrew/apibrew/pkg/service"
	backend_event_handler "github.com/apibrew/apibrew/pkg/service/backend-event-handler"
	"github.com/dop251/goja"
)

type GlobalObject interface {
	Define(name string, value interface{})
	Get(name string) interface{}
}

type VmOptions struct {
}

type CodeExecutorService interface {
	GetContainer() service.Container
	GetBackendEventHandler() backend_event_handler.BackendEventHandler
	GetGlobalObject() GlobalObject
	NewVm(options VmOptions) (*goja.Runtime, error)
}

type CodeExecutionContext interface {
	AddHandlerId(id string)
	RemoveHandlerId(id string)
	Context() context.Context
	GetCodeIdentifier() string
}
