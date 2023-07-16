package resource_model

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/types"
	"google.golang.org/protobuf/types/known/structpb"
)

import "github.com/google/uuid"
import "github.com/apibrew/apibrew/pkg/formats/unstructured"

type RecordMapper struct {
}

func NewRecordMapper() *RecordMapper {
	return &RecordMapper{}
}

var RecordMapperInstance = NewRecordMapper()

func (m *RecordMapper) New() *Record {
	return &Record{}
}

func (m *RecordMapper) ToRecord(record *Record) *model.Record {
	var rec = &model.Record{}
	rec.Properties = m.ToProperties(record)

	if record.Id != nil {
		rec.Id = record.Id.String()
	}

	return rec
}

func (m *RecordMapper) FromRecord(record *model.Record) *Record {
	return m.FromProperties(record.Properties)
}

func (m *RecordMapper) ToProperties(record *Record) map[string]*structpb.Value {
	var properties = make(map[string]*structpb.Value)

	var_c46cb1a24f5d := record.Id

	if var_c46cb1a24f5d != nil {
		var var_c46cb1a24f5d_mapped *structpb.Value

		var var_c46cb1a24f5d_err error
		var_c46cb1a24f5d_mapped, var_c46cb1a24f5d_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_c46cb1a24f5d)
		if var_c46cb1a24f5d_err != nil {
			panic(var_c46cb1a24f5d_err)
		}
		properties["id"] = var_c46cb1a24f5d_mapped
	}

	var_9badc1cf7ea5 := record.Properties

	var var_9badc1cf7ea5_mapped *structpb.Value

	var var_9badc1cf7ea5_err error
	var_9badc1cf7ea5_mapped, var_9badc1cf7ea5_err = types.ByResourcePropertyType(model.ResourceProperty_OBJECT).Pack(var_9badc1cf7ea5)
	if var_9badc1cf7ea5_err != nil {
		panic(var_9badc1cf7ea5_err)
	}
	properties["properties"] = var_9badc1cf7ea5_mapped

	var_f1c5cff28eb8 := record.PackedProperties

	if var_f1c5cff28eb8 != nil {
		var var_f1c5cff28eb8_mapped *structpb.Value

		var var_f1c5cff28eb8_l []*structpb.Value
		for _, value := range var_f1c5cff28eb8 {

			var_bfc5db834da5 := value
			var var_bfc5db834da5_mapped *structpb.Value

			var var_bfc5db834da5_err error
			var_bfc5db834da5_mapped, var_bfc5db834da5_err = types.ByResourcePropertyType(model.ResourceProperty_OBJECT).Pack(var_bfc5db834da5)
			if var_bfc5db834da5_err != nil {
				panic(var_bfc5db834da5_err)
			}

			var_f1c5cff28eb8_l = append(var_f1c5cff28eb8_l, var_bfc5db834da5_mapped)
		}
		var_f1c5cff28eb8_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_f1c5cff28eb8_l})
		properties["packedProperties"] = var_f1c5cff28eb8_mapped
	}
	return properties
}

func (m *RecordMapper) FromProperties(properties map[string]*structpb.Value) *Record {
	var s = m.New()
	if properties["id"] != nil {

		var_184bf597e04d := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_184bf597e04d)

		if err != nil {
			panic(err)
		}

		var_184bf597e04d_mapped := new(uuid.UUID)
		*var_184bf597e04d_mapped = val.(uuid.UUID)

		s.Id = var_184bf597e04d_mapped
	}
	if properties["properties"] != nil {

		var_737b56a82fff := properties["properties"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_OBJECT).UnPack(var_737b56a82fff)

		if err != nil {
			panic(err)
		}

		var_737b56a82fff_mapped := val.(unstructured.Unstructured)

		s.Properties = var_737b56a82fff_mapped
	}
	if properties["packedProperties"] != nil {

		var_3c5dc60d039e := properties["packedProperties"]
		var_3c5dc60d039e_mapped := []unstructured.Unstructured{}
		for _, v := range var_3c5dc60d039e.GetListValue().Values {

			var_14e51431feed := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_OBJECT).UnPack(var_14e51431feed)

			if err != nil {
				panic(err)
			}

			var_14e51431feed_mapped := val.(unstructured.Unstructured)

			var_3c5dc60d039e_mapped = append(var_3c5dc60d039e_mapped, var_14e51431feed_mapped)
		}

		s.PackedProperties = var_3c5dc60d039e_mapped
	}
	return s
}
