package execute

import (
	"github.com/dop251/goja"
	log "github.com/sirupsen/logrus"
)

func executeFn(level log.Level) interface{} {
	log.Println("Called")

	return nil
}

func Register(vm *goja.Runtime) error {
	return vm.Set("execute", executeFn)
}
