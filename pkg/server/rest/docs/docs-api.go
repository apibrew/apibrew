package docs

import (
	"github.com/apibrew/apibrew/pkg/abs"
	"github.com/gorilla/mux"
)

type api struct {
	resourceService abs.ResourceService
	swaggerApi      SwaggerApi
}

type Api interface {
	Handler() *mux.Router
}

func (a *api) Handler() *mux.Router {
	var r = mux.NewRouter()

	a.swaggerApi.ConfigureRouter(r)
	index(r)
	yaml(a.resourceService)(r)

	return r
}

func NewApi(resourceService abs.ResourceService) Api {
	return &api{
		resourceService: resourceService,
		swaggerApi:      NewSwaggerApi(resourceService),
	}
}
