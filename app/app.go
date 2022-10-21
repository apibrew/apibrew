package app

import (
	"data-handler/api"
	"data-handler/service"
	"data-handler/service/backend"
	"data-handler/service/backend/postgres"
	"data-handler/stub"
	"data-handler/stub/model"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

type App struct {
	initData                       *model.InitData
	grpcServer                     *grpc.Server
	authenticationService          service.AuthenticationService
	dataSourceService              service.DataSourceService
	resourceService                service.ResourceService
	recordService                  service.RecordService
	workSpaceService               stub.WorkSpaceServiceServer
	userService                    service.UserService
	postgresResourceServiceBackend backend.ResourceServiceBackend
	recordApi                      api.RecordApi
	authenticationApi              api.AuthenticationApi
	apiServer                      api.Server
	grpcLis                        net.Listener
	httpLis                        net.Listener
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
	app.postgresResourceServiceBackend = postgres.NewPostgresResourceServiceBackend()
	app.workSpaceService = service.NewWorkSpaceService()
	app.userService = service.NewUserService()
	app.recordApi = api.NewRecordApi()
	app.authenticationApi = api.NewAuthenticationApi()
	app.apiServer = api.NewServer()

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

	var opts []grpc.ServerOption
	app.grpcServer = grpc.NewServer(opts...)
	stub.RegisterResourceServiceServer(app.grpcServer, app.resourceService)
	stub.RegisterAuthenticationServiceServer(app.grpcServer, app.authenticationService)
	stub.RegisterDataSourceServiceServer(app.grpcServer, app.dataSourceService)
	stub.RegisterRecordServiceServer(app.grpcServer, app.recordService)
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
	app.dataSourceService.Init()
	app.resourceService.Init(app.initData)
	//workSpaceService.Init(initData)
	app.recordService.Init(app.initData)
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

	app.recordService.InjectPostgresResourceServiceBackend(app.postgresResourceServiceBackend)
	app.recordService.InjectDataSourceService(app.dataSourceService)
	app.recordService.InjectAuthenticationService(app.authenticationService)
	app.recordService.InjectResourceService(app.resourceService)

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
