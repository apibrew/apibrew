package test

import (
	"data-handler/model"
	"data-handler/server/stub"
	log "github.com/sirupsen/logrus"
	"reflect"
	"testing"
)

func TestCreateAndReadDataSource(t *testing.T) {
	res2, err := dataSourceServiceClient.Get(ctx, &stub.GetDataSourceRequest{
		Id: dataSource1.Id,
	})

	if err != nil {
		t.Error(err)
		return
	}

	if res2.DataSource == nil {
		t.Error("Data source must not be null")
		return
	}

	dataSource1.AuditData = res2.DataSource.AuditData

	if !reflect.DeepEqual(dataSource1, res2.DataSource) {
		log.Println(dataSource1)
		log.Println(res2.DataSource)
		t.Error("Backend is different")
		return
	}
}

func TestCreateDataSourceStatusTest(t *testing.T) {

	newDataSource := &model.DataSource{
		Backend:     systemDataSource.Backend,
		Type:        model.DataType_USER,
		Name:        "test-data-source",
		Description: "test-data-source",
		Options:     systemDataSource.Options,
	}

	defer func() {
		if newDataSource.Id != "" {
			_, err := dataSourceServiceClient.Delete(ctx, &stub.DeleteDataSourceRequest{
				Ids: []string{newDataSource.Id},
			})

			if err != nil {
				t.Error(err)
				return
			}
		}
	}()

	resp, err := dataSourceServiceClient.Create(ctx, &stub.CreateDataSourceRequest{
		DataSources: []*model.DataSource{newDataSource},
	})

	if err != nil {
		t.Error(err)
		return
	}

	newDataSource.Id = resp.DataSources[0].Id

	checkNewCreatedDatasourceStatus(newDataSource, t)
}

func TestCreateDataSourceWithWrongPasswordStatusTest(t *testing.T) {

	newDataSource := &model.DataSource{
		Backend:     systemDataSource.Backend,
		Type:        model.DataType_USER,
		Name:        "test-data-source",
		Description: "test-data-source",
		Options:     dhTestWrongPassword.Options,
	}

	defer func() {
		if newDataSource.Id != "" {
			_, err := dataSourceServiceClient.Delete(ctx, &stub.DeleteDataSourceRequest{
				Ids: []string{newDataSource.Id},
			})

			if err != nil {
				t.Error(err)
				return
			}
		}
	}()

	resp, err := dataSourceServiceClient.Create(ctx, &stub.CreateDataSourceRequest{
		DataSources: []*model.DataSource{newDataSource},
	})

	if err != nil {
		t.Error(err)
		return
	}

	newDataSource.Id = resp.DataSources[0].Id

	checkNewCreatedDatasourceStatusPasswordWrong(newDataSource, t)
}

func TestListCreatedDataSources(t *testing.T) {

	res, err := dataSourceServiceClient.List(ctx, &stub.ListDataSourceRequest{})

	if err != nil {
		t.Error(err)
		return
	}

	if len(res.Content) < 3 {
		t.Error("DataSourceList does not match: ", len(res.Content), 3)
	}
}

func TestUpdateDataSource(t *testing.T) {

	newDataSource := &model.DataSource{
		Backend:     systemDataSource.Backend,
		Type:        model.DataType_USER,
		Name:        "test-data-source",
		Description: "test-data-source",
		Options:     dataSource1.Options,
		Version:     1,
	}

	defer func() {
		if newDataSource.Id != "" {
			_, err := dataSourceServiceClient.Delete(ctx, &stub.DeleteDataSourceRequest{
				Ids: []string{newDataSource.Id},
			})

			if err != nil {
				t.Error(err)
				return
			}
		}
	}()

	resp, err := dataSourceServiceClient.Create(ctx, &stub.CreateDataSourceRequest{
		DataSources: []*model.DataSource{newDataSource},
	})

	if err != nil {
		t.Error(err)
		return
	}

	newDataSource.Id = resp.DataSources[0].Id

	checkNewCreatedDatasourceStatus(newDataSource, t)

	newDataSource.Options = &model.DataSource_PostgresqlParams{
		PostgresqlParams: &model.PostgresqlOptions{
			Username:      "dhtest2",
			Password:      "dhtest2",
			Host:          "127.0.0.1",
			Port:          5432,
			DbName:        "market",
			DefaultSchema: "public",
		},
	}

	res, err := dataSourceServiceClient.Update(ctx, &stub.UpdateDataSourceRequest{
		DataSources: []*model.DataSource{newDataSource},
	})

	if err != nil {
		t.Error(err)
	}

	if len(res.DataSources) != 1 {
		t.Error("Invalid datasource length on update response", len(res.DataSources))
	}

	updatedOptions := res.DataSources[0].Options.(*model.DataSource_PostgresqlParams)

	if updatedOptions.PostgresqlParams.Username != "dhtest2" {
		t.Error("Username is not updated")
	}

	if updatedOptions.PostgresqlParams.Host != "127.0.0.1" {
		t.Error("Host is corrupted")
	}

	if res.DataSources[0].Version != 2 {
		t.Error("Version is wrong")
	}

	getRes, err := dataSourceServiceClient.Get(ctx, &stub.GetDataSourceRequest{
		Id: newDataSource.Id,
	})

	if err != nil {
		t.Error(err)
	}

	getOptions := getRes.DataSource.Options.(*model.DataSource_PostgresqlParams)

	if getOptions.PostgresqlParams.Username != "dhtest2" {
		t.Error("Username is not updated")
	}

	if getOptions.PostgresqlParams.Host != "127.0.0.1" {
		t.Error("Host is corrupted")
	}

	if getRes.DataSource.Version != 2 {
		t.Error("Version is wrong")
	}

	checkNewCreatedDatasourceStatusPasswordWrong(getRes.DataSource, t)

}

func TestUpdateDataSourceStatus(t *testing.T) {

	newDataSource := &model.DataSource{
		Backend:     systemDataSource.Backend,
		Type:        model.DataType_USER,
		Name:        "test-data-source",
		Description: "test-data-source",
		Options: &model.DataSource_PostgresqlParams{
			PostgresqlParams: &model.PostgresqlOptions{
				Username:      "dh_test2",
				Password:      "dh_test",
				Host:          "127.0.0.1",
				Port:          5432,
				DbName:        "dh_test",
				DefaultSchema: "public",
			},
		},
		Version: 1,
	}

	defer func() {
		if newDataSource.Id != "" {
			_, err := dataSourceServiceClient.Delete(ctx, &stub.DeleteDataSourceRequest{
				Ids: []string{newDataSource.Id},
			})

			if err != nil {
				t.Error(err)
				return
			}
		}
	}()

	resp, err := dataSourceServiceClient.Create(ctx, &stub.CreateDataSourceRequest{
		DataSources: []*model.DataSource{newDataSource},
	})

	if err != nil {
		t.Error(err)
		return
	}

	newDataSource.Id = resp.DataSources[0].Id
	createdDataSource1 := resp.DataSources[0]

	log.Info("Step 1")
	checkNewCreatedDatasourceStatusPasswordWrong(newDataSource, t)
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

	dataSourceServiceClient.Update(ctx, &stub.UpdateDataSourceRequest{
		DataSources: []*model.DataSource{createdDataSource1},
	})
	log.Info("Step 4")

	checkNewCreatedDatasourceStatusPasswordWrong(createdDataSource1, t)

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

	dataSourceServiceClient.Update(ctx, &stub.UpdateDataSourceRequest{
		DataSources: []*model.DataSource{createdDataSource1},
	})

	log.Info("Step 6")

	checkNewCreatedDatasourceStatus(createdDataSource1, t)
	log.Info("Step 7")
}

func checkNewCreatedDatasourceStatus(createdDataSource *model.DataSource, t *testing.T) {

	res, err := dataSourceServiceClient.Status(ctx, &stub.StatusRequest{
		Id: createdDataSource.Id,
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

func checkNewCreatedDatasourceStatusPasswordWrong(createdDataSource *model.DataSource, t *testing.T) {

	_, err := dataSourceServiceClient.Status(ctx, &stub.StatusRequest{
		Id: createdDataSource.Id,
	})

	if err == nil {
		t.Error("It should be unable to login to database")
		return
	}
}
