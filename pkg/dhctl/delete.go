package dhctl

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tislib/data-handler/pkg/stub"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete",
}

var deleteResourceMigrate *bool
var deleteResourceForce *bool
var deleteResourceId *string
var deleteResourceName *string
var deleteResourceNamespace *string
var deleteDataSourceId *string
var deleteDataSourceName *string

var deleteResource = &cobra.Command{
	Use: "resource",
	Run: func(cmd *cobra.Command, args []string) {
		if *deleteResourceId == "" && *deleteResourceName == "" {
			log.Fatal("Id or name must be provided")
		}

		if *deleteResourceId == "" {
			resp, err := GetDhClient().GetResourceServiceClient().GetByName(cmd.Context(), &stub.GetResourceByNameRequest{
				Token:     GetDhClient().GetToken(),
				Namespace: *deleteResourceNamespace,
				Name:      *deleteResourceName,
			})

			if err != nil {
				log.Fatal(err)
			}

			*deleteResourceId = resp.Resource.Id
		}

		_, err := GetDhClient().GetResourceServiceClient().Delete(cmd.Context(), &stub.DeleteResourceRequest{
			Token:          GetDhClient().GetToken(),
			Namespace:      *deleteResourceNamespace,
			Ids:            []string{*deleteResourceId},
			DoMigration:    *deleteResourceMigrate,
			ForceMigration: *deleteResourceForce,
		})

		if err != nil {
			log.Fatal(err)
		}
	},
}

var deleteDataSource = &cobra.Command{
	Use: "data-source",
	Run: func(cmd *cobra.Command, args []string) {
		if *deleteDataSourceId == "" && *deleteDataSourceName == "" {
			log.Fatal("Id or name must be provided")
		}

		if *deleteDataSourceId == "" {
			resp, err := GetDhClient().GetDataSourceServiceClient().List(cmd.Context(), &stub.ListDataSourceRequest{
				Token: GetDhClient().GetToken(),
			})

			if err != nil {
				log.Fatal(err)
			}

			for _, item := range resp.Content {
				if item.Name == *deleteDataSourceName {
					*deleteDataSourceId = item.Id
				}
			}

			if *deleteDataSourceId == "" {
				log.Fatal("Datasource not found with name: " + *deleteDataSourceName)
			}
		}

		_, err := GetDhClient().GetDataSourceServiceClient().Delete(cmd.Context(), &stub.DeleteDataSourceRequest{
			Token: GetDhClient().GetToken(),
			Ids:   []string{*deleteDataSourceId},
		})

		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	deleteResourceMigrate = deleteResource.PersistentFlags().BoolP("migrate", "m", true, "")
	deleteResourceForce = deleteResource.PersistentFlags().BoolP("force", "f", false, "")

	deleteResourceId = deleteResource.PersistentFlags().String("id", "", "Id of resource")
	deleteResourceName = deleteResource.PersistentFlags().String("name", "", "Id of resource")
	deleteResourceNamespace = deleteResource.PersistentFlags().StringP("namespace", "n", "default", "Namespace")

	deleteDataSourceId = deleteDataSource.PersistentFlags().String("id", "", "Id of resource")
	deleteDataSourceName = deleteDataSource.PersistentFlags().String("name", "", "Id of resource")

	deleteCmd.AddCommand(deleteResource)
	deleteCmd.AddCommand(deleteDataSource)
}
