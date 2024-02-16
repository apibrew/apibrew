package addons

import (
	"github.com/apibrew/apibrew/pkg/nano/abs"
	"github.com/apibrew/apibrew/pkg/nano/addons/auth"
	"github.com/apibrew/apibrew/pkg/nano/addons/aws"
	"github.com/apibrew/apibrew/pkg/nano/addons/console"
	"github.com/apibrew/apibrew/pkg/nano/addons/execute"
	"github.com/apibrew/apibrew/pkg/nano/addons/global"
	"github.com/apibrew/apibrew/pkg/nano/addons/http"
	"github.com/apibrew/apibrew/pkg/nano/addons/lambda"
	"github.com/apibrew/apibrew/pkg/nano/addons/mail"
	"github.com/apibrew/apibrew/pkg/nano/addons/resource"
	"github.com/apibrew/apibrew/pkg/service"
	"github.com/dop251/goja"
)

func Register(vm *goja.Runtime, cec abs.CodeExecutionContext, s abs.CodeExecutorService, codeName string, container service.Container) error {
	if err := console.Register(vm, codeName); err != nil {
		return err
	}

	if err := resource.Register(vm, cec, s); err != nil {
		return err
	}

	if err := lambda.Register(vm, cec, s); err != nil {
		return err
	}

	if err := mail.Register(vm); err != nil {
		return err
	}

	if err := http.Register(vm); err != nil {
		return err
	}

	if err := auth.Register(vm, container); err != nil {
		return err
	}

	if err := execute.Register(vm, s); err != nil {
		return err
	}

	if err := global.Register(vm, s); err != nil {
		return err
	}

	if err := aws.Register(vm); err != nil {
		return err
	}

	return nil
}
