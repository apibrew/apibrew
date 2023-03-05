package dhctl

import (
	"context"
	"github.com/olekukonko/tablewriter"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tislib/data-handler/pkg/dhctl/flags"
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/service/annotations"
	"github.com/tislib/data-handler/pkg/stub"
	"os"
	"strconv"
	"strings"
)

var dataSourceId *string
var dataSourceName *string
var dataSourcePrepareEntityNames *string
var dataSourcePrepareCatalogs *string
var dataSourcePrepareApplyMigrate *bool

func prepareResourcesFromDataSource(ctx context.Context, dataSource *model.DataSource) <-chan *model.Resource {
	ch := make(chan *model.Resource)

	go func() {
		defer func() {
			close(ch)
		}()
		catalogs := check2(GetDhClient().GetDataSourceServiceClient().ListEntities(ctx, &stub.ListEntitiesRequest{
			Token: GetDhClient().GetToken(),
			Id:    dataSource.Id,
		})).Catalogs

		for _, catalog := range catalogs {

			if *dataSourcePrepareCatalogs != "*" {
				catalogs := strings.Split(*dataSourcePrepareCatalogs, ",")

				found := false

				for _, cc := range catalogs {
					if cc == catalog.Name {
						found = true
						break
					}
				}

				if !found {
					continue
				}
			}

			for _, entity := range catalog.Entities {
				if *dataSourcePrepareEntityNames != "*" {
					entities := strings.Split(*dataSourcePrepareEntityNames, ",")

					found := false

					for _, ec := range entities {
						if ec == entity.Name {
							found = true
							break
						}
					}

					if !found {
						continue
					}
				}

				res := check2(GetDhClient().GetDataSourceServiceClient().PrepareResourceFromEntity(ctx, &stub.PrepareResourceFromEntityRequest{
					Token:   GetDhClient().GetToken(),
					Id:      dataSource.Id,
					Catalog: catalog.Name,
					Entity:  entity.Name,
				}))

				resource := res.Resource

				if entity.ReadOnly {
					annotations.Enable(resource, annotations.DisableBackup)
				}

				ch <- resource
			}
		}
	}()

	return ch
}

var dataSourceCmd = &cobra.Command{
	Use:   "data-source",
	Short: "data-source - Data source related operations",
}

var dataSourceStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "status - Data source status",
	Run: func(cmd *cobra.Command, args []string) {
		dataSource := loadDataSourceByNameOrId(cmd.Context(), *dataSourceId, *dataSourceName)

		resp := check2(GetDhClient().GetDataSourceServiceClient().Status(cmd.Context(), &stub.StatusRequest{
			Token: GetDhClient().GetToken(),
			Id:    dataSource.Id,
		}))

		log.Println("DataSource name: " + dataSource.Name)
		log.Println("ConnectionAlreadyInitiated: " + strconv.FormatBool(resp.ConnectionAlreadyInitiated))
		log.Println("TestConnection: " + strconv.FormatBool(resp.TestConnection))
	},
}

var dataSourceListEntitiesCmd = &cobra.Command{
	Use:   "list-entities",
	Short: "list-entities - List existing entities in data source",
	Run: func(cmd *cobra.Command, args []string) {
		dataSource := loadDataSourceByNameOrId(cmd.Context(), *dataSourceId, *dataSourceName)

		catalogs := check2(GetDhClient().GetDataSourceServiceClient().ListEntities(cmd.Context(), &stub.ListEntitiesRequest{
			Token: GetDhClient().GetToken(),
			Id:    dataSource.Id,
		})).Catalogs

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Catalog", "Entity", "Editable"})
		table.SetAutoWrapText(false)
		table.SetAutoFormatHeaders(true)
		table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
		table.SetAlignment(tablewriter.ALIGN_LEFT)
		table.SetCenterSeparator("")
		table.SetColumnSeparator("")
		table.SetRowSeparator("")
		table.SetHeaderLine(false)
		table.SetBorder(false)
		table.SetTablePadding("\t") // pad with tabs
		table.SetNoWhiteSpace(true)

		for _, catalog := range catalogs {
			for _, entity := range catalog.Entities {
				status := "editable"

				if entity.ReadOnly {
					status = "read-only"
				}

				table.Append([]string{
					catalog.Name, entity.Name, status,
				})
			}
		}

		table.Render() // Send output
	},
}

var dataSourcePrepareCmd = &cobra.Command{
	Use:   "prepare",
	Short: "prepare - Prepare resources from existing data source entities",
}

var dataSourcePrepareDescribe = &cobra.Command{
	Use:   "describe",
	Short: "describe - Describe resource candidate from data source prepare",
	Run: func(cmd *cobra.Command, args []string) {
		dataSource := loadDataSourceByNameOrId(cmd.Context(), *dataSourceId, *dataSourceName)

		ch := prepareResourcesFromDataSource(cmd.Context(), dataSource)

		var overrideConfig = new(flags.OverrideConfig)
		overrideFlags.Parse(overrideConfig, cmd, args)

		for item := range ch {
			if overrideConfig.Namespace != "" {
				item.Namespace = overrideConfig.Namespace
			}

			if overrideConfig.DataSource != "" {
				item.SourceConfig.DataSource = overrideConfig.DataSource
			}

			describeWriter.WriteResources([]*model.Resource{item})
		}
	},
}

var dataSourcePrepareApply = &cobra.Command{
	Use:   "apply",
	Short: "apply - Apply resource candidate from data source prepare",
	Run: func(cmd *cobra.Command, args []string) {
		dataSource := loadDataSourceByNameOrId(cmd.Context(), *dataSourceId, *dataSourceName)

		ch := prepareResourcesFromDataSource(cmd.Context(), dataSource)

		var overrideConfig = new(flags.OverrideConfig)
		overrideFlags.Parse(overrideConfig, cmd, args)

		for item := range ch {
			if overrideConfig.Namespace != "" {
				item.Namespace = overrideConfig.Namespace
			}

			if overrideConfig.DataSource != "" {
				item.SourceConfig.DataSource = overrideConfig.DataSource
			}

			resource, err := GetDhClient().GetResourceServiceClient().GetByName(cmd.Context(), &stub.GetResourceByNameRequest{
				Token:     GetDhClient().GetToken(),
				Namespace: item.Namespace,
				Name:      item.Name,
			})

			if err == nil {
				item.Id = resource.Resource.Id
				check2(GetDhClient().GetResourceServiceClient().Update(cmd.Context(), &stub.UpdateResourceRequest{
					Token: GetDhClient().GetToken(),
					Resources: []*model.Resource{
						item,
					},
					DoMigration: *dataSourcePrepareApplyMigrate,
				}))
				log.Printf("Resource Updated: %s/%s \n", item.Namespace, item.Name)
			} else {
				check2(GetDhClient().GetResourceServiceClient().Create(cmd.Context(), &stub.CreateResourceRequest{
					Token: GetDhClient().GetToken(),
					Resources: []*model.Resource{
						item,
					},
					DoMigration: *dataSourcePrepareApplyMigrate,
				}))
				log.Printf("Resource Created: %s/%s \n", item.Namespace, item.Name)
			}
		}
	},
}

func init() {
	dataSourceName = dataSourceCmd.PersistentFlags().String("name", "", "Data Source name")
	dataSourceId = dataSourceCmd.PersistentFlags().String("id", "", "Data Source Id")

	dataSourcePrepareEntityNames = dataSourcePrepareCmd.PersistentFlags().StringP("entity", "e", "*", "Select entities for resource preparation, default value is * for selection of all entities, you can use comma to select multiple entities")
	dataSourcePrepareCatalogs = dataSourcePrepareCmd.PersistentFlags().StringP("catalog", "c", "*", "Select catalogs for resource preparation, default value is * for selection of all catalogs, you can use comma to select multiple catalogs")
	dataSourcePrepareApplyMigrate = dataSourcePrepareApply.PersistentFlags().BoolP("migrate", "m", false, "migrate")

	dataSourceCmd.AddCommand(dataSourceStatusCmd)
	dataSourceCmd.AddCommand(dataSourcePrepareCmd)
	dataSourceCmd.AddCommand(dataSourceListEntitiesCmd)

	dataSourcePrepareCmd.AddCommand(dataSourcePrepareDescribe)
	dataSourcePrepareCmd.AddCommand(dataSourcePrepareApply)

	overrideFlags.Declare(dataSourcePrepareCmd)
}
