package service

import (
	"github.com/apibrew/apibrew/pkg/abs"
	"github.com/apibrew/apibrew/pkg/logging"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resources"
	backend_event_handler "github.com/apibrew/apibrew/pkg/service/backend-event-handler"
	"github.com/apibrew/apibrew/pkg/service/handlers"
	log "github.com/sirupsen/logrus"
)

type App struct {
	config                   *model.AppConfig
	authenticationService    abs.AuthenticationService
	authorizationService     abs.AuthorizationService
	dataSourceService        abs.DataSourceService
	resourceService          abs.ResourceService
	recordService            abs.RecordService
	backendProviderService   abs.BackendProviderService
	namespaceService         abs.NamespaceService
	userService              abs.GenericRecordService[*model.User]
	roleService              abs.GenericRecordService[*model.Role]
	stdHandler               handlers.StdHandlers
	watchService             abs.WatchService
	extensionService         abs.ExtensionService
	resourceMigrationService abs.ResourceMigrationService
	externalService          abs.ExternalService
	backendEventHandler      backend_event_handler.BackendEventHandler
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

func (app *App) GetUserService() abs.GenericRecordService[*model.User] {
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

func (app *App) Init() <-chan interface{} {
	app.backendEventHandler = backend_event_handler.NewBackendEventHandler()

	app.authorizationService = NewAuthorizationService()
	app.backendProviderService = NewBackendProviderService(app.backendEventHandler)
	app.resourceMigrationService = NewResourceMigrationService()

	app.resourceService = NewResourceService(app.backendProviderService, app.resourceMigrationService, app.authorizationService)
	app.recordService = NewRecordService(app.resourceService, app.backendProviderService, app.authorizationService)

	app.dataSourceService = NewDataSourceService(app.resourceService, app.recordService, app.backendProviderService)

	app.namespaceService = NewNamespaceService(app.resourceService, app.recordService, app.backendProviderService)

	app.userService = NewGenericRecordService[*model.User](app.recordService, resources.UserResource, func() *model.User {
		return &model.User{}
	})

	app.roleService = NewGenericRecordService[*model.Role](app.recordService, resources.RoleResource, func() *model.Role {
		return &model.Role{}
	})

	app.stdHandler = handlers.NewStdHandler(app.backendEventHandler)
	app.watchService = NewWatchService(app.backendEventHandler, app.authorizationService)
	app.externalService = NewExternalService()
	app.extensionService = NewExtensionService(app.recordService, app.backendProviderService, app.backendEventHandler, app.externalService)
	app.authenticationService = NewAuthenticationService(app.recordService, app.authorizationService)

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
	app.namespaceService.Init(app.config)
	app.userService.Init([]*model.User{})
	app.roleService.Init([]*model.Role{{
		Name: "root",
		SecurityConstraints: []*model.SecurityConstraint{
			{
				Namespace: "*",
				Resource:  "*",
				Property:  "*",
				Operation: model.OperationType_FULL,
				Permit:    model.PermitType_PERMIT_TYPE_ALLOW,
			},
		},
	}})
	app.authenticationService.Init(app.config)

	app.extensionService.Init(app.config)
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
