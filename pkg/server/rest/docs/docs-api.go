package docs

import (
	"github.com/apibrew/apibrew/pkg/service"
	"github.com/gorilla/mux"
)

type api struct {
	resourceService service.ResourceService
	swaggerApi      SwaggerApi
}

type Api interface {
	Handler() *mux.Router
}

func (a *api) Handler() *mux.Router {
	var r = mux.NewRouter()

	a.swaggerApi.ConfigureRouter(r)
	index(r)

	return r
}

func NewApi(resourceService service.ResourceService) Api {
	return &api{
		resourceService: resourceService,
		swaggerApi:      NewSwaggerApi(resourceService),
	}
}
