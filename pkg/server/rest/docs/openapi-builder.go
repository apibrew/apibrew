package docs

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resource_model"
	"github.com/apibrew/apibrew/pkg/resource_model/extramappings"
	"github.com/apibrew/apibrew/pkg/resources"
	"github.com/apibrew/apibrew/pkg/resources/mapping"
	"github.com/apibrew/apibrew/pkg/service"
	"github.com/apibrew/apibrew/pkg/service/annotations"
	"github.com/apibrew/apibrew/pkg/util"
	"github.com/getkin/kin-openapi/openapi3"
	"strings"
)

type openApiBuilder struct {
	resourceService service.ResourceService
	recordService   service.RecordService
}

type OpenApiDocPrepareConfig struct {
	group      string
	namespaces []string
	resources  []string
}

var error401 = &openapi3.ResponseRef{
	Value: &openapi3.Response{
		Description: util.Pointer("Unauthorized"),
		Content: openapi3.NewContentWithJSONSchemaRef(&openapi3.SchemaRef{
			Ref: "#/components/schemas/Error",
		}),
	},
}

var error400 = &openapi3.ResponseRef{
	Value: &openapi3.Response{
		Description: util.Pointer("Bad Request"),
		Content: openapi3.NewContentWithJSONSchemaRef(&openapi3.SchemaRef{
			Ref: "#/components/schemas/Error",
		}),
	},
}

var authenticationTokenApi = &openapi3.PathItem{
	Post: &openapi3.Operation{
		Tags:        []string{"Authentication"},
		OperationID: "getAuthenticationToken",
		Summary:     "This endpoint is used to authenticate the user and get the access token.",
		Description: "The access token is used to authenticate the user for all the endpoints which needs authentication.",
		RequestBody: &openapi3.RequestBodyRef{
			Value: &openapi3.RequestBody{
				Required: true,
				Content: openapi3.NewContentWithJSONSchemaRef(&openapi3.SchemaRef{
					Ref: "#/components/schemas/AuthenticationRequest",
				}),
			},
		},
		Responses: map[string]*openapi3.ResponseRef{
			"200": {
				Value: &openapi3.Response{
					Description: util.Pointer("Authentication Response"),
					Content: openapi3.NewContentWithJSONSchemaRef(&openapi3.SchemaRef{
						Ref: "#/components/schemas/AuthenticationResponse",
					}),
				},
			},
			"401": {
				Value: &openapi3.Response{
					Description: util.Pointer("Unauthorized"),
					Content: openapi3.Content{
						"application/json": &openapi3.MediaType{
							Example: fromJson(`{
						  "code": "AUTHENTICATION_FAILED",
						  "message": "Authentication failed"
						}`),
						},
					},
				},
			},
			"400": {
				Value: &openapi3.Response{
					Description: util.Pointer("Unauthorized"),
					Content: openapi3.Content{
						"application/json": &openapi3.MediaType{
							Example: fromJson(`{
							"code":"RECORD_VALIDATION_ERROR",
							"message":"json: cannot unmarshal number into Go struct field AuthenticationRequest.password of type string"
						}`),
						},
					},
				},
			},
		},
	},
	Put: &openapi3.Operation{
		Tags:        []string{"Authentication"},
		OperationID: "refreshAuthenticationToken",
		Summary:     "This endpoint is used to refresh the access token.",
		Description: "The access token is used to authenticate the user for all the endpoints which needs authentication.",
		RequestBody: &openapi3.RequestBodyRef{
			Value: &openapi3.RequestBody{
				Required: true,
				Content: openapi3.NewContentWithJSONSchemaRef(&openapi3.SchemaRef{
					Ref: "#/components/schemas/RefreshTokenRequest",
				}),
			},
		},
		Responses: map[string]*openapi3.ResponseRef{
			"200": {
				Value: &openapi3.Response{
					Description: util.Pointer("Authentication Response"),
					Content: openapi3.NewContentWithJSONSchemaRef(&openapi3.SchemaRef{
						Ref: "#/components/schemas/RefreshTokenResponse",
					}),
				},
			},
			"401": error401,
			"400": error400,
		},
	},
	Get: &openapi3.Operation{
		Tags:        []string{"Authentication"},
		OperationID: "ViewToken",
		Summary:     "This endpoint is used to view the content of the access token.",
		Description: "The access token is used to authenticate the user for all the endpoints which needs authentication.",
		Security: &openapi3.SecurityRequirements{
			{
				"bearerAuth": []string{},
			},
		},
		Responses: map[string]*openapi3.ResponseRef{
			"200": {
				Value: &openapi3.Response{
					Description: util.Pointer("Authentication Response"),
					Content: openapi3.NewContentWithJSONSchemaRef(&openapi3.SchemaRef{
						Ref: "#/components/schemas/ViewTokenResponse",
					}),
				},
			},
			"401": error401,
			"400": error400,
		},
	},
}

var resourcesApi = &openapi3.PathItem{
	Get: &openapi3.Operation{
		Tags:        []string{"Resources"},
		OperationID: "getResources",
		Summary:     "This endpoint is used to get the list of resources.",
		Description: "The access token is used to authenticate the user for all the endpoints which needs authentication.",
		Responses: map[string]*openapi3.ResponseRef{
			"200": {
				Value: &openapi3.Response{
					Description: util.Pointer("List of resources"),
					Content: openapi3.NewContentWithJSONSchemaRef(&openapi3.SchemaRef{
						Value: &openapi3.Schema{
							Type: openapi3.TypeArray,
							Items: &openapi3.SchemaRef{
								Ref: "#/components/schemas/SystemResource",
							},
						},
					}),
				},
			},
			"401": error401,
			"400": error400,
		},
	},
	Post: &openapi3.Operation{
		Tags:        []string{"Resources"},
		OperationID: "createResource",
		Summary:     "This endpoint is used to create a new resource.",
		Description: "The access token is used to authenticate the user for all the endpoints which needs authentication.",
		RequestBody: &openapi3.RequestBodyRef{
			Value: &openapi3.RequestBody{
				Required: true,
				Content: openapi3.NewContentWithJSONSchemaRef(&openapi3.SchemaRef{
					Ref: "#/components/schemas/SystemResource",
				}),
			},
		},
		Responses: map[string]*openapi3.ResponseRef{
			"200": {
				Value: &openapi3.Response{
					Description: util.Pointer("Created resource"),
					Content: openapi3.NewContentWithJSONSchemaRef(&openapi3.SchemaRef{
						Ref: "#/components/schemas/SystemResource",
					}),
				},
			},
			"401": error401,
			"400": error400,
		},
	},
}

var resourceItemApi = &openapi3.PathItem{
	Get: &openapi3.Operation{
		Tags:        []string{"Resources"},
		OperationID: "getResource",
		Summary:     "This endpoint is used to get a resource.",
		Description: "The access token is used to authenticate the user for all the endpoints which needs authentication.",
		Parameters: []*openapi3.ParameterRef{
			{
				Value: &openapi3.Parameter{
					Name:     "id",
					In:       "path",
					Required: true,
					Schema: &openapi3.SchemaRef{
						Value: &openapi3.Schema{
							Type: "string",
						},
					},
				},
			},
		},
		Responses: map[string]*openapi3.ResponseRef{
			"200": {
				Value: &openapi3.Response{
					Description: util.Pointer("Resource"),
					Content: openapi3.NewContentWithJSONSchemaRef(&openapi3.SchemaRef{
						Ref: "#/components/schemas/SystemResource",
					}),
				},
			},
			"401": error401,
			"400": error400,
		},
	},
	Put: &openapi3.Operation{
		Tags:        []string{"Resources"},
		OperationID: "updateResource",
		Summary:     "This endpoint is used to update a resource.",
		Description: "The access token is used to authenticate the user for all the endpoints which needs authentication.",
		Parameters: []*openapi3.ParameterRef{
			{
				Value: &openapi3.Parameter{
					Name:     "id",
					In:       "path",
					Required: true,
					Schema: &openapi3.SchemaRef{
						Value: &openapi3.Schema{
							Type: "string",
						},
					},
				},
			},
		},
		RequestBody: &openapi3.RequestBodyRef{
			Value: &openapi3.RequestBody{
				Required: true,
				Content: openapi3.NewContentWithJSONSchemaRef(&openapi3.SchemaRef{
					Ref: "#/components/schemas/SystemResource",
				}),
			},
		},
	},
	Delete: &openapi3.Operation{
		Tags:        []string{"Resources"},
		OperationID: "deleteResource",
		Summary:     "This endpoint is used to delete a resource.",
		Description: "The access token is used to authenticate the user for all the endpoints which needs authentication.",
		Parameters: []*openapi3.ParameterRef{
			{
				Value: &openapi3.Parameter{
					Name:     "id",
					In:       "path",
					Required: true,
					Schema: &openapi3.SchemaRef{
						Value: &openapi3.Schema{
							Type: "string",
						},
					},
				},
			},
		},
	},
}

var resourceItemByNameApi = &openapi3.PathItem{
	Get: &openapi3.Operation{
		Tags:        []string{"Resources"},
		OperationID: "getResourceByName",
		Summary:     "This endpoint is used to get a resource by name.",
		Description: "The access token is used to authenticate the user for all the endpoints which needs authentication.",
		Parameters: []*openapi3.ParameterRef{
			{
				Value: &openapi3.Parameter{
					Name:     "namespace",
					In:       "path",
					Required: true,
					Schema: &openapi3.SchemaRef{
						Value: &openapi3.Schema{
							Type: "string",
						},
					},
				},
			},
			{
				Value: &openapi3.Parameter{
					Name:     "name",
					In:       "path",
					Required: true,
					Schema: &openapi3.SchemaRef{
						Value: &openapi3.Schema{
							Type: "string",
						},
					},
				},
			},
		},
		Responses: map[string]*openapi3.ResponseRef{
			"200": {
				Value: &openapi3.Response{
					Description: util.Pointer("Resource"),
					Content: openapi3.NewContentWithJSONSchemaRef(&openapi3.SchemaRef{
						Ref: "#/components/schemas/SystemResource",
					}),
				},
			},
		},
	},
}

func (s *openApiBuilder) prepareDoc(ctx context.Context, config OpenApiDocPrepareConfig) (*openapi3.T, errors.ServiceError) {
	doc := &openapi3.T{
		OpenAPI: "3.0.0",
		Info: &openapi3.Info{
			Title:       "API Brew",
			Description: "OpenApi 3.0 definition for API Brew Rest API",
			Version:     "1.0",
			License: &openapi3.License{
				Name: "Apache License",
				URL:  "https://github.com/google/gnostic/blob/master/LICENSE",
			},
			Contact: &openapi3.Contact{
				Name: "API Brew",
			},
		},
		Paths: openapi3.Paths{
			"/authentication/token": authenticationTokenApi,
		},
		Components: openapi3.Components{
			Schemas: openapi3.Schemas{
				"Error": {
					Value: &openapi3.Schema{
						Required: []string{"code", "message"},
						Properties: map[string]*openapi3.SchemaRef{
							"code": {
								Value: &openapi3.Schema{
									Type: "string",
								},
							},
							"message": {
								Value: &openapi3.Schema{
									Type: "string",
								},
							},
							"details": {
								Value: &openapi3.Schema{
									Type: "object",
								},
							},
						},
					},
				},
				"AuthenticationRequest": {
					Value: &openapi3.Schema{
						Required: []string{"username", "password"},
						Properties: map[string]*openapi3.SchemaRef{
							"username": {
								Value: &openapi3.Schema{
									Type:    "string",
									Example: "admin",
								},
							},
							"password": {
								Value: &openapi3.Schema{
									Type:    "string",
									Example: "admin",
								},
							},
							"term": {
								Ref: "#/components/schemas/TokenTerm",
							},
						},
					},
				},
				"AuthenticationResponse": {
					Value: &openapi3.Schema{
						Required: []string{"token"},
						Properties: map[string]*openapi3.SchemaRef{
							"token": {
								Ref: "#/components/schemas/Token",
							},
						},
					},
				},
				"RefreshTokenRequest": {
					Value: &openapi3.Schema{
						Required: []string{"username", "password"},
						Properties: map[string]*openapi3.SchemaRef{
							"token": {
								Value: &openapi3.Schema{
									Type:    "string",
									Example: "admin",
								},
							},
							"term": {
								Ref: "#/components/schemas/TokenTerm",
							},
						},
					},
				},
				"RefreshTokenResponse": {
					Value: &openapi3.Schema{
						Required: []string{"token"},
						Properties: map[string]*openapi3.SchemaRef{
							"token": {
								Ref: "#/components/schemas/Token",
							},
						},
					},
				},
				"ViewTokenResponse": {
					Value: &openapi3.Schema{},
				},
				"Token": {
					Value: &openapi3.Schema{
						Type: "object",
						Properties: map[string]*openapi3.SchemaRef{
							"term": {
								Ref: "#/components/schemas/TokenTerm",
							},
							"content": {
								Value: &openapi3.Schema{
									Type: "string",
								},
							},
							"expiration": {
								Value: &openapi3.Schema{
									Type:   "string",
									Format: "date-time",
								},
							},
						},
						Example: fromJson(`{
							"term": "SHORT",
							"content": "eyJhbGciOiJSUzI1NiIsI-8pl63y8h-ZOLM",
							"expiration": "2023-10-16T15:14:11.161530Z"
						  }`),
					},
				},
				"TokenTerm": {
					Value: &openapi3.Schema{
						Type:   "string",
						Format: "enum",
						Enum: []interface{}{
							"VERY_SHORT", "SHORT", "MIDDLE", "LONG", "VERY_LONG",
						},
						Example: "LONG",
					},
				},
				"BooleanExpression": {
					Value: &openapi3.Schema{
						Type: "object",
					},
				},
				"SearchRecordRequest": {
					Value: &openapi3.Schema{
						Type: "object",
						Properties: map[string]*openapi3.SchemaRef{
							"query": {
								Ref: "#/components/schemas/BooleanExpression",
							},
							"limit": {
								Value: &openapi3.Schema{
									Type: "integer",
								},
							},
							"offset": {
								Value: &openapi3.Schema{
									Type: "integer",
								},
							},
							"useHistory": {
								Value: &openapi3.Schema{
									Type: "boolean",
								},
							},
							"resolveReferences": {
								Value: &openapi3.Schema{
									Type: "array",
									Items: &openapi3.SchemaRef{
										Value: &openapi3.Schema{
											Type: "string",
										},
									},
								},
							},
							"annotations": {
								Value: &openapi3.Schema{
									Type: "object",
									AdditionalProperties: &openapi3.SchemaRef{
										Value: &openapi3.Schema{
											Type: "string",
										},
									},
								},
							},
						},
					},
				},
			},
			SecuritySchemes: map[string]*openapi3.SecuritySchemeRef{
				"bearerAuth": {
					Value: &openapi3.SecurityScheme{
						Type:         "http",
						Scheme:       "bearer",
						BearerFormat: "JWT",
					},
				},
			},
		},
		Security: []openapi3.SecurityRequirement{
			{
				"bearerAuth": []string{},
			},
		},
		Tags: []*openapi3.Tag{
			{
				Name:        "Authentication",
				Description: "Authentication APIs are used to authenticate users and get access to the resources. For all endpoints, which needs you to be authenticated, you need to pass the access token in the header.\n                The access token is obtained by calling the authenticate endpoint.",
			},
		},
	}

	if checkResourceAllowed(config, resources.ResourceResource, false) {
		doc.Paths["/resources"] = resourcesApi
		doc.Paths["/resources/{id}"] = resourceItemApi
		doc.Paths["/resources/by-name/{namespace}/{name}"] = resourceItemByNameApi
	}

	list, _ := s.resourceService.List(util.WithSystemContext(ctx))

	var resourceActionsRecords, _, err = s.recordService.List(util.WithSystemContext(context.TODO()), service.RecordListParams{
		Namespace: resources.ResourceActionResource.Namespace,
		Resource:  resources.ResourceActionResource.Name,
		Limit:     1000000,
	})

	if err != nil {
		return nil, err
	}

	var resourceActions = util.ArrayMap(resourceActionsRecords, resource_model.ResourceActionMapperInstance.FromRecord)

	for _, item := range list {
		if annotations.IsEnabled(item, annotations.RestApiDisabled) {
			continue
		}

		if !checkResourceAllowed(config, item, false) {
			continue
		}

		var resourceActionsForItem []*resource_model.ResourceAction

		for _, resourceAction := range resourceActions {
			if resourceAction.Resource.Id.String() == item.Id {
				resourceAction.Resource = extramappings.ResourceTo(item)
				resourceActionsForItem = append(resourceActionsForItem, resourceAction)

			}
		}

		s.appendResourceApis(doc, item, resourceActionsForItem)
	}

	s.appendRecordeGenericApi(doc)

	// post processing
	var security = &openapi3.SecurityRequirements{
		{
			"bearerAuth": []string{},
		},
	}

	for pathKey, path := range doc.Paths {
		for operationKey, operation := range path.Operations() {
			if pathKey == "/authentication/token" && operationKey == "POST" {
				continue
			}

			operation.Security = security
		}

		if strings.HasPrefix(pathKey, "/records/") {
			delete(doc.Paths, pathKey)
		}
	}

	for _, item := range list {
		if !checkResourceAllowed(config, item, true) {
			continue
		}

		if annotations.IsEnabled(item, annotations.OpenApiHide) {
			continue
		}

		schemaName := util.ResourceJsonSchemaName(item)

		doc.Components.Schemas[schemaName] = &openapi3.SchemaRef{
			Value: util.PropertiesWithTitleToJsonSchema(item, item),
		}

		for _, subType := range item.Types {
			doc.Components.Schemas[schemaName+subType.Name] = &openapi3.SchemaRef{
				Value: util.PropertiesWithTitleToJsonSchema(item, subType),
			}
		}

		for _, resourceAction := range resourceActions {
			if resourceAction.Resource.Id.String() == item.Id && resourceAction.Input != nil {
				var resourceActionAsResource = &resource_model.Resource{
					Name:       item.Name + "_" + resourceAction.Name,
					Properties: resourceAction.Input,
					Types:      resourceAction.Types,
				}

				resourceActionAsResourceInt := mapping.ResourceFromRecord(resource_model.ResourceMapperInstance.ToRecord(resourceActionAsResource))

				doc.Components.Schemas[schemaName+"_"+resourceAction.Name] = &openapi3.SchemaRef{
					Value: util.PropertiesWithTitleToJsonSchema(resourceActionAsResourceInt, resourceActionAsResourceInt),
				}
			}
		}
	}

	return doc, nil
}

func fromJson(s string) interface{} {
	var result interface{}
	err := json.Unmarshal([]byte(s), &result)

	if err != nil {
		panic(err)
	}

	return result
}

func checkResourceAllowed(config OpenApiDocPrepareConfig, resource *model.Resource, forResource bool) bool {
	if len(config.namespaces) > 0 {
		if !util.ArrayContains(config.namespaces, resource.Namespace) {
			return false
		}
	}

	if len(config.resources) > 0 {
		if !util.ArrayContains(config.resources, resource.Name) {
			return false
		}
	}

	if config.group == "full" {
		return true
	}

	resourceApiGroup := annotations.Get(resource, annotations.OpenApiGroup)

	if resourceApiGroup == "" {
		resourceApiGroup = "user"
	}

	if forResource && resource.Namespace == "system" && config.group == "meta" {
		return true
	}

	return resourceApiGroup == config.group
}

func (s *openApiBuilder) appendResourceApis(doc *openapi3.T, resource *model.Resource, resourceActions []*resource_model.ResourceAction) {
	jsonSchemaRef := "#/components/schemas/" + util.ResourceJsonSchemaName(resource)

	var tags []string
	if resource.GetNamespace() == "default" {
		tags = []string{resource.GetName() + " API"}
	} else {
		tags = []string{resource.GetNamespace() + " / " + resource.GetName() + " API"}
	}

	title := resource.GetTitle()
	description := resource.GetDescription()

	if title == "" {
		title = resource.GetName()
	}

	if description == "" {
		description = "Api for " + resource.GetName()
	}

	var itemSchema = &openapi3.Schema{
		Properties: map[string]*openapi3.SchemaRef{
			"content": {
				Ref: jsonSchemaRef,
			},
		},
	}

	doc.Paths["/"+s.getResourceFQN(resource)] = &openapi3.PathItem{
		Summary:     title,
		Description: description,
		Get: &openapi3.Operation{
			Tags:        tags,
			Summary:     fmt.Sprintf("%s - List items", title),
			Description: fmt.Sprintf("%s - List items", description),
			OperationID: "list" + resource.GetName(),
			Responses: map[string]*openapi3.ResponseRef{
				"200": {
					Value: &openapi3.Response{
						Description: util.Pointer("List of items"),
						Content: openapi3.NewContentWithJSONSchemaRef(&openapi3.SchemaRef{
							Value: &openapi3.Schema{
								Required: []string{"content"},
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
						}),
					},
				},
			},
		},
		Post: &openapi3.Operation{
			Tags:        tags,
			Summary:     fmt.Sprintf("%s - Create new item", title),
			Description: fmt.Sprintf("%s - Create new item", description),
			OperationID: "create" + resource.GetName(),
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
						Description: util.Pointer("Created item"),
						Content: openapi3.NewContentWithJSONSchemaRef(&openapi3.SchemaRef{
							Ref: jsonSchemaRef,
						}),
					},
				},
			},
		},
		Patch: &openapi3.Operation{
			Tags:        tags,
			Summary:     fmt.Sprintf("%s - Apply an item", title),
			Description: fmt.Sprintf("%s - Apply an item, it will check id and unique properties, if such item is exists, update operation will be executed, if not create operation is executed. If There are no change between updating record and existing record, nothing will be done", description),
			OperationID: "apply" + resource.GetName(),
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
						Description: util.Pointer("Updated item"),
						Content: openapi3.NewContentWithJSONSchemaRef(&openapi3.SchemaRef{
							Ref: jsonSchemaRef,
						}),
					},
				},
			},
		},
	}

	doc.Paths["/"+s.getResourceFQN(resource)+"/{id}"] = &openapi3.PathItem{
		Summary:     title,
		Description: description,
		Get: &openapi3.Operation{
			Tags:        tags,
			Summary:     fmt.Sprintf("%s - Get_ item", title),
			Description: fmt.Sprintf("%s - Get_ item", description),
			OperationID: "get" + resource.GetName(),
			Parameters: []*openapi3.ParameterRef{
				{
					Value: &openapi3.Parameter{
						Name:     "id",
						In:       "path",
						Required: true,
						Schema: &openapi3.SchemaRef{
							Value: &openapi3.Schema{
								Type: "string",
							},
						},
					},
				},
			},
			Responses: map[string]*openapi3.ResponseRef{
				"200": {
					Value: &openapi3.Response{
						Description: util.Pointer("Item"),
						Content: openapi3.NewContentWithJSONSchemaRef(&openapi3.SchemaRef{
							Value: itemSchema,
						}),
					},
				},
			},
		},
		Delete: &openapi3.Operation{
			Tags:        tags,
			Summary:     fmt.Sprintf("%s - Delete_ item", title),
			Description: fmt.Sprintf("%s - Delete_ item", description),
			OperationID: "delete" + resource.GetName(),
			Parameters: []*openapi3.ParameterRef{
				{
					Value: &openapi3.Parameter{
						Name:     "id",
						In:       "path",
						Required: true,
						Schema: &openapi3.SchemaRef{
							Value: &openapi3.Schema{
								Type: "string",
							},
						},
					},
				},
			},
			Responses: map[string]*openapi3.ResponseRef{
				"200": {
					Value: &openapi3.Response{
						Description: util.Pointer("Deleted item"),
						Content: openapi3.NewContentWithJSONSchemaRef(&openapi3.SchemaRef{
							Value: itemSchema,
						}),
					},
				},
			},
		},
		Put: &openapi3.Operation{
			Tags:        tags,
			Summary:     fmt.Sprintf("%s - Update item", title),
			Description: fmt.Sprintf("%s - Update item", description),
			OperationID: "update" + resource.GetName(),
			Parameters: []*openapi3.ParameterRef{
				{
					Value: &openapi3.Parameter{
						Name:     "id",
						In:       "path",
						Required: true,
						Schema: &openapi3.SchemaRef{
							Value: &openapi3.Schema{
								Type: "string",
							},
						},
					},
				},
			},
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
						Description: util.Pointer("Updated item"),
						Content: openapi3.NewContentWithJSONSchemaRef(&openapi3.SchemaRef{
							Ref: jsonSchemaRef,
						}),
					},
				},
			},
		},
	}

	if resource == resources.ResourceResource {
		return
	}

	doc.Paths["/"+s.getResourceFQN(resource)+"/_search"] = &openapi3.PathItem{
		Summary:     title + " - Search",
		Description: description + " - Search",
		Post: &openapi3.Operation{
			Tags:        tags,
			Summary:     fmt.Sprintf("%s - Search items", title),
			Description: fmt.Sprintf("%s - Search items", description),
			OperationID: "search" + resource.GetName(),
			RequestBody: &openapi3.RequestBodyRef{
				Value: &openapi3.RequestBody{
					Required: true,
					Content: openapi3.NewContentWithJSONSchemaRef(&openapi3.SchemaRef{
						Ref: "#/components/schemas/SearchRecordRequest",
					}),
				},
			},
			Responses: map[string]*openapi3.ResponseRef{
				"200": {
					Value: &openapi3.Response{
						Description: util.Pointer("List of items"),
						Content: openapi3.NewContentWithJSONSchemaRef(&openapi3.SchemaRef{
							Value: itemSchema,
						}),
					},
				},
			},
		},
	}

	for _, resourceAction := range resourceActions {
		doc.Paths["/"+s.getResourceFQN(resource)+"/{id}/_"+resourceAction.Name] = &openapi3.PathItem{
			Summary:     title + " - " + util.DePointer(resourceAction.Title, resourceAction.Name),
			Description: util.DePointer(resourceAction.Description, ""),
			Post:        s.resourceActionOperation(tags, title, resourceAction, jsonSchemaRef, itemSchema),
		}
	}
}

func (s *openApiBuilder) appendRecordeGenericApi(doc *openapi3.T) {
	var resource = resources.RecordResource

	var genericRecordSchema = &openapi3.SchemaRef{
		Value: &openapi3.Schema{
			Title:                       "Generic Record",
			Description:                 "It is a generic payload, it has not a specific structure without knowledge of resource. For each resource this structure is shaped accordingly",
			AdditionalPropertiesAllowed: util.Pointer(true),
		},
	}

	var tags = []string{"Records Generic API"}

	title := resource.GetTitle()
	description := resource.GetDescription()

	if title == "" {
		title = resource.GetName()
	}

	if description == "" {
		description = "Api for " + resource.GetName()
	}

	var itemSchema = &openapi3.Schema{
		Properties: map[string]*openapi3.SchemaRef{
			"content": genericRecordSchema,
		},
	}

	var genericParameters = []*openapi3.ParameterRef{
		{
			Value: &openapi3.Parameter{
				Name:     "namespace",
				In:       "path",
				Required: true,
				Schema: &openapi3.SchemaRef{
					Value: &openapi3.Schema{
						Type: "string",
					},
				},
			},
		},
		{
			Value: &openapi3.Parameter{
				Name:     "resource",
				In:       "path",
				Required: true,
				Schema: &openapi3.SchemaRef{
					Value: &openapi3.Schema{
						Type: "string",
					},
				},
			},
		},
	}

	doc.Paths["/{namespace}-{resource}"] = &openapi3.PathItem{
		Summary:     title,
		Description: description,
		Parameters:  genericParameters,
		Get: &openapi3.Operation{
			Tags:        tags,
			Summary:     fmt.Sprintf("%s - List items", title),
			Description: fmt.Sprintf("%s - List items", description),
			OperationID: "list" + resource.GetName(),
			Responses: map[string]*openapi3.ResponseRef{
				"200": {
					Value: &openapi3.Response{
						Description: util.Pointer("List of items"),
						Content: openapi3.NewContentWithJSONSchemaRef(&openapi3.SchemaRef{
							Value: &openapi3.Schema{
								Required: []string{"content"},
								Properties: map[string]*openapi3.SchemaRef{
									"content": {
										Value: &openapi3.Schema{
											Type:  openapi3.TypeArray,
											Items: genericRecordSchema,
										},
									},
								},
							},
						}),
					},
				},
			},
		},
		Post: &openapi3.Operation{
			Tags:        tags,
			Summary:     fmt.Sprintf("%s - Create new item", title),
			Description: fmt.Sprintf("%s - Create new item", description),
			OperationID: "create" + resource.GetName(),
			RequestBody: &openapi3.RequestBodyRef{
				Value: &openapi3.RequestBody{
					Required: true,
					Content:  openapi3.NewContentWithJSONSchemaRef(genericRecordSchema),
				},
			},
			Responses: map[string]*openapi3.ResponseRef{
				"200": {
					Value: &openapi3.Response{
						Description: util.Pointer("Created item"),
						Content:     openapi3.NewContentWithJSONSchemaRef(genericRecordSchema),
					},
				},
			},
		},
		Patch: &openapi3.Operation{
			Tags:        tags,
			Summary:     fmt.Sprintf("%s - Apply an item", title),
			Description: fmt.Sprintf("%s - Apply an item, it will check id and unique properties, if such item is exists, update operation will be executed, if not create operation is executed. If There are no change between updating record and existing record, nothing will be done", description),
			OperationID: "apply" + resource.GetName(),
			RequestBody: &openapi3.RequestBodyRef{
				Value: &openapi3.RequestBody{
					Required: true,
					Content:  openapi3.NewContentWithJSONSchemaRef(genericRecordSchema),
				},
			},
			Responses: map[string]*openapi3.ResponseRef{
				"200": {
					Value: &openapi3.Response{
						Description: util.Pointer("Updated item"),
						Content:     openapi3.NewContentWithJSONSchemaRef(genericRecordSchema),
					},
				},
			},
		},
	}

	doc.Paths["/{namespace}-{resource}/{id}"] = &openapi3.PathItem{
		Summary:     title,
		Description: description,
		Parameters:  genericParameters,
		Get: &openapi3.Operation{
			Tags:        tags,
			Summary:     fmt.Sprintf("%s - Get_ item", title),
			Description: fmt.Sprintf("%s - Get_ item", description),
			OperationID: "get" + resource.GetName(),
			Parameters: []*openapi3.ParameterRef{
				{
					Value: &openapi3.Parameter{
						Name:     "id",
						In:       "path",
						Required: true,
						Schema: &openapi3.SchemaRef{
							Value: &openapi3.Schema{
								Type: "string",
							},
						},
					},
				},
			},
			Responses: map[string]*openapi3.ResponseRef{
				"200": {
					Value: &openapi3.Response{
						Description: util.Pointer("Item"),
						Content: openapi3.NewContentWithJSONSchemaRef(&openapi3.SchemaRef{
							Value: itemSchema,
						}),
					},
				},
			},
		},
		Delete: &openapi3.Operation{
			Tags:        tags,
			Summary:     fmt.Sprintf("%s - Delete_ item", title),
			Description: fmt.Sprintf("%s - Delete_ item", description),
			OperationID: "delete" + resource.GetName(),
			Parameters: []*openapi3.ParameterRef{
				{
					Value: &openapi3.Parameter{
						Name:     "id",
						In:       "path",
						Required: true,
						Schema: &openapi3.SchemaRef{
							Value: &openapi3.Schema{
								Type: "string",
							},
						},
					},
				},
			},
			Responses: map[string]*openapi3.ResponseRef{
				"200": {
					Value: &openapi3.Response{
						Description: util.Pointer("Deleted item"),
						Content: openapi3.NewContentWithJSONSchemaRef(&openapi3.SchemaRef{
							Value: itemSchema,
						}),
					},
				},
			},
		},
		Put: &openapi3.Operation{
			Tags:        tags,
			Summary:     fmt.Sprintf("%s - Update item", title),
			Description: fmt.Sprintf("%s - Update item", description),
			OperationID: "update" + resource.GetName(),
			Parameters: []*openapi3.ParameterRef{
				{
					Value: &openapi3.Parameter{
						Name:     "id",
						In:       "path",
						Required: true,
						Schema: &openapi3.SchemaRef{
							Value: &openapi3.Schema{
								Type: "string",
							},
						},
					},
				},
			},
			RequestBody: &openapi3.RequestBodyRef{
				Value: &openapi3.RequestBody{
					Required: true,
					Content:  openapi3.NewContentWithJSONSchemaRef(genericRecordSchema),
				},
			},
			Responses: map[string]*openapi3.ResponseRef{
				"200": {
					Value: &openapi3.Response{
						Description: util.Pointer("Updated item"),
						Content:     openapi3.NewContentWithJSONSchemaRef(genericRecordSchema),
					},
				},
			},
		},
	}

	if resource == resources.ResourceResource {
		return
	}

	doc.Paths["/{namespace}-{resource}/{id}/_search"] = &openapi3.PathItem{
		Summary:     title + " - Search",
		Description: description + " - Search",
		Parameters:  genericParameters,
		Post: &openapi3.Operation{
			Tags:        tags,
			Summary:     fmt.Sprintf("%s - Search items", title),
			Description: fmt.Sprintf("%s - Search items", description),
			OperationID: "search" + resource.GetName(),
			RequestBody: &openapi3.RequestBodyRef{
				Value: &openapi3.RequestBody{
					Required: true,
					Content: openapi3.NewContentWithJSONSchemaRef(&openapi3.SchemaRef{
						Ref: "#/components/schemas/SearchRecordRequest",
					}),
				},
			},
			Responses: map[string]*openapi3.ResponseRef{
				"200": {
					Value: &openapi3.Response{
						Description: util.Pointer("List of items"),
						Content: openapi3.NewContentWithJSONSchemaRef(&openapi3.SchemaRef{
							Value: itemSchema,
						}),
					},
				},
			},
		},
	}

	doc.Paths["/{namespace}-{resource}/{id}/{action}"] = &openapi3.PathItem{
		Summary:    title + " - " + "Action",
		Parameters: genericParameters,
		Post: &openapi3.Operation{
			Tags: tags,
			Parameters: []*openapi3.ParameterRef{
				{
					Value: &openapi3.Parameter{
						Name:     "id",
						In:       "path",
						Required: true,
						Schema: &openapi3.SchemaRef{
							Value: &openapi3.Schema{
								Type: "string",
							},
						},
					},
				},
				{
					Value: &openapi3.Parameter{
						Name:     "action",
						In:       "path",
						Required: true,
						Schema: &openapi3.SchemaRef{
							Value: &openapi3.Schema{
								Type: "string",
							},
						},
					},
				},
			},
			OperationID: "action" + resource.GetName(),
			Responses: map[string]*openapi3.ResponseRef{
				"200": {
					Value: &openapi3.Response{
						Description: util.Pointer("List of items"),
						Content: openapi3.NewContentWithJSONSchemaRef(&openapi3.SchemaRef{
							Value: itemSchema,
						}),
					},
				},
			},
		},
	}
}

func (s *openApiBuilder) resourceActionOperation(tags []string, title string, resourceAction *resource_model.ResourceAction, jsonSchemaRef string, itemSchema *openapi3.Schema) *openapi3.Operation {
	var result = &openapi3.Operation{
		Tags:        tags,
		Summary:     title + " - " + util.DePointer(resourceAction.Title, resourceAction.Name),
		Description: util.DePointer(resourceAction.Description, ""),
		OperationID: "action" + resourceAction.Resource.Name + resourceAction.Name,
		Parameters: []*openapi3.ParameterRef{
			{
				Value: &openapi3.Parameter{
					Name:     "id",
					In:       "path",
					Required: true,
					Schema: &openapi3.SchemaRef{
						Value: &openapi3.Schema{
							Type: "string",
						},
					},
				},
			},
		},
		Responses: map[string]*openapi3.ResponseRef{
			"200": {
				Value: &openapi3.Response{
					Description: util.Pointer("List of items"),
					Content: openapi3.NewContentWithJSONSchemaRef(&openapi3.SchemaRef{
						Value: itemSchema,
					}),
				},
			},
		},
	}

	if resourceAction.Input != nil {
		result.RequestBody = &openapi3.RequestBodyRef{
			Value: &openapi3.RequestBody{
				Required: true,
				Content: openapi3.NewContentWithJSONSchemaRef(&openapi3.SchemaRef{
					Ref: jsonSchemaRef + "_" + resourceAction.Name,
				}),
			},
		}
	}

	return result
}
