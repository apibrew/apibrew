package test2

import (
	"context"
	"data-handler/helper"
	"data-handler/logging"
	"data-handler/model"
	"data-handler/server/stub"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/metadata"
)

func Setup() {
	setupDataSources := []*model.DataSource{
		dataSourceDhTest,
	}
	// creating data sources
	listDataSourceResp, err := dataSourceServiceClient.List(context.TODO(), &stub.ListDataSourceRequest{})

	check(err)
	for _, cd := range setupDataSources {
		found := false
		for _, ds := range listDataSourceResp.Content {

			if cd.Name == ds.Name {
				found = true
				break
			}
		}

		if !found {

		}
	}

	dataSourceServiceClient.Create(context.TODO(), &stub.CreateDataSourceRequest{
		Token:       "test-token",
		DataSources: nil,
	})

}

func prepareTextContext() context.Context {
	ctx := context.TODO()

	clientTrackId := helper.RandStringRunes(8)

	ctx = logging.WithLogField(ctx, "clientTrackId", clientTrackId)
	ctx = context.WithValue(ctx, "clientTrackId", clientTrackId)
	ctx = metadata.AppendToOutgoingContext(ctx, "clientTrackId", clientTrackId)

	log.Info("Init test clientTrackId: ", clientTrackId)

	return ctx
}
