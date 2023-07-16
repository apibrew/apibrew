package test

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resource_model"
	"github.com/apibrew/apibrew/pkg/resources"
	"github.com/apibrew/apibrew/pkg/stub"
	"github.com/apibrew/apibrew/pkg/test/setup"
	"github.com/google/uuid"
	"testing"
)

func TestCreateAndReadDataSource(t *testing.T) {
	res2, err := recordClient.Get(setup.Ctx, &stub.GetRecordRequest{
		Namespace: resources.DataSourceResource.Namespace,
		Resource:  resources.DataSourceResource.Name,
		Id:        setup.DataSource1.Id.String(),
	})

	if err != nil {
		t.Error(err)
		return
	}

	if res2.Record == nil {
		t.Error("Data source must not be null")
		return
	}

	DeepEqual(t, resource_model.DataSourceMapperInstance.ToRecord(setup.DataSource1), res2.Record, "")
}

func TestCreateRecordstatusTest(t *testing.T) {
	newDataSource := &resource_model.DataSource{
		Backend:     setup.SystemDataSource.Backend,
		Name:        "test-data-source",
		Description: "test-data-source",
		Options:     setup.SystemDataSource.Options,
	}

	defer func() {
		if newDataSource.Id != nil {
			_, err := recordClient.Delete(setup.Ctx, &stub.DeleteRecordRequest{
				Namespace: resources.DataSourceResource.Namespace,
				Resource:  resources.DataSourceResource.Name,
				Ids:       []string{newDataSource.Id.String()},
			})

			if err != nil {
				t.Error(err)
				return
			}
		}
	}()

	resp, err := recordClient.Create(setup.Ctx, &stub.CreateRecordRequest{
		Records: []*model.Record{resource_model.DataSourceMapperInstance.ToRecord(newDataSource)},
	})

	if err != nil {
		t.Error(err)
		return
	}

	newDataSource.Id = new(uuid.UUID)
	*newDataSource.Id = uuid.MustParse(resp.Records[0].Id)

	checkNewCreatedRecordStatus(newDataSource, t)
}

func TestCreateDataSourceWithWrongPasswordStatusTest(t *testing.T) {

	newDataSource := &resource_model.DataSource{
		Backend:     setup.SystemDataSource.Backend,
		Name:        "test-data-source",
		Description: "test-data-source",
		Options:     setup.DhTestWrongPassword.Options,
	}

	defer func() {
		if newDataSource.Id != nil {
			_, err := recordClient.Delete(setup.Ctx, &stub.DeleteRecordRequest{
				Namespace: resources.DataSourceResource.Namespace,
				Resource:  resources.DataSourceResource.Name,
				Ids:       []string{newDataSource.Id.String()},
			})

			if err != nil {
				t.Error(err)
				return
			}
		}
	}()

	resp, err := recordClient.Create(setup.Ctx, &stub.CreateRecordRequest{
		Namespace: resources.DataSourceResource.Namespace,
		Resource:  resources.DataSourceResource.Name,
		Records:   []*model.Record{resource_model.DataSourceMapperInstance.ToRecord(newDataSource)},
	})

	if err != nil {
		t.Error(err)
		return
	}

	newDataSource.Id = new(uuid.UUID)
	*newDataSource.Id = uuid.MustParse(resp.Records[0].Id)

	checkNewCreatedRecordStatusPasswordWrong(newDataSource, t)
}

func TestListCreatedRecords(t *testing.T) {

	res, err := recordClient.List(setup.Ctx, &stub.ListRecordRequest{

		Namespace: resources.DataSourceResource.Namespace,
		Resource:  resources.DataSourceResource.Name,
	})

	if err != nil {
		t.Error(err)
		return
	}

	if len(res.Content) < 3 {
		t.Error("DataSourceList does not match: ", len(res.Content), 3)
	}
}

func TestUpdateDataSource(t *testing.T) {

	newDataSource := &resource_model.DataSource{
		Backend:     setup.SystemDataSource.Backend,
		Name:        "test-data-source",
		Description: "test-data-source",
		Options:     setup.DataSource1.Options,
		Version:     1,
	}

	defer func() {
		if newDataSource.Id != nil {
			_, err := recordClient.Delete(setup.Ctx, &stub.DeleteRecordRequest{
				Ids: []string{newDataSource.Id.String()},
			})

			if err != nil {
				t.Error(err)
				return
			}
		}
	}()

	resp, err := recordClient.Create(setup.Ctx, &stub.CreateRecordRequest{
		Records: []*model.Record{resource_model.DataSourceMapperInstance.ToRecord(newDataSource)},
	})

	if err != nil {
		t.Error(err)
		return
	}

	newDataSource.Id = new(uuid.UUID)
	*newDataSource.Id = uuid.MustParse(resp.Records[0].Id)

	checkNewCreatedRecordStatus(newDataSource, t)

	newDataSource.Options = map[string]string{
		"Username":      "dhtest2",
		"Password":      "dhtest2",
		"Host":          "127.0.0.1",
		"Port":          "5432",
		"DbName":        "market",
		"DefaultSchema": "public",
	}

	res, err := recordClient.Update(setup.Ctx, &stub.UpdateRecordRequest{
		Records: []*model.Record{
			resource_model.DataSourceMapperInstance.ToRecord(newDataSource),
		},
	})

	if err != nil {
		t.Error(err)
		return
	}

	if len(res.Records) != 1 {
		t.Error("Invalid datasource length on update response", len(res.Records))
	}

	updatedParams := resource_model.DataSourceMapperInstance.FromRecord(res.Records[0])

	if updatedParams.Options["Username"] != "dhtest2" {
		t.Error("Username is not updated")
	}

	if updatedParams.Options["Host"] != "127.0.0.1" {
		t.Error("Host is corrupted")
	}

	if res.Records[0].Properties["Version"].GetNumberValue() != 2 {
		t.Error("Version is wrong")
	}

	getRes, err := recordClient.Get(setup.Ctx, &stub.GetRecordRequest{
		Id: newDataSource.Id.String(),
	})

	if err != nil {
		t.Error(err)
	}

	getParams := resource_model.DataSourceMapperInstance.FromRecord(getRes.Record).Options

	if getParams["Username"] != "dhtest2" {
		t.Error("Username is not updated")
	}

	if getParams["Host"] != "127.0.0.1" {
		t.Error("Host is corrupted")
	}

	if getRes.Record.Properties["Version"].GetNumberValue() != 2 {
		t.Error("Version is wrong")
	}

	checkNewCreatedRecordStatusPasswordWrong(resource_model.DataSourceMapperInstance.FromRecord(getRes.Record), t)
}

func TestUpdateRecordstatus(t *testing.T) {

	newDataSource := &resource_model.DataSource{
		Backend:     setup.SystemDataSource.Backend,
		Name:        "test-data-source",
		Description: "test-data-source",
		Options: map[string]string{
			"Username":      "dh_test2",
			"Password":      "dh_test",
			"Host":          "127.0.0.1",
			"Port":          "5432",
			"DbName":        "dh_test",
			"DefaultSchema": "public",
		},
		Version: 1,
	}

	defer func() {
		if newDataSource.Id != nil {
			_, err := recordClient.Delete(setup.Ctx, &stub.DeleteRecordRequest{
				Ids: []string{newDataSource.Id.String()},
			})

			if err != nil {
				t.Error(err)
				return
			}
		}
	}()

	resp, err := recordClient.Create(setup.Ctx, &stub.CreateRecordRequest{
		Records: []*model.Record{resource_model.DataSourceMapperInstance.ToRecord(newDataSource)},
	})

	if err != nil {
		t.Error(err)
		return
	}

	newDataSource.Id = new(uuid.UUID)
	*newDataSource.Id = uuid.MustParse(resp.Records[0].Id)
	createdDataSource1 := resource_model.DataSourceMapperInstance.FromRecord(resp.Records[0])

	checkNewCreatedRecordStatusPasswordWrong(newDataSource, t)

	createdDataSource1.Options = map[string]string{
		"Username":      "dh_test2",
		"Password":      "dh_test",
		"Host":          "127.0.0.1",
		"Port":          "5432",
		"DbName":        "dh_test",
		"DefaultSchema": "public",
	}

	_, _ = recordClient.Update(setup.Ctx, &stub.UpdateRecordRequest{
		Records: []*model.Record{resource_model.DataSourceMapperInstance.ToRecord(createdDataSource1)},
	})

	checkNewCreatedRecordStatusPasswordWrong(createdDataSource1, t)

	createdDataSource1.Options = map[string]string{
		"Username":      "dh_test",
		"Password":      "dh_test",
		"Host":          "127.0.0.1",
		"Port":          "5432",
		"DbName":        "dh_test",
		"DefaultSchema": "public",
	}
	createdDataSource1.Version++

	_, err = recordClient.Update(setup.Ctx, &stub.UpdateRecordRequest{
		Records: []*model.Record{resource_model.DataSourceMapperInstance.ToRecord(createdDataSource1)},
	})

	if err != nil {
		t.Error(err)
		return
	}

	checkNewCreatedRecordStatus(createdDataSource1, t)
}

func checkNewCreatedRecordStatus(createdDataSource *resource_model.DataSource, t *testing.T) {

	res, err := dataSourceClient.Status(setup.Ctx, &stub.StatusRequest{
		Id: createdDataSource.Id.String(),
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

func checkNewCreatedRecordStatusPasswordWrong(createdDataSource *resource_model.DataSource, t *testing.T) {
	resp, err := dataSourceClient.Status(setup.Ctx, &stub.StatusRequest{
		Id: createdDataSource.Id.String(),
	})

	if err == nil {
		t.Error("It should be unable to login to database", resp)
		return
	}
}
