package flags

import (
	"context"
	"errors"
	"github.com/apibrew/apibrew/pkg/client"
	"github.com/apibrew/apibrew/pkg/formats/unstructured"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/service"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"strings"
)

type SelectorFlags struct {
	client       func() client.Client
	Filters      []string
	Limit        int64
	Offset       int64
	resourceName string
	namespace    string
	PackRecords  bool
}

func (s *SelectorFlags) Declare(cmd *cobra.Command) {
	cmd.PersistentFlags().StringP("namespace", "n", "default", "Namespace")
	cmd.PersistentFlags().String("name", "", "item name")
	cmd.PersistentFlags().String("names", "", "item names")
}

func (s *SelectorFlags) Parse(result *SelectedRecordsResult, cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return errors.New("resourceName name is required as argument")
	}

	if len(args) > 2 {
		return errors.New("too many arguments, it must be `namespace resourceName` or `resourceName` or `namespace/resourceName`")
	}

	if len(args) == 1 {
		s.resourceName = args[0]

		if strings.Contains(s.resourceName, "/") {
			parts := strings.Split(s.resourceName, "/")
			s.namespace = parts[0]
			s.resourceName = parts[1]
		} else {
			s.namespace = ""
		}
	} else {
		s.namespace = args[0]
		s.resourceName = args[1]
	}

	if s.namespace == "" && (s.resourceName == "all" || s.resourceName == "*") {

		resources, err := s.client().ListResources(cmd.Context())

		check(err)

		result.Resources = resources

		for _, resource := range resources {
			log.Println(resource.Name)
			if resource.Virtual {
				continue
			}

			var data, err = s.readSelectData3(cmd.Context(), resource)

			if err != nil {
				return err
			}

			result.Records = append(result.Records, data)
		}
	} else if s.namespace == "" && (s.resourceName == "type" || s.resourceName == "types" || s.resourceName == "resource" || s.resourceName == "resources") {
		s.resourceName = ""
		s.namespace = "default"

		for _, filter := range s.Filters {
			filterKey, filterValue := parseFilter(filter)

			if filterKey == "namespace" || filterKey == "namespace.name" {
				s.namespace = filterValue
			} else if filterKey == "name" {
				s.resourceName = filterValue
			}
		}

		resources, err := s.client().ListResources(cmd.Context())

		check(err)

		var filteredResources []*model.Resource

		for _, resource := range resources {
			if s.namespace != "" && resource.Namespace != s.namespace {
				continue
			}

			if s.resourceName != "" && resource.Name != s.resourceName {
				continue
			}

			filteredResources = append(filteredResources, resource)
		}

		result.Resources = filteredResources
	} else {
		resource, err := s.client().GetResourceByName(cmd.Context(), s.namespace, s.resourceName)

		if err != nil {
			return err
		}

		data, err := s.readSelectData3(cmd.Context(), resource)

		if err != nil {
			return err
		}

		result.Records = append(result.Records, data)
	}

	return nil
}

func parseFilter(filter string) (string, string) {
	middle := strings.Index(filter, "=")

	if middle == -1 {
		return filter, ""
	}

	return filter[:middle], filter[middle+1:]
}

type SelectedRecordData struct {
	Total    uint32
	Resource *model.Resource
	Records  []unstructured.Unstructured
}

type SelectedRecordsResult struct {
	Records   []SelectedRecordData
	Resources []*model.Resource
}

func (s *SelectorFlags) readSelectData3(ctx context.Context, resource *model.Resource) (SelectedRecordData, error) {
	var filters = make(map[string]string)

	for _, filter := range s.Filters {
		filterKey, filterValue := parseFilter(filter)
		filters[filterKey] = filterValue
	}

	result, total, err := s.client().ListRecords(ctx, service.RecordListParams{
		Namespace:   resource.Namespace,
		Resource:    resource.Name,
		Limit:       uint32(s.Limit),
		Offset:      uint64(s.Offset),
		PackRecords: s.PackRecords,
		Filters:     filters,
	})

	if err != nil {
		return SelectedRecordData{}, err
	}

	var res = SelectedRecordData{
		Resource: resource,
		Total:    total,
		Records:  result,
	}

	return res, nil
}

func NewSelectorFlags(clientGetter func() client.Client) *SelectorFlags {
	return &SelectorFlags{client: clientGetter}
}
