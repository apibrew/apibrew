package rest

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resource_model"
	"github.com/apibrew/apibrew/pkg/resources/mapping"
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
	subRoute.HandleFunc("/by-name/{namespace}/{name}", r.handleResourceByName).Methods("GET")
}

func (r *resourceApi) handleResourceList(writer http.ResponseWriter, request *http.Request) {
	var resources, err = r.resourceService.List(request.Context())

	ServiceResponder[*stub.ListResourceRequest]().
		Writer(writer).
		Request(request).
		Respond(util.ArrayMap(resources, resourceTo), err)
}

func resourceTo(resource *model.Resource) *resource_model.Resource {
	if resource == nil {
		return nil
	}
	resourceRec := mapping.ResourceToRecord(resource)
	return resource_model.ResourceMapperInstance.FromRecord(resourceRec)
}

func resourceFrom(resource *resource_model.Resource) *model.Resource {
	if resource == nil {
		return nil
	}
	resourceRec := resource_model.ResourceMapperInstance.ToRecord(resource)
	return mapping.ResourceFromRecord(resourceRec)
}

func (r *resourceApi) handleResourceCreate(writer http.ResponseWriter, request *http.Request) {
	rw := new(resource_model.Resource)

	err := parseRequestMessage(request, rw)

	if err != nil {
		handleClientError(writer, err)
		return
	}

	res, serviceErr := r.resourceService.Create(request.Context(), resourceFrom(rw), true, true)

	ServiceResponder[*stub.CreateResourceRequest]().
		Writer(writer).
		Request(request).
		Respond(resourceTo(res), serviceErr)
}

func (r *resourceApi) handleResourceGet(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]

	resource, serviceErr := r.resourceService.Get(request.Context(), id)

	ServiceResponder[*stub.GetResourceRequest]().
		Writer(writer).
		Request(request).
		Respond(resourceTo(resource), serviceErr)
}

func (r *resourceApi) handleResourceByName(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	namespace := vars["namespace"]
	name := vars["name"]

	resource, serviceErr := r.resourceService.GetResourceByName(request.Context(), namespace, name)

	ServiceResponder[*stub.GetResourceRequest]().
		Writer(writer).
		Request(request).
		Respond(resourceTo(resource), serviceErr)
}

func (r *resourceApi) handleResourceUpdate(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]

	resourceForUpdate := new(resource_model.Resource)

	err := parseRequestMessage(request, resourceForUpdate)

	if err != nil {
		handleClientError(writer, err)
		return
	}

	resource, serviceErr := r.resourceService.Get(request.Context(), id)

	if serviceErr != nil {
		ServiceResponder[*stub.UpdateResourceRequest]().
			Writer(writer).
			Request(request).
			Respond(nil, serviceErr)
		return
	}

	resource.Id = id

	serviceErr = r.resourceService.Update(request.Context(), resourceFrom(resourceForUpdate), true, true)

	if serviceErr != nil {
		resource = nil
	}

	ServiceResponder[*stub.UpdateResourceRequest]().
		Writer(writer).
		Request(request).
		Respond(resourceTo(resource), serviceErr)
}

func (r *resourceApi) handleResourceDelete(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]

	serviceErr := r.resourceService.Delete(request.Context(), []string{id}, true, true)

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
