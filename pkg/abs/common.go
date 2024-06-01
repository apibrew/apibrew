package abs

import (
	"github.com/apibrew/apibrew/pkg/model"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/structpb"
)

type ResourceLike interface {
	GetProperties() []*model.ResourceProperty
	GetTypes() []*model.ResourceSubType
}

type RecordLike interface {
	Keys() []string
	GetStructProperty(key string) *structpb.Value
	SetStructProperty(key string, value *structpb.Value) RecordLike
	WithProperties(properties map[string]interface{}) RecordLike
	WithStructProperties(properties map[string]*structpb.Value) RecordLike
	WithStringProperty(key string, value string) RecordLike
	EqualTo(updated RecordLike) bool
	HasProperty(key string) bool
	GetStringProperty(s string) string
	ToStruct() *structpb.Struct
	DeleteProperty(s string)
	Merge(appliedRecord RecordLike) RecordLike
	AsInterface(key string) interface{}
	Self() map[string]interface{}
}

type record map[string]interface{}

func (r *record) Self() map[string]interface{} {
	return *r
}

func (r *record) WithProperties(properties map[string]interface{}) RecordLike {
	for key, value := range properties {
		(*r)[key] = value
	}

	return r
}

func (r *record) AsInterface(key string) interface{} {
	return (*r)[key]
}

func (r *record) Merge(appliedRecord RecordLike) RecordLike {
	for _, key := range appliedRecord.Keys() {
		r.SetStructProperty(key, appliedRecord.GetStructProperty(key))
	}

	return r
}

func (r *record) DeleteProperty(s string) {
	delete(*r, s)
}

func (r *record) ToStruct() *structpb.Struct {
	var result = make(map[string]*structpb.Value)

	for key := range *r {
		result[key] = r.GetStructProperty(key)
	}

	return &structpb.Struct{Fields: result}
}

func (r *record) GetStringProperty(s string) string {
	val, ok := (*r)[s].(string)

	if !ok {
		return ""
	}

	return val
}

func (r *record) HasProperty(key string) bool {
	_, ok := (*r)[key]

	return ok
}

func (r *record) EqualTo(updated RecordLike) bool {
	if updated == nil {
		return false
	}

	if len(*r) != len(updated.Keys()) {
		return false
	}

	for key, value := range *r {
		if updated.GetStructProperty(key) == nil {
			return false
		}

		if value != updated.GetStructProperty(key) {
			if !proto.Equal(updated.GetStructProperty(key), r.GetStructProperty(key)) {
				return false
			}
		}
	}

	return true
}

func (r *record) Keys() []string {
	var result = make([]string, 0, len(*r))

	for key := range *r {
		result = append(result, key)
	}

	return result
}

func (r *record) GetStructProperty(key string) *structpb.Value {
	value, ok := (*r)[key]

	if !ok {
		return nil
	}

	if value == nil {
		return nil
	}

	stValue, err := structpb.NewValue(value)

	if err != nil {
		panic(err)
	}

	return stValue
}

func (r *record) SetStructProperty(key string, value *structpb.Value) RecordLike {
	(*r)[key] = value.AsInterface()

	return r
}

func (r *record) WithStructProperties(properties map[string]*structpb.Value) RecordLike {
	for key, value := range properties {
		r.SetStructProperty(key, value)
	}

	return r
}

func (r *record) WithStringProperty(key string, value string) RecordLike {
	(*r)[key] = value

	return r
}

func RecordLikeAsRecord(record RecordLike) *model.Record {
	if record == nil {
		return nil
	}

	var properties = make(map[string]*structpb.Value)

	for _, key := range record.Keys() {
		properties[key] = record.GetStructProperty(key)
	}

	return &model.Record{
		Properties: properties,
	}
}

func RecordAsRecordLike(record *model.Record) RecordLike {
	if record == nil {
		return nil
	}

	return NewRecordLikeWithProperties(record.Properties)
}

func RecordLikeAsRecords(record []RecordLike) []*model.Record {
	records := make([]*model.Record, 0, len(record))

	for _, r := range record {
		records = append(records, RecordLikeAsRecord(r))
	}

	return records
}

func RecordLikeAsRecords2(record []*model.Record) []RecordLike {
	records := make([]RecordLike, 0, len(record))

	for _, r := range record {
		records = append(records, NewRecordLike().WithStructProperties(r.Properties))
	}

	return records
}

func UpdateRecordsProperties(record RecordLike, properties map[string]*structpb.Value) {
	for key, value := range properties {
		record.SetStructProperty(key, value)
	}
}

func NewRecordLike() RecordLike {
	var result = make(record)

	return &result
}

func NewRecordLikeWithProperties(properties map[string]*structpb.Value) RecordLike {
	if properties == nil {
		properties = make(map[string]*structpb.Value)
	}

	result := make(record)

	result.WithStructProperties(properties)

	return &result
}

func NewRecordLikeFromProperties(properties map[string]interface{}) RecordLike {
	var result = record(properties)

	return &result
}
