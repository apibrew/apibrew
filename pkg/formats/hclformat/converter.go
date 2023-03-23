package hclformat

import (
	"encoding/base64"

	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/gocty"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// ToProtobufMessage writes the values from the given value into the
// fields of the given protobuf message.
//
// The given value must have an object type matching what
// ImpliedTypeForMessageDesc would return for the message descriptor
// associated with the message given in "into", or else decoding will
// fail.
//
// The types in the protocol buffers type system can have a smaller range
// than the corresponding cty types we convert from, so this function might
// return an error if the given values are out of range. In those cases,
// the returned error will be a cty.PathError with a message written to
// be understood by an end-user who provided whatever data was converted
// to cty.Value, without mentioning protobuf implementation details.
//
// In case of any error, the given message may be partially updated.
//
// Protocol buffers has no concept of an unknown value, so ToProtobufMessage
// will return an error if there are any unknown values in the given object.
// Don't pass marked values to ToProtobufMessage; it will panic if it
// encounters any values that are marked.
func ToProtobufMessage(obj cty.Value, into protoreflect.Message) error {
	path := make(cty.Path, 0, 4)
	if obj.IsNull() {
		return path.NewErrorf("must not be null")
	}
	if !obj.IsKnown() {
		return path.NewErrorf("value must be known")
	}
	return toProtobufMessage(obj, into, path)
}

func toProtobufMessage(obj cty.Value, into protoreflect.Message, path cty.Path) error {

	desc := into.Descriptor()
	fields := desc.Fields()
	ty := obj.Type()

	if !ty.IsObjectType() {
		return path.NewErrorf("an object is required")
	}

	// TODO: Verify that any "oneofs" are well-formed, such
	// that each one has only one of its fields non-null.

	for i := 0; i < fields.Len(); i++ {
		field := fields.Get(i)
		name := string(field.Name())

		if !ty.HasAttribute(name) {
			return path.NewErrorf("missing required attribute %q", name)
		}

		// Temporarily extend path with new attribute name
		path := append(path, cty.GetAttrStep{Name: name})

		av := obj.GetAttr(name)
		err := toProtobufMessageField(into, field, av, path)
		if err != nil {
			return err
		}
	}

	return nil
}

func toProtobufMessageField(msg protoreflect.Message, field protoreflect.FieldDescriptor, v cty.Value, path cty.Path) error {
	if v.IsNull() {
		msg.Clear(field)
		if !field.HasPresence() {
			return path.NewErrorf("must not be null")
		}
		return nil
	}
	if !v.IsKnown() {
		return path.NewErrorf("value must be known")
	}
	ty := v.Type()

	switch {
	case field.IsMap():
		keyField := field.MapKey()
		valField := field.MapValue()
		// We have a different representation for maps with string keys vs.
		// maps with other key types.
		switch {
		case keyField.Kind() == protoreflect.StringKind:
			// Should be a cty.Map whose element type corresponds with valField.
			if !ty.IsMapType() {
				return path.NewErrorf("a map is required")
			}
			protoMap := msg.NewField(field).Map()
			for it := v.ElementIterator(); it.Next(); {
				ek, ev := it.Element()
				path := append(path, cty.IndexStep{Key: ek})
				ekProto := protoreflect.MapKey(protoreflect.ValueOfString(ek.AsString()))
				evProto, err := toProtobufValue(ev, valField, func() protoreflect.Value {
					return protoMap.Mutable(ekProto)
				}, path)
				if err != nil {
					return err
				}
				protoMap.Set(ekProto, evProto)
			}
			msg.Set(field, protoreflect.ValueOfMap(protoMap))
		default:
			// Should be a cty.Set whose element type is an object with
			// key and value attributes.
			if !ty.IsSetType() {
				return path.NewErrorf("a set of objects is required")
			}
			ety := ty.ElementType()
			if !ety.IsObjectType() {
				return path.NewErrorf("a set of objects is required")
			}
			atys := ety.AttributeTypes()
			if _, exists := atys["key"]; !exists {
				return path.NewErrorf("set element type must have attribute \"key\"")
			}
			if _, exists := atys["value"]; !exists {
				return path.NewErrorf("set element type must have attribute \"value\"")
			}
			if len(atys) != 2 {
				return path.NewErrorf("set element type must only have attributes \"key\" and \"value\"")
			}
			protoMap := msg.NewField(field).Map()
			// In this case we'll decode into the message type that the
			// proto compiler generated to represent the map elements,
			// since our element type ought to be compatible with it.
			msg.Clear(field)
			for it := v.ElementIterator(); it.Next(); {
				_, ev := it.Element()
				path := append(path, cty.IndexStep{Key: ev})

				keyVal := ev.GetAttr("key")
				valVal := ev.GetAttr("value")

				keyProto, err := toProtobufValue(keyVal, keyField, nil, path)
				if err != nil {
					return err
				}
				valProto, err := toProtobufValue(valVal, valField, func() protoreflect.Value {
					return protoMap.Mutable(protoreflect.MapKey(keyProto))
				}, path)
				if err != nil {
					return err
				}

				protoMap.Set(protoreflect.MapKey(keyProto), valProto)
			}
			msg.Set(field, protoreflect.ValueOfMap(protoMap))
		}
	case field.IsList():
		if !ty.IsListType() {
			return path.NewErrorf("a list is required")
		}
		msg.Clear(field)
		protoList := msg.NewField(field).List()
		for it := v.ElementIterator(); it.Next(); {
			_, ev := it.Element()
			path := append(path, cty.IndexStep{Key: ev})

			alreadyAppended := false
			evProto, err := toProtobufValue(ev, field, func() protoreflect.Value {
				alreadyAppended = true
				return protoList.AppendMutable()
			}, path)
			if err != nil {
				return err
			}
			if !alreadyAppended {
				protoList.Append(evProto)
			}
		}
		msg.Set(field, protoreflect.ValueOfList(protoList))
	default:
		vProto, err := toProtobufValue(v, field, func() protoreflect.Value {
			return msg.Mutable(field)
		}, path)
		if err != nil {
			return err
		}
		msg.Set(field, vProto)
	}

	return nil
}

// toProtobufValue is a pretty awkward function that deals with decoding
// individual cty values into arbitrary protocol buffers values. This is
// made particularly awkward because protoreflect handles differently
// primitive type values, nested message values, and the various compound
// types represented by "repeated".
//
// This method can only deal with the first two categories. If a field is
// repeated (a list or a map) then the caller must deal with the cardinality
// business.
//
// To deal with nested message values this function requires a callback to
// obtain a new message value that's already assigned into wherever it's
// going to end up. In that case, this function guarantees to return that
// same value, so the caller might choose not to reassign it in that case,
// but it also doesn't hurt to assign it again for simplicity's sake.
//
// toProtobufValue can't deal with null or unknown values. The caller
// should deal with that first, before calling.
func toProtobufValue(v cty.Value, field protoreflect.FieldDescriptor, mut func() protoreflect.Value, path cty.Path) (protoreflect.Value, error) {
	var nothing protoreflect.Value
	kind := field.Kind()
	ty := v.Type()
	switch kind {
	case protoreflect.BoolKind:
		if !cty.Bool.Equals(ty) {
			return nothing, path.NewErrorf("a boolean value is required")
		}
		return protoreflect.ValueOfBool(v.True()), nil
	case protoreflect.StringKind:
		if !cty.String.Equals(ty) {
			return nothing, path.NewErrorf("a string is required")
		}
		return protoreflect.ValueOfString(v.AsString()), nil
	case protoreflect.BytesKind:
		if !cty.String.Equals(ty) {
			return nothing, path.NewErrorf("a string containing base64 bytes is required")
		}
		b64s := v.AsString()
		bytes, err := base64.StdEncoding.DecodeString(b64s)
		if err != nil {
			return nothing, path.NewErrorf("string must contain base64-encoded bytes")
		}
		return protoreflect.ValueOfBytes(bytes), nil
	case protoreflect.EnumKind:
		if !cty.String.Equals(ty) {
			return nothing, path.NewErrorf("a string containing a keyword is required")
		}
		name := protoreflect.Name(v.AsString())
		enumDesc := field.Enum()
		optionDesc := enumDesc.Values().ByName(name)
		if optionDesc == nil {
			return nothing, path.NewErrorf("value isn't one of the expected keywords")
		}
		return protoreflect.ValueOfEnum(optionDesc.Number()), nil
	case protoreflect.MessageKind:
		msg := mut().Message()
		err := toProtobufMessage(v, msg, path)
		if err != nil {
			return nothing, err
		}
		return protoreflect.ValueOfMessage(msg), nil
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		var n int32
		err := gocty.FromCtyValue(v, &n)
		if err != nil {
			return nothing, path.NewError(err)
		}
		return protoreflect.ValueOfInt32(n), nil
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		var n uint32
		err := gocty.FromCtyValue(v, &n)
		if err != nil {
			return nothing, path.NewError(err)
		}
		return protoreflect.ValueOfUint32(n), nil
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		var n int64
		err := gocty.FromCtyValue(v, &n)
		if err != nil {
			return nothing, path.NewError(err)
		}
		return protoreflect.ValueOfInt64(n), nil
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		var n uint64
		err := gocty.FromCtyValue(v, &n)
		if err != nil {
			return nothing, path.NewError(err)
		}
		return protoreflect.ValueOfUint64(n), nil
	case protoreflect.FloatKind:
		var n float32
		err := gocty.FromCtyValue(v, &n)
		if err != nil {
			return nothing, path.NewError(err)
		}
		return protoreflect.ValueOfFloat32(n), nil
	case protoreflect.DoubleKind:
		var n float64
		err := gocty.FromCtyValue(v, &n)
		if err != nil {
			return nothing, path.NewError(err)
		}
		return protoreflect.ValueOfFloat64(n), nil
	default:
		return nothing, path.NewErrorf("no cty equivalent for protobuf kind %s", kind.String())
	}
}
