package apbr

import (
	"github.com/apibrew/apibrew/pkg/apbr/flags"
	"github.com/apibrew/apibrew/pkg/generator/golang"
	"github.com/apibrew/apibrew/pkg/generator/typescript"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
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

		var selection = &flags.SelectedRecordsResult{}

		selectorFlags.Parse(selection, cmd, args)

		check(err)

		if pkg == "" {
			pkg = "model"
		}

		switch platform {
		case "golang":
			err = golang.GenerateGoResourceCode(golang.GenerateResourceCodeParams{
				Namespace: namespace,
				Package:   pkg,
				Resources: selection.Resources,
				Path:      path,
			})
		case "typescript":
			err = typescript.GenerateResourceCode(typescript.GenerateResourceCodeParams{
				Namespace: namespace,
				Package:   pkg,
				Resources: selection.Resources,
				Path:      path,
			})
		default:
			log.Fatal("Unknown platform: " + platform)
		}

		check(err)
	},
}

func init() {
	generatorCmd.PersistentFlags().StringP("path", "p", ".", "Path")
	generatorCmd.PersistentFlags().String("package", "", "Package")
	generatorCmd.PersistentFlags().String("platform", "", "Platform: [golang, nodejs, typescript]")
	selectorFlags.Declare(generatorCmd)
}
