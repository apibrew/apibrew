package test

import (
	"github.com/apibrew/apibrew/modules/nano/pkg/model"
	model2 "github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/stub"
	"github.com/apibrew/apibrew/pkg/test/setup"
	"github.com/apibrew/apibrew/pkg/util"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/structpb"
	"testing"
)

func TestTypescriptNanoScriptBasic(t *testing.T) {
	result := runScriptWithLanguage(t, `const a: number = 5; const b: number = 6; a + b;`, model.ScriptLanguage_TYPESCRIPT)
	if t.Failed() {
		return
	}

	assert.NotNil(t, result["output"])

	if t.Failed() {
		return
	}

	output := result["output"]

	assert.Equal(t, float64(11), output)
}

func TestTypescriptNanoPersonCreate(t *testing.T) {
	err := apbrApply("data/person-typescript.yml")
	if err != nil {
		t.Error(err)
		return
	}

	record1 := new(model2.Record)
	st, err := structpb.NewStruct(map[string]interface{}{
		"firstName": "Taleh",
	})

	record1.Properties = st.GetFields()

	if err != nil {
		t.Error(err)
		return
	}

	res, err := recordClient.Create(setup.Ctx, &stub.CreateRecordRequest{
		Resource: "PersonTypescript",
		Records:  []*model2.Record{record1},
	})

	if err != nil {
		t.Error(err)
		return
	}

	getRes, err := recordClient.Get(setup.Ctx, &stub.GetRecordRequest{
		Resource: "PersonTypescript",
		Id:       util.GetRecordId(res.Records[0]),
	})

	if err != nil {
		t.Error(err)
		return
	}

	assert.Equal(t, getRes.Record.Properties["lastName"].GetStringValue(), "Unknown")
}
