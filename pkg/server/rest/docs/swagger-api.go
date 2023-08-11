package docs

import (
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/service"
	"github.com/apibrew/apibrew/pkg/util"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type SwaggerApi interface {
	ConfigureRouter(r *mux.Router)
}

type swaggerApi struct {
	resourceService service.ResourceService
}

func (s *swaggerApi) ConfigureRouter(r *mux.Router) {
	var oab = &openApiBuilder{
		resourceService: s.resourceService,
	}

	r.HandleFunc("/docs/openapi.json", func(w http.ResponseWriter, req *http.Request) {
		doc, serviceErr := oab.prepareDoc(req.Context(), OpenApiDocPrepareConfig{group: "user"})

		s.writeDocResult(w, serviceErr, doc)
	})

	r.HandleFunc("/docs/openapi/{group}.json", func(w http.ResponseWriter, req *http.Request) {
		var group = mux.Vars(req)["group"]

		doc, serviceErr := oab.prepareDoc(req.Context(), OpenApiDocPrepareConfig{group: group})

		s.writeDocResult(w, serviceErr, doc)
	})

	r.HandleFunc("/docs/resources/jsonschema/{namespace}/{resourceName}.json", func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		namespace := vars["namespace"]
		resourceName := vars["resourceName"]

		resource, _ := s.resourceService.GetResourceByName(util.WithSystemContext(req.Context()), namespace, resourceName)

		if resource == nil {
			http.Error(w, errors.ResourceNotFoundError.GetFullMessage(), 404)
			return
		}

		doc := util.PropertiesWithTitleToJsonSchema(resource, resource)

		data, err := doc.MarshalJSON()

		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		_, err = w.Write(data)

		if err != nil {
			log.Error(err)
		}

	})
}

func (s *swaggerApi) writeDocResult(w http.ResponseWriter, serviceErr errors.ServiceError, doc *openapi3.T) {
	if serviceErr != nil {
		http.Error(w, serviceErr.GetFullMessage(), 500)
		return
	}

	data, err := doc.MarshalJSON()

	if err != nil {
		http.Error(w, serviceErr.GetFullMessage(), 400)
		return
	}

	_, err = w.Write(data)

	if err != nil {
		log.Error(err)
	}
}
func NewSwaggerApi(resourceService service.ResourceService) SwaggerApi {
	return &swaggerApi{
		resourceService: resourceService,
	}
}
