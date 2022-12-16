package grpc_service

import (
	"context"
	"data-handler/grpc/stub"
	"data-handler/model"
	"data-handler/service"
	"data-handler/service/errors"
	"data-handler/service/security"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type GrpcServerInjectionConstructorParams struct {
	ResourceService       service.ResourceService
	RecordService         service.RecordService
	AuthenticationService service.AuthenticationService
	DataSourceService     service.DataSourceService
	WorkspaceService      service.WorkspaceService
	UserService           service.UserService
	WatchService          service.WatchService
}

type GrpcServer interface {
	Serve(lis net.Listener) error
	Init(*model.InitData)
	Stop()
}

type grpcServer struct {
	grpcServer            *grpc.Server
	resourceService       service.ResourceService
	recordService         service.RecordService
	authenticationService service.AuthenticationService
	dataSourceService     service.DataSourceService
	workspaceService      service.WorkspaceService
	userService           service.UserService
	initData              *model.InitData
	watchService          service.WatchService
}

func (g *grpcServer) Stop() {
	g.grpcServer.Stop()
}

func (g *grpcServer) Init(initData *model.InitData) {
	g.initData = initData
	var opts = []grpc.ServerOption{
		grpc.UnaryInterceptor(g.grpcIntercept),
	}

	g.grpcServer = grpc.NewServer(opts...)

	reflection.Register(g.grpcServer)

	stub.RegisterResourceServiceServer(g.grpcServer, NewResourceServiceServer(g.resourceService))
	stub.RegisterAuthenticationServiceServer(g.grpcServer, NewAuthenticationServiceServer(g.authenticationService))
	stub.RegisterDataSourceServiceServer(g.grpcServer, NewDataSourceServiceServer(g.dataSourceService))
	stub.RegisterRecordServiceServer(g.grpcServer, NewRecordServiceServer(g.recordService))
	stub.RegisterWatchServiceServer(g.grpcServer, NewWatchServiceServer(g.watchService))
}

func (g *grpcServer) Serve(lis net.Listener) error {
	return g.grpcServer.Serve(lis)
}

func (g *grpcServer) grpcIntercept(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	if rtw, ok := req.(RequestWithToken); ok && !g.initData.Config.DisableAuthentication {
		token := rtw.GetToken()

		userDetails, err := g.authenticationService.ParseAndVerifyToken(token)

		if err != nil {
			return nil, errors.AuthenticationFailedError
		}

		userCtx := security.WithUserDetails(ctx, *userDetails)

		return handler(userCtx, req)
	}
	return handler(ctx, req)
}

func NewGrpcServer(params GrpcServerInjectionConstructorParams) GrpcServer {
	return &grpcServer{
		resourceService:       params.ResourceService,
		recordService:         params.RecordService,
		watchService:          params.WatchService,
		authenticationService: params.AuthenticationService,
		dataSourceService:     params.DataSourceService,
		workspaceService:      params.WorkspaceService,
		userService:           params.UserService,
	}
}
