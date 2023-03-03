package dhctl

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tislib/data-handler/pkg/stub"
	"strconv"
)

var dataSourceId *string
var dataSourceName *string

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

func init() {
	dataSourceName = dataSourceCmd.PersistentFlags().StringP("name", "m", "", "Data Source name")
	dataSourceId = dataSourceCmd.PersistentFlags().StringP("id", "", "", "Data Source Id")

	dataSourceCmd.AddCommand(dataSourceStatusCmd)
}
