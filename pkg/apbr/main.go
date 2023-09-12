package apbr

import (
	"github.com/apibrew/apibrew/pkg/apbr/flags"
	"github.com/apibrew/apibrew/pkg/client"
	"github.com/spf13/cobra"
)

var selectorFlags = flags.NewSelectorFlags(GetDhClient)
var overrideFlags = flags.NewOverrideFlags()
var dhClient client.Client

func GetDhClient() client.Client {
	return dhClient
}

var rootCmd = &cobra.Command{
	Use:   "apbr",
	Short: "apbr - client for apibrew",
	Long:  `apbr is cli util for apibrew, you can do various of operations with apbr`,
	Run: func(cmd *cobra.Command, args []string) {
		check(cmd.Usage())
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	rootCmd.AddCommand(applyCmd)
	rootCmd.AddCommand(generatorCmd)
	rootCmd.AddCommand(deleteCmd)
	//rootCmd.AddCommand(dataSourceCmd)
	defineRootFlags(rootCmd)
}

func PrepareRootCmd() *cobra.Command {
	initGetCmd()

	return rootCmd
}
