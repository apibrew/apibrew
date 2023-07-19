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

	var_e0a88246e815 := namespace.Id

	if var_e0a88246e815 != nil {
		var var_e0a88246e815_mapped *structpb.Value

		var var_e0a88246e815_err error
		var_e0a88246e815_mapped, var_e0a88246e815_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_e0a88246e815)
		if var_e0a88246e815_err != nil {
			panic(var_e0a88246e815_err)
		}
		properties["id"] = var_e0a88246e815_mapped
	}

	var_5ab6a3057b58 := namespace.Version

	var var_5ab6a3057b58_mapped *structpb.Value

	var var_5ab6a3057b58_err error
	var_5ab6a3057b58_mapped, var_5ab6a3057b58_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(var_5ab6a3057b58)
	if var_5ab6a3057b58_err != nil {
		panic(var_5ab6a3057b58_err)
	}
	properties["version"] = var_5ab6a3057b58_mapped

	var_1a391c705929 := namespace.CreatedBy

	if var_1a391c705929 != nil {
		var var_1a391c705929_mapped *structpb.Value

		var var_1a391c705929_err error
		var_1a391c705929_mapped, var_1a391c705929_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_1a391c705929)
		if var_1a391c705929_err != nil {
			panic(var_1a391c705929_err)
		}
		properties["createdBy"] = var_1a391c705929_mapped
	}

	var_92b2222e316e := namespace.UpdatedBy

	if var_92b2222e316e != nil {
		var var_92b2222e316e_mapped *structpb.Value

		var var_92b2222e316e_err error
		var_92b2222e316e_mapped, var_92b2222e316e_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_92b2222e316e)
		if var_92b2222e316e_err != nil {
			panic(var_92b2222e316e_err)
		}
		properties["updatedBy"] = var_92b2222e316e_mapped
	}

	var_93239db2a1ae := namespace.CreatedOn

	if var_93239db2a1ae != nil {
		var var_93239db2a1ae_mapped *structpb.Value

		var var_93239db2a1ae_err error
		var_93239db2a1ae_mapped, var_93239db2a1ae_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_93239db2a1ae)
		if var_93239db2a1ae_err != nil {
			panic(var_93239db2a1ae_err)
		}
		properties["createdOn"] = var_93239db2a1ae_mapped
	}

	var_5ab0fbe80439 := namespace.UpdatedOn

	if var_5ab0fbe80439 != nil {
		var var_5ab0fbe80439_mapped *structpb.Value

		var var_5ab0fbe80439_err error
		var_5ab0fbe80439_mapped, var_5ab0fbe80439_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_5ab0fbe80439)
		if var_5ab0fbe80439_err != nil {
			panic(var_5ab0fbe80439_err)
		}
		properties["updatedOn"] = var_5ab0fbe80439_mapped
	}

	var_8700f156eee4 := namespace.Name

	var var_8700f156eee4_mapped *structpb.Value

	var var_8700f156eee4_err error
	var_8700f156eee4_mapped, var_8700f156eee4_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_8700f156eee4)
	if var_8700f156eee4_err != nil {
		panic(var_8700f156eee4_err)
	}
	properties["name"] = var_8700f156eee4_mapped

	var_298ab2f603d6 := namespace.Description

	if var_298ab2f603d6 != nil {
		var var_298ab2f603d6_mapped *structpb.Value

		var var_298ab2f603d6_err error
		var_298ab2f603d6_mapped, var_298ab2f603d6_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_298ab2f603d6)
		if var_298ab2f603d6_err != nil {
			panic(var_298ab2f603d6_err)
		}
		properties["description"] = var_298ab2f603d6_mapped
	}

	var_038845acaeb6 := namespace.Details

	if var_038845acaeb6 != nil {
		var var_038845acaeb6_mapped *structpb.Value

		var var_038845acaeb6_err error
		var_038845acaeb6_mapped, var_038845acaeb6_err = types.ByResourcePropertyType(model.ResourceProperty_OBJECT).Pack(*var_038845acaeb6)
		if var_038845acaeb6_err != nil {
			panic(var_038845acaeb6_err)
		}
		properties["details"] = var_038845acaeb6_mapped
	}

	var_af35b272552a := namespace.SecurityConstraints

	if var_af35b272552a != nil {
		var var_af35b272552a_mapped *structpb.Value

		var var_af35b272552a_l []*structpb.Value
		for _, value := range var_af35b272552a {

			var_b99b51629834 := value
			var var_b99b51629834_mapped *structpb.Value

			var_b99b51629834_mapped = structpb.NewStructValue(&structpb.Struct{Fields: SecurityConstraintMapperInstance.ToProperties(var_b99b51629834)})

			var_af35b272552a_l = append(var_af35b272552a_l, var_b99b51629834_mapped)
		}
		var_af35b272552a_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_af35b272552a_l})
		properties["securityConstraints"] = var_af35b272552a_mapped
	}
	return properties
}

func (m *NamespaceMapper) FromProperties(properties map[string]*structpb.Value) *Namespace {
	var s = m.New()
	if properties["id"] != nil {

		var_1b72428826c3 := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_1b72428826c3)

		if err != nil {
			panic(err)
		}

		var_1b72428826c3_mapped := new(uuid.UUID)
		*var_1b72428826c3_mapped = val.(uuid.UUID)

		s.Id = var_1b72428826c3_mapped
	}
	if properties["version"] != nil {

		var_3865f5b19fb9 := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_3865f5b19fb9)

		if err != nil {
			panic(err)
		}

		var_3865f5b19fb9_mapped := val.(int32)

		s.Version = var_3865f5b19fb9_mapped
	}
	if properties["createdBy"] != nil {

		var_448e74fbd40b := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_448e74fbd40b)

		if err != nil {
			panic(err)
		}

		var_448e74fbd40b_mapped := new(string)
		*var_448e74fbd40b_mapped = val.(string)

		s.CreatedBy = var_448e74fbd40b_mapped
	}
	if properties["updatedBy"] != nil {

		var_a079a3623269 := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_a079a3623269)

		if err != nil {
			panic(err)
		}

		var_a079a3623269_mapped := new(string)
		*var_a079a3623269_mapped = val.(string)

		s.UpdatedBy = var_a079a3623269_mapped
	}
	if properties["createdOn"] != nil {

		var_db9defb81796 := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_db9defb81796)

		if err != nil {
			panic(err)
		}

		var_db9defb81796_mapped := new(time.Time)
		*var_db9defb81796_mapped = val.(time.Time)

		s.CreatedOn = var_db9defb81796_mapped
	}
	if properties["updatedOn"] != nil {

		var_96798f342a99 := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_96798f342a99)

		if err != nil {
			panic(err)
		}

		var_96798f342a99_mapped := new(time.Time)
		*var_96798f342a99_mapped = val.(time.Time)

		s.UpdatedOn = var_96798f342a99_mapped
	}
	if properties["name"] != nil {

		var_8bd3ee210468 := properties["name"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_8bd3ee210468)

		if err != nil {
			panic(err)
		}

		var_8bd3ee210468_mapped := val.(string)

		s.Name = var_8bd3ee210468_mapped
	}
	if properties["description"] != nil {

		var_af3f3353e4ab := properties["description"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_af3f3353e4ab)

		if err != nil {
			panic(err)
		}

		var_af3f3353e4ab_mapped := new(string)
		*var_af3f3353e4ab_mapped = val.(string)

		s.Description = var_af3f3353e4ab_mapped
	}
	if properties["details"] != nil {

		var_cf83a3437af1 := properties["details"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_OBJECT).UnPack(var_cf83a3437af1)

		if err != nil {
			panic(err)
		}

		var_cf83a3437af1_mapped := new(unstructured.Unstructured)
		*var_cf83a3437af1_mapped = val.(unstructured.Unstructured)

		s.Details = var_cf83a3437af1_mapped
	}
	if properties["securityConstraints"] != nil {

		var_4c5ac1b78416 := properties["securityConstraints"]
		var_4c5ac1b78416_mapped := []*SecurityConstraint{}
		for _, v := range var_4c5ac1b78416.GetListValue().Values {

			var_77d52377fca8 := v
			var_77d52377fca8_mapped := SecurityConstraintMapperInstance.FromProperties(var_77d52377fca8.GetStructValue().Fields)

			var_4c5ac1b78416_mapped = append(var_4c5ac1b78416_mapped, var_77d52377fca8_mapped)
		}

		s.SecurityConstraints = var_4c5ac1b78416_mapped
	}
	return s
}
