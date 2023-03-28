package model

import "time"
import "github.com/tislib/data-handler/pkg/model"
import "github.com/tislib/data-handler/pkg/client"
import "github.com/google/uuid"
import "github.com/tislib/data-handler/pkg/types"
import "google.golang.org/protobuf/types/known/structpb"

type RichTest3995 struct {
	Double    float64
	Text      string
	Id        uuid.UUID
	Timestamp time.Time
	Bool      bool
	Int64     int64
	Float     float32
	String    string
	Uuid      uuid.UUID
	Date      time.Time
	Int32O    *int32
	Int32     int32
	Object    map[string]interface{}
	UpdatedOn *time.Time
	Version   int32
	Time      time.Time
	Bytes     *[]uint8
	CreatedBy string
	UpdatedBy *string
	CreatedOn time.Time
}

func (s *RichTest3995) GetId() string {
	valStr := types.ByResourcePropertyType(model.ResourceProperty_UUID).String(s.Id)
	return valStr
}

func (s *RichTest3995) ToRecord() *model.Record {
	var rec = &model.Record{}
	rec.Properties = s.ToProperties()

	return rec
}

func (s *RichTest3995) FromRecord(record *model.Record) {
	s.FromProperties(record.Properties)
}

func (s *RichTest3995) FromProperties(properties map[string]*structpb.Value) {
	if properties["double"] != nil {
		val0, _ := types.ByResourcePropertyType(model.ResourceProperty_FLOAT64).UnPack(properties["double"])
		s.Double = val0.(float64)
	}

	if properties["text"] != nil {
		val1, _ := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(properties["text"])
		s.Text = val1.(string)
	}

	if properties["id"] != nil {
		val2, _ := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(properties["id"])
		s.Id = val2.(uuid.UUID)
	}

	if properties["timestamp"] != nil {
		val3, _ := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(properties["timestamp"])
		s.Timestamp = val3.(time.Time)
	}

	if properties["bool"] != nil {
		val4, _ := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(properties["bool"])
		s.Bool = val4.(bool)
	}

	if properties["int64"] != nil {
		val5, _ := types.ByResourcePropertyType(model.ResourceProperty_INT64).UnPack(properties["int64"])
		s.Int64 = val5.(int64)
	}

	if properties["float"] != nil {
		val6, _ := types.ByResourcePropertyType(model.ResourceProperty_FLOAT32).UnPack(properties["float"])
		s.Float = val6.(float32)
	}

	if properties["string"] != nil {
		val7, _ := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(properties["string"])
		s.String = val7.(string)
	}

	if properties["uuid"] != nil {
		val8, _ := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(properties["uuid"])
		s.Uuid = val8.(uuid.UUID)
	}

	if properties["date"] != nil {
		val9, _ := types.ByResourcePropertyType(model.ResourceProperty_DATE).UnPack(properties["date"])
		s.Date = val9.(time.Time)
	}

	if properties["int32_o"] != nil {
		val10, _ := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(properties["int32_o"])
		s.Int32O = new(int32)
		*s.Int32O = val10.(int32)
	}

	if properties["int32"] != nil {
		val11, _ := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(properties["int32"])
		s.Int32 = val11.(int32)
	}

	if properties["object"] != nil {
		val12, _ := types.ByResourcePropertyType(model.ResourceProperty_OBJECT).UnPack(properties["object"])
		s.Object = val12.(map[string]interface{})
	}

	if properties["updatedOn"] != nil {
		val13, _ := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(properties["updatedOn"])
		s.UpdatedOn = new(time.Time)
		*s.UpdatedOn = val13.(time.Time)
	}

	if properties["version"] != nil {
		val14, _ := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(properties["version"])
		s.Version = val14.(int32)
	}

	if properties["time"] != nil {
		val15, _ := types.ByResourcePropertyType(model.ResourceProperty_TIME).UnPack(properties["time"])
		s.Time = val15.(time.Time)
	}

	if properties["bytes"] != nil {
		val16, _ := types.ByResourcePropertyType(model.ResourceProperty_BYTES).UnPack(properties["bytes"])
		s.Bytes = new([]uint8)
		*s.Bytes = val16.([]uint8)
	}

	if properties["createdBy"] != nil {
		val17, _ := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(properties["createdBy"])
		s.CreatedBy = val17.(string)
	}

	if properties["updatedBy"] != nil {
		val18, _ := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(properties["updatedBy"])
		s.UpdatedBy = new(string)
		*s.UpdatedBy = val18.(string)
	}

	if properties["createdOn"] != nil {
		val19, _ := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(properties["createdOn"])
		s.CreatedOn = val19.(time.Time)
	}

}

func (s *RichTest3995) ToProperties() map[string]*structpb.Value {
	var properties = make(map[string]*structpb.Value)

	val0, err := types.ByResourcePropertyType(model.ResourceProperty_FLOAT64).Pack(s.Double)
	if err != nil {
		panic(err)
	}
	properties["double"] = val0

	val1, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(s.Text)
	if err != nil {
		panic(err)
	}
	properties["text"] = val1

	val2, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(s.Id)
	if err != nil {
		panic(err)
	}
	properties["id"] = val2

	val3, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(s.Timestamp)
	if err != nil {
		panic(err)
	}
	properties["timestamp"] = val3

	val4, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(s.Bool)
	if err != nil {
		panic(err)
	}
	properties["bool"] = val4

	val5, err := types.ByResourcePropertyType(model.ResourceProperty_INT64).Pack(s.Int64)
	if err != nil {
		panic(err)
	}
	properties["int64"] = val5

	val6, err := types.ByResourcePropertyType(model.ResourceProperty_FLOAT32).Pack(s.Float)
	if err != nil {
		panic(err)
	}
	properties["float"] = val6

	val7, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(s.String)
	if err != nil {
		panic(err)
	}
	properties["string"] = val7

	val8, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(s.Uuid)
	if err != nil {
		panic(err)
	}
	properties["uuid"] = val8

	val9, err := types.ByResourcePropertyType(model.ResourceProperty_DATE).Pack(s.Date)
	if err != nil {
		panic(err)
	}
	properties["date"] = val9

	if s.Int32O != nil {
		val10, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(*s.Int32O)
		if err != nil {
			panic(err)
		}
		properties["int32_o"] = val10
	}

	val11, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(s.Int32)
	if err != nil {
		panic(err)
	}
	properties["int32"] = val11

	val12, err := types.ByResourcePropertyType(model.ResourceProperty_OBJECT).Pack(s.Object)
	if err != nil {
		panic(err)
	}
	properties["object"] = val12

	if s.UpdatedOn != nil {
		val13, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*s.UpdatedOn)
		if err != nil {
			panic(err)
		}
		properties["updatedOn"] = val13
	}

	val14, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(s.Version)
	if err != nil {
		panic(err)
	}
	properties["version"] = val14

	val15, err := types.ByResourcePropertyType(model.ResourceProperty_TIME).Pack(s.Time)
	if err != nil {
		panic(err)
	}
	properties["time"] = val15

	if s.Bytes != nil {
		val16, err := types.ByResourcePropertyType(model.ResourceProperty_BYTES).Pack(*s.Bytes)
		if err != nil {
			panic(err)
		}
		properties["bytes"] = val16
	}

	val17, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(s.CreatedBy)
	if err != nil {
		panic(err)
	}
	properties["createdBy"] = val17

	if s.UpdatedBy != nil {
		val18, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*s.UpdatedBy)
		if err != nil {
			panic(err)
		}
		properties["updatedBy"] = val18
	}

	val19, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(s.CreatedOn)
	if err != nil {
		panic(err)
	}
	properties["createdOn"] = val19

	return properties
}

func (s *RichTest3995) GetResourceName() string {
	return "rich-test-3995"
}

func (s *RichTest3995) GetNamespace() string {
	return "default"
}

func NewRichTest3995Repository(dhClient client.DhClient) client.Repository[*RichTest3995] {
	return client.NewRepository[*RichTest3995](dhClient, client.RepositoryParams[*RichTest3995]{Instance: new(RichTest3995)})
}
