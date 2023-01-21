package test2

import (
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
