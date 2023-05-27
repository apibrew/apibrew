package unstructured

import (
	"encoding/json"
	"github.com/apibrew/apibrew/pkg/model"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/runtime/protoimpl"
	"google.golang.org/protobuf/types/known/structpb"
	"unicode/utf8"
)

type Unstructured map[string]interface{}

func (u Unstructured) MergeInto(un Unstructured, nested bool) {
	for key, value := range un {
		if !nested || u[key] == nil {
			u[key] = value
		} else {
			if subU, ok := u[key].(Unstructured); ok {
				subU.MergeInto(value.(Unstructured), nested)
			} else if subU, ok := u[key].([]interface{}); ok {
				subU = append(subU, value.([]interface{})...)
				u[key] = subU
			} else {
				u[key] = value
			}
		}
	}
}

func (u Unstructured) MergeOut(un Unstructured, nested bool) {
	for key, value := range un {
		if !nested || u[key] == nil {
			u[key] = value
		} else {
			if subU, ok := u[key].(Unstructured); ok {
				subU.MergeOut(value.(Unstructured), nested)
			} else if subU, ok := u[key].([]interface{}); ok {
				subU = append(subU, value.([]interface{})...)
				u[key] = subU
			}
		}
	}
}

func (u Unstructured) ToProtoMessage(msg proto.Message) error {
	b, err := json.Marshal(u)
	if err != nil {
		return err
	}

	return jsonUMo.Unmarshal(b, msg)
}

func (u Unstructured) FromProtoMessage(msg proto.Message) error {
	b, err := jsonMo.Marshal(msg)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, &u)
}

var jsonMo = protojson.MarshalOptions{
	Multiline:       true,
	EmitUnpopulated: false,
}

var jsonUMo = protojson.UnmarshalOptions{
	AllowPartial:   true,
	DiscardUnknown: false,
	Resolver:       nil,
}

func (u Unstructured) ToRecord() (*model.Record, error) {
	record := &model.Record{}
	properties, err := u.ToProperties()

	if err != nil {
		return nil, err
	}

	record.Properties = properties

	return record, nil
}

func (u Unstructured) ToProperties() (map[string]*structpb.Value, error) {
	var properties = make(map[string]*structpb.Value)

	for key, value := range u {
		var err error
		properties[key], err = NewStructValue(value)

		if err != nil {
			return nil, err
		}
	}

	return properties, nil
}

func (u Unstructured) Keys() []string {
	var keys []string

	for key := range u {
		keys = append(keys, key)
	}

	return keys
}

func (u Unstructured) DeleteKey(key string) {
	delete(u, key)
}

func NewStructValue(v interface{}) (*structpb.Value, error) {
	switch v := v.(type) {
	case Unstructured:
		x := &structpb.Struct{Fields: make(map[string]*structpb.Value, len(v))}
		for k, v := range v {
			if !utf8.ValidString(k) {
				return nil, protoimpl.X.NewError("invalid UTF-8 in string: %q", k)
			}
			var err error
			x.Fields[k], err = NewStructValue(v)
			if err != nil {
				return nil, err
			}
		}
		return structpb.NewStructValue(x), nil
	case []interface{}:
		x := &structpb.ListValue{Values: make([]*structpb.Value, len(v))}
		for i, v := range v {
			var err error
			x.Values[i], err = NewStructValue(v)
			if err != nil {
				return nil, err
			}
		}
		return structpb.NewListValue(x), nil
	default:
		return structpb.NewValue(v)
	}
}
