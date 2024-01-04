package lambda

import (
	"github.com/apibrew/apibrew/pkg/nano/abs"
	"github.com/dop251/goja"
)

func Register(vm *goja.Runtime, cec abs.CodeExecutionContext, s abs.CodeExecutorService) error {
	return vm.Set("lambda", lambdaFn(s.GetContainer(), vm, cec, s.GetBackendEventHandler()))
}
