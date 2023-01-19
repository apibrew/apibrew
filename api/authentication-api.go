package api

import (
	"data-handler/grpc/stub"
	"data-handler/model"
	"data-handler/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"strings"
)

type LoginData struct {
	Username string
	Password string
	Term     model.TokenTerm
}

type AuthenticationApi interface {
	ConfigureRouter(r *mux.Router)
	AuthenticationMiddleWare(http.Handler) http.Handler
}

type authenticationApi struct {
	authenticationService service.AuthenticationService
	initData              *model.InitData
}

func (r *authenticationApi) AuthenticationMiddleWare(next http.Handler) http.Handler {
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

func NewAuthenticationApi(authenticationService service.AuthenticationService, initData *model.InitData) AuthenticationApi {
	return &authenticationApi{authenticationService: authenticationService, initData: initData}
}
