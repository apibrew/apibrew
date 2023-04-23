package yamlformat

import (
	"context"
	"encoding/json"
	"errors"
	log "github.com/sirupsen/logrus"
	"github.com/tislib/apibrew/pkg/client"
	"github.com/tislib/apibrew/pkg/formats"
	"github.com/tislib/apibrew/pkg/model"
	"github.com/tislib/apibrew/pkg/stub"
	"google.golang.org/protobuf/encoding/protojson"
	"gopkg.in/yaml.v3"
	"io"
)

type executor struct {
	params          ExecutorParams
	resources       []*model.Resource
	resourceNameMap map[string]*model.Resource
}

func (e *executor) Restore(ctx context.Context) error {
	var jsonUMo = protojson.UnmarshalOptions{
		AllowPartial:   false,
		DiscardUnknown: false,
		Resolver:       nil,
	}

	decoder := yaml.NewDecoder(e.params.Input)

	for {
		var body map[string]interface{}
		var err = decoder.Decode(&body)

		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		body = convert(body).(map[string]interface{})

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

		jsonData, err := json.Marshal(body)

		if err != nil {
			return err
		}

		switch elemType {
		case "resource":
			var resource = new(model.Resource)
			err = jsonUMo.Unmarshal(jsonData, resource)

			if err != nil {
				log.Print(string(jsonData))
				return err
			}

			err = e.params.DhClient.ApplyResource(ctx, resource, e.params.DoMigration, e.params.ForceMigration)

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
			err = jsonUMo.Unmarshal(jsonData, record)

			if err != nil {
				return err
			}

			err = e.params.DhClient.ApplyRecord(ctx, resource, record)

			if err != nil {
				return err
			}
		case "datasource", "data-source", "dataSource":
			var dataSource = new(model.DataSource)

			err = jsonUMo.Unmarshal(jsonData, dataSource)

			if err != nil {
				return err
			}

			err = e.params.DhClient.Apply(ctx, dataSource)

			if err != nil {
				return err
			}
		case "namespace":
			var dataSource = new(model.Namespace)

			err = jsonUMo.Unmarshal(jsonData, dataSource)

			if err != nil {
				return err
			}

			err = e.params.DhClient.Apply(ctx, dataSource)

			if err != nil {
				return err
			}
		case "extension":
			var dataSource = new(model.Extension)

			err = jsonUMo.Unmarshal(jsonData, dataSource)

			if err != nil {
				return err
			}

			err = e.params.DhClient.Apply(ctx, dataSource)

			if err != nil {
				return err
			}
		case "user":
			var dataSource = new(model.User)

			err = jsonUMo.Unmarshal(jsonData, dataSource)

			if err != nil {
				return err
			}

			err = e.params.DhClient.Apply(ctx, dataSource)

			if err != nil {
				return err
			}
		default:
			return errors.New("unknown type: " + elemType)
		}
	}

	return nil
}

func (e *executor) init() error {
	resp, err := e.params.DhClient.GetResourceClient().List(context.TODO(), &stub.ListResourceRequest{})

	if err != nil {
		return err
	}

	e.resources = resp.Resources
	e.resourceNameMap = make(map[string]*model.Resource)

	for _, item := range e.resources {
		e.resourceNameMap[item.Namespace+"/"+item.Name] = item
	}

	return nil
}

func convert(i interface{}) interface{} {
	switch x := i.(type) {
	case map[interface{}]interface{}:
		m2 := map[string]interface{}{}
		for k, v := range x {
			// TODO: check if key is string
			m2[k.(string)] = convert(v)
		}
		return m2
	case map[string]interface{}:
		m2 := map[string]interface{}{}
		for k, v := range x {
			m2[k] = convert(v)
		}
		return m2
	case []interface{}:
		for i, v := range x {
			x[i] = convert(v)
		}
	}
	return i
}

type OverrideConfig struct {
	Namespace  string
	DataSource string
}

type ExecutorParams struct {
	Input          io.Reader
	DhClient       client.DhClient
	OverrideConfig OverrideConfig
	Token          string
	DoMigration    bool
	ForceMigration bool
	DataOnly       bool
}

func NewExecutor(params ExecutorParams) (formats.Executor, error) {
	exec := &executor{
		params: params,
	}

	return exec, exec.init()
}
