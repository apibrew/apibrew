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

var umo = protojson.UnmarshalOptions{
	AllowPartial: true,
}

type PropertyValueWrapper struct {
	Value *structpb.Value
}

func (pvw *PropertyValueWrapper) MarshalJSON() ([]byte, error) {
	if pvw.Value == nil {
		return []byte("null"), nil
	}
	data, err := mo.Marshal(pvw.Value)

	if err != nil {
		return nil, err
	}

	return data, err
}

func (pvw *PropertyValueWrapper) UnmarshalJSON(data []byte) error {
	pvw.Value = new(structpb.Value)
	return umo.Unmarshal(data, pvw.Value)
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

func (rw *RecordWrapper) toRecord() *model.Record {
	record := new(model.Record)
	record.Properties = make(map[string]*structpb.Value)

	for key, value := range rw.properties {
		if value == nil {
			record.Properties[key] = nil
		} else {
			record.Properties[key] = value.Value
		}
	}

	return record
}

func NewEmptyRecordWrapper() *RecordWrapper {
	var rw = new(RecordWrapper)

	rw.properties = make(map[string]*PropertyValueWrapper)

	return rw
}
func NewRecordWrapper(record *model.Record) *RecordWrapper {
	if record == nil {
		return nil
	}

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

type BooleanExpressionWrapper struct {
	expr *model.BooleanExpression
}

func (pvw *BooleanExpressionWrapper) MarshalJSON() ([]byte, error) {
	return mo.Marshal(pvw.expr)
}

func (pvw *BooleanExpressionWrapper) UnmarshalJSON(data []byte) error {
	pvw.expr = new(model.BooleanExpression)
	return umo.Unmarshal(data, pvw.expr)
}

type SearchRecordRequest struct {
	Query             BooleanExpressionWrapper `json:"query,omitempty"`
	Limit             uint32                   `json:"limit,omitempty"`
	Offset            uint64                   `json:"offset,omitempty"`
	UseHistory        bool                     `json:"useHistory,omitempty"`
	ResolveReferences []string                 `json:"resolveReferences,omitempty"`
	Annotations       map[string]string        `json:"annotations,omitempty"`
}

type HealthResponse struct {
	Status string `json:"status"`
}
