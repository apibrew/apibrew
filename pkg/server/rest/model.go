package rest

import (
	"encoding/json"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resource_model"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/structpb"
	"time"
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

type SearchRecordRequest struct {
	Query             *resource_model.BooleanExpression `json:"query,omitempty"`
	Limit             uint32                            `json:"limit,omitempty"`
	Offset            uint64                            `json:"offset,omitempty"`
	UseHistory        bool                              `json:"useHistory,omitempty"`
	ResolveReferences []string                          `json:"resolveReferences,omitempty"`
	Annotations       map[string]string                 `json:"annotations,omitempty"`
	Filters           map[string]string                 `json:"filters,omitempty"`
}

type HealthResponse struct {
	Status string `json:"status"`
}

type AuthenticationRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Term     string `json:"term"`
}

type Token struct {
	Term       string `json:"term"`
	Content    string `json:"content"`
	Expiration time.Time
}

type AuthenticationResponse struct {
	Token Token `json:"token"`
}

type RefreshTokenRequest struct {
	Token string `json:"token"`
	Term  string `json:"term"`
}

type RefreshTokenResponse struct {
	Token Token `json:"token"`
}
