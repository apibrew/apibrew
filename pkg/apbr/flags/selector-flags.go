package flags

import (
	"context"
	"github.com/apibrew/apibrew/pkg/client"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/service/annotations"
	"github.com/apibrew/apibrew/pkg/stub"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"io"
	"strings"
)

type selectorFlags struct {
	client func() client.DhClient
}

func (s selectorFlags) Declare(cmd *cobra.Command) {
	cmd.PersistentFlags().StringP("namespace", "n", "default", "Namespace")
	cmd.PersistentFlags().String("name", "", "Item name")
	cmd.PersistentFlags().String("names", "", "Item names")
}

func (s selectorFlags) Parse(result *SelectedRecordsResult, cmd *cobra.Command, args []string) {
	name := getFlag(cmd, "name", false)
	names := getFlag(cmd, "names", false)
	namespace := getFlag(cmd, "namespace", false)

	backup, _ := cmd.PersistentFlags().GetBool("backup")
	limit, _ := cmd.PersistentFlags().GetInt64("limit")
	offset, _ := cmd.PersistentFlags().GetInt64("offset")

	if len(args) == 0 {
		log.Fatal("type should be provided")
	}

	getType := args[0]

	if getType == "all" || getType == "*" {
		resp, err := s.client().GetResourceClient().List(cmd.Context(), &stub.ListResourceRequest{
			Token: s.client().GetToken(),
		})

		check(err)

		result.Resources = resp.Resources

		for _, resource := range resp.Resources {
			log.Println(resource.Name)
			if resource.Virtual {
				continue
			}

			if namespace != "" && resource.Namespace != namespace {
				continue
			}

			if backup && annotations.IsEnabled(resource, annotations.DisableBackup) {
				log.Printf("Skipping %s/%s [backup mode & Disable backup annotation enabled]\n", resource.Namespace, resource.Name)
				continue
			}
			func(localResource *model.Resource) {
				result.RecordProviders = append(result.RecordProviders, func() SelectedRecordData {
					return s.readSelectData3(cmd.Context(), localResource, backup, limit, offset)
				})
			}(resource)
		}
	} else if getType == "type" || getType == "types" || getType == "resource" || getType == "resources" {
		resp, err := s.client().GetResourceClient().List(cmd.Context(), &stub.ListResourceRequest{
			Token: s.client().GetToken(),
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

		result.Resources = filteredResources
	} else {
		resourceResp, err := s.client().GetResourceClient().GetByName(cmd.Context(), &stub.GetResourceByNameRequest{
			Token:     s.client().GetToken(),
			Namespace: namespace,
			Name:      getType,
		})

		check(err)

		if backup && annotations.IsEnabled(resourceResp.Resource, annotations.DisableBackup) {
			log.Printf("Skipping %s/%s [backup mode & Disable backup annotation enabled]\n", resourceResp.Resource.Namespace, resourceResp.Resource.Name)
			return
		}

		result.RecordProviders = []func() SelectedRecordData{
			func() SelectedRecordData {
				return s.readSelectData3(cmd.Context(), resourceResp.Resource, backup, limit, offset)
			},
		}
	}

}

type SelectedRecordData struct {
	Total    uint32
	Resource *model.Resource
	Records  chan *model.Record
}

type SelectedRecordsResult struct {
	RecordProviders []func() SelectedRecordData
	Resources       []*model.Resource
}

func (s selectorFlags) readSelectData3(ctx context.Context, resource *model.Resource, backup bool, limit int64, offset int64) SelectedRecordData {
	log.Println("readSelectData3 1 " + resource.Name)

	countRes, err := s.client().GetRecordClient().List(ctx, &stub.ListRecordRequest{
		Token:     s.client().GetToken(),
		Namespace: resource.Namespace,
		Resource:  resource.Name,
		Limit:     1,
		Offset:    uint64(offset),
	})

	check(err)

	log.Println("readSelectData3 2")
	resp, err := s.client().GetRecordClient().ReadStream(ctx, &stub.ReadStreamRequest{
		Token:       s.client().GetToken(),
		Namespace:   resource.Namespace,
		Resource:    resource.Name,
		Limit:       uint32(limit),
		Offset:      uint64(offset),
		PackRecords: backup,
	})

	check(err)

	log.Println("readSelectData3 3")
	var res = SelectedRecordData{
		Resource: resource,
		Total:    countRes.Total,
		Records:  make(chan *model.Record),
	}

	go func() {
		defer func() {
			close(res.Records)
			err := resp.CloseSend()

			if err != nil {
				log.Fatal(err)
			}
		}()

		for {
			record := new(model.Record)

			err = resp.RecvMsg(record)

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatal(err)
			}

			res.Records <- record
		}
	}()

	return res
}

func NewSelectorFlags(clientGetter func() client.DhClient) FlagHelper[*SelectedRecordsResult] {
	return &selectorFlags{client: clientGetter}
}
