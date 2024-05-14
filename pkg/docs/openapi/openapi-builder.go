package openapi

import (
	"context"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resources"
	"github.com/apibrew/apibrew/pkg/service"
	annotations "github.com/apibrew/apibrew/pkg/service/annotations"
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

func (s *OpenApiBuilder) PrepareDoc(ctx context.Context, config OpenApiDocPrepareConfig) (*openapi3.T, error) {
	var doc = s.GetBaseDocCopy()

	list, _ := s.ResourceService.List(util.WithSystemContext(ctx))

	for _, item := range list {
		if annotations.IsEnabled(item, annotations.RestApiDisabled) {
			continue
		}

		if !checkResourceAllowed(config, item) {
			continue
		}

		s.appendResourceApis(doc, item)
	}

	for path := range doc.Paths {
		if strings.HasPrefix(path, "/{resourceRestPath}") {
			delete(doc.Paths, path)
		}
	}

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
		if !checkResourceAllowed(config, item) {
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

func checkResourceAllowed(config OpenApiDocPrepareConfig, resource *model.Resource) bool {
	if resource.Namespace == resources.ResourceResource.Namespace && resource.Name == resources.ResourceResource.Name {
		return true
	}

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

	return resourceApiGroup == config.Group
}

func (s *OpenApiBuilder) appendResourceApis(doc *openapi3.T, resource *model.Resource) {
	if resource == resources.ResourceResource {
		return
	}

	var docCopy = s.GetBaseDocCopy()

	if annotations.IsEnabled(resource, annotations.ActionApi) {
		for path, pathData := range docCopy.Paths {
			if strings.HasPrefix(path, "/{resourceActionPath}") {
				applyTemplate(pathData, resource)

				path = strings.ReplaceAll(path, "/{resourceActionPath}", "/"+util.ResourceRestPath(resource))

				doc.Paths[path] = pathData
			}
		}
	} else {
		for path, pathData := range docCopy.Paths {
			if strings.HasPrefix(path, "/{resourceRestPath}") {
				applyTemplate(pathData, resource)

				path = strings.ReplaceAll(path, "/{resourceRestPath}", "/"+util.ResourceRestPath(resource))

				doc.Paths[path] = pathData
			}
		}

		var filterParameters []*openapi3.ParameterRef

		for _, property := range resource.Properties {
			if util.IsFilterableProperty(property.Type) {
				filterParameters = append(filterParameters, &openapi3.ParameterRef{
					Value: &openapi3.Parameter{
						Name:        property.Name,
						In:          "query",
						Description: "Filter by " + property.Name,
						Schema:      &openapi3.SchemaRef{Value: &openapi3.Schema{Type: "string"}},
					},
				})
			}
		}

		doc.Paths["/"+util.ResourceRestPath(resource)].Get.Parameters = append(doc.Paths["/"+util.ResourceRestPath(resource)].Get.Parameters, filterParameters...)
		doc.Paths["/"+util.ResourceRestPath(resource)+"/_watch"].Get.Parameters = filterParameters
	}

}

func (s *OpenApiBuilder) GetBaseDocCopy() *openapi3.T {
	var doc = new(openapi3.T)

	err := doc.UnmarshalJSON([]byte(openApiBaseContent))

	if err != nil {
		log.Fatal(err)
	}

	return doc
}

type templateCandidate interface {
	MarshalJSON() ([]byte, error)
	UnmarshalJSON(data []byte) error
}

func applyTemplate[T templateCandidate](data T, resource *model.Resource) {
	var byts, err = data.MarshalJSON()

	var tag string
	if resource.GetNamespace() == "default" {
		tag = resource.GetName() + " API"
	} else {
		tag = resource.GetNamespace() + " / " + resource.GetName() + " API"
	}

	if err != nil {
		log.Fatal(err)
	}

	var str = string(byts)

	schemaName := util.ResourceJsonSchemaName(resource)

	str = strings.ReplaceAll(str, "{resourceRestPath}", util.ResourceRestPath(resource))
	str = strings.ReplaceAll(str, "{tag}", tag)
	str = strings.ReplaceAll(str, "{schemaName}", schemaName)
	str = strings.ReplaceAll(str, "{name}", resource.Name)
	str = strings.ReplaceAll(str, "{namespace}", resource.Namespace)
	str = strings.ReplaceAll(str, "{title}", util.DePointer(resource.Title, resource.Name))
	str = strings.ReplaceAll(str, "{description}", util.DePointer(resource.Description, ""))

	err = data.UnmarshalJSON([]byte(str))

	if err != nil {
		log.Fatal(err)
	}
}
