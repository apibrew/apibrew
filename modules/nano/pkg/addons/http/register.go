package http

import (
	"github.com/dop251/goja"
)

func Register(vm *goja.Runtime) error {
	if err := vm.Set("http", &httpObject{vm: vm}); err != nil {
		return err
	}

	return nil
}
