package model

import "reflect"
import "github.com/tislib/apibrew/pkg/helper"
import "github.com/tislib/apibrew/pkg/model"
import "github.com/tislib/apibrew/pkg/client"
import "github.com/google/uuid"
import "github.com/tislib/apibrew/pkg/types"
import "google.golang.org/protobuf/types/known/structpb"

type TaxRate struct {
	Id uuid.UUID

	Name string

	Country *Country

	City *City

	Order int32

	Until int32

	Rate float32

	Version int32
}

func (s *TaxRate) GetId() uuid.UUID {
	return s.Id
}

func (s *TaxRate) GetName() string {
	return s.Name
}

func (s *TaxRate) GetCountry() *Country {
	return s.Country
}

func (s *TaxRate) GetCity() *City {
	return s.City
}

func (s *TaxRate) GetOrder() int32 {
	return s.Order
}

func (s *TaxRate) GetUntil() int32 {
	return s.Until
}

func (s *TaxRate) GetRate() float32 {
	return s.Rate
}

func (s *TaxRate) GetVersion() int32 {
	return s.Version
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

		s.Name = val.(string)

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

		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(properties["order"])

		if err != nil {
			panic(err)
		}

		s.Order = val.(int32)

	}

	if properties["until"] != nil {

		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(properties["until"])

		if err != nil {
			panic(err)
		}

		s.Until = val.(int32)

	}

	if properties["rate"] != nil {

		val, err := types.ByResourcePropertyType(model.ResourceProperty_FLOAT32).UnPack(properties["rate"])

		if err != nil {
			panic(err)
		}

		s.Rate = val.(float32)

	}

	if properties["version"] != nil {

		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(properties["version"])

		if err != nil {
			panic(err)
		}

		s.Version = val.(int32)

	}

}

func (s *TaxRate) ToProperties() map[string]*structpb.Value {
	var properties = make(map[string]*structpb.Value)

	Id, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(s.Id)
	if err != nil {
		panic(err)
	}
	properties["id"] = Id

	Name, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(s.Name)
	if err != nil {
		panic(err)
	}
	properties["name"] = Name

	if s.Country != nil {

		properties["country"] = structpb.NewStructValue(&structpb.Struct{Fields: s.Country.ToProperties()})

	}

	if s.City != nil {

		properties["city"] = structpb.NewStructValue(&structpb.Struct{Fields: s.City.ToProperties()})

	}

	Order, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(s.Order)
	if err != nil {
		panic(err)
	}
	properties["order"] = Order

	Until, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(s.Until)
	if err != nil {
		panic(err)
	}
	properties["until"] = Until

	Rate, err := types.ByResourcePropertyType(model.ResourceProperty_FLOAT32).Pack(s.Rate)
	if err != nil {
		panic(err)
	}
	properties["rate"] = Rate

	Version, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(s.Version)
	if err != nil {
		panic(err)
	}
	properties["version"] = Version

	return properties
}

func (s *TaxRate) GetResourceName() string {
	return "tax_rate"
}

func (s *TaxRate) GetNamespace() string {
	return "default"
}

func (s *TaxRate) Equals(other *TaxRate) bool {
	return reflect.DeepEqual(s, other)
}

func (s *TaxRate) Same(other *TaxRate) bool {
	return s.Equals(other)
}

func NewTaxRateRepository(dhClient client.DhClient) client.Repository[*TaxRate] {
	return client.NewRepository[*TaxRate](dhClient, client.RepositoryParams[*TaxRate]{InstanceProvider: func() *TaxRate {
		return new(TaxRate)
	}})
}

var TaxRateId = client.DefineProperty[uuid.UUID, helper.UuidQueryBuilder]("id", model.ResourceProperty_UUID, helper.UuidQueryBuilder{PropName: "id"})

var TaxRateName = client.DefineProperty[string, helper.StringQueryBuilder]("name", model.ResourceProperty_STRING, helper.StringQueryBuilder{PropName: "name"})

var TaxRateCountry = client.DefineProperty[*Country, helper.ReferenceQueryBuilder[*Country]]("country", model.ResourceProperty_REFERENCE, helper.ReferenceQueryBuilder[*Country]{PropName: "country"})

var TaxRateCity = client.DefineProperty[*City, helper.ReferenceQueryBuilder[*City]]("city", model.ResourceProperty_REFERENCE, helper.ReferenceQueryBuilder[*City]{PropName: "city"})

var TaxRateOrder = client.DefineProperty[int32, helper.Int32QueryBuilder]("order", model.ResourceProperty_INT32, helper.Int32QueryBuilder{PropName: "order"})

var TaxRateUntil = client.DefineProperty[int32, helper.Int32QueryBuilder]("until", model.ResourceProperty_INT32, helper.Int32QueryBuilder{PropName: "until"})

var TaxRateRate = client.DefineProperty[float32, helper.Float32QueryBuilder]("rate", model.ResourceProperty_FLOAT32, helper.Float32QueryBuilder{PropName: "rate"})

var TaxRateVersion = client.DefineProperty[int32, helper.Int32QueryBuilder]("version", model.ResourceProperty_INT32, helper.Int32QueryBuilder{PropName: "version"})
