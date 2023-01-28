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

var ctx context.Context

func TestMain(m *testing.M) {
	Setup()

	code := m.Run()
	os.Exit(code)
}

func Setup() {
	log.SetLevel(log.TraceLevel)
	log.SetReportCaller(true)
	initTextContext()
	setupDataSources(ctx)
	setupResources(ctx)
}

func setupDataSources(ctx context.Context) {
	dataSources := []*model.DataSource{
		dataSourceDhTest,
		dhTest,
		dataSource1,
	}
	// creating data sources
	listDataSourceResp, err := dataSourceServiceClient.List(ctx, &stub.ListDataSourceRequest{})

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

	createRes, err := dataSourceServiceClient.Create(ctx, &stub.CreateDataSourceRequest{
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

func setupResources(ctx context.Context) {
	resources := []*model.Resource{
		richResource1,
	}
	richResource1.SourceConfig.DataSource = dhTest.Id
	// creating data sources
	listResourceResp, err := resourceServiceClient.List(ctx, &stub.ListResourceRequest{})

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

	createRes, err := resourceServiceClient.Create(ctx, &stub.CreateResourceRequest{
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

func initTextContext() {
	ctx = context.TODO()

	clientTrackId := helper.RandStringRunes(8)

	ctx = logging.WithLogField(ctx, "clientTrackId", clientTrackId)
	ctx = context.WithValue(ctx, "clientTrackId", clientTrackId)
	ctx = metadata.AppendToOutgoingContext(ctx, "clientTrackId", clientTrackId)

	log.Info("Init test clientTrackId: ", clientTrackId)
	ctx = withUserAuthenticationContext(ctx, "admin", "admin")
}

func withUserAuthenticationContext(ctx context.Context, username, password string) context.Context {
	resp, err := authenticationServiceClient.Authenticate(ctx, &stub.AuthenticationRequest{
		Username: username,
		Password: password,
	})

	if err != nil {
		panic(err)
	}

	return metadata.AppendToOutgoingContext(context.TODO(), "token", resp.Token.Content)
}
