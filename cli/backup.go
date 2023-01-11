package main

import (
	"data-handler/grpc/stub"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "backup - for doing backup operation on data-handler",
	Run: func(cmd *cobra.Command, args []string) {
		parseRootFlags(cmd)
		initClient(cmd.Context())

		resp, err := resourceServiceClient.List(cmd.Context(), &stub.ListResourceRequest{
			Token: authToken,
		})

		log.Print("Started")
	},
}
