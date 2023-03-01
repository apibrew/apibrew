package dhctl

import (
	"github.com/gosimple/slug"
	"github.com/spf13/cobra"
	"github.com/tislib/data-handler/pkg/generator"
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/stub"
	"os"
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

		resp, err := resourceServiceClient.List(cmd.Context(), &stub.ListResourceRequest{
			Token: authToken,
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

		for _, resource := range filteredResources {
			if namespace != resource.Namespace {
				continue
			}

			code := generator.GenerateGoResourceCode(resource, generator.GenerateResourceCodeParams{
				Package:   pkg,
				Resources: resp.Resources,
			})

			resourceFileName := slug.Make(resource.Namespace) + "-" + slug.Make(resource.Name) + ".go"

			err = os.WriteFile(path+"/"+resourceFileName, []byte(code), 0777)

			check(err)
		}
	},
}

func init() {
	generatorCmd.PersistentFlags().StringP("namespace", "n", "default", "Namespace")
	generatorCmd.PersistentFlags().StringP("path", "p", "", "Path")
	generatorCmd.PersistentFlags().String("package", "", "Package")
}
