package openapi

import (
	"context"
	"fmt"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resources"
	"github.com/apibrew/apibrew/pkg/service"
	"github.com/apibrew/apibrew/pkg/service/annotations"
	"github.com/apibrew/apibrew/pkg/util"
	"github.com/getkin/kin-openapi/openapi3"
	log "github.com/sirupsen/logrus"
	"strings"
)

type OpenApiBuilder struct {
	ResourceService service.ResourceService
	RecordService   service.RecordService
}

type OpenApiDocPrepareConfig struct {
	Group      string
	Namespaces []string
	Resources  []string
}

func (s *OpenApiBuilder) PrepareDoc(ctx context.Context, config OpenApiDocPrepareConfig) (*openapi3.T, errors.ServiceError) {
	var doc = new(openapi3.T)

	err := doc.UnmarshalJSON([]byte(openApiBaseContent))

	if err != nil {
		return nil, errors.InternalError.WithMessage(err.Error())
	}

	list, _ := s.ResourceService.List(util.WithSystemContext(ctx))

	for _, item := range list {
		if annotations.IsEnabled(item, annotations.RestApiDisabled) {
			continue
		}

		if !checkResourceAllowed(config, item, false) {
			continue
		}

		s.appendResourceApis(doc, item)
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
	}

	return doc, nil
}

func checkResourceAllowed(config OpenApiDocPrepareConfig, resource *model.Resource, forResource bool) bool {
	if len(config.Namespaces) > 0 {
		if !util.ArrayContains(config.Namespaces, resource.Namespace) {
			return false
		}
	}

	if len(config.Resources) > 0 {
		if !util.ArrayContains(config.Resources, resource.Name) {
			return false
		}
	}

	if config.Group == "full" {
		return true
	}

	resourceApiGroup := annotations.Get(resource, annotations.OpenApiGroup)

	if resourceApiGroup == "" {
		resourceApiGroup = "user"
	}

	if forResource && resource.Namespace == "system" && config.Group == "meta" {
		return true
	}

	return resourceApiGroup == config.Group
}

func (s *OpenApiBuilder) appendResourceApis(doc *openapi3.T, resource *model.Resource) {
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

	if !annotations.IsEnabled(resource, annotations.ActionApi) {
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
	}

	if resource == resources.ResourceResource {
		return
	}

	if !annotations.IsEnabled(resource, annotations.ActionApi) {
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
	}

	if annotations.IsEnabled(resource, annotations.ActionApi) {
		doc.Paths["/"+s.getResourceFQN(resource)] = &openapi3.PathItem{
			Summary:     title,
			Description: description,
			Post: &openapi3.Operation{
				Tags:        tags,
				Summary:     title,
				Description: description,
				OperationID: resource.GetName(),
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
							Description: util.Pointer("Result"),
							Content: openapi3.NewContentWithJSONSchemaRef(&openapi3.SchemaRef{
								Ref: jsonSchemaRef,
							}),
						},
					},
				},
			},
		}
	}
}

func (s *OpenApiBuilder) appendRecordeGenericApi(doc *openapi3.T) {
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
}

func (s *OpenApiBuilder) Init() {
	var doc = new(openapi3.T)

	err := doc.UnmarshalJSON([]byte(openApiBaseContent))

	log.Println(err)
}
