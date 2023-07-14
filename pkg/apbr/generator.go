package apbr

import (
	"github.com/apibrew/apibrew/pkg/apbr/flags"
	"github.com/apibrew/apibrew/pkg/generator"
	"github.com/apibrew/apibrew/pkg/model"
	resources2 "github.com/apibrew/apibrew/pkg/resources"
	"github.com/spf13/cobra"
)

var generatorCmd = &cobra.Command{
	Use:   "generate",
	Short: "generate - Generate codes",
	Run: func(cmd *cobra.Command, args []string) {
		parseRootFlags(cmd)

		namespace, err := cmd.Flags().GetString("namespace")
		check(err)

		system, err := cmd.Flags().GetBool("system")
		check(err)

		path, err := cmd.Flags().GetString("path")
		check(err)

		pkg, err := cmd.Flags().GetString("package")
		check(err)

		platform, err := cmd.Flags().GetString("platform")
		check(err)

		var resources []*model.Resource

		if !system {
			var selection = &flags.SelectedRecordsResult{}
			selectorFlags.Parse(selection, cmd, args)
			resources = selection.Resources
		} else {
			resources = resources2.GetAllSystemResources()
		}

		if pkg == "" {
			pkg = "model"
		}

		err = generator.GenerateResourceCodes(platform, pkg, resources, path, namespace)

		check(err)
	},
}

func init() {
	generatorCmd.PersistentFlags().StringP("path", "p", ".", "Path")
	generatorCmd.PersistentFlags().String("package", "", "Package")
	generatorCmd.PersistentFlags().Bool("system", false, "System only")
	generatorCmd.PersistentFlags().String("platform", "", "Platform: [golang, nodejs, typescript]")
	selectorFlags.Declare(generatorCmd)
}
