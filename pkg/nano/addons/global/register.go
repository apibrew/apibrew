package global

import (
	"github.com/apibrew/apibrew/pkg/nano/abs"
	"github.com/dop251/goja"
)

func Register(vm *goja.Runtime, s abs.CodeExecutorService) error {
	return vm.Set("global", s.GetGlobalObject())
}
