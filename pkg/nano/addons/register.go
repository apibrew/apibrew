package addons

import (
	"github.com/apibrew/apibrew/pkg/nano/abs"
	"github.com/apibrew/apibrew/pkg/nano/addons/console"
	"github.com/apibrew/apibrew/pkg/nano/addons/http"
	"github.com/apibrew/apibrew/pkg/nano/addons/lambda"
	"github.com/apibrew/apibrew/pkg/nano/addons/mail"
	"github.com/apibrew/apibrew/pkg/nano/addons/resource"
	"github.com/dop251/goja"
)

func Register(vm *goja.Runtime, cec abs.CodeExecutionContext, s abs.CodeExecutorService, codeName string) error {
	err := console.Register(vm, codeName)

	if err != nil {
		return err
	}

	err = resource.Register(vm, cec, s)

	if err != nil {
		return err
	}

	err = lambda.Register(vm, cec, s)

	if err != nil {
		return err
	}

	err = mail.Register(vm)

	if err != nil {
		return err
	}

	err = vm.Set("http", http.Register(vm))

	if err != nil {
		return err
	}
	return nil
}
