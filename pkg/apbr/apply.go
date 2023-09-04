package apbr

import (
	"context"
	"errors"
	"fmt"
	"github.com/apibrew/apibrew/pkg/apbr/flags"
	"github.com/apibrew/apibrew/pkg/formats/apply"
	log "github.com/sirupsen/logrus"
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

		overrideConfig := new(flags.OverrideConfig)
		overrideFlags.Parse(overrideConfig, cmd, args)

		applier := apply.NewApplier(GetDhClient(), doMigration, dataOnly, force, *overrideConfig)

		for _, inputFilePath := range inputFilePathArr {
			err = applier.ApplyWithPattern(context.TODO(), inputFilePath, format)
			if err != nil {
				log.Error(err)
			}
		}

		log.Info("Done")

		return nil
	},
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
