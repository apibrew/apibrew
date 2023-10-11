// AUTOGENERATED FILE

//go:build !codeanalysis

package resource_model

import (
	"github.com/apibrew/apibrew/pkg/abs"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/types"
	"google.golang.org/protobuf/types/known/structpb"
)

import "github.com/google/uuid"
import "time"

type ResourceActionMapper struct {
}

func NewResourceActionMapper() *ResourceActionMapper {
	return &ResourceActionMapper{}
}

var ResourceActionMapperInstance = NewResourceActionMapper()

func (m *ResourceActionMapper) New() *ResourceAction {
	return &ResourceAction{}
}

func (m *ResourceActionMapper) ResourceIdentity() abs.ResourceIdentity {
	return abs.ResourceIdentity{
		Namespace: "system",
		Name:      "ResourceAction",
	}
}

func (m *ResourceActionMapper) ToRecord(resourceAction *ResourceAction) *model.Record {
	var rec = &model.Record{}
	rec.Properties = m.ToProperties(resourceAction)

	if resourceAction.Id != nil {
		rec.Id = resourceAction.Id.String()
	}

	return rec
}

func (m *ResourceActionMapper) FromRecord(record *model.Record) *ResourceAction {
	return m.FromProperties(record.Properties)
}

func (m *ResourceActionMapper) ToProperties(resourceAction *ResourceAction) map[string]*structpb.Value {
	var properties = make(map[string]*structpb.Value)

	var_Id := resourceAction.Id

	if var_Id != nil {
		var var_Id_mapped *structpb.Value

		var var_Id_err error
		var_Id_mapped, var_Id_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_Id)
		if var_Id_err != nil {
			panic(var_Id_err)
		}
		properties["id"] = var_Id_mapped
	}

	var_Version := resourceAction.Version

	var var_Version_mapped *structpb.Value

	var var_Version_err error
	var_Version_mapped, var_Version_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(var_Version)
	if var_Version_err != nil {
		panic(var_Version_err)
	}
	properties["version"] = var_Version_mapped

	var_CreatedBy := resourceAction.CreatedBy

	if var_CreatedBy != nil {
		var var_CreatedBy_mapped *structpb.Value

		var var_CreatedBy_err error
		var_CreatedBy_mapped, var_CreatedBy_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_CreatedBy)
		if var_CreatedBy_err != nil {
			panic(var_CreatedBy_err)
		}
		properties["createdBy"] = var_CreatedBy_mapped
	}

	var_UpdatedBy := resourceAction.UpdatedBy

	if var_UpdatedBy != nil {
		var var_UpdatedBy_mapped *structpb.Value

		var var_UpdatedBy_err error
		var_UpdatedBy_mapped, var_UpdatedBy_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_UpdatedBy)
		if var_UpdatedBy_err != nil {
			panic(var_UpdatedBy_err)
		}
		properties["updatedBy"] = var_UpdatedBy_mapped
	}

	var_CreatedOn := resourceAction.CreatedOn

	if var_CreatedOn != nil {
		var var_CreatedOn_mapped *structpb.Value

		var var_CreatedOn_err error
		var_CreatedOn_mapped, var_CreatedOn_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_CreatedOn)
		if var_CreatedOn_err != nil {
			panic(var_CreatedOn_err)
		}
		properties["createdOn"] = var_CreatedOn_mapped
	}

	var_UpdatedOn := resourceAction.UpdatedOn

	if var_UpdatedOn != nil {
		var var_UpdatedOn_mapped *structpb.Value

		var var_UpdatedOn_err error
		var_UpdatedOn_mapped, var_UpdatedOn_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_UpdatedOn)
		if var_UpdatedOn_err != nil {
			panic(var_UpdatedOn_err)
		}
		properties["updatedOn"] = var_UpdatedOn_mapped
	}

	var_Resource := resourceAction.Resource

	if var_Resource != nil {
		var var_Resource_mapped *structpb.Value

		var_Resource_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ResourceMapperInstance.ToProperties(var_Resource)})
		properties["resource"] = var_Resource_mapped
	}

	var_Name := resourceAction.Name

	var var_Name_mapped *structpb.Value

	var var_Name_err error
	var_Name_mapped, var_Name_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_Name)
	if var_Name_err != nil {
		panic(var_Name_err)
	}
	properties["name"] = var_Name_mapped

	var_Title := resourceAction.Title

	if var_Title != nil {
		var var_Title_mapped *structpb.Value

		var var_Title_err error
		var_Title_mapped, var_Title_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_Title)
		if var_Title_err != nil {
			panic(var_Title_err)
		}
		properties["title"] = var_Title_mapped
	}

	var_Description := resourceAction.Description

	if var_Description != nil {
		var var_Description_mapped *structpb.Value

		var var_Description_err error
		var_Description_mapped, var_Description_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_Description)
		if var_Description_err != nil {
			panic(var_Description_err)
		}
		properties["description"] = var_Description_mapped
	}

	var_Internal := resourceAction.Internal

	var var_Internal_mapped *structpb.Value

	var var_Internal_err error
	var_Internal_mapped, var_Internal_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(var_Internal)
	if var_Internal_err != nil {
		panic(var_Internal_err)
	}
	properties["internal"] = var_Internal_mapped

	var_Types := resourceAction.Types

	if var_Types != nil {
		var var_Types_mapped *structpb.Value

		var var_Types_l []*structpb.Value
		for _, value := range var_Types {

			var_5x := value
			var var_5x_mapped *structpb.Value

			var_5x_mapped = structpb.NewStructValue(&structpb.Struct{Fields: SubTypeMapperInstance.ToProperties(&var_5x)})

			var_Types_l = append(var_Types_l, var_5x_mapped)
		}
		var_Types_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_Types_l})
		properties["types"] = var_Types_mapped
	}

	var_Input := resourceAction.Input

	if var_Input != nil {
		var var_Input_mapped *structpb.Value

		var var_Input_l []*structpb.Value
		for _, value := range var_Input {

			var_5x := value
			var var_5x_mapped *structpb.Value

			var_5x_mapped = structpb.NewStructValue(&structpb.Struct{Fields: PropertyMapperInstance.ToProperties(&var_5x)})

			var_Input_l = append(var_Input_l, var_5x_mapped)
		}
		var_Input_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_Input_l})
		properties["input"] = var_Input_mapped
	}

	var_Output := resourceAction.Output

	if var_Output != nil {
		var var_Output_mapped *structpb.Value

		var_Output_mapped = structpb.NewStructValue(&structpb.Struct{Fields: PropertyMapperInstance.ToProperties(var_Output)})
		properties["output"] = var_Output_mapped
	}

	var_Annotations := resourceAction.Annotations

	if var_Annotations != nil {
		var var_Annotations_mapped *structpb.Value

		var var_Annotations_st *structpb.Struct = new(structpb.Struct)
		var_Annotations_st.Fields = make(map[string]*structpb.Value)
		for key, value := range var_Annotations {

			var_1x := value
			var var_1x_mapped *structpb.Value

			var var_1x_err error
			var_1x_mapped, var_1x_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_1x)
			if var_1x_err != nil {
				panic(var_1x_err)
			}

			var_Annotations_st.Fields[key] = var_1x_mapped
		}
		var_Annotations_mapped = structpb.NewStructValue(var_Annotations_st)
		properties["annotations"] = var_Annotations_mapped
	}
	return properties
}

func (m *ResourceActionMapper) FromProperties(properties map[string]*structpb.Value) *ResourceAction {
	var s = m.New()
	if properties["id"] != nil && properties["id"].AsInterface() != nil {

		var_Id := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_Id)

		if err != nil {
			panic(err)
		}

		var_Id_mapped := new(uuid.UUID)
		*var_Id_mapped = val.(uuid.UUID)

		s.Id = var_Id_mapped
	}
	if properties["version"] != nil && properties["version"].AsInterface() != nil {

		var_Version := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_Version)

		if err != nil {
			panic(err)
		}

		var_Version_mapped := val.(int32)

		s.Version = var_Version_mapped
	}
	if properties["createdBy"] != nil && properties["createdBy"].AsInterface() != nil {

		var_CreatedBy := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_CreatedBy)

		if err != nil {
			panic(err)
		}

		var_CreatedBy_mapped := new(string)
		*var_CreatedBy_mapped = val.(string)

		s.CreatedBy = var_CreatedBy_mapped
	}
	if properties["updatedBy"] != nil && properties["updatedBy"].AsInterface() != nil {

		var_UpdatedBy := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_UpdatedBy)

		if err != nil {
			panic(err)
		}

		var_UpdatedBy_mapped := new(string)
		*var_UpdatedBy_mapped = val.(string)

		s.UpdatedBy = var_UpdatedBy_mapped
	}
	if properties["createdOn"] != nil && properties["createdOn"].AsInterface() != nil {

		var_CreatedOn := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_CreatedOn)

		if err != nil {
			panic(err)
		}

		var_CreatedOn_mapped := new(time.Time)
		*var_CreatedOn_mapped = val.(time.Time)

		s.CreatedOn = var_CreatedOn_mapped
	}
	if properties["updatedOn"] != nil && properties["updatedOn"].AsInterface() != nil {

		var_UpdatedOn := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_UpdatedOn)

		if err != nil {
			panic(err)
		}

		var_UpdatedOn_mapped := new(time.Time)
		*var_UpdatedOn_mapped = val.(time.Time)

		s.UpdatedOn = var_UpdatedOn_mapped
	}
	if properties["resource"] != nil && properties["resource"].AsInterface() != nil {

		var_Resource := properties["resource"]
		var_Resource_mapped := ResourceMapperInstance.FromProperties(var_Resource.GetStructValue().Fields)

		s.Resource = var_Resource_mapped
	}
	if properties["name"] != nil && properties["name"].AsInterface() != nil {

		var_Name := properties["name"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_Name)

		if err != nil {
			panic(err)
		}

		var_Name_mapped := val.(string)

		s.Name = var_Name_mapped
	}
	if properties["title"] != nil && properties["title"].AsInterface() != nil {

		var_Title := properties["title"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_Title)

		if err != nil {
			panic(err)
		}

		var_Title_mapped := new(string)
		*var_Title_mapped = val.(string)

		s.Title = var_Title_mapped
	}
	if properties["description"] != nil && properties["description"].AsInterface() != nil {

		var_Description := properties["description"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_Description)

		if err != nil {
			panic(err)
		}

		var_Description_mapped := new(string)
		*var_Description_mapped = val.(string)

		s.Description = var_Description_mapped
	}
	if properties["internal"] != nil && properties["internal"].AsInterface() != nil {

		var_Internal := properties["internal"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_Internal)

		if err != nil {
			panic(err)
		}

		var_Internal_mapped := val.(bool)

		s.Internal = var_Internal_mapped
	}
	if properties["types"] != nil && properties["types"].AsInterface() != nil {

		var_Types := properties["types"]
		var_Types_mapped := []SubType{}
		for _, v := range var_Types.GetListValue().Values {

			var_4x := v
			var mappedValue = SubTypeMapperInstance.FromProperties(var_4x.GetStructValue().Fields)

			var_4x_mapped := *mappedValue

			var_Types_mapped = append(var_Types_mapped, var_4x_mapped)
		}

		s.Types = var_Types_mapped
	}
	if properties["input"] != nil && properties["input"].AsInterface() != nil {

		var_Input := properties["input"]
		var_Input_mapped := []Property{}
		for _, v := range var_Input.GetListValue().Values {

			var_4x := v
			var mappedValue = PropertyMapperInstance.FromProperties(var_4x.GetStructValue().Fields)

			var_4x_mapped := *mappedValue

			var_Input_mapped = append(var_Input_mapped, var_4x_mapped)
		}

		s.Input = var_Input_mapped
	}
	if properties["output"] != nil && properties["output"].AsInterface() != nil {

		var_Output := properties["output"]
		var mappedValue = PropertyMapperInstance.FromProperties(var_Output.GetStructValue().Fields)

		var_Output_mapped := mappedValue

		s.Output = var_Output_mapped
	}
	if properties["annotations"] != nil && properties["annotations"].AsInterface() != nil {

		var_Annotations := properties["annotations"]
		var_Annotations_mapped := make(map[string]string)
		for k, v := range var_Annotations.GetStructValue().Fields {

			var_3x := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_3x)

			if err != nil {
				panic(err)
			}

			var_3x_mapped := val.(string)

			var_Annotations_mapped[k] = var_3x_mapped
		}

		s.Annotations = var_Annotations_mapped
	}
	return s
}
