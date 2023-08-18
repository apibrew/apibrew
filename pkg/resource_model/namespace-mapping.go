package resource_model

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/types"
	"google.golang.org/protobuf/types/known/structpb"
)

import "github.com/google/uuid"
import "time"
import "github.com/apibrew/apibrew/pkg/formats/unstructured"

type NamespaceMapper struct {
}

func NewNamespaceMapper() *NamespaceMapper {
	return &NamespaceMapper{}
}

var NamespaceMapperInstance = NewNamespaceMapper()

func (m *NamespaceMapper) New() *Namespace {
	return &Namespace{}
}

func (m *NamespaceMapper) ToRecord(namespace *Namespace) *model.Record {
	var rec = &model.Record{}
	rec.Properties = m.ToProperties(namespace)

	if namespace.Id != nil {
		rec.Id = namespace.Id.String()
	}

	return rec
}

func (m *NamespaceMapper) FromRecord(record *model.Record) *Namespace {
	return m.FromProperties(record.Properties)
}

func (m *NamespaceMapper) ToProperties(namespace *Namespace) map[string]*structpb.Value {
	var properties = make(map[string]*structpb.Value)

	var_1e01c8de88ac := namespace.Id

	if var_1e01c8de88ac != nil {
		var var_1e01c8de88ac_mapped *structpb.Value

		var var_1e01c8de88ac_err error
		var_1e01c8de88ac_mapped, var_1e01c8de88ac_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_1e01c8de88ac)
		if var_1e01c8de88ac_err != nil {
			panic(var_1e01c8de88ac_err)
		}
		properties["id"] = var_1e01c8de88ac_mapped
	}

	var_7eeb295f5bbb := namespace.Version

	var var_7eeb295f5bbb_mapped *structpb.Value

	var var_7eeb295f5bbb_err error
	var_7eeb295f5bbb_mapped, var_7eeb295f5bbb_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(var_7eeb295f5bbb)
	if var_7eeb295f5bbb_err != nil {
		panic(var_7eeb295f5bbb_err)
	}
	properties["version"] = var_7eeb295f5bbb_mapped

	var_47fa09e78062 := namespace.CreatedBy

	if var_47fa09e78062 != nil {
		var var_47fa09e78062_mapped *structpb.Value

		var var_47fa09e78062_err error
		var_47fa09e78062_mapped, var_47fa09e78062_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_47fa09e78062)
		if var_47fa09e78062_err != nil {
			panic(var_47fa09e78062_err)
		}
		properties["createdBy"] = var_47fa09e78062_mapped
	}

	var_3983a1e61692 := namespace.UpdatedBy

	if var_3983a1e61692 != nil {
		var var_3983a1e61692_mapped *structpb.Value

		var var_3983a1e61692_err error
		var_3983a1e61692_mapped, var_3983a1e61692_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_3983a1e61692)
		if var_3983a1e61692_err != nil {
			panic(var_3983a1e61692_err)
		}
		properties["updatedBy"] = var_3983a1e61692_mapped
	}

	var_6c2ee0146206 := namespace.CreatedOn

	if var_6c2ee0146206 != nil {
		var var_6c2ee0146206_mapped *structpb.Value

		var var_6c2ee0146206_err error
		var_6c2ee0146206_mapped, var_6c2ee0146206_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_6c2ee0146206)
		if var_6c2ee0146206_err != nil {
			panic(var_6c2ee0146206_err)
		}
		properties["createdOn"] = var_6c2ee0146206_mapped
	}

	var_adb9fb8c10c2 := namespace.UpdatedOn

	if var_adb9fb8c10c2 != nil {
		var var_adb9fb8c10c2_mapped *structpb.Value

		var var_adb9fb8c10c2_err error
		var_adb9fb8c10c2_mapped, var_adb9fb8c10c2_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_adb9fb8c10c2)
		if var_adb9fb8c10c2_err != nil {
			panic(var_adb9fb8c10c2_err)
		}
		properties["updatedOn"] = var_adb9fb8c10c2_mapped
	}

	var_6948c52d79f6 := namespace.Name

	var var_6948c52d79f6_mapped *structpb.Value

	var var_6948c52d79f6_err error
	var_6948c52d79f6_mapped, var_6948c52d79f6_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_6948c52d79f6)
	if var_6948c52d79f6_err != nil {
		panic(var_6948c52d79f6_err)
	}
	properties["name"] = var_6948c52d79f6_mapped

	var_fa380deea4c6 := namespace.Description

	if var_fa380deea4c6 != nil {
		var var_fa380deea4c6_mapped *structpb.Value

		var var_fa380deea4c6_err error
		var_fa380deea4c6_mapped, var_fa380deea4c6_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_fa380deea4c6)
		if var_fa380deea4c6_err != nil {
			panic(var_fa380deea4c6_err)
		}
		properties["description"] = var_fa380deea4c6_mapped
	}

	var_a8a57f776be8 := namespace.Details

	if var_a8a57f776be8 != nil {
		var var_a8a57f776be8_mapped *structpb.Value

		var var_a8a57f776be8_err error
		var_a8a57f776be8_mapped, var_a8a57f776be8_err = types.ByResourcePropertyType(model.ResourceProperty_OBJECT).Pack(*var_a8a57f776be8)
		if var_a8a57f776be8_err != nil {
			panic(var_a8a57f776be8_err)
		}
		properties["details"] = var_a8a57f776be8_mapped
	}
	return properties
}

func (m *NamespaceMapper) FromProperties(properties map[string]*structpb.Value) *Namespace {
	var s = m.New()
	if properties["id"] != nil {

		var_f7820fd43569 := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_f7820fd43569)

		if err != nil {
			panic(err)
		}

		var_f7820fd43569_mapped := new(uuid.UUID)
		*var_f7820fd43569_mapped = val.(uuid.UUID)

		s.Id = var_f7820fd43569_mapped
	}
	if properties["version"] != nil {

		var_3f49fffc6938 := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_3f49fffc6938)

		if err != nil {
			panic(err)
		}

		var_3f49fffc6938_mapped := val.(int32)

		s.Version = var_3f49fffc6938_mapped
	}
	if properties["createdBy"] != nil {

		var_990cec7ee165 := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_990cec7ee165)

		if err != nil {
			panic(err)
		}

		var_990cec7ee165_mapped := new(string)
		*var_990cec7ee165_mapped = val.(string)

		s.CreatedBy = var_990cec7ee165_mapped
	}
	if properties["updatedBy"] != nil {

		var_51e2537bd9ac := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_51e2537bd9ac)

		if err != nil {
			panic(err)
		}

		var_51e2537bd9ac_mapped := new(string)
		*var_51e2537bd9ac_mapped = val.(string)

		s.UpdatedBy = var_51e2537bd9ac_mapped
	}
	if properties["createdOn"] != nil {

		var_152c7ea7dccf := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_152c7ea7dccf)

		if err != nil {
			panic(err)
		}

		var_152c7ea7dccf_mapped := new(time.Time)
		*var_152c7ea7dccf_mapped = val.(time.Time)

		s.CreatedOn = var_152c7ea7dccf_mapped
	}
	if properties["updatedOn"] != nil {

		var_1a549985d0e9 := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_1a549985d0e9)

		if err != nil {
			panic(err)
		}

		var_1a549985d0e9_mapped := new(time.Time)
		*var_1a549985d0e9_mapped = val.(time.Time)

		s.UpdatedOn = var_1a549985d0e9_mapped
	}
	if properties["name"] != nil {

		var_7b77ac3f17be := properties["name"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_7b77ac3f17be)

		if err != nil {
			panic(err)
		}

		var_7b77ac3f17be_mapped := val.(string)

		s.Name = var_7b77ac3f17be_mapped
	}
	if properties["description"] != nil {

		var_aa357ba62f09 := properties["description"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_aa357ba62f09)

		if err != nil {
			panic(err)
		}

		var_aa357ba62f09_mapped := new(string)
		*var_aa357ba62f09_mapped = val.(string)

		s.Description = var_aa357ba62f09_mapped
	}
	if properties["details"] != nil {

		var_9b18fd648c18 := properties["details"]
		var_9b18fd648c18_mapped := new(unstructured.Unstructured)
		*var_9b18fd648c18_mapped = unstructured.FromStructValue(var_9b18fd648c18.GetStructValue())

		s.Details = var_9b18fd648c18_mapped
	}
	return s
}
