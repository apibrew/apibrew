package api

import (
	"context"
	"data-handler/grpc/stub"
	"data-handler/helper"
	"data-handler/logging"
	"data-handler/model"
	"data-handler/params"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"net/http"
)

//import _ "net/http/pprof"

type Router interface {
	ConfigureRouter(router *mux.Router)
}

type Server interface {
	Serve(lis net.Listener)
}

type server struct {
	recordApi         RecordApi
	authenticationApi AuthenticationApi
	swaggerApi        SwaggerApi
}

func (s *server) Serve(lis net.Listener) {
	r := mux.NewRouter()

	r.Use(s.authenticationApi.AuthenticationMiddleWare)
	r.Use(s.TrackingMiddleWare)

	c := cors.New(cors.Options{
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
		AllowedOrigins:   []string{"*"},
		Debug:            true,
	})

	s.swaggerApi.ConfigureRouter(r)
	s.authenticationApi.ConfigureRouter(r)
	s.recordApi.ConfigureRouter(r)

	mux := runtime.NewServeMux()

	r.PathPrefix("/v1").Handler(mux)

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	stub.RegisterAuthenticationServiceHandlerFromEndpoint(context.TODO(), mux, "localhost:9009", opts)
	stub.RegisterUserServiceHandlerFromEndpoint(context.TODO(), mux, "localhost:9009", opts)

	if err := http.Serve(lis, c.Handler(r)); err != nil {
		panic(err)
	}
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

func NewServer(serverInjectionParams params.ServerInjectionConstructorParams, initData *model.InitData) Server {
	return &server{
		recordApi:         NewRecordApi(serverInjectionParams.RecordService, serverInjectionParams.ResourceService),
		authenticationApi: NewAuthenticationApi(serverInjectionParams.AuthenticationService, initData),
		swaggerApi:        NewSwaggerApi(serverInjectionParams.ResourceService),
	}
}
