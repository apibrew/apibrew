package api

import (
	"github.com/gorilla/mux"
	"net"
	"net/http"
)

//import _ "net/http/pprof"

type Router interface {
	ConfigureRouter(router *mux.Router)
}

type Server interface {
	InjectRecordApi(api RecordApi)
	Serve(lis net.Listener)
	InjectAuthenticationApi(api AuthenticationApi)
}

type server struct {
	recordApi         RecordApi
	authenticationApi AuthenticationApi
}

func (s *server) InjectAuthenticationApi(api AuthenticationApi) {
	s.authenticationApi = api
}

func (s *server) Serve(lis net.Listener) {
	r := mux.NewRouter()
	s.recordApi.ConfigureRouter(r)
	s.authenticationApi.ConfigureRouter(r)

	//go func() {
	//	log.Println(http.ListenAndServe("localhost:6060", nil))
	//}()

	if err := http.Serve(lis, r); err != nil {
		panic(err)
	}
}

func (s *server) InjectRecordApi(api RecordApi) {
	s.recordApi = api
}

func NewServer() Server {
	return &server{}
}
