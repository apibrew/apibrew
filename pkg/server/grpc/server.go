package grpc

import (
	"context"
	"github.com/apibrew/apibrew/pkg/abs"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/helper"
	"github.com/apibrew/apibrew/pkg/logging"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/service/security"
	"github.com/apibrew/apibrew/pkg/stub"
	log "github.com/sirupsen/logrus"
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
	grpcServer               *grpc.Server
	resourceService          abs.ResourceService
	resourceMigrationService abs.ResourceMigrationService
	recordService            abs.RecordService
	authenticationService    abs.AuthenticationService
	dataSourceService        abs.DataSourceService
	namespaceService         abs.NamespaceService
	userService              abs.GenericRecordService[*model.User]
	initData                 *model.InitData
	watchService             abs.WatchService
	extensionService         abs.ExtensionService
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

	stub.RegisterResourceServer(g.grpcServer, NewResourceServer(g.resourceService))
	stub.RegisterAuthenticationServer(g.grpcServer, NewAuthenticationServer(g.authenticationService))
	stub.RegisterDataSourceServer(g.grpcServer, NewDataSourceServer(g.dataSourceService))
	stub.RegisterRecordServer(g.grpcServer, NewRecordServer(g.recordService, g.authenticationService))
	stub.RegisterUserServer(g.grpcServer, NewUserServer(g.userService))
	stub.RegisterNamespaceServer(g.grpcServer, NewNamespaceServer(g.namespaceService))
	stub.RegisterWatchServer(g.grpcServer, NewWatchServer(g.watchService))
	stub.RegisterExtensionServer(g.grpcServer, NewExtensionServer(g.extensionService))
	stub.RegisterGenericServer(g.grpcServer, NewGenericService(g.recordService))
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

	if authenticationService.AuthenticationDisabled() {
		ctx = security.WithSystemContext(ctx)
	} else if token != "" {
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
		resourceService:          container.GetResourceService(),
		resourceMigrationService: container.GetResourceMigrationService(),
		recordService:            container.GetRecordService(),
		watchService:             container.GetWatchService(),
		authenticationService:    container.GetAuthenticationService(),
		dataSourceService:        container.GetDataSourceService(),
		namespaceService:         container.GetNamespaceService(),
		userService:              container.GetUserService(),
		extensionService:         container.GetExtensionService(),
	}
}
