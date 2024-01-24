package execute

import (
	"github.com/dop251/goja"
	"strings"
)

type Params struct {
	This goja.Value             `json:"this"`
	Args map[string]interface{} `json:"args"`
}

func executeFn(vm *goja.Runtime) func(script string, params Params) interface{} {
	return func(script string, params Params) interface{} {

		var argNames []string
		var argValues []goja.Value

		for key, _ := range params.Args {
			argNames = append(argNames, key)
			argValues = append(argValues, vm.ToValue(params.Args[key]))
		}

		fnContent := `function execute__(` + strings.Join(argNames, ",") + `) {` + script + `}`

		_, err := vm.RunString(fnContent)

		if err != nil {
			panic(err)
		}

		fn := vm.Get("execute__").Export().(func(call goja.FunctionCall) goja.Value)

		value := fn(goja.FunctionCall{
			This:      params.This,
			Arguments: argValues,
		})

		val := value.Export()

		return val
	}
}

func Register(vm *goja.Runtime) error {
	return vm.Set("execute", executeFn(vm))
}
