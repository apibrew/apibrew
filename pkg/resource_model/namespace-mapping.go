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

	var_a3fed6e597e9 := namespace.Id

	if var_a3fed6e597e9 != nil {
		var var_a3fed6e597e9_mapped *structpb.Value

		var var_a3fed6e597e9_err error
		var_a3fed6e597e9_mapped, var_a3fed6e597e9_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_a3fed6e597e9)
		if var_a3fed6e597e9_err != nil {
			panic(var_a3fed6e597e9_err)
		}
		properties["id"] = var_a3fed6e597e9_mapped
	}

	var_8953580f2d54 := namespace.Version

	var var_8953580f2d54_mapped *structpb.Value

	var var_8953580f2d54_err error
	var_8953580f2d54_mapped, var_8953580f2d54_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(var_8953580f2d54)
	if var_8953580f2d54_err != nil {
		panic(var_8953580f2d54_err)
	}
	properties["version"] = var_8953580f2d54_mapped

	var_07275d713031 := namespace.CreatedBy

	if var_07275d713031 != nil {
		var var_07275d713031_mapped *structpb.Value

		var var_07275d713031_err error
		var_07275d713031_mapped, var_07275d713031_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_07275d713031)
		if var_07275d713031_err != nil {
			panic(var_07275d713031_err)
		}
		properties["createdBy"] = var_07275d713031_mapped
	}

	var_21e7efb47c9c := namespace.UpdatedBy

	if var_21e7efb47c9c != nil {
		var var_21e7efb47c9c_mapped *structpb.Value

		var var_21e7efb47c9c_err error
		var_21e7efb47c9c_mapped, var_21e7efb47c9c_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_21e7efb47c9c)
		if var_21e7efb47c9c_err != nil {
			panic(var_21e7efb47c9c_err)
		}
		properties["updatedBy"] = var_21e7efb47c9c_mapped
	}

	var_f418c1102937 := namespace.CreatedOn

	if var_f418c1102937 != nil {
		var var_f418c1102937_mapped *structpb.Value

		var var_f418c1102937_err error
		var_f418c1102937_mapped, var_f418c1102937_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_f418c1102937)
		if var_f418c1102937_err != nil {
			panic(var_f418c1102937_err)
		}
		properties["createdOn"] = var_f418c1102937_mapped
	}

	var_d98d4d27b815 := namespace.UpdatedOn

	if var_d98d4d27b815 != nil {
		var var_d98d4d27b815_mapped *structpb.Value

		var var_d98d4d27b815_err error
		var_d98d4d27b815_mapped, var_d98d4d27b815_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_d98d4d27b815)
		if var_d98d4d27b815_err != nil {
			panic(var_d98d4d27b815_err)
		}
		properties["updatedOn"] = var_d98d4d27b815_mapped
	}

	var_854c9390a37d := namespace.Name

	var var_854c9390a37d_mapped *structpb.Value

	var var_854c9390a37d_err error
	var_854c9390a37d_mapped, var_854c9390a37d_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_854c9390a37d)
	if var_854c9390a37d_err != nil {
		panic(var_854c9390a37d_err)
	}
	properties["name"] = var_854c9390a37d_mapped

	var_5d3411884d85 := namespace.Description

	if var_5d3411884d85 != nil {
		var var_5d3411884d85_mapped *structpb.Value

		var var_5d3411884d85_err error
		var_5d3411884d85_mapped, var_5d3411884d85_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_5d3411884d85)
		if var_5d3411884d85_err != nil {
			panic(var_5d3411884d85_err)
		}
		properties["description"] = var_5d3411884d85_mapped
	}

	var_f4517c0d460a := namespace.Details

	if var_f4517c0d460a != nil {
		var var_f4517c0d460a_mapped *structpb.Value

		var var_f4517c0d460a_err error
		var_f4517c0d460a_mapped, var_f4517c0d460a_err = types.ByResourcePropertyType(model.ResourceProperty_OBJECT).Pack(*var_f4517c0d460a)
		if var_f4517c0d460a_err != nil {
			panic(var_f4517c0d460a_err)
		}
		properties["details"] = var_f4517c0d460a_mapped
	}

	var_7546f9565ae1 := namespace.SecurityConstraints

	if var_7546f9565ae1 != nil {
		var var_7546f9565ae1_mapped *structpb.Value

		var var_7546f9565ae1_l []*structpb.Value
		for _, value := range var_7546f9565ae1 {

			var_f9d3739158c5 := value
			var var_f9d3739158c5_mapped *structpb.Value

			var_f9d3739158c5_mapped = structpb.NewStructValue(&structpb.Struct{Fields: SecurityConstraintMapperInstance.ToProperties(var_f9d3739158c5)})

			var_7546f9565ae1_l = append(var_7546f9565ae1_l, var_f9d3739158c5_mapped)
		}
		var_7546f9565ae1_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_7546f9565ae1_l})
		properties["securityConstraints"] = var_7546f9565ae1_mapped
	}
	return properties
}

func (m *NamespaceMapper) FromProperties(properties map[string]*structpb.Value) *Namespace {
	var s = m.New()
	if properties["id"] != nil {

		var_6e9dbedc8165 := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_6e9dbedc8165)

		if err != nil {
			panic(err)
		}

		var_6e9dbedc8165_mapped := new(uuid.UUID)
		*var_6e9dbedc8165_mapped = val.(uuid.UUID)

		s.Id = var_6e9dbedc8165_mapped
	}
	if properties["version"] != nil {

		var_ae646c4d652e := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_ae646c4d652e)

		if err != nil {
			panic(err)
		}

		var_ae646c4d652e_mapped := val.(int32)

		s.Version = var_ae646c4d652e_mapped
	}
	if properties["createdBy"] != nil {

		var_4b9bc77cc5e9 := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_4b9bc77cc5e9)

		if err != nil {
			panic(err)
		}

		var_4b9bc77cc5e9_mapped := new(string)
		*var_4b9bc77cc5e9_mapped = val.(string)

		s.CreatedBy = var_4b9bc77cc5e9_mapped
	}
	if properties["updatedBy"] != nil {

		var_576f7537d187 := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_576f7537d187)

		if err != nil {
			panic(err)
		}

		var_576f7537d187_mapped := new(string)
		*var_576f7537d187_mapped = val.(string)

		s.UpdatedBy = var_576f7537d187_mapped
	}
	if properties["createdOn"] != nil {

		var_0bdf884c8764 := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_0bdf884c8764)

		if err != nil {
			panic(err)
		}

		var_0bdf884c8764_mapped := new(time.Time)
		*var_0bdf884c8764_mapped = val.(time.Time)

		s.CreatedOn = var_0bdf884c8764_mapped
	}
	if properties["updatedOn"] != nil {

		var_1f7d23c42ff6 := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_1f7d23c42ff6)

		if err != nil {
			panic(err)
		}

		var_1f7d23c42ff6_mapped := new(time.Time)
		*var_1f7d23c42ff6_mapped = val.(time.Time)

		s.UpdatedOn = var_1f7d23c42ff6_mapped
	}
	if properties["name"] != nil {

		var_1ebe5372077b := properties["name"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_1ebe5372077b)

		if err != nil {
			panic(err)
		}

		var_1ebe5372077b_mapped := val.(string)

		s.Name = var_1ebe5372077b_mapped
	}
	if properties["description"] != nil {

		var_605e4cd42ed1 := properties["description"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_605e4cd42ed1)

		if err != nil {
			panic(err)
		}

		var_605e4cd42ed1_mapped := new(string)
		*var_605e4cd42ed1_mapped = val.(string)

		s.Description = var_605e4cd42ed1_mapped
	}
	if properties["details"] != nil {

		var_9ee746d37c01 := properties["details"]
		var_9ee746d37c01_mapped := new(unstructured.Unstructured)
		*var_9ee746d37c01_mapped = unstructured.FromStructValue(var_9ee746d37c01.GetStructValue())

		s.Details = var_9ee746d37c01_mapped
	}
	if properties["securityConstraints"] != nil {

		var_da5237daf3da := properties["securityConstraints"]
		var_da5237daf3da_mapped := []*SecurityConstraint{}
		for _, v := range var_da5237daf3da.GetListValue().Values {

			var_221f4ffe4d8f := v
			var_221f4ffe4d8f_mapped := SecurityConstraintMapperInstance.FromProperties(var_221f4ffe4d8f.GetStructValue().Fields)

			var_da5237daf3da_mapped = append(var_da5237daf3da_mapped, var_221f4ffe4d8f_mapped)
		}

		s.SecurityConstraints = var_da5237daf3da_mapped
	}
	return s
}
