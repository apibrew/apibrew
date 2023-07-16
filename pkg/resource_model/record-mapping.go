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
	return rec
}

func (m *RecordMapper) FromRecord(record *model.Record) *Record {
	return m.FromProperties(record.Properties)
}

func (m *RecordMapper) ToProperties(record *Record) map[string]*structpb.Value {
	var properties = make(map[string]*structpb.Value)

	var_8d10081b8b8d := record.Id

	if var_8d10081b8b8d != nil {
		var var_8d10081b8b8d_mapped *structpb.Value

		var var_8d10081b8b8d_err error
		var_8d10081b8b8d_mapped, var_8d10081b8b8d_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_8d10081b8b8d)
		if var_8d10081b8b8d_err != nil {
			panic(var_8d10081b8b8d_err)
		}
		properties["id"] = var_8d10081b8b8d_mapped
	}

	var_05cc9c7982eb := record.Properties

	var var_05cc9c7982eb_mapped *structpb.Value

	var var_05cc9c7982eb_err error
	var_05cc9c7982eb_mapped, var_05cc9c7982eb_err = types.ByResourcePropertyType(model.ResourceProperty_OBJECT).Pack(var_05cc9c7982eb)
	if var_05cc9c7982eb_err != nil {
		panic(var_05cc9c7982eb_err)
	}
	properties["properties"] = var_05cc9c7982eb_mapped

	var_be3939add396 := record.PackedProperties

	if var_be3939add396 != nil {
		var var_be3939add396_mapped *structpb.Value

		var var_be3939add396_l []*structpb.Value
		for _, value := range var_be3939add396 {

			var_d48f66ee3269 := value
			var var_d48f66ee3269_mapped *structpb.Value

			var var_d48f66ee3269_err error
			var_d48f66ee3269_mapped, var_d48f66ee3269_err = types.ByResourcePropertyType(model.ResourceProperty_OBJECT).Pack(var_d48f66ee3269)
			if var_d48f66ee3269_err != nil {
				panic(var_d48f66ee3269_err)
			}

			var_be3939add396_l = append(var_be3939add396_l, var_d48f66ee3269_mapped)
		}
		var_be3939add396_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_be3939add396_l})
		properties["packedProperties"] = var_be3939add396_mapped
	}
	return properties
}

func (m *RecordMapper) FromProperties(properties map[string]*structpb.Value) *Record {
	var s = m.New()
	if properties["id"] != nil {

		var_f41413a02c18 := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_f41413a02c18)

		if err != nil {
			panic(err)
		}

		var_f41413a02c18_mapped := new(uuid.UUID)
		*var_f41413a02c18_mapped = val.(uuid.UUID)

		s.Id = var_f41413a02c18_mapped
	}
	if properties["properties"] != nil {

		var_f0f58ceb521c := properties["properties"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_OBJECT).UnPack(var_f0f58ceb521c)

		if err != nil {
			panic(err)
		}

		var_f0f58ceb521c_mapped := val.(unstructured.Unstructured)

		s.Properties = var_f0f58ceb521c_mapped
	}
	if properties["packedProperties"] != nil {

		var_71f926d804e1 := properties["packedProperties"]
		var_71f926d804e1_mapped := []unstructured.Unstructured{}
		for _, v := range var_71f926d804e1.GetListValue().Values {

			var_6599e0ad0f3f := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_OBJECT).UnPack(var_6599e0ad0f3f)

			if err != nil {
				panic(err)
			}

			var_6599e0ad0f3f_mapped := val.(unstructured.Unstructured)

			var_71f926d804e1_mapped = append(var_71f926d804e1_mapped, var_6599e0ad0f3f_mapped)
		}

		s.PackedProperties = var_71f926d804e1_mapped
	}
	return s
}
