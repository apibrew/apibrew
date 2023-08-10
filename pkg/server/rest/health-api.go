package rest

import (
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type HealthApi interface {
	ConfigureRouter(r *mux.Router)
}

type healthApi struct {
}

func (m healthApi) ConfigureRouter(router *mux.Router) {
	subRoute := router.PathPrefix("/health").Subrouter()
	// collection level operations
	subRoute.HandleFunc("", m.handleHealth).Methods("GET")
}

func (m healthApi) handleHealth(writer http.ResponseWriter, request *http.Request) {
	var response = &HealthResponse{
		Status: "ok",
	}

	err := json.NewEncoder(writer).Encode(response)

	if err != nil {
		log.Warn("Error encoding health response: ", err)
	}
}

func NewHealthApi() HealthApi {
	return &healthApi{}
}
