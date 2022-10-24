package api

import (
	"data-handler/service"
	"data-handler/stub"
	"data-handler/stub/model"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
	"net/http"
	"strconv"
	"strings"
)

type RecordApi interface {
	InjectRecordService(service service.RecordService)
	InjectResourceService(service service.ResourceService)
	ConfigureRouter(r *mux.Router)
}

type recordApi struct {
	recordService   service.RecordService
	resourceService service.ResourceService
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
	exists, err := r.resourceService.CheckResourceExists("", resourceName)
	if err != nil {
		log.Println(err)
		return false
	}
	return exists
}

func (r *recordApi) InjectRecordService(service service.RecordService) {
	r.recordService = service
}

func (r *recordApi) InjectResourceService(service service.ResourceService) {
	r.resourceService = service
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

	ServiceResponder[*stub.ListRecordRequest, *stub.ListRecordResponse]().
		Writer(writer).
		Request(request).
		ServiceCall(r.recordService.List).
		Payload(&stub.ListRecordRequest{
			Token:      getToken(request),
			Resource:   resourceName,
			Query:      query,
			Limit:      uint32(limit),
			Offset:     uint64(offset),
			UseHistory: getRequestBoolFlag(request, "useHistory"),
		}).
		Respond()
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
	record1.Type = model.DataType_USER

	if err != nil {
		handleClientError(writer, err)
		return
	}

	ServiceResponder[*stub.CreateRecordRequest, *stub.CreateRecordResponse]().
		Writer(writer).
		Request(request).
		ServiceCall(r.recordService.Create).
		Payload(&stub.CreateRecordRequest{
			Token:   getToken(request),
			Records: []*model.Record{record1},
		}).
		ResponseMapper(func(response *stub.CreateRecordResponse) proto.Message {
			return response.Records[0]
		}).
		Respond()
}

func (r *recordApi) handleRecordGet(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	resourceName := vars["resourceName"]
	id := vars["id"]

	ServiceResponder[*stub.GetRecordRequest, *stub.GetRecordResponse]().
		Writer(writer).
		Request(request).
		ServiceCall(r.recordService.Get).
		Payload(&stub.GetRecordRequest{
			Token:    getToken(request),
			Resource: resourceName,
			Id:       id,
		}).
		ResponseMapper(func(response *stub.GetRecordResponse) proto.Message {
			return response.Record
		}).
		Respond()
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

	ServiceResponder[*stub.UpdateRecordRequest, *stub.UpdateRecordResponse]().
		Writer(writer).
		Request(request).
		ServiceCall(r.recordService.Update).
		Payload(&stub.UpdateRecordRequest{
			Token:        getToken(request),
			Records:      []*model.Record{record},
			CheckVersion: getRequestBoolFlag(request, "checkVersion"),
		}).
		ResponseMapper(func(response *stub.UpdateRecordResponse) proto.Message {
			return response.Records[0]
		}).
		Respond()
}

func (r *recordApi) handleRecordDelete(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	resourceName := vars["resourceName"]
	id := vars["id"]

	ServiceResponder[*stub.DeleteRecordRequest, *stub.DeleteRecordResponse]().
		Writer(writer).
		Request(request).
		ServiceCall(r.recordService.Delete).
		Payload(&stub.DeleteRecordRequest{
			Token:    getToken(request),
			Resource: resourceName,
			Ids:      []string{id},
		}).
		Respond()
}

func (r *recordApi) handleRecordSearch(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	resourceName := vars["resourceName"]

	listRecordRequest := new(stub.ListRecordRequest)

	err := parseRequestMessage(request, listRecordRequest)
	listRecordRequest.Resource = resourceName

	if err != nil {
		handleClientError(writer, err)
		return
	}

	ServiceResponder[*stub.ListRecordRequest, *stub.ListRecordResponse]().
		Writer(writer).
		Request(request).
		ServiceCall(r.recordService.List).
		Payload(listRecordRequest).
		Respond()
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

func NewRecordApi() RecordApi {
	return &recordApi{}
}
