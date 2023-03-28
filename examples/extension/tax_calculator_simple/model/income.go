package model

import "time"
import "github.com/tislib/data-handler/pkg/model"
import "github.com/tislib/data-handler/pkg/client"
import "github.com/google/uuid"
import "github.com/tislib/data-handler/pkg/types"
import "google.golang.org/protobuf/types/known/structpb"

type Income struct {
	UpdatedBy   *string
	CreatedOn   time.Time
	UpdatedOn   *time.Time
	Version     int32
	Id          uuid.UUID
	GrossIncome int32
	Tax         *int32
	NetIncome   *int32
	CreatedBy   string
}

func (s *Income) GetId() string {
	valStr := types.ByResourcePropertyType(model.ResourceProperty_UUID).String(s.Id)
	return valStr
}

func (s *Income) ToRecord() *model.Record {
	var rec = &model.Record{}
	rec.Properties = s.ToProperties()

	return rec
}

func (s *Income) FromRecord(record *model.Record) {
	s.FromProperties(record.Properties)
}

func (s *Income) FromProperties(properties map[string]*structpb.Value) {
	if properties["updatedBy"] != nil {
		val0, _ := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(properties["updatedBy"])
		s.UpdatedBy = new(string)
		*s.UpdatedBy = val0.(string)
	}

	if properties["createdOn"] != nil {
		val1, _ := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(properties["createdOn"])
		s.CreatedOn = val1.(time.Time)
	}

	if properties["updatedOn"] != nil {
		val2, _ := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(properties["updatedOn"])
		s.UpdatedOn = new(time.Time)
		*s.UpdatedOn = val2.(time.Time)
	}

	if properties["version"] != nil {
		val3, _ := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(properties["version"])
		s.Version = val3.(int32)
	}

	if properties["id"] != nil {
		val4, _ := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(properties["id"])
		s.Id = val4.(uuid.UUID)
	}

	if properties["gross_income"] != nil {
		val5, _ := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(properties["gross_income"])
		s.GrossIncome = val5.(int32)
	}

	if properties["tax"] != nil {
		val6, _ := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(properties["tax"])
		s.Tax = new(int32)
		*s.Tax = val6.(int32)
	}

	if properties["net_income"] != nil {
		val7, _ := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(properties["net_income"])
		s.NetIncome = new(int32)
		*s.NetIncome = val7.(int32)
	}

	if properties["createdBy"] != nil {
		val8, _ := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(properties["createdBy"])
		s.CreatedBy = val8.(string)
	}

}

func (s *Income) ToProperties() map[string]*structpb.Value {
	var properties = make(map[string]*structpb.Value)

	if s.UpdatedBy != nil {
		val0, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*s.UpdatedBy)
		if err != nil {
			panic(err)
		}
		properties["updatedBy"] = val0
	}

	val1, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(s.CreatedOn)
	if err != nil {
		panic(err)
	}
	properties["createdOn"] = val1

	if s.UpdatedOn != nil {
		val2, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*s.UpdatedOn)
		if err != nil {
			panic(err)
		}
		properties["updatedOn"] = val2
	}

	val3, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(s.Version)
	if err != nil {
		panic(err)
	}
	properties["version"] = val3

	val4, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(s.Id)
	if err != nil {
		panic(err)
	}
	properties["id"] = val4

	val5, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(s.GrossIncome)
	if err != nil {
		panic(err)
	}
	properties["gross_income"] = val5

	if s.Tax != nil {
		val6, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(*s.Tax)
		if err != nil {
			panic(err)
		}
		properties["tax"] = val6
	}

	if s.NetIncome != nil {
		val7, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(*s.NetIncome)
		if err != nil {
			panic(err)
		}
		properties["net_income"] = val7
	}

	val8, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(s.CreatedBy)
	if err != nil {
		panic(err)
	}
	properties["createdBy"] = val8

	return properties
}

func (s *Income) GetResourceName() string {
	return "income"
}

func (s *Income) GetNamespace() string {
	return "default"
}

func NewIncomeRepository(dhClient client.DhClient) client.Repository[*Income] {
	return client.NewRepository[*Income](dhClient, client.RepositoryParams[*Income]{Instance: new(Income)})
}
