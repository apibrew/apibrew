package docs

import (
	"github.com/apibrew/apibrew/pkg/generator/templates/typescript"
	"github.com/apibrew/apibrew/pkg/service"
	"github.com/apibrew/apibrew/pkg/util"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

type TypescriptTypesApi interface {
	ConfigureRouter(r *mux.Router)
}

type typescriptTypesApi struct {
	resourceService service.ResourceService
	recordService   service.RecordService
}

func (s *typescriptTypesApi) ConfigureRouter(r *mux.Router) {
	r.HandleFunc("/docs/typescript-types.d.ts", func(w http.ResponseWriter, req *http.Request) {
		resources, err := s.resourceService.List(util.SystemContext)

		var pkg = "model"

		if err != nil {
			http.Error(w, err.GetFullMessage(), 500)
			return
		}

		for _, resource := range resources {
			code := typescript.GenerateClassCode(pkg, resource)

			code = strings.Replace(code, "export ", "declare ", -1)

			_, err := w.Write([]byte(code))

			if err != nil {
				log.Error(err)
				return
			}
		}
	})
}

func NewTypescriptTypesApi(resourceService service.ResourceService, recordService service.RecordService) TypescriptTypesApi {
	return &typescriptTypesApi{
		resourceService: resourceService,
		recordService:   recordService,
	}
}
