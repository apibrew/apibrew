package mapping

import (
	"data-handler/model"
	"data-handler/service/system"
	"google.golang.org/protobuf/types/known/structpb"
)

func ResourceToRecord(resource *model.Resource) *model.Record {
	properties := make(map[string]interface{})

	properties["name"] = resource.Name
	properties["namespace"] = resource.Namespace
	properties["dataSource"] = resource.SourceConfig.DataSource
	properties["entity"] = resource.SourceConfig.Entity
	properties["catalog"] = resource.SourceConfig.Catalog
	properties["type"] = int32(resource.DataType.Number())
	properties["annotations"] = convertMap(resource.Annotations, func(v string) interface{} {
		return v
	})

	structProperties, err := structpb.NewStruct(properties)

	if err != nil {
		panic(err)
	}

	return &model.Record{
		Id:         resource.Id,
		Resource:   system.ResourceResource.Name,
		DataType:   resource.DataType,
		Properties: structProperties,
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
		DataType:  model.DataType(int32(record.Properties.AsMap()["type"].(float64))),
		AuditData: record.AuditData,
		Version:   record.Version,
		Name:      record.Properties.AsMap()["name"].(string),
		Namespace: record.Properties.AsMap()["namespace"].(string),
		SourceConfig: &model.ResourceSourceConfig{
			DataSource: record.Properties.AsMap()["dataSource"].(string),
			Entity:     record.Properties.AsMap()["entity"].(string),
			Catalog:    record.Properties.AsMap()["catalog"].(string),
		},
		Annotations: convertMap(record.Properties.AsMap()["annotations"].(map[string]interface{}), func(v interface{}) string {
			return v.(string)
		}),
	}

	return resource
}
