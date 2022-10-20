package api

import (
	"data-handler/service"
	"data-handler/stub"
	"data-handler/stub/model"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
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
	subRoute.HandleFunc("/{resourceName}/_search", r.handleRecordSearch).Methods("POST")
	subRoute.HandleFunc("/{resourceName}", r.handleRecordCreate).Methods("POST")
	subRoute.HandleFunc("/{resourceName}", r.handleRecordBatchUpdate).Methods("PUT")
	subRoute.HandleFunc("/{resourceName}", r.handleRecordBatchDelete).Methods("DELETE")

	// record level operations
	subRoute.HandleFunc("/{resourceName}/{id}", r.handleRecordGet).Methods("GET")
	subRoute.HandleFunc("/{resourceName}/{id}", r.handleRecordUpdate).Methods("PUT")
	subRoute.HandleFunc("/{resourceName}/{id}", r.handleRecordDelete).Methods("DELETE")
}

func (r *recordApi) matchFunc(request *http.Request, match *mux.RouteMatch) bool {
	pathParts := strings.Split(request.URL.Path, "/")
	resourceName := pathParts[1]
	exists, err := r.resourceService.CheckResourceExists(resourceName)
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

	respondMessage(writer)(r.recordService.List(request.Context(), &stub.ListRecordRequest{
		Token:    getToken(request),
		Resource: resourceName,
		Query:    nil,
		Limit:    10,
		Offset:   0,
	}))
}

func (r *recordApi) handleRecordCreate(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	resourceName := vars["resourceName"]

	data, err := io.ReadAll(request.Body)

	if err != nil {
		log.Error(err)
		writer.WriteHeader(400)
		return
	}

	record1 := new(model.Record)

	err = umo.Unmarshal(data, record1)

	record1.Resource = resourceName
	record1.Type = model.DataType_USER

	if err != nil {
		log.Error(err)
		writer.WriteHeader(400)
		return
	}

	respondMessage(writer)(r.recordService.Create(request.Context(), &stub.CreateRecordRequest{
		Token:   getToken(request),
		Records: []*model.Record{record1},
	}))
}

func (r *recordApi) handleRecordGet(writer http.ResponseWriter, request *http.Request) {

}

func (r *recordApi) handleRecordUpdate(writer http.ResponseWriter, request *http.Request) {

}

func (r *recordApi) handleRecordDelete(writer http.ResponseWriter, request *http.Request) {

}

func (r *recordApi) handleRecordBatchDelete(writer http.ResponseWriter, request *http.Request) {

}

func (r *recordApi) handleRecordSearch(writer http.ResponseWriter, request *http.Request) {

}

func (r *recordApi) handleRecordBatchUpdate(writer http.ResponseWriter, request *http.Request) {

}

func NewRecordApi() RecordApi {
	return &recordApi{}
}
