package dhctl

import (
	"github.com/gosimple/slug"
	"github.com/spf13/cobra"
	"github.com/tislib/data-handler/pkg/generator"
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/stub"
	"log"
	"os"
)

var generatorCmd = &cobra.Command{
	Use:   "generate",
	Short: "generate - Generate codes",
	Run: func(cmd *cobra.Command, args []string) {
		parseRootFlags(cmd)
		initClient(cmd.Context())

		namespace, err := cmd.Flags().GetString("namespace")
		check(err)

		path, err := cmd.Flags().GetString("path")
		check(err)

		pkg, err := cmd.Flags().GetString("package")
		check(err)

		generateObjects, err := cmd.Flags().GetStringArray("generateObjects")
		check(err)

		log.Println(namespace, path, pkg, generateObjects)

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
			if resource.Namespace == "system" {
				continue
			}
			code := generator.GenerateResourceCode(resource, generator.GenerateResourceCodeParams{
				Package: pkg,
			})

			resourceFileName := slug.Make(resource.Namespace) + "-" + slug.Make(resource.Name) + ".go"

			err = os.WriteFile(path+"/"+resourceFileName, []byte(code), 0777)

			check(err)
		}
	},
}

func init() {
	generatorCmd.PersistentFlags().StringP("namespace", "n", "empty", "Namespace")
	generatorCmd.PersistentFlags().StringP("path", "p", "", "Path")
	generatorCmd.PersistentFlags().String("package", "", "Package")
	generatorCmd.PersistentFlags().StringArray("generateObjects", []string{}, "Generate Objects(resource, mapping)")
}
