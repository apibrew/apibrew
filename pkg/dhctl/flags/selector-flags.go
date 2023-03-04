package flags

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tislib/data-handler/pkg/client"
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/stub"
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

	limit, _ := cmd.PersistentFlags().GetInt64("limit")
	offset, _ := cmd.PersistentFlags().GetInt64("offset")

	if len(args) == 0 {
		log.Fatal("type should be provided")
	}

	getType := args[0]

	if getType == "all" || getType == "*" {
		resp, err := s.client().GetResourceServiceClient().List(cmd.Context(), &stub.ListResourceRequest{
			Token: s.client().GetToken(),
		})

		check(err)

		result.Resources = resp.Resources

		for _, resource := range resp.Resources {
			if resource.Virtual {
				continue
			}

			s.readSelectData3(cmd.Context(), resource, result, limit, offset)
		}
	} else if getType == "type" || getType == "types" || getType == "resource" || getType == "resources" {
		resp, err := s.client().GetResourceServiceClient().List(cmd.Context(), &stub.ListResourceRequest{
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
		resourceResp, err := s.client().GetResourceServiceClient().GetByName(cmd.Context(), &stub.GetResourceByNameRequest{
			Token:     s.client().GetToken(),
			Namespace: namespace,
			Name:      getType,
		})

		check(err)

		s.readSelectData3(cmd.Context(), resourceResp.Resource, result, limit, offset)
	}

}

type SelectedRecordData struct {
	Resource *model.Resource
	Records  chan *model.Record
}

type SelectedRecordsResult struct {
	Records   []SelectedRecordData
	Resources []*model.Resource
}

func (s selectorFlags) readSelectData3(ctx context.Context, resource *model.Resource, result *SelectedRecordsResult, limit int64, offset int64) {
	resp, err := s.client().GetRecordServiceClient().ReadStream(ctx, &stub.ReadStreamRequest{
		Token:     s.client().GetToken(),
		Namespace: resource.Namespace,
		Resource:  resource.Name,
		Limit:     uint32(limit),
		Offset:    uint64(offset),
	})

	check(err)

	var res = SelectedRecordData{
		Resource: resource,
		Records:  make(chan *model.Record),
	}

	result.Records = append(result.Records, res)

	go func() {
		defer func() {
			close(res.Records)
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
}

func NewSelectorFlags(clientGetter func() client.DhClient) FlagHelper[*SelectedRecordsResult] {
	return &selectorFlags{client: clientGetter}
}
