package apbr

import (
	"fmt"
	"github.com/apibrew/apibrew/pkg/apbr/flags"
	"github.com/apibrew/apibrew/pkg/formats/executor"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use: "delete",
	RunE: func(cmd *cobra.Command, args []string) error {
		parseRootFlags(cmd)

		inputFilePathArr, err := cmd.Flags().GetStringArray("file")
		if err != nil {
			return fmt.Errorf("failed to get input file path: %w", err)
		}
		doMigration, err := cmd.Flags().GetBool("migrate")
		if err != nil {
			return fmt.Errorf("failed to get migration flag: %w", err)
		}

		force, err := cmd.Flags().GetBool("force")
		if err != nil {
			return fmt.Errorf("failed to get force flag: %w", err)
		}

		if len(inputFilePathArr) > 0 {
			format, err := cmd.Flags().GetString("format")

			if err != nil {
				return fmt.Errorf("failed to get format flag: %w", err)
			}

			applier := executor.NewExecutor(executor.DELETE, GetClient(), doMigration, false, force, flags.OverrideConfig{})

			for _, inputFilePath := range inputFilePathArr {
				err = applier.ApplyWithPattern(cmd.Context(), inputFilePath, format)
				if err != nil {
					log.Error(err)
				}
			}
		} else {
			var selection = &flags.SelectedRecordsResult{}

			filters, err := cmd.Flags().GetStringSlice("filter")

			if err != nil {
				return err
			}

			selectorFlags.Filters = filters

			err = selectorFlags.Parse(selection, cmd, args)

			if err != nil {
				return err
			}

			for _, resource := range selection.Resources {
				err = GetClient().DeleteResource(cmd.Context(), resource.Id, doMigration, force)
				if err != nil {
					return err
				}
			}

			for _, record := range selection.Records {
				for _, r := range record.Records {
					err = GetClient().DeleteRecord(cmd.Context(), record.Resource.Namespace, record.Resource.Name, r)
					if err != nil {
						return err
					}
				}
			}
		}

		return nil
	},
}

func init() {
	deleteCmd.PersistentFlags().StringSlice("filter", nil, "filter")
	deleteCmd.PersistentFlags().StringArrayP("file", "f", nil, "path to file or directory to apply")
	deleteCmd.PersistentFlags().BoolP("migrate", "m", true, "Migrate")
	deleteCmd.PersistentFlags().Bool("force", false, "Force")
	deleteCmd.PersistentFlags().String("format", "yaml", "[yaml, json]")

	selectorFlags.Declare(deleteCmd)

}
