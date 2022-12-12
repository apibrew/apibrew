package app

import (
	"data-handler/api"
	grpc_server "data-handler/grpc"
	"data-handler/model"
	"data-handler/service"
	"data-handler/service/backend"
	"data-handler/service/backend/postgres"
	"data-handler/service/handler"
	log "github.com/sirupsen/logrus"
	"net"
)

type App struct {
	initData                       *model.InitData
	grpcServer                     grpc_server.GrpcServer
	authenticationService          service.AuthenticationService
	dataSourceService              service.DataSourceService
	resourceService                service.ResourceService
	recordService                  service.RecordService
	workspaceService               service.WorkspaceService
	userService                    service.UserService
	postgresResourceServiceBackend backend.ResourceServiceBackend
	recordApi                      api.RecordApi
	authenticationApi              api.AuthenticationApi
	apiServer                      api.Server
	grpcLis                        net.Listener
	httpLis                        net.Listener
	genericHandler                 handler.GenericHandler
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
	app.recordService = service.NewRecordService()
	app.genericHandler = handler.NewGenericHandler()
	app.postgresResourceServiceBackend = postgres.NewPostgresResourceServiceBackend()
	app.workspaceService = service.NewWorkspaceService()
	app.userService = service.NewUserService()
	app.recordApi = api.NewRecordApi()
	app.authenticationApi = api.NewAuthenticationApi()
	app.apiServer = api.NewServer()

	app.grpcServer = grpc_server.NewGrpcServer(grpc_server.GrpcServerInjectionConstructorParams{
		ResourceService:       app.resourceService,
		RecordService:         app.recordService,
		AuthenticationService: app.authenticationService,
		DataSourceService:     app.dataSourceService,
		WorkspaceService:      app.workspaceService,
		UserService:           app.userService,
	})

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
	app.postgresResourceServiceBackend.Init()
	app.dataSourceService.Init(app.initData)
	app.resourceService.Init(app.initData)
	app.workspaceService.Init(app.initData)
	app.recordService.Init(app.initData)
	app.userService.Init(app.initData)
	app.authenticationService.Init(app.initData)
}

func (app *App) InjectServices() {
	app.dataSourceService.InjectResourceService(app.resourceService)
	app.dataSourceService.InjectRecordService(app.recordService)
	app.dataSourceService.InjectInitData(app.initData)
	app.dataSourceService.InjectPostgresResourceServiceBackend(app.postgresResourceServiceBackend)
	app.dataSourceService.InjectAuthenticationService(app.authenticationService)

	app.postgresResourceServiceBackend.InjectDataSourceService(app.dataSourceService)

	app.resourceService.InjectDataSourceService(app.dataSourceService)
	app.resourceService.InjectPostgresResourceServiceBackend(app.postgresResourceServiceBackend)
	app.resourceService.InjectAuthenticationService(app.authenticationService)

	app.userService.InjectAuthenticationService(app.authenticationService)
	app.userService.InjectRecordService(app.recordService)
	app.userService.InjectResourceService(app.resourceService)

	app.workspaceService.InjectAuthenticationService(app.authenticationService)
	app.workspaceService.InjectRecordService(app.recordService)
	app.workspaceService.InjectResourceService(app.resourceService)

	app.recordService.InjectPostgresResourceServiceBackend(app.postgresResourceServiceBackend)
	app.recordService.InjectDataSourceService(app.dataSourceService)
	app.recordService.InjectAuthenticationService(app.authenticationService)
	app.recordService.InjectResourceService(app.resourceService)
	app.recordService.InjectGenericHandler(app.genericHandler)

	app.authenticationService.InjectRecordService(app.recordService)

	app.authenticationApi.InjectAuthenticationService(app.authenticationService)

	app.recordApi.InjectRecordService(app.recordService)
	app.recordApi.InjectResourceService(app.resourceService)
	app.apiServer.InjectAuthenticationApi(app.authenticationApi)
	app.apiServer.InjectRecordApi(app.recordApi)
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
