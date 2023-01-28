package main

import (
	"github.com/spf13/cobra"
	"github.com/tislib/data-handler/server/stub"
	"log"
)

var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "backup - for doing backup operation on data-handler",
	Run: func(cmd *cobra.Command, args []string) {
		parseRootFlags(cmd)
		initClient(cmd.Context())

		output, err := cmd.Flags().GetString("output")
		check(err)

		format, err := cmd.Flags().GetString("format")
		check(err)

		selectResource, err := cmd.Flags().GetStringArray("select-resource")
		check(err)

		selectDataSource, err := cmd.Flags().GetStringArray("select-data-source")
		check(err)

		if output == "" {
			log.Fatal("output should provided")
		}

		log.Print(output, format, selectResource, selectDataSource)

		resourceServiceClient.List(cmd.Context(), &stub.ListResourceRequest{
			Token: authToken,
		})

	},
}

func init() {
	backupCmd.PersistentFlags().StringP("output", "o", "", "Output file")
	backupCmd.PersistentFlags().StringP("format", "f", "protobuf", "Backup format")
	backupCmd.PersistentFlags().StringArray("select-resource", []string{}, "Select specific resources")
	backupCmd.PersistentFlags().StringArray("select-data-source", []string{}, "Select specific data sources")
	backupCmd.PersistentFlags().StringArray("query", []string{}, "Backup format")
}
