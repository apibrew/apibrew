// Code generated by apbr generate. DO NOT EDIT.
// versions:
// 	apbr generate v1.2

//go:build !codeanalysis

package model

import (
	"github.com/apibrew/apibrew/pkg/abs"
	"github.com/apibrew/apibrew/pkg/formats/unstructured"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/types"
	"google.golang.org/protobuf/types/known/structpb"
)

import "github.com/google/uuid"

type TenantMapper struct {
}

func NewTenantMapper() *TenantMapper {
	return &TenantMapper{}
}

var TenantMapperInstance = NewTenantMapper()

func (m *TenantMapper) New() *Tenant {
	return &Tenant{}
}

func (m *TenantMapper) ResourceIdentity() abs.ResourceIdentity {
	return abs.ResourceIdentity{
		Namespace: "default",
		Name:      "Tenant",
	}
}

func (m *TenantMapper) ToRecord(tenant *Tenant) *model.Record {
	var rec = &model.Record{}
	rec.Properties = m.ToProperties(tenant)
	return rec
}

func (m *TenantMapper) FromRecord(record *model.Record) *Tenant {
	return m.FromProperties(record.Properties)
}

func (m *TenantMapper) ToProperties(tenant *Tenant) map[string]*structpb.Value {
	var properties = make(map[string]*structpb.Value)

	var_Id := tenant.Id

	if var_Id != nil {
		var var_Id_mapped *structpb.Value

		var var_Id_err error
		var_Id_mapped, var_Id_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_Id)
		if var_Id_err != nil {
			panic(var_Id_err)
		}
		properties["id"] = var_Id_mapped
	}

	var_Name := tenant.Name

	if var_Name != nil {
		var var_Name_mapped *structpb.Value

		var var_Name_err error
		var_Name_mapped, var_Name_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_Name)
		if var_Name_err != nil {
			panic(var_Name_err)
		}
		properties["name"] = var_Name_mapped
	}

	var_Description := tenant.Description

	if var_Description != nil {
		var var_Description_mapped *structpb.Value

		var var_Description_err error
		var_Description_mapped, var_Description_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_Description)
		if var_Description_err != nil {
			panic(var_Description_err)
		}
		properties["description"] = var_Description_mapped
	}

	var_Version := tenant.Version

	var var_Version_mapped *structpb.Value

	var var_Version_err error
	var_Version_mapped, var_Version_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(var_Version)
	if var_Version_err != nil {
		panic(var_Version_err)
	}
	properties["version"] = var_Version_mapped
	return properties
}

func (m *TenantMapper) FromProperties(properties map[string]*structpb.Value) *Tenant {
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
	if properties["name"] != nil && properties["name"].AsInterface() != nil {

		var_Name := properties["name"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_Name)

		if err != nil {
			panic(err)
		}

		var_Name_mapped := new(string)
		*var_Name_mapped = val.(string)

		s.Name = var_Name_mapped
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
	if properties["version"] != nil && properties["version"].AsInterface() != nil {

		var_Version := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_Version)

		if err != nil {
			panic(err)
		}

		var_Version_mapped := val.(int32)

		s.Version = var_Version_mapped
	}
	return s
}

func (m *TenantMapper) ToUnstructured(tenant *Tenant) unstructured.Unstructured {
	var properties unstructured.Unstructured = make(unstructured.Unstructured)
	properties["type"] = "default/Tenant"

	if tenant == nil {
		return properties
	}

	var_Id := tenant.Id

	if var_Id != nil {
		var var_Id_mapped interface{}

		var_Id_mapped = var_Id.String()
		properties["id"] = var_Id_mapped
	}

	var_Name := tenant.Name

	if var_Name != nil {
		var var_Name_mapped interface{}

		var_Name_mapped = *var_Name
		properties["name"] = var_Name_mapped
	}

	var_Description := tenant.Description

	if var_Description != nil {
		var var_Description_mapped interface{}

		var_Description_mapped = *var_Description
		properties["description"] = var_Description_mapped
	}

	var_Version := tenant.Version

	var var_Version_mapped interface{}

	var_Version_mapped = var_Version
	properties["version"] = var_Version_mapped

	return properties
}
