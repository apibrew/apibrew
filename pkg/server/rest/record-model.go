package rest

import (
	"encoding/json"
	"github.com/apibrew/apibrew/pkg/model"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/structpb"
)

var mo = protojson.MarshalOptions{
	Multiline:     true,
	Indent:        "  ",
	AllowPartial:  true,
	UseProtoNames: true,
}

type PropertyValueWrapper struct {
	Value *structpb.Value
}

func (pvw *PropertyValueWrapper) MarshalJSON() ([]byte, error) {
	return mo.Marshal(pvw.Value)
}

func (pvw *PropertyValueWrapper) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &pvw.Value)
}

type RecordWrapper struct {
	properties map[string]*PropertyValueWrapper
}

func (rw *RecordWrapper) MarshalJSON() ([]byte, error) {
	return json.Marshal(rw.properties)
}

func (rw *RecordWrapper) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &rw.properties)
}

func NewRecordWrapper(record *model.Record) *RecordWrapper {
	var rw = new(RecordWrapper)

	rw.properties = make(map[string]*PropertyValueWrapper)

	for k, v := range record.Properties {
		rw.properties[k] = &PropertyValueWrapper{Value: v}
	}

	return rw
}

type RecordList struct {
	Total   uint64           `json:"total"`
	Records []*RecordWrapper `json:"content"`
}
