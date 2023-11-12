package ops

import (
	"context"
	"errors"
	"github.com/apibrew/apibrew/pkg/client"
	"github.com/apibrew/apibrew/pkg/formats/unstructured"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resource_model"
	"github.com/apibrew/apibrew/pkg/resource_model/extramappings"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/structpb"
	"gopkg.in/yaml.v3"
	"io"
	"os"
	"strings"
)

type Executor struct {
	Client              client.Client
	resources           []*model.Resource
	resourcePropertyMap map[string]*model.ResourceProperty
	preprocessor        preprocessor

	ResourceHandler func(resource *model.Resource) error
	RecordHandler   func(namespace string, resource string, record *model.Record) error
}

func (e *Executor) RestoreItem(body unstructured.Unstructured) error {
	body, err := e.preprocessor.preprocess(body)

	if err != nil {
		return err
	}

	bodyStr, err := yaml.Marshal(body)

	log.Debug("Restoring item: \n", string(bodyStr))

	var elemType string
	var namespace string
	var resourceName string
	var ok bool

	if elemType, ok = body["type"].(string); !ok {
		return errors.New("type field is required on record yaml definition")
	}

	if namespace, ok = body["namespace"].(string); !ok {
		namespace = "default"
	}

	resourceName, _ = body["resource"].(string)

	delete(body, "type")
	delete(body, "resource")
	delete(body, "namespace")

	if err != nil {
		return err
	}

	if elemType == "resource" {
		record, err := unstructured.ToRecord(body)

		if err != nil {
			return err
		}

		resourceModel := resource_model.ResourceMapperInstance.FromRecord(record)

		resource := extramappings.ResourceFrom(resourceModel)
		resource.Namespace = namespace

		err = e.ResourceHandler(resource)

		if err != nil {
			log.Error("Error applying resource: ", resource.Namespace+"/"+resource.Name)
			return err
		}
	} else {
		if elemType != "record" {
			if strings.Contains(elemType, "/") {
				namespace = strings.Split(elemType, "/")[0]
				resourceName = strings.Split(elemType, "/")[1]
			} else {
				resourceName = elemType
			}
		}
		if resourceName == "" {
			return errors.New("resource field is required on record yaml definition")
		}

		var record = new(model.Record)

		if _, exists := body["properties"]; !exists {
			body["properties"] = make(unstructured.Unstructured)
		}

		for key, value := range body {
			if key != "properties" {
				body["properties"].(unstructured.Unstructured)[key] = value
				delete(body, key)
			}
		}

		err = unstructured.ToProtoMessage(body, record)

		if err != nil {
			return err
		}

		// fix type BYTES
		for key, value := range record.Properties {
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

							record.Properties[key] = fileContentStr
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
		writer:   &Writer{},
	}

	return nil
}

type OverrideConfig struct {
	Namespace  string
	DataSource string
}

type ParserFunc func(reader io.Reader, consumer func(data unstructured.Unstructured) error) error
