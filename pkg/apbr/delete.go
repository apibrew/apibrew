package apbr

import (
	"github.com/apibrew/apibrew/pkg/stub"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
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
var deleteNamespaceId *string
var deleteNamespaceName *string
var deleteUserId *string
var deleteUserName *string
var deleteExtensionId *string
var deleteExtensionName *string

var deleteResource = &cobra.Command{
	Use: "resource",
	Run: func(cmd *cobra.Command, args []string) {
		if *deleteResourceId == "" && *deleteResourceName == "" {
			log.Fatal("Id or name must be provided")
		}

		if *deleteResourceId == "" {
			resp, err := GetDhClient().GetResourceClient().GetByName(cmd.Context(), &stub.GetResourceByNameRequest{
				Token:     GetDhClient().GetToken(),
				Namespace: *deleteResourceNamespace,
				Name:      *deleteResourceName,
			})

			if err != nil {
				log.Fatal(err)
			}

			*deleteResourceId = resp.Resource.Id
		}

		_, err := GetDhClient().GetResourceClient().Delete(cmd.Context(), &stub.DeleteResourceRequest{
			Token:          GetDhClient().GetToken(),
			Ids:            []string{*deleteResourceId},
			DoMigration:    *deleteResourceMigrate,
			ForceMigration: *deleteResourceForce,
		})

		if err != nil {
			log.Fatal(err)
		}
	},
}

//var deleteDataSource = &cobra.Command{
//	Use: "data-source",
//	Run: func(cmd *cobra.Command, args []string) {
//		if *deleteDataSourceId == "" && *deleteDataSourceName == "" {
//			log.Fatal("Id or name must be provided")
//		}
//
//		if *deleteDataSourceId == "" {
//			resp, err := GetDhClient().GetDataSourceClient().List(cmd.Context(), &stub.ListDataSourceRequest{
//				Token: GetDhClient().GetToken(),
//			})
//
//			if err != nil {
//				log.Fatal(err)
//			}
//
//			for _, item := range resp.Content {
//				if item.Name == *deleteDataSourceName {
//					*deleteDataSourceId = item.Id
//				}
//			}
//
//			if *deleteDataSourceId == "" {
//				log.Fatal("Datasource not found with name: " + *deleteDataSourceName)
//			}
//		}
//
//		_, err := GetDhClient().GetDataSourceClient().Delete(cmd.Context(), &stub.DeleteDataSourceRequest{
//			Token: GetDhClient().GetToken(),
//			Ids:   []string{*deleteDataSourceId},
//		})
//
//		if err != nil {
//			log.Fatal(err)
//		}
//	},
//}

var deleteNamespace = &cobra.Command{
	Use: "namespace",
	Run: func(cmd *cobra.Command, args []string) {
		if *deleteNamespaceId == "" && *deleteNamespaceName == "" {
			log.Fatal("Id or name must be provided")
		}

		if *deleteNamespaceId == "" {
			resp, err := GetDhClient().GetNamespaceClient().List(cmd.Context(), &stub.ListNamespaceRequest{
				Token: GetDhClient().GetToken(),
			})

			if err != nil {
				log.Fatal(err)
			}

			for _, item := range resp.Content {
				if item.Name == *deleteNamespaceName {
					*deleteNamespaceId = item.Id
				}
			}

			if *deleteNamespaceId == "" {
				log.Fatal("Namespace not found with name: " + *deleteNamespaceName)
			}
		}

		_, err := GetDhClient().GetNamespaceClient().Delete(cmd.Context(), &stub.DeleteNamespaceRequest{
			Token: GetDhClient().GetToken(),
			Ids:   []string{*deleteNamespaceId},
		})

		if err != nil {
			log.Fatal(err)
		}
	},
}

var deleteUser = &cobra.Command{
	Use: "user",
	Run: func(cmd *cobra.Command, args []string) {
		if *deleteUserId == "" && *deleteUserName == "" {
			log.Fatal("Id or name must be provided")
		}

		if *deleteUserId == "" {
			resp, err := GetDhClient().GetUserClient().List(cmd.Context(), &stub.ListUserRequest{
				Token: GetDhClient().GetToken(),
			})

			if err != nil {
				log.Fatal(err)
			}

			for _, item := range resp.Content {
				if item.Username == *deleteUserName {
					*deleteUserId = item.Id
				}
			}

			if *deleteUserId == "" {
				log.Fatal("User not found with name: " + *deleteUserName)
			}
		}

		_, err := GetDhClient().GetUserClient().Delete(cmd.Context(), &stub.DeleteUserRequest{
			Token: GetDhClient().GetToken(),
			Ids:   []string{*deleteUserId},
		})

		if err != nil {
			log.Fatal(err)
		}
	},
}

var deleteExtension = &cobra.Command{
	Use: "extension",
	Run: func(cmd *cobra.Command, args []string) {
		if *deleteExtensionId == "" && *deleteExtensionName == "" {
			log.Fatal("Id or name must be provided")
		}

		if *deleteExtensionId == "" {
			resp, err := GetDhClient().GetExtensionClient().List(cmd.Context(), &stub.ListExtensionRequest{
				Token: GetDhClient().GetToken(),
			})

			if err != nil {
				log.Fatal(err)
			}

			for _, item := range resp.Content {
				if item.Name == *deleteExtensionName {
					*deleteExtensionId = item.Id
				}
			}

			if *deleteExtensionId == "" {
				log.Fatal("Extension not found with name: " + *deleteExtensionName)
			}
		}

		_, err := GetDhClient().GetExtensionClient().Delete(cmd.Context(), &stub.DeleteExtensionRequest{
			Token: GetDhClient().GetToken(),
			Ids:   []string{*deleteExtensionId},
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

	//deleteDataSourceId = deleteDataSource.PersistentFlags().String("id", "", "Id of data-source")
	//deleteDataSourceName = deleteDataSource.PersistentFlags().String("name", "", "Id of data-source")

	deleteNamespaceId = deleteNamespace.PersistentFlags().String("id", "", "Id of namespace")
	deleteNamespaceName = deleteNamespace.PersistentFlags().String("name", "", "Id of namespace")

	deleteUserId = deleteUser.PersistentFlags().String("id", "", "Id of user")
	deleteUserName = deleteUser.PersistentFlags().String("username", "", "Username of user")

	deleteExtensionId = deleteExtension.PersistentFlags().String("id", "", "Id of extension")
	deleteExtensionName = deleteExtension.PersistentFlags().String("name", "", "name of extension")

	deleteCmd.AddCommand(deleteResource)
	//deleteCmd.AddCommand(deleteDataSource)
	deleteCmd.AddCommand(deleteNamespace)
	deleteCmd.AddCommand(deleteUser)
	deleteCmd.AddCommand(deleteExtension)
}
