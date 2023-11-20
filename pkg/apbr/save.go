package apbr

import (
	"errors"
	"fmt"
	"github.com/apibrew/apibrew/pkg/apbr/flags"
	"github.com/apibrew/apibrew/pkg/formats/executor"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func execSave(mode executor.Mode, cmd *cobra.Command, args []string) error {
	parseRootFlags(cmd)

	inputFilePathArr, err := cmd.Flags().GetStringArray("file")
	if err != nil {
		return fmt.Errorf("failed to get input file path: %w", err)
	}
	if len(inputFilePathArr) == 0 {
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
	err = overrideFlags.Parse(overrideConfig, cmd, args)

	if err != nil {
		return fmt.Errorf("failed to parse override flags: %w", err)
	}

	var typ = ""

	if args != nil && len(args) > 0 {
		typ = args[0]
	}

	applier := executor.NewExecutor(mode, GetClient(), doMigration, dataOnly, force, typ, *overrideConfig)

	for _, inputFilePath := range inputFilePathArr {
		err = applier.ApplyWithPattern(cmd.Context(), inputFilePath, format)
		if err != nil {
			log.Error(err)
		}
	}

	log.Info("Done")

	return nil
}

var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "apply - apply resources/records docs: https://apibrew.io/docs/cli#apply",
	RunE: func(cmd *cobra.Command, args []string) error {
		return execSave(executor.APPLY, cmd, args)
	},
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create - update resources/records docs: https://apibrew.io/docs/cli#create",
	RunE: func(cmd *cobra.Command, args []string) error {
		return execSave(executor.CREATE, cmd, args)
	},
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update - update resources/records docs: https://apibrew.io/docs/cli#update",
	RunE: func(cmd *cobra.Command, args []string) error {
		return execSave(executor.UPDATE, cmd, args)
	},
}

func initCmd(cmd *cobra.Command) {
	cmd.PersistentFlags().StringArrayP("file", "f", nil, "path to file or directory to apply")
	cmd.PersistentFlags().StringP("namespace", "n", "default", "Namespace")
	cmd.PersistentFlags().BoolP("migrate", "m", true, "Migrate")
	cmd.PersistentFlags().Bool("force", false, "Force")
	cmd.PersistentFlags().Bool("data-only", false, "Data Only")
	cmd.PersistentFlags().Bool("schema-only", false, "Schema Only")
	cmd.PersistentFlags().String("format", "yaml", "[yaml, json]")

	overrideFlags.Declare(cmd)
}

func init() {
	initCmd(applyCmd)
	initCmd(createCmd)
	initCmd(updateCmd)
}
