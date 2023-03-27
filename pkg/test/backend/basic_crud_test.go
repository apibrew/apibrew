package backend

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/tislib/data-handler/pkg/errors"
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/stub"
	"github.com/tislib/data-handler/pkg/test/setup"
	"github.com/tislib/data-handler/pkg/types"
	"github.com/tislib/data-handler/pkg/util"
	"google.golang.org/protobuf/types/known/structpb"
	"testing"
)

var dataSources = []*model.DataSource{
	setup.DhTest,
	dhTestMysql,
	dhTestRedis,
	dhTestMongo,
	dhTestSqlite,
}

var resources = make(map[*model.DataSource]*model.Resource)

func TestMain(t *testing.M) {
	setup.SetupDataSources(setup.Ctx, dataSources)

	var pendingResources []*model.Resource

	for _, dataSource := range dataSources {
		newRes := setup.PrepareRichResource1()

		newRes.SourceConfig.DataSource = dataSource.Name
		newRes.Name = dataSource.Name + "-" + newRes.Name
		newRes.SourceConfig.Entity = newRes.Name

		pendingResources = append(pendingResources, newRes)
		resources[dataSource] = newRes
	}

	setup.SetupResources(setup.Ctx, pendingResources)

	t.Run()

	setup.DestroyResources(setup.Ctx, pendingResources)
}

func TestCreateRecord(t *testing.T) {
	for _, dataSource := range dataSources {
		t.Run(fmt.Sprintf("%s[%s]", dataSource.Backend.String(), dataSource.Name), func(t *testing.T) {
			record1 := new(model.Record)
			st, err := structpb.NewStruct(map[string]interface{}{
				"bool":   true,
				"bytes":  "YXNk",
				"date":   "2001-01-02",
				"double": 12.3,
				"float":  31.200000762939453,
				"int32":  12,
				"int64":  34,
				"object": map[string]interface{}{ //@todo fixme double packing problem
					"test1": "test-123",
				},
				"string":    "asdasdksadjsakldksal",
				"text":      "test1233321",
				"time":      "17:04:05",
				"timestamp": "2006-01-02T15:04:05Z",
				"uuid":      "bdedf5b8-5179-11ed-bdc3-0242ac120002",
			})

			record1.Properties = st.GetFields()

			if err != nil {
				t.Error(err)
			}

			res, err := setup.GetTestDhClient().GetRecordClient().Create(setup.Ctx, &stub.CreateRecordRequest{
				Token:    "",
				Resource: resources[dataSource].Name,
				Records:  []*model.Record{record1},
			})

			if err != nil {
				t.Error(err)
				return
			}

			getRes, err := setup.GetTestDhClient().GetRecordClient().Get(setup.Ctx, &stub.GetRecordRequest{
				Token:    "",
				Resource: resources[dataSource].Name,
				Id:       res.Records[0].Id,
			})

			if err != nil {
				t.Error(err)
				return
			}

			if !util.IsSameRecord(getRes.Record, record1) {
				t.Error("created and get records has different property count")
				return
			}

			for _, property := range resources[dataSource].Properties {
				propertyType := types.ByResourcePropertyType(property.Type)
				val1, _ := propertyType.UnPack(record1.Properties[property.Name])
				val2, _ := propertyType.UnPack(getRes.Record.Properties[property.Name])

				if !propertyType.Equals(val1, val2) {
					t.Errorf("created and get records has different values: %v <=> %v", val1, val2)
					return
				}
			}
		})
	}
}

func TestDeleteRecord(t *testing.T) {
	for _, dataSource := range dataSources {
		t.Run(fmt.Sprintf("%s[%s]", dataSource.Backend.String(), dataSource.Name), func(t *testing.T) {
			record1 := new(model.Record)
			st, err := structpb.NewStruct(map[string]interface{}{
				"bool":   true,
				"bytes":  "YXNk",
				"date":   "2001-01-02",
				"double": 12.3,
				"float":  31.200000762939453,
				"int32":  12,
				"int64":  34,
				"object": map[string]interface{}{ //@todo fixme double packing problem
					"test1": "test-123",
				},
				"string":    "asdasdksadjsakldksal",
				"text":      "test1233321",
				"time":      "17:04:05",
				"timestamp": "2006-01-02T15:04:05Z",
				"uuid":      "bdedf5b8-5179-11ed-bdc3-0242ac120002",
			})

			record1.Properties = st.GetFields()

			if err != nil {
				t.Error(err)
			}

			res, err := setup.GetTestDhClient().GetRecordClient().Create(setup.Ctx, &stub.CreateRecordRequest{
				Token:    "",
				Resource: resources[dataSource].Name,
				Records:  []*model.Record{record1},
			})

			if err != nil {
				t.Error(err)
				return
			}

			_, err = setup.GetTestDhClient().GetRecordClient().Delete(setup.Ctx, &stub.DeleteRecordRequest{
				Token:    "",
				Resource: resources[dataSource].Name,
				Ids:      []string{res.Records[0].Id},
			})

			if err != nil {
				t.Error(err)
				return
			}

			_, err = setup.GetTestDhClient().GetRecordClient().Get(setup.Ctx, &stub.GetRecordRequest{
				Token:    "",
				Resource: resources[dataSource].Name,
				Id:       res.Records[0].Id,
			})

			assert.Equal(t, util.GetErrorCode(err), model.ErrorCode_RECORD_NOT_FOUND)
		})
	}
}

func TestUpdateRecord(t *testing.T) {
	for _, dataSource := range dataSources {
		t.Run(fmt.Sprintf("%s[%s]", dataSource.Backend.String(), dataSource.Name), func(t *testing.T) {
			record1 := new(model.Record)
			st, err := structpb.NewStruct(map[string]interface{}{
				"bool":   true,
				"bytes":  "YXNk",
				"date":   "2001-01-02",
				"double": 12.3,
				"float":  31.200000762939453,
				"int32":  12,
				"int64":  34,
				"object": map[string]interface{}{ //@todo fixme double packing problem
					"test1": "test-123",
				},
				"string":    "asdasdksadjsakldksal",
				"text":      "test1233321",
				"time":      "17:04:05",
				"timestamp": "2006-01-02T15:04:05Z",
				"uuid":      "bdedf5b8-5179-11ed-bdc3-0242ac120002",
			})

			record1.Properties = st.GetFields()

			if err != nil {
				t.Error(err)
			}

			res, err := setup.GetTestDhClient().GetRecordClient().Create(setup.Ctx, &stub.CreateRecordRequest{
				Token:    "",
				Resource: resources[dataSource].Name,
				Records:  []*model.Record{record1},
			})

			if err != nil {
				t.Error(err)
				return
			}

			record1.Id = res.Records[0].Id

			record1.Properties["string"] = structpb.NewStringValue("Updated Value")

			_, err = setup.GetTestDhClient().GetRecordClient().Update(setup.Ctx, &stub.UpdateRecordRequest{
				Token:    "",
				Resource: resources[dataSource].Name,
				Records:  []*model.Record{record1},
			})

			if err != nil {
				t.Error(err)
				return
			}

			getRes, err := setup.GetTestDhClient().GetRecordClient().Get(setup.Ctx, &stub.GetRecordRequest{
				Token:    "",
				Resource: resources[dataSource].Name,
				Id:       res.Records[0].Id,
			})

			if err != nil {
				t.Error(err)
				return
			}

			if !util.IsSameRecord(getRes.Record, record1) {
				t.Error("created and get records has different property count")
				return
			}

			for _, property := range resources[dataSource].Properties {
				propertyType := types.ByResourcePropertyType(property.Type)
				val1, _ := propertyType.UnPack(record1.Properties[property.Name])
				val2, _ := propertyType.UnPack(getRes.Record.Properties[property.Name])

				if !propertyType.Equals(val1, val2) {
					t.Errorf("created and get records has different values: %v <=> %v", val1, val2)
					return
				}
			}
		})
	}
}

func TestQueryRecord(t *testing.T) {
	for _, dataSource := range dataSources {
		t.Run(fmt.Sprintf("%s[%s]", dataSource.Backend.String(), dataSource.Name), func(t *testing.T) {
			var list = []map[string]interface{}{
				{
					"bool":   true,
					"bytes":  "YXNk",
					"date":   "2001-01-02",
					"double": 12.3,
					"float":  31.200000762939453,
					"int32":  12,
					"int64":  34,
					"object": map[string]interface{}{ //@todo fixme double packing problem
						"test1": "test-123",
					},
					"string":    "param-1",
					"text":      "test1233321",
					"time":      "17:04:05",
					"timestamp": "2006-01-02T15:04:05Z",
					"uuid":      "bdedf5b8-5179-11ed-bdc3-0242ac120002",
				},
				{
					"bool":   true,
					"bytes":  "YXNk",
					"date":   "2001-01-02",
					"double": 12.3,
					"float":  31.200000762939453,
					"int32":  12,
					"int64":  34,
					"object": map[string]interface{}{ //@todo fixme double packing problem
						"test1": "test-123",
					},
					"string":    "param-1",
					"text":      "test1233321",
					"time":      "17:04:05",
					"timestamp": "2006-01-02T15:04:05Z",
					"uuid":      "bdedf5b8-5179-11ed-bdc3-0242ac120002",
				},
				{
					"bool":   true,
					"bytes":  "YXNk",
					"date":   "2001-01-02",
					"double": 12.3,
					"float":  31.200000762939453,
					"int32":  12,
					"int64":  34,
					"object": map[string]interface{}{ //@todo fixme double packing problem
						"test1": "test-123",
					},
					"string":    "param-2",
					"text":      "test1233321",
					"time":      "17:04:05",
					"timestamp": "2006-01-02T15:04:05Z",
					"uuid":      "bdedf5b8-5179-11ed-bdc3-0242ac120002",
				},
				{
					"bool":   true,
					"bytes":  "YXNk",
					"date":   "2001-01-02",
					"double": 12.3,
					"float":  31.200000762939453,
					"int32":  121,
					"int64":  34,
					"object": map[string]interface{}{ //@todo fixme double packing problem
						"test1": "test-123",
					},
					"string":    "param-1",
					"text":      "test1233321",
					"time":      "17:04:05",
					"timestamp": "2006-01-02T15:04:05Z",
					"uuid":      "bdedf5b8-5179-11ed-bdc3-0242ac120002",
				},
			}
			var records []*model.Record

			for _, item := range list {
				var st, err = structpb.NewStruct(item)

				if err != nil {
					t.Error(err)
					return
				}

				records = append(records, &model.Record{
					Properties: st.GetFields(),
				})
			}

			_, err := setup.GetTestDhClient().GetRecordClient().Create(setup.Ctx, &stub.CreateRecordRequest{
				Resource: resources[dataSource].Name,
				Records:  records,
			})

			if err != nil {
				t.Error(err)
			}

			res, err := setup.GetTestDhClient().GetRecordClient().List(setup.Ctx, &stub.ListRecordRequest{
				Resource: resources[dataSource].Name,
				Filters: map[string]string{
					"string": "param-1",
					"int32":  "12",
				},
			})

			if errors.UnsupportedOperation.Is(err) {
				t.SkipNow()
				return
			}

			if err != nil {
				t.Error(err)
				return
			}

			assert.Equal(t, int(res.Total), 2)
			assert.Len(t, res.Content, 2)
		})
	}
}
