package rest

import (
	"context"
	"data-handler/app"
	"data-handler/helper"
	"data-handler/logging"
	"data-handler/server/stub"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"net/http"
	"strings"
)

//import _ "net/http/pprof"

type Router interface {
	ConfigureRouter(router *mux.Router)
}

type Server interface {
	Serve(lis net.Listener)
}

type server struct {
	recordApi  RecordApi
	swaggerApi SwaggerApi
}

func (r *server) AuthenticationMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		authorizationHeader := req.Header.Get("Authorization")

		if authorizationHeader != "" {
			tokenParts := strings.Split(authorizationHeader, " ")

			if len(tokenParts) != 2 {
				handleClientErrorText(w, "authorization header should contain two part") //@todo fixme
				return
			}

			if strings.ToLower(tokenParts[0]) != "bearer" {
				handleClientErrorText(w, "authorization token type should be bearer") //@todo fixme
				return
			}

			token := tokenParts[1]

			if req.URL.RawQuery == "" {
				req.URL.RawQuery = "token=" + token
			} else {
				req.URL.RawQuery = req.URL.RawQuery + "token=" + token
			}
		}

		next.ServeHTTP(w, req)
	})
}

func (s *server) Serve(lis net.Listener) {
	r := mux.NewRouter()

	r.Use(s.AuthenticationMiddleWare)
	r.Use(s.TrackingMiddleWare)

	c := cors.New(cors.Options{
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
		AllowedOrigins:   []string{"*"},
		AllowedHeaders:   []string{"Authorization"},
		Debug:            true,
	})

	m := runtime.NewServeMux()

	r.PathPrefix("/records").Handler(m)
	r.PathPrefix("/users").Handler(m)
	r.PathPrefix("/authentication").Handler(m)

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	stub.RegisterAuthenticationServiceHandlerFromEndpoint(context.TODO(), m, "localhost:9009", opts)
	stub.RegisterUserServiceHandlerFromEndpoint(context.TODO(), m, "localhost:9009", opts)
	stub.RegisterRecordServiceHandlerFromEndpoint(context.TODO(), m, "localhost:9009", opts)

	s.swaggerApi.ConfigureRouter(r)
	s.recordApi.ConfigureRouter(r)

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

func NewServer(container app.Container) Server {
	return &server{
		recordApi:  NewRecordApi(container.GetRecordService(), container.GetResourceService()),
		swaggerApi: NewSwaggerApi(container.GetResourceService()),
	}
}
