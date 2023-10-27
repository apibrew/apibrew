package impl

import (
	"context"
	"github.com/apibrew/apibrew/pkg/apbr/flags"
	"github.com/apibrew/apibrew/pkg/client"
	"github.com/apibrew/apibrew/pkg/formats/executor"
	"github.com/apibrew/apibrew/pkg/logging"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/service"
	backend_event_handler "github.com/apibrew/apibrew/pkg/service/backend-event-handler"
	"github.com/apibrew/apibrew/pkg/service/handlers"
	"github.com/apibrew/apibrew/pkg/util"
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
	eventChannelService      service.EventChannelService
	metricsService           service.MetricsService
	auditService             service.AuditService
	statsService             service.StatsService
}

func (app *App) GetAuthorizationService() service.AuthorizationService {
	return app.authorizationService
}

func (app *App) GetEventChannelService() service.EventChannelService {
	return app.eventChannelService
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

func (app *App) GetMetricsService() service.MetricsService {
	return app.metricsService
}

func (app *App) GetAuditService() service.AuditService {
	return app.auditService
}

func (app *App) GetStatsService() service.StatsService {
	return app.statsService
}

func (app *App) Init() <-chan interface{} {
	app.authorizationService = NewAuthorizationService()

	app.backendEventHandler = backend_event_handler.NewBackendEventHandler(app.authorizationService)

	app.backendProviderService = NewBackendProviderService(app.backendEventHandler)
	app.resourceMigrationService = NewResourceMigrationService()

	app.resourceService = NewResourceService(app.backendProviderService, app.resourceMigrationService, app.authorizationService)
	app.recordService = NewRecordService(app.resourceService, app.backendProviderService, app.authorizationService, app.backendEventHandler)

	app.dataSourceService = NewDataSourceService(app.resourceService, app.recordService, app.backendProviderService)
	app.eventChannelService = NewEventChannelService(app.authorizationService)
	app.externalService = NewExternalService(app.eventChannelService)
	app.extensionService = NewExtensionService(app.recordService, app.backendProviderService, app.backendEventHandler, app.externalService)

	app.stdHandler = handlers.NewStdHandler(app.backendEventHandler, app.backendProviderService, app.extensionService)
	app.watchService = NewWatchService(app.backendEventHandler, app.authorizationService, app.resourceService)
	app.authenticationService = NewAuthenticationService(app.recordService)
	app.metricsService = NewMetricService(app.recordService, app.resourceService)
	app.auditService = NewAuditService(app.backendEventHandler, app.recordService)
	app.statsService = NewStatsService(app.backendEventHandler)

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
	app.metricsService.Init(app.config)
	app.extensionService.Init(app.config)
	app.eventChannelService.Init(app.config)
	app.auditService.Init(app.config)
	app.statsService.Init(app.config)

	// run apply paths

	if app.config.ApplyPaths != nil {
		applier := executor.NewExecutor(executor.APPLY, client.NewLocalClient(app), true, false, false, flags.OverrideConfig{})
		for _, path := range app.config.ApplyPaths {
			err := applier.ApplyWithPattern(util.WithSystemContext(context.TODO()), path, "")

			if err != nil {
				log.Fatalf("failed to apply file: %s", err)
			}
		}
	}
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
