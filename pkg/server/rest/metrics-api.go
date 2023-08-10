package rest

import (
	"github.com/apibrew/apibrew/pkg/service"
	"github.com/gorilla/mux"
	"net/http"
)

type MetricsApi interface {
	ConfigureRouter(r *mux.Router)
}

type metricsApi struct {
	metricsService service.MetricsService
}

func (m metricsApi) ConfigureRouter(router *mux.Router) {
	subRoute := router.PathPrefix("/metrics").Subrouter()
	// collection level operations
	subRoute.HandleFunc("", m.handleMetrics).Methods("GET")
}

func (m metricsApi) handleMetrics(writer http.ResponseWriter, request *http.Request) {

}

func NewMetricsApi(service service.MetricsService) MetricsApi {
	return &metricsApi{
		metricsService: service,
	}
}
