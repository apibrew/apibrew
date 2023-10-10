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

		var resourceActionsRecords, _, err2 = GetDhClient().ListRecords(cmd.Context(), service.RecordListParams{
			Namespace: resources2.ResourceActionResource.Namespace,
			Resource:  resources2.ResourceActionResource.Name,
			Limit:     1000000,
		})

		check(err2)

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

		err = generator.GenerateResourceCodes(platform, pkg, resources, mappedResourceActions, path, namespace)

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
