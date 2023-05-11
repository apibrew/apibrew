package apbr

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tislib/apibrew/pkg/model"
	"github.com/tislib/apibrew/pkg/stub"
	"google.golang.org/grpc/status"
)

func loadDataSourceByNameOrId(ctx context.Context, id string, name string) *model.DataSource {
	if id == "" {
		resp := check2(GetDhClient().GetDataSourceClient().List(ctx, &stub.ListDataSourceRequest{
			Token: GetDhClient().GetToken(),
		}))

		for _, item := range resp.Content {
			if item.Name == name {
				return item
			}
		}

		log.Fatal("Datasource not found with name: " + name)
	}

	return check2(GetDhClient().GetDataSourceClient().Get(ctx, &stub.GetDataSourceRequest{
		Token: GetDhClient().GetToken(),
		Id:    id,
	})).DataSource
}

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

func check2[T any](val T, err error) T {
	if err != nil {
		st, isStatus := status.FromError(err)

		if isStatus {
			log.Fatalf(st.Message())
		} else {
			log.Fatal(err)
		}
	}

	return val
}

func getFlag(cmd *cobra.Command, commandName string, required bool) string {
	o, err := cmd.Flags().GetString(commandName)
	check(err)

	if o == "" && required {
		log.Fatalf("%s is required but not provided", commandName)
	}

	return o
}

func contains[T comparable](arr []T, item T) bool {
	for _, a := range arr {
		if a == item {
			return true
		}
	}

	return false
}
