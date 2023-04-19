package dhctl

import (
	"context"
	"errors"
	"fmt"
	"github.com/tislib/data-handler/pkg/formats"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tislib/data-handler/pkg/dhctl/flags"
	"github.com/tislib/data-handler/pkg/formats/batch"
	"github.com/tislib/data-handler/pkg/formats/hclformat"
	"github.com/tislib/data-handler/pkg/formats/yaml"
)

var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "apply - apply resources",
	RunE: func(cmd *cobra.Command, args []string) error {
		parseRootFlags(cmd)

		inputFilePath, err := cmd.Flags().GetString("file")
		if err != nil {
			return fmt.Errorf("failed to get input file path: %w", err)
		}
		if inputFilePath == "" {
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

		overrideConfig := new(flags.OverrideConfig)
		overrideFlags.Parse(overrideConfig, cmd, args)

		var executor formats.Executor
		switch {
		case strings.HasSuffix(inputFilePath, ".hcl"):
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

		case strings.HasSuffix(inputFilePath, ".pbe"):
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

		case strings.HasSuffix(inputFilePath, ".yml"), strings.HasSuffix(inputFilePath, ".yaml"):
			in, err := os.Open(inputFilePath)
			if err != nil {
				return fmt.Errorf("failed to open YAML file: %w", err)
			}
			defer in.Close()

			executor, err = yamlformat.NewExecutor(yamlformat.ExecutorParams{
				Input:          in,
				Token:          GetDhClient().GetToken(),
				DhClient:       GetDhClient(),
				DoMigration:    doMigration,
				ForceMigration: force,
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
			return fmt.Errorf("failed to restore resources: %w", err)
		}

		return nil
	},
}

func init() {
	applyCmd.PersistentFlags().StringP("file", "f", "", "Input file")
	applyCmd.PersistentFlags().StringP("namespace", "n", "default", "Namespace")
	applyCmd.PersistentFlags().BoolP("migrate", "m", true, "Migrate")
	applyCmd.PersistentFlags().Bool("force", false, "Force")
	applyCmd.PersistentFlags().Bool("data-only", false, "Data Only")

	overrideFlags.Declare(applyCmd)
}
