package proto

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resources/special"
	"github.com/apibrew/apibrew/pkg/service/annotations"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

type ResourceHelper struct {
}

// convert proto message to resource
func (h ResourceHelper) ProtoToResource(reflect protoreflect.Message) *model.Resource {
	resource := new(model.Resource)

	messageOptions := reflect.Descriptor().Options().(*descriptorpb.MessageOptions)

	resource.Name = proto.GetExtension(messageOptions, model.E_ResourceName).(string)
	resource.Namespace = proto.GetExtension(messageOptions, model.E_ResourceNamespace).(string)

	resource.SourceConfig = &model.ResourceSourceConfig{
		DataSource: proto.GetExtension(messageOptions, model.E_ResourceDataSource).(string),
		Entity:     proto.GetExtension(messageOptions, model.E_ResourceEntity).(string),
		Catalog:    proto.GetExtension(messageOptions, model.E_ResourceCatalog).(string),
	}

	resource.Annotations = make(map[string]string)

	if proto.GetExtension(messageOptions, model.E_SecurityContextDisallowAll).(bool) {
		resource.SecurityContext = special.SecurityContextDisallowAll
	}

	fields := reflect.Descriptor().Fields()

	for i := 0; i < fields.Len(); i++ {
		field := fields.Get(i)

		// special properties
		if field.Kind() == protoreflect.MessageKind && field.Message().FullName() == "model.AuditData" {
			resource.Properties = append(resource.Properties, special.AuditProperties...)
			annotations.Enable(resource, annotations.EnableAudit)
			continue
		}
		if field.Kind() == protoreflect.Uint32Kind && field.Name() == "version" {
			resource.Properties = append(resource.Properties, special.VersionProperty)
			continue
		}
		if field.Kind() == protoreflect.StringKind && field.Name() == "id" {
			resource.Properties = append(resource.Properties, special.IdProperty)
			continue
		}

		resource.Properties = append(resource.Properties, h.prepareResourcePropertyFromField(field))
	}

	return resource
}

func (h ResourceHelper) prepareResourcePropertyFromField(field protoreflect.FieldDescriptor) *model.ResourceProperty {
	property := new(model.ResourceProperty)

	fieldOptions := field.Options().(*descriptorpb.FieldOptions)

	fieldOptionsType := proto.GetExtension(fieldOptions, model.E_PropertyType).(model.ResourceProperty_Type)

	property.Name = string(field.Name())
	property.Unique = proto.GetExtension(fieldOptions, model.E_PropertyUnique).(bool)
	property.Immutable = proto.GetExtension(fieldOptions, model.E_PropertyImmutable).(bool)
	property.Mapping = proto.GetExtension(fieldOptions, model.E_PropertyMapping).(string)
	property.Length = proto.GetExtension(fieldOptions, model.E_PropertyLength).(uint32)
	property.Required = !field.HasOptionalKeyword()

	propertyAnnotations := proto.GetExtension(fieldOptions, model.E_PropertyAnnotations).([]*model.Annotation)

	for _, annotation := range propertyAnnotations {
		if property.Annotations == nil {
			property.Annotations = make(map[string]string)
		}
		property.Annotations[annotation.Name] = annotation.Value
	}

	switch field.Kind() {
	case protoreflect.MessageKind:
		property.Type = model.ResourceProperty_OBJECT
	case protoreflect.EnumKind:
		property.Type = model.ResourceProperty_ENUM
	case protoreflect.StringKind:
		property.Type = model.ResourceProperty_STRING
		if fieldOptionsType == model.ResourceProperty_UUID {
			property.Type = model.ResourceProperty_UUID
		}
	case protoreflect.BoolKind:
		property.Type = model.ResourceProperty_BOOL
	case protoreflect.Int32Kind, protoreflect.Uint32Kind:
		property.Type = model.ResourceProperty_INT32
	case protoreflect.Int64Kind, protoreflect.Uint64Kind:
		property.Type = model.ResourceProperty_INT64
	case protoreflect.FloatKind:
		property.Type = model.ResourceProperty_FLOAT32
	case protoreflect.DoubleKind:
		property.Type = model.ResourceProperty_FLOAT64
	case protoreflect.BytesKind:
		property.Type = model.ResourceProperty_BYTES
	default:
		panic("unknown field type" + field.Kind().String())
	}

	if property.Type == model.ResourceProperty_STRING && property.Length == 0 {
		property.Length = 256
	}

	if property.Mapping == "" {
		property.Mapping = property.Name
	}

	if field.IsList() {
		listProperty := new(model.ResourceProperty)
		listProperty.Type = model.ResourceProperty_LIST
		listProperty.Name = property.Name
		listProperty.Mapping = property.Mapping
		property.Name = ""
		property.Mapping = ""
		property.Required = false
		property.Unique = false

		listProperty.Item = property
		return listProperty
	}

	return property
}
