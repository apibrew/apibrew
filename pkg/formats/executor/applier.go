package executor

import (
	"context"
	"errors"
	"github.com/apibrew/apibrew/pkg/apbr/flags"
	"github.com/apibrew/apibrew/pkg/client"
	"github.com/apibrew/apibrew/pkg/formats/reader"
	"github.com/apibrew/apibrew/pkg/formats/unstructured/ops"
	"github.com/apibrew/apibrew/pkg/model"
	log "github.com/sirupsen/logrus"
	"github.com/yargevad/filepathx"
	"strings"
)

type Mode string

const (
	APPLY  Mode = "APPLY"
	CREATE Mode = "CREATE"
	UPDATE Mode = "UPDATE"
	DELETE Mode = "DELETE"
)

type Executor struct {
	client         client.Client
	doMigration    bool
	dataOnly       bool
	force          bool
	overrideConfig flags.OverrideConfig
	mode           Mode
}

func (a *Executor) Apply(ctx context.Context, inputFilePath string, format string) error {
	return a.ApplyWithType(ctx, inputFilePath, format, "")
}

func (a *Executor) ApplyWithType(ctx context.Context, inputFilePath string, format string, givenType string) error {
	reader := reader.Reader{}

	ctx, cancel := context.WithCancel(ctx)

	defer cancel()

	unstructuredExecutor := &ops.Executor{
		Client: a.client,
		Type:   givenType,
		RecordHandler: func(namespace string, resource string, record *model.Record) error {
			if a.mode == APPLY {
				appliedRecord, err := a.client.ApplyRecord(ctx, namespace, resource, record)

				if err != nil {
					return err
				}

				record.Properties = appliedRecord.Properties
				return err
			} else if a.mode == CREATE {
				appliedRecord, err := a.client.CreateRecord(ctx, namespace, resource, record)

				if err != nil {
					return err
				}

				record.Properties = appliedRecord.Properties
				return err
			} else if a.mode == UPDATE {
				appliedRecord, err := a.client.UpdateRecord(ctx, namespace, resource, record)

				if err != nil {
					return err
				}

				record.Properties = appliedRecord.Properties
				return err
			} else if a.mode == DELETE {
				return a.client.DeleteRecord(ctx, namespace, resource, record) // fixme locate id if not exists
			} else {
				return errors.New("unknown mode")
			}
		},
		ResourceHandler: func(resource *model.Resource) error {
			defer func() {
				log.Info("Resource: ", resource.Name, " "+string(a.mode))
			}()
			if a.mode == APPLY {
				return a.client.ApplyResource(ctx, resource, a.doMigration, a.force)
			} else if a.mode == CREATE {
				return a.client.CreateResource(ctx, resource, a.doMigration, a.force)
			} else if a.mode == UPDATE {
				return a.client.UpdateResource(ctx, resource, a.doMigration, a.force)
			} else if a.mode == DELETE {
				return a.client.DeleteResource(ctx, resource.Id, a.doMigration, a.force)
			} else {
				return errors.New("unknown mode")
			}
		},
	}

	err := unstructuredExecutor.Init(ctx)

	if err != nil {
		return err
	}

	return reader.Read(ctx, inputFilePath, format, unstructuredExecutor.RestoreItem)
}

func (a *Executor) ApplyWithPattern(ctx context.Context, inputFilePath string, format string) error {
	log.Info("Apply pattern: ", inputFilePath, " ...")
	if strings.Contains(inputFilePath, "*") {
		filenames, err := filepathx.Glob(inputFilePath)

		if err != nil {
			log.Fatalf("failed to get files: %s", err)
			return nil
		}

		for _, filename := range filenames {
			log.Info("Apply file: ", filename)
			err = a.Apply(ctx, filename, format)

			if err != nil {
				log.Fatalf("failed to apply file: %s", err)
			}
		}
	} else {
		log.Info("Apply file: ", inputFilePath)
		err := a.Apply(ctx, inputFilePath, format)

		if err != nil {
			log.Fatalf("failed to apply file: %s", err)
		}
	}

	return nil
}

func NewExecutor(mode Mode, dhClient client.Client, doMigration bool, dataOnly bool, force bool, overrideConfig flags.OverrideConfig) *Executor {
	return &Executor{
		client:         dhClient,
		doMigration:    doMigration,
		dataOnly:       dataOnly,
		force:          force,
		overrideConfig: overrideConfig,
		mode:           mode,
	}
}
