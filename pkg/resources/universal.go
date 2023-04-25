package resources

import (
	"github.com/tislib/apibrew/pkg/model"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func ProtoDescriptorToResource(descriptor protoreflect.Message) *model.Resource {
	return &model.Resource{
		Name:            string(descriptor.Descriptor().Name()),
		Namespace:       "system",
		Properties:      mapProperties(descriptor),
		Indexes:         nil,
		SecurityContext: nil,
		Virtual:         false,
		Immutable:       false,
		Abstract:        false,
		Title:           nil,
		Description:     nil,
		AuditData:       nil,
		Version:         0,
		Annotations:     nil,
	}
}

func mapProperties(descriptor protoreflect.Message) []*model.ResourceProperty {
	var properties []*model.ResourceProperty

	for i := 0; i < descriptor.Descriptor().Fields().Len(); i++ {
		field := descriptor.Descriptor().Fields().Get(i)
		properties = append(properties, mapProperty(field))
	}

	return properties
}

func mapProperty(field protoreflect.FieldDescriptor) *model.ResourceProperty {
	return &model.ResourceProperty{
		Name:     string(field.Name()),
		Mapping:  string(field.Name()),
		Primary:  false,
		Type:     mapType(field),
		Length:   0,
		Required: false,
		Unique:   false,
	}
}

func mapType(field protoreflect.FieldDescriptor) model.ResourceProperty_Type {
	switch field.Kind() {
	case protoreflect.BoolKind:
		return model.ResourceProperty_BOOL
	case protoreflect.Int32Kind:
		return model.ResourceProperty_INT32
	case protoreflect.Int64Kind:
		return model.ResourceProperty_INT64
	case protoreflect.Uint32Kind:
		return model.ResourceProperty_INT32
	case protoreflect.Uint64Kind:
		return model.ResourceProperty_INT64
	case protoreflect.FloatKind:
		return model.ResourceProperty_FLOAT32
	case protoreflect.DoubleKind:
		return model.ResourceProperty_FLOAT64
	case protoreflect.StringKind:
		return model.ResourceProperty_STRING
	case protoreflect.BytesKind:
		return model.ResourceProperty_BYTES
	case protoreflect.MessageKind:
		return model.ResourceProperty_REFERENCE
	case protoreflect.EnumKind:
		return model.ResourceProperty_ENUM
	default:
		panic("Unknown type")
		//return model.ResourceProperty_OBJECT
	}
}
