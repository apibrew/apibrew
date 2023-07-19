package apbr

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/apibrew/apibrew/pkg/formats"
	"github.com/apibrew/apibrew/pkg/formats/yamlformat"
	"github.com/apibrew/apibrew/pkg/util"
	log "github.com/sirupsen/logrus"
	"github.com/yargevad/filepathx"
	"google.golang.org/grpc/status"
	"os"
	"strings"

	"github.com/apibrew/apibrew/pkg/apbr/flags"
	"github.com/apibrew/apibrew/pkg/formats/batch"
	"github.com/apibrew/apibrew/pkg/formats/hclformat"
	"github.com/spf13/cobra"
)

var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "apply - apply resources",
	RunE: func(cmd *cobra.Command, args []string) error {
		parseRootFlags(cmd)

		inputFilePathArr, err := cmd.Flags().GetStringArray("file")
		if err != nil {
			return fmt.Errorf("failed to get input file path: %w", err)
		}
		if inputFilePathArr == nil {
			return errors.New("file must be provided")
		}

		doMigration, err := cmd.Flags().GetBool("migrate")
		if err != nil {
			return fmt.Errorf("failed to get migration flag: %w", err)
		}

		dataOnly, err := cmd.Flags().GetBool("data-only")
		if err != nil {
			return fmt.Errorf("failed to get data-only flag: %w", err)
		}

		force, err := cmd.Flags().GetBool("force")
		if err != nil {
			return fmt.Errorf("failed to get force flag: %w", err)
		}

		format, err := cmd.Flags().GetString("format")

		if err != nil {
			return fmt.Errorf("failed to get format flag: %w", err)
		}

		for _, inputFilePath := range inputFilePathArr {
			log.Info("Apply pattern: ", inputFilePath, " ...")
			if strings.Contains(inputFilePath, "*") {
				filenames, err := filepathx.Glob(inputFilePath)

				if err != nil {
					log.Fatalf("failed to get files: %s", err)
					return nil
				}

				for _, filename := range filenames {
					log.Info("Apply file: ", filename)
					err = applyLocal(filename, doMigration, dataOnly, force, format, cmd, args)

					if err != nil {
						log.Fatalf("failed to apply file: %s", err)
						return nil
					}
				}
			} else {
				log.Info("Apply file: ", inputFilePath)
				err := applyLocal(inputFilePath, doMigration, dataOnly, force, format, cmd, args)

				if err != nil {
					log.Fatalf("failed to apply file: %s", err)
					return nil
				}
			}
		}

		log.Info("Done")

		return nil
	},
}

func applyLocal(inputFilePath string, doMigration bool, dataOnly bool, force bool, format string, cmd *cobra.Command, args []string) error {
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

	overrideConfig := new(flags.OverrideConfig)
	overrideFlags.Parse(overrideConfig, cmd, args)

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
			Token:          GetDhClient().GetToken(),
			DhClient:       GetDhClient(),
			DoMigration:    doMigration,
			ForceMigration: force,
			OverrideConfig: hclformat.OverrideConfig{
				Namespace:  overrideConfig.Namespace,
				DataSource: overrideConfig.DataSource,
			},
		})
		if err != nil {
			return fmt.Errorf("failed to create HCL executor: %w", err)
		}

	case format == "pbe":
		in, err := os.Open(inputFilePath)
		if err != nil {
			return fmt.Errorf("failed to open PBE file: %w", err)
		}
		defer in.Close()

		executor = batch.NewExecutor(batch.ExecutorParams{
			Input:          in,
			Token:          GetDhClient().GetToken(),
			ResourceClient: GetDhClient().GetResourceClient(),
			RecordClient:   GetDhClient().GetRecordClient(),
			DataOnly:       dataOnly,
			OverrideConfig: batch.OverrideConfig{
				Namespace:  overrideConfig.Namespace,
				DataSource: overrideConfig.DataSource,
			},
		})
		if err != nil {
			return fmt.Errorf("failed to create PBE executor: %w", err)
		}

	case format == "yaml":
		in, err := os.Open(inputFilePath)
		if err != nil {
			return fmt.Errorf("failed to open YAML file: %w", err)
		}
		defer in.Close()

		executor, err = yamlformat.NewExecutor(yamlformat.ExecutorParams{
			Input:          in,
			DhClient:       GetDhClient(),
			DoMigration:    doMigration,
			ForceMigration: force,
			Token:          GetDhClient().GetToken(),
			OverrideConfig: yamlformat.OverrideConfig{
				Namespace:  overrideConfig.Namespace,
				DataSource: overrideConfig.DataSource,
			},
		})
		if err != nil {
			return fmt.Errorf("failed to create YAML executor: %w", err)
		}

	default:
		return fmt.Errorf("unsupported file format: %s", inputFilePath)
	}

	if err := executor.Restore(context.Background()); err != nil {
		if _, found := status.FromError(err); found {
			errorCode := util.GetErrorCode(err)
			errorFields := util.GetErrorFields(err)

			errorFieldsJson, _ := json.MarshalIndent(errorFields, "", "  ")

			errorStr := fmt.Sprintf("errorCode: %s, errorFields: \n%s", errorCode, string(errorFieldsJson))

			return errors.New(errorStr)
		} else {
			return err
		}
	}

	return nil
}

func init() {
	applyCmd.PersistentFlags().StringArrayP("file", "f", nil, "Input file")
	applyCmd.PersistentFlags().StringP("namespace", "n", "default", "Namespace")
	applyCmd.PersistentFlags().BoolP("migrate", "m", true, "Migrate")
	applyCmd.PersistentFlags().Bool("force", false, "Force")
	applyCmd.PersistentFlags().Bool("data-only", false, "Data Only")
	applyCmd.PersistentFlags().String("format", "yaml", "[yaml, yml, hcl, pbe]")

	overrideFlags.Declare(applyCmd)
}
