package yamlformat

import (
	"context"
	"github.com/apibrew/apibrew/pkg/client"
	"github.com/apibrew/apibrew/pkg/formats"
	"github.com/apibrew/apibrew/pkg/formats/unstructured"
	"github.com/apibrew/apibrew/pkg/formats/unstructured/ops"
	"gopkg.in/yaml.v3"
	"io"
)

type executor struct {
	params               ExecutorParams
	unstructuredExecutor *ops.Executor
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
	DoMigration    bool
	ForceMigration bool
	DataOnly       bool
	Recursive      bool
}

func NewExecutor(params ExecutorParams, ctx context.Context) (formats.Executor, error) {
	unstructuredExecutor := &ops.Executor{
		Params: ops.ExecutorParams{
			DhClient: params.DhClient,
			OverrideConfig: ops.OverrideConfig{
				Namespace:  params.OverrideConfig.Namespace,
				DataSource: params.OverrideConfig.DataSource,
			},
			DoMigration:    params.DoMigration,
			ForceMigration: params.ForceMigration,
			DataOnly:       params.DataOnly,
		},
	}

	err := unstructuredExecutor.Init(ctx)

	if err != nil {
		return nil, err
	}

	return &executor{
		params:               params,
		unstructuredExecutor: unstructuredExecutor,
	}, nil
}
