package mapping

import (
	"data-handler/model"
	"data-handler/service/system"
	"google.golang.org/protobuf/types/known/structpb"
)

func ResourceToRecord(resource *model.Resource) *model.Record {
	properties := make(map[string]*structpb.Value)

	properties["name"] = structpb.NewStringValue(resource.Name)
	properties["namespace"] = structpb.NewStringValue(resource.Namespace)
	properties["dataSource"] = structpb.NewStringValue(resource.SourceConfig.DataSource)
	properties["entity"] = structpb.NewStringValue(resource.SourceConfig.Entity)
	properties["catalog"] = structpb.NewStringValue(resource.SourceConfig.Catalog)
	properties["type"] = structpb.NewNumberValue(float64(resource.DataType.Number()))
	properties["annotations"], _ = structpb.NewValue(convertMap(resource.Annotations, func(v string) interface{} {
		return v
	}))

	return &model.Record{
		Id:         resource.Id,
		Resource:   system.ResourceResource.Name,
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
		Namespace: record.Properties["namespace"].GetStringValue(),
		SourceConfig: &model.ResourceSourceConfig{
			DataSource: record.Properties["dataSource"].GetStringValue(),
			Entity:     record.Properties["entity"].GetStringValue(),
			Catalog:    record.Properties["catalog"].GetStringValue(),
		},
		Annotations: convertMap(record.Properties["annotations"].GetStructValue().AsMap(), func(v interface{}) string {
			return v.(string)
		}),
	}

	return resource
}
