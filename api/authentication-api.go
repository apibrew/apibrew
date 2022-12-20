package api

import (
	"data-handler/grpc/stub"
	"data-handler/model"
	"data-handler/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

type LoginData struct {
	Username string
	Password string
	Term     model.TokenTerm
}

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
		var loginData = LoginData{}
		data, err := io.ReadAll(request.Body)

		if err != nil {
			handleClientError(writer, err)
			return
		}

		err = json.Unmarshal(data, &loginData)

		if err != nil {
			handleClientError(writer, err)
			return
		}

		token, serviceErr := r.authenticationService.Authenticate(request.Context(), loginData.Username, loginData.Password, loginData.Term)

		ServiceResponder[*stub.AuthenticationRequest, *stub.AuthenticationResponse]().
			Writer(writer).
			Request(request).
			Respond(&stub.AuthenticationResponse{
				Token: token,
				Error: toProtoError(serviceErr),
			}, serviceErr)

	}).Methods("POST")
}

func (r *authenticationApi) InjectAuthenticationService(service service.AuthenticationService) {
	r.authenticationService = service
}

func NewAuthenticationApi() AuthenticationApi {
	return &authenticationApi{}
}
