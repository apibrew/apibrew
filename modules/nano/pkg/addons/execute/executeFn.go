package execute

import (
	"context"
	"github.com/apibrew/apibrew/modules/nano/pkg/abs"
	"github.com/apibrew/apibrew/modules/nano/pkg/addons/util"
	"github.com/dop251/goja"
)

type Params struct {
	This    goja.Value             `json:"this"`
	Args    map[string]interface{} `json:"args"`
	Isolate bool                   `json:"isolate"`
}

func executeFn(vm *goja.Runtime, s abs.CodeExecutorService) func(script string, params Params) interface{} {
	return func(script string, params Params) interface{} {

		if params.Isolate {
			if (len(params.Args)) > 0 {
				util.ThrowError(vm, "Isolated execution does not support arguments")
			}

			res, err := s.RunInlineScript(context.TODO(), "execute", script)

			if err != nil {
				util.ThrowError(vm, err.Error())
			}

			return res
		} else {
			for key, value := range params.Args {
				if err := vm.Set(key, value); err != nil {
					util.ThrowError(vm, err.Error())
				}
			}

			value, err := vm.RunString(script)

			if err != nil {
				util.ThrowError(vm, err.Error())
			}

			return value.Export()
		}
	}
}

func Register(vm *goja.Runtime, s abs.CodeExecutorService) error {
	if err := vm.Set("execute", executeFn(vm, s)); err != nil {
		return err
	}

	return nil
}
