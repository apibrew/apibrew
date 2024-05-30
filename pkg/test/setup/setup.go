package setup

import (
	"context"
	"errors"
	"github.com/apibrew/apibrew/pkg/abs"
	"github.com/apibrew/apibrew/pkg/helper"
	"github.com/apibrew/apibrew/pkg/logging"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resource_model"
	"github.com/apibrew/apibrew/pkg/resources"
	"github.com/apibrew/apibrew/pkg/stub"
	"github.com/apibrew/apibrew/pkg/util"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/metadata"
	"testing"
)

var Ctx context.Context

func init() {
	initClient()
	initTextContext()
	Records := []*resource_model.DataSource{
		DataSourceDhTest,
		DhTest,
		DataSource1,
	}
	resources := []*model.Resource{
		RichResource1,
		SimpleVirtualResource1,
	}
	SetupRecords(Ctx, Records)
	SetupResources(Ctx, resources)
}

func SetupRecords(ctx context.Context, Records []*resource_model.DataSource) {
	// creating data sources
	listDataSourceResp, err := RecordClient.List(ctx, &stub.ListRecordRequest{
		Namespace: resources.DataSourceResource.Namespace,
		Resource:  resources.DataSourceResource.Name,
	})

	if err != nil {
		log.Fatal(err)
	}

	var RecordsForCreate []*resource_model.DataSource

	for _, cd := range Records {
		found := false
		for _, dso := range listDataSourceResp.Content {
			ds := resource_model.DataSourceMapperInstance.FromRecord(dso)

			if cd.Name == ds.Name {
				found = true
				cd.Id = ds.Id
				break
			}
		}

		if !found {
			RecordsForCreate = append(RecordsForCreate, cd)
		}
	}

	if len(RecordsForCreate) > 0 {
		createRes, err := RecordClient.Create(ctx, &stub.CreateRecordRequest{
			Namespace: resources.DataSourceResource.Namespace,
			Resource:  resources.DataSourceResource.Name,
			Records:   abs.RecordLikeAsRecords(util.ArrayMap(RecordsForCreate, resource_model.DataSourceMapperInstance.ToRecord)),
		})

		if err != nil {
			panic(err)
		}

		for _, cd := range Records {
			if cd.Id != nil {
				continue
			}

			found := false
			for _, dsr := range createRes.Records {
				ds := resource_model.DataSourceMapperInstance.FromRecord(dsr)

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
	listResourceResp, err := ResourceClient.List(ctx, &stub.ListResourceRequest{})

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
		_, err := ResourceClient.Delete(ctx, &stub.DeleteResourceRequest{
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
				resourcesForUpdate = append(resourcesForUpdate, cd)
				break
			}
		}

		if !found {
			resourcesForCreate = append(resourcesForCreate, cd)
		}
	}

	if len(resourcesForUpdate) > 0 {
		updatedRes, err := ResourceClient.Update(ctx, &stub.UpdateResourceRequest{
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
		createRes, err := ResourceClient.Create(ctx, &stub.CreateResourceRequest{
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
	_, err := ResourceClient.Delete(ctx, &stub.DeleteResourceRequest{
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

	dhClient.UpdateTokenFromContext(Ctx)
}

func WithUserAuthenticationContext(ctx context.Context, username, password string) context.Context {
	resp, err := AuthenticationClient.Authenticate(ctx, &stub.AuthenticationRequest{
		Username: username,
		Password: password,
		Term:     model.TokenTerm_MIDDLE,
	})

	if err != nil {
		panic(err)
	}

	//nolint:all
	ctx = context.WithValue(ctx, "token", resp.Token.Content)

	return metadata.AppendToOutgoingContext(ctx, "token", resp.Token.Content)
}
