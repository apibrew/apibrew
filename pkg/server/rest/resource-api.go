package rest

import (
	"encoding/json"
	"github.com/apibrew/apibrew/pkg/core"
	"github.com/apibrew/apibrew/pkg/resource_model"
	"github.com/apibrew/apibrew/pkg/resource_model/extramappings"
	"github.com/apibrew/apibrew/pkg/resources"
	"github.com/apibrew/apibrew/pkg/service"
	"github.com/apibrew/apibrew/pkg/util"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type ResourceApi interface {
	ConfigureRouter(r *mux.Router)
}

type resourceApi struct {
	resourceService service.ResourceService
	watchService    service.WatchService
}

func (r *resourceApi) ConfigureRouter(router *mux.Router) {
	subRoute := router.PathPrefix("/resources").Subrouter()
	// collection level operations
	subRoute.HandleFunc("", r.handleResourceList).Methods("GET")
	subRoute.HandleFunc("", r.handleResourceCreate).Methods("POST")
	subRoute.HandleFunc("/", r.handleResourceList).Methods("GET")
	subRoute.HandleFunc("/", r.handleResourceCreate).Methods("POST")

	// resource level operations
	subRoute.HandleFunc("/_watch", r.handleResourceWatch).Methods("GET")
	subRoute.HandleFunc("/{id}", r.handleResourceGet).Methods("GET")
	subRoute.HandleFunc("/{id}", r.handleResourceUpdate).Methods("PUT")
	subRoute.HandleFunc("/{id}", r.handleResourceDelete).Methods("DELETE")

	// by name
	subRoute.HandleFunc("/by-name/{namespace}/{name}", r.handleResourceByName).Methods("GET")
}

func (r *resourceApi) handleResourceList(writer http.ResponseWriter, request *http.Request) {
	var resources, err = r.resourceService.List(request.Context())

	ServiceResponder().
		Writer(writer).
		Respond(map[string]interface{}{
			"total":   len(resources),
			"content": util.ArrayMap(resources, extramappings.ResourceTo),
		}, err)
}

func (r *resourceApi) handleResourceCreate(writer http.ResponseWriter, request *http.Request) {
	rw := new(resource_model.Resource)

	err := parseRequestMessage(request, rw)

	if err != nil {
		handleError(writer, err)
		return
	}

	var forceMigrate = request.Header.Get("X-Force-Migrate") == "true" || request.URL.Query().Get("forceMigrate") == "true"

	res, serviceErr := r.resourceService.Create(request.Context(), extramappings.ResourceFrom(rw), true, forceMigrate)

	ServiceResponder().
		Writer(writer).
		Respond(extramappings.ResourceTo(res), serviceErr)
}

func (r *resourceApi) handleResourceGet(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]

	resource, serviceErr := r.resourceService.Get(request.Context(), id)

	ServiceResponder().
		Writer(writer).
		Respond(extramappings.ResourceTo(resource), serviceErr)
}

func (r *resourceApi) handleResourceByName(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	namespace := vars["namespace"]
	name := vars["name"]

	resource, serviceErr := r.resourceService.GetResourceByName(request.Context(), namespace, name)

	ServiceResponder().
		Writer(writer).
		Respond(extramappings.ResourceTo(resource), serviceErr)
}

func (r *resourceApi) handleResourceUpdate(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]

	resourceForUpdate := new(resource_model.Resource)

	err := parseRequestMessage(request, resourceForUpdate)

	if err != nil {
		handleError(writer, err)
		return
	}

	resource, serviceErr := r.resourceService.Get(request.Context(), id)

	if serviceErr != nil {
		ServiceResponder().
			Writer(writer).
			Respond(nil, serviceErr)
		return
	}

	resource.Id = id

	var forceMigrate = request.Header.Get("X-Force-Migrate") == "true" || request.URL.Query().Get("forceMigrate") == "true"

	serviceErr = r.resourceService.Update(request.Context(), extramappings.ResourceFrom(resourceForUpdate), true, forceMigrate)

	if serviceErr != nil {
		resource = nil
	}

	ServiceResponder().
		Writer(writer).
		Respond(extramappings.ResourceTo(resource), serviceErr)
}

func (r *resourceApi) handleResourceDelete(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]

	var forceMigrate = request.Header.Get("X-Force-Migrate") == "true" || request.URL.Query().Get("forceMigrate") == "true"

	serviceErr := r.resourceService.Delete(request.Context(), []string{id}, true, forceMigrate)

	ServiceResponder().
		Writer(writer).
		Respond(nil, serviceErr)
}

func (r *resourceApi) handleResourceWatch(writer http.ResponseWriter, request *http.Request) {
	res, err := r.watchService.WatchResource(request.Context(), service.WatchParams{
		Selector: &core.EventSelector{
			Actions: []core.Event_Action{
				core.Event_CREATE,
				core.Event_UPDATE,
				core.Event_DELETE,
			},
			Namespaces: []string{resources.ResourceResource.Namespace},
			Resources:  []string{resources.ResourceResource.Name},
		},
	})

	if err != nil {
		handleServiceError(writer, err)
		return
	}

	writer.WriteHeader(200)

	for eventProto := range res {
		select {
		case <-request.Context().Done():
			return
		default:
		}

		event := extramappings.EventFromProto(eventProto)

		if eventProto.Resource != nil {
			event.Resource = &resource_model.Resource{
				Name: eventProto.Resource.Name,
				Namespace: &resource_model.Namespace{
					Name: eventProto.Resource.Namespace,
				},
			}
		}

		data, err := json.Marshal(event)

		if err != nil {
			log.Print(err)
			return
		}

		_, _ = writer.Write(data)

		_, _ = writer.Write([]byte("\n"))

		if f, ok := writer.(http.Flusher); ok {
			f.Flush()
		}

	}
}

func NewResourceApi(container service.Container) ResourceApi {
	return &resourceApi{
		resourceService: container.GetResourceService(),
		watchService:    container.GetWatchService(),
	}
}
