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
		res, err := container.dataSourceService.Create(context.TODO(), &stub.CreateDataSourceRequest{
			Token: "test-token",
			DataSources: []*model.DataSource{
				{
					Backend: model.DataSourceBackend_POSTGRESQL,
					Options: &model.DataSource_PostgresqlParams{
						PostgresqlParams: &model.PostgresqlOptions{
							Username:      "root",
							Password:      "52fa536f0c5b85f9d806633937f06446",
							Host:          "tiswork.tisserv.net",
							Port:          5432,
							DbName:        "market",
							DefaultSchema: "public",
						},
					},
				},
			},
		})

		if err != nil {
			t.Error(err)
			return
		}

		log.Print(res)
	})
}

func TestCreateAndReadDataSource(t *testing.T) {
	withClient(func(container *SimpleAppGrpcContainer) {
		var dataSource = &model.DataSource{
			Backend: model.DataSourceBackend_POSTGRESQL,
			Options: &model.DataSource_PostgresqlParams{
				PostgresqlParams: &model.PostgresqlOptions{
					Username:      "root",
					Password:      "52fa536f0c5b85f9d806633937f06446",
					Host:          "tiswork.tisserv.net",
					Port:          5432,
					DbName:        "market",
					DefaultSchema: "public",
				},
			},
		}

		res1, err := container.dataSourceService.Create(context.TODO(), &stub.CreateDataSourceRequest{
			Token: "test-token",
			DataSources: []*model.DataSource{
				dataSource,
			},
		})

		if err != nil {
			t.Error(err)
			return
		}

		if len(res1.DataSources) != 1 {
			t.Error("data-sources count must be 1")
			return
		}

		res2, err := container.dataSourceService.Get(context.TODO(), &stub.GetDataSourceRequest{
			Token: "test-token",
			Id:    res1.DataSources[0].Id,
		})

		if err != nil {
			t.Error(err)
			return
		}

		if res2.DataSource == nil {
			t.Error("Data source must not be null")
			return
		}

		if !reflect.DeepEqual(res1.DataSources[0], res2.DataSource) {
			log.Println(res1.DataSources[0])
			log.Println(res2.DataSource)
			t.Error("Backend is different")
			return
		}
	})
}
