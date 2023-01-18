package lib

import (
	"context"
	"data-handler/grpc/stub"
)

func Setup() {
	// creating data sources
	listDataSourceResp, err := dataSourceServiceClient.List(context.TODO(), &stub.ListDataSourceRequest{})

	checkResp(listDataSourceResp, err)

	for _, ds := listDataSourceResp.Content {}

}
