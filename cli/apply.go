package main

import (
	"data-handler/grpc/stub"
	"github.com/spf13/cobra"
	"log"
)

var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "apply - apply resources",
	Run: func(cmd *cobra.Command, args []string) {
		parseRootFlags(cmd)
		initClient(cmd.Context())

		file, err := cmd.Flags().GetString("file")
		check(err)

		if file == "" {
			log.Fatal("file should provided")
		}

		log.Print(file)

		resourceServiceClient.List(cmd.Context(), &stub.ListResourceRequest{
			Token: authToken,
		})

	},
}

func init() {
	applyCmd.PersistentFlags().StringP("file", "f", "", "Output file")
}
