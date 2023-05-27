package yamlformat

import (
	"context"
	"github.com/apibrew/apibrew/pkg/client"
	"github.com/apibrew/apibrew/pkg/formats"
	"github.com/apibrew/apibrew/pkg/formats/unstructured"
	"github.com/apibrew/apibrew/pkg/model"
	"gopkg.in/yaml.v3"
	"io"
)

type executor struct {
	params               ExecutorParams
	resources            []*model.Resource
	resourceNameMap      map[string]*model.Resource
	resourcePropertyMap  map[string]*model.ResourceProperty
	unstructuredExecutor *unstructured.Executor
}

func (e *executor) parser(r io.Reader, consumer func(data unstructured.Unstructured) error) error {
	decoder := yaml.NewDecoder(r)

	for {
		var body unstructured.Unstructured
		var err = decoder.Decode(&body)

		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		err = consumer(body)

		if err != nil {
			return err
		}
	}

	return nil
}

func (e *executor) Restore(ctx context.Context) error {
	return e.parser(e.params.Input, func(data unstructured.Unstructured) error {
		return e.unstructuredExecutor.RestoreItem(ctx, data)
	})
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
	unstructuredExecutor := &unstructured.Executor{
		Params: unstructured.ExecutorParams{
			DhClient: params.DhClient,
			OverrideConfig: unstructured.OverrideConfig{
				Namespace:  params.OverrideConfig.Namespace,
				DataSource: params.OverrideConfig.DataSource,
			},
			Token:          params.Token,
			DoMigration:    params.DoMigration,
			ForceMigration: params.ForceMigration,
			DataOnly:       params.DataOnly,
		},
	}

	err := unstructuredExecutor.Init()

	if err != nil {
		return nil, err
	}

	return &executor{
		params:               params,
		unstructuredExecutor: unstructuredExecutor,
	}, nil
}
