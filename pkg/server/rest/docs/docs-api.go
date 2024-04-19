package docs

import (
	"github.com/apibrew/apibrew/pkg/service"
	"github.com/gorilla/mux"
)

type api struct {
	resourceService    service.ResourceService
	swaggerApi         SwaggerApi
	typescriptTypesApi TypescriptTypesApi
}

func (a *api) ConfigureRouter(r *mux.Router) {
	r.PathPrefix("/docs").Handler(a.Handler())
}

type Api interface {
	Handler() *mux.Router
	ConfigureRouter(r *mux.Router)
}

func (a *api) Handler() *mux.Router {
	var r = mux.NewRouter()

	a.swaggerApi.ConfigureRouter(r)
	a.typescriptTypesApi.ConfigureRouter(r)
	index(r)

	return r
}

func NewApi(resourceService service.ResourceService, recordService service.RecordService) Api {
	return &api{
		resourceService:    resourceService,
		swaggerApi:         NewSwaggerApi(resourceService, recordService),
		typescriptTypesApi: NewTypescriptTypesApi(resourceService, recordService),
	}
}
