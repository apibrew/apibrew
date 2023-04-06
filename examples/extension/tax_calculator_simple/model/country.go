package model

import "reflect"
import "github.com/tislib/data-handler/pkg/helper"
import "github.com/tislib/data-handler/pkg/model"
import "github.com/tislib/data-handler/pkg/client"
import "github.com/google/uuid"
import "github.com/tislib/data-handler/pkg/types"
import "google.golang.org/protobuf/types/known/structpb"

type Country struct {
	Id uuid.UUID

	Name *string

	Description *string

	Population *int64

	Area *int64

	Version int32
}

func (s *Country) GetId() uuid.UUID {
	return s.Id
}

func (s *Country) GetName() *string {
	return s.Name
}

func (s *Country) GetDescription() *string {
	return s.Description
}

func (s *Country) GetPopulation() *int64 {
	return s.Population
}

func (s *Country) GetArea() *int64 {
	return s.Area
}

func (s *Country) GetVersion() int32 {
	return s.Version
}

func (s *Country) ToRecord() *model.Record {
	var rec = &model.Record{}
	rec.Properties = s.ToProperties()

	return rec
}

func (s *Country) FromRecord(record *model.Record) {
	s.FromProperties(record.Properties)
}

func (s *Country) FromProperties(properties map[string]*structpb.Value) {

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

	if properties["description"] != nil {

		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(properties["description"])

		if err != nil {
			panic(err)
		}

		s.Description = new(string)
		*s.Description = val.(string)

	}

	if properties["population"] != nil {

		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT64).UnPack(properties["population"])

		if err != nil {
			panic(err)
		}

		s.Population = new(int64)
		*s.Population = val.(int64)

	}

	if properties["area"] != nil {

		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT64).UnPack(properties["area"])

		if err != nil {
			panic(err)
		}

		s.Area = new(int64)
		*s.Area = val.(int64)

	}

	if properties["version"] != nil {

		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(properties["version"])

		if err != nil {
			panic(err)
		}

		s.Version = val.(int32)

	}

}

func (s *Country) ToProperties() map[string]*structpb.Value {
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

	if s.Description != nil {

		Description, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*s.Description)
		if err != nil {
			panic(err)
		}
		properties["description"] = Description

	}

	if s.Population != nil {

		Population, err := types.ByResourcePropertyType(model.ResourceProperty_INT64).Pack(*s.Population)
		if err != nil {
			panic(err)
		}
		properties["population"] = Population

	}

	if s.Area != nil {

		Area, err := types.ByResourcePropertyType(model.ResourceProperty_INT64).Pack(*s.Area)
		if err != nil {
			panic(err)
		}
		properties["area"] = Area

	}

	Version, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(s.Version)
	if err != nil {
		panic(err)
	}
	properties["version"] = Version

	return properties
}

func (s *Country) GetResourceName() string {
	return "country"
}

func (s *Country) GetNamespace() string {
	return "default"
}

func (s *Country) Equals(other *Country) bool {
	return reflect.DeepEqual(s, other)
}

func (s *Country) Same(other *Country) bool {
	return s.Equals(other)
}

func NewCountryRepository(dhClient client.DhClient) client.Repository[*Country] {
	return client.NewRepository[*Country](dhClient, client.RepositoryParams[*Country]{InstanceProvider: func() *Country {
		return new(Country)
	}})
}

var CountryId = client.DefineProperty[uuid.UUID, helper.UuidQueryBuilder]("id", model.ResourceProperty_UUID, helper.UuidQueryBuilder{PropName: "id"})

var CountryName = client.DefineProperty[string, helper.StringQueryBuilder]("name", model.ResourceProperty_STRING, helper.StringQueryBuilder{PropName: "name"})

var CountryDescription = client.DefineProperty[string, helper.StringQueryBuilder]("description", model.ResourceProperty_STRING, helper.StringQueryBuilder{PropName: "description"})

var CountryPopulation = client.DefineProperty[int64, helper.Int64QueryBuilder]("population", model.ResourceProperty_INT64, helper.Int64QueryBuilder{PropName: "population"})

var CountryArea = client.DefineProperty[int64, helper.Int64QueryBuilder]("area", model.ResourceProperty_INT64, helper.Int64QueryBuilder{PropName: "area"})

var CountryVersion = client.DefineProperty[int32, helper.Int32QueryBuilder]("version", model.ResourceProperty_INT32, helper.Int32QueryBuilder{PropName: "version"})
