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
