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
	return rec
}

func (m *NamespaceMapper) FromRecord(record *model.Record) *Namespace {
	return m.FromProperties(record.Properties)
}

func (m *NamespaceMapper) ToProperties(namespace *Namespace) map[string]*structpb.Value {
	var properties = make(map[string]*structpb.Value)

	var_2cd168f5589b := namespace.Id

	if var_2cd168f5589b != nil {
		var var_2cd168f5589b_mapped *structpb.Value

		var var_2cd168f5589b_err error
		var_2cd168f5589b_mapped, var_2cd168f5589b_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_2cd168f5589b)
		if var_2cd168f5589b_err != nil {
			panic(var_2cd168f5589b_err)
		}
		properties["id"] = var_2cd168f5589b_mapped
	}

	var_ff2075e94cbf := namespace.Version

	var var_ff2075e94cbf_mapped *structpb.Value

	var var_ff2075e94cbf_err error
	var_ff2075e94cbf_mapped, var_ff2075e94cbf_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(var_ff2075e94cbf)
	if var_ff2075e94cbf_err != nil {
		panic(var_ff2075e94cbf_err)
	}
	properties["version"] = var_ff2075e94cbf_mapped

	var_465042e21992 := namespace.CreatedBy

	if var_465042e21992 != nil {
		var var_465042e21992_mapped *structpb.Value

		var var_465042e21992_err error
		var_465042e21992_mapped, var_465042e21992_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_465042e21992)
		if var_465042e21992_err != nil {
			panic(var_465042e21992_err)
		}
		properties["createdBy"] = var_465042e21992_mapped
	}

	var_672d27e6e12c := namespace.UpdatedBy

	if var_672d27e6e12c != nil {
		var var_672d27e6e12c_mapped *structpb.Value

		var var_672d27e6e12c_err error
		var_672d27e6e12c_mapped, var_672d27e6e12c_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_672d27e6e12c)
		if var_672d27e6e12c_err != nil {
			panic(var_672d27e6e12c_err)
		}
		properties["updatedBy"] = var_672d27e6e12c_mapped
	}

	var_46266cc181e4 := namespace.CreatedOn

	if var_46266cc181e4 != nil {
		var var_46266cc181e4_mapped *structpb.Value

		var var_46266cc181e4_err error
		var_46266cc181e4_mapped, var_46266cc181e4_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_46266cc181e4)
		if var_46266cc181e4_err != nil {
			panic(var_46266cc181e4_err)
		}
		properties["createdOn"] = var_46266cc181e4_mapped
	}

	var_d23988d47d6d := namespace.UpdatedOn

	if var_d23988d47d6d != nil {
		var var_d23988d47d6d_mapped *structpb.Value

		var var_d23988d47d6d_err error
		var_d23988d47d6d_mapped, var_d23988d47d6d_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_d23988d47d6d)
		if var_d23988d47d6d_err != nil {
			panic(var_d23988d47d6d_err)
		}
		properties["updatedOn"] = var_d23988d47d6d_mapped
	}

	var_6d7b396cda02 := namespace.Name

	var var_6d7b396cda02_mapped *structpb.Value

	var var_6d7b396cda02_err error
	var_6d7b396cda02_mapped, var_6d7b396cda02_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_6d7b396cda02)
	if var_6d7b396cda02_err != nil {
		panic(var_6d7b396cda02_err)
	}
	properties["name"] = var_6d7b396cda02_mapped

	var_5cc781a0263f := namespace.Description

	if var_5cc781a0263f != nil {
		var var_5cc781a0263f_mapped *structpb.Value

		var var_5cc781a0263f_err error
		var_5cc781a0263f_mapped, var_5cc781a0263f_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_5cc781a0263f)
		if var_5cc781a0263f_err != nil {
			panic(var_5cc781a0263f_err)
		}
		properties["description"] = var_5cc781a0263f_mapped
	}

	var_29fb57d9c1ec := namespace.Details

	if var_29fb57d9c1ec != nil {
		var var_29fb57d9c1ec_mapped *structpb.Value

		var var_29fb57d9c1ec_err error
		var_29fb57d9c1ec_mapped, var_29fb57d9c1ec_err = types.ByResourcePropertyType(model.ResourceProperty_OBJECT).Pack(*var_29fb57d9c1ec)
		if var_29fb57d9c1ec_err != nil {
			panic(var_29fb57d9c1ec_err)
		}
		properties["details"] = var_29fb57d9c1ec_mapped
	}

	var_c139be01d90f := namespace.SecurityConstraints

	if var_c139be01d90f != nil {
		var var_c139be01d90f_mapped *structpb.Value

		var var_c139be01d90f_l []*structpb.Value
		for _, value := range var_c139be01d90f {

			var_86a6f7183661 := value
			var var_86a6f7183661_mapped *structpb.Value

			var_86a6f7183661_mapped = structpb.NewStructValue(&structpb.Struct{Fields: SecurityConstraintMapperInstance.ToProperties(var_86a6f7183661)})

			var_c139be01d90f_l = append(var_c139be01d90f_l, var_86a6f7183661_mapped)
		}
		var_c139be01d90f_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_c139be01d90f_l})
		properties["securityConstraints"] = var_c139be01d90f_mapped
	}
	return properties
}

func (m *NamespaceMapper) FromProperties(properties map[string]*structpb.Value) *Namespace {
	var s = m.New()
	if properties["id"] != nil {

		var_16fb22310cea := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_16fb22310cea)

		if err != nil {
			panic(err)
		}

		var_16fb22310cea_mapped := new(uuid.UUID)
		*var_16fb22310cea_mapped = val.(uuid.UUID)

		s.Id = var_16fb22310cea_mapped
	}
	if properties["version"] != nil {

		var_b0137cf0447e := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_b0137cf0447e)

		if err != nil {
			panic(err)
		}

		var_b0137cf0447e_mapped := val.(int32)

		s.Version = var_b0137cf0447e_mapped
	}
	if properties["createdBy"] != nil {

		var_8f5b7a653b08 := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_8f5b7a653b08)

		if err != nil {
			panic(err)
		}

		var_8f5b7a653b08_mapped := new(string)
		*var_8f5b7a653b08_mapped = val.(string)

		s.CreatedBy = var_8f5b7a653b08_mapped
	}
	if properties["updatedBy"] != nil {

		var_96dfbc9fc564 := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_96dfbc9fc564)

		if err != nil {
			panic(err)
		}

		var_96dfbc9fc564_mapped := new(string)
		*var_96dfbc9fc564_mapped = val.(string)

		s.UpdatedBy = var_96dfbc9fc564_mapped
	}
	if properties["createdOn"] != nil {

		var_e1134479d4aa := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_e1134479d4aa)

		if err != nil {
			panic(err)
		}

		var_e1134479d4aa_mapped := new(time.Time)
		*var_e1134479d4aa_mapped = val.(time.Time)

		s.CreatedOn = var_e1134479d4aa_mapped
	}
	if properties["updatedOn"] != nil {

		var_0455ae90cc9a := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_0455ae90cc9a)

		if err != nil {
			panic(err)
		}

		var_0455ae90cc9a_mapped := new(time.Time)
		*var_0455ae90cc9a_mapped = val.(time.Time)

		s.UpdatedOn = var_0455ae90cc9a_mapped
	}
	if properties["name"] != nil {

		var_8fb05cf80971 := properties["name"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_8fb05cf80971)

		if err != nil {
			panic(err)
		}

		var_8fb05cf80971_mapped := val.(string)

		s.Name = var_8fb05cf80971_mapped
	}
	if properties["description"] != nil {

		var_8cd3ef33bf57 := properties["description"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_8cd3ef33bf57)

		if err != nil {
			panic(err)
		}

		var_8cd3ef33bf57_mapped := new(string)
		*var_8cd3ef33bf57_mapped = val.(string)

		s.Description = var_8cd3ef33bf57_mapped
	}
	if properties["details"] != nil {

		var_8a2b28f7c936 := properties["details"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_OBJECT).UnPack(var_8a2b28f7c936)

		if err != nil {
			panic(err)
		}

		var_8a2b28f7c936_mapped := new(unstructured.Unstructured)
		*var_8a2b28f7c936_mapped = val.(unstructured.Unstructured)

		s.Details = var_8a2b28f7c936_mapped
	}
	if properties["securityConstraints"] != nil {

		var_cf1cae8bd515 := properties["securityConstraints"]
		var_cf1cae8bd515_mapped := []*SecurityConstraint{}
		for _, v := range var_cf1cae8bd515.GetListValue().Values {

			var_02fb93acba40 := v
			var_02fb93acba40_mapped := SecurityConstraintMapperInstance.FromProperties(var_02fb93acba40.GetStructValue().Fields)

			var_cf1cae8bd515_mapped = append(var_cf1cae8bd515_mapped, var_02fb93acba40_mapped)
		}

		s.SecurityConstraints = var_cf1cae8bd515_mapped
	}
	return s
}
