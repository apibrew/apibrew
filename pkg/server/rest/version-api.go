package rest

import (
	"github.com/apibrew/apibrew/module"
	"github.com/gorilla/mux"
	"net/http"
)

type VersionApi interface {
	ConfigureRouter(r *mux.Router)
}

type versionApi struct {
}

func (m versionApi) ConfigureRouter(router *mux.Router) {
	subRoute := router.PathPrefix("/_version").Subrouter()
	// collection level operations
	subRoute.HandleFunc("", m.handleVersion).Methods("GET")
}

type VersionResponse struct {
	Version string            `json:"version"`
	Modules map[string]string `json:"modules"`
}

func (m versionApi) handleVersion(writer http.ResponseWriter, request *http.Request) {
	resp := VersionResponse{
		Version: module.Version,
		Modules: module.Modules,
	}

	respondSuccess(writer, resp)
}

func NewVersionApi() VersionApi {
	return &versionApi{}
}
