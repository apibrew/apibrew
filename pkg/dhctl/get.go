package dhctl

import (
	"github.com/spf13/cobra"
	"github.com/tislib/data-handler/pkg/dhctl/output"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get - get type",
	Run: func(cmd *cobra.Command, args []string) {
		parseRootFlags(cmd)
		initClient(cmd.Context())

		o := getFlag(cmd, "output", true)

		selection := selectData(cmd, args)

		writer := output.NewOutputWriter(o)

		if selection.resources != nil {
			writer.WriteResources(selection.resources)
		}

		for _, records := range selection.records {
			writer.WriteRecords(records.resource, records.records)
		}
	},
}

func init() {
	getCmd.PersistentFlags().StringP("output", "o", "console", "Output format")
	initSelectorFlags(getCmd)
}
