package resource_model

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/types"
	"google.golang.org/protobuf/types/known/structpb"
)

import "github.com/google/uuid"
import "time"
import "github.com/apibrew/apibrew/pkg/formats/unstructured"

type UserMapper struct {
}

func NewUserMapper() *UserMapper {
	return &UserMapper{}
}

var UserMapperInstance = NewUserMapper()

func (m *UserMapper) New() *User {
	return &User{}
}

func (m *UserMapper) ToRecord(user *User) *model.Record {
	var rec = &model.Record{}
	rec.Properties = m.ToProperties(user)
	return rec
}

func (m *UserMapper) FromRecord(record *model.Record) *User {
	return m.FromProperties(record.Properties)
}

func (m *UserMapper) ToProperties(user *User) map[string]*structpb.Value {
	var properties = make(map[string]*structpb.Value)

	var_765df8e0789d := user.Id

	if var_765df8e0789d != nil {
		var var_765df8e0789d_mapped *structpb.Value

		var var_765df8e0789d_err error
		var_765df8e0789d_mapped, var_765df8e0789d_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_765df8e0789d)
		if var_765df8e0789d_err != nil {
			panic(var_765df8e0789d_err)
		}
		properties["id"] = var_765df8e0789d_mapped
	}

	var_afc019e8a6df := user.Version

	var var_afc019e8a6df_mapped *structpb.Value

	var var_afc019e8a6df_err error
	var_afc019e8a6df_mapped, var_afc019e8a6df_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(var_afc019e8a6df)
	if var_afc019e8a6df_err != nil {
		panic(var_afc019e8a6df_err)
	}
	properties["version"] = var_afc019e8a6df_mapped

	var_7c78575b0382 := user.CreatedBy

	if var_7c78575b0382 != nil {
		var var_7c78575b0382_mapped *structpb.Value

		var var_7c78575b0382_err error
		var_7c78575b0382_mapped, var_7c78575b0382_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_7c78575b0382)
		if var_7c78575b0382_err != nil {
			panic(var_7c78575b0382_err)
		}
		properties["createdBy"] = var_7c78575b0382_mapped
	}

	var_3c13dcbf5fe2 := user.UpdatedBy

	if var_3c13dcbf5fe2 != nil {
		var var_3c13dcbf5fe2_mapped *structpb.Value

		var var_3c13dcbf5fe2_err error
		var_3c13dcbf5fe2_mapped, var_3c13dcbf5fe2_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_3c13dcbf5fe2)
		if var_3c13dcbf5fe2_err != nil {
			panic(var_3c13dcbf5fe2_err)
		}
		properties["updatedBy"] = var_3c13dcbf5fe2_mapped
	}

	var_71853828af7f := user.CreatedOn

	if var_71853828af7f != nil {
		var var_71853828af7f_mapped *structpb.Value

		var var_71853828af7f_err error
		var_71853828af7f_mapped, var_71853828af7f_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_71853828af7f)
		if var_71853828af7f_err != nil {
			panic(var_71853828af7f_err)
		}
		properties["createdOn"] = var_71853828af7f_mapped
	}

	var_1cd278e2f264 := user.UpdatedOn

	if var_1cd278e2f264 != nil {
		var var_1cd278e2f264_mapped *structpb.Value

		var var_1cd278e2f264_err error
		var_1cd278e2f264_mapped, var_1cd278e2f264_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_1cd278e2f264)
		if var_1cd278e2f264_err != nil {
			panic(var_1cd278e2f264_err)
		}
		properties["updatedOn"] = var_1cd278e2f264_mapped
	}

	var_14ace208d203 := user.Username

	var var_14ace208d203_mapped *structpb.Value

	var var_14ace208d203_err error
	var_14ace208d203_mapped, var_14ace208d203_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_14ace208d203)
	if var_14ace208d203_err != nil {
		panic(var_14ace208d203_err)
	}
	properties["username"] = var_14ace208d203_mapped

	var_605a7462aebd := user.Password

	if var_605a7462aebd != nil {
		var var_605a7462aebd_mapped *structpb.Value

		var var_605a7462aebd_err error
		var_605a7462aebd_mapped, var_605a7462aebd_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_605a7462aebd)
		if var_605a7462aebd_err != nil {
			panic(var_605a7462aebd_err)
		}
		properties["password"] = var_605a7462aebd_mapped
	}

	var_e77498d50640 := user.Roles

	if var_e77498d50640 != nil {
		var var_e77498d50640_mapped *structpb.Value

		var var_e77498d50640_l []*structpb.Value
		for _, value := range var_e77498d50640 {

			var_bbab2b9164c0 := value
			var var_bbab2b9164c0_mapped *structpb.Value

			var var_bbab2b9164c0_err error
			var_bbab2b9164c0_mapped, var_bbab2b9164c0_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_bbab2b9164c0)
			if var_bbab2b9164c0_err != nil {
				panic(var_bbab2b9164c0_err)
			}

			var_e77498d50640_l = append(var_e77498d50640_l, var_bbab2b9164c0_mapped)
		}
		var_e77498d50640_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_e77498d50640_l})
		properties["roles"] = var_e77498d50640_mapped
	}

	var_c870a8b4f5eb := user.SecurityConstraints

	if var_c870a8b4f5eb != nil {
		var var_c870a8b4f5eb_mapped *structpb.Value

		var var_c870a8b4f5eb_l []*structpb.Value
		for _, value := range var_c870a8b4f5eb {

			var_974a004e0a92 := value
			var var_974a004e0a92_mapped *structpb.Value

			var_974a004e0a92_mapped = structpb.NewStructValue(&structpb.Struct{Fields: SecurityConstraintMapperInstance.ToProperties(var_974a004e0a92)})

			var_c870a8b4f5eb_l = append(var_c870a8b4f5eb_l, var_974a004e0a92_mapped)
		}
		var_c870a8b4f5eb_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_c870a8b4f5eb_l})
		properties["securityConstraints"] = var_c870a8b4f5eb_mapped
	}

	var_870a16ea1b31 := user.Details

	if var_870a16ea1b31 != nil {
		var var_870a16ea1b31_mapped *structpb.Value

		var var_870a16ea1b31_err error
		var_870a16ea1b31_mapped, var_870a16ea1b31_err = types.ByResourcePropertyType(model.ResourceProperty_OBJECT).Pack(*var_870a16ea1b31)
		if var_870a16ea1b31_err != nil {
			panic(var_870a16ea1b31_err)
		}
		properties["details"] = var_870a16ea1b31_mapped
	}
	return properties
}

func (m *UserMapper) FromProperties(properties map[string]*structpb.Value) *User {
	var s = m.New()
	if properties["id"] != nil {

		var_1deb6e5b2e3e := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_1deb6e5b2e3e)

		if err != nil {
			panic(err)
		}

		var_1deb6e5b2e3e_mapped := new(uuid.UUID)
		*var_1deb6e5b2e3e_mapped = val.(uuid.UUID)

		s.Id = var_1deb6e5b2e3e_mapped
	}
	if properties["version"] != nil {

		var_0cfcc924afc6 := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_0cfcc924afc6)

		if err != nil {
			panic(err)
		}

		var_0cfcc924afc6_mapped := val.(int32)

		s.Version = var_0cfcc924afc6_mapped
	}
	if properties["createdBy"] != nil {

		var_5d16813f3957 := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_5d16813f3957)

		if err != nil {
			panic(err)
		}

		var_5d16813f3957_mapped := new(string)
		*var_5d16813f3957_mapped = val.(string)

		s.CreatedBy = var_5d16813f3957_mapped
	}
	if properties["updatedBy"] != nil {

		var_4137a836703b := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_4137a836703b)

		if err != nil {
			panic(err)
		}

		var_4137a836703b_mapped := new(string)
		*var_4137a836703b_mapped = val.(string)

		s.UpdatedBy = var_4137a836703b_mapped
	}
	if properties["createdOn"] != nil {

		var_ad5c8a7a613d := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_ad5c8a7a613d)

		if err != nil {
			panic(err)
		}

		var_ad5c8a7a613d_mapped := new(time.Time)
		*var_ad5c8a7a613d_mapped = val.(time.Time)

		s.CreatedOn = var_ad5c8a7a613d_mapped
	}
	if properties["updatedOn"] != nil {

		var_6c7dba1b9eee := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_6c7dba1b9eee)

		if err != nil {
			panic(err)
		}

		var_6c7dba1b9eee_mapped := new(time.Time)
		*var_6c7dba1b9eee_mapped = val.(time.Time)

		s.UpdatedOn = var_6c7dba1b9eee_mapped
	}
	if properties["username"] != nil {

		var_92d978ad00a6 := properties["username"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_92d978ad00a6)

		if err != nil {
			panic(err)
		}

		var_92d978ad00a6_mapped := val.(string)

		s.Username = var_92d978ad00a6_mapped
	}
	if properties["password"] != nil {

		var_7945ef24a249 := properties["password"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_7945ef24a249)

		if err != nil {
			panic(err)
		}

		var_7945ef24a249_mapped := new(string)
		*var_7945ef24a249_mapped = val.(string)

		s.Password = var_7945ef24a249_mapped
	}
	if properties["roles"] != nil {

		var_db230176c3c3 := properties["roles"]
		var_db230176c3c3_mapped := []string{}
		for _, v := range var_db230176c3c3.GetListValue().Values {

			var_aaf24db80c5a := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_aaf24db80c5a)

			if err != nil {
				panic(err)
			}

			var_aaf24db80c5a_mapped := val.(string)

			var_db230176c3c3_mapped = append(var_db230176c3c3_mapped, var_aaf24db80c5a_mapped)
		}

		s.Roles = var_db230176c3c3_mapped
	}
	if properties["securityConstraints"] != nil {

		var_1782b9348e8a := properties["securityConstraints"]
		var_1782b9348e8a_mapped := []*SecurityConstraint{}
		for _, v := range var_1782b9348e8a.GetListValue().Values {

			var_eec229c83911 := v
			var_eec229c83911_mapped := SecurityConstraintMapperInstance.FromProperties(var_eec229c83911.GetStructValue().Fields)

			var_1782b9348e8a_mapped = append(var_1782b9348e8a_mapped, var_eec229c83911_mapped)
		}

		s.SecurityConstraints = var_1782b9348e8a_mapped
	}
	if properties["details"] != nil {

		var_33a2f16850e4 := properties["details"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_OBJECT).UnPack(var_33a2f16850e4)

		if err != nil {
			panic(err)
		}

		var_33a2f16850e4_mapped := new(unstructured.Unstructured)
		*var_33a2f16850e4_mapped = val.(unstructured.Unstructured)

		s.Details = var_33a2f16850e4_mapped
	}
	return s
}
