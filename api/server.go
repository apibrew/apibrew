package api

import (
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
	s.recordApi.ConfigureRouter(r)
	s.authenticationApi.ConfigureRouter(r)

	//go func() {
	//	log.Println(http.ListenAndServe("localhost:6060", nil))
	//}()

	c := cors.New(cors.Options{
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	})

	handler := c.Handler(r)
	if err := http.Serve(lis, handler); err != nil {
		panic(err)
	}
}

func NewServer(serverInjectionParams params.ServerInjectionConstructorParams) Server {
	return &server{}
}
