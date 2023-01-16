package main

import (
	"data-handler/cli/output"
	"data-handler/grpc/stub"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

var describeCmd = &cobra.Command{
	Use:   "describe",
	Short: "describe - describe resource",
	Run: func(cmd *cobra.Command, args []string) {
		parseRootFlags(cmd)
		initClient(cmd.Context())

		if len(args) == 0 {
			log.Fatal("type should be provided")
		}

		resourceName := args[0]

		namespace := "default"

		if len(args) > 1 {
			namespace = args[0]
			resourceName = args[1]
		}

		if strings.Contains(resourceName, "/") {
			parts := strings.Split(resourceName, "/")
			namespace = parts[0]
			resourceName = parts[1]
		}

		writer := output.NewOutputWriter("console")

		resp, err := resourceServiceClient.GetByName(cmd.Context(), &stub.GetResourceByNameRequest{
			Token:     authToken,
			Namespace: namespace,
			Name:      resourceName,
		})

		check(err)

		checkError(resp.Error)

		writer.DescribeResource(resp.Resource)
	},
}

func init() {
	//getCmd.PersistentFlags().StringP("output", "o", "console", "Output format")
}
