package api

import (
	"github.com/apibrew/apibrew/pkg/api"
	"github.com/apibrew/apibrew/pkg/nano/abs"
	"github.com/apibrew/apibrew/pkg/service"
	"github.com/dop251/goja"
)

func Register(vm *goja.Runtime, cec abs.CodeExecutionContext, container service.Container) error {
	apiInterface := api.NewInterface(container)

	if err := vm.Set("create", create(cec, apiInterface)); err != nil {
		return err
	}

	if err := vm.Set("update", update(cec, apiInterface)); err != nil {
		return err
	}

	if err := vm.Set("apply", apply(cec, apiInterface)); err != nil {
		return err
	}

	if err := vm.Set("delete", delete(cec, apiInterface)); err != nil {
		return err
	}

	if err := vm.Set("load", load(cec, apiInterface)); err != nil {
		return err
	}

	if err := vm.Set("list", list(cec, apiInterface)); err != nil {
		return err
	}

	if err := vm.Set("resourceByName", resourceByName(cec, apiInterface)); err != nil {
		return err
	}

	return nil
}
