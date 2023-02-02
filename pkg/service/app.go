package service

import (
	log "github.com/sirupsen/logrus"
	"github.com/tislib/data-handler/pkg/logging"
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/service/handler"
	"github.com/tislib/data-handler/pkg/service/handlers"
)

type App struct {
	initData               *model.InitData
	authenticationService  AuthenticationService
	dataSourceService      DataSourceService
	resourceService        ResourceService
	recordService          RecordService
	backendProviderService BackendProviderService
	namespaceService       NamespaceService
	userService            UserService
	genericHandler         *handler.GenericHandler
	stdHandler             handlers.StdHandler
	watchService           WatchService
	extensionService       ExtensionService
	pluginService          PluginService
}

type Container interface {
	GetRecordService() RecordService
	GetAuthenticationService() AuthenticationService
	GetResourceService() ResourceService
	GetDataSourceService() DataSourceService
	GetWatchService() WatchService
	GetNamespaceService() NamespaceService
	GetUserService() UserService
	GetExtensionService() ExtensionService
	GetPluginService() PluginService
}

func (app *App) GetWatchService() WatchService {
	return app.watchService
}

func (app *App) GetNamespaceService() NamespaceService {
	return app.namespaceService
}

func (app *App) GetExtensionService() ExtensionService {
	return app.extensionService
}

func (app *App) GetUserService() UserService {
	return app.userService
}

func (app *App) GetRecordService() RecordService {
	return app.recordService
}

func (app *App) GetAuthenticationService() AuthenticationService {
	return app.authenticationService
}

func (app *App) GetResourceService() ResourceService {
	return app.resourceService
}

func (app *App) GetDataSourceService() DataSourceService {
	return app.dataSourceService
}

func (app *App) GetPluginService() PluginService {
	return app.pluginService
}

func (app *App) Init() {
	app.backendProviderService = NewBackendProviderService()
	app.genericHandler = handler.NewGenericHandler()

	app.resourceService = NewResourceService(app.backendProviderService)
	app.recordService = NewRecordService(app.resourceService, app.backendProviderService, app.genericHandler)

	app.dataSourceService = NewDataSourceService(app.resourceService, app.recordService, app.backendProviderService)

	app.namespaceService = NewNamespaceService(app.resourceService, app.recordService, app.backendProviderService)
	app.userService = NewUserService(app.resourceService, app.recordService, app.backendProviderService)
	app.stdHandler = handlers.NewStdHandler(app.genericHandler)
	app.watchService = NewWatchService(app.genericHandler)
	app.extensionService = NewExtensionService(app.recordService, app.backendProviderService, app.genericHandler)
	app.pluginService = NewPluginService()
	app.authenticationService = NewAuthenticationService(app.recordService)

	app.initServices()
}

func (app *App) initServices() {
	app.backendProviderService.Init(app.initData)
	app.dataSourceService.Init(app.initData)
	app.resourceService.Init(app.initData)
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

func (app *App) SetGrayLogAddr(addr string) {
	logging.SetupGrayLog("tiswork.tisserv.net:12201", "test")
}
