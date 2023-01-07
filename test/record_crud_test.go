package test

import (
	"data-handler/grpc/stub"
	"data-handler/model"
	"google.golang.org/protobuf/types/known/structpb"
	"log"
	"testing"
)

func TestComplexPayload1Fail(t *testing.T) {
	ctx := prepareTextContext()

	withDataSource(ctx, t, container, dataSourceDhTest, func(dataSource *model.DataSource) {
		richResource1.SourceConfig.DataSource = dataSource.Id
		withResource(ctx, t, richResource1, func() {

			record1 := new(model.Record)
			record1.Resource = richResource1.Name

			res, err := container.recordService.Create(ctx, &stub.CreateRecordRequest{
				Token:   "",
				Records: []*model.Record{record1},
			})

			if err != nil {
				t.Error(err)
			}

			if res.Error == nil {
				t.Error("Save should fail")
			}

			if res.Error.Code != model.ErrorCode_RECORD_VALIDATION_ERROR {
				t.Error("Error code is wrong: " + res.Error.Code.String())
			}

			if len(res.Error.Fields) != 14 {
				t.Error("error field length should be 15 but is ", len(res.Error.Fields))
			}

			log.Print(res.Error)

		})
	})
}

func TestComplexPayload1Success(t *testing.T) {
	ctx := prepareTextContext()

	withDataSource(ctx, t, container, dataSourceDhTest, func(dataSource *model.DataSource) {
		richResource1.SourceConfig.DataSource = dataSource.Id
		withResource(ctx, t, richResource1, func() {
			record1 := new(model.Record)
			record1.Resource = richResource1.Name
			var err error
			record1.Properties, err = structpb.NewStruct(map[string]interface{}{
				"bool":    true,
				"bytes":   "YXNk",
				"date":    "2006-01-02",
				"double":  12.3,
				"float":   31.200000762939453,
				"int32":   12,
				"int64":   34,
				"numeric": 99,
				"object": map[string]interface{}{
					"test1": "test-123",
				},
				"string":    "asdasdksadjsakldksal",
				"text":      "test1233321",
				"time":      "17:04:05",
				"timestamp": "2006-01-02T15:04:05Z",
				"uuid":      "bdedf5b8-5179-11ed-bdc3-0242ac120002",
			})

			if err != nil {
				t.Error(err)
			}

			res, err := container.recordService.Create(ctx, &stub.CreateRecordRequest{
				Token:   "",
				Records: []*model.Record{record1},
			})

			if err != nil {
				t.Error(err)
				return
			}

			if res.Error != nil {
				t.Error(res.Error)
				return
			}

			getRes, err := container.recordService.Get(ctx, &stub.GetRecordRequest{
				Token:    "",
				Resource: richResource1.Name,
				Id:       res.Records[0].Id,
			})

			if err != nil {
				t.Error(err)
			}

			if getRes.Error != nil {
				t.Error(getRes.Error)
			}

			createJson, err := record1.Properties.MarshalJSON()

			if err != nil {
				t.Error(err)
			}

			getJson, err := getRes.Record.Properties.MarshalJSON()

			if err != nil {
				t.Error(err)
			}

			if string(createJson) != string(getJson) {
				log.Println(string(createJson))
				log.Println(string(getJson))
				t.Error("Created and get records has different properties")
			}
		})
	})
}
