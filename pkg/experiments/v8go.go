package main

import (
	"fmt"
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/console"
	"github.com/dop251/goja_nodejs/require"
	"io/ioutil"
	"time"
)

func main() {
	vm := goja.New()
	vm.SetFieldNameMapper(goja.UncapFieldNameMapper())

	registry := new(require.Registry) // this can be shared by multiple runtimes

	registry.Enable(vm)
	console.Enable(vm)

	registerSource(vm, "builtin.js", "/Users/taleh/Projects/apibrew/apibrew/pkg/experiments/builtin.js")
	registerSource(vm, "out.js", "/Users/taleh/Projects/apibrew/apibrew-manager/storage/out.js")

	time.Sleep(10 * time.Second)
}

func registerSource(vm *goja.Runtime, name string, srcFile string) (goja.Value, bool) {
	var content, err = ioutil.ReadFile(srcFile)
	//
	if err != nil {
		fmt.Println(err)
		return nil, true
	}

	script := string(content)
	res, err := vm.RunScript(name, script) // executes a script on the global context

	if err != nil {
		fmt.Println(err)
		return nil, true
	}
	return res, false
}
