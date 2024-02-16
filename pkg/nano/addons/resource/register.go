package resource

import (
	"github.com/apibrew/apibrew/pkg/nano/abs"
	"github.com/dop251/goja"
)

func Register(vm *goja.Runtime, cec abs.CodeExecutionContext, s abs.CodeExecutorService) error {
	if err := vm.Set("resource", resourceFn(
		s.GetContainer(),
		vm,
		cec,
		s.GetBackendEventHandler(),
		s.GetGlobalObject(),
	)); err != nil {
		return err
	}

	return nil
}
