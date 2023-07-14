package rest

import (
	"fmt"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/service"
	"net/http"
	"strings"
)

type recordsApiFiltersMiddleWare struct {
	resourceService service.ResourceService
}

func (w recordsApiFiltersMiddleWare) handler(handler http.Handler) http.Handler {
	resourceService := w.resourceService
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		path := req.URL.Path

		pathParts := strings.Split(path, "/")
		var newPathParts []string

		var resource *model.Resource

		if len(pathParts) == 2 {
			resource = resourceService.GetResourceByName(req.Context(), "default", pathParts[1])
		} else if len(pathParts) > 2 {
			resource = resourceService.GetResourceByName(req.Context(), pathParts[1], pathParts[2])
			newPathParts = pathParts[3:]

			if resource == nil {
				resource = resourceService.GetResourceByName(req.Context(), "default", pathParts[1])
				newPathParts = pathParts[2:]
			}
		}

		if resource != nil {
			newPath := fmt.Sprintf("/records/%s/%s", resource.Namespace, resource.Name)
			if len(newPathParts) > 0 {
				newPath = newPath + "/" + strings.Join(newPathParts, "/")
			}
			req.URL.RawPath = newPath
			req.URL.Path = newPath
		}

		handler.ServeHTTP(w, req)
	})
}

func newRecordsApiFiltersMiddleWare(resourceService service.ResourceService) *recordsApiFiltersMiddleWare {
	return &recordsApiFiltersMiddleWare{resourceService: resourceService}
}
