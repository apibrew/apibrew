package test

import (
	"context"
	"data-handler/app"
	"data-handler/helper"
	"data-handler/logging"
	"data-handler/model"
	"data-handler/server/stub"
	"data-handler/util"
	"fmt"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"reflect"
	"testing"
	"time"
)

type SimpleAppGrpcContainer struct {
	app.Container

	authenticationService stub.AuthenticationServiceClient
	dataSourceService     stub.DataSourceServiceClient
	resourceService       stub.ResourceServiceClient
	recordService         stub.RecordServiceClient
}

func (receiver SimpleAppGrpcContainer) GetRecordService() stub.RecordServiceClient {
	return receiver.recordService
}
func (receiver SimpleAppGrpcContainer) GetAuthenticationService() stub.AuthenticationServiceClient {
	return receiver.authenticationService
}
func (receiver SimpleAppGrpcContainer) GetResourceService() stub.ResourceServiceClient {
	return receiver.resourceService
}
func (receiver SimpleAppGrpcContainer) GetDataSourceService() stub.DataSourceServiceClient {
	return receiver.dataSourceService
}

var initData = prepareInitData()

var container *SimpleAppGrpcContainer
var application *app.App = new(app.App)

func init() {
	application = new(app.App)

	application.SetInitData(initData)

	application.Init()

	go application.Serve()
	time.Sleep(10 * time.Millisecond)

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(initData.Config.GrpcAddr, opts...)

	if err != nil {
		panic(err)
	}

	container = &SimpleAppGrpcContainer{
		recordService:         stub.NewRecordServiceClient(conn),
		authenticationService: stub.NewAuthenticationServiceClient(conn),
		resourceService:       stub.NewResourceServiceClient(conn),
		dataSourceService:     stub.NewDataSourceServiceClient(conn),
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

func withDataSource(ctx context.Context, t testing.TB, container *SimpleAppGrpcContainer, dataSource *model.DataSource, exec func(dataSource *model.DataSource)) {
	log.Print("[withDataSource]Step 1")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
			t.Error(r)
		}
	}()
	log.Print("Create data-source:", dataSource)
	res, err := container.dataSourceService.Create(ctx, &stub.CreateDataSourceRequest{
		Token:       "test-token",
		DataSources: []*model.DataSource{dataSource},
	})
	log.Print("[withDataSource]Step 2")

	if err != nil {
		t.Error(err)
		return
	}

	if res.Error != nil {
		t.Error(res.Error.Message)
		return
	}

	log.Print("[withDataSource]Step 3")

	log.Print("data-source created", res.DataSources[0].Id)

	if !reflect.DeepEqual(len(res.DataSources), 1) {
		t.Error("Created datasource length is wrong", len(res.DataSources), 1)
	}

	log.Print("[withDataSource]Step 4")

	exists := checkDataSourceExists(ctx, container, res.DataSources[0].Id)
	if !exists {
		t.Error("Datasource created but not exists")
		return
	}

	log.Print("[withDataSource]Step 5")

	exec(res.DataSources[0])

	log.Print("[withDataSource]Step 6")

	log.Print("data-source deleting", res.DataSources[0].Id)
	res2, err := container.dataSourceService.Delete(ctx, &stub.DeleteDataSourceRequest{
		Token: "test-token",
		Ids: util.ArrayMap(res.DataSources, func(t *model.DataSource) string {
			return t.Id
		}),
	})

	log.Print("[withDataSource]Step 7")

	if err != nil {
		t.Error(err)
		return
	}

	log.Print("[withDataSource]Step 8")

	if res2.Error != nil {
		t.Error(res.Error.Message)
		return
	}

	log.Print("[withDataSource]Step 9")

	log.Print("data-source deleted", res.DataSources[0].Id)

	exists = checkDataSourceExists(ctx, container, res.DataSources[0].Id)

	log.Print("[withDataSource]Step 10")

	if exists {
		t.Error("Datasource removed but exists")
	}

	log.Print("[withDataSource]Step 11")
}

func checkDataSourceExists(ctx context.Context, container *SimpleAppGrpcContainer, id string) bool {
	res, err := container.dataSourceService.Get(ctx, &stub.GetDataSourceRequest{
		Token: "test-token",
		Id:    id,
	})

	return res.Error == nil && err == nil && res.DataSource != nil
}

func withResource(ctx context.Context, t testing.TB, resource *model.Resource, exec func()) {
	log.Print("resource creating", resource)
	resource.Flags = &model.ResourceFlags{}

	res, err := container.resourceService.Create(ctx, &stub.CreateResourceRequest{
		Token:          "test-token",
		Resources:      []*model.Resource{resource},
		DoMigration:    true,
		ForceMigration: true,
	})
	if err != nil {
		t.Error(err)
		return
	}

	if res.Error != nil {
		if res.Error.Code == model.ErrorCode_ALREADY_EXISTS {
			res2, _ := container.resourceService.GetByName(ctx, &stub.GetResourceByNameRequest{
				Token:     "test-token",
				Namespace: resource.Namespace,
				Name:      resource.Name,
			})
			resource = res2.Resource

		} else {
			t.Error(res.Error.Message)
			return
		}
	} else {
		resource.Id = res.Resources[0].Id
	}

	log.Print("resource created", resource.Name)

	defer func() {
		log.Print("resource deleting", resource.Name)
		_, err = container.resourceService.Delete(ctx, &stub.DeleteResourceRequest{
			Token:          "test-token",
			Ids:            []string{resource.Id},
			DoMigration:    true,
			ForceMigration: true,
		})
	}()

	exec()

	log.Print("resource deleted", resource.Name)

}

func withAutoLoadedResource(ctx context.Context, t testing.TB, container *SimpleAppGrpcContainer, dataSource *model.DataSource, catalog, entity string, exec func(resource *model.Resource)) {
	withDataSource(ctx, t, container, dataSource, func(dataSource *model.DataSource) {
		log.Print("begin PrepareResourceFromEntity", catalog, entity, dataSource.Id)
		res, err := container.dataSourceService.PrepareResourceFromEntity(ctx, &stub.PrepareResourceFromEntityRequest{
			Token:   "test-token",
			Id:      dataSource.Id,
			Catalog: catalog,
			Entity:  entity,
		})

		if err != nil {
			t.Error(err)
			return
		}

		if res.Error != nil {
			t.Error(res.Error)
			return
		}

		var resourceId string

		defer func() {
			log.Print("begin delete resource without migration", res.Resource.Namespace, res.Resource.Name)
			deleteRes, err := container.resourceService.Delete(ctx, &stub.DeleteResourceRequest{
				Token:          "test-token",
				Ids:            []string{resourceId},
				DoMigration:    false,
				ForceMigration: false,
			})

			if err != nil {
				t.Error(err)
				return
			}

			if deleteRes.Error != nil {
				t.Error(deleteRes.Error)
				return
			}

			log.Info("resource deleted: " + res.Resource.Name)
		}()

		log.Print("finish PrepareResourceFromEntity", catalog, entity, dataSource.Id)

		log.Print("begin create resource without migration", res.Resource.Namespace, res.Resource.Name)
		createRes, err := container.resourceService.Create(ctx, &stub.CreateResourceRequest{
			Token:          "test-token",
			Resources:      []*model.Resource{res.Resource},
			DoMigration:    false,
			ForceMigration: false,
		})

		if err != nil {
			t.Error(err)
			return
		}

		if createRes.Error != nil {
			if createRes.Error.Code == model.ErrorCode_ALREADY_EXISTS {
				res2, _ := container.resourceService.GetByName(ctx, &stub.GetResourceByNameRequest{
					Token:     "test-token",
					Namespace: res.Resource.Namespace,
					Name:      res.Resource.Name,
				})
				resourceId = res2.Resource.Id

			} else {
				t.Error(res.Error.Message)
				return
			}
		} else {
			resourceId = createRes.Resources[0].Id
		}

		log.Print("finish create resource without migration", res.Resource.Namespace, res.Resource.Name)

		log.Print("Calling exec: ", res.Resource.Namespace, " ", res.Resource.Name, " ", res.Resource.SourceConfig.DataSource)
		exec(res.Resource)
		log.Print("Finished exec: ", res.Resource.Namespace, " ", res.Resource.Name, " ", res.Resource.SourceConfig.DataSource)
	})
}
