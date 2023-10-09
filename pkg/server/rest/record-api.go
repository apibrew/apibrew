package rest

import (
	"github.com/apibrew/apibrew/pkg/formats/unstructured"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resource_model/extramappings"
	"github.com/apibrew/apibrew/pkg/service"
	"github.com/apibrew/apibrew/pkg/service/annotations"
	"github.com/apibrew/apibrew/pkg/util"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"strings"
)

type RecordApi interface {
	ConfigureRouter(r *mux.Router)
}

type recordApi struct {
	recordService   service.RecordService
	resourceService service.ResourceService
}

func (r *recordApi) ConfigureRouter(router *mux.Router) {
	subRoute := router.MatcherFunc(r.matchFunc).Subrouter()
	// collection level operations
	subRoute.HandleFunc("/{resourceSlug}", r.handleRecordList).Methods("GET")
	subRoute.HandleFunc("/{resourceSlug}", r.handleRecordCreate).Methods("POST")
	subRoute.HandleFunc("/{resourceSlug}", r.handleRecordApply).Methods("PATCH")

	// search
	subRoute.HandleFunc("/{resourceSlug}/_search", r.handleRecordSearch).Methods("POST")
	subRoute.HandleFunc("/{resourceSlug}/_resource", r.handleRecordResource).Methods("GET")
	subRoute.HandleFunc("/{resourceSlug}/{id}/_{action}", r.handleAction).Methods("GET")
	subRoute.HandleFunc("/{resourceSlug}/{id}/_{action}", r.handleAction).Methods("POST")

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

	resolveReferences := request.URL.Query().Get("resolve-references")

	// handle query parameters

	queryMap := make(map[string]interface{})

	for key := range request.URL.Query() {
		queryMap[key] = request.URL.Query().Get(key)
	}

	query, srvErr := r.recordService.PrepareQuery(resource, queryMap)

	if srvErr != nil {
		handleClientError(writer, srvErr)
		return
	}

	limit := 10
	offset := 0

	if request.URL.Query().Get("limit") != "" {
		var _err error
		limit, _err = strconv.Atoi(request.URL.Query().Get("limit"))

		if _err != nil {
			handleClientError(writer, _err)
			return
		}
	}

	if request.URL.Query().Get("offset") != "" {
		var _err error
		offset, _err = strconv.Atoi(request.URL.Query().Get("offset"))

		if _err != nil {
			handleClientError(writer, _err)
			return
		}
	}

	result, total, serviceErr := r.recordService.List(request.Context(), service.RecordListParams{
		Query:             query,
		Namespace:         resource.Namespace,
		Resource:          resource.Name,
		Limit:             uint32(limit),
		Offset:            uint64(offset),
		UseHistory:        getRequestBoolFlag(request, "useHistory"),
		ResolveReferences: strings.Split(resolveReferences, ","),
	})

	if serviceErr != nil {
		handleServiceError(writer, serviceErr)
		return
	}

	respondSuccess(writer, &RecordList{
		Total:   uint64(total),
		Records: util.ArrayMap(result, NewRecordWrapper),
	})
}

func (r *recordApi) handleRecordCreate(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	resource := r.resourceService.GetSchema().ResourceBySlug[vars["resourceSlug"]]

	record1 := NewEmptyRecordWrapper()

	err := parseRequestMessage(request, record1)

	if err != nil {
		handleClientError(writer, err)
		return
	}

	res, serviceErr := r.recordService.Create(request.Context(), service.RecordCreateParams{
		Namespace: resource.Namespace,
		Resource:  resource.Name,
		Records:   []*model.Record{record1.toRecord()},
	})

	if serviceErr != nil {
		handleServiceError(writer, serviceErr)
		return
	}

	var createdRecord *model.Record = nil

	if len(res) > 0 {
		createdRecord = res[0]
	}

	respondSuccess(writer, NewRecordWrapper(createdRecord))
}

func (r *recordApi) handleRecordApply(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	resource := r.resourceService.GetSchema().ResourceBySlug[vars["resourceSlug"]]

	record1 := NewEmptyRecordWrapper()

	err := parseRequestMessage(request, record1)

	if err != nil {
		handleClientError(writer, err)
		return
	}

	if err != nil {
		handleClientError(writer, err)
		return
	}

	res, serviceErr := r.recordService.Apply(request.Context(), service.RecordUpdateParams{
		Namespace: resource.Namespace,
		Resource:  resource.Name,
		Records:   []*model.Record{record1.toRecord()},
	})

	if serviceErr != nil {
		handleServiceError(writer, serviceErr)
		return
	}

	var appliedRecord *model.Record = nil

	if len(res) > 0 {
		appliedRecord = res[0]
	}

	respondSuccess(writer, NewRecordWrapper(appliedRecord))
}

func (r *recordApi) handleRecordGet(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	resource := r.resourceService.GetSchema().ResourceBySlug[vars["resourceSlug"]]
	id := vars["id"]

	resolveReferences := request.URL.Query().Get("resolve-references")

	record, serviceErr := r.recordService.Get(request.Context(), service.RecordGetParams{
		Namespace:         resource.Namespace,
		Resource:          resource.Name,
		Id:                id,
		ResolveReferences: strings.Split(resolveReferences, ","),
	})

	if serviceErr != nil {
		handleServiceError(writer, serviceErr)
		return
	}

	respondSuccess(writer, NewRecordWrapper(record))
}

func (r *recordApi) handleRecordUpdate(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	resource := r.resourceService.GetSchema().ResourceBySlug[vars["resourceSlug"]]
	id := vars["id"]

	recordWrap := NewEmptyRecordWrapper()

	err := parseRequestMessage(request, recordWrap)

	record := recordWrap.toRecord()

	if err != nil {
		handleClientError(writer, err)
		return
	}

	record.Id = id

	result, serviceErr := r.recordService.Update(request.Context(), service.RecordUpdateParams{
		Namespace: resource.Namespace,
		Resource:  resource.Name,
		Records:   []*model.Record{record},
	})

	if serviceErr != nil {
		handleServiceError(writer, serviceErr)
		return
	}

	var updatedRecord *model.Record = nil

	if len(result) == 1 {
		updatedRecord = result[0]
	}

	respondSuccess(writer, NewRecordWrapper(updatedRecord))
}

func (r *recordApi) handleRecordDelete(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	resource := r.resourceService.GetSchema().ResourceBySlug[vars["resourceSlug"]]
	id := vars["id"]

	serviceErr := r.recordService.Delete(request.Context(), service.RecordDeleteParams{
		Namespace: resource.Namespace,
		Resource:  resource.Name,
		Ids:       []string{id},
	})

	ServiceResponder().
		Writer(writer).
		Respond(nil, serviceErr)
}

func (r *recordApi) handleRecordSearch(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	resource := r.resourceService.GetSchema().ResourceBySlug[vars["resourceSlug"]]

	listRecordRequest := new(SearchRecordRequest)

	err := parseRequestMessage(request, listRecordRequest)

	if err != nil {
		handleClientError(writer, err)
		return
	}

	var query *model.BooleanExpression

	if listRecordRequest.Query != nil {
		query = extramappings.BooleanExpressionToProto(*listRecordRequest.Query)
	}

	result, total, serviceErr := r.recordService.List(request.Context(), service.RecordListParams{
		Query:             query,
		Namespace:         resource.Namespace,
		Resource:          resource.Name,
		Limit:             listRecordRequest.Limit,
		Offset:            listRecordRequest.Offset,
		UseHistory:        listRecordRequest.UseHistory,
		ResolveReferences: listRecordRequest.ResolveReferences,
	})

	if serviceErr != nil {
		handleServiceError(writer, serviceErr)
		return
	}

	respondSuccess(writer, &RecordList{
		Total:   uint64(total),
		Records: util.ArrayMap(result, NewRecordWrapper),
	})
}

func (r *recordApi) handleRecordResource(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	resource := r.resourceService.GetSchema().ResourceBySlug[vars["resourceSlug"]]

	ServiceResponder().
		Writer(writer).
		Respond(resourceTo(resource), nil)
}

func (r *recordApi) handleAction(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	resource := r.resourceService.GetSchema().ResourceBySlug[vars["resourceSlug"]]

	var input = new(unstructured.Unstructured)

	serr := parseRequestMessage(request, input)

	if serr != nil {
		handleClientError(writer, serr)
		return
	}

	result, err := r.recordService.ExecuteAction(request.Context(), service.ExecuteActionParams{
		Namespace:  resource.Namespace,
		Resource:   resource.Name,
		Id:         vars["id"],
		ActionName: vars["action"],
		Input:      *input,
	})

	if err != nil {
		handleClientError(writer, err)
		return
	}

	respondSuccess(writer, result)
}

func NewRecordApi(container service.Container) RecordApi {
	return &recordApi{
		recordService:   container.GetRecordService(),
		resourceService: container.GetResourceService(),
	}
}
