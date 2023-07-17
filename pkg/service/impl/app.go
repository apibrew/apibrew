package impl

import (
	"github.com/apibrew/apibrew/pkg/logging"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/service"
	backend_event_handler "github.com/apibrew/apibrew/pkg/service/backend-event-handler"
	"github.com/apibrew/apibrew/pkg/service/handlers"
	log "github.com/sirupsen/logrus"
)

type App struct {
	config                   *model.AppConfig
	authenticationService    service.AuthenticationService
	authorizationService     service.AuthorizationService
	dataSourceService        service.DataSourceService
	resourceService          service.ResourceService
	recordService            service.RecordService
	backendProviderService   service.BackendProviderService
	stdHandler               handlers.StdHandlers
	watchService             service.WatchService
	resourceMigrationService service.ResourceMigrationService
	externalService          service.ExternalService
	backendEventHandler      backend_event_handler.BackendEventHandler
	extensionService         service.ExtensionService
}

func (app *App) GetWatchService() service.WatchService {
	return app.watchService
}

func (app *App) GetResourceMigrationService() service.ResourceMigrationService {
	return app.resourceMigrationService
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

func (app *App) GetExtensionService() service.ExtensionService {
	return app.extensionService
}

func (app *App) Init() <-chan interface{} {
	app.backendEventHandler = backend_event_handler.NewBackendEventHandler()

	app.authorizationService = NewAuthorizationService()
	app.backendProviderService = NewBackendProviderService(app.backendEventHandler)
	app.resourceMigrationService = NewResourceMigrationService()

	app.resourceService = NewResourceService(app.backendProviderService, app.resourceMigrationService, app.authorizationService)
	app.recordService = NewRecordService(app.resourceService, app.backendProviderService, app.authorizationService)

	app.dataSourceService = NewDataSourceService(app.resourceService, app.recordService, app.backendProviderService)

	app.stdHandler = handlers.NewStdHandler(app.backendEventHandler, app.backendProviderService)
	app.watchService = NewWatchService(app.backendEventHandler, app.authorizationService)
	app.externalService = NewExternalService()
	app.authenticationService = NewAuthenticationService(app.recordService, app.authorizationService)
	app.extensionService = NewExtensionService(app.recordService, app.backendProviderService, app.backendEventHandler, app.externalService)

	initSignal := make(chan interface{})
	go func() {
		app.initServices()
		initSignal <- nil
	}()

	return initSignal
}

func (app *App) initServices() {
	app.backendProviderService.Init(app.config)
	app.stdHandler.Init(app.config)
	app.resourceService.Init(app.config)
	app.recordService.Init(app.config)
	app.dataSourceService.Init(app.config)
	app.authenticationService.Init(app.config)
}

func (app *App) SetConfig(config *model.AppConfig) {
	app.config = config

	app.CheckInitData(config)
}

func (app *App) CheckInitData(config *model.AppConfig) {
	if config.SystemDataSource == nil {
		log.Fatal("System dataSource is not set")
	}
}

//goland:noinspection GoUnusedParameter
func (app *App) SetGrayLogAddr(addr string) {
	logging.SetupGrayLog("tiswork.tisserv.net:12201", "test")
}
