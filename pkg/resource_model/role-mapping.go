package resource_model

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/types"
	"google.golang.org/protobuf/types/known/structpb"
)

import "github.com/google/uuid"
import "time"
import "github.com/apibrew/apibrew/pkg/formats/unstructured"

type RoleMapper struct {
}

func NewRoleMapper() *RoleMapper {
	return &RoleMapper{}
}

var RoleMapperInstance = NewRoleMapper()

func (m *RoleMapper) New() *Role {
	return &Role{}
}

func (m *RoleMapper) ToRecord(role *Role) *model.Record {
	var rec = &model.Record{}
	rec.Properties = m.ToProperties(role)

	if role.Id != nil {
		rec.Id = role.Id.String()
	}

	return rec
}

func (m *RoleMapper) FromRecord(record *model.Record) *Role {
	return m.FromProperties(record.Properties)
}

func (m *RoleMapper) ToProperties(role *Role) map[string]*structpb.Value {
	var properties = make(map[string]*structpb.Value)

	var_1be9f42f2b37 := role.Id

	if var_1be9f42f2b37 != nil {
		var var_1be9f42f2b37_mapped *structpb.Value

		var var_1be9f42f2b37_err error
		var_1be9f42f2b37_mapped, var_1be9f42f2b37_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_1be9f42f2b37)
		if var_1be9f42f2b37_err != nil {
			panic(var_1be9f42f2b37_err)
		}
		properties["id"] = var_1be9f42f2b37_mapped
	}

	var_64293df864ca := role.Version

	var var_64293df864ca_mapped *structpb.Value

	var var_64293df864ca_err error
	var_64293df864ca_mapped, var_64293df864ca_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(var_64293df864ca)
	if var_64293df864ca_err != nil {
		panic(var_64293df864ca_err)
	}
	properties["version"] = var_64293df864ca_mapped

	var_1e17d4989880 := role.CreatedBy

	if var_1e17d4989880 != nil {
		var var_1e17d4989880_mapped *structpb.Value

		var var_1e17d4989880_err error
		var_1e17d4989880_mapped, var_1e17d4989880_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_1e17d4989880)
		if var_1e17d4989880_err != nil {
			panic(var_1e17d4989880_err)
		}
		properties["createdBy"] = var_1e17d4989880_mapped
	}

	var_6d21c6c911b7 := role.UpdatedBy

	if var_6d21c6c911b7 != nil {
		var var_6d21c6c911b7_mapped *structpb.Value

		var var_6d21c6c911b7_err error
		var_6d21c6c911b7_mapped, var_6d21c6c911b7_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_6d21c6c911b7)
		if var_6d21c6c911b7_err != nil {
			panic(var_6d21c6c911b7_err)
		}
		properties["updatedBy"] = var_6d21c6c911b7_mapped
	}

	var_5464b7646df6 := role.CreatedOn

	if var_5464b7646df6 != nil {
		var var_5464b7646df6_mapped *structpb.Value

		var var_5464b7646df6_err error
		var_5464b7646df6_mapped, var_5464b7646df6_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_5464b7646df6)
		if var_5464b7646df6_err != nil {
			panic(var_5464b7646df6_err)
		}
		properties["createdOn"] = var_5464b7646df6_mapped
	}

	var_61a6eba184c7 := role.UpdatedOn

	if var_61a6eba184c7 != nil {
		var var_61a6eba184c7_mapped *structpb.Value

		var var_61a6eba184c7_err error
		var_61a6eba184c7_mapped, var_61a6eba184c7_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_61a6eba184c7)
		if var_61a6eba184c7_err != nil {
			panic(var_61a6eba184c7_err)
		}
		properties["updatedOn"] = var_61a6eba184c7_mapped
	}

	var_b9062fbba026 := role.Name

	var var_b9062fbba026_mapped *structpb.Value

	var var_b9062fbba026_err error
	var_b9062fbba026_mapped, var_b9062fbba026_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_b9062fbba026)
	if var_b9062fbba026_err != nil {
		panic(var_b9062fbba026_err)
	}
	properties["name"] = var_b9062fbba026_mapped

	var_53b23b20da07 := role.SecurityConstraints

	if var_53b23b20da07 != nil {
		var var_53b23b20da07_mapped *structpb.Value

		var var_53b23b20da07_l []*structpb.Value
		for _, value := range var_53b23b20da07 {

			var_7d55a26579b2 := value
			var var_7d55a26579b2_mapped *structpb.Value

			var_7d55a26579b2_mapped = structpb.NewStructValue(&structpb.Struct{Fields: SecurityConstraintMapperInstance.ToProperties(var_7d55a26579b2)})

			var_53b23b20da07_l = append(var_53b23b20da07_l, var_7d55a26579b2_mapped)
		}
		var_53b23b20da07_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_53b23b20da07_l})
		properties["securityConstraints"] = var_53b23b20da07_mapped
	}

	var_06b118de9b24 := role.Details

	if var_06b118de9b24 != nil {
		var var_06b118de9b24_mapped *structpb.Value

		var var_06b118de9b24_err error
		var_06b118de9b24_mapped, var_06b118de9b24_err = types.ByResourcePropertyType(model.ResourceProperty_OBJECT).Pack(*var_06b118de9b24)
		if var_06b118de9b24_err != nil {
			panic(var_06b118de9b24_err)
		}
		properties["details"] = var_06b118de9b24_mapped
	}
	return properties
}

func (m *RoleMapper) FromProperties(properties map[string]*structpb.Value) *Role {
	var s = m.New()
	if properties["id"] != nil {

		var_7fc7f2c4e680 := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_7fc7f2c4e680)

		if err != nil {
			panic(err)
		}

		var_7fc7f2c4e680_mapped := new(uuid.UUID)
		*var_7fc7f2c4e680_mapped = val.(uuid.UUID)

		s.Id = var_7fc7f2c4e680_mapped
	}
	if properties["version"] != nil {

		var_d21e02a52f99 := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_d21e02a52f99)

		if err != nil {
			panic(err)
		}

		var_d21e02a52f99_mapped := val.(int32)

		s.Version = var_d21e02a52f99_mapped
	}
	if properties["createdBy"] != nil {

		var_c9821da108c4 := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_c9821da108c4)

		if err != nil {
			panic(err)
		}

		var_c9821da108c4_mapped := new(string)
		*var_c9821da108c4_mapped = val.(string)

		s.CreatedBy = var_c9821da108c4_mapped
	}
	if properties["updatedBy"] != nil {

		var_0a1414d448ba := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_0a1414d448ba)

		if err != nil {
			panic(err)
		}

		var_0a1414d448ba_mapped := new(string)
		*var_0a1414d448ba_mapped = val.(string)

		s.UpdatedBy = var_0a1414d448ba_mapped
	}
	if properties["createdOn"] != nil {

		var_8a06a461f73d := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_8a06a461f73d)

		if err != nil {
			panic(err)
		}

		var_8a06a461f73d_mapped := new(time.Time)
		*var_8a06a461f73d_mapped = val.(time.Time)

		s.CreatedOn = var_8a06a461f73d_mapped
	}
	if properties["updatedOn"] != nil {

		var_bb6631a86655 := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_bb6631a86655)

		if err != nil {
			panic(err)
		}

		var_bb6631a86655_mapped := new(time.Time)
		*var_bb6631a86655_mapped = val.(time.Time)

		s.UpdatedOn = var_bb6631a86655_mapped
	}
	if properties["name"] != nil {

		var_b456e6017dfb := properties["name"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_b456e6017dfb)

		if err != nil {
			panic(err)
		}

		var_b456e6017dfb_mapped := val.(string)

		s.Name = var_b456e6017dfb_mapped
	}
	if properties["securityConstraints"] != nil {

		var_19c86411096d := properties["securityConstraints"]
		var_19c86411096d_mapped := []*SecurityConstraint{}
		for _, v := range var_19c86411096d.GetListValue().Values {

			var_a99eb4f35770 := v
			var_a99eb4f35770_mapped := SecurityConstraintMapperInstance.FromProperties(var_a99eb4f35770.GetStructValue().Fields)

			var_19c86411096d_mapped = append(var_19c86411096d_mapped, var_a99eb4f35770_mapped)
		}

		s.SecurityConstraints = var_19c86411096d_mapped
	}
	if properties["details"] != nil {

		var_be00f876ea2a := properties["details"]
		var_be00f876ea2a_mapped := new(unstructured.Unstructured)
		*var_be00f876ea2a_mapped = unstructured.FromStructValue(var_be00f876ea2a.GetStructValue())

		s.Details = var_be00f876ea2a_mapped
	}
	return s
}
