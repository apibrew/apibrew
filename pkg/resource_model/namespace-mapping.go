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

	var_c56d8415670a := namespace.Id

	if var_c56d8415670a != nil {
		var var_c56d8415670a_mapped *structpb.Value

		var var_c56d8415670a_err error
		var_c56d8415670a_mapped, var_c56d8415670a_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_c56d8415670a)
		if var_c56d8415670a_err != nil {
			panic(var_c56d8415670a_err)
		}
		properties["id"] = var_c56d8415670a_mapped
	}

	var_f1fef1756538 := namespace.Version

	var var_f1fef1756538_mapped *structpb.Value

	var var_f1fef1756538_err error
	var_f1fef1756538_mapped, var_f1fef1756538_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(var_f1fef1756538)
	if var_f1fef1756538_err != nil {
		panic(var_f1fef1756538_err)
	}
	properties["version"] = var_f1fef1756538_mapped

	var_8a9b02bc4f0b := namespace.CreatedBy

	if var_8a9b02bc4f0b != nil {
		var var_8a9b02bc4f0b_mapped *structpb.Value

		var var_8a9b02bc4f0b_err error
		var_8a9b02bc4f0b_mapped, var_8a9b02bc4f0b_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_8a9b02bc4f0b)
		if var_8a9b02bc4f0b_err != nil {
			panic(var_8a9b02bc4f0b_err)
		}
		properties["createdBy"] = var_8a9b02bc4f0b_mapped
	}

	var_17bdb4565276 := namespace.UpdatedBy

	if var_17bdb4565276 != nil {
		var var_17bdb4565276_mapped *structpb.Value

		var var_17bdb4565276_err error
		var_17bdb4565276_mapped, var_17bdb4565276_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_17bdb4565276)
		if var_17bdb4565276_err != nil {
			panic(var_17bdb4565276_err)
		}
		properties["updatedBy"] = var_17bdb4565276_mapped
	}

	var_e9825c038622 := namespace.CreatedOn

	if var_e9825c038622 != nil {
		var var_e9825c038622_mapped *structpb.Value

		var var_e9825c038622_err error
		var_e9825c038622_mapped, var_e9825c038622_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_e9825c038622)
		if var_e9825c038622_err != nil {
			panic(var_e9825c038622_err)
		}
		properties["createdOn"] = var_e9825c038622_mapped
	}

	var_1197f377e853 := namespace.UpdatedOn

	if var_1197f377e853 != nil {
		var var_1197f377e853_mapped *structpb.Value

		var var_1197f377e853_err error
		var_1197f377e853_mapped, var_1197f377e853_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_1197f377e853)
		if var_1197f377e853_err != nil {
			panic(var_1197f377e853_err)
		}
		properties["updatedOn"] = var_1197f377e853_mapped
	}

	var_50e6c3464ce0 := namespace.Name

	var var_50e6c3464ce0_mapped *structpb.Value

	var var_50e6c3464ce0_err error
	var_50e6c3464ce0_mapped, var_50e6c3464ce0_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_50e6c3464ce0)
	if var_50e6c3464ce0_err != nil {
		panic(var_50e6c3464ce0_err)
	}
	properties["name"] = var_50e6c3464ce0_mapped

	var_ff9c011fc36d := namespace.Description

	if var_ff9c011fc36d != nil {
		var var_ff9c011fc36d_mapped *structpb.Value

		var var_ff9c011fc36d_err error
		var_ff9c011fc36d_mapped, var_ff9c011fc36d_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_ff9c011fc36d)
		if var_ff9c011fc36d_err != nil {
			panic(var_ff9c011fc36d_err)
		}
		properties["description"] = var_ff9c011fc36d_mapped
	}

	var_aef627d44f74 := namespace.Details

	if var_aef627d44f74 != nil {
		var var_aef627d44f74_mapped *structpb.Value

		var var_aef627d44f74_err error
		var_aef627d44f74_mapped, var_aef627d44f74_err = types.ByResourcePropertyType(model.ResourceProperty_OBJECT).Pack(*var_aef627d44f74)
		if var_aef627d44f74_err != nil {
			panic(var_aef627d44f74_err)
		}
		properties["details"] = var_aef627d44f74_mapped
	}

	var_3ffb4b97af35 := namespace.SecurityConstraints

	if var_3ffb4b97af35 != nil {
		var var_3ffb4b97af35_mapped *structpb.Value

		var var_3ffb4b97af35_l []*structpb.Value
		for _, value := range var_3ffb4b97af35 {

			var_46e7e23d79a9 := value
			var var_46e7e23d79a9_mapped *structpb.Value

			var_46e7e23d79a9_mapped = structpb.NewStructValue(&structpb.Struct{Fields: SecurityConstraintMapperInstance.ToProperties(var_46e7e23d79a9)})

			var_3ffb4b97af35_l = append(var_3ffb4b97af35_l, var_46e7e23d79a9_mapped)
		}
		var_3ffb4b97af35_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_3ffb4b97af35_l})
		properties["securityConstraints"] = var_3ffb4b97af35_mapped
	}
	return properties
}

func (m *NamespaceMapper) FromProperties(properties map[string]*structpb.Value) *Namespace {
	var s = m.New()
	if properties["id"] != nil {

		var_4502d646715d := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_4502d646715d)

		if err != nil {
			panic(err)
		}

		var_4502d646715d_mapped := new(uuid.UUID)
		*var_4502d646715d_mapped = val.(uuid.UUID)

		s.Id = var_4502d646715d_mapped
	}
	if properties["version"] != nil {

		var_b61961e5de80 := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_b61961e5de80)

		if err != nil {
			panic(err)
		}

		var_b61961e5de80_mapped := val.(int32)

		s.Version = var_b61961e5de80_mapped
	}
	if properties["createdBy"] != nil {

		var_0e17464bdc25 := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_0e17464bdc25)

		if err != nil {
			panic(err)
		}

		var_0e17464bdc25_mapped := new(string)
		*var_0e17464bdc25_mapped = val.(string)

		s.CreatedBy = var_0e17464bdc25_mapped
	}
	if properties["updatedBy"] != nil {

		var_10708d1fe65f := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_10708d1fe65f)

		if err != nil {
			panic(err)
		}

		var_10708d1fe65f_mapped := new(string)
		*var_10708d1fe65f_mapped = val.(string)

		s.UpdatedBy = var_10708d1fe65f_mapped
	}
	if properties["createdOn"] != nil {

		var_d524bb475c25 := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_d524bb475c25)

		if err != nil {
			panic(err)
		}

		var_d524bb475c25_mapped := new(time.Time)
		*var_d524bb475c25_mapped = val.(time.Time)

		s.CreatedOn = var_d524bb475c25_mapped
	}
	if properties["updatedOn"] != nil {

		var_ff8bb65b8f94 := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_ff8bb65b8f94)

		if err != nil {
			panic(err)
		}

		var_ff8bb65b8f94_mapped := new(time.Time)
		*var_ff8bb65b8f94_mapped = val.(time.Time)

		s.UpdatedOn = var_ff8bb65b8f94_mapped
	}
	if properties["name"] != nil {

		var_3186850d9b5d := properties["name"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_3186850d9b5d)

		if err != nil {
			panic(err)
		}

		var_3186850d9b5d_mapped := val.(string)

		s.Name = var_3186850d9b5d_mapped
	}
	if properties["description"] != nil {

		var_350077b5af5f := properties["description"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_350077b5af5f)

		if err != nil {
			panic(err)
		}

		var_350077b5af5f_mapped := new(string)
		*var_350077b5af5f_mapped = val.(string)

		s.Description = var_350077b5af5f_mapped
	}
	if properties["details"] != nil {

		var_a29a5b887868 := properties["details"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_OBJECT).UnPack(var_a29a5b887868)

		if err != nil {
			panic(err)
		}

		var_a29a5b887868_mapped := new(unstructured.Unstructured)
		*var_a29a5b887868_mapped = val.(unstructured.Unstructured)

		s.Details = var_a29a5b887868_mapped
	}
	if properties["securityConstraints"] != nil {

		var_e57873df6965 := properties["securityConstraints"]
		var_e57873df6965_mapped := []*SecurityConstraint{}
		for _, v := range var_e57873df6965.GetListValue().Values {

			var_cd5b91d28dce := v
			var_cd5b91d28dce_mapped := SecurityConstraintMapperInstance.FromProperties(var_cd5b91d28dce.GetStructValue().Fields)

			var_e57873df6965_mapped = append(var_e57873df6965_mapped, var_cd5b91d28dce_mapped)
		}

		s.SecurityConstraints = var_e57873df6965_mapped
	}
	return s
}
