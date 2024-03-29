package rest

import (
	"encoding/json"
	"fmt"
	"github.com/apibrew/apibrew/pkg/resource_model"
	"github.com/apibrew/apibrew/pkg/resource_model/extramappings"
	"github.com/apibrew/apibrew/pkg/service"
	"github.com/gorilla/mux"
	"net/http"
)

type EventChannelApi interface {
	ConfigureRouter(r *mux.Router)
}

type eventChannelApi struct {
	service service.EventChannelService
}

func (r *eventChannelApi) ConfigureRouter(router *mux.Router) {
	router.HandleFunc("/_events", r.pollEvents).Methods("GET")
	router.HandleFunc("/_events", r.writeEvent).Methods("POST")
}

func (r *eventChannelApi) pollEvents(writer http.ResponseWriter, request *http.Request) {
	channelKey := request.URL.Query().Get("channelKey")

	if channelKey == "" {
		handleError(writer, fmt.Errorf("channelKey is required"))
		return
	}

	events, err := r.service.PollEvents(request.Context(), channelKey)

	if err != nil {
		handleError(writer, err)
		return
	}

	writer.WriteHeader(200)

	for eventProto := range events {
		event := extramappings.EventFromProto(eventProto)

		if eventProto.Resource != nil {
			event.Resource = &resource_model.Resource{
				Name: eventProto.Resource.Name,
				Namespace: &resource_model.Namespace{
					Name: eventProto.Resource.Namespace,
				},
			}
		}

		select {
		case <-request.Context().Done():
			return
		default:
		}

		data, err := json.Marshal(event)

		if err != nil {
			handleError(writer, err)
			return
		}

		_, _ = writer.Write(data)

		_, _ = writer.Write([]byte("\n"))

		if f, ok := writer.(http.Flusher); ok {
			f.Flush()
		}

	}
}

func (r *eventChannelApi) writeEvent(writer http.ResponseWriter, request *http.Request) {
	channelKey := request.URL.Query().Get("channelKey")

	if channelKey == "" {
		handleError(writer, fmt.Errorf("channelKey is required"))
		return
	}

	var event = &resource_model.Event{}

	err := json.NewDecoder(request.Body).Decode(event)

	if err != nil {
		handleError(writer, err)
		return
	}

	err = r.service.WriteEvent(request.Context(), extramappings.EventToProto(event))

	if err != nil {
		handleError(writer, err)
		return
	}

	writer.WriteHeader(200)
}

func NewEventChannelApi(container service.Container) EventChannelApi {
	return &eventChannelApi{
		service: container.GetEventChannelService(),
	}
}
