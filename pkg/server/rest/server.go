package rest

import (
	"context"
	"github.com/apibrew/apibrew/pkg/abs"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/helper"
	"github.com/apibrew/apibrew/pkg/logging"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/server/grpc"
	"github.com/apibrew/apibrew/pkg/server/rest/docs"
	"github.com/apibrew/apibrew/pkg/service/security"
	"github.com/apibrew/apibrew/pkg/stub"
	"github.com/apibrew/apibrew/pkg/stub/rest"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"net"
	"net/http"
	"strings"
)

type Router interface {
	ConfigureRouter(router *mux.Router)
}

type Server interface {
	ServeH2C(lis net.Listener)
	ServeHttp(lis net.Listener)
	Init(data *model.InitData)
	ServeHttp2Tls(tls net.Listener)
}

type server struct {
	handler                     http.Handler
	certFile                    string
	keyFile                     string
	recordsApiFiltersMiddleWare *recordsApiFiltersMiddleWare
	container                   abs.Container
	docsApi                     docs.Api
}

func (s *server) Init(*model.InitData) {
	s.configureRoutes()
}

func (s *server) AuthenticationMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		authorizationHeader := req.Header.Get("Authorization")

		if authorizationHeader != "" {
			tokenParts := strings.Split(authorizationHeader, " ")

			if len(tokenParts) != 2 {
				http.Error(w, "authorization header should contain two part", 400)
				return
			}

			if strings.ToLower(tokenParts[0]) != "bearer" {
				http.Error(w, "authorization token type should be bearer", 400)
				return
			}

			token := tokenParts[1]

			req.Header.Set("Grpc-Metadata-token", token)

			userDetails, err := s.container.GetAuthenticationService().ParseAndVerifyToken(token)

			if err == nil {
				ctx := security.WithUserDetails(req.Context(), *userDetails)

				ctx = logging.WithLogField(ctx, "User", userDetails.Username)

				req = req.WithContext(ctx)
			} else {
				http.Error(w, errors.AuthenticationFailedError.Error(), 401)
				return
			}
		}

		next.ServeHTTP(w, req)
	})
}

func (s *server) ServeH2C(lis net.Listener) {
	h2s := &http2.Server{}

	srv := &http.Server{
		Handler: h2c.NewHandler(s.handler, h2s),
	}

	if err := srv.Serve(lis); err != nil {
		panic(err)
	}
}

func (s *server) ServeHttp(lis net.Listener) {
	srv := &http.Server{
		Handler: s.handler,
	}

	if err := srv.Serve(lis); err != nil {
		panic(err)
	}
}

func (s *server) ServeHttp2Tls(tls net.Listener) {
	srv := &http.Server{
		Handler: s.handler,
	}

	if s.certFile == "" || s.keyFile == "" {
		log.Warn("Cert and Key is not provided: TLS will be disabled")
		return
	}

	if err := srv.ServeTLS(tls, s.certFile, s.keyFile); err != nil {
		panic(err)
	}
}

func (s *server) configureRoutes() {
	r := mux.NewRouter()

	r.Use(s.AuthenticationMiddleWare)
	r.Use(s.TrackingMiddleWare)

	c := cors.New(cors.Options{
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowCredentials: true,
		AllowedOrigins:   []string{"*"},
		AllowedHeaders:   []string{"*"},
	})

	if log.GetLevel() >= log.DebugLevel {
		c.Log = log.New()
	}

	m := runtime.NewServeMux()

	r.PathPrefix("/records").Handler(m)
	r.PathPrefix("/authentication").Handler(m)
	r.PathPrefix("/system").Handler(m)

	if err := stub.RegisterAuthenticationHandlerServer(context.TODO(), m, grpc.NewAuthenticationServer(s.container.GetAuthenticationService())); err != nil {
		log.Fatal(err)
	}
	if err := rest.RegisterRecordHandlerServer(context.TODO(), m, newRecordService(s.container.GetRecordService())); err != nil {
		log.Fatal(err)
	}
	if err := stub.RegisterUserHandlerServer(context.TODO(), m, grpc.NewUserServer(s.container.GetUserService())); err != nil {
		log.Fatal(err)
	}
	if err := stub.RegisterRecordHandlerServer(context.TODO(), m, grpc.NewRecordServer(s.container.GetRecordService(), s.container.GetAuthenticationService())); err != nil {
		log.Fatal(err)
	}
	if err := stub.RegisterResourceHandlerServer(context.TODO(), m, grpc.NewResourceServer(s.container.GetResourceService())); err != nil {
		log.Fatal(err)
	}
	if err := stub.RegisterNamespaceHandlerServer(context.TODO(), m, grpc.NewNamespaceServer(s.container.GetNamespaceService())); err != nil {
		log.Fatal(err)
	}
	if err := stub.RegisterDataSourceHandlerServer(context.TODO(), m, grpc.NewDataSourceServer(s.container.GetDataSourceService())); err != nil {
		log.Fatal(err)
	}
	if err := stub.RegisterExtensionHandlerServer(context.TODO(), m, grpc.NewExtensionServer(s.container.GetExtensionService())); err != nil {
		log.Fatal(err)
	}
	if err := stub.RegisterWatchHandlerServer(context.TODO(), m, grpc.NewWatchServer(s.container.GetWatchService())); err != nil {
		log.Fatal(err)
	}

	r.PathPrefix("/docs").Handler(s.docsApi.Handler())

	s.handler = c.Handler(s.recordsApiFiltersMiddleWare.handler(r))
}

func (s *server) TrackingMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		trackId := helper.RandStringRunes(8)
		w.Header().Set("TrackId", trackId)

		req = req.WithContext(logging.WithLogField(req.Context(), "TrackId", trackId))

		if req.Header.Get("ClientTrackId") != "" {
			req = req.WithContext(logging.WithLogField(req.Context(), "ClientTrackId", req.Header.Get("ClientTrackId")))
		}

		next.ServeHTTP(w, req)
	})
}

func NewServer(container abs.Container) Server {
	return &server{
		container:                   container,
		docsApi:                     docs.NewApi(container.GetResourceService()),
		recordsApiFiltersMiddleWare: newRecordsApiFiltersMiddleWare(container.GetResourceService()),
	}
}
