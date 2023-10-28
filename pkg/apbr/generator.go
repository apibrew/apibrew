package apbr

import (
	"github.com/apibrew/apibrew/pkg/apbr/flags"
	"github.com/apibrew/apibrew/pkg/generator"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resource_model"
	"github.com/apibrew/apibrew/pkg/resource_model/extramappings"
	resources2 "github.com/apibrew/apibrew/pkg/resources"
	"github.com/apibrew/apibrew/pkg/service"
	"github.com/spf13/cobra"
)

var generatorCmd = &cobra.Command{
	Use:   "generate",
	Short: "generate - Generate codes",
	RunE: func(cmd *cobra.Command, args []string) error {
		parseRootFlags(cmd)

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

		var selection = &flags.SelectedRecordsResult{}
		err = selectorFlags.Parse(selection, cmd, []string{"resources"})

		if err != nil {
			return err
		}

		filters, err := cmd.Flags().GetStringSlice("filter")

		if err != nil {
			return err
		}

		selectorFlags.Filters = filters

		resources = selection.Resources

		if pkg == "" {
			pkg = "model"
		}

		var resourceActionsRecords, _, err2 = GetClient().ListRecords(cmd.Context(), service.RecordListParams{
			Namespace: resources2.ResourceActionResource.Namespace,
			Resource:  resources2.ResourceActionResource.Name,
			Limit:     1000000,
		})

		if err2 != nil {
			return err2
		}

		var mappedResourceActions = map[*model.Resource][]*model.Resource{}

		for _, resourceActionRecord := range resourceActionsRecords {
			resourceAction := resource_model.ResourceActionMapperInstance.FromRecord(resourceActionRecord)

			res := &resource_model.Resource{
				Name:  resourceAction.Name,
				Types: resourceAction.Types,
			}

			if resourceAction.Input != nil {
				res.Types = append(res.Types, resource_model.SubType{
					Name:       resourceAction.Name + "Input",
					Properties: resourceAction.Input,
				})
			}

			if resourceAction.Output != nil {
				res.Properties = []resource_model.Property{
					*resourceAction.Output,
				}
			}

			for _, resource := range resources {
				if resource.Id == resourceAction.Resource.Id.String() {
					mappedResourceActions[resource] = append(mappedResourceActions[resource], extramappings.ResourceFrom(res))
				}
			}
		}

		return generator.GenerateResourceCodes(platform, pkg, resources, mappedResourceActions, path, namespace)
	},
}

func init() {
	generatorCmd.PersistentFlags().StringP("path", "p", ".", "Path")
	generatorCmd.PersistentFlags().String("package", "", "Package")
	generatorCmd.PersistentFlags().String("platform", "", "Platform: [golang, javascript, typescript, java]")
	generatorCmd.PersistentFlags().StringSlice("filter", nil, "filter")
	selectorFlags.Declare(generatorCmd)
}
