package util

import (
	"github.com/apibrew/apibrew/modules/nano/pkg/abs"
	"github.com/apibrew/apibrew/pkg/service"
	"github.com/dop251/goja"
)

func Register(vm *goja.Runtime, cec abs.CodeExecutionContext, container service.Container) error {
	if err := vm.Set("resourcePath", resourcePath); err != nil {
		return err
	}

	if err := vm.Set("throwError", throwError(vm)); err != nil {
		return err
	}

	return nil
}
