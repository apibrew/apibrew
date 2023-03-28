package model

import "time"
import "reflect"
import "github.com/tislib/data-handler/pkg/model"
import "github.com/tislib/data-handler/pkg/client"
import "github.com/google/uuid"
import "github.com/tislib/data-handler/pkg/types"
import "google.golang.org/protobuf/types/known/structpb"

type Country struct {
	Id          uuid.UUID
	Name        *string
	Description *string
	Population  *int64
	Area        *int64
	CreatedBy   string
	UpdatedBy   *string
	CreatedOn   time.Time
	UpdatedOn   *time.Time
	Version     int32
}

func (s *Country) GetId() string {
	valStr := types.ByResourcePropertyType(model.ResourceProperty_UUID).String(s.Id)
	return valStr
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
		val0, _ := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(properties["id"])
		s.Id = val0.(uuid.UUID)
	}

	if properties["name"] != nil {
		val1, _ := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(properties["name"])
		s.Name = new(string)
		*s.Name = val1.(string)
	}

	if properties["description"] != nil {
		val2, _ := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(properties["description"])
		s.Description = new(string)
		*s.Description = val2.(string)
	}

	if properties["population"] != nil {
		val3, _ := types.ByResourcePropertyType(model.ResourceProperty_INT64).UnPack(properties["population"])
		s.Population = new(int64)
		*s.Population = val3.(int64)
	}

	if properties["area"] != nil {
		val4, _ := types.ByResourcePropertyType(model.ResourceProperty_INT64).UnPack(properties["area"])
		s.Area = new(int64)
		*s.Area = val4.(int64)
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

func (s *Country) ToProperties() map[string]*structpb.Value {
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

	if s.Description != nil {
		val2, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*s.Description)
		if err != nil {
			panic(err)
		}
		properties["description"] = val2
	}

	if s.Population != nil {
		val3, err := types.ByResourcePropertyType(model.ResourceProperty_INT64).Pack(*s.Population)
		if err != nil {
			panic(err)
		}
		properties["population"] = val3
	}

	if s.Area != nil {
		val4, err := types.ByResourcePropertyType(model.ResourceProperty_INT64).Pack(*s.Area)
		if err != nil {
			panic(err)
		}
		properties["area"] = val4
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

func (s *Country) GetResourceName() string {
	return "country"
}

func (s *Country) GetNamespace() string {
	return "default"
}

func (s *Country) Clone() *Country {
	var newInstance = new(Country)
	newInstance.Id = s.Id
	if s.Name != nil {
		newInstance.Name = s.Name
	}

	if s.Description != nil {
		newInstance.Description = s.Description
	}

	if s.Population != nil {
		newInstance.Population = s.Population
	}

	if s.Area != nil {
		newInstance.Area = s.Area
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

func (s *Country) Equals(other *Country) bool {
	return reflect.DeepEqual(s, other)
}

func (s *Country) Same(other *Country) bool {
	return s.Equals(other)
}

func NewCountryRepository(dhClient client.DhClient) client.Repository[*Country] {
	return client.NewRepository[*Country](dhClient, client.RepositoryParams[*Country]{Instance: new(Country)})
}
