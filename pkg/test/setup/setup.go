package setup

import (
	"context"
	"errors"
	log "github.com/sirupsen/logrus"
	"github.com/tislib/data-handler/pkg/abs"
	"github.com/tislib/data-handler/pkg/helper"
	"github.com/tislib/data-handler/pkg/logging"
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/stub"
	"github.com/tislib/data-handler/pkg/util"
	"google.golang.org/grpc/metadata"
	"testing"
)

var Ctx context.Context

func init() {
	initClient()
	initTextContext()
	dataSources := []*model.DataSource{
		DataSourceDhTest,
		DhTest,
		DataSource1,
	}
	resources := []*model.Resource{
		RichResource1,
		SimpleVirtualResource1,
	}
	SetupDataSources(Ctx, dataSources)
	SetupResources(Ctx, resources)
}

func SetupDataSources(ctx context.Context, dataSources []*model.DataSource) {
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

	if len(dataSourcesForCreate) > 0 {
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
}

func PrepareResourcesForTest(t *testing.T, resources []*model.Resource) func() {
	return ResourcesWithErrorHandler(Ctx, resources, func(err error) {
		t.Error(err)
	})
}

func ResourcesWithErrorHandler(ctx context.Context, resources []*model.Resource, eh func(err error)) func() {
	// creating data sources
	listResourceResp, err := resourceClient.List(ctx, &stub.ListResourceRequest{})

	if err != nil {
		panic(err)
	}

	var resourcesForCreate []*model.Resource
	var resourcesForUpdate []*model.Resource
	var exists = false

	destroyHandler := func() {
		if !exists {
			return
		}
		_, err := resourceClient.Delete(ctx, &stub.DeleteResourceRequest{
			Token: GetTestDhClient().GetToken(),
			Ids: util.ArrayMap[*model.Resource, string](resources, func(t *model.Resource) string {
				return t.Id
			}),
			DoMigration:    true,
			ForceMigration: true,
		})

		if err != nil {
			log.Fatal(err)
		}
	}

	for _, cd := range resources {
		found := false
		for _, ds := range listResourceResp.Resources {
			if cd.Name == ds.Name {
				found = true
				cd.Id = ds.Id
				util.NormalizeResource(ds)
				resourcesForUpdate = append(resourcesForUpdate, ds)
				break
			}
		}

		if !found {
			resourcesForCreate = append(resourcesForCreate, cd)
		}
	}

	if len(resourcesForUpdate) > 0 {
		updatedRes, err := resourceClient.Update(ctx, &stub.UpdateResourceRequest{
			Resources:      resourcesForUpdate,
			DoMigration:    true,
			ForceMigration: true,
		})

		if err != nil {
			eh(err)
			return destroyHandler
		}

		exists = true

		for _, cd := range resourcesForUpdate {
			if cd.Id != "" {
				continue
			}

			found := false
			for _, ds := range updatedRes.Resources {
				if cd.Name == ds.Name {
					found = true
					cd.Id = ds.Id
					break
				}
			}

			if !found {
				eh(errors.New("Could not create resource: " + cd.Name))
				return destroyHandler
			}
		}
	}

	if len(resourcesForCreate) > 0 {
		createRes, err := resourceClient.Create(ctx, &stub.CreateResourceRequest{
			Resources:      resourcesForCreate,
			DoMigration:    true,
			ForceMigration: true,
		})

		if err != nil {
			eh(err)
			return destroyHandler
		}

		exists = true

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
				eh(errors.New("Could not create resource: " + cd.Name))
				return destroyHandler
			}
		}
	}

	return destroyHandler
}

func SetupResources(ctx context.Context, resources []*model.Resource) {
	ResourcesWithErrorHandler(ctx, resources, func(err error) {
		panic(err)
	})
}

func DestroyResources(ctx context.Context, resources []*model.Resource) {
	_, err := resourceClient.Delete(ctx, &stub.DeleteResourceRequest{
		Token: GetTestDhClient().GetToken(),
		Ids: util.ArrayMap[*model.Resource, string](resources, func(t *model.Resource) string {
			return t.Id
		}),
		DoMigration:    true,
		ForceMigration: true,
	})

	if err != nil {
		log.Fatal(err)
	}
}

func initTextContext() {
	Ctx = context.TODO()

	clientTrackId := helper.RandStringRunes(8)

	Ctx = logging.WithLogField(Ctx, "clientTrackId", clientTrackId)
	Ctx = context.WithValue(Ctx, abs.ClientTrackIdContextKey, clientTrackId)
	Ctx = metadata.AppendToOutgoingContext(Ctx, "clientTrackId", clientTrackId)

	log.Info("Init test clientTrackId: ", clientTrackId)
	Ctx = WithUserAuthenticationContext(Ctx, "admin", "admin")
}

func WithUserAuthenticationContext(ctx context.Context, username, password string) context.Context {
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
