package flags

import (
	"context"
	"github.com/apibrew/apibrew/pkg/client"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/service"
	"github.com/apibrew/apibrew/pkg/service/annotations"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"strings"
)

type selectorFlags struct {
	client func() client.Client
}

func (s selectorFlags) Declare(cmd *cobra.Command) {
	cmd.PersistentFlags().StringP("namespace", "n", "default", "Namespace")
	cmd.PersistentFlags().String("name", "", "item name")
	cmd.PersistentFlags().String("names", "", "item names")
}

func (s selectorFlags) Parse(result *SelectedRecordsResult, cmd *cobra.Command, args []string) {
	name := getFlag(cmd, "name", false)
	names := getFlag(cmd, "names", false)
	namespace := getFlag(cmd, "namespace", false)

	backup, _ := cmd.PersistentFlags().GetBool("backup")
	limit, _ := cmd.PersistentFlags().GetInt64("limit")
	offset, _ := cmd.PersistentFlags().GetInt64("offset")

	var getType = "resource"
	if len(args) > 0 {
		getType = args[0]
	}

	if getType == "all" || getType == "*" {
		resources, err := s.client().ListResources(cmd.Context())

		check(err)

		result.Resources = resources

		for _, resource := range resources {
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
		resources, err := s.client().ListResources(cmd.Context())

		check(err)

		var filteredResources []*model.Resource

		if name != "" {
			for _, item := range resources {
				if item.Name == name {
					if namespace == "" || item.Namespace == namespace {
						filteredResources = append(filteredResources, item)
					}
				}
			}
		} else if names != "" {
			for _, ni := range strings.Split(names, ",") {
				for _, item := range resources {
					if item.Name == ni {
						if namespace == "" || item.Namespace == namespace {
							filteredResources = append(filteredResources, item)
						}
					}
				}
			}
		} else {
			for _, item := range resources {
				if namespace == "" || item.Namespace == namespace {
					filteredResources = append(filteredResources, item)
				}
			}
		}

		result.Resources = filteredResources
	} else {
		resource, err := s.client().GetResourceByName(cmd.Context(), namespace, getType)

		check(err)

		if backup && annotations.IsEnabled(resource, annotations.DisableBackup) {
			log.Printf("Skipping %s/%s [backup mode & Disable backup annotation enabled]\n", resource.Namespace, resource.Name)
			return
		}

		result.RecordProviders = []func() SelectedRecordData{
			func() SelectedRecordData {
				return s.readSelectData3(cmd.Context(), resource, backup, limit, offset)
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

	_, total, err := s.client().ListRecords(ctx, service.RecordListParams{
		Namespace: resource.Namespace,
		Resource:  resource.Name,
		Limit:     1,
		Offset:    uint64(offset),
	})

	check(err)

	recordsChan := make(chan *model.Record)

	err = s.client().ReadRecordStream(ctx, service.RecordListParams{
		Namespace:   resource.Namespace,
		Resource:    resource.Name,
		Limit:       uint32(limit),
		Offset:      uint64(offset),
		PackRecords: backup,
	}, recordsChan)
	check(err)

	var res = SelectedRecordData{
		Resource: resource,
		Total:    total,
		Records:  recordsChan,
	}

	return res
}

func NewSelectorFlags(clientGetter func() client.Client) FlagHelper[*SelectedRecordsResult] {
	return &selectorFlags{client: clientGetter}
}
