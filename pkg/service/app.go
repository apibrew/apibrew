package service

import (
	log "github.com/sirupsen/logrus"
	"github.com/tislib/apibrew/pkg/abs"
	"github.com/tislib/apibrew/pkg/logging"
	"github.com/tislib/apibrew/pkg/model"
	"github.com/tislib/apibrew/pkg/service/handler"
	"github.com/tislib/apibrew/pkg/service/handlers"
)

type App struct {
	initData                 *model.InitData
	authenticationService    abs.AuthenticationService
	dataSourceService        abs.DataSourceService
	resourceService          abs.ResourceService
	recordService            abs.RecordService
	backendProviderService   abs.BackendProviderService
	namespaceService         abs.NamespaceService
	userService              abs.UserService
	genericHandler           *handler.GenericHandler
	stdHandler               handlers.StdHandler
	watchService             abs.WatchService
	extensionService         abs.ExtensionService
	pluginService            abs.PluginService
	resourceMigrationService abs.ResourceMigrationService
	externalService          abs.ExternalService
}

func (app *App) GetWatchService() abs.WatchService {
	return app.watchService
}

func (app *App) GetResourceMigrationService() abs.ResourceMigrationService {
	return app.resourceMigrationService
}

func (app *App) GetNamespaceService() abs.NamespaceService {
	return app.namespaceService
}

func (app *App) GetExtensionService() abs.ExtensionService {
	return app.extensionService
}

func (app *App) GetUserService() abs.UserService {
	return app.userService
}

func (app *App) GetRecordService() abs.RecordService {
	return app.recordService
}

func (app *App) GetAuthenticationService() abs.AuthenticationService {
	return app.authenticationService
}

func (app *App) GetResourceService() abs.ResourceService {
	return app.resourceService
}

func (app *App) GetDataSourceService() abs.DataSourceService {
	return app.dataSourceService
}

func (app *App) GetPluginService() abs.PluginService {
	return app.pluginService
}

func (app *App) Init() {
	app.backendProviderService = NewBackendProviderService()
	app.genericHandler = handler.NewGenericHandler()
	app.resourceMigrationService = NewResourceMigrationService()

	app.resourceService = NewResourceService(app.backendProviderService, app.resourceMigrationService)
	app.recordService = NewRecordService(app.resourceService, app.backendProviderService, app.genericHandler)

	app.dataSourceService = NewDataSourceService(app.resourceService, app.recordService, app.backendProviderService)

	app.namespaceService = NewNamespaceService(app.resourceService, app.recordService, app.backendProviderService)
	app.userService = NewUserService(app.resourceService, app.recordService, app.backendProviderService)
	app.stdHandler = handlers.NewStdHandler(app.genericHandler)
	app.watchService = NewWatchService(app.genericHandler)
	app.externalService = NewExternalService()
	app.extensionService = NewExtensionService(app.recordService, app.backendProviderService, app.genericHandler, app.externalService)
	app.pluginService = NewPluginService()
	app.authenticationService = NewAuthenticationService(app.recordService)

	app.initServices()
}

func (app *App) initServices() {
	app.backendProviderService.Init(app.initData)
	app.resourceService.Init(app.initData)
	app.dataSourceService.Init(app.initData)
	app.namespaceService.Init(app.initData)
	app.recordService.Init(app.initData)
	app.userService.Init(app.initData)
	app.authenticationService.Init(app.initData)
	app.stdHandler.Init(app.initData)

	app.extensionService.Init(app.initData)
	app.pluginService.Init(app.initData)
}

func (app *App) SetInitData(data *model.InitData) {
	app.initData = data

	app.CheckInitData(data)
}

func (app *App) CheckInitData(data *model.InitData) {
	if data.SystemDataSource == nil {
		log.Fatal("System dataSource is not set")
	}
}

//goland:noinspection GoUnusedParameter
func (app *App) SetGrayLogAddr(addr string) {
	logging.SetupGrayLog("tiswork.tisserv.net:12201", "test")
}
