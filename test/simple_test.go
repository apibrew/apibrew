package test

import (
	"context"
	"data-handler/stub"
	"data-handler/stub/model"
	log "github.com/sirupsen/logrus"
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
		}

		log.Print(res)
	})
}
