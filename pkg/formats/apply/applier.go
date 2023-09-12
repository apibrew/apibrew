package apply

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/apibrew/apibrew/pkg/apbr/flags"
	"github.com/apibrew/apibrew/pkg/client"
	"github.com/apibrew/apibrew/pkg/formats"
	"github.com/apibrew/apibrew/pkg/formats/hclformat"
	"github.com/apibrew/apibrew/pkg/formats/yamlformat"
	"github.com/apibrew/apibrew/pkg/util"
	log "github.com/sirupsen/logrus"
	"github.com/yargevad/filepathx"
	"google.golang.org/grpc/status"
	"os"
	"strings"
)

type Applier struct {
	dhClient       client.Client
	doMigration    bool
	dataOnly       bool
	force          bool
	overrideConfig flags.OverrideConfig
}

func (a *Applier) Apply(ctx context.Context, inputFilePath string, format string) error {
	if strings.HasSuffix(inputFilePath, ".hcl") {
		format = "hcl"
	} else if strings.HasSuffix(inputFilePath, ".pbe") {
		format = "pbe"
	} else if strings.HasSuffix(inputFilePath, ".yaml") || strings.HasSuffix(inputFilePath, ".yml") {
		format = "yaml"
	}

	if format == "yml" {
		format = "yaml"
	}

	var executor formats.Executor
	switch {
	case format == "hcl":
		in, err := os.Open(inputFilePath)
		if err != nil {
			return fmt.Errorf("failed to open HCL file: %w", err)
		}
		defer in.Close()

		executor, err = hclformat.NewExecutor(hclformat.ExecutorParams{
			Input:          in,
			DhClient:       a.dhClient,
			DoMigration:    a.doMigration,
			ForceMigration: a.force,
			OverrideConfig: hclformat.OverrideConfig{
				Namespace:  a.overrideConfig.Namespace,
				DataSource: a.overrideConfig.DataSource,
			},
		})
		if err != nil {
			return fmt.Errorf("failed to create HCL executor: %w", err)
		}

	case format == "yaml":
		in, err := os.Open(inputFilePath)
		if err != nil {
			return fmt.Errorf("failed to open YAML file: %w", err)
		}
		defer in.Close()

		executor, err = yamlformat.NewExecutor(yamlformat.ExecutorParams{
			Input:          in,
			DhClient:       a.dhClient,
			DoMigration:    a.doMigration,
			ForceMigration: a.force,
			OverrideConfig: yamlformat.OverrideConfig{
				Namespace:  a.overrideConfig.Namespace,
				DataSource: a.overrideConfig.DataSource,
			},
		}, ctx)
		if err != nil {
			return fmt.Errorf("failed to create YAML executor: %w", err)
		}

	default:
		return fmt.Errorf("unsupported file format: %s", inputFilePath)
	}

	if err := executor.Restore(ctx); err != nil {
		if _, found := status.FromError(err); found {
			errorCode := util.GetErrorCode(err)
			errorFields := util.GetErrorFields(err)

			errorFieldsJson, err2 := json.MarshalIndent(errorFields, "", "  ")

			if err2 != nil {
				return err2
			}

			return fmt.Errorf("message: %s, errorCode: %s, errorFields: \n%s", util.GetErrorMessage(err), errorCode, string(errorFieldsJson))
		} else {
			return err
		}
	}

	return nil
}

func (a *Applier) ApplyWithPattern(ctx context.Context, inputFilePath string, format string) error {
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

func NewApplier(dhClient client.Client, doMigration bool, dataOnly bool, force bool, overrideConfig flags.OverrideConfig) *Applier {
	return &Applier{
		dhClient:       dhClient,
		doMigration:    doMigration,
		dataOnly:       dataOnly,
		force:          force,
		overrideConfig: overrideConfig,
	}
}
