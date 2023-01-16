package main

import (
	"context"
	"data-handler/cli/output"
	"data-handler/grpc/stub"
	"data-handler/model"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get - get type",
	Run: func(cmd *cobra.Command, args []string) {
		parseRootFlags(cmd)
		initClient(cmd.Context())

		o := getFlag(cmd, "output", true)
		name := getFlag(cmd, "name", false)
		names := getFlag(cmd, "names", false)
		namespace := getFlag(cmd, "namespace", false)

		if len(args) == 0 {
			log.Fatal("type should be provided")
		}

		getType := args[0]

		writer := output.NewOutputWriter(o)

		if getType == "type" || getType == "types" || getType == "resource" || getType == "resources" {
			resp, err := resourceServiceClient.List(cmd.Context(), &stub.ListResourceRequest{
				Token: authToken,
			})

			check(err)

			checkError(resp.Error)

			var filteredResources []*model.Resource

			if name != "" {
				for _, item := range resp.Resources {
					if item.Name == name {
						filteredResources = append(filteredResources, item)
					}
				}
			} else if names != "" {
				for _, ni := range strings.Split(names, ",") {
					for _, item := range resp.Resources {
						if item.Name == ni {
							filteredResources = append(filteredResources, item)
						}
					}
				}
			} else {
				filteredResources = resp.Resources
			}

			writer.WriteResources(filteredResources)
		} else {
			resp, err := recordServiceClient.List(context.TODO(), &stub.ListRecordRequest{
				Token:     authToken,
				Namespace: namespace,
				Resource:  getType,
			})

			check(err)

			checkError(resp.Error)

			resourceResp, err := resourceServiceClient.GetByName(context.TODO(), &stub.GetResourceByNameRequest{
				Token:     authToken,
				Namespace: namespace,
				Name:      getType,
			})

			check(err)

			checkError(resourceResp.Error)

			writer.WriteRecords(resourceResp.Resource, resp.Content)
		}
	},
}

func init() {
	getCmd.PersistentFlags().StringP("output", "o", "console", "Output format")
	getCmd.PersistentFlags().StringP("namespace", "w", "default", "Namespace")
	getCmd.PersistentFlags().StringP("name", "n", "", "Item name")
	getCmd.PersistentFlags().String("names", "", "Item names")
}
