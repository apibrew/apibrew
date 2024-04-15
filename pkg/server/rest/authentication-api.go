package rest

import (
	"github.com/apibrew/apibrew/pkg/errors"
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
	subRoute.HandleFunc("/oauth2", m.handleOauth2).Methods("POST")
	subRoute.HandleFunc("", m.handleRefreshToken).Methods("PUT")
	subRoute.HandleFunc("", m.handleGetToken).Methods("GET")
}

func (m authenticationApi) handleOauth2(writer http.ResponseWriter, request *http.Request) {
	if err := request.ParseForm(); err != nil {
		handleError(writer, err)
		return
	}
	var params = request.Form

	var clientId = params.Get("client_id")
	var clientSecret = params.Get("client_secret")
	var grantType = params.Get("grant_type")
	//var code = params.Get("code")
	//var redirectUri = params.Get("redirect_uri")
	//var refreshToken = params.Get("refresh_token")
	//var scope = params.Get("scope")
	var term = params.Get("term")
	var minimizeToken = params.Get("minimize") == "true"

	if term == "" {
		term = "MIDDLE"
	}

	switch grantType {
	case "client_credentials":
		{
			token, serr := m.service.Authenticate(request.Context(), clientId, clientSecret, model.TokenTerm(model.TokenTerm_value[term]), minimizeToken)

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
				Respond(map[string]interface{}{
					"access_token": token.Content,
				}, serr)
		}
	default:
		ServiceResponder().
			Writer(writer).
			Respond(nil, errors.RecordValidationError.WithMessage("unsupported grant type: "+grantType))
	}
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

func (m authenticationApi) handleGetToken(writer http.ResponseWriter, request *http.Request) {
	token, serr := m.service.GetToken(request.Context())

	ServiceResponder().
		Writer(writer).
		Respond(token, serr)
}

func NewAuthenticationApi(service service.AuthenticationService) AuthenticationApi {
	return &authenticationApi{service: service}
}
