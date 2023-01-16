package grpc_service

import (
	"context"
	"data-handler/grpc/stub"
	"data-handler/helper"
	"data-handler/logging"
	"data-handler/model"
	"data-handler/params"
	"data-handler/service"
	"data-handler/service/errors"
	"data-handler/service/security"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"net"
)

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
	namespaceService      service.NamespaceService
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
	// pass authentication context
	if rtw, ok := req.(RequestWithToken); ok && !g.initData.Config.DisableAuthentication {
		token := rtw.GetToken()

		userDetails, err := g.authenticationService.ParseAndVerifyToken(token)

		if err != nil {
			return nil, errors.AuthenticationFailedError
		}

		ctx = security.WithUserDetails(ctx, *userDetails)
	}

	// server track id
	trackId := helper.RandStringRunes(8)
	header := metadata.Pairs("TrackId", trackId)
	err := grpc.SetHeader(ctx, header)

	ctx = logging.WithLogField(ctx, "TrackId", trackId)

	if err != nil {
		return nil, err
	}

	// client track id
	md, mdExists := metadata.FromIncomingContext(ctx)

	if mdExists {
		if len(md.Get("ClientTrackId")) > 0 {
			ctx = logging.WithLogField(ctx, "ClientTrackId", md.Get("ClientTrackId")[0])
		}
	}

	return handler(ctx, req)
}

func NewGrpcServer(params params.ServerInjectionConstructorParams) GrpcServer {
	return &grpcServer{
		resourceService:       params.ResourceService,
		recordService:         params.RecordService,
		watchService:          params.WatchService,
		authenticationService: params.AuthenticationService,
		dataSourceService:     params.DataSourceService,
		namespaceService:      params.NamespaceService,
		userService:           params.UserService,
	}
}
