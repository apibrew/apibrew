package test

import (
	"data-handler/model"
	"data-handler/server/stub"
	"data-handler/server/util"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/structpb"
	"testing"
)

func TestComplexPayload1Fail(t *testing.T) {
	ctx := prepareTextContext()

	record1 := new(model.Record)
	record1.Resource = richResource1.Name

	_, err := recordServiceClient.Create(ctx, &stub.CreateRecordRequest{
		Token:   "",
		Records: []*model.Record{record1},
	})

	if err == nil {
		t.Error("Save should fail")
	}

	if util.GetErrorCode(err) != model.ErrorCode_RECORD_VALIDATION_ERROR {
		t.Error("Error code should be: " + model.ErrorCode_RECORD_VALIDATION_ERROR.String())
	}

	errorFields := util.GetErrorFields(err)

	if len(errorFields) != 14 {
		t.Error("There must be 14 error field")
	}
}

func TestComplexPayload1Success(t *testing.T) {
	ctx := prepareTextContext()

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

	res, err := recordServiceClient.Create(ctx, &stub.CreateRecordRequest{
		Token:   "",
		Records: []*model.Record{record1},
	})

	if err != nil {
		t.Error(err)
		return
	}

	getRes, err := recordServiceClient.Get(ctx, &stub.GetRecordRequest{
		Token:    "",
		Resource: richResource1.Name,
		Id:       res.Records[0].Id,
	})

	if err != nil {
		t.Error(err)
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
}
