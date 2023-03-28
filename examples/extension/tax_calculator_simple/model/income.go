package model

import "time"
import "reflect"
import "github.com/tislib/data-handler/pkg/model"
import "github.com/tislib/data-handler/pkg/client"
import "github.com/google/uuid"
import "github.com/tislib/data-handler/pkg/types"
import "google.golang.org/protobuf/types/known/structpb"

type Income struct {
	Country     *Country
	City        *City
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
	if properties["country"] != nil {
		s.Country = new(Country)
		s.Country.FromProperties(properties["country"].GetStructValue().Fields)
	}

	if properties["city"] != nil {
		s.City = new(City)
		s.City.FromProperties(properties["city"].GetStructValue().Fields)
	}

	if properties["updatedBy"] != nil {
		val2, _ := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(properties["updatedBy"])
		s.UpdatedBy = new(string)
		*s.UpdatedBy = val2.(string)
	}

	if properties["createdOn"] != nil {
		val3, _ := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(properties["createdOn"])
		s.CreatedOn = val3.(time.Time)
	}

	if properties["updatedOn"] != nil {
		val4, _ := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(properties["updatedOn"])
		s.UpdatedOn = new(time.Time)
		*s.UpdatedOn = val4.(time.Time)
	}

	if properties["version"] != nil {
		val5, _ := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(properties["version"])
		s.Version = val5.(int32)
	}

	if properties["id"] != nil {
		val6, _ := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(properties["id"])
		s.Id = val6.(uuid.UUID)
	}

	if properties["gross_income"] != nil {
		val7, _ := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(properties["gross_income"])
		s.GrossIncome = val7.(int32)
	}

	if properties["tax"] != nil {
		val8, _ := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(properties["tax"])
		s.Tax = new(int32)
		*s.Tax = val8.(int32)
	}

	if properties["net_income"] != nil {
		val9, _ := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(properties["net_income"])
		s.NetIncome = new(int32)
		*s.NetIncome = val9.(int32)
	}

	if properties["createdBy"] != nil {
		val10, _ := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(properties["createdBy"])
		s.CreatedBy = val10.(string)
	}

}

func (s *Income) ToProperties() map[string]*structpb.Value {
	var properties = make(map[string]*structpb.Value)

	if s.Country != nil {
		properties["country"] = structpb.NewStructValue(&structpb.Struct{Fields: s.Country.ToProperties()})
	}

	if s.City != nil {
		properties["city"] = structpb.NewStructValue(&structpb.Struct{Fields: s.City.ToProperties()})
	}

	if s.UpdatedBy != nil {
		val2, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*s.UpdatedBy)
		if err != nil {
			panic(err)
		}
		properties["updatedBy"] = val2
	}

	val3, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(s.CreatedOn)
	if err != nil {
		panic(err)
	}
	properties["createdOn"] = val3

	if s.UpdatedOn != nil {
		val4, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*s.UpdatedOn)
		if err != nil {
			panic(err)
		}
		properties["updatedOn"] = val4
	}

	val5, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(s.Version)
	if err != nil {
		panic(err)
	}
	properties["version"] = val5

	val6, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(s.Id)
	if err != nil {
		panic(err)
	}
	properties["id"] = val6

	val7, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(s.GrossIncome)
	if err != nil {
		panic(err)
	}
	properties["gross_income"] = val7

	if s.Tax != nil {
		val8, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(*s.Tax)
		if err != nil {
			panic(err)
		}
		properties["tax"] = val8
	}

	if s.NetIncome != nil {
		val9, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(*s.NetIncome)
		if err != nil {
			panic(err)
		}
		properties["net_income"] = val9
	}

	val10, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(s.CreatedBy)
	if err != nil {
		panic(err)
	}
	properties["createdBy"] = val10

	return properties
}

func (s *Income) GetResourceName() string {
	return "income"
}

func (s *Income) GetNamespace() string {
	return "default"
}

func (s *Income) Clone() *Income {
	var newInstance = new(Income)
	if s.Country != nil {
		newInstance.Country = s.Country
	}

	if s.City != nil {
		newInstance.City = s.City
	}

	if s.UpdatedBy != nil {
		newInstance.UpdatedBy = s.UpdatedBy
	}

	newInstance.CreatedOn = s.CreatedOn
	if s.UpdatedOn != nil {
		newInstance.UpdatedOn = s.UpdatedOn
	}

	newInstance.Version = s.Version
	newInstance.Id = s.Id
	newInstance.GrossIncome = s.GrossIncome
	if s.Tax != nil {
		newInstance.Tax = s.Tax
	}

	if s.NetIncome != nil {
		newInstance.NetIncome = s.NetIncome
	}

	newInstance.CreatedBy = s.CreatedBy
	return newInstance
}

func (s *Income) Equals(other *Income) bool {
	return reflect.DeepEqual(s, other)
}

func (s *Income) Same(other *Income) bool {
	return s.Equals(other)
}

func NewIncomeRepository(dhClient client.DhClient) client.Repository[*Income] {
	return client.NewRepository[*Income](dhClient, client.RepositoryParams[*Income]{Instance: new(Income)})
}
