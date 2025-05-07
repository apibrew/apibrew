package test

import (
	"github.com/apibrew/apibrew/modules/tenant"
	"github.com/apibrew/apibrew/pkg/api"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/service"
	"github.com/apibrew/apibrew/pkg/service/impl"
	"github.com/apibrew/apibrew/pkg/test/setup"
)

var container service.Container

var apiInterface api.Interface

func init() {
	container = setup.GetContainer()

	app := container.(*impl.App)

	app.GetAppConfig().Modules = map[string]*model.ModuleConfig{
		"tenant": {},
	}

	app.RegisterModule(pkg.NewModule)

	apiInterface = api.NewInterface(container)
}
