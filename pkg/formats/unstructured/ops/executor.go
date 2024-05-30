package ops

import (
	"context"
	"errors"
	"github.com/apibrew/apibrew/pkg/abs"
	"github.com/apibrew/apibrew/pkg/client"
	"github.com/apibrew/apibrew/pkg/formats/unstructured"
	"github.com/apibrew/apibrew/pkg/formats/writer"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resource_model"
	"github.com/apibrew/apibrew/pkg/resource_model/extramappings"
	resources2 "github.com/apibrew/apibrew/pkg/resources"
	"github.com/apibrew/apibrew/pkg/service/validate"
	"github.com/apibrew/apibrew/pkg/util"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/structpb"
	"gopkg.in/yaml.v3"
	"io"
	"os"
	"reflect"
)

type Executor struct {
	Client              client.Client
	resources           []*model.Resource
	resourcePropertyMap map[string]*model.ResourceProperty
	preprocessor        preprocessor
	Type                string

	ResourceHandler func(resource *model.Resource) error
	RecordHandler   func(namespace string, resource string, record abs.RecordLike) error
}

func (e *Executor) RestoreItem(in unstructured.Unstructured) error {
	cfPath, err := os.Getwd()

	if err != nil {
		return err
	}

	processed, err := e.preprocessor.preprocess(cfPath+"/main.go", in)

	if err != nil {
		return err
	}

	var list []unstructured.Unstructured

	switch value := processed.(type) {
	case unstructured.Unstructured:
		list = append(list, value)
	case []unstructured.Unstructured:
		list = value
	case []interface{}:
		for _, item := range value {
			list = append(list, item.(unstructured.Unstructured))
		}
	default:
		return errors.New("Invalid type: " + reflect.TypeOf(processed).String())
	}

	for _, body := range list {
		bodyStr, err := yaml.Marshal(body)

		log.Debug("Restoring item: \n", string(bodyStr))

		var elemType = e.Type

		if typ, ok := body["type"].(string); ok {
			elemType = typ
			delete(body, "type")
		}

		if elemType == "" {
			return errors.New("type field or arg is required")
		}

		if err != nil {
			return err
		}

		if elemType == "resource" {
			record, err := unstructured.ToRecord(body)

			if err != nil {
				return err
			}

			if err = validate.Records(resources2.ResourceResource, []abs.RecordLike{record}, false); err != nil {
				return err
			}

			resourceModel := resource_model.ResourceMapperInstance.FromRecord(record)

			resource := extramappings.ResourceFrom(resourceModel)

			e.fixResource(resource)

			err = e.ResourceHandler(resource)

			if err != nil {
				log.Error("Error applying resource: ", resource.Namespace+"/"+resource.Name)
				return err
			}
		} else {
			resourceIdentity := util.ParseType(elemType)

			var namespace, resourceName string

			namespace = resourceIdentity.Namespace
			resourceName = resourceIdentity.Name

			if resourceName == "" || namespace == "" {
				log.Println(body)

				return errors.New("Resource not set: " + namespace + "/" + resourceName)
			}

			record, err := unstructured.ToRecord(body)

			if err != nil {
				return err
			}

			// fix type BYTES
			for key, value := range record.GetProperties() {
				var property = e.resourcePropertyMap[namespace+"/"+resourceName+"/"+key]

				if property == nil {
					return errors.New("Property not found: " + namespace + "/" + resourceName + "/" + key)
				}

				if property.Type == model.ResourceProperty_BYTES {
					if value.GetStructValue() != nil {
						if value.GetStructValue().Fields["include"] != nil {
							if value.GetStructValue().Fields["include"].GetStringValue() != "" {
								fileContent, err := os.ReadFile(value.GetStructValue().Fields["include"].GetStringValue())

								if err != nil {
									return err
								}

								fileContentStr, err := structpb.NewValue(fileContent)

								if err != nil {
									return err
								}

								record.GetProperties()[key] = fileContentStr
							}
						}
					}
				}
			}

			err = e.RecordHandler(namespace, resourceName, record)

			if err != nil {
				return err
			}
		}
	}

	return nil
}
func (e *Executor) Init(ctx context.Context) error {
	resources, err := e.Client.ListResources(ctx)

	if err != nil {
		return err
	}

	e.resources = resources
	e.resourcePropertyMap = make(map[string]*model.ResourceProperty)

	if len(e.resources) == 0 {
		panic("Could not load resources")
	}

	for _, item := range e.resources {
		for _, field := range item.Properties {
			e.resourcePropertyMap[item.Namespace+"/"+item.Name+"/"+field.Name] = field
		}
	}

	e.preprocessor = preprocessor{
		dhClient: e.Client,
		writer:   &writer.Writer{},
	}

	return nil
}

func (e *Executor) InitSystemOnly() {
	resources := resources2.GetAllSystemResources()

	e.resources = resources
	e.resourcePropertyMap = make(map[string]*model.ResourceProperty)

	if len(e.resources) == 0 {
		panic("Could not load resources")
	}

	for _, item := range e.resources {
		for _, field := range item.Properties {
			e.resourcePropertyMap[item.Namespace+"/"+item.Name+"/"+field.Name] = field
		}
	}

	e.preprocessor = preprocessor{
		dhClient: e.Client,
		writer:   &writer.Writer{},
	}
}

func (e *Executor) fixResource(resource *model.Resource) {
	var namedTypes = util.GetNamedMap(resource.Types)

	util.ResourceWalkProperties(resource, func(path string, prop *model.ResourceProperty) {
		if prop.Type == model.ResourceProperty_STRUCT {
			if util.DePointer(prop.TypeRef, "") == "$property" {
				if namedTypes["Property"] == nil {
					resource.Types = append(resource.Types, resources2.PropertyType)
					namedTypes = util.GetNamedMap(resource.Types)
				}
				prop.TypeRef = util.Pointer("Property")
			}
		}
	})
}

type OverrideConfig struct {
	Namespace  string
	DataSource string
}

type ParserFunc func(reader io.Reader, consumer func(data unstructured.Unstructured) error) error
