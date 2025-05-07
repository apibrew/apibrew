package handler

import (
	"github.com/apibrew/apibrew/modules/nano/pkg/abs"
	"github.com/dop251/goja"
)

func Register(vm *goja.Runtime, cec abs.CodeExecutionContext, s abs.CodeExecutorService) error {
	if err := vm.Set("handle", handle(vm, cec, s.GetBackendEventHandler())); err != nil {
		return err
	}

	if err := vm.Set("handleExported", handle(vm, cec, s.GetBackendEventHandler())); err != nil {
		return err
	}

	return nil
}
