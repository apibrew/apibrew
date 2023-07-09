package apbr

import (
	"github.com/spf13/cobra"
)

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
