package apbr

import (
	"context"
	"github.com/apibrew/apibrew/pkg/apbr/flags"
	yamlformat "github.com/apibrew/apibrew/pkg/formats/yaml"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/service/annotations"
	"github.com/apibrew/apibrew/pkg/stub"
	"github.com/olekukonko/tablewriter"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"strconv"
	"strings"
)

var dataSourceId *string
var dataSourceName *string
var dataSourcePrepareEntityNames *string
var dataSourcePrepareCatalogs *string

func prepareResourcesFromDataSource(ctx context.Context, dataSource *model.DataSource) <-chan *model.Resource {
	ch := make(chan *model.Resource)

	go func() {
		defer func() {
			close(ch)
		}()
		catalogs := check2(GetDhClient().GetDataSourceClient().ListEntities(ctx, &stub.ListEntitiesRequest{
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

				res := check2(GetDhClient().GetDataSourceClient().PrepareResourceFromEntity(ctx, &stub.PrepareResourceFromEntityRequest{
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

		resp := check2(GetDhClient().GetDataSourceClient().Status(cmd.Context(), &stub.StatusRequest{
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

		catalogs := check2(GetDhClient().GetDataSourceClient().ListEntities(cmd.Context(), &stub.ListEntitiesRequest{
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
	RunE: func(cmd *cobra.Command, args []string) error {
		dataSource := loadDataSourceByNameOrId(cmd.Context(), *dataSourceId, *dataSourceName)

		ch := prepareResourcesFromDataSource(cmd.Context(), dataSource)

		var yamlWriter = yamlformat.NewWriter(os.Stdout)

		var overrideConfig = new(flags.OverrideConfig)
		overrideFlags.Parse(overrideConfig, cmd, args)

		for item := range ch {
			if overrideConfig.Namespace != "" {
				item.Namespace = overrideConfig.Namespace
			}

			if overrideConfig.DataSource != "" {
				item.SourceConfig.DataSource = overrideConfig.DataSource
			}

			err := yamlWriter.WriteResource(item)

			if err != nil {
				return err
			}
		}

		return nil
	},
}

func init() {
	dataSourceName = dataSourceCmd.PersistentFlags().String("name", "", "Data Source name")
	dataSourceId = dataSourceCmd.PersistentFlags().String("id", "", "Data Source Id")

	dataSourcePrepareEntityNames = dataSourcePrepareCmd.PersistentFlags().StringP("entity", "e", "*", "Select entities for resource preparation, default value is * for selection of all entities, you can use comma to select multiple entities")
	dataSourcePrepareCatalogs = dataSourcePrepareCmd.PersistentFlags().StringP("catalog", "c", "*", "Select catalogs for resource preparation, default value is * for selection of all catalogs, you can use comma to select multiple catalogs")

	dataSourceCmd.AddCommand(dataSourceStatusCmd)
	dataSourceCmd.AddCommand(dataSourcePrepareCmd)
	dataSourceCmd.AddCommand(dataSourceListEntitiesCmd)

	overrideFlags.Declare(dataSourcePrepareCmd)
}
