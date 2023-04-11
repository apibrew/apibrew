package dhctl

import (
	"github.com/spf13/cobra"
	"github.com/tislib/data-handler/pkg/generator/golang"
	"github.com/tislib/data-handler/pkg/generator/nodejs"
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/stub"
)

var generatorCmd = &cobra.Command{
	Use:   "generate",
	Short: "generate - Generate codes",
	Run: func(cmd *cobra.Command, args []string) {
		parseRootFlags(cmd)

		namespace, err := cmd.Flags().GetString("namespace")
		check(err)

		path, err := cmd.Flags().GetString("path")
		check(err)

		pkg, err := cmd.Flags().GetString("package")
		check(err)

		platform, err := cmd.Flags().GetString("platform")
		check(err)

		resp, err := GetDhClient().GetResourceClient().List(cmd.Context(), &stub.ListResourceRequest{
			Token: GetDhClient().GetToken(),
		})

		check(err)

		var filteredResources []*model.Resource

		if len(args) == 0 {
			filteredResources = resp.Resources
		} else {
			for _, resource := range resp.Resources {
				if contains(args, resource.Name) {
					filteredResources = append(filteredResources, resource)
				}
			}
		}

		if pkg == "" {
			pkg = "model"
		}

		switch platform {
		case "golang":
			err = golang.GenerateGoResourceCode(golang.GenerateResourceCodeParams{
				Namespace: namespace,
				Package:   pkg,
				Resources: resp.Resources,
				Path:      path,
			})
		case "nodejs":
			err = nodejs.GenerateResourceCode(nodejs.GenerateResourceCodeParams{
				Namespace: namespace,
				Package:   pkg,
				Resources: resp.Resources,
				Path:      path,
			})
		}

		check(err)
	},
}

func init() {
	generatorCmd.PersistentFlags().StringP("namespace", "n", "default", "Namespace")
	generatorCmd.PersistentFlags().StringP("path", "p", "", "Path")
	generatorCmd.PersistentFlags().String("package", "", "Package")
	generatorCmd.PersistentFlags().String("platform", "", "Platform: [golang, nodejs]")
}
