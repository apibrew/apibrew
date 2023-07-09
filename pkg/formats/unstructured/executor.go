package unstructured

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/apibrew/apibrew/pkg/client"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/stub"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/structpb"
	"gopkg.in/yaml.v3"
	"io"
	"os"
	"strings"
)

type Executor struct {
	Params              ExecutorParams
	resources           []*model.Resource
	resourceNameMap     map[string]*model.Resource
	resourcePropertyMap map[string]*model.ResourceProperty
	parser              Parser
	preprocessor        preprocessor
}

func (e *Executor) RestoreItem(ctx context.Context, body Unstructured) error {
	body = fixMaps(body).(Unstructured)

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

	switch elemType {
	case "resource":
		var resource = new(model.Resource)
		err = body.ToProtoMessage(resource)

		if err != nil {
			jsonData, _ := json.MarshalIndent(body, " ", "  ")
			for index, line := range strings.Split(strings.TrimSuffix(string(jsonData), "\n"), "\n") {
				fmt.Printf("%d: %s\n", index+1, line)
			}
			return err
		}

		resource.Namespace = namespace

		err = e.Params.DhClient.ApplyResource(ctx, resource, e.Params.DoMigration, e.Params.ForceMigration)

		if err != nil {
			return err
		}
	case "record":
		if resourceName == "" {
			return errors.New("resource field is required on record yaml definition")
		}

		var resource = e.resourceNameMap[namespace+"/"+resourceName]

		if resource == nil {
			return errors.New("Resource not found: " + namespace + "/" + resourceName)
		}

		// locating resource

		var record = new(model.Record)
		err = body.ToProtoMessage(record)

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

		err = e.Params.DhClient.ApplyRecord(ctx, resource, record)

		if err != nil {
			return err
		}
	case "namespace":
		var namespace = new(model.Namespace)

		err = body.ToProtoMessage(namespace)

		if err != nil {
			return err
		}

		err = e.Params.DhClient.Apply(ctx, namespace)

		if err != nil {
			return err
		}
	case "extension":
		var dataSource = new(model.Extension)

		err = body.ToProtoMessage(dataSource)

		if err != nil {
			return err
		}

		err = e.Params.DhClient.Apply(ctx, dataSource)

		if err != nil {
			return err
		}
	case "user":
		var user = new(model.User)

		err = body.ToProtoMessage(user)

		if err != nil {
			return err
		}

		err = e.Params.DhClient.Apply(ctx, user)

		if err != nil {
			return err
		}
	default:
		return errors.New("unknown type: " + elemType)
	}

	return nil
}
func (e *Executor) Init() error {
	resp, err := e.Params.DhClient.GetResourceClient().List(context.TODO(), &stub.ListResourceRequest{
		Token: e.Params.Token,
	})

	if err != nil {
		return err
	}

	e.resources = resp.Resources
	e.resourceNameMap = make(map[string]*model.Resource)
	e.resourcePropertyMap = make(map[string]*model.ResourceProperty)

	for _, item := range e.resources {
		e.resourceNameMap[item.Namespace+"/"+item.Name] = item

		for _, field := range item.Properties {
			e.resourcePropertyMap[item.Namespace+"/"+item.Name+"/"+field.Name] = field
		}
	}

	e.preprocessor = preprocessor{
		dhClient: e.Params.DhClient,
		writer:   &Writer{},
	}

	return nil
}

type OverrideConfig struct {
	Namespace  string
	DataSource string
}

type ExecutorParams struct {
	DhClient       client.DhClient
	OverrideConfig OverrideConfig
	Token          string
	DoMigration    bool
	ForceMigration bool
	DataOnly       bool
}

type Parser func(reader io.Reader, consumer func(data Unstructured) error) error
