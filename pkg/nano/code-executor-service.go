package nano

import (
	"encoding/base64"
	"fmt"
	"github.com/apibrew/apibrew/pkg/service"
	"github.com/dop251/goja"
	log "github.com/sirupsen/logrus"
)

type codeExecutorService struct {
	container service.Container
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

	err = s.registerBuiltIns(code, vm)

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
	return nil
}

func (s codeExecutorService) registerBuiltIns(code *Code, vm *goja.Runtime) error {
	err := vm.Set("console", newConsoleObject(code.Name))

	if err != nil {
		return err
	}

	err = vm.Set("resource", resourceFn(s.container))

	if err != nil {
		return err
	}

	return nil
}

func newCodeExecutorService(container service.Container) *codeExecutorService {
	return &codeExecutorService{container: container}
}
