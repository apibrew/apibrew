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

	var_d0e8e4f89ef3 := namespace.Id

	if var_d0e8e4f89ef3 != nil {
		var var_d0e8e4f89ef3_mapped *structpb.Value

		var var_d0e8e4f89ef3_err error
		var_d0e8e4f89ef3_mapped, var_d0e8e4f89ef3_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_d0e8e4f89ef3)
		if var_d0e8e4f89ef3_err != nil {
			panic(var_d0e8e4f89ef3_err)
		}
		properties["id"] = var_d0e8e4f89ef3_mapped
	}

	var_c8334d5f844f := namespace.Version

	var var_c8334d5f844f_mapped *structpb.Value

	var var_c8334d5f844f_err error
	var_c8334d5f844f_mapped, var_c8334d5f844f_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(var_c8334d5f844f)
	if var_c8334d5f844f_err != nil {
		panic(var_c8334d5f844f_err)
	}
	properties["version"] = var_c8334d5f844f_mapped

	var_c190ce5c05e7 := namespace.CreatedBy

	if var_c190ce5c05e7 != nil {
		var var_c190ce5c05e7_mapped *structpb.Value

		var var_c190ce5c05e7_err error
		var_c190ce5c05e7_mapped, var_c190ce5c05e7_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_c190ce5c05e7)
		if var_c190ce5c05e7_err != nil {
			panic(var_c190ce5c05e7_err)
		}
		properties["createdBy"] = var_c190ce5c05e7_mapped
	}

	var_4e6887b2bb6f := namespace.UpdatedBy

	if var_4e6887b2bb6f != nil {
		var var_4e6887b2bb6f_mapped *structpb.Value

		var var_4e6887b2bb6f_err error
		var_4e6887b2bb6f_mapped, var_4e6887b2bb6f_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_4e6887b2bb6f)
		if var_4e6887b2bb6f_err != nil {
			panic(var_4e6887b2bb6f_err)
		}
		properties["updatedBy"] = var_4e6887b2bb6f_mapped
	}

	var_c71ea5027902 := namespace.CreatedOn

	if var_c71ea5027902 != nil {
		var var_c71ea5027902_mapped *structpb.Value

		var var_c71ea5027902_err error
		var_c71ea5027902_mapped, var_c71ea5027902_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_c71ea5027902)
		if var_c71ea5027902_err != nil {
			panic(var_c71ea5027902_err)
		}
		properties["createdOn"] = var_c71ea5027902_mapped
	}

	var_a6e33b16dcf1 := namespace.UpdatedOn

	if var_a6e33b16dcf1 != nil {
		var var_a6e33b16dcf1_mapped *structpb.Value

		var var_a6e33b16dcf1_err error
		var_a6e33b16dcf1_mapped, var_a6e33b16dcf1_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_a6e33b16dcf1)
		if var_a6e33b16dcf1_err != nil {
			panic(var_a6e33b16dcf1_err)
		}
		properties["updatedOn"] = var_a6e33b16dcf1_mapped
	}

	var_90a1f111c7f4 := namespace.Name

	var var_90a1f111c7f4_mapped *structpb.Value

	var var_90a1f111c7f4_err error
	var_90a1f111c7f4_mapped, var_90a1f111c7f4_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_90a1f111c7f4)
	if var_90a1f111c7f4_err != nil {
		panic(var_90a1f111c7f4_err)
	}
	properties["name"] = var_90a1f111c7f4_mapped

	var_e9f806b9e84f := namespace.Description

	if var_e9f806b9e84f != nil {
		var var_e9f806b9e84f_mapped *structpb.Value

		var var_e9f806b9e84f_err error
		var_e9f806b9e84f_mapped, var_e9f806b9e84f_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_e9f806b9e84f)
		if var_e9f806b9e84f_err != nil {
			panic(var_e9f806b9e84f_err)
		}
		properties["description"] = var_e9f806b9e84f_mapped
	}

	var_5ccbb0f3761a := namespace.Details

	if var_5ccbb0f3761a != nil {
		var var_5ccbb0f3761a_mapped *structpb.Value

		var var_5ccbb0f3761a_err error
		var_5ccbb0f3761a_mapped, var_5ccbb0f3761a_err = types.ByResourcePropertyType(model.ResourceProperty_OBJECT).Pack(*var_5ccbb0f3761a)
		if var_5ccbb0f3761a_err != nil {
			panic(var_5ccbb0f3761a_err)
		}
		properties["details"] = var_5ccbb0f3761a_mapped
	}

	var_b83aed275cc2 := namespace.SecurityConstraints

	if var_b83aed275cc2 != nil {
		var var_b83aed275cc2_mapped *structpb.Value

		var var_b83aed275cc2_l []*structpb.Value
		for _, value := range var_b83aed275cc2 {

			var_3abf1da1b257 := value
			var var_3abf1da1b257_mapped *structpb.Value

			var_3abf1da1b257_mapped = structpb.NewStructValue(&structpb.Struct{Fields: SecurityConstraintMapperInstance.ToProperties(var_3abf1da1b257)})

			var_b83aed275cc2_l = append(var_b83aed275cc2_l, var_3abf1da1b257_mapped)
		}
		var_b83aed275cc2_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_b83aed275cc2_l})
		properties["securityConstraints"] = var_b83aed275cc2_mapped
	}
	return properties
}

func (m *NamespaceMapper) FromProperties(properties map[string]*structpb.Value) *Namespace {
	var s = m.New()
	if properties["id"] != nil {

		var_619fa514d937 := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_619fa514d937)

		if err != nil {
			panic(err)
		}

		var_619fa514d937_mapped := new(uuid.UUID)
		*var_619fa514d937_mapped = val.(uuid.UUID)

		s.Id = var_619fa514d937_mapped
	}
	if properties["version"] != nil {

		var_4be59e99ae4e := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_4be59e99ae4e)

		if err != nil {
			panic(err)
		}

		var_4be59e99ae4e_mapped := val.(int32)

		s.Version = var_4be59e99ae4e_mapped
	}
	if properties["createdBy"] != nil {

		var_15471c79482a := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_15471c79482a)

		if err != nil {
			panic(err)
		}

		var_15471c79482a_mapped := new(string)
		*var_15471c79482a_mapped = val.(string)

		s.CreatedBy = var_15471c79482a_mapped
	}
	if properties["updatedBy"] != nil {

		var_224f84cf1025 := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_224f84cf1025)

		if err != nil {
			panic(err)
		}

		var_224f84cf1025_mapped := new(string)
		*var_224f84cf1025_mapped = val.(string)

		s.UpdatedBy = var_224f84cf1025_mapped
	}
	if properties["createdOn"] != nil {

		var_2f8b1e7fff90 := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_2f8b1e7fff90)

		if err != nil {
			panic(err)
		}

		var_2f8b1e7fff90_mapped := new(time.Time)
		*var_2f8b1e7fff90_mapped = val.(time.Time)

		s.CreatedOn = var_2f8b1e7fff90_mapped
	}
	if properties["updatedOn"] != nil {

		var_1d652b1e2561 := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_1d652b1e2561)

		if err != nil {
			panic(err)
		}

		var_1d652b1e2561_mapped := new(time.Time)
		*var_1d652b1e2561_mapped = val.(time.Time)

		s.UpdatedOn = var_1d652b1e2561_mapped
	}
	if properties["name"] != nil {

		var_1ce3fbe0b193 := properties["name"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_1ce3fbe0b193)

		if err != nil {
			panic(err)
		}

		var_1ce3fbe0b193_mapped := val.(string)

		s.Name = var_1ce3fbe0b193_mapped
	}
	if properties["description"] != nil {

		var_cb8518fd8143 := properties["description"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_cb8518fd8143)

		if err != nil {
			panic(err)
		}

		var_cb8518fd8143_mapped := new(string)
		*var_cb8518fd8143_mapped = val.(string)

		s.Description = var_cb8518fd8143_mapped
	}
	if properties["details"] != nil {

		var_70d58e3eebb4 := properties["details"]
		var_70d58e3eebb4_mapped := new(unstructured.Unstructured)
		*var_70d58e3eebb4_mapped = unstructured.FromStructValue(var_70d58e3eebb4.GetStructValue())

		s.Details = var_70d58e3eebb4_mapped
	}
	if properties["securityConstraints"] != nil {

		var_0ba6f5ac57e5 := properties["securityConstraints"]
		var_0ba6f5ac57e5_mapped := []*SecurityConstraint{}
		for _, v := range var_0ba6f5ac57e5.GetListValue().Values {

			var_eae3ee996d5b := v
			var_eae3ee996d5b_mapped := SecurityConstraintMapperInstance.FromProperties(var_eae3ee996d5b.GetStructValue().Fields)

			var_0ba6f5ac57e5_mapped = append(var_0ba6f5ac57e5_mapped, var_eae3ee996d5b_mapped)
		}

		s.SecurityConstraints = var_0ba6f5ac57e5_mapped
	}
	return s
}
