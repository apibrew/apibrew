package test

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/stub"
	"github.com/apibrew/apibrew/pkg/test/setup"
	"github.com/apibrew/apibrew/pkg/types"
	"github.com/apibrew/apibrew/pkg/util"
	"google.golang.org/protobuf/types/known/structpb"
	"log"
	"strconv"
	"testing"
)

func TestComplexPayload1Fail(t *testing.T) {
	record1 := new(model.Record)

	_, err := recordClient.Create(setup.Ctx, &stub.CreateRecordRequest{
		Resource: setup.RichResource1.Name,
		Records:  []*model.Record{record1},
	})

	if err == nil {
		t.Error("Apply should fail")
		return
	}

	if util.GetErrorCode(err) != model.ErrorCode_RECORD_VALIDATION_ERROR {
		t.Error("Error code should be: " + model.ErrorCode_RECORD_VALIDATION_ERROR.String())
	}

	errorFields := util.GetErrorFields(err)

	if len(errorFields) != 12 {
		t.Error("There must be 12 error field but: " + strconv.Itoa(len(errorFields)))
	}
}

func TestComplexPayload1Success(t *testing.T) {
	record1 := new(model.Record)
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

	res, err := recordClient.Create(setup.Ctx, &stub.CreateRecordRequest{
		Token:    "",
		Resource: setup.RichResource1.Name,
		Records:  []*model.Record{record1},
	})

	if err != nil {
		t.Error(err)
		return
	}

	getRes, err := recordClient.Get(setup.Ctx, &stub.GetRecordRequest{
		Token:    "",
		Resource: setup.RichResource1.Name,
		Id:       res.Records[0].Id,
	})

	if err != nil {
		t.Error(err)
	}

	for _, property := range setup.RichResource1.Properties {
		propertyType := types.ByResourcePropertyType(property.Type)
		val1, _ := propertyType.UnPack(record1.Properties[property.Name])
		val2, _ := propertyType.UnPack(getRes.Record.Properties[property.Name])

		if !propertyType.Equals(val1, val2) {
			t.Errorf("created and get records has different values: %v <=> %v", val1, val2)
			return
		}
	}
}

func TestComplexPayload1Success1(t *testing.T) {

	record1 := new(model.Record)
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

	//  temp

	resp, err := recordClient.List(setup.Ctx, &stub.ListRecordRequest{
		Token:    "",
		Resource: setup.RichResource1.Name,
	})

	log.Print(resp, err)
}
