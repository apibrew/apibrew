package api

import (
	"github.com/apibrew/apibrew/modules/nano/pkg/abs"
	"github.com/apibrew/apibrew/pkg/api"
	"github.com/apibrew/apibrew/pkg/service"
	"github.com/dop251/goja"
)

func Register(vm *goja.Runtime, cec abs.CodeExecutionContext, container service.Container) error {
	apiInterface := api.NewInterface(container)

	if err := vm.Set("create", create(cec, vm, apiInterface)); err != nil {
		return err
	}

	if err := vm.Set("update", update(cec, vm, apiInterface)); err != nil {
		return err
	}

	if err := vm.Set("apply", apply(cec, vm, apiInterface)); err != nil {
		return err
	}

	if err := vm.Set("delete_", delete_(cec, vm, apiInterface)); err != nil {
		return err
	}

	if err := vm.Set("delete", delete_(cec, vm, apiInterface)); err != nil {
		return err
	}

	if err := vm.Set("load", load(cec, vm, apiInterface)); err != nil {
		return err
	}

	if err := vm.Set("list", list(cec, vm, apiInterface)); err != nil {
		return err
	}

	if err := vm.Set("resourceByName", resourceByName(cec, vm, apiInterface)); err != nil {
		return err
	}

	if err := vm.Set("begin", begin(cec, vm, apiInterface)); err != nil {
		return err
	}

	if err := vm.Set("commit", commit(cec, vm, apiInterface)); err != nil {
		return err
	}

	if err := vm.Set("rollback", rollback(cec, vm, apiInterface)); err != nil {
		return err
	}

	return nil
}
