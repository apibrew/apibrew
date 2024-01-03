package nano

import (
	"encoding/base64"
	"fmt"
	"github.com/apibrew/apibrew/pkg/service"
	backend_event_handler "github.com/apibrew/apibrew/pkg/service/backend-event-handler"
	"github.com/dop251/goja"
	log "github.com/sirupsen/logrus"
	"time"
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

	if err = s.registerTimeoutFunctions(code, vm, cec); err != nil {
		return err
	}

	return nil
}

func (s codeExecutorService) registerTimeoutFunctions(code *Code, vm *goja.Runtime, cec *codeExecutionContext) error {
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
