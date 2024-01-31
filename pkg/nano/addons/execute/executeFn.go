package execute

import (
	"github.com/apibrew/apibrew/pkg/nano/abs"
	"github.com/dop251/goja"
	"strings"
)

type Params struct {
	This    goja.Value             `json:"this"`
	Args    map[string]interface{} `json:"args"`
	Isolate bool                   `json:"isolate"`
	Vm      *goja.Runtime          `json:"vm"`
}

func executeFn(vm *goja.Runtime) func(script string, params Params) interface{} {
	return func(script string, params Params) interface{} {

		executionVm := vm

		if params.Vm != nil {
			executionVm = params.Vm
		}

		var argNames []string
		var argValues []goja.Value

		for key, _ := range params.Args {
			argNames = append(argNames, key)
			argValues = append(argValues, executionVm.ToValue(params.Args[key]))
		}

		if params.Isolate {
			fnContent := `function execute__(` + strings.Join(argNames, ",") + `) {` + script + `}`

			_, err := executionVm.RunString(fnContent)

			if err != nil {
				panic(err)
			}

			fn := executionVm.Get("execute__").Export().(func(call goja.FunctionCall) goja.Value)

			value := fn(goja.FunctionCall{
				This:      params.This,
				Arguments: argValues,
			})

			val := value.Export()

			return val
		} else {
			for key, value := range params.Args {
				if err := executionVm.Set(key, value); err != nil {
					panic(err)
				}
			}

			value, err := executionVm.RunString(script)

			if err != nil {
				panic(err)
			}

			return value.Export()
		}
	}
}

func initVmFn(s abs.CodeExecutorService) func(options abs.VmOptions) *goja.Runtime {
	return func(options abs.VmOptions) *goja.Runtime {
		vm, err := s.NewVm(options)

		if err != nil {
			panic(err)
		}

		return vm
	}
}

func Register(vm *goja.Runtime, s abs.CodeExecutorService) error {
	if err := vm.Set("execute", executeFn(vm)); err != nil {
		return err
	}

	if err := vm.Set("initVm", initVmFn(s)); err != nil {
		return err
	}

	return nil
}
