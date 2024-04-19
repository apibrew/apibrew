package docs

import (
	"github.com/apibrew/apibrew/pkg/docs/openapi"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/service"
	"github.com/apibrew/apibrew/pkg/util"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

type SwaggerApi interface {
	ConfigureRouter(r *mux.Router)
}

type swaggerApi struct {
	resourceService service.ResourceService
	recordService   service.RecordService
}

func (s *swaggerApi) ConfigureRouter(r *mux.Router) {
	var oab = &openapi.OpenApiBuilder{
		ResourceService: s.resourceService,
		RecordService:   s.recordService,
	}

	oab.Init()

	r.HandleFunc("/docs/openapi.json", func(w http.ResponseWriter, req *http.Request) {
		config := openapi.OpenApiDocPrepareConfig{Group: "user"}

		if req.URL.Query().Get("group") != "" {
			config.Group = req.URL.Query().Get("group")
		}

		if req.URL.Query().Get("namespace") != "" {
			config.Namespaces = strings.Split(req.URL.Query().Get("namespace"), ",")
		}

		if req.URL.Query().Get("resource") != "" {
			config.Resources = strings.Split(req.URL.Query().Get("resource"), ",")
		}

		doc, serviceErr := oab.PrepareDoc(req.Context(), config)

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
func NewSwaggerApi(resourceService service.ResourceService, recordService service.RecordService) SwaggerApi {
	return &swaggerApi{
		resourceService: resourceService,
		recordService:   recordService,
	}
}
