package nano

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/service"
	backend_event_handler "github.com/apibrew/apibrew/pkg/service/backend-event-handler"
	"github.com/dop251/goja"
	log "github.com/sirupsen/logrus"
)

type httpObject struct {
	goja.DynamicObject
	container service.Container
	resource  *model.Resource

	Fire   func(event map[string]interface{})       `json:"fire"`
	Listen func(func(event map[string]interface{})) `json:"listen"`

	vm                  *goja.Runtime
	cec                 *codeExecutionContext
	backendEventHandler backend_event_handler.BackendEventHandler
}

func (h *httpObject) get() {
	log.Print("Get called")
}

func NewHttpObject(vm *goja.Runtime) goja.Value {
	return vm.ToValue(&httpObject{})
}
