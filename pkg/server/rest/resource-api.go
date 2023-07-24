package rest

import (
	"github.com/apibrew/apibrew/pkg/service"
	"github.com/apibrew/apibrew/pkg/stub"
	"github.com/apibrew/apibrew/pkg/util"
	"github.com/gorilla/mux"
	"net/http"
)

type ResourceApi interface {
	ConfigureRouter(r *mux.Router)
}

type resourceApi struct {
	resourceService service.ResourceService
}

func (r *resourceApi) ConfigureRouter(router *mux.Router) {
	subRoute := router.PathPrefix("/resources").Subrouter()
	// collection level operations
	subRoute.HandleFunc("", r.handleResourceList).Methods("GET")
	subRoute.HandleFunc("", r.handleResourceCreate).Methods("POST")
	subRoute.HandleFunc("/", r.handleResourceList).Methods("GET")
	subRoute.HandleFunc("/", r.handleResourceCreate).Methods("POST")

	// resource level operations
	subRoute.HandleFunc("/{id}", r.handleResourceGet).Methods("GET")
	subRoute.HandleFunc("/{id}", r.handleResourceUpdate).Methods("PUT")
	subRoute.HandleFunc("/{id}", r.handleResourceDelete).Methods("DELETE")

	// by name
	subRoute.HandleFunc("/resources/by-name/{namespace}/{name}", r.handleResourceByName).Methods("GET")
}

func (r *resourceApi) handleResourceList(writer http.ResponseWriter, request *http.Request) {
	var resources, err = r.resourceService.List(request.Context())

	ServiceResponder[*stub.ListResourceRequest]().
		Writer(writer).
		Request(request).
		Respond(util.ArrayMap(resources, NewResourceWrapper), err)
}

func (r *resourceApi) handleResourceCreate(writer http.ResponseWriter, request *http.Request) {
	rw := new(ResourceWrapper)

	err := parseRequestMessage(request, rw)

	if err != nil {
		handleClientError(writer, err)
		return
	}

	res, serviceErr := r.resourceService.Create(request.Context(), rw.resource, true, false)

	ServiceResponder[*stub.CreateResourceRequest]().
		Writer(writer).
		Request(request).
		Respond(NewResourceWrapper(res), serviceErr)
}

func (r *resourceApi) handleResourceGet(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	resource := r.resourceService.GetSchema().ResourceBySlug[vars["resourceSlug"]]
	id := vars["id"]

	resource, serviceErr := r.resourceService.Get(request.Context(), id)

	ServiceResponder[*stub.GetResourceRequest]().
		Writer(writer).
		Request(request).
		Respond(NewResourceWrapper(resource), serviceErr)
}

func (r *resourceApi) handleResourceByName(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	resource := r.resourceService.GetSchema().ResourceBySlug[vars["resourceSlug"]]
	namespace := vars["namespace"]
	name := vars["name"]

	resource, serviceErr := r.resourceService.GetResourceByName(request.Context(), namespace, name)

	ServiceResponder[*stub.GetResourceRequest]().
		Writer(writer).
		Request(request).
		Respond(NewResourceWrapper(resource), serviceErr)
}

func (r *resourceApi) handleResourceUpdate(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	resource := r.resourceService.GetSchema().ResourceBySlug[vars["resourceSlug"]]
	id := vars["id"]

	resourceWrap := new(ResourceWrapper)

	err := parseRequestMessage(request, resourceWrap)

	if err != nil {
		handleClientError(writer, err)
		return
	}

	resource.Id = id

	serviceErr := r.resourceService.Update(request.Context(), resource, true, false)

	if serviceErr != nil {
		resource = nil
	}

	ServiceResponder[*stub.UpdateResourceRequest]().
		Writer(writer).
		Request(request).
		Respond(NewResourceWrapper(resource), serviceErr)
}

func (r *resourceApi) handleResourceDelete(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]

	serviceErr := r.resourceService.Delete(request.Context(), []string{id}, true, false)

	ServiceResponder[*stub.DeleteResourceRequest]().
		Writer(writer).
		Request(request).
		Respond(nil, serviceErr)
}

func NewResourceApi(container service.Container) ResourceApi {
	return &resourceApi{
		resourceService: container.GetResourceService(),
	}
}
