package apbradm

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "apbradm",
	Short: "apbradm - client for apibrew",
	Long:  `apbradm is cli util for apibrew cluster admin, you can do various of operations with apbradm`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Usage()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}

func PrepareRootCmd() *cobra.Command {
	return rootCmd
}
