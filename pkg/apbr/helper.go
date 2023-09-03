package apbr

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/status"
)

//func loadDataSourceByNameOrId(ctx context.Context, id string, name string) *resource_model.DataSource {
//	if id == "" {
//		resp := check2(GetDhClient().GetDataSourceClient().List(ctx, &stub.ListDataSourceRequest{
//			Token: GetDhClient().GetToken(),
//		}))
//
//		for _, item := range resp.Content {
//			if item.Name == name {
//				return item
//			}
//		}
//
//		log.Fatal("Datasource not found with name: " + name)
//	}
//
//	return check2(GetDhClient().GetDataSourceClient().Get(ctx, &stub.GetDataSourceRequest{
//		Token: GetDhClient().GetToken(),
//		Id:    id,
//	})).DataSource
//}

func check(err error) {
	if err != nil {
		st, isStatus := status.FromError(err)

		if isStatus {
			log.Fatalf(st.Message())
		} else {
			log.Fatal(err)
		}
	}
}

func getFlag(cmd *cobra.Command, commandName string, required bool) string {
	o, err := cmd.Flags().GetString(commandName)
	check(err)

	if o == "" && required {
		log.Fatalf("%s is required but not provided", commandName)
	}

	return o
}
