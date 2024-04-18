package rest

import (
	"context"
	"encoding/json"
	"github.com/apibrew/apibrew/pkg/api"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/formats/unstructured"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resource_model"
	"github.com/apibrew/apibrew/pkg/resource_model/extramappings"
	"github.com/apibrew/apibrew/pkg/service"
	"github.com/apibrew/apibrew/pkg/service/annotations"
	"github.com/apibrew/apibrew/pkg/util"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"strings"
)

type RecordApi interface {
	ConfigureRouter(r *mux.Router)
}

type recordApi struct {
	api             api.Interface
	resourceService service.ResourceService
	watchService    service.WatchService
}

func (r *recordApi) ConfigureRouter(router *mux.Router) {
	subRoute := router.MatcherFunc(r.matchFunc).Subrouter()
	// collection level operations
	subRoute.HandleFunc("/{resourceSlug}", r.handleRecordList).Methods("GET")
	subRoute.HandleFunc("/{resourceSlug}", r.handleRecordCreate).Methods("POST")
	subRoute.HandleFunc("/{resourceSlug}", r.handleRecordApply).Methods("PATCH")

	// internal actions
	subRoute.HandleFunc("/{resourceSlug}/_search", r.handleRecordSearch).Methods("POST")
	subRoute.HandleFunc("/{resourceSlug}/_load", r.handleRecordLoad).Methods("POST")
	subRoute.HandleFunc("/{resourceSlug}/_watch", r.handleRecordWatch).Methods("GET")
	subRoute.HandleFunc("/{resourceSlug}/_resource", r.handleRecordResource).Methods("GET")

	// record level operations
	subRoute.HandleFunc("/{resourceSlug}/{id}", r.handleRecordGet).Methods("GET")
	subRoute.HandleFunc("/{resourceSlug}/{id}", r.handleRecordUpdate).Methods("PUT")
	subRoute.HandleFunc("/{resourceSlug}/{id}", r.handleRecordDelete).Methods("DELETE")
}

func (r *recordApi) matchFunc(request *http.Request, match *mux.RouteMatch) bool {
	pathParts := strings.Split(request.URL.Path, "/")

	if len(pathParts) < 2 {
		return false
	}

	slug := pathParts[1]

	resource := r.resourceService.GetSchema().ResourceBySlug[slug]

	if resource == nil {
		return false
	}

	if annotations.IsEnabled(resource, annotations.RestApiDisabled) {
		return false
	}

	return true
}

func (r *recordApi) handleRecordList(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	resource := r.resourceService.GetSchema().ResourceBySlug[vars["resourceSlug"]]

	if annotations.IsEnabled(resource, annotations.ActionApi) {
		handleError(writer, errors.LogicalError.WithMessage("Get action is not supported for this resource."))
		return
	}

	resolveReferences := request.URL.Query().Get("resolve-references")

	// handle query parameters

	limit := 10
	offset := 0

	if request.URL.Query().Get("limit") != "" {
		var _err error
		limit, _err = strconv.Atoi(request.URL.Query().Get("limit"))

		if _err != nil {
			handleError(writer, _err)
			return
		}
	}

	if request.URL.Query().Get("offset") != "" {
		var _err error
		offset, _err = strconv.Atoi(request.URL.Query().Get("offset"))

		if _err != nil {
			handleError(writer, _err)
			return
		}
	}

	filters := r.makeFilters(request)

	result, serviceErr := r.api.List(r.prepareContext(request), api.ListParams{
		Type:              resource.Namespace + "/" + resource.Name,
		Filters:           filters,
		Limit:             uint32(limit),
		Offset:            uint64(offset),
		UseHistory:        getRequestBoolFlag(request, "useHistory"),
		ResolveReferences: strings.Split(resolveReferences, ","),
	})

	if serviceErr != nil {
		handleServiceError(writer, serviceErr)
		return
	}

	respondSuccess(writer, result)
}

func (r *recordApi) makeFilters(request *http.Request) map[string]string {
	filters := make(map[string]string)

	for key := range request.URL.Query() {
		if key == "limit" || key == "offset" || key == "resolve-references" || key == "useHistory" {
			continue
		}
		filters[key] = request.URL.Query().Get(key)
	}
	return filters
}

func (r *recordApi) handleRecordCreate(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	resource := r.resourceService.GetSchema().ResourceBySlug[vars["resourceSlug"]]

	record := new(unstructured.Unstructured)

	err := parseRequestMessage(request, record)

	if err != nil {
		handleError(writer, err)
		return
	}

	(*record)["type"] = resource.Namespace + "/" + resource.Name

	result, serviceErr := r.api.Create(r.prepareContext(request), *record)

	if serviceErr != nil {
		handleServiceError(writer, serviceErr)
		return
	}

	respondSuccess(writer, result)
}

func (r *recordApi) handleRecordApply(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	resource := r.resourceService.GetSchema().ResourceBySlug[vars["resourceSlug"]]

	if annotations.IsEnabled(resource, annotations.ActionApi) {
		handleError(writer, errors.LogicalError.WithMessage("Get action is not supported for this resource."))
		return
	}

	record := new(unstructured.Unstructured)

	err := parseRequestMessage(request, record)

	if err != nil {
		handleError(writer, err)
		return
	}

	(*record)["type"] = resource.Namespace + "/" + resource.Name

	result, serviceErr := r.api.Apply(r.prepareContext(request), *record)

	if serviceErr != nil {
		handleServiceError(writer, serviceErr)
		return
	}

	respondSuccess(writer, result)
}

func (r *recordApi) handleRecordGet(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	resource := r.resourceService.GetSchema().ResourceBySlug[vars["resourceSlug"]]

	if annotations.IsEnabled(resource, annotations.ActionApi) {
		handleError(writer, errors.LogicalError.WithMessage("Get action is not supported for this resource."))
		return
	}

	id := vars["id"]

	resolveReferences := request.URL.Query().Get("resolve-references")

	record, err := r.api.Load(r.prepareContext(request), map[string]interface{}{
		"type": resource.Namespace + "/" + resource.Name,
		"id":   id,
	}, api.LoadParams{
		UseHistory:        getRequestBoolFlag(request, "useHistory"),
		ResolveReferences: strings.Split(resolveReferences, ","),
	})

	if err != nil {
		panic(err)
	}

	respondSuccess(writer, record)
}

func (r *recordApi) handleRecordLoad(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	resource := r.resourceService.GetSchema().ResourceBySlug[vars["resourceSlug"]]

	if annotations.IsEnabled(resource, annotations.ActionApi) {
		handleError(writer, errors.LogicalError.WithMessage("Get action is not supported for this resource."))
		return
	}

	record := new(unstructured.Unstructured)

	err := parseRequestMessage(request, record)

	if err != nil {
		handleError(writer, err)
		return
	}

	(*record)["type"] = resource.Namespace + "/" + resource.Name

	resolveReferences := request.URL.Query().Get("resolve-references")

	result, err := r.api.Load(r.prepareContext(request), *record, api.LoadParams{
		UseHistory:        getRequestBoolFlag(request, "useHistory"),
		ResolveReferences: strings.Split(resolveReferences, ","),
	})

	if err != nil {
		panic(err)
	}

	respondSuccess(writer, result)
}

func (r *recordApi) handleRecordUpdate(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	resource := r.resourceService.GetSchema().ResourceBySlug[vars["resourceSlug"]]

	if annotations.IsEnabled(resource, annotations.ActionApi) {
		handleError(writer, errors.LogicalError.WithMessage("Get action is not supported for this resource."))
		return
	}

	record := new(unstructured.Unstructured)

	err := parseRequestMessage(request, record)

	if err != nil {
		handleError(writer, err)
		return
	}

	(*record)["type"] = resource.Namespace + "/" + resource.Name

	result, serviceErr := r.api.Update(r.prepareContext(request), *record)

	if serviceErr != nil {
		handleServiceError(writer, serviceErr)
		return
	}

	respondSuccess(writer, result)
}

func (r *recordApi) handleRecordDelete(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	resource := r.resourceService.GetSchema().ResourceBySlug[vars["resourceSlug"]]

	if annotations.IsEnabled(resource, annotations.ActionApi) {
		handleError(writer, errors.LogicalError.WithMessage("Get action is not supported for this resource."))
		return
	}

	id := vars["id"]

	serviceErr := r.api.Delete(r.prepareContext(request), unstructured.Unstructured{
		"type": resource.Namespace + "/" + resource.Name,
		"id":   id,
	})

	ServiceResponder().
		Writer(writer).
		Respond(nil, serviceErr)
}

func (r *recordApi) handleRecordSearch(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	resource := r.resourceService.GetSchema().ResourceBySlug[vars["resourceSlug"]]

	if annotations.IsEnabled(resource, annotations.ActionApi) {
		handleError(writer, errors.LogicalError.WithMessage("Get action is not supported for this resource."))
		return
	}

	listRecordRequest := new(api.ListParams)
	listRecordRequest.Type = resource.Namespace + "/" + resource.Name

	err := parseRequestMessage(request, listRecordRequest)

	if err != nil {
		handleError(writer, err)
		return
	}

	result, serviceErr := r.api.List(r.prepareContext(request), *listRecordRequest)

	if serviceErr != nil {
		handleServiceError(writer, serviceErr)
		return
	}

	respondSuccess(writer, result)
}

func (r *recordApi) handleRecordResource(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	resource := r.resourceService.GetSchema().ResourceBySlug[vars["resourceSlug"]]

	ServiceResponder().
		Writer(writer).
		Respond(extramappings.ResourceTo(resource), nil)
}

func (r *recordApi) handleRecordWatch(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	resource := r.resourceService.GetSchema().ResourceBySlug[vars["resourceSlug"]]

	filters := r.makeFilters(request)

	var useEventSource = request.URL.Query().Get("use-event-source") == "true"

	var recordSelector *model.BooleanExpression

	if len(filters) > 0 {
		var err errors.ServiceError
		recordSelector, err = util.PrepareQuery(resource, filters)

		if err != nil {
			handleServiceError(writer, err)
			return
		}
	}

	res, err := r.watchService.WatchResource(r.prepareContext(request), service.WatchParams{
		Selector: &model.EventSelector{
			Actions: []model.Event_Action{
				model.Event_CREATE,
				model.Event_UPDATE,
				model.Event_DELETE,
			},
			Namespaces:     []string{resource.Namespace},
			Resources:      []string{resource.Name},
			RecordSelector: recordSelector,
		},
	})

	if err != nil {
		handleServiceError(writer, err)
		return
	}

	if useEventSource {
		writer.Header().Set("Content-Type", "text/event-stream")
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

		if useEventSource {
			_, _ = writer.Write([]byte("data: "))
		}
		_, _ = writer.Write(data)

		_, _ = writer.Write([]byte("\n\n"))

		if f, ok := writer.(http.Flusher); ok {
			f.Flush()
		}

	}
}

func (r *recordApi) prepareContext(request *http.Request) context.Context {
	return annotations.SetWithContext(request.Context(), "requestUrl", request.URL.String())
}

func NewRecordApi(container service.Container) RecordApi {
	return &recordApi{
		api:             api.NewInterface(container),
		resourceService: container.GetResourceService(),
		watchService:    container.GetWatchService(),
	}
}
