package rest

import (
	"context"
	"fmt"
	"github.com/apibrew/apibrew/pkg/helper"
	"github.com/apibrew/apibrew/pkg/logging"
	"github.com/apibrew/apibrew/pkg/server/grpc"
	"github.com/apibrew/apibrew/pkg/server/rest/docs"
	"github.com/apibrew/apibrew/pkg/service"
	"github.com/apibrew/apibrew/pkg/service/annotations"
	"github.com/apibrew/apibrew/pkg/stub"
	"github.com/apibrew/apibrew/pkg/util"
	jwt_model "github.com/apibrew/apibrew/pkg/util/jwt-model"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
)

type Router interface {
	ConfigureRouter(router *mux.Router)
}

type Server interface {
	Init()
	ServeH2C(lis net.Listener)
	ServeHttp(lis net.Listener)
	ServeHttp2Tls(tls net.Listener)
}

type server struct {
	handler         http.Handler
	certFile        string
	keyFile         string
	container       service.Container
	docsApi         docs.Api
	metricsApi      MetricsApi
	healthApi       HealthApi
	recordApi       RecordApi
	resourceApi     ResourceApi
	eventChannelApi EventChannelApi
}

func (s *server) Init() {
	s.configureRoutes()
}

func (s *server) AuthenticationMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if s.container.GetAuthenticationService().AuthenticationDisabled() {
			req = req.WithContext(util.WithSystemContext(req.Context()))

			next.ServeHTTP(w, req)
			return
		}
		authorizationHeader := req.Header.Get("Authorization")
		tokenUrlParam := req.URL.Query().Get("token")

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

			var notAccepted bool
			req, notAccepted = s.setRequestToken(w, req, token)
			if notAccepted {
				return
			}
		} else if tokenUrlParam != "" {
			var notAccepted bool
			req, notAccepted = s.setRequestToken(w, req, tokenUrlParam)
			if notAccepted {
				return
			}
		}

		next.ServeHTTP(w, req)
	})
}

func (s *server) setRequestToken(w http.ResponseWriter, req *http.Request, token string) (*http.Request, bool) {
	req.Header.Set("Grpc-Metadata-token", token)

	userDetails, err := s.container.GetAuthenticationService().ParseAndVerifyToken(token)

	if err == nil {
		ctx := jwt_model.WithUserDetails(req.Context(), *userDetails)

		ctx = logging.WithLogField(ctx, "User", userDetails.Username)

		req = req.WithContext(ctx)
	} else {
		http.Error(w, err.Error(), 401)
		return nil, true
	}
	return req, false
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

	if log.GetLevel() >= log.TraceLevel {
		r.Use(s.TraceLogMiddleWare)
	}
	r.Use(s.AuthenticationMiddleWare)
	r.Use(s.TrackingMiddleWare)
	r.Use(s.AnnotationsMiddleWare)

	c := cors.New(cors.Options{
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowCredentials: true,
		AllowedOrigins:   []string{"*"},
		AllowedHeaders:   []string{"*"},
	})

	if log.GetLevel() >= log.TraceLevel {
		c.Log = log.New()
	}

	m := runtime.NewServeMux()

	s.recordApi.ConfigureRouter(r)
	s.resourceApi.ConfigureRouter(r)
	s.metricsApi.ConfigureRouter(r)
	s.healthApi.ConfigureRouter(r)
	s.eventChannelApi.ConfigureRouter(r)

	r.PathPrefix("/authentication").Handler(m)

	if err := stub.RegisterAuthenticationHandlerServer(context.TODO(), m, grpc.NewAuthenticationServer(s.container.GetAuthenticationService())); err != nil {
		log.Fatal(err)
	}

	r.PathPrefix("/docs").Handler(s.docsApi.Handler())

	healthEndpoint := new(HealthEndpoint)
	r.PathPrefix("/health").Handler(healthEndpoint.Handler())

	s.handler = c.Handler(r)
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

func (s *server) AnnotationsMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

		ctx := req.Context()

		for k, v := range req.Header {
			if len(v) > 0 {
				for ext, exists := range annotations.ClientAllowedAnnotations {
					if exists && strings.ToLower(k) == strings.ToLower(ext) {
						ctx = annotations.SetWithContext(ctx, ext, v[0])
					}
				}
			}
		}

		req = req.WithContext(ctx)

		next.ServeHTTP(w, req)
	})
}

func (s *server) TraceLogMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		x, err := httputil.DumpRequest(req, true)
		if err != nil {
			http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
			return
		}
		log.Tracef("Request: \n===============\n%s\n===============", string(x))

		next.ServeHTTP(w, req)
	})
}

func NewServer(container service.Container) Server {
	return &server{
		container:       container,
		docsApi:         docs.NewApi(container.GetResourceService()),
		recordApi:       NewRecordApi(container),
		resourceApi:     NewResourceApi(container),
		metricsApi:      NewMetricsApi(container.GetMetricsService()),
		healthApi:       NewHealthApi(),
		eventChannelApi: NewEventChannelApi(container),
	}
}
