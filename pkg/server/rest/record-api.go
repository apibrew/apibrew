package rest

import (
	"context"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/tislib/data-handler/pkg/abs"
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/stub"
	"net/http"
	"strconv"
	"strings"
)

type RecordApi interface {
	ConfigureRouter(r *mux.Router)
}

type recordApi struct {
	recordService   abs.RecordService
	resourceService abs.ResourceService
}

func (r *recordApi) ConfigureRouter(router *mux.Router) {
	subRoute := router.MatcherFunc(r.matchFunc).Subrouter()
	// collection level operations
	subRoute.HandleFunc("/{resourceName}", r.handleRecordList).Methods("GET")
	subRoute.HandleFunc("/{resourceName}", r.handleRecordCreate).Methods("POST")

	// search
	subRoute.HandleFunc("/{resourceName}/_search", r.handleRecordSearch).Methods("POST")

	// collection bulk operations
	subRoute.HandleFunc("/{resourceName}/_bulk", r.handleRecordBatchCreate).Methods("POST")
	subRoute.HandleFunc("/{resourceName}/_bulk", r.handleRecordBatchUpdate).Methods("PUT")
	subRoute.HandleFunc("/{resourceName}/_bulk", r.handleRecordBatchDelete).Methods("DELETE")

	// record level operations
	subRoute.HandleFunc("/{resourceName}/{id}", r.handleRecordGet).Methods("GET")
	subRoute.HandleFunc("/{resourceName}/{id}", r.handleRecordUpdate).Methods("PUT")
	subRoute.HandleFunc("/{resourceName}/{id}", r.handleRecordDelete).Methods("DELETE")
}

func (r *recordApi) matchFunc(request *http.Request, match *mux.RouteMatch) bool {
	pathParts := strings.Split(request.URL.Path, "/")
	resourceName := pathParts[1]
	exists, err := r.resourceService.CheckResourceExists(context.TODO(), "default", resourceName)
	if err != nil {
		log.Println(err)
		return false
	}
	return exists
}

func (r *recordApi) handleRecordList(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	resourceName := vars["resourceName"]

	resource, err := r.resourceService.GetResourceByName(request.Context(), "", resourceName)

	if err != nil {
		handleClientError(writer, err)
		return
	}

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

	result, total, serviceErr := r.recordService.List(request.Context(), abs.RecordListParams{
		Query:      query,
		Namespace:  "default",
		Resource:   resourceName,
		Limit:      uint32(limit),
		Offset:     uint64(offset),
		UseHistory: getRequestBoolFlag(request, "useHistory"),
	})

	ServiceResponder[*stub.ListRecordRequest]().
		Writer(writer).
		Request(request).
		Respond(&stub.ListRecordResponse{
			Total:   total,
			Content: result,
		}, serviceErr)
}

func (r *recordApi) handleRecordCreate(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	resourceName := vars["resourceName"]

	record1 := new(model.Record)

	err := parseRequestMessage(request, record1)

	if err != nil {
		handleClientError(writer, err)
		return
	}

	record1.Resource = resourceName
	record1.DataType = model.DataType_USER

	if err != nil {
		handleClientError(writer, err)
		return
	}

	res, inserted, serviceErr := r.recordService.Create(request.Context(), abs.RecordCreateParams{
		Namespace:      "default",
		Records:        []*model.Record{record1},
		IgnoreIfExists: false,
	})

	ServiceResponder[*stub.CreateRecordRequest]().
		Writer(writer).
		Request(request).
		Respond(&stub.CreateRecordResponse{
			Records:  res,
			Inserted: inserted,
		}, serviceErr)
}

func (r *recordApi) handleRecordGet(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	resourceName := vars["resourceName"]
	id := vars["id"]

	record, serviceErr := r.recordService.Get(request.Context(), abs.RecordGetParams{
		Namespace: "default",
		Resource:  resourceName,
		Id:        id,
	})

	ServiceResponder[*stub.GetRecordRequest]().
		Writer(writer).
		Request(request).
		Respond(&stub.GetRecordResponse{
			Record: record,
		}, serviceErr)
}

func (r *recordApi) handleRecordUpdate(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	resourceName := vars["resourceName"]
	id := vars["id"]

	record := new(model.Record)

	err := parseRequestMessage(request, record)

	if err != nil {
		handleClientError(writer, err)
		return
	}

	record.Resource = resourceName
	record.Id = id

	result, serviceErr := r.recordService.Update(request.Context(), abs.RecordUpdateParams{
		Namespace:    "",
		Records:      []*model.Record{record},
		CheckVersion: false,
	})

	var updatedRecord *model.Record = nil

	if len(result) == 1 {
		updatedRecord = result[0]
	}

	ServiceResponder[*stub.UpdateRecordRequest]().
		Writer(writer).
		Request(request).
		Respond(updatedRecord, serviceErr)
}

func (r *recordApi) handleRecordDelete(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	resourceName := vars["resourceName"]
	id := vars["id"]

	serviceErr := r.recordService.Delete(request.Context(), abs.RecordDeleteParams{
		Namespace: "default",
		Resource:  resourceName,
		Ids:       []string{id},
	})

	ServiceResponder[*stub.DeleteRecordRequest]().
		Writer(writer).
		Request(request).
		Respond(&stub.DeleteRecordResponse{}, serviceErr)
}

func (r *recordApi) handleRecordSearch(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	resourceName := vars["resourceName"]

	listRecordRequest := new(stub.SearchRecordRequest)

	err := parseRequestMessage(request, listRecordRequest)
	listRecordRequest.Resource = resourceName

	if err != nil {
		handleClientError(writer, err)
		return
	}

	result, total, serviceErr := r.recordService.List(request.Context(), abs.RecordListParams{
		Query:      listRecordRequest.Query,
		Namespace:  "default",
		Resource:   listRecordRequest.Resource,
		Limit:      listRecordRequest.Limit,
		Offset:     listRecordRequest.Offset,
		UseHistory: listRecordRequest.UseHistory,
	})

	ServiceResponder[*stub.ListRecordRequest]().
		Writer(writer).
		Request(request).
		Respond(&stub.ListRecordResponse{
			Total:   total,
			Content: result,
		}, serviceErr)
}

func (r *recordApi) handleRecordBatchDelete(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Not implemented"))
}

func (r *recordApi) handleRecordBatchUpdate(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Not implemented"))
}

func (r *recordApi) handleRecordBatchCreate(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Not implemented"))
}

func NewRecordApi(recordService abs.RecordService, resourceService abs.ResourceService) RecordApi {
	return &recordApi{
		recordService:   recordService,
		resourceService: resourceService,
	}
}