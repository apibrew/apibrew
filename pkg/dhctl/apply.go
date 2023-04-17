package dhctl

import (
	"context"
	"github.com/spf13/cobra"
	"github.com/tislib/data-handler/pkg/dhctl/flags"
	"github.com/tislib/data-handler/pkg/formats/batch"
	"github.com/tislib/data-handler/pkg/formats/hclformat"
	yamlformat "github.com/tislib/data-handler/pkg/formats/yaml"
	"log"
	"os"
	"strings"
)

var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "apply - apply resources",
	Run: func(cmd *cobra.Command, args []string) {
		parseRootFlags(cmd)

		file, err := cmd.Flags().GetString("file")
		check(err)

		migrate, err := cmd.Flags().GetBool("migrate")
		check(err)

		dataOnly, err := cmd.Flags().GetBool("data-only")
		check(err)

		force, err := cmd.Flags().GetBool("force")
		check(err)

		var overrideConfig = new(flags.OverrideConfig)
		overrideFlags.Parse(overrideConfig, cmd, args)

		if file == "" {
			log.Fatal("file should provided")
		}

		if strings.HasSuffix(file, ".hcl") {
			in, err := os.Open(file)

			check(err)

			hclExecutor, err := hclformat.NewExecutor(hclformat.ExecutorParams{
				Input:          in,
				Token:          GetDhClient().GetToken(),
				DhClient:       GetDhClient(),
				DoMigration:    migrate,
				ForceMigration: force,
				OverrideConfig: hclformat.OverrideConfig{
					Namespace:  overrideConfig.Namespace,
					DataSource: overrideConfig.DataSource,
				},
			})

			check(err)

			err = hclExecutor.Restore(context.TODO(), in)

			check(err)

			return
		} else if strings.HasSuffix(file, ".pbe") {
			in, err := os.Open(file)

			check(err)

			batchExecutor := batch.NewExecutor(batch.ExecutorParams{
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

			err = batchExecutor.Restore(context.TODO(), in)

			check(err)

			return
		} else if strings.HasSuffix(file, "yml") || strings.HasSuffix(file, "yaml") {
			in, err := os.Open(file)

			check(err)

			yamlExecutor, err := yamlformat.NewExecutor(yamlformat.ExecutorParams{
				Input:          in,
				Token:          GetDhClient().GetToken(),
				DhClient:       GetDhClient(),
				DoMigration:    migrate,
				ForceMigration: force,
				OverrideConfig: yamlformat.OverrideConfig{
					Namespace:  overrideConfig.Namespace,
					DataSource: overrideConfig.DataSource,
				},
			})

			check(err)

			err = yamlExecutor.Restore(context.TODO(), in)

			check(err)

			return
		}
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
