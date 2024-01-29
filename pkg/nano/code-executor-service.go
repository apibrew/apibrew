package nano

import (
	"encoding/base64"
	"fmt"
	"github.com/apibrew/apibrew/pkg/nano/abs"
	"github.com/apibrew/apibrew/pkg/nano/addons"
	"github.com/apibrew/apibrew/pkg/service"
	backend_event_handler "github.com/apibrew/apibrew/pkg/service/backend-event-handler"
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
	log "github.com/sirupsen/logrus"
	"time"
)

type codeExecutorService struct {
	container           service.Container
	backendEventHandler backend_event_handler.BackendEventHandler
	codeContext         map[string]*codeExecutionContext
	globalObject        *globalObject
}

func (s codeExecutorService) NewVm(options abs.VmOptions) (*goja.Runtime, error) {
	vm := goja.New()
	vm.SetFieldNameMapper(goja.UncapFieldNameMapper())

	registry := new(require.Registry) // this can be shared by multiple runtimes

	runtime := goja.New()
	registry.Enable(runtime)

	cec := &codeExecutionContext{}
	err := s.registerBuiltIns("", vm, cec)

	if err != nil {
		return nil, err
	}

	return vm, nil
}

func (s codeExecutorService) GetContainer() service.Container {
	return s.container
}

func (s codeExecutorService) GetBackendEventHandler() backend_event_handler.BackendEventHandler {
	return s.backendEventHandler
}

func (s codeExecutorService) GetGlobalObject() abs.GlobalObject {
	return s.globalObject
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

	registry := new(require.Registry) // this can be shared by multiple runtimes

	runtime := goja.New()
	registry.Enable(runtime)

	cec := &codeExecutionContext{}
	err = s.registerBuiltIns(code.Name, vm, cec)

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
	cec := s.codeContext[code.Name]
	if len(cec.handlerIds) > 0 {
		for _, handlerId := range s.codeContext[code.Name].handlerIds {
			s.backendEventHandler.UnRegisterHandler(backend_event_handler.Handler{
				Id: handlerId,
			})
		}
	}

	for _, cancelFn := range cec.closeHandlers {
		cancelFn()
	}

	return nil
}

func (s codeExecutorService) registerBuiltIns(codeName string, vm *goja.Runtime, cec *codeExecutionContext) error {
	if err := addons.Register(vm, cec, s, codeName, s.container); err != nil {
		return err
	}

	if err := vm.Set("global", s.globalObject); err != nil {
		return err
	}

	if err := s.registerTimeoutFunctions(vm, cec); err != nil {
		return err
	}

	return nil
}

func (s codeExecutorService) registerTimeoutFunctions(vm *goja.Runtime, cec *codeExecutionContext) error {
	if err := vm.Set("setTimeout", s.setTimeoutFn(cec)); err != nil {
		return err
	}

	if err := vm.Set("clearTimeout", s.clearTimeoutFn(cec)); err != nil {
		return err
	}

	if err := vm.Set("setInterval", s.setIntervalFn(cec)); err != nil {
		return err
	}

	if err := vm.Set("clearInterval", s.clearIntervalFn(cec)); err != nil {
		return err
	}

	if err := vm.Set("sleep", s.sleepFn(cec)); err != nil {
		return err
	}

	return nil
}

func (s codeExecutorService) setTimeoutFn(cec *codeExecutionContext) func(fn func(), duration int64) func() {
	return func(fn func(), duration int64) func() {
		cancel := make(chan struct{})
		cancelFn := func() {
			close(cancel)
		}
		cec.closeHandlers = append(cec.closeHandlers, cancelFn) // fixme (potential memory leak)

		go func() {
			defer func() {
				if r := recover(); r != nil {
					log.Warn(r)
				}
			}()
			select {
			case <-time.After(time.Duration(duration) * time.Millisecond):
				fn()
			case <-cancel:
				// Cancel the timeout
			}
		}()
		return cancelFn
	}
}

func (s codeExecutorService) clearTimeoutFn(cec *codeExecutionContext) func(clearFn func()) {
	return func(clearFn func()) {
		defer func() {
			if r := recover(); r != nil {
				log.Warn(r)
			}
		}()
		clearFn()
	}
}

func (s codeExecutorService) setIntervalFn(cec *codeExecutionContext) func(fn func(), duration int64) func() {
	return func(fn func(), duration int64) func() {
		cancel := make(chan struct{})
		cancelFn := func() {
			close(cancel)
		}
		cec.closeHandlers = append(cec.closeHandlers, cancelFn) // fixme (potential memory leak)

		go func() {
		Loop:
			for {
				defer func() {
					if r := recover(); r != nil {
						log.Warn(r)
					}
				}()
				select {
				case <-time.After(time.Duration(duration) * time.Millisecond):
					fn()
				case <-cancel:
					break Loop
				}
			}
		}()
		return cancelFn
	}
}

func (s codeExecutorService) clearIntervalFn(cec *codeExecutionContext) func(clearFn func()) {
	return func(clearFn func()) {
		defer func() {
			if r := recover(); r != nil {
				log.Warn(r)
			}
		}()
		clearFn()
	}
}

func (s codeExecutorService) sleepFn(cec *codeExecutionContext) func(duration int32) {
	return func(duration int32) {
		time.Sleep(time.Duration(duration) * time.Millisecond)
	}
}

func newCodeExecutorService(container service.Container, backendEventHandler backend_event_handler.BackendEventHandler) *codeExecutorService {
	return &codeExecutorService{container: container, backendEventHandler: backendEventHandler, codeContext: make(map[string]*codeExecutionContext), globalObject: newGlobalObject()}
}
