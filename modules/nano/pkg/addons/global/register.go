package global

import (
	"github.com/apibrew/apibrew/modules/nano/pkg/abs"
	"github.com/dop251/goja"
)

func Register(vm *goja.Runtime, s abs.CodeExecutorService) error {
	return vm.Set("global", s.GetGlobalObject())
}
