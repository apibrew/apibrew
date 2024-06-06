package abs

import (
	"github.com/apibrew/apibrew/pkg/model"
	"google.golang.org/protobuf/types/known/structpb"
	"reflect"
)

type ResourceLike interface {
	GetProperties() []*model.ResourceProperty
	GetTypes() []*model.ResourceSubType
}

// types:
// string
// int32
// int64
// float32
// float64
// bool
// []interface{}
// map[string]interface{}

type RecordLike interface {
	Keys() []string
	GetProperty(key string) interface{}
	GetPropertyWithDefault(key string, defaultValue interface{}) interface{}
	SetProperty(key string, value interface{}) RecordLike
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
	MapCopy() map[string]interface{}
}

type record map[string]interface{}

func (r *record) GetPropertyWithDefault(key string, defaultValue interface{}) interface{} {
	value, ok := (*r)[key]

	if !ok {
		return defaultValue
	}

	return value
}

func (r *record) GetProperty(key string) interface{} {
	return (*r)[key]
}

func (r *record) SetProperty(key string, value interface{}) RecordLike {
	switch typedValue := value.(type) {
	case bool:
	case float64:
	case string:
	case []interface{}:
	case map[string]interface{}:
	case interface{}:
		if value != nil {
			panic("unsupported type: " + reflect.TypeOf(typedValue).String())
		}
	default:
		if value != nil {
			panic("unsupported type: " + reflect.TypeOf(typedValue).String())
		}
	}

	(*r)[key] = value

	return r
}

func (r *record) MapCopy() map[string]interface{} {
	var result = make(map[string]interface{})

	for key, value := range *r {
		result[key] = value
	}

	return result
}

func (r *record) WithProperties(properties map[string]interface{}) RecordLike {
	for key, value := range properties {
		r.SetProperty(key, value)
	}

	return r
}

func (r *record) AsInterface(key string) interface{} {
	return (*r)[key]
}

func (r *record) Merge(appliedRecord RecordLike) RecordLike {
	for _, key := range appliedRecord.Keys() {
		r.SetProperty(key, appliedRecord.GetProperty(key))
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
	if !reflect.DeepEqual(*r, updated.MapCopy()) {
		return false
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
	r.SetProperty(key, value.AsInterface())

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

	return NewRecordLikeWithStructProperties(record.Properties)
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

func UpdateRecordsStructProperties(record RecordLike, properties map[string]*structpb.Value) {
	for key, value := range properties {
		record.SetStructProperty(key, value)
	}
}

func UpdateRecordsProperties(record RecordLike, properties map[string]interface{}) {
	for key, value := range properties {
		record.SetProperty(key, value)
	}
}

func NewRecordLike() RecordLike {
	var result = make(record)

	return &result
}

func NewRecordLikeWithStructProperties(properties map[string]*structpb.Value) RecordLike {
	if properties == nil {
		properties = make(map[string]*structpb.Value)
	}

	result := make(record)

	result.WithStructProperties(properties)

	return &result
}

func NewRecordLikeWithProperties(properties map[string]interface{}) RecordLike {
	if properties == nil {
		properties = make(map[string]interface{})
	}

	result := make(record)

	result.WithProperties(properties)

	return &result
}

func NewRecordLikeFromProperties(properties map[string]interface{}) RecordLike {
	var result = record(properties)

	return &result
}
