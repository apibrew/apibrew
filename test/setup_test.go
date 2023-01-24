package test

import (
	"context"
	"data-handler/helper"
	"data-handler/logging"
	"data-handler/model"
	"data-handler/server/stub"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/metadata"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	Setup()

	code := m.Run()
	os.Exit(code)
}

func Setup() {
	log.SetLevel(log.TraceLevel)
	log.SetReportCaller(true)
	setupDataSources()
	setupResources()
}

func setupDataSources() {
	dataSources := []*model.DataSource{
		dataSourceDhTest,
		dhTest,
		dataSource1,
	}
	// creating data sources
	listDataSourceResp, err := dataSourceServiceClient.List(context.TODO(), &stub.ListDataSourceRequest{})

	if err != nil {
		panic(err)
		return
	}

	var dataSourcesForCreate []*model.DataSource

	for _, cd := range dataSources {
		found := false
		for _, ds := range listDataSourceResp.Content {

			if cd.Name == ds.Name {
				found = true
				*cd = *ds
				break
			}
		}

		if !found {
			dataSourcesForCreate = append(dataSourcesForCreate, cd)
		}
	}

	createRes, err := dataSourceServiceClient.Create(context.TODO(), &stub.CreateDataSourceRequest{
		Token:       "test-token",
		DataSources: dataSourcesForCreate,
	})

	if err != nil {
		panic(err)
		return
	}

	for _, cd := range dataSources {
		if cd.Id != "" {
			continue
		}

		found := false
		for _, ds := range createRes.DataSources {

			if cd.Name == ds.Name {
				found = true
				*cd = *ds
				break
			}
		}

		if !found {
			panic("Could not create data source: " + cd.Name)
		}
	}
}

func setupResources() {
	resources := []*model.Resource{
		richResource1,
	}
	richResource1.SourceConfig.DataSource = dhTest.Id
	// creating data sources
	listResourceResp, err := resourceServiceClient.List(context.TODO(), &stub.ListResourceRequest{})

	if err != nil {
		panic(err)
		return
	}

	var resourcesForCreate []*model.Resource

	for _, cd := range resources {
		found := false
		for _, ds := range listResourceResp.Resources {

			if cd.Name == ds.Name {
				found = true
				*cd = *ds
				cd.Id = ds.Id
				cd.SourceConfig.DataSource = ds.SourceConfig.DataSource
				break
			}
		}

		if !found {
			resourcesForCreate = append(resourcesForCreate, cd)
		}
	}

	createRes, err := resourceServiceClient.Create(context.TODO(), &stub.CreateResourceRequest{
		Token:          "test-token",
		Resources:      resourcesForCreate,
		DoMigration:    true,
		ForceMigration: true,
	})

	if err != nil {
		panic(err)
		return
	}

	for _, cd := range resources {
		if cd.Id != "" {
			continue
		}

		found := false
		for _, ds := range createRes.Resources {

			if cd.Name == ds.Name {
				found = true
				*cd = *ds
				break
			}
		}

		if !found {
			panic("Could not create data source: " + cd.Name)
		}
	}
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
