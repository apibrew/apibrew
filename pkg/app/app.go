package app

import (
	log "github.com/sirupsen/logrus"
	"github.com/tislib/data-handler/pkg/logging"
	"github.com/tislib/data-handler/pkg/model"
	service2 "github.com/tislib/data-handler/pkg/service"
	"github.com/tislib/data-handler/pkg/service/handler"
	"github.com/tislib/data-handler/pkg/service/handlers"
)

type App struct {
	initData               *model.InitData
	authenticationService  service2.AuthenticationService
	dataSourceService      service2.DataSourceService
	resourceService        service2.ResourceService
	recordService          service2.RecordService
	backendProviderService service2.BackendProviderService
	namespaceService       service2.NamespaceService
	userService            service2.UserService
	genericHandler         *handler.GenericHandler
	stdHandler             handlers.StdHandler
	watchService           service2.WatchService
}

type Container interface {
	GetRecordService() service2.RecordService
	GetAuthenticationService() service2.AuthenticationService
	GetResourceService() service2.ResourceService
	GetDataSourceService() service2.DataSourceService
	GetWatchService() service2.WatchService
	GetNamespaceService() service2.NamespaceService
	GetUserService() service2.UserService
}

func (app *App) GetWatchService() service2.WatchService {
	return app.watchService
}

func (app *App) GetNamespaceService() service2.NamespaceService {
	return app.namespaceService
}

func (app *App) GetUserService() service2.UserService {
	return app.userService
}

func (app *App) GetRecordService() service2.RecordService {
	return app.recordService
}

func (app *App) GetAuthenticationService() service2.AuthenticationService {
	return app.authenticationService
}

func (app *App) GetResourceService() service2.ResourceService {
	return app.resourceService
}

func (app *App) GetDataSourceService() service2.DataSourceService {
	return app.dataSourceService
}

func (app *App) Init() {
	app.authenticationService = service2.NewAuthenticationService()
	app.dataSourceService = service2.NewDataSourceService()
	app.resourceService = service2.NewResourceService()
	app.backendProviderService = service2.NewBackendProviderService()
	app.recordService = service2.NewRecordService()
	app.genericHandler = handler.NewGenericHandler()
	app.namespaceService = service2.NewNamespaceService()
	app.userService = service2.NewUserService()
	app.stdHandler = handlers.NewStdHandler(app.genericHandler, app.dataSourceService, app.userService, app.recordService)
	app.watchService = service2.NewWatchService(app.genericHandler)

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
	app.userService.InjectBackendProviderService(app.backendProviderService)

	app.namespaceService.InjectRecordService(app.recordService)
	app.namespaceService.InjectResourceService(app.resourceService)
	app.namespaceService.InjectBackendProviderService(app.backendProviderService)

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
