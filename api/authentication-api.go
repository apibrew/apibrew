package api

import (
	"data-handler/service"
	"github.com/gorilla/mux"
	"net/http"
)

type AuthenticationApi interface {
	InjectAuthenticationService(service service.AuthenticationService)
	ConfigureRouter(r *mux.Router)
}

type authenticationApi struct {
	authenticationService service.AuthenticationService
}

func (r *authenticationApi) ConfigureRouter(router *mux.Router) {
	subRouter := router.PathPrefix("/authentication/").Subrouter()

	subRouter.HandleFunc("/token", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("TOKENINGGGG"))
	})
}

func (r *authenticationApi) InjectAuthenticationService(service service.AuthenticationService) {
	r.authenticationService = service
}

func NewAuthenticationApi() AuthenticationApi {
	return &authenticationApi{}
}
