package test

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/stub"
	"github.com/apibrew/apibrew/pkg/test/setup"
	"testing"
)

func TestCreateAndReadDataSource(t *testing.T) {
	res2, err := dataSourceClient.Get(setup.Ctx, &stub.GetDataSourceRequest{
		Id: setup.DataSource1.Id,
	})

	if err != nil {
		t.Error(err)
		return
	}

	if res2.DataSource == nil {
		t.Error("Data source must not be null")
		return
	}

	setup.DataSource1.AuditData = res2.DataSource.AuditData

	DeepEqual(t, setup.DataSource1, res2.DataSource, "")
}

func TestCreateDataSourceStatusTest(t *testing.T) {
	newDataSource := &modelnew.DataSource{
		Backend:     setup.SystemDataSource.Backend,
		Name:        "test-data-source",
		Description: "test-data-source",
		Params:      setup.SystemDataSource.Params,
	}

	defer func() {
		if newDataSource.Id != "" {
			_, err := dataSourceClient.Delete(setup.Ctx, &stub.DeleteDataSourceRequest{
				Ids: []string{newDataSource.Id},
			})

			if err != nil {
				t.Error(err)
				return
			}
		}
	}()

	resp, err := dataSourceClient.Create(setup.Ctx, &stub.CreateDataSourceRequest{
		DataSources: []*modelnew.DataSource{newDataSource},
	})

	if err != nil {
		t.Error(err)
		return
	}

	newDataSource.Id = resp.DataSources[0].Id

	checkNewCreatedDatasourceStatus(newDataSource, t)
}

func TestCreateDataSourceWithWrongPasswordStatusTest(t *testing.T) {

	newDataSource := &modelnew.DataSource{
		Backend:     setup.SystemDataSource.Backend,
		Name:        "test-data-source",
		Description: "test-data-source",
		Params:      setup.DhTestWrongPassword.Params,
	}

	defer func() {
		if newDataSource.Id != "" {
			_, err := dataSourceClient.Delete(setup.Ctx, &stub.DeleteDataSourceRequest{
				Ids: []string{newDataSource.Id},
			})

			if err != nil {
				t.Error(err)
				return
			}
		}
	}()

	resp, err := dataSourceClient.Create(setup.Ctx, &stub.CreateDataSourceRequest{
		DataSources: []*modelnew.DataSource{newDataSource},
	})

	if err != nil {
		t.Error(err)
		return
	}

	newDataSource.Id = resp.DataSources[0].Id

	checkNewCreatedDatasourceStatusPasswordWrong(newDataSource, t)
}

func TestListCreatedDataSources(t *testing.T) {

	res, err := dataSourceClient.List(setup.Ctx, &stub.ListDataSourceRequest{})

	if err != nil {
		t.Error(err)
		return
	}

	if len(res.Content) < 3 {
		t.Error("DataSourceList does not match: ", len(res.Content), 3)
	}
}

func TestUpdateDataSource(t *testing.T) {

	newDataSource := &modelnew.DataSource{
		Backend:     setup.SystemDataSource.Backend,
		Name:        "test-data-source",
		Description: "test-data-source",
		Params:      setup.DataSource1.Params,
		Version:     1,
	}

	defer func() {
		if newDataSource.Id != "" {
			_, err := dataSourceClient.Delete(setup.Ctx, &stub.DeleteDataSourceRequest{
				Ids: []string{newDataSource.Id},
			})

			if err != nil {
				t.Error(err)
				return
			}
		}
	}()

	resp, err := dataSourceClient.Create(setup.Ctx, &stub.CreateDataSourceRequest{
		DataSources: []*modelnew.DataSource{newDataSource},
	})

	if err != nil {
		t.Error(err)
		return
	}

	newDataSource.Id = resp.DataSources[0].Id

	checkNewCreatedDatasourceStatus(newDataSource, t)

	newDataSource.Params = &modelnew.DataSource_PostgresqlParams{
		PostgresqlParams: &model.PostgresqlParams{
			Username:      "dhtest2",
			Password:      "dhtest2",
			Host:          "127.0.0.1",
			Port:          5432,
			DbName:        "market",
			DefaultSchema: "public",
		},
	}

	res, err := dataSourceClient.Update(setup.Ctx, &stub.UpdateDataSourceRequest{
		DataSources: []*modelnew.DataSource{newDataSource},
	})

	if err != nil {
		t.Error(err)
		return
	}

	if len(res.DataSources) != 1 {
		t.Error("Invalid datasource length on update response", len(res.DataSources))
	}

	updatedParams := res.DataSources[0].Params.(*modelnew.DataSource_PostgresqlParams)

	if updatedParams.PostgresqlParams.Username != "dhtest2" {
		t.Error("Username is not updated")
	}

	if updatedParams.PostgresqlParams.Host != "127.0.0.1" {
		t.Error("Host is corrupted")
	}

	if res.DataSources[0].Version != 2 {
		t.Error("Version is wrong")
	}

	getRes, err := dataSourceClient.Get(setup.Ctx, &stub.GetDataSourceRequest{
		Id: newDataSource.Id,
	})

	if err != nil {
		t.Error(err)
	}

	getParams := getRes.DataSource.Params.(*modelnew.DataSource_PostgresqlParams)

	if getParams.PostgresqlParams.Username != "dhtest2" {
		t.Error("Username is not updated")
	}

	if getParams.PostgresqlParams.Host != "127.0.0.1" {
		t.Error("Host is corrupted")
	}

	if getRes.DataSource.Version != 2 {
		t.Error("Version is wrong")
	}

	checkNewCreatedDatasourceStatusPasswordWrong(getRes.DataSource, t)
}

func TestUpdateDataSourceStatus(t *testing.T) {

	newDataSource := &modelnew.DataSource{
		Backend:     setup.SystemDataSource.Backend,
		Name:        "test-data-source",
		Description: "test-data-source",
		Params: &modelnew.DataSource_PostgresqlParams{
			PostgresqlParams: &model.PostgresqlParams{
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
			_, err := dataSourceClient.Delete(setup.Ctx, &stub.DeleteDataSourceRequest{
				Ids: []string{newDataSource.Id},
			})

			if err != nil {
				t.Error(err)
				return
			}
		}
	}()

	resp, err := dataSourceClient.Create(setup.Ctx, &stub.CreateDataSourceRequest{
		DataSources: []*modelnew.DataSource{newDataSource},
	})

	if err != nil {
		t.Error(err)
		return
	}

	newDataSource.Id = resp.DataSources[0].Id
	createdDataSource1 := resp.DataSources[0]

	checkNewCreatedDatasourceStatusPasswordWrong(newDataSource, t)

	createdDataSource1.Params = &modelnew.DataSource_PostgresqlParams{
		PostgresqlParams: &model.PostgresqlParams{
			Username:      "dh_test2",
			Password:      "dh_test",
			Host:          "127.0.0.1",
			Port:          5432,
			DbName:        "dh_test",
			DefaultSchema: "public",
		},
	}

	_, _ = dataSourceClient.Update(setup.Ctx, &stub.UpdateDataSourceRequest{
		DataSources: []*modelnew.DataSource{createdDataSource1},
	})

	checkNewCreatedDatasourceStatusPasswordWrong(createdDataSource1, t)

	createdDataSource1.Params = &modelnew.DataSource_PostgresqlParams{
		PostgresqlParams: &model.PostgresqlParams{
			Username:      "dh_test",
			Password:      "dh_test",
			Host:          "127.0.0.1",
			Port:          5432,
			DbName:        "dh_test",
			DefaultSchema: "public",
		},
	}
	createdDataSource1.Version++

	_, err = dataSourceClient.Update(setup.Ctx, &stub.UpdateDataSourceRequest{
		DataSources: []*modelnew.DataSource{createdDataSource1},
	})

	if err != nil {
		t.Error(err)
		return
	}

	checkNewCreatedDatasourceStatus(createdDataSource1, t)
}

func checkNewCreatedDatasourceStatus(createdDataSource *modelnew.DataSource, t *testing.T) {

	res, err := dataSourceClient.Status(setup.Ctx, &stub.StatusRequest{
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

func checkNewCreatedDatasourceStatusPasswordWrong(createdDataSource *modelnew.DataSource, t *testing.T) {
	resp, err := dataSourceClient.Status(setup.Ctx, &stub.StatusRequest{
		Id: createdDataSource.Id,
	})

	if err == nil {
		t.Error("It should be unable to login to database", resp)
		return
	}
}
