package mapping

import (
	"encoding/json"
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/util"
	"google.golang.org/protobuf/types/known/structpb"
)

func ResourceToRecord(resource *model.Resource) *model.Record {
	properties := make(map[string]*structpb.Value)

	properties["name"] = structpb.NewStringValue(resource.Name)
	properties["namespace"] = util.StructKv("name", resource.Namespace)
	properties["virtual"] = structpb.NewBoolValue(resource.Virtual)
	properties["immutable"] = structpb.NewBoolValue(resource.Immutable)
	if resource.SourceConfig != nil {
		properties["dataSource"] = util.StructKv("name", resource.SourceConfig.DataSource)
		properties["entity"] = structpb.NewStringValue(resource.SourceConfig.Entity)
		properties["catalog"] = structpb.NewStringValue(resource.SourceConfig.Catalog)
	}
	properties["type"] = structpb.NewNumberValue(float64(resource.DataType.Number()))
	properties["annotations"], _ = structpb.NewValue(convertMap(resource.Annotations, func(v string) interface{} {
		return v
	}))

	properties["securityContext"] = SecurityContextToValue(resource.SecurityContext)

	if resource.Indexes != nil {
		var lv []*structpb.Value

		for _, index := range resource.Indexes {
			lv = append(lv, ResourceIndexToValue(index))
		}

		properties["indexes"] = structpb.NewListValue(&structpb.ListValue{Values: lv})
	}

	return &model.Record{
		Id:         resource.Id,
		DataType:   resource.DataType,
		Properties: properties,
		AuditData:  resource.AuditData,
		Version:    resource.Version,
	}
}

func convertMap[T interface{}, K interface{}](annotations map[string]T, mapper func(v T) K) map[string]K {
	var result = make(map[string]K)

	for k, v := range annotations {
		result[k] = mapper(v)
	}

	return result
}

func ResourceFromRecord(record *model.Record) *model.Resource {
	if record == nil {
		return nil
	}

	var resource = &model.Resource{
		Id:        record.Id,
		DataType:  model.DataType(int32(record.Properties["type"].GetNumberValue())),
		AuditData: record.AuditData,
		Version:   record.Version,
		Name:      record.Properties["name"].GetStringValue(),
		Namespace: record.Properties["namespace"].GetStructValue().GetFields()["name"].GetStringValue(),
		Virtual:   record.Properties["virtual"].GetBoolValue(),
		Immutable: record.Properties["immutable"].GetBoolValue(),
		SourceConfig: &model.ResourceSourceConfig{
			DataSource: record.Properties["dataSource"].GetStructValue().GetFields()["name"].GetStringValue(),
			Entity:     record.Properties["entity"].GetStringValue(),
			Catalog:    record.Properties["catalog"].GetStringValue(),
		},
		SecurityContext: SecurityContextFromValue(record.Properties["securityContext"]),
		Annotations: convertMap(record.Properties["annotations"].GetStructValue().AsMap(), func(v interface{}) string {
			return v.(string)
		}),
	}

	if record.Properties["indexes"] != nil {
		list := record.Properties["indexes"].GetListValue()

		for _, val := range list.Values {
			resource.Indexes = append(resource.Indexes, ResourceIndexFromValue(val))
		}
	}

	return resource
}

func ResourceIndexFromValue(val *structpb.Value) *model.ResourceIndex {
	jData, err := val.MarshalJSON()

	if err != nil {
		panic(err)
	}

	var ri = new(model.ResourceIndex)

	err = json.Unmarshal(jData, ri)

	if err != nil {
		panic(err)
	}

	return ri
}

func ResourceIndexToValue(index *model.ResourceIndex) *structpb.Value {
	jData, err := json.Marshal(index)

	if err != nil {
		panic(err)
	}

	var val = &structpb.Value{}

	err = val.UnmarshalJSON(jData)

	if err != nil {
		panic(err)
	}

	return val
}
