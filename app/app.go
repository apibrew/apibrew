package app

import (
	"data-handler/logging"
	"data-handler/model"
	"data-handler/params"
	grpc_server "data-handler/server/grpc"
	"data-handler/server/rest"
	"data-handler/service"
	"data-handler/service/handler"
	"data-handler/service/handlers"
	log "github.com/sirupsen/logrus"
	"net"
)

type App struct {
	initData               *model.InitData
	grpcServer             grpc_server.GrpcServer
	authenticationService  service.AuthenticationService
	dataSourceService      service.DataSourceService
	resourceService        service.ResourceService
	recordService          service.RecordService
	backendProviderService service.BackendProviderService
	namespaceService       service.NamespaceService
	userService            service.UserService
	apiServer              rest.Server
	grpcLis                net.Listener
	httpLis                net.Listener
	genericHandler         *handler.GenericHandler
	stdHandler             handlers.StdHandler
	watchService           service.WatchService
}

type GrpcContainer interface {
	GetRecordService() service.RecordService
	GetAuthenticationService() service.AuthenticationService
	GetResourceService() service.ResourceService
	GetDataSourceService() service.DataSourceService
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

	app.grpcServer = grpc_server.NewGrpcServer(params.ServerInjectionConstructorParams{
		ResourceService:       app.resourceService,
		RecordService:         app.recordService,
		AuthenticationService: app.authenticationService,
		DataSourceService:     app.dataSourceService,
		NamespaceService:      app.namespaceService,
		UserService:           app.userService,
		WatchService:          app.watchService,
	})

	app.apiServer = rest.NewServer(params.ServerInjectionConstructorParams{
		ResourceService:       app.resourceService,
		RecordService:         app.recordService,
		AuthenticationService: app.authenticationService,
		DataSourceService:     app.dataSourceService,
		NamespaceService:      app.namespaceService,
		UserService:           app.userService,
		WatchService:          app.watchService,
	}, app.initData)

	app.InjectServices()
	app.initServices()

	var err error
	app.grpcLis, err = net.Listen("tcp", app.initData.Config.GrpcAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	app.httpLis, err = net.Listen("tcp", app.initData.Config.HttpAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	app.grpcServer.Init(app.initData)
}

func (app *App) Serve() {
	go app.apiServer.Serve(app.httpLis)

	err := app.grpcServer.Serve(app.grpcLis)

	if err != nil {
		panic(err)
	}
}

func (app *App) Stop() {
	app.grpcServer.Stop()
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
