package test

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/stub"
	"github.com/apibrew/apibrew/pkg/test/setup"
	"github.com/apibrew/apibrew/pkg/util"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/structpb"
	"testing"
)

func TestNanoPersonCreate(t *testing.T) {
	err := apbrApply("data/person.yml")
	if err != nil {
		t.Error(err)
		return
	}

	record1 := new(model.Record)
	st, err := structpb.NewStruct(map[string]interface{}{
		"firstName": "Taleh",
	})

	record1.Properties = st.GetFields()

	if err != nil {
		t.Error(err)
		return
	}

	res, err := recordClient.Create(setup.Ctx, &stub.CreateRecordRequest{
		Resource: "Person",
		Records:  []*model.Record{record1},
	})

	if err != nil {
		t.Error(err)
		return
	}

	getRes, err := recordClient.Get(setup.Ctx, &stub.GetRecordRequest{
		Resource: "Person",
		Id:       util.GetRecordId(res.Records[0]),
	})

	if err != nil {
		t.Error(err)
		return
	}

	assert.Equal(t, getRes.Record.Properties["lastName"].GetStringValue(), "Unknown")
}
