package model

import "reflect"
import "github.com/tislib/apibrew/pkg/helper"
import "github.com/tislib/apibrew/pkg/model"
import "github.com/tislib/apibrew/pkg/client"
import "github.com/google/uuid"
import "github.com/tislib/apibrew/pkg/types"
import "google.golang.org/protobuf/types/known/structpb"

type Income struct {
	Id uuid.UUID

	Country *Country

	City *City

	GrossIncome int32

	Tax *int32

	NetIncome *int32

	Version int32
}

func (s *Income) GetId() uuid.UUID {
	return s.Id
}

func (s *Income) GetCountry() *Country {
	return s.Country
}

func (s *Income) GetCity() *City {
	return s.City
}

func (s *Income) GetGrossIncome() int32 {
	return s.GrossIncome
}

func (s *Income) GetTax() *int32 {
	return s.Tax
}

func (s *Income) GetNetIncome() *int32 {
	return s.NetIncome
}

func (s *Income) GetVersion() int32 {
	return s.Version
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

	if properties["id"] != nil {

		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(properties["id"])

		if err != nil {
			panic(err)
		}

		s.Id = val.(uuid.UUID)

	}

	if properties["country"] != nil {

		s.Country = new(Country)
		s.Country.FromProperties(properties["country"].GetStructValue().Fields)

	}

	if properties["city"] != nil {

		s.City = new(City)
		s.City.FromProperties(properties["city"].GetStructValue().Fields)

	}

	if properties["gross_income"] != nil {

		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(properties["gross_income"])

		if err != nil {
			panic(err)
		}

		s.GrossIncome = val.(int32)

	}

	if properties["tax"] != nil {

		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(properties["tax"])

		if err != nil {
			panic(err)
		}

		s.Tax = new(int32)
		*s.Tax = val.(int32)

	}

	if properties["net_income"] != nil {

		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(properties["net_income"])

		if err != nil {
			panic(err)
		}

		s.NetIncome = new(int32)
		*s.NetIncome = val.(int32)

	}

	if properties["version"] != nil {

		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(properties["version"])

		if err != nil {
			panic(err)
		}

		s.Version = val.(int32)

	}

}

func (s *Income) ToProperties() map[string]*structpb.Value {
	var properties = make(map[string]*structpb.Value)

	Id, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(s.Id)
	if err != nil {
		panic(err)
	}
	properties["id"] = Id

	if s.Country != nil {

		properties["country"] = structpb.NewStructValue(&structpb.Struct{Fields: s.Country.ToProperties()})

	}

	if s.City != nil {

		properties["city"] = structpb.NewStructValue(&structpb.Struct{Fields: s.City.ToProperties()})

	}

	GrossIncome, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(s.GrossIncome)
	if err != nil {
		panic(err)
	}
	properties["gross_income"] = GrossIncome

	if s.Tax != nil {

		Tax, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(*s.Tax)
		if err != nil {
			panic(err)
		}
		properties["tax"] = Tax

	}

	if s.NetIncome != nil {

		NetIncome, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(*s.NetIncome)
		if err != nil {
			panic(err)
		}
		properties["net_income"] = NetIncome

	}

	Version, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(s.Version)
	if err != nil {
		panic(err)
	}
	properties["version"] = Version

	return properties
}

func (s *Income) GetResourceName() string {
	return "income"
}

func (s *Income) GetNamespace() string {
	return "default"
}

func (s *Income) Equals(other *Income) bool {
	return reflect.DeepEqual(s, other)
}

func (s *Income) Same(other *Income) bool {
	return s.Equals(other)
}

func NewIncomeRepository(dhClient client.DhClient) client.Repository[*Income] {
	return client.NewRepository[*Income](dhClient, client.RepositoryParams[*Income]{InstanceProvider: func() *Income {
		return new(Income)
	}})
}

var IncomeId = client.DefineProperty[uuid.UUID, helper.UuidQueryBuilder]("id", model.ResourceProperty_UUID, helper.UuidQueryBuilder{PropName: "id"})

var IncomeCountry = client.DefineProperty[*Country, helper.ReferenceQueryBuilder[*Country]]("country", model.ResourceProperty_REFERENCE, helper.ReferenceQueryBuilder[*Country]{PropName: "country"})

var IncomeCity = client.DefineProperty[*City, helper.ReferenceQueryBuilder[*City]]("city", model.ResourceProperty_REFERENCE, helper.ReferenceQueryBuilder[*City]{PropName: "city"})

var IncomeGrossIncome = client.DefineProperty[int32, helper.Int32QueryBuilder]("gross_income", model.ResourceProperty_INT32, helper.Int32QueryBuilder{PropName: "gross_income"})

var IncomeTax = client.DefineProperty[int32, helper.Int32QueryBuilder]("tax", model.ResourceProperty_INT32, helper.Int32QueryBuilder{PropName: "tax"})

var IncomeNetIncome = client.DefineProperty[int32, helper.Int32QueryBuilder]("net_income", model.ResourceProperty_INT32, helper.Int32QueryBuilder{PropName: "net_income"})

var IncomeVersion = client.DefineProperty[int32, helper.Int32QueryBuilder]("version", model.ResourceProperty_INT32, helper.Int32QueryBuilder{PropName: "version"})
