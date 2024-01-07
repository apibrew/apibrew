package apbr

import (
	"errors"
	"github.com/apibrew/apibrew/pkg/apbr/flags"
	"github.com/apibrew/apibrew/pkg/formats/executor"
	"github.com/apibrew/apibrew/pkg/generator"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/util"
	"github.com/spf13/cobra"
)

var generatorCmd = &cobra.Command{
	Use:   "generate",
	Short: "generate - Generate codes",
	RunE: func(cmd *cobra.Command, args []string) error {
		sourceFile, err := cmd.Flags().GetString("source-file")

		if sourceFile == "" {
			parseRootFlags(cmd)
		}

		namespace, err := cmd.Flags().GetString("namespace")

		if err != nil {
			return err
		}

		path, err := cmd.Flags().GetString("path")

		if err != nil {
			return err
		}

		pkg, err := cmd.Flags().GetString("package")

		if err != nil {
			return err
		}

		platform, err := cmd.Flags().GetString("platform")

		if err != nil {
			return err
		}

		var resources []*model.Resource

		filters, err := cmd.Flags().GetStringSlice("filter")

		if err != nil {
			return err
		}

		selectorFlags.Filters = filters

		if sourceFile == "" {
			var selection = &flags.SelectedRecordsResult{}
			err = selectorFlags.Parse(selection, cmd, []string{"resources"})

			if err != nil {
				return err
			}

			resources = selection.Resources
		} else {
			applier := executor.NewExecutor(executor.COLLECT, GetClient(), false, false, false, "", flags.OverrideConfig{})

			err = applier.Apply(cmd.Context(), sourceFile, "yaml")

			if err != nil {
				return err
			}

			resources = applier.CollectedResources

			for _, resource := range resources {
				util.NormalizeResource(resource)
			}
		}

		if pkg == "" {
			pkg = "model"
		}

		var mappedResourceActions = map[*model.Resource][]*model.Resource{}

		if len(resources) == 0 {
			return errors.New("no resources matched the filter")
		}

		return generator.GenerateResourceCodes(platform, pkg, resources, mappedResourceActions, path, namespace)
	},
}

func init() {
	generatorCmd.PersistentFlags().StringP("path", "p", ".", "Path")
	generatorCmd.PersistentFlags().String("package", "", "Package")
	generatorCmd.PersistentFlags().String("platform", "", "Platform: [golang, javascript, typescript, java]")
	generatorCmd.PersistentFlags().StringSlice("filter", nil, "filter")
	generatorCmd.PersistentFlags().String("source-file", "", "Generate models from a source file")
	selectorFlags.Declare(generatorCmd)
}
