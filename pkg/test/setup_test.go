package test

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/tislib/data-handler/pkg/abs"
	"github.com/tislib/data-handler/pkg/helper"
	"github.com/tislib/data-handler/pkg/logging"
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/stub"
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
	listDataSourceResp, err := dataSourceClient.List(ctx, &stub.ListDataSourceRequest{})

	if err != nil {
		panic(err)
	}

	var dataSourcesForCreate []*model.DataSource

	for _, cd := range dataSources {
		found := false
		for _, ds := range listDataSourceResp.Content {

			if cd.Name == ds.Name {
				found = true
				cd.Id = ds.Id
				break
			}
		}

		if !found {
			dataSourcesForCreate = append(dataSourcesForCreate, cd)
		}
	}

	createRes, err := dataSourceClient.Create(ctx, &stub.CreateDataSourceRequest{
		DataSources: dataSourcesForCreate,
	})

	if err != nil {
		panic(err)
	}

	for _, cd := range dataSources {
		if cd.Id != "" {
			continue
		}

		found := false
		for _, ds := range createRes.DataSources {

			if cd.Name == ds.Name {
				found = true
				cd.Id = ds.Id
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
		simpleVirtualResource1,
	}
	// creating data sources
	listResourceResp, err := resourceClient.List(ctx, &stub.ListResourceRequest{})

	if err != nil {
		panic(err)
	}

	var resourcesForCreate []*model.Resource

	for _, cd := range resources {
		found := false
		for _, ds := range listResourceResp.Resources {

			if cd.Name == ds.Name {
				found = true
				break
			}
		}

		if !found {
			resourcesForCreate = append(resourcesForCreate, cd)
		}
	}

	createRes, err := resourceClient.Create(ctx, &stub.CreateResourceRequest{
		Resources:      resourcesForCreate,
		DoMigration:    true,
		ForceMigration: true,
	})

	if err != nil {
		panic(err)
	}

	for _, cd := range resourcesForCreate {
		if cd.Id != "" {
			continue
		}

		found := false
		for _, ds := range createRes.Resources {
			if cd.Name == ds.Name {
				found = true
				cd.Id = ds.Id
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
	ctx = context.WithValue(ctx, abs.ClientTrackIdContextKey, clientTrackId)
	ctx = metadata.AppendToOutgoingContext(ctx, "clientTrackId", clientTrackId)

	log.Info("Init test clientTrackId: ", clientTrackId)
	ctx = withUserAuthenticationContext(ctx, "admin", "admin")
}

func withUserAuthenticationContext(ctx context.Context, username, password string) context.Context {
	resp, err := authenticationClient.Authenticate(ctx, &stub.AuthenticationRequest{
		Username: username,
		Password: password,
		Term:     model.TokenTerm_MIDDLE,
	})

	if err != nil {
		panic(err)
	}

	return metadata.AppendToOutgoingContext(context.TODO(), "token", resp.Token.Content)
}
