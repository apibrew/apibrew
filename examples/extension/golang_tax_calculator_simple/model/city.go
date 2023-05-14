package model

import "reflect"
import "github.com/apibrew/apibrew/pkg/helper"
import "github.com/apibrew/apibrew/pkg/model"
import "github.com/apibrew/apibrew/pkg/client"
import "github.com/google/uuid"
import "github.com/apibrew/apibrew/pkg/types"
import "google.golang.org/protobuf/types/known/structpb"

type City struct {
	Id uuid.UUID

	Name *string

	Country *Country

	Description *string

	Version int32
}

func (s *City) GetId() uuid.UUID {
	return s.Id
}

func (s *City) GetName() *string {
	return s.Name
}

func (s *City) GetCountry() *Country {
	return s.Country
}

func (s *City) GetDescription() *string {
	return s.Description
}

func (s *City) GetVersion() int32 {
	return s.Version
}

func (s *City) ToRecord() *model.Record {
	var rec = &model.Record{}
	rec.Properties = s.ToProperties()

	return rec
}

func (s *City) FromRecord(record *model.Record) {
	s.FromProperties(record.Properties)
}

func (s *City) FromProperties(properties map[string]*structpb.Value) {

	if properties["id"] != nil {

		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(properties["id"])

		if err != nil {
			panic(err)
		}

		s.Id = val.(uuid.UUID)

	}

	if properties["name"] != nil {

		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(properties["name"])

		if err != nil {
			panic(err)
		}

		s.Name = new(string)
		*s.Name = val.(string)

	}

	if properties["country"] != nil {

		s.Country = new(Country)
		s.Country.FromProperties(properties["country"].GetStructValue().Fields)

	}

	if properties["description"] != nil {

		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(properties["description"])

		if err != nil {
			panic(err)
		}

		s.Description = new(string)
		*s.Description = val.(string)

	}

	if properties["version"] != nil {

		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(properties["version"])

		if err != nil {
			panic(err)
		}

		s.Version = val.(int32)

	}

}

func (s *City) ToProperties() map[string]*structpb.Value {
	var properties = make(map[string]*structpb.Value)

	Id, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(s.Id)
	if err != nil {
		panic(err)
	}
	properties["id"] = Id

	if s.Name != nil {

		Name, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*s.Name)
		if err != nil {
			panic(err)
		}
		properties["name"] = Name

	}

	if s.Country != nil {

		properties["country"] = structpb.NewStructValue(&structpb.Struct{Fields: s.Country.ToProperties()})

	}

	if s.Description != nil {

		Description, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*s.Description)
		if err != nil {
			panic(err)
		}
		properties["description"] = Description

	}

	Version, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(s.Version)
	if err != nil {
		panic(err)
	}
	properties["version"] = Version

	return properties
}

func (s *City) GetResourceName() string {
	return "city"
}

func (s *City) GetNamespace() string {
	return "default"
}

func (s *City) Equals(other *City) bool {
	return reflect.DeepEqual(s, other)
}

func (s *City) Same(other *City) bool {
	return s.Equals(other)
}

func NewCityRepository(dhClient client.DhClient) client.Repository[*City] {
	return client.NewRepository[*City](dhClient, client.RepositoryParams[*City]{InstanceProvider: func() *City {
		return new(City)
	}})
}

var CityId = client.DefineProperty[uuid.UUID, helper.UuidQueryBuilder]("id", model.ResourceProperty_UUID, helper.UuidQueryBuilder{PropName: "id"})

var CityName = client.DefineProperty[string, helper.StringQueryBuilder]("name", model.ResourceProperty_STRING, helper.StringQueryBuilder{PropName: "name"})

var CityCountry = client.DefineProperty[*Country, helper.ReferenceQueryBuilder[*Country]]("country", model.ResourceProperty_REFERENCE, helper.ReferenceQueryBuilder[*Country]{PropName: "country"})

var CityDescription = client.DefineProperty[string, helper.StringQueryBuilder]("description", model.ResourceProperty_STRING, helper.StringQueryBuilder{PropName: "description"})

var CityVersion = client.DefineProperty[int32, helper.Int32QueryBuilder]("version", model.ResourceProperty_INT32, helper.Int32QueryBuilder{PropName: "version"})
