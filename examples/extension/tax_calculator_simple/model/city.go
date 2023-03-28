package model

import "time"
import "reflect"
import "github.com/tislib/data-handler/pkg/model"
import "github.com/tislib/data-handler/pkg/client"
import "github.com/google/uuid"
import "github.com/tislib/data-handler/pkg/types"
import "google.golang.org/protobuf/types/known/structpb"

type City struct {
	Id          uuid.UUID
	Name        *string
	Country     *Country
	State       *State
	Description *string
	CreatedBy   string
	UpdatedBy   *string
	CreatedOn   time.Time
	UpdatedOn   *time.Time
	Version     int32
}

func (s *City) GetId() string {
	valStr := types.ByResourcePropertyType(model.ResourceProperty_UUID).String(s.Id)
	return valStr
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
		val0, _ := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(properties["id"])
		s.Id = val0.(uuid.UUID)
	}

	if properties["name"] != nil {
		val1, _ := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(properties["name"])
		s.Name = new(string)
		*s.Name = val1.(string)
	}

	if properties["country"] != nil {
		s.Country = new(Country)
		s.Country.FromProperties(properties["country"].GetStructValue().Fields)
	}

	if properties["state"] != nil {
		s.State = new(State)
		s.State.FromProperties(properties["state"].GetStructValue().Fields)
	}

	if properties["description"] != nil {
		val4, _ := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(properties["description"])
		s.Description = new(string)
		*s.Description = val4.(string)
	}

	if properties["createdBy"] != nil {
		val5, _ := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(properties["createdBy"])
		s.CreatedBy = val5.(string)
	}

	if properties["updatedBy"] != nil {
		val6, _ := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(properties["updatedBy"])
		s.UpdatedBy = new(string)
		*s.UpdatedBy = val6.(string)
	}

	if properties["createdOn"] != nil {
		val7, _ := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(properties["createdOn"])
		s.CreatedOn = val7.(time.Time)
	}

	if properties["updatedOn"] != nil {
		val8, _ := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(properties["updatedOn"])
		s.UpdatedOn = new(time.Time)
		*s.UpdatedOn = val8.(time.Time)
	}

	if properties["version"] != nil {
		val9, _ := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(properties["version"])
		s.Version = val9.(int32)
	}

}

func (s *City) ToProperties() map[string]*structpb.Value {
	var properties = make(map[string]*structpb.Value)

	val0, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(s.Id)
	if err != nil {
		panic(err)
	}
	properties["id"] = val0

	if s.Name != nil {
		val1, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*s.Name)
		if err != nil {
			panic(err)
		}
		properties["name"] = val1
	}

	if s.Country != nil {
		properties["country"] = structpb.NewStructValue(&structpb.Struct{Fields: s.Country.ToProperties()})
	}

	if s.State != nil {
		properties["state"] = structpb.NewStructValue(&structpb.Struct{Fields: s.State.ToProperties()})
	}

	if s.Description != nil {
		val4, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*s.Description)
		if err != nil {
			panic(err)
		}
		properties["description"] = val4
	}

	val5, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(s.CreatedBy)
	if err != nil {
		panic(err)
	}
	properties["createdBy"] = val5

	if s.UpdatedBy != nil {
		val6, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*s.UpdatedBy)
		if err != nil {
			panic(err)
		}
		properties["updatedBy"] = val6
	}

	val7, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(s.CreatedOn)
	if err != nil {
		panic(err)
	}
	properties["createdOn"] = val7

	if s.UpdatedOn != nil {
		val8, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*s.UpdatedOn)
		if err != nil {
			panic(err)
		}
		properties["updatedOn"] = val8
	}

	val9, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(s.Version)
	if err != nil {
		panic(err)
	}
	properties["version"] = val9

	return properties
}

func (s *City) GetResourceName() string {
	return "city"
}

func (s *City) GetNamespace() string {
	return "default"
}

func (s *City) Clone() *City {
	var newInstance = new(City)
	newInstance.Id = s.Id
	if s.Name != nil {
		newInstance.Name = s.Name
	}

	if s.Country != nil {
		newInstance.Country = s.Country
	}

	if s.State != nil {
		newInstance.State = s.State
	}

	if s.Description != nil {
		newInstance.Description = s.Description
	}

	newInstance.CreatedBy = s.CreatedBy
	if s.UpdatedBy != nil {
		newInstance.UpdatedBy = s.UpdatedBy
	}

	newInstance.CreatedOn = s.CreatedOn
	if s.UpdatedOn != nil {
		newInstance.UpdatedOn = s.UpdatedOn
	}

	newInstance.Version = s.Version
	return newInstance
}

func (s *City) Equals(other *City) bool {
	return reflect.DeepEqual(s, other)
}

func (s *City) Same(other *City) bool {
	return s.Equals(other)
}

func NewCityRepository(dhClient client.DhClient) client.Repository[*City] {
	return client.NewRepository[*City](dhClient, client.RepositoryParams[*City]{Instance: new(City)})
}
