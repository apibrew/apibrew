package main

import (
	"context"
	"data-handler/cli/output"
	"data-handler/grpc/stub"
	"data-handler/model"
	"github.com/spf13/cobra"
	"log"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get - get type",
	Run: func(cmd *cobra.Command, args []string) {
		parseRootFlags(cmd)
		initClient(cmd.Context())

		o := getFlag(cmd, "output", true)
		name := getFlag(cmd, "name", false)
		workspace := getFlag(cmd, "workspace", false)

		if len(args) == 0 {
			log.Fatal("type should be provided")
		}

		getType := args[0]

		writer := output.NewOutputWriter(o)

		if getType == "type" || getType == "types" || getType == "resource" || getType == "resources" {
			if name != "" {
				resp, err := resourceServiceClient.GetByName(cmd.Context(), &stub.GetResourceByNameRequest{
					Token:     authToken,
					Workspace: workspace,
					Name:      name,
				})

				check(err)

				writer.WriteResources([]*model.Resource{resp.Resource})
			} else {
				resp, err := resourceServiceClient.List(cmd.Context(), &stub.ListResourceRequest{
					Token: authToken,
				})

				check(err)

				writer.WriteResources(resp.Resources)
			}
		} else {
			resp, err := recordServiceClient.List(context.TODO(), &stub.ListRecordRequest{
				Token:     authToken,
				Workspace: workspace,
				Resource:  getType,
			})

			check(err)

			writer.WriteRecords(resp.Content)
		}
	},
}

func init() {
	getCmd.PersistentFlags().StringP("output", "o", "console", "Output format")
	getCmd.PersistentFlags().StringP("workspace", "w", "default", "Output format")
	getCmd.PersistentFlags().StringP("name", "n", "", "Output format")
}
