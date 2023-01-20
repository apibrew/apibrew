package app

import (
	"data-handler/logging"
	"data-handler/model"
	"data-handler/service"
	"data-handler/service/handler"
	"data-handler/service/handlers"
	log "github.com/sirupsen/logrus"
)

type App struct {
	initData               *model.InitData
	authenticationService  service.AuthenticationService
	dataSourceService      service.DataSourceService
	resourceService        service.ResourceService
	recordService          service.RecordService
	backendProviderService service.BackendProviderService
	namespaceService       service.NamespaceService
	userService            service.UserService
	genericHandler         *handler.GenericHandler
	stdHandler             handlers.StdHandler
	watchService           service.WatchService
}

type Container interface {
	GetRecordService() service.RecordService
	GetAuthenticationService() service.AuthenticationService
	GetResourceService() service.ResourceService
	GetDataSourceService() service.DataSourceService
	GetWatchService() service.WatchService
	GetNamespaceService() service.NamespaceService
	GetUserService() service.UserService
}

func (app *App) GetWatchService() service.WatchService {
	return app.watchService
}

func (app *App) GetNamespaceService() service.NamespaceService {
	return app.namespaceService
}

func (app *App) GetUserService() service.UserService {
	return app.userService
}

func (app *App) GetRecordService() service.RecordService {
	return app.recordService
}

func (app *App) GetAuthenticationService() service.AuthenticationService {
	return app.authenticationService
}

func (app *App) GetResourceService() service.ResourceService {
	return app.resourceService
}

func (app *App) GetDataSourceService() service.DataSourceService {
	return app.dataSourceService
}

func (app *App) Init() {
	app.authenticationService = service.NewAuthenticationService()
	app.dataSourceService = service.NewDataSourceService()
	app.resourceService = service.NewResourceService()
	app.backendProviderService = service.NewBackendProviderService()
	app.recordService = service.NewRecordService()
	app.genericHandler = handler.NewGenericHandler()
	app.namespaceService = service.NewNamespaceService()
	app.userService = service.NewUserService()
	app.stdHandler = handlers.NewStdHandler(app.genericHandler, app.dataSourceService)
	app.watchService = service.NewWatchService(app.genericHandler)

	app.InjectServices()
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
}

func (app *App) InjectServices() {
	app.dataSourceService.InjectResourceService(app.resourceService)
	app.dataSourceService.InjectRecordService(app.recordService)
	app.dataSourceService.InjectBackendProviderService(app.backendProviderService)

	app.resourceService.InjectBackendProviderService(app.backendProviderService)

	app.userService.InjectRecordService(app.recordService)
	app.userService.InjectResourceService(app.resourceService)

	app.namespaceService.InjectRecordService(app.recordService)
	app.namespaceService.InjectResourceService(app.resourceService)

	app.recordService.InjectBackendProviderService(app.backendProviderService)
	app.recordService.InjectResourceService(app.resourceService)
	app.recordService.InjectGenericHandler(app.genericHandler)

	app.authenticationService.InjectRecordService(app.recordService)

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
