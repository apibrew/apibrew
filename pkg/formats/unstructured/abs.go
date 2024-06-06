package unstructured

import (
	"github.com/apibrew/apibrew/pkg/abs"
	"google.golang.org/protobuf/runtime/protoimpl"
	"google.golang.org/protobuf/types/known/structpb"
	"unicode/utf8"
)

type Any = interface{}
type Unstructured = map[string]Any

func MergeInto(u Unstructured, un Unstructured, nested bool) {
	for key, value := range un {
		if !nested || u[key] == nil {
			u[key] = value
		} else {
			if subU, ok := u[key].(Unstructured); ok {
				MergeInto(subU, value.(Unstructured), nested)
			} else if subU, ok := u[key].([]Any); ok {
				subU = append(subU, value.([]Any)...)
				u[key] = subU
			} else {
				u[key] = value
			}
		}
	}
}

func MergeOut(u Unstructured, un Unstructured, nested bool) {
	for key, value := range un {
		if !nested || u[key] == nil {
			u[key] = value
		} else {
			if subU, ok := u[key].(Unstructured); ok {
				MergeOut(subU, value.(Unstructured), nested)
			} else if subU, ok := u[key].([]Any); ok {
				subU = append(subU, value.([]Any)...)
				u[key] = subU
			}
		}
	}
}

func FromRecord(record abs.RecordLike) (Unstructured, error) {
	return record.MapCopy(), nil
}

func ToRecord(u Unstructured) (abs.RecordLike, error) {
	return abs.NewRecordLikeFromProperties(u), nil
}

func Keys(u Unstructured) []string {
	var keys []string

	for key := range u {
		keys = append(keys, key)
	}

	return keys
}

func DeleteKey(u Unstructured, key string) {
	delete(u, key)
}

func ToValue(v Any) (*structpb.Value, error) {
	switch v := v.(type) {
	case Unstructured:
		x := &structpb.Struct{Fields: make(map[string]*structpb.Value, len(v))}
		for k, v := range v {
			if !utf8.ValidString(k) {
				return nil, protoimpl.X.NewError("invalid UTF-8 in string: %q", k)
			}
			var err error
			x.Fields[k], err = ToValue(v)
			if err != nil {
				return nil, err
			}
		}
		return structpb.NewStructValue(x), nil
	case []Any:
		x := &structpb.ListValue{Values: make([]*structpb.Value, len(v))}
		for i, v := range v {
			var err error
			x.Values[i], err = ToValue(v)
			if err != nil {
				return nil, err
			}
		}
		return structpb.NewListValue(x), nil
	default:
		return structpb.NewValue(v)
	}
}

func FromStructValue(v *structpb.Struct) Unstructured {
	if v == nil {
		return nil
	}

	u := make(Unstructured, len(v.Fields))

	for k, v := range v.Fields {
		u[k] = FromValue(v)
	}

	return u
}

func FromValue(v *structpb.Value) Any {
	if v == nil {
		return nil
	}
	switch v := v.Kind.(type) {
	case *structpb.Value_NullValue:
		return nil
	case *structpb.Value_NumberValue:
		return v.NumberValue
	case *structpb.Value_StringValue:
		return v.StringValue
	case *structpb.Value_BoolValue:
		return v.BoolValue
	case *structpb.Value_StructValue:
		return FromStructValue(v.StructValue)
	case *structpb.Value_ListValue:
		return FromListValue(v.ListValue)
	default:
		return nil
	}
}

func FromListValue(value *structpb.ListValue) Any {
	var list []Any

	for _, v := range value.Values {
		list = append(list, FromValue(v))
	}

	return list
}
