package lib

import (
	"context"
	"data-handler/server/stub"
)

func Setup() {
	// creating data sources
	listDataSourceResp, err := dataSourceServiceClient.List(context.TODO(), &stub.ListDataSourceRequest{})

	checkResp(listDataSourceResp, err)

	for _, ds := listDataSourceResp.Content {}

}
