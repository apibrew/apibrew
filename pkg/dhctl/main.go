package dhctl

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dhctl",
	Short: "dhctl - client for data-handler",
	Long:  `dhctl is cli util for data-handler, you can do various of operations with dhctl`,
	Run: func(cmd *cobra.Command, args []string) {
		check(cmd.Usage())
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	rootCmd.AddCommand(applyCmd)
	rootCmd.AddCommand(generatorCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(dataSourceCmd)
	defineRootFlags(rootCmd)
}

func PrepareRootCmd() *cobra.Command {
	initGetCmd()

	return rootCmd
}
