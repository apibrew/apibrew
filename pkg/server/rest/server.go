package rest

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
	"github.com/tislib/data-handler/pkg/abs"
	"github.com/tislib/data-handler/pkg/helper"
	"github.com/tislib/data-handler/pkg/logging"
	"github.com/tislib/data-handler/pkg/model"
	_ "github.com/tislib/data-handler/pkg/server/rest/statik"
	"github.com/tislib/data-handler/pkg/stub"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
	swaggerApi                  SwaggerApi
	handler                     http.Handler
	certFile                    string
	keyFile                     string
	recordsApiFiltersMiddleWare *recordsApiFiltersMiddleWare
}

func (s *server) Init(*model.InitData) {
	s.configureRoutes()
	s.keyFile = "/Users/taleh/Projects/data-handler/dev/server.key"
	s.certFile = "/Users/taleh/Projects/data-handler/dev/server.crt"
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

	if err := srv.ServeTLS(tls, s.certFile, s.keyFile); err != nil {
		panic(err)
	}
}

func (s *server) configureRoutes() {
	r := mux.NewRouter()

	r.Use(s.AuthenticationMiddleWare)
	r.Use(s.TrackingMiddleWare)

	c := cors.New(cors.Options{
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
		AllowedOrigins:   []string{"*"},
		AllowedHeaders:   []string{"Authorization"},
	})

	m := runtime.NewServeMux()

	r.PathPrefix("/records").Handler(m)
	r.PathPrefix("/authentication").Handler(m)
	r.PathPrefix("/system").Handler(m)

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	if err := stub.RegisterAuthenticationServiceHandlerFromEndpoint(context.TODO(), m, "localhost:9009", opts); err != nil {
		log.Fatal(err)
	}
	if err := stub.RegisterUserServiceHandlerFromEndpoint(context.TODO(), m, "localhost:9009", opts); err != nil {
		log.Fatal(err)
	}
	if err := stub.RegisterRecordServiceHandlerFromEndpoint(context.TODO(), m, "localhost:9009", opts); err != nil {
		log.Fatal(err)
	}
	if err := stub.RegisterResourceServiceHandlerFromEndpoint(context.TODO(), m, "localhost:9009", opts); err != nil {
		log.Fatal(err)
	}
	if err := stub.RegisterNamespaceServiceHandlerFromEndpoint(context.TODO(), m, "localhost:9009", opts); err != nil {
		log.Fatal(err)
	}
	if err := stub.RegisterDataSourceServiceHandlerFromEndpoint(context.TODO(), m, "localhost:9009", opts); err != nil {
		log.Fatal(err)
	}
	if err := stub.RegisterWatchServiceHandlerFromEndpoint(context.TODO(), m, "localhost:9009", opts); err != nil {
		log.Fatal(err)
	}

	s.swaggerApi.ConfigureRouter(r)

	s.handler = s.recordsApiFiltersMiddleWare.handler(c.Handler(r))
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
		swaggerApi:                  NewSwaggerApi(container.GetResourceService()),
		recordsApiFiltersMiddleWare: newRecordsApiFiltersMiddleWare(container.GetResourceService()),
	}
}
