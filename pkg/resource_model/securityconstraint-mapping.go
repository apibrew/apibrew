package resource_model

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/types"
	"google.golang.org/protobuf/types/known/structpb"
)

import "github.com/google/uuid"
import "time"
import "github.com/apibrew/apibrew/pkg/formats/unstructured"

type SecurityConstraintMapper struct {
}

func NewSecurityConstraintMapper() *SecurityConstraintMapper {
	return &SecurityConstraintMapper{}
}

var SecurityConstraintMapperInstance = NewSecurityConstraintMapper()

func (m *SecurityConstraintMapper) New() *SecurityConstraint {
	return &SecurityConstraint{}
}

func (m *SecurityConstraintMapper) ToRecord(securityConstraint *SecurityConstraint) *model.Record {
	var rec = &model.Record{}
	rec.Properties = m.ToProperties(securityConstraint)

	if securityConstraint.Id != nil {
		rec.Id = securityConstraint.Id.String()
	}

	return rec
}

func (m *SecurityConstraintMapper) FromRecord(record *model.Record) *SecurityConstraint {
	return m.FromProperties(record.Properties)
}

func (m *SecurityConstraintMapper) ToProperties(securityConstraint *SecurityConstraint) map[string]*structpb.Value {
	var properties = make(map[string]*structpb.Value)

	var_495607012ecb := securityConstraint.Id

	if var_495607012ecb != nil {
		var var_495607012ecb_mapped *structpb.Value

		var var_495607012ecb_err error
		var_495607012ecb_mapped, var_495607012ecb_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_495607012ecb)
		if var_495607012ecb_err != nil {
			panic(var_495607012ecb_err)
		}
		properties["id"] = var_495607012ecb_mapped
	}

	var_297528748f8c := securityConstraint.Version

	var var_297528748f8c_mapped *structpb.Value

	var var_297528748f8c_err error
	var_297528748f8c_mapped, var_297528748f8c_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(var_297528748f8c)
	if var_297528748f8c_err != nil {
		panic(var_297528748f8c_err)
	}
	properties["version"] = var_297528748f8c_mapped

	var_0e7f2ef9227d := securityConstraint.CreatedBy

	if var_0e7f2ef9227d != nil {
		var var_0e7f2ef9227d_mapped *structpb.Value

		var var_0e7f2ef9227d_err error
		var_0e7f2ef9227d_mapped, var_0e7f2ef9227d_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_0e7f2ef9227d)
		if var_0e7f2ef9227d_err != nil {
			panic(var_0e7f2ef9227d_err)
		}
		properties["createdBy"] = var_0e7f2ef9227d_mapped
	}

	var_f4de992beb17 := securityConstraint.UpdatedBy

	if var_f4de992beb17 != nil {
		var var_f4de992beb17_mapped *structpb.Value

		var var_f4de992beb17_err error
		var_f4de992beb17_mapped, var_f4de992beb17_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_f4de992beb17)
		if var_f4de992beb17_err != nil {
			panic(var_f4de992beb17_err)
		}
		properties["updatedBy"] = var_f4de992beb17_mapped
	}

	var_0360fc0471de := securityConstraint.CreatedOn

	if var_0360fc0471de != nil {
		var var_0360fc0471de_mapped *structpb.Value

		var var_0360fc0471de_err error
		var_0360fc0471de_mapped, var_0360fc0471de_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_0360fc0471de)
		if var_0360fc0471de_err != nil {
			panic(var_0360fc0471de_err)
		}
		properties["createdOn"] = var_0360fc0471de_mapped
	}

	var_8cb1ce104f60 := securityConstraint.UpdatedOn

	if var_8cb1ce104f60 != nil {
		var var_8cb1ce104f60_mapped *structpb.Value

		var var_8cb1ce104f60_err error
		var_8cb1ce104f60_mapped, var_8cb1ce104f60_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_8cb1ce104f60)
		if var_8cb1ce104f60_err != nil {
			panic(var_8cb1ce104f60_err)
		}
		properties["updatedOn"] = var_8cb1ce104f60_mapped
	}

	var_8947112460d5 := securityConstraint.Namespace

	var var_8947112460d5_mapped *structpb.Value

	var var_8947112460d5_err error
	var_8947112460d5_mapped, var_8947112460d5_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_8947112460d5)
	if var_8947112460d5_err != nil {
		panic(var_8947112460d5_err)
	}
	properties["namespace"] = var_8947112460d5_mapped

	var_8ddb9adf8265 := securityConstraint.Resource

	var var_8ddb9adf8265_mapped *structpb.Value

	var var_8ddb9adf8265_err error
	var_8ddb9adf8265_mapped, var_8ddb9adf8265_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_8ddb9adf8265)
	if var_8ddb9adf8265_err != nil {
		panic(var_8ddb9adf8265_err)
	}
	properties["resource"] = var_8ddb9adf8265_mapped

	var_a582b4a751f6 := securityConstraint.Property

	var var_a582b4a751f6_mapped *structpb.Value

	var var_a582b4a751f6_err error
	var_a582b4a751f6_mapped, var_a582b4a751f6_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_a582b4a751f6)
	if var_a582b4a751f6_err != nil {
		panic(var_a582b4a751f6_err)
	}
	properties["property"] = var_a582b4a751f6_mapped

	var_e3bef5d674bd := securityConstraint.PropertyValue

	if var_e3bef5d674bd != nil {
		var var_e3bef5d674bd_mapped *structpb.Value

		var var_e3bef5d674bd_err error
		var_e3bef5d674bd_mapped, var_e3bef5d674bd_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_e3bef5d674bd)
		if var_e3bef5d674bd_err != nil {
			panic(var_e3bef5d674bd_err)
		}
		properties["propertyValue"] = var_e3bef5d674bd_mapped
	}

	var_395593026cce := securityConstraint.PropertyMode

	if var_395593026cce != nil {
		var var_395593026cce_mapped *structpb.Value

		var var_395593026cce_err error
		var_395593026cce_mapped, var_395593026cce_err = types.ByResourcePropertyType(model.ResourceProperty_ENUM).Pack(string(*var_395593026cce))
		if var_395593026cce_err != nil {
			panic(var_395593026cce_err)
		}
		properties["propertyMode"] = var_395593026cce_mapped
	}

	var_0e9b2143bf11 := securityConstraint.Operation

	var var_0e9b2143bf11_mapped *structpb.Value

	var var_0e9b2143bf11_err error
	var_0e9b2143bf11_mapped, var_0e9b2143bf11_err = types.ByResourcePropertyType(model.ResourceProperty_ENUM).Pack(string(var_0e9b2143bf11))
	if var_0e9b2143bf11_err != nil {
		panic(var_0e9b2143bf11_err)
	}
	properties["operation"] = var_0e9b2143bf11_mapped

	var_01a366816c29 := securityConstraint.RecordIds

	if var_01a366816c29 != nil {
		var var_01a366816c29_mapped *structpb.Value

		var var_01a366816c29_l []*structpb.Value
		for _, value := range var_01a366816c29 {

			var_b5438987b769 := value
			var var_b5438987b769_mapped *structpb.Value

			var var_b5438987b769_err error
			var_b5438987b769_mapped, var_b5438987b769_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_b5438987b769)
			if var_b5438987b769_err != nil {
				panic(var_b5438987b769_err)
			}

			var_01a366816c29_l = append(var_01a366816c29_l, var_b5438987b769_mapped)
		}
		var_01a366816c29_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_01a366816c29_l})
		properties["recordIds"] = var_01a366816c29_mapped
	}

	var_1cd56297263c := securityConstraint.Before

	if var_1cd56297263c != nil {
		var var_1cd56297263c_mapped *structpb.Value

		var var_1cd56297263c_err error
		var_1cd56297263c_mapped, var_1cd56297263c_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_1cd56297263c)
		if var_1cd56297263c_err != nil {
			panic(var_1cd56297263c_err)
		}
		properties["before"] = var_1cd56297263c_mapped
	}

	var_5224b4a372ab := securityConstraint.After

	if var_5224b4a372ab != nil {
		var var_5224b4a372ab_mapped *structpb.Value

		var var_5224b4a372ab_err error
		var_5224b4a372ab_mapped, var_5224b4a372ab_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_5224b4a372ab)
		if var_5224b4a372ab_err != nil {
			panic(var_5224b4a372ab_err)
		}
		properties["after"] = var_5224b4a372ab_mapped
	}

	var_e4adaddd902b := securityConstraint.Username

	if var_e4adaddd902b != nil {
		var var_e4adaddd902b_mapped *structpb.Value

		var var_e4adaddd902b_err error
		var_e4adaddd902b_mapped, var_e4adaddd902b_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_e4adaddd902b)
		if var_e4adaddd902b_err != nil {
			panic(var_e4adaddd902b_err)
		}
		properties["username"] = var_e4adaddd902b_mapped
	}

	var_0e578704898e := securityConstraint.Role

	if var_0e578704898e != nil {
		var var_0e578704898e_mapped *structpb.Value

		var var_0e578704898e_err error
		var_0e578704898e_mapped, var_0e578704898e_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_0e578704898e)
		if var_0e578704898e_err != nil {
			panic(var_0e578704898e_err)
		}
		properties["role"] = var_0e578704898e_mapped
	}

	var_56d0cccf3e25 := securityConstraint.Permit

	var var_56d0cccf3e25_mapped *structpb.Value

	var var_56d0cccf3e25_err error
	var_56d0cccf3e25_mapped, var_56d0cccf3e25_err = types.ByResourcePropertyType(model.ResourceProperty_ENUM).Pack(string(var_56d0cccf3e25))
	if var_56d0cccf3e25_err != nil {
		panic(var_56d0cccf3e25_err)
	}
	properties["permit"] = var_56d0cccf3e25_mapped

	var_10abff803026 := securityConstraint.LocalFlags

	if var_10abff803026 != nil {
		var var_10abff803026_mapped *structpb.Value

		var var_10abff803026_err error
		var_10abff803026_mapped, var_10abff803026_err = types.ByResourcePropertyType(model.ResourceProperty_OBJECT).Pack(*var_10abff803026)
		if var_10abff803026_err != nil {
			panic(var_10abff803026_err)
		}
		properties["localFlags"] = var_10abff803026_mapped
	}
	return properties
}

func (m *SecurityConstraintMapper) FromProperties(properties map[string]*structpb.Value) *SecurityConstraint {
	var s = m.New()
	if properties["id"] != nil {

		var_321aa2d1b268 := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_321aa2d1b268)

		if err != nil {
			panic(err)
		}

		var_321aa2d1b268_mapped := new(uuid.UUID)
		*var_321aa2d1b268_mapped = val.(uuid.UUID)

		s.Id = var_321aa2d1b268_mapped
	}
	if properties["version"] != nil {

		var_9ebbdddfae9e := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_9ebbdddfae9e)

		if err != nil {
			panic(err)
		}

		var_9ebbdddfae9e_mapped := val.(int32)

		s.Version = var_9ebbdddfae9e_mapped
	}
	if properties["createdBy"] != nil {

		var_7f03e558d856 := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_7f03e558d856)

		if err != nil {
			panic(err)
		}

		var_7f03e558d856_mapped := new(string)
		*var_7f03e558d856_mapped = val.(string)

		s.CreatedBy = var_7f03e558d856_mapped
	}
	if properties["updatedBy"] != nil {

		var_517537a13178 := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_517537a13178)

		if err != nil {
			panic(err)
		}

		var_517537a13178_mapped := new(string)
		*var_517537a13178_mapped = val.(string)

		s.UpdatedBy = var_517537a13178_mapped
	}
	if properties["createdOn"] != nil {

		var_d8eca0e9435d := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_d8eca0e9435d)

		if err != nil {
			panic(err)
		}

		var_d8eca0e9435d_mapped := new(time.Time)
		*var_d8eca0e9435d_mapped = val.(time.Time)

		s.CreatedOn = var_d8eca0e9435d_mapped
	}
	if properties["updatedOn"] != nil {

		var_1de93ca9766f := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_1de93ca9766f)

		if err != nil {
			panic(err)
		}

		var_1de93ca9766f_mapped := new(time.Time)
		*var_1de93ca9766f_mapped = val.(time.Time)

		s.UpdatedOn = var_1de93ca9766f_mapped
	}
	if properties["namespace"] != nil {

		var_abf2ef19ea6f := properties["namespace"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_abf2ef19ea6f)

		if err != nil {
			panic(err)
		}

		var_abf2ef19ea6f_mapped := val.(string)

		s.Namespace = var_abf2ef19ea6f_mapped
	}
	if properties["resource"] != nil {

		var_98a1322b599d := properties["resource"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_98a1322b599d)

		if err != nil {
			panic(err)
		}

		var_98a1322b599d_mapped := val.(string)

		s.Resource = var_98a1322b599d_mapped
	}
	if properties["property"] != nil {

		var_2f680233e340 := properties["property"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_2f680233e340)

		if err != nil {
			panic(err)
		}

		var_2f680233e340_mapped := val.(string)

		s.Property = var_2f680233e340_mapped
	}
	if properties["propertyValue"] != nil {

		var_a7238cdccccf := properties["propertyValue"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_a7238cdccccf)

		if err != nil {
			panic(err)
		}

		var_a7238cdccccf_mapped := new(string)
		*var_a7238cdccccf_mapped = val.(string)

		s.PropertyValue = var_a7238cdccccf_mapped
	}
	if properties["propertyMode"] != nil {

		var_ea0adf7cd923 := properties["propertyMode"]
		var_ea0adf7cd923_mapped := new(SecurityConstraintPropertyMode)
		*var_ea0adf7cd923_mapped = (SecurityConstraintPropertyMode)(var_ea0adf7cd923.GetStringValue())

		s.PropertyMode = var_ea0adf7cd923_mapped
	}
	if properties["operation"] != nil {

		var_6a92b6d7d071 := properties["operation"]
		var_6a92b6d7d071_mapped := (SecurityConstraintOperation)(var_6a92b6d7d071.GetStringValue())

		s.Operation = var_6a92b6d7d071_mapped
	}
	if properties["recordIds"] != nil {

		var_1d082b1f7829 := properties["recordIds"]
		var_1d082b1f7829_mapped := []string{}
		for _, v := range var_1d082b1f7829.GetListValue().Values {

			var_f42329da6042 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_f42329da6042)

			if err != nil {
				panic(err)
			}

			var_f42329da6042_mapped := val.(string)

			var_1d082b1f7829_mapped = append(var_1d082b1f7829_mapped, var_f42329da6042_mapped)
		}

		s.RecordIds = var_1d082b1f7829_mapped
	}
	if properties["before"] != nil {

		var_7b1cf34e1cf9 := properties["before"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_7b1cf34e1cf9)

		if err != nil {
			panic(err)
		}

		var_7b1cf34e1cf9_mapped := new(time.Time)
		*var_7b1cf34e1cf9_mapped = val.(time.Time)

		s.Before = var_7b1cf34e1cf9_mapped
	}
	if properties["after"] != nil {

		var_f3e188874787 := properties["after"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_f3e188874787)

		if err != nil {
			panic(err)
		}

		var_f3e188874787_mapped := new(time.Time)
		*var_f3e188874787_mapped = val.(time.Time)

		s.After = var_f3e188874787_mapped
	}
	if properties["username"] != nil {

		var_a04827146b9e := properties["username"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_a04827146b9e)

		if err != nil {
			panic(err)
		}

		var_a04827146b9e_mapped := new(string)
		*var_a04827146b9e_mapped = val.(string)

		s.Username = var_a04827146b9e_mapped
	}
	if properties["role"] != nil {

		var_1d297062713e := properties["role"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_1d297062713e)

		if err != nil {
			panic(err)
		}

		var_1d297062713e_mapped := new(string)
		*var_1d297062713e_mapped = val.(string)

		s.Role = var_1d297062713e_mapped
	}
	if properties["permit"] != nil {

		var_f0c17dbc3978 := properties["permit"]
		var_f0c17dbc3978_mapped := (SecurityConstraintPermit)(var_f0c17dbc3978.GetStringValue())

		s.Permit = var_f0c17dbc3978_mapped
	}
	if properties["localFlags"] != nil {

		var_2ec3ccae1aa7 := properties["localFlags"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_OBJECT).UnPack(var_2ec3ccae1aa7)

		if err != nil {
			panic(err)
		}

		var_2ec3ccae1aa7_mapped := new(unstructured.Unstructured)
		*var_2ec3ccae1aa7_mapped = val.(unstructured.Unstructured)

		s.LocalFlags = var_2ec3ccae1aa7_mapped
	}
	return s
}
