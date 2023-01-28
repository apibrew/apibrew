package test

import (
	"data-handler/model"
	"data-handler/server/stub"
	"data-handler/server/util"
	"data-handler/service/types"
	"google.golang.org/protobuf/types/known/structpb"
	"testing"
)

func TestComplexPayload1Fail(t *testing.T) {

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

	if len(errorFields) != 13 {
		t.Error("There must be 14 error field")
	}
}

func TestComplexPayload1Success(t *testing.T) {

	record1 := new(model.Record)
	record1.Resource = richResource1.Name
	st, err := structpb.NewStruct(map[string]interface{}{
		"bool":   true,
		"bytes":  "YXNk",
		"date":   "2006-01-02",
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
	if len(record1.Properties) != len(getRes.Record.Properties) {
		t.Error("created and get records has different property count")
		return
	}

	for _, property := range richResource1.Properties {
		propertyType := types.ByResourcePropertyType(property.Type)
		val1, _ := propertyType.UnPack(record1.Properties[property.Name])
		val2, _ := propertyType.UnPack(getRes.Record.Properties[property.Name])

		if !propertyType.Equals(val1, val2) {
			t.Errorf("created and get records has different values: %v <=> %v", val1, val2)
			return
		}
	}
}
