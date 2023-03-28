package model

import "time"
import "reflect"
import "github.com/tislib/data-handler/pkg/model"
import "github.com/tislib/data-handler/pkg/client"
import "github.com/google/uuid"
import "github.com/tislib/data-handler/pkg/types"
import "google.golang.org/protobuf/types/known/structpb"

type TaxRate struct {
	Rate      float32
	CreatedBy string
	UpdatedBy *string
	CreatedOn time.Time
	Id        uuid.UUID
	Name      string
	Country   *Country
	City      *City
	Order     int32
	UpdatedOn *time.Time
	Until     int32
	Version   int32
}

func (s *TaxRate) GetId() string {
	valStr := types.ByResourcePropertyType(model.ResourceProperty_UUID).String(s.Id)
	return valStr
}

func (s *TaxRate) ToRecord() *model.Record {
	var rec = &model.Record{}
	rec.Properties = s.ToProperties()

	return rec
}

func (s *TaxRate) FromRecord(record *model.Record) {
	s.FromProperties(record.Properties)
}

func (s *TaxRate) FromProperties(properties map[string]*structpb.Value) {
	if properties["rate"] != nil {
		val0, _ := types.ByResourcePropertyType(model.ResourceProperty_FLOAT32).UnPack(properties["rate"])
		s.Rate = val0.(float32)
	}

	if properties["createdBy"] != nil {
		val1, _ := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(properties["createdBy"])
		s.CreatedBy = val1.(string)
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

	if properties["id"] != nil {
		val4, _ := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(properties["id"])
		s.Id = val4.(uuid.UUID)
	}

	if properties["name"] != nil {
		val5, _ := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(properties["name"])
		s.Name = val5.(string)
	}

	if properties["country"] != nil {
		s.Country = new(Country)
		s.Country.FromProperties(properties["country"].GetStructValue().Fields)
	}

	if properties["city"] != nil {
		s.City = new(City)
		s.City.FromProperties(properties["city"].GetStructValue().Fields)
	}

	if properties["order"] != nil {
		val8, _ := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(properties["order"])
		s.Order = val8.(int32)
	}

	if properties["updatedOn"] != nil {
		val9, _ := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(properties["updatedOn"])
		s.UpdatedOn = new(time.Time)
		*s.UpdatedOn = val9.(time.Time)
	}

	if properties["until"] != nil {
		val10, _ := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(properties["until"])
		s.Until = val10.(int32)
	}

	if properties["version"] != nil {
		val11, _ := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(properties["version"])
		s.Version = val11.(int32)
	}

}

func (s *TaxRate) ToProperties() map[string]*structpb.Value {
	var properties = make(map[string]*structpb.Value)

	val0, err := types.ByResourcePropertyType(model.ResourceProperty_FLOAT32).Pack(s.Rate)
	if err != nil {
		panic(err)
	}
	properties["rate"] = val0

	val1, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(s.CreatedBy)
	if err != nil {
		panic(err)
	}
	properties["createdBy"] = val1

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

	val4, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(s.Id)
	if err != nil {
		panic(err)
	}
	properties["id"] = val4

	val5, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(s.Name)
	if err != nil {
		panic(err)
	}
	properties["name"] = val5

	if s.Country != nil {
		properties["country"] = structpb.NewStructValue(&structpb.Struct{Fields: s.Country.ToProperties()})
	}

	if s.City != nil {
		properties["city"] = structpb.NewStructValue(&structpb.Struct{Fields: s.City.ToProperties()})
	}

	val8, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(s.Order)
	if err != nil {
		panic(err)
	}
	properties["order"] = val8

	if s.UpdatedOn != nil {
		val9, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*s.UpdatedOn)
		if err != nil {
			panic(err)
		}
		properties["updatedOn"] = val9
	}

	val10, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(s.Until)
	if err != nil {
		panic(err)
	}
	properties["until"] = val10

	val11, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(s.Version)
	if err != nil {
		panic(err)
	}
	properties["version"] = val11

	return properties
}

func (s *TaxRate) GetResourceName() string {
	return "tax_rate"
}

func (s *TaxRate) GetNamespace() string {
	return "default"
}

func (s *TaxRate) Clone() *TaxRate {
	var newInstance = new(TaxRate)
	newInstance.Rate = s.Rate
	newInstance.CreatedBy = s.CreatedBy
	if s.UpdatedBy != nil {
		newInstance.UpdatedBy = s.UpdatedBy
	}

	newInstance.CreatedOn = s.CreatedOn
	newInstance.Id = s.Id
	newInstance.Name = s.Name
	if s.Country != nil {
		newInstance.Country = s.Country
	}

	if s.City != nil {
		newInstance.City = s.City
	}

	newInstance.Order = s.Order
	if s.UpdatedOn != nil {
		newInstance.UpdatedOn = s.UpdatedOn
	}

	newInstance.Until = s.Until
	newInstance.Version = s.Version
	return newInstance
}

func (s *TaxRate) Equals(other *TaxRate) bool {
	return reflect.DeepEqual(s, other)
}

func (s *TaxRate) Same(other *TaxRate) bool {
	return s.Equals(other)
}

func NewTaxRateRepository(dhClient client.DhClient) client.Repository[*TaxRate] {
	return client.NewRepository[*TaxRate](dhClient, client.RepositoryParams[*TaxRate]{Instance: new(TaxRate)})
}
