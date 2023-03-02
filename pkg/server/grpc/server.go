package grpc

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/tislib/data-handler/pkg/abs"
	"github.com/tislib/data-handler/pkg/errors"
	"github.com/tislib/data-handler/pkg/helper"
	"github.com/tislib/data-handler/pkg/logging"
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/service/security"
	"github.com/tislib/data-handler/pkg/stub"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"net"
)

type Server interface {
	Serve(lis net.Listener)
	Init(*model.InitData)
	Stop()
}

type grpcServer struct {
	grpcServer            *grpc.Server
	resourceService       abs.ResourceService
	recordService         abs.RecordService
	authenticationService abs.AuthenticationService
	dataSourceService     abs.DataSourceService
	namespaceService      abs.NamespaceService
	userService           abs.UserService
	initData              *model.InitData
	watchService          abs.WatchService
	extensionService      abs.ExtensionService
}

func (g *grpcServer) Stop() {
	g.grpcServer.Stop()
}

func (g *grpcServer) Init(initData *model.InitData) {
	g.initData = initData
	var opts = []grpc.ServerOption{
		grpc.UnaryInterceptor(g.grpcIntercept),
		grpc.StreamInterceptor(g.grpcStreamIntercept),
	}

	g.grpcServer = grpc.NewServer(opts...)

	reflection.Register(g.grpcServer)

	stub.RegisterResourceServiceServer(g.grpcServer, NewResourceServiceServer(g.resourceService))
	stub.RegisterAuthenticationServiceServer(g.grpcServer, NewAuthenticationServiceServer(g.authenticationService))
	stub.RegisterDataSourceServiceServer(g.grpcServer, NewDataSourceServiceServer(g.dataSourceService))
	stub.RegisterRecordServiceServer(g.grpcServer, NewRecordServiceServer(g.recordService, g.authenticationService))
	stub.RegisterUserServiceServer(g.grpcServer, NewUserServiceServer(g.userService))
	stub.RegisterNamespaceServiceServer(g.grpcServer, NewNamespaceServiceServer(g.namespaceService))
	stub.RegisterWatchServiceServer(g.grpcServer, NewWatchServiceServer(g.watchService))
	stub.RegisterExtensionServiceServer(g.grpcServer, NewExtensionServiceServer(g.extensionService))
	stub.RegisterGenericServiceServer(g.grpcServer, NewGenericService(g.recordService))
}

func (g *grpcServer) Serve(lis net.Listener) {
	if err := g.grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}

func (g *grpcServer) grpcIntercept(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	ctx, err := interceptRequest(g.authenticationService, ctx, req)

	if err != nil {
		return nil, err
	}

	return handler(ctx, req)
}

func (g *grpcServer) grpcStreamIntercept(req interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	return handler(req, ss)
}

func interceptRequest(authenticationService abs.AuthenticationService, ctx context.Context, req interface{}) (context.Context, error) {
	// pass authentication context
	var token string

	if rtw, ok := req.(RequestWithToken); ok {
		token = rtw.GetToken()
	}

	// client track id
	md, mdExists := metadata.FromIncomingContext(ctx)

	if mdExists {
		if len(md.Get("ClientTrackId")) > 0 {
			ctx = logging.WithLogField(ctx, "ClientTrackId", md.Get("ClientTrackId")[0])
		}

		if token == "" {
			if len(md.Get("Token")) > 0 {
				token = md.Get("Token")[0]
			}
		}
	}

	if token != "" {
		userDetails, err := authenticationService.ParseAndVerifyToken(token)

		if err != nil {
			return ctx, errors.AuthenticationFailedError
		}

		ctx = security.WithUserDetails(ctx, *userDetails)

		ctx = logging.WithLogField(ctx, "User", userDetails.Username)
	}

	// server track id
	trackId := helper.RandStringRunes(8)
	header := metadata.Pairs("TrackId", trackId)
	err := grpc.SetHeader(ctx, header)

	ctx = logging.WithLogField(ctx, "TrackId", trackId)

	if err != nil {
		return ctx, err
	}

	return ctx, nil
}

func NewGrpcServer(container abs.Container) Server {
	return &grpcServer{
		resourceService:       container.GetResourceService(),
		recordService:         container.GetRecordService(),
		watchService:          container.GetWatchService(),
		authenticationService: container.GetAuthenticationService(),
		dataSourceService:     container.GetDataSourceService(),
		namespaceService:      container.GetNamespaceService(),
		userService:           container.GetUserService(),
		extensionService:      container.GetExtensionService(),
	}
}
