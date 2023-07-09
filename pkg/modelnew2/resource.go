package modelnew2

import "time"
import "reflect"
import "github.com/apibrew/apibrew/pkg/model"
import "github.com/apibrew/apibrew/pkg/types"
import "google.golang.org/protobuf/types/known/structpb"

type Resource struct {
	id                  aa
	version             aa
	createdBy           aa
	updatedBy           aa
	createdOn           aa
	updatedOn           aa
	name                aa
	namespace           aa
	virtual             aa
	types               aa
	immutable           aa
	abstract            aa
	dataSource          aa
	entity              aa
	catalog             aa
	annotations         aa
	indexes             aa
	securityConstraints aa
	title               aa
	description         aa
}

func (s *Resource) GetId() string {
	return s.Id
}

func (s *Resource) GetVersion() int32 {
	return s.Version
}

func (s *Resource) GetCreatedBy() string {
	return s.CreatedBy
}

func (s *Resource) GetUpdatedBy() *string {
	return s.UpdatedBy
}

func (s *Resource) GetCreatedOn() time.Time {
	return s.CreatedOn
}

func (s *Resource) GetUpdatedOn() *time.Time {
	return s.UpdatedOn
}

func (s *Resource) GetName() string {
	return s.Name
}

func (s *Resource) GetDescription() *string {
	return s.Description
}

func (s *Resource) GetBackend() ResourceBackendType {
	return s.Backend
}

func (s *Resource) GetOptions() map[string]string {
	return s.Options
}

func (s *Resource) ToRecord() *model.Record {
	var rec = &model.Record{}
	rec.Properties = s.ToProperties()

	return rec
}

func (s *Resource) FromRecord(record *model.Record) {
	s.FromProperties(record.Properties)
}

func (s *Resource) FromProperties(properties map[string]*structpb.Value) {

	if properties["id"] != nil {

		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(properties["id"])

		if err != nil {
			panic(err)
		}

		s.Id = val.(string)

	}

	if properties["version"] != nil {

		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(properties["version"])

		if err != nil {
			panic(err)
		}

		s.Version = val.(int32)

	}

	if properties["createdBy"] != nil {

		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(properties["createdBy"])

		if err != nil {
			panic(err)
		}

		s.CreatedBy = val.(string)

	}

	if properties["updatedBy"] != nil {

		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(properties["updatedBy"])

		if err != nil {
			panic(err)
		}

		s.UpdatedBy = new(string)
		*s.UpdatedBy = val.(string)

	}

	if properties["createdOn"] != nil {

		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(properties["createdOn"])

		if err != nil {
			panic(err)
		}

		s.CreatedOn = val.(time.Time)

	}

	if properties["updatedOn"] != nil {

		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(properties["updatedOn"])

		if err != nil {
			panic(err)
		}

		s.UpdatedOn = new(time.Time)
		*s.UpdatedOn = val.(time.Time)

	}

	if properties["name"] != nil {

		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(properties["name"])

		if err != nil {
			panic(err)
		}

		s.Name = val.(string)

	}

	if properties["description"] != nil {

		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(properties["description"])

		if err != nil {
			panic(err)
		}

		s.Description = new(string)
		*s.Description = val.(string)

	}

	if properties["backend"] != nil {

		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRUCT).UnPack(properties["backend"])

		if err != nil {
			panic(err)
		}

		s.Backend = ResourceBackendType(val.(string))

	}

	if properties["options"] != nil {

		val, err := types.ByResourcePropertyType(model.ResourceProperty_MAP).UnPack(properties["options"])

		if err != nil {
			panic(err)
		}

		mapData := val.(map[string]interface{})
		s.Options = make(map[string]string)

		for k, v := range mapData {
			s.Options[k] = v.(string)
		}
	}

}

func (s *Resource) ToProperties() map[string]*structpb.Value {
	var properties = make(map[string]*structpb.Value)

	Id, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(s.Id)
	if err != nil {
		panic(err)
	}
	properties["id"] = Id

	Version, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(s.Version)
	if err != nil {
		panic(err)
	}
	properties["version"] = Version

	CreatedBy, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(s.CreatedBy)
	if err != nil {
		panic(err)
	}
	properties["createdBy"] = CreatedBy

	if s.UpdatedBy != nil {

		UpdatedBy, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*s.UpdatedBy)
		if err != nil {
			panic(err)
		}
		properties["updatedBy"] = UpdatedBy

	}

	CreatedOn, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(s.CreatedOn)
	if err != nil {
		panic(err)
	}
	properties["createdOn"] = CreatedOn

	if s.UpdatedOn != nil {

		UpdatedOn, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*s.UpdatedOn)
		if err != nil {
			panic(err)
		}
		properties["updatedOn"] = UpdatedOn

	}

	Name, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(s.Name)
	if err != nil {
		panic(err)
	}
	properties["name"] = Name

	if s.Description != nil {

		Description, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*s.Description)
		if err != nil {
			panic(err)
		}
		properties["description"] = Description

	}

	Backend, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(s.Backend)
	if err != nil {
		panic(err)
	}
	properties["backend"] = Backend

	return properties
}

func (s *Resource) GetResourceName() string {
	return "data-source"
}

func (s *Resource) GetNamespace() string {
	return "system"
}

func (s *Resource) Equals(other *Resource) bool {
	return reflect.DeepEqual(s, other)
}

func (s *Resource) Same(other *Resource) bool {
	return s.Equals(other)
}
