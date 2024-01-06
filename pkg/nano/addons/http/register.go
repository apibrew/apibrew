package http

import (
	"github.com/dop251/goja"
)

func Register(vm *goja.Runtime) error {
	return vm.Set("http", &httpObject{vm: vm})
}
