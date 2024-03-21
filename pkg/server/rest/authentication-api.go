package rest

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/service"
	"github.com/gorilla/mux"
	"net/http"
)

type AuthenticationApi interface {
	ConfigureRouter(r *mux.Router)
}

type authenticationApi struct {
	service service.AuthenticationService
}

func (m authenticationApi) ConfigureRouter(router *mux.Router) {
	subRoute := router.PathPrefix("/authentication/token").Subrouter()
	// collection level operations
	subRoute.HandleFunc("", m.handleAuthentication).Methods("POST")
	subRoute.HandleFunc("", m.handleRefreshToken).Methods("PUT")
}

func (m authenticationApi) handleAuthentication(writer http.ResponseWriter, request *http.Request) {
	authReq := new(AuthenticationRequest)

	err := parseRequestMessage(request, authReq)

	if err != nil {
		handleError(writer, err)
		return
	}

	token, serr := m.service.Authenticate(request.Context(), authReq.Username, authReq.Password, model.TokenTerm(model.TokenTerm_value[authReq.Term]), authReq.MinimizeToken)

	var resp = &AuthenticationResponse{}

	if token != nil {
		resp.Token = Token{
			Term:       token.Term.String(),
			Content:    token.Content,
			Expiration: token.Expiration.AsTime(),
		}
	}

	ServiceResponder().
		Writer(writer).
		Respond(resp, serr)
}

func (m authenticationApi) handleRefreshToken(writer http.ResponseWriter, request *http.Request) {
	authReq := new(RefreshTokenRequest)

	err := parseRequestMessage(request, authReq)

	if err != nil {
		handleError(writer, err)
		return
	}

	token, serr := m.service.RenewToken(request.Context(), authReq.Token, model.TokenTerm(model.TokenTerm_value[authReq.Term]))

	var resp = &RefreshTokenResponse{}

	if token != nil {
		resp.Token = Token{
			Term:       token.Term.String(),
			Content:    token.Content,
			Expiration: token.Expiration.AsTime(),
		}
	}

	ServiceResponder().
		Writer(writer).
		Respond(resp, serr)
}

func NewAuthenticationApi(service service.AuthenticationService) AuthenticationApi {
	return &authenticationApi{service: service}
}
