package test

import (
	"context"
	"data-handler/stub"
	"data-handler/stub/model"
	log "github.com/sirupsen/logrus"
	"reflect"
	"testing"
)

func TestCreateDataSource(t *testing.T) {
	withClient(func(container *SimpleAppGrpcContainer) {
		withDataSource(t, container, dataSource1, func(createdDataSource *model.DataSource) {
			// testing is done
		})
	})
}

func TestCreateAndReadDataSource(t *testing.T) {
	withClient(func(container *SimpleAppGrpcContainer) {
		withDataSource(t, container, dataSource1, func(createdDataSource *model.DataSource) {
			res2, err := container.dataSourceService.Get(context.TODO(), &stub.GetDataSourceRequest{
				Token: "test-token",
				Id:    createdDataSource.Id,
			})

			if err != nil {
				t.Error(err)
				return
			}

			if res2.DataSource == nil {
				t.Error("Data source must not be null")
				return
			}

			if !reflect.DeepEqual(createdDataSource, res2.DataSource) {
				log.Println(createdDataSource)
				log.Println(res2.DataSource)
				t.Error("Backend is different")
				return
			}
		})
	})
}

func TestCreateDataSourceStatusTest(t *testing.T) {
	withClient(func(container *SimpleAppGrpcContainer) {
		withDataSource(t, container, systemDataSource, func(createdDataSource *model.DataSource) {
			res, err := container.dataSourceService.Status(context.TODO(), &stub.StatusRequest{
				Token: "test-token",
				Id:    createdDataSource.Id,
			})

			if err != nil {
				t.Error(err)
			}

			if res.ConnectionAlreadyInitiated {
				t.Error("New created datasource should have initiated connection")
			}

			if !res.TestConnection {
				t.Error("New created connection should pass test connection")
			}
		})
	})
}

func TestCreateDataSourceWithWrongPasswordStatusTest(t *testing.T) {
	withClient(func(container *SimpleAppGrpcContainer) {
		withDataSource(t, container, dataSource1WrongPassword, func(createdDataSource *model.DataSource) {
			_, err := container.dataSourceService.Status(context.TODO(), &stub.StatusRequest{
				Token: "test-token",
				Id:    createdDataSource.Id,
			})

			if err == nil {
				t.Error("It should be unable to login to database")
				return
			}
		})
	})
}

func TestListCreatedDataSources(t *testing.T) {
	withClient(func(container *SimpleAppGrpcContainer) {
		withDataSource(t, container, dataSource1, func(createdDataSource1 *model.DataSource) {
			withDataSource(t, container, dataSource1, func(createdDataSource2 *model.DataSource) {
				withDataSource(t, container, dataSource1, func(createdDataSource3 *model.DataSource) {
					res, err := container.dataSourceService.List(context.TODO(), &stub.ListDataSourceRequest{
						Token: "test-token",
					})

					if err != nil {
						t.Error(err)
						return
					}

					if len(res.Content) <= 3 {
						t.Error("DataSourceList does not match: ", len(res.Content), 3)
					}
				})
			})
		})
	})
}
