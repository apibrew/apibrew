package test

import (
	"context"
	"data-handler/app"
	"data-handler/stub"
	"data-handler/stub/model"
	"data-handler/util"
	"fmt"
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

func withClient(fn func(container *SimpleAppGrpcContainer)) {
	withApp(func(application *app.App) {
		var opts []grpc.DialOption
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

		conn, err := grpc.Dial(application.Addr, opts...)

		defer conn.Close()

		container := &SimpleAppGrpcContainer{
			recordService:         stub.NewRecordServiceClient(conn),
			authenticationService: stub.NewAuthenticationServiceClient(conn),
			resourceService:       stub.NewResourceServiceClient(conn),
			dataSourceService:     stub.NewDataSourceServiceClient(conn),
		}

		fn(container)

		if err != nil {
			panic(err)
		}

	})
}

func withApp(exec func(application *app.App)) {
	application := new(app.App)

	application.Addr = "127.0.0.1:17912"

	application.SetInitData(prepareInitData())

	application.Init()

	go application.Serve()
	time.Sleep(10 * time.Millisecond)

	exec(application)

	defer application.Stop()
}

func withDataSource(t *testing.T, container *SimpleAppGrpcContainer, dataSource *model.DataSource, exec func(dataSource *model.DataSource)) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
			t.Error(r)
		}
	}()
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

	if !reflect.DeepEqual(len(res.DataSources), 1) {
		t.Error("Created datasource length is wrong", len(res.DataSources), 1)
	}

	exists := checkDataSourceExists(container, res.DataSources[0].Id)
	if !exists {
		t.Error("Datasource created but not exists")
		return
	}

	exec(res.DataSources[0])

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

	exists = checkDataSourceExists(container, res.DataSources[0].Id)
	if exists {
		t.Error("Datasource removed but exists")
	}
}

func checkDataSourceExists(container *SimpleAppGrpcContainer, id string) bool {
	_, err := container.dataSourceService.Get(context.TODO(), &stub.GetDataSourceRequest{
		Token: "test-token",
		Id:    id,
	})

	if err != nil {
		return false
	}
	return true
}

func withAutoLoadedResource(t *testing.T, container *SimpleAppGrpcContainer, dataSource *model.DataSource, mappingName string, exec func(resource *model.Resource)) {
	withDataSource(t, container, dataSource, func(dataSource *model.DataSource) {
		res, err := container.dataSourceService.PrepareResourceFromEntity(context.TODO(), &stub.PrepareResourceFromEntityRequest{
			Token:  "test-token",
			Id:     dataSource.Id,
			Entity: mappingName,
		})

		if err != nil {
			t.Error(err)
			return
		}

		container.resourceService.Create(context.TODO(), &stub.CreateResourceRequest{
			Token:          "test-token",
			Resources:      []*model.Resource{res.Resource},
			DoMigration:    false,
			ForceMigration: false,
		})

		exec(res.Resource)

		_, err = container.resourceService.Delete(context.TODO(), &stub.DeleteResourceRequest{
			Token:          "test-token",
			Ids:            []string{res.Resource.Name},
			DoMigration:    false,
			ForceMigration: false,
		})

		if err != nil {
			t.Error(err)
			return
		}
	})
}
