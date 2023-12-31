package nano

import (
	"encoding/base64"
	"fmt"
	"github.com/apibrew/apibrew/pkg/service"
	backend_event_handler "github.com/apibrew/apibrew/pkg/service/backend-event-handler"
	"github.com/dop251/goja"
	log "github.com/sirupsen/logrus"
)

type codeExecutorService struct {
	container           service.Container
	backendEventHandler backend_event_handler.BackendEventHandler
	codeContext         map[string]*codeExecutionContext
	globalObject        *globalObject
}

func (s codeExecutorService) registerCode(code *Code) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()
	decodedBytes, err := base64.StdEncoding.DecodeString(code.Content)

	if err == nil {
		code.Content = string(decodedBytes)
	}

	log.Debug("Registering code: " + code.Name + " / " + code.Content)

	vm := goja.New()
	vm.SetFieldNameMapper(goja.UncapFieldNameMapper())

	cec := &codeExecutionContext{}

	err = s.registerBuiltIns(code, vm, cec)

	s.codeContext[code.Name] = cec

	if err != nil {
		return err
	}

	_, err = vm.RunString(code.Content)
	if err != nil {
		return err
	}

	return nil
}

func (s codeExecutorService) updateCode(code *Code) error {
	if err := s.unRegisterCode(code); err != nil {
		return err
	}

	if err := s.registerCode(code); err != nil {
		return err
	}

	return nil
}

func (s codeExecutorService) unRegisterCode(code *Code) error {
	if len(s.codeContext[code.Name].handlerIds) > 0 {
		for _, handlerId := range s.codeContext[code.Name].handlerIds {
			s.backendEventHandler.UnRegisterHandler(backend_event_handler.Handler{
				Id: handlerId,
			})
		}
	}
	return nil
}

func (s codeExecutorService) registerBuiltIns(code *Code, vm *goja.Runtime, cec *codeExecutionContext) error {
	err := vm.Set("console", newConsoleObject(code.Name, vm, cec))

	if err != nil {
		return err
	}

	err = vm.Set("resource", resourceFn(s.container, vm, cec, s.backendEventHandler, s.globalObject))

	if err != nil {
		return err
	}

	err = vm.Set("lambda", lambdaFn(s.container, vm, cec, s.backendEventHandler))

	if err != nil {
		return err
	}

	err = vm.Set("http", NewHttpObject(vm))

	if err != nil {
		return err
	}

	err = vm.Set("global", s.globalObject)

	if err != nil {
		return err
	}

	return nil
}

func newCodeExecutorService(container service.Container, backendEventHandler backend_event_handler.BackendEventHandler) *codeExecutorService {
	return &codeExecutorService{container: container, backendEventHandler: backendEventHandler, codeContext: make(map[string]*codeExecutionContext), globalObject: newGlobalObject()}
}
