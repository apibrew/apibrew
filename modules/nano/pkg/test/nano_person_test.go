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

func TestNanoPersonPreventDelete(t *testing.T) {
	err := apbrApply("data/person.yml")
	if err != nil {
		t.Error(err)
		return
	}

	record1 := new(model.Record)
	st, err := structpb.NewStruct(map[string]interface{}{
		"firstName": "PreventDelete",
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

	if res.Record != nil {
		t.Error("Record should not be created")
	}
}

func TestPersonBindCreate(t *testing.T) {
	err := apbrApply("data/human.yml")
	if err != nil {
		t.Error(err)
		return
	}

	record1 := new(model.Record)
	st, err := structpb.NewStruct(map[string]interface{}{
		"name": "Taleh Ibrahimli",
	})

	record1.Properties = st.GetFields()

	if err != nil {
		t.Error(err)
		return
	}

	res, err := recordClient.Create(setup.Ctx, &stub.CreateRecordRequest{
		Resource: "Human",
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

	assert.Equal(t, getRes.Record.Properties["firstName"].GetStringValue(), "Taleh")
	assert.Equal(t, getRes.Record.Properties["lastName"].GetStringValue(), "Ibrahimli")
}

func TestPersonBindUpdate(t *testing.T) {
	err := apbrApply("data/human.yml")
	if err != nil {
		t.Error(err)
		return
	}

	record1 := new(model.Record)
	st, err := structpb.NewStruct(map[string]interface{}{
		"name": "Taleh Ibrahimli",
	})

	record1.Properties = st.GetFields()

	if err != nil {
		t.Error(err)
		return
	}

	resp1, err := recordClient.Create(setup.Ctx, &stub.CreateRecordRequest{
		Resource: "Human",
		Records:  []*model.Record{record1},
	})

	if err != nil {
		t.Error(err)
		return
	}

	record2 := new(model.Record)
	st2, err := structpb.NewStruct(map[string]interface{}{
		"id":   util.GetRecordId(resp1.Record),
		"name": "Talehx Ibrahimlix",
	})

	record2.Properties = st2.GetFields()

	if err != nil {
		t.Error(err)
		return
	}

	res2, err := recordClient.Update(setup.Ctx, &stub.UpdateRecordRequest{
		Resource: "Human",
		Records:  []*model.Record{record2},
	})

	if err != nil {
		t.Error(err)
		return
	}

	getRes, err := recordClient.Get(setup.Ctx, &stub.GetRecordRequest{
		Resource: "Person",
		Id:       util.GetRecordId(res2.Records[0]),
	})

	if err != nil {
		t.Error(err)
		return
	}

	assert.Equal(t, getRes.Record.Properties["firstName"].GetStringValue(), "Talehx")
	assert.Equal(t, getRes.Record.Properties["lastName"].GetStringValue(), "Ibrahimlix")
}
