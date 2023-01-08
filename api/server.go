package api

import (
	"data-handler/api/swagger"
	"data-handler/helper"
	"data-handler/logging"
	"data-handler/model"
	"data-handler/params"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
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
}

func (s *server) Serve(lis net.Listener) {
	r := mux.NewRouter()

	swagger.ConfigureRouter(r)

	r.Use(s.authenticationApi.AuthenticationMiddleWare)
	r.Use(s.TrackingMiddleWare)

	s.authenticationApi.ConfigureRouter(r)
	s.recordApi.ConfigureRouter(r)

	c := cors.New(cors.Options{
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	})

	r.Use(c.Handler)

	if err := http.Serve(lis, r); err != nil {
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
	}
}
