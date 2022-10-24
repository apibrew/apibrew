package test

import (
	"context"
	"data-handler/app"
	"data-handler/stub"
	"data-handler/stub/model"
	"data-handler/util"
	"fmt"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"reflect"
	"testing"
	"time"
)

type SimpleAppGrpcContainer struct {
	app.GrpcContainer

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
	log.SetLevel(log.TraceLevel)
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

func withDataSource(t testing.TB, container *SimpleAppGrpcContainer, dataSource *model.DataSource, exec func(dataSource *model.DataSource)) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
			t.Error(r)
		}
	}()
	log.Print("Create data-source:", dataSource)
	res, err := container.dataSourceService.Create(context.TODO(), &stub.CreateDataSourceRequest{
		Token:       "test-token",
		DataSources: []*model.DataSource{dataSource},
	})

	if err != nil {
		t.Error(err)
		return
	}

	if res.Error != nil {
		t.Error(res.Error.Message)
		return
	}

	log.Print("data-source created", res.DataSources[0].Id)

	if !reflect.DeepEqual(len(res.DataSources), 1) {
		t.Error("Created datasource length is wrong", len(res.DataSources), 1)
	}

	exists := checkDataSourceExists(container, res.DataSources[0].Id)
	if !exists {
		t.Error("Datasource created but not exists")
		return
	}

	exec(res.DataSources[0])

	log.Print("data-source deleting", res.DataSources[0].Id)
	res2, err := container.dataSourceService.Delete(context.TODO(), &stub.DeleteDataSourceRequest{
		Token: "test-token",
		Ids: util.ArrayMap(res.DataSources, func(t *model.DataSource) string {
			return t.Id
		}),
	})

	if err != nil {
		t.Error(err)
		return
	}

	if res2.Error != nil {
		t.Error(res.Error.Message)
		return
	}

	log.Print("data-source deleted", res.DataSources[0].Id)

	exists = checkDataSourceExists(container, res.DataSources[0].Id)
	if exists {
		t.Error("Datasource removed but exists")
	}
}

func checkDataSourceExists(container *SimpleAppGrpcContainer, id string) bool {
	res, err := container.dataSourceService.Get(context.TODO(), &stub.GetDataSourceRequest{
		Token: "test-token",
		Id:    id,
	})

	return res.Error == nil && err == nil && res.DataSource != nil
}

func withResource(t testing.TB, resource *model.Resource, exec func()) {
	log.Print("resource creating", resource)
	res, err := container.resourceService.Create(context.TODO(), &stub.CreateResourceRequest{
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
		t.Error(res.Error.Message)
		return
	}
	log.Print("resource created", res.Resources[0].Name)

	exec()

	log.Print("resource deleting", res.Resources[0].Name)
	_, err = container.resourceService.Delete(context.TODO(), &stub.DeleteResourceRequest{
		Token:          "test-token",
		Ids:            []string{resource.Name},
		DoMigration:    true,
		ForceMigration: true,
	})

	if err != nil {
		t.Error(err)
		return
	}
	log.Print("resource deleted", res.Resources[0].Name)

}

func withAutoLoadedResource(t testing.TB, container *SimpleAppGrpcContainer, dataSource *model.DataSource, mappingName string, exec func(resource *model.Resource)) {
	withDataSource(t, container, dataSource, func(dataSource *model.DataSource) {
		log.Print("begin PrepareResourceFromEntity", mappingName, dataSource.Id)
		res, err := container.dataSourceService.PrepareResourceFromEntity(context.TODO(), &stub.PrepareResourceFromEntityRequest{
			Token:  "test-token",
			Id:     dataSource.Id,
			Entity: mappingName,
		})

		if err != nil {
			t.Error(err)
			return
		}

		if res.Error != nil {
			t.Error(res.Error)
			return
		}

		log.Print("finish PrepareResourceFromEntity", mappingName, dataSource.Id)

		log.Print("begin create resource without migration", res.Resource.Workspace, res.Resource.Name)
		createRes, err := container.resourceService.Create(context.TODO(), &stub.CreateResourceRequest{
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
			t.Error(createRes.Error)
			return
		}

		log.Print("finish create resource without migration", res.Resource.Workspace, res.Resource.Name)

		log.Print("Calling exec: ", res.Resource.Workspace, " ", res.Resource.Name, " ", res.Resource.SourceConfig.DataSource)
		exec(res.Resource)
		log.Print("Finished exec: ", res.Resource.Workspace, " ", res.Resource.Name, " ", res.Resource.SourceConfig.DataSource)

		log.Print("begin delete resource without migration", res.Resource.Workspace, res.Resource.Name)
		deleteRes, err := container.resourceService.Delete(context.TODO(), &stub.DeleteResourceRequest{
			Token:          "test-token",
			Ids:            []string{res.Resource.Name},
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
	})
}
