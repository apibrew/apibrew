package util

import (
	"fmt"
	"github.com/dop251/goja"
)

func throwError(vm *goja.Runtime) func(msg string) {
	return func(msg string) {
		ThrowError(vm, msg)
	}
}

func ThrowError(vm *goja.Runtime, msg string) {
	err := vm.NewGoError(fmt.Errorf(msg))
	panic(vm.ToValue(err))
}
