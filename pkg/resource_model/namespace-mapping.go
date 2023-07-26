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

	var_61e83db65f2e := namespace.Id

	if var_61e83db65f2e != nil {
		var var_61e83db65f2e_mapped *structpb.Value

		var var_61e83db65f2e_err error
		var_61e83db65f2e_mapped, var_61e83db65f2e_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_61e83db65f2e)
		if var_61e83db65f2e_err != nil {
			panic(var_61e83db65f2e_err)
		}
		properties["id"] = var_61e83db65f2e_mapped
	}

	var_87f9fc9486ff := namespace.Version

	var var_87f9fc9486ff_mapped *structpb.Value

	var var_87f9fc9486ff_err error
	var_87f9fc9486ff_mapped, var_87f9fc9486ff_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(var_87f9fc9486ff)
	if var_87f9fc9486ff_err != nil {
		panic(var_87f9fc9486ff_err)
	}
	properties["version"] = var_87f9fc9486ff_mapped

	var_719420e33abd := namespace.CreatedBy

	if var_719420e33abd != nil {
		var var_719420e33abd_mapped *structpb.Value

		var var_719420e33abd_err error
		var_719420e33abd_mapped, var_719420e33abd_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_719420e33abd)
		if var_719420e33abd_err != nil {
			panic(var_719420e33abd_err)
		}
		properties["createdBy"] = var_719420e33abd_mapped
	}

	var_7d592c3227b7 := namespace.UpdatedBy

	if var_7d592c3227b7 != nil {
		var var_7d592c3227b7_mapped *structpb.Value

		var var_7d592c3227b7_err error
		var_7d592c3227b7_mapped, var_7d592c3227b7_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_7d592c3227b7)
		if var_7d592c3227b7_err != nil {
			panic(var_7d592c3227b7_err)
		}
		properties["updatedBy"] = var_7d592c3227b7_mapped
	}

	var_e399eda625c9 := namespace.CreatedOn

	if var_e399eda625c9 != nil {
		var var_e399eda625c9_mapped *structpb.Value

		var var_e399eda625c9_err error
		var_e399eda625c9_mapped, var_e399eda625c9_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_e399eda625c9)
		if var_e399eda625c9_err != nil {
			panic(var_e399eda625c9_err)
		}
		properties["createdOn"] = var_e399eda625c9_mapped
	}

	var_9c65e9f7e867 := namespace.UpdatedOn

	if var_9c65e9f7e867 != nil {
		var var_9c65e9f7e867_mapped *structpb.Value

		var var_9c65e9f7e867_err error
		var_9c65e9f7e867_mapped, var_9c65e9f7e867_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_9c65e9f7e867)
		if var_9c65e9f7e867_err != nil {
			panic(var_9c65e9f7e867_err)
		}
		properties["updatedOn"] = var_9c65e9f7e867_mapped
	}

	var_2797ebebdb50 := namespace.Name

	var var_2797ebebdb50_mapped *structpb.Value

	var var_2797ebebdb50_err error
	var_2797ebebdb50_mapped, var_2797ebebdb50_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_2797ebebdb50)
	if var_2797ebebdb50_err != nil {
		panic(var_2797ebebdb50_err)
	}
	properties["name"] = var_2797ebebdb50_mapped

	var_1b28c764a5dc := namespace.Description

	if var_1b28c764a5dc != nil {
		var var_1b28c764a5dc_mapped *structpb.Value

		var var_1b28c764a5dc_err error
		var_1b28c764a5dc_mapped, var_1b28c764a5dc_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_1b28c764a5dc)
		if var_1b28c764a5dc_err != nil {
			panic(var_1b28c764a5dc_err)
		}
		properties["description"] = var_1b28c764a5dc_mapped
	}

	var_38e5aab20757 := namespace.Details

	if var_38e5aab20757 != nil {
		var var_38e5aab20757_mapped *structpb.Value

		var var_38e5aab20757_err error
		var_38e5aab20757_mapped, var_38e5aab20757_err = types.ByResourcePropertyType(model.ResourceProperty_OBJECT).Pack(*var_38e5aab20757)
		if var_38e5aab20757_err != nil {
			panic(var_38e5aab20757_err)
		}
		properties["details"] = var_38e5aab20757_mapped
	}

	var_666b9aff17dd := namespace.SecurityConstraints

	if var_666b9aff17dd != nil {
		var var_666b9aff17dd_mapped *structpb.Value

		var var_666b9aff17dd_l []*structpb.Value
		for _, value := range var_666b9aff17dd {

			var_4c54ec5454df := value
			var var_4c54ec5454df_mapped *structpb.Value

			var_4c54ec5454df_mapped = structpb.NewStructValue(&structpb.Struct{Fields: SecurityConstraintMapperInstance.ToProperties(var_4c54ec5454df)})

			var_666b9aff17dd_l = append(var_666b9aff17dd_l, var_4c54ec5454df_mapped)
		}
		var_666b9aff17dd_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_666b9aff17dd_l})
		properties["securityConstraints"] = var_666b9aff17dd_mapped
	}
	return properties
}

func (m *NamespaceMapper) FromProperties(properties map[string]*structpb.Value) *Namespace {
	var s = m.New()
	if properties["id"] != nil {

		var_a33edbd84c11 := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_a33edbd84c11)

		if err != nil {
			panic(err)
		}

		var_a33edbd84c11_mapped := new(uuid.UUID)
		*var_a33edbd84c11_mapped = val.(uuid.UUID)

		s.Id = var_a33edbd84c11_mapped
	}
	if properties["version"] != nil {

		var_b7e20a132345 := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_b7e20a132345)

		if err != nil {
			panic(err)
		}

		var_b7e20a132345_mapped := val.(int32)

		s.Version = var_b7e20a132345_mapped
	}
	if properties["createdBy"] != nil {

		var_05b9f73a67a4 := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_05b9f73a67a4)

		if err != nil {
			panic(err)
		}

		var_05b9f73a67a4_mapped := new(string)
		*var_05b9f73a67a4_mapped = val.(string)

		s.CreatedBy = var_05b9f73a67a4_mapped
	}
	if properties["updatedBy"] != nil {

		var_156529a7b0da := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_156529a7b0da)

		if err != nil {
			panic(err)
		}

		var_156529a7b0da_mapped := new(string)
		*var_156529a7b0da_mapped = val.(string)

		s.UpdatedBy = var_156529a7b0da_mapped
	}
	if properties["createdOn"] != nil {

		var_d8c5648046bb := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_d8c5648046bb)

		if err != nil {
			panic(err)
		}

		var_d8c5648046bb_mapped := new(time.Time)
		*var_d8c5648046bb_mapped = val.(time.Time)

		s.CreatedOn = var_d8c5648046bb_mapped
	}
	if properties["updatedOn"] != nil {

		var_d40f2aceb314 := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_d40f2aceb314)

		if err != nil {
			panic(err)
		}

		var_d40f2aceb314_mapped := new(time.Time)
		*var_d40f2aceb314_mapped = val.(time.Time)

		s.UpdatedOn = var_d40f2aceb314_mapped
	}
	if properties["name"] != nil {

		var_79e923195f73 := properties["name"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_79e923195f73)

		if err != nil {
			panic(err)
		}

		var_79e923195f73_mapped := val.(string)

		s.Name = var_79e923195f73_mapped
	}
	if properties["description"] != nil {

		var_299e14ba2b59 := properties["description"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_299e14ba2b59)

		if err != nil {
			panic(err)
		}

		var_299e14ba2b59_mapped := new(string)
		*var_299e14ba2b59_mapped = val.(string)

		s.Description = var_299e14ba2b59_mapped
	}
	if properties["details"] != nil {

		var_35e98951658b := properties["details"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_OBJECT).UnPack(var_35e98951658b)

		if err != nil {
			panic(err)
		}

		var_35e98951658b_mapped := new(unstructured.Unstructured)
		*var_35e98951658b_mapped = val.(unstructured.Unstructured)

		s.Details = var_35e98951658b_mapped
	}
	if properties["securityConstraints"] != nil {

		var_94d9cb205f9c := properties["securityConstraints"]
		var_94d9cb205f9c_mapped := []*SecurityConstraint{}
		for _, v := range var_94d9cb205f9c.GetListValue().Values {

			var_d1a020cab833 := v
			var_d1a020cab833_mapped := SecurityConstraintMapperInstance.FromProperties(var_d1a020cab833.GetStructValue().Fields)

			var_94d9cb205f9c_mapped = append(var_94d9cb205f9c_mapped, var_d1a020cab833_mapped)
		}

		s.SecurityConstraints = var_94d9cb205f9c_mapped
	}
	return s
}
