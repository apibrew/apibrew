package test2

import (
	"data-handler/model"
	"data-handler/server/stub"
	log "github.com/sirupsen/logrus"
	"reflect"
	"testing"
)

func TestCreateAndReadDataSource(t *testing.T) {
	ctx := prepareTextContext()

	res2, err := dataSourceServiceClient.Get(ctx, &stub.GetDataSourceRequest{
		Token: "test-token",
		Id:    dataSource1.Id,
	})

	if err != nil {
		t.Error(err)
		return
	}

	if res2.DataSource == nil {
		t.Error("Data source must not be null")
		return
	}

	if !reflect.DeepEqual(dataSource1, res2.DataSource) {
		log.Println(dataSource1)
		log.Println(res2.DataSource)
		t.Error("Backend is different")
		return
	}
}

func TestCreateDataSourceStatusTest(t *testing.T) {
	ctx := prepareTextContext()

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
	ctx := prepareTextContext()

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

func checkNewCreatedDatasourceStatus(createdDataSource *model.DataSource, t *testing.T) {
	ctx := prepareTextContext()

	res, err := dataSourceServiceClient.Status(ctx, &stub.StatusRequest{
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

func checkNewCreatedDatasourceStatusPasswordWrong(createdDataSource *model.DataSource, t *testing.T) {
	ctx := prepareTextContext()

	_, err := dataSourceServiceClient.Status(ctx, &stub.StatusRequest{
		Token: "test-token",
		Id:    createdDataSource.Id,
	})

	if err == nil {
		t.Error("It should be unable to login to database")
		return
	}
}
