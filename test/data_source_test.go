package test

import (
	"data-handler/grpc/stub"
	"data-handler/model"
	log "github.com/sirupsen/logrus"
	"reflect"
	"testing"
)

func TestCreateDataSource(t *testing.T) {
	ctx := prepareTextContext()

	withDataSource(ctx, t, container, dataSource1, func(createdDataSource *model.DataSource) {
		// testing is done
	})
}

func TestCreateAndReadDataSource(t *testing.T) {
	ctx := prepareTextContext()

	withDataSource(ctx, t, container, dataSource1, func(createdDataSource *model.DataSource) {
		res2, err := container.dataSourceService.Get(ctx, &stub.GetDataSourceRequest{
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
}

func TestCreateDataSourceStatusTest(t *testing.T) {
	ctx := prepareTextContext()

	withDataSource(ctx, t, container, systemDataSource, func(createdDataSource *model.DataSource) {
		checkNewCreatedDatasourceStatus(createdDataSource, container, t)
	})
}

func TestCreateDataSourceWithWrongPasswordStatusTest(t *testing.T) {
	ctx := prepareTextContext()

	withDataSource(ctx, t, container, dataSource1WrongPassword, func(createdDataSource *model.DataSource) {
		checkNewCreatedDatasourceStatusPasswordWrong(createdDataSource, container, t)
	})
}

func TestListCreatedDataSources(t *testing.T) {
	ctx := prepareTextContext()

	withDataSource(ctx, t, container, dataSource1, func(createdDataSource1 *model.DataSource) {
		withDataSource(ctx, t, container, dataSource1, func(createdDataSource2 *model.DataSource) {
			withDataSource(ctx, t, container, dataSource1, func(createdDataSource3 *model.DataSource) {
				res, err := container.dataSourceService.List(ctx, &stub.ListDataSourceRequest{
					Token: "test-token",
				})

				if err != nil {
					t.Error(err)
					return
				}

				if len(res.Content) < 3 {
					t.Error("DataSourceList does not match: ", len(res.Content), 3)
				}
			})
		})
	})
}

func TestUpdateDataSource(t *testing.T) {
	ctx := prepareTextContext()

	withDataSource(ctx, t, container, dataSource1, func(createdDataSource1 *model.DataSource) {
		createdDataSource1.Options = &model.DataSource_PostgresqlParams{
			PostgresqlParams: &model.PostgresqlOptions{
				Username:      "root2",
				Password:      "52fa536f0c5b85f9d806633937f064462",
				Host:          "tiswork.tisserv.net",
				Port:          5432,
				DbName:        "market",
				DefaultSchema: "public",
			},
		}

		res, err := container.dataSourceService.Update(ctx, &stub.UpdateDataSourceRequest{
			Token:       "test-token",
			DataSources: []*model.DataSource{createdDataSource1},
		})

		if err != nil {
			t.Error(err)
		}

		if len(res.DataSources) != 1 {
			t.Error("Invalid datasource length on update response", len(res.DataSources))
		}

		updatedOptions := res.DataSources[0].Options.(*model.DataSource_PostgresqlParams)

		if updatedOptions.PostgresqlParams.Username != "root2" {
			t.Error("Username is not updated")
		}

		if updatedOptions.PostgresqlParams.Host != "tiswork.tisserv.net" {
			t.Error("Host is corrupted")
		}

		if res.DataSources[0].Version != 2 {
			t.Error("Version is wrong")
		}

		getRes, err := container.dataSourceService.Get(ctx, &stub.GetDataSourceRequest{
			Token: "test-token",
			Id:    createdDataSource1.Id,
		})

		if err != nil {
			t.Error(err)
		}

		getOptions := getRes.DataSource.Options.(*model.DataSource_PostgresqlParams)

		if getOptions.PostgresqlParams.Username != "root2" {
			t.Error("Username is not updated")
		}

		if getOptions.PostgresqlParams.Host != "tiswork.tisserv.net" {
			t.Error("Host is corrupted")
		}

		if getRes.DataSource.Version != 2 {
			t.Error("Version is wrong")
		}

	})
}

func TestUpdateDataSourceStatus(t *testing.T) {
	ctx := prepareTextContext()

	log.Info("Begin test")
	withDataSource(ctx, t, container, dataSourceDhTest, func(createdDataSource1 *model.DataSource) {
		log.Info("Step 1")
		checkNewCreatedDatasourceStatus(createdDataSource1, container, t)
		log.Info("Step 2")

		createdDataSource1.Options = &model.DataSource_PostgresqlParams{
			PostgresqlParams: &model.PostgresqlOptions{
				Username:      "dh_test2",
				Password:      "dh_test",
				Host:          "127.0.0.1",
				Port:          5432,
				DbName:        "dh_test",
				DefaultSchema: "public",
			},
		}
		log.Info("Step 3")

		container.dataSourceService.Update(ctx, &stub.UpdateDataSourceRequest{
			Token:       "test-token",
			DataSources: []*model.DataSource{createdDataSource1},
		})
		log.Info("Step 4")

		checkNewCreatedDatasourceStatusPasswordWrong(createdDataSource1, container, t)

		log.Info("Step 5")

		createdDataSource1.Options = &model.DataSource_PostgresqlParams{
			PostgresqlParams: &model.PostgresqlOptions{
				Username:      "dh_test",
				Password:      "dh_test",
				Host:          "127.0.0.1",
				Port:          5432,
				DbName:        "dh_test",
				DefaultSchema: "public",
			},
		}
		createdDataSource1.Version++

		container.dataSourceService.Update(ctx, &stub.UpdateDataSourceRequest{
			Token:       "test-token",
			DataSources: []*model.DataSource{createdDataSource1},
		})

		log.Info("Step 6")

		checkNewCreatedDatasourceStatus(createdDataSource1, container, t)
		log.Info("Step 7")
	})

	log.Info("End Test")
}

func checkNewCreatedDatasourceStatus(createdDataSource *model.DataSource, container *SimpleAppGrpcContainer, t *testing.T) {
	ctx := prepareTextContext()

	res, err := container.dataSourceService.Status(ctx, &stub.StatusRequest{
		Token: "test-token",
		Id:    createdDataSource.Id,
	})

	if err != nil {
		t.Error(err)
		return

	}

	if res.ConnectionAlreadyInitiated {
		t.Error("New created datasource should have initiated connection")
	}

	if !res.TestConnection {
		t.Error("New created connection should pass test connection")
	}
}

func checkNewCreatedDatasourceStatusPasswordWrong(createdDataSource *model.DataSource, container *SimpleAppGrpcContainer, t *testing.T) {
	ctx := prepareTextContext()

	res, _ := container.dataSourceService.Status(ctx, &stub.StatusRequest{
		Token: "test-token",
		Id:    createdDataSource.Id,
	})

	if res.Error == nil {
		t.Error("It should be unable to login to database")
		return
	}
}
