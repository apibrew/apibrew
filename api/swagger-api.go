package api

import (
	"context"
	"data-handler/model"
	"data-handler/service"
	"data-handler/service/errors"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gorilla/mux"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/http-swagger"
	"net/http"
)

type SwaggerApi interface {
	ConfigureRouter(r *mux.Router)
}

type swaggerApi struct {
	resourceService service.ResourceService
}

func (s *swaggerApi) ConfigureRouter(r *mux.Router) {
	swaggerFiles.Handler.Prefix = "/docs/"

	r.HandleFunc("/docs/api.json", func(w http.ResponseWriter, req *http.Request) {
		doc, serviceErr := s.prepareDoc(req.Context())

		if serviceErr != nil {
			handleServiceError(w, serviceErr)
			return
		}

		data, err := doc.MarshalJSON()

		if err != nil {
			handleClientError(w, err)
			return
		}

		w.Write(data)
	})

	r.HandleFunc("/docs/resources/{namespace}/{resourceName}.json", func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		namespace := vars["namespace"]
		resourceName := vars["resourceName"]

		resource, serviceErr := s.resourceService.GetResourceByName(req.Context(), namespace, resourceName)

		if serviceErr != nil {
			handleServiceError(w, serviceErr)
			return
		}

		doc := s.prepareResourceSchema(resource)

		data, err := doc.MarshalJSON()

		if err != nil {
			handleClientError(w, err)
			return
		}

		w.Write(data)
	})

	r.HandleFunc("/docs", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Location", "/docs/index.html")
		writer.WriteHeader(301)
	})

	r.PathPrefix("/docs").HandlerFunc(httpSwagger.Handler(
		httpSwagger.URL("/docs/api.json"), //The url pointing to API definition
	))
}

func (s *swaggerApi) prepareDoc(ctx context.Context) (*openapi3.T, errors.ServiceError) {
	loader := openapi3.NewLoader()

	doc, err := loader.LoadFromFile("openapi.base.yml")

	if err != nil {
		panic(err)
	}

	list, serviceErr := s.resourceService.List(ctx)

	if serviceErr != nil {
		return nil, serviceErr
	}

	for _, item := range list {
		s.appendResourceApis(ctx, doc, item)
	}

	return doc, nil
}

func (s *swaggerApi) appendResourceApis(ctx context.Context, doc *openapi3.T, resource *model.Resource) {
	jsonSchemaRef := "/docs/resources/" + resource.Namespace + "/" + resource.Name + ".json"

	doc.Paths["/"+s.getResourceFQN(resource)] = &openapi3.PathItem{
		Summary:     resource.Name,
		Description: "Api for " + resource.Name,
		Get: &openapi3.Operation{
			Responses: map[string]*openapi3.ResponseRef{
				"200": {
					Value: &openapi3.Response{
						Content: openapi3.NewContentWithJSONSchemaRef(&openapi3.SchemaRef{
							Ref: "#/components/schemas/list-" + s.getResourceFQN(resource),
						}),
					},
				},
			},
		},
		Post: &openapi3.Operation{
			RequestBody: &openapi3.RequestBodyRef{
				Value: &openapi3.RequestBody{
					Required: true,
					Content: openapi3.NewContentWithJSONSchemaRef(&openapi3.SchemaRef{
						Ref: jsonSchemaRef,
					}),
				},
			},
			Responses: map[string]*openapi3.ResponseRef{
				"200": {
					Value: &openapi3.Response{
						Content: openapi3.NewContentWithJSONSchemaRef(&openapi3.SchemaRef{
							Ref: jsonSchemaRef,
						}),
					},
				},
			},
		},
	}

	doc.Components.Schemas["list-"+s.getResourceFQN(resource)] = &openapi3.SchemaRef{
		Value: &openapi3.Schema{
			Properties: map[string]*openapi3.SchemaRef{
				"content": {
					Value: &openapi3.Schema{
						Type: openapi3.TypeArray,
						Items: &openapi3.SchemaRef{
							Ref: jsonSchemaRef,
						},
					},
				},
			},
		},
	}
}

func (s *swaggerApi) getResourceFQN(resource *model.Resource) string {
	if resource.Namespace == "default" {
		return resource.Name
	} else {
		return resource.Namespace + "-" + resource.Name
	}
}

func (s *swaggerApi) prepareResourceSchema(resource *model.Resource) *openapi3.Schema {
	var requiredItems []string

	schema := &openapi3.Schema{
		Properties: map[string]*openapi3.SchemaRef{},
	}

	for _, property := range resource.Properties {
		schema.Properties[property.Name] = &openapi3.SchemaRef{
			Value: &openapi3.Schema{
				//Type: property.Type.String(),
				Type: "string",
			},
		}

		if property.Required {
			requiredItems = append(requiredItems, property.Name)
		}
	}

	schema.Required = requiredItems

	return schema
}

func NewSwaggerApi(resourceService service.ResourceService) SwaggerApi {
	return &swaggerApi{
		resourceService: resourceService,
	}
}
