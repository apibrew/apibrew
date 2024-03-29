package grpc

import (
	"context"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/helper"
	"github.com/apibrew/apibrew/pkg/logging"
	"github.com/apibrew/apibrew/pkg/service"
	"github.com/apibrew/apibrew/pkg/service/annotations"
	"github.com/apibrew/apibrew/pkg/stub"
	"github.com/apibrew/apibrew/pkg/util"
	jwt_model "github.com/apibrew/apibrew/pkg/util/jwt-model"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"net"
	"strings"
)

type Server interface {
	Serve(lis net.Listener)
	Init()
	Stop()
}

type grpcServer struct {
	grpcServer               *grpc.Server
	resourceService          service.ResourceService
	resourceMigrationService service.ResourceMigrationService
	recordService            service.RecordService
	authenticationService    service.AuthenticationService
	dataSourceService        service.DataSourceService
	watchService             service.WatchService
	eventChannelService      service.EventChannelService
}

func (g *grpcServer) Stop() {
	g.grpcServer.Stop()
}

func (g *grpcServer) Init() {
	var opts = []grpc.ServerOption{
		grpc.UnaryInterceptor(g.grpcIntercept),
		grpc.StreamInterceptor(g.grpcStreamIntercept),
		grpc.MaxRecvMsgSize(10 * 1024 * 1024),
	}

	g.grpcServer = grpc.NewServer(opts...)

	reflection.Register(g.grpcServer)

	stub.RegisterResourceServer(g.grpcServer, NewResourceServer(g.resourceService))
	stub.RegisterAuthenticationServer(g.grpcServer, NewAuthenticationServer(g.authenticationService))
	stub.RegisterDataSourceServer(g.grpcServer, NewDataSourceServer(g.dataSourceService))
	stub.RegisterRecordServer(g.grpcServer, NewRecordServer(g.recordService, g.authenticationService))
	stub.RegisterWatchServer(g.grpcServer, NewWatchServer(g.watchService, g.authenticationService))
	stub.RegisterEventChannelServer(g.grpcServer, NewEventChannelGrpcService(g.eventChannelService, g.authenticationService))
}

func (g *grpcServer) Serve(lis net.Listener) {
	if err := g.grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}

func (g *grpcServer) grpcIntercept(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	ctx, err := interceptRequest(g.authenticationService, ctx, req)

	if err != nil {
		return nil, err
	}

	return handler(ctx, req)
}

func (g *grpcServer) grpcStreamIntercept(req interface{}, ss grpc.ServerStream, _ *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	return handler(req, ss)
}

func interceptRequest(authenticationService service.AuthenticationService, ctx context.Context, req interface{}) (context.Context, error) {
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

		for k, v := range md {
			if len(v) > 0 {
				for ext, exists := range annotations.ClientAllowedAnnotations {
					if exists && strings.EqualFold(k, ext) {
						ctx = annotations.SetWithContext(ctx, ext, v[0])
					}
				}
			}
		}

	}

	if authenticationService.AuthenticationDisabled() {
		ctx = util.WithSystemContext(ctx)
	} else if token != "" {
		userDetails, err := authenticationService.ParseAndVerifyToken(token)

		if err != nil {
			return ctx, errors.AuthenticationFailedError
		}

		ctx = jwt_model.WithUserDetails(ctx, *userDetails)

		ctx = logging.WithLogField(ctx, "User", userDetails.Username)
	}

	// apbr-server track id
	trackId := helper.RandStringRunes(8)
	header := metadata.Pairs("TrackId", trackId)
	err := grpc.SetHeader(ctx, header)

	ctx = logging.WithLogField(ctx, "TrackId", trackId)

	if err != nil {
		return ctx, err
	}

	return ctx, nil
}

func NewGrpcServer(container service.Container) Server {
	return &grpcServer{
		resourceService:          container.GetResourceService(),
		resourceMigrationService: container.GetResourceMigrationService(),
		recordService:            container.GetRecordService(),
		watchService:             container.GetWatchService(),
		authenticationService:    container.GetAuthenticationService(),
		dataSourceService:        container.GetDataSourceService(),
		eventChannelService:      container.GetEventChannelService(),
	}
}
