package rest

import (
	"encoding/json"
	"github.com/apibrew/apibrew/pkg/service"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
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
	req, err := m.parseMetricsRequest(request)

	if err != nil {
		_, _ = writer.Write([]byte(err.Error()))
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	metrics, err := m.metricsService.GetMetrics(req)

	err = json.NewEncoder(writer).Encode(metrics)

	if err != nil {
		log.Print(err)

		_, _ = writer.Write([]byte(err.Error()))
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (m metricsApi) parseMetricsRequest(request *http.Request) (service.MetricsRequest, error) {
	var req service.MetricsRequest

	if request.URL.Query().Get("namespace") != "" {
		req.Namespace = new(string)
		*req.Namespace = request.URL.Query().Get("namespace")
	}

	if request.URL.Query().Get("resource") != "" {
		req.Resource = new(string)
		*req.Resource = request.URL.Query().Get("resource")
	}

	if request.URL.Query().Get("operation") != "" {
		req.Operation = new(service.MetricsOperation)
		*req.Operation = service.MetricsOperation(request.URL.Query().Get("operation"))
	}

	if request.URL.Query().Get("interval") != "" {
		req.Interval = new(service.MetricsInterval)
		*req.Interval = service.MetricsInterval(request.URL.Query().Get("interval"))
	}

	if request.URL.Query().Get("from") != "" {
		from, err := time.Parse(time.RFC3339, request.URL.Query().Get("from"))

		if err != nil {
			return service.MetricsRequest{}, err
		}

		req.From = &from
	}

	if request.URL.Query().Get("to") != "" {
		to, err := time.Parse(time.RFC3339, request.URL.Query().Get("to"))

		if err != nil {
			return service.MetricsRequest{}, err
		}

		req.To = &to
	}

	return req, nil
}

func NewMetricsApi(service service.MetricsService) MetricsApi {
	return &metricsApi{
		metricsService: service,
	}
}
