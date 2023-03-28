package model

import "time"
import "github.com/tislib/data-handler/pkg/model"
import "github.com/tislib/data-handler/pkg/client"
import "github.com/google/uuid"
import "github.com/tislib/data-handler/pkg/types"
import "google.golang.org/protobuf/types/known/structpb"

type VirtualResource struct {
	Version     int32
	Id          uuid.UUID
	Name        string
	Description *string
	CreatedBy   string
	UpdatedBy   *string
	CreatedOn   time.Time
	UpdatedOn   *time.Time
}

func (s *VirtualResource) GetId() string {
	valStr := types.ByResourcePropertyType(model.ResourceProperty_UUID).String(s.Id)
	return valStr
}

func (s *VirtualResource) ToRecord() *model.Record {
	var rec = &model.Record{}
	rec.Properties = s.ToProperties()

	return rec
}

func (s *VirtualResource) FromRecord(record *model.Record) {
	s.FromProperties(record.Properties)
}

func (s *VirtualResource) FromProperties(properties map[string]*structpb.Value) {
	if properties["version"] != nil {
		val0, _ := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(properties["version"])
		s.Version = val0.(int32)
	}

	if properties["id"] != nil {
		val1, _ := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(properties["id"])
		s.Id = val1.(uuid.UUID)
	}

	if properties["name"] != nil {
		val2, _ := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(properties["name"])
		s.Name = val2.(string)
	}

	if properties["description"] != nil {
		val3, _ := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(properties["description"])
		s.Description = new(string)
		*s.Description = val3.(string)
	}

	if properties["createdBy"] != nil {
		val4, _ := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(properties["createdBy"])
		s.CreatedBy = val4.(string)
	}

	if properties["updatedBy"] != nil {
		val5, _ := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(properties["updatedBy"])
		s.UpdatedBy = new(string)
		*s.UpdatedBy = val5.(string)
	}

	if properties["createdOn"] != nil {
		val6, _ := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(properties["createdOn"])
		s.CreatedOn = val6.(time.Time)
	}

	if properties["updatedOn"] != nil {
		val7, _ := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(properties["updatedOn"])
		s.UpdatedOn = new(time.Time)
		*s.UpdatedOn = val7.(time.Time)
	}

}

func (s *VirtualResource) ToProperties() map[string]*structpb.Value {
	var properties = make(map[string]*structpb.Value)

	val0, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(s.Version)
	if err != nil {
		panic(err)
	}
	properties["version"] = val0

	val1, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(s.Id)
	if err != nil {
		panic(err)
	}
	properties["id"] = val1

	val2, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(s.Name)
	if err != nil {
		panic(err)
	}
	properties["name"] = val2

	if s.Description != nil {
		val3, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*s.Description)
		if err != nil {
			panic(err)
		}
		properties["description"] = val3
	}

	val4, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(s.CreatedBy)
	if err != nil {
		panic(err)
	}
	properties["createdBy"] = val4

	if s.UpdatedBy != nil {
		val5, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*s.UpdatedBy)
		if err != nil {
			panic(err)
		}
		properties["updatedBy"] = val5
	}

	val6, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(s.CreatedOn)
	if err != nil {
		panic(err)
	}
	properties["createdOn"] = val6

	if s.UpdatedOn != nil {
		val7, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*s.UpdatedOn)
		if err != nil {
			panic(err)
		}
		properties["updatedOn"] = val7
	}

	return properties
}

func (s *VirtualResource) GetResourceName() string {
	return "virtualResource"
}

func (s *VirtualResource) GetNamespace() string {
	return "default"
}

func NewVirtualResourceRepository(dhClient client.DhClient) client.Repository[*VirtualResource] {
	return client.NewRepository[*VirtualResource](dhClient, client.RepositoryParams[*VirtualResource]{Instance: new(VirtualResource)})
}
