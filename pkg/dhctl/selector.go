package dhctl

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/stub"
	"io"
	"strings"
)

type SelectedRecordData struct {
	resource *model.Resource
	records  chan *model.Record
}

type SelectedRecordsResult struct {
	records   []SelectedRecordData
	resources []*model.Resource
}

func selectData(cmd *cobra.Command, args []string) SelectedRecordsResult {
	name := getFlag(cmd, "name", false)
	names := getFlag(cmd, "names", false)
	namespace := getFlag(cmd, "namespace", false)

	if len(args) == 0 {
		log.Fatal("type should be provided")
	}

	getType := args[0]

	var result SelectedRecordsResult

	if getType == "all" || getType == "*" {
		resp, err := resourceServiceClient.List(cmd.Context(), &stub.ListResourceRequest{
			Token: authToken,
		})

		check(err)

		result.resources = resp.Resources

		for _, resource := range resp.Resources {
			if resource.Virtual {
				continue
			}

			readSelectData3(cmd.Context(), resource, &result)
		}
	} else if getType == "type" || getType == "types" || getType == "resource" || getType == "resources" {
		resp, err := resourceServiceClient.List(cmd.Context(), &stub.ListResourceRequest{
			Token: authToken,
		})

		check(err)

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

		result.resources = filteredResources
	} else {
		resourceResp, err := resourceServiceClient.GetByName(cmd.Context(), &stub.GetResourceByNameRequest{
			Token:     authToken,
			Namespace: namespace,
			Name:      getType,
		})

		check(err)

		readSelectData3(cmd.Context(), resourceResp.Resource, &result)
	}

	return result
}

func readSelectData3(ctx context.Context, resource *model.Resource, result *SelectedRecordsResult) {
	resp, err := recordServiceClient.ReadStream(ctx, &stub.ReadStreamRequest{
		Token:     authToken,
		Namespace: resource.Namespace,
		Resource:  resource.Name,
	})

	check(err)

	var res = SelectedRecordData{
		resource: resource,
		records:  make(chan *model.Record),
	}

	result.records = append(result.records, res)

	go func() {
		defer func() {
			close(res.records)
		}()

		for {
			record := new(model.Record)

			err = resp.RecvMsg(record)

			if err == io.EOF {
				break
			}

			if err != nil {
				panic(err)
			}

			res.records <- record
		}
	}()
}

func initSelectorFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().StringP("namespace", "n", "default", "Namespace")
	cmd.PersistentFlags().String("name", "", "Item name")
	cmd.PersistentFlags().String("names", "", "Item names")
}
