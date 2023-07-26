package mapping

import (
	"encoding/json"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/util"
	"google.golang.org/protobuf/types/known/structpb"
)

func ResourceToRecord(resource *model.Resource) *model.Record {
	properties := make(map[string]*structpb.Value)

	properties["name"] = structpb.NewStringValue(resource.Name)
	if resource.Title != nil {
		properties["title"] = structpb.NewStringValue(*resource.Title)
	}
	if resource.Description != nil {
		properties["description"] = structpb.NewStringValue(*resource.Description)
	}
	properties["namespace"] = util.StructKv("name", resource.Namespace)
	properties["virtual"] = structpb.NewBoolValue(resource.Virtual)
	properties["abstract"] = structpb.NewBoolValue(resource.Abstract)
	properties["immutable"] = structpb.NewBoolValue(resource.Immutable)
	if resource.SourceConfig != nil {
		properties["dataSource"] = util.StructKv("name", resource.SourceConfig.DataSource)
		properties["entity"] = structpb.NewStringValue(resource.SourceConfig.Entity)
		properties["catalog"] = structpb.NewStringValue(resource.SourceConfig.Catalog)
	}
	properties["annotations"], _ = structpb.NewValue(convertMap(resource.Annotations, func(v string) interface{} {
		return v
	}))

	if resource.Indexes != nil {
		var lv []*structpb.Value

		for _, index := range resource.Indexes {
			lv = append(lv, ResourceIndexToValue(index))
		}

		properties["indexes"] = structpb.NewListValue(&structpb.ListValue{Values: lv})
	}

	var propertyStructList []*structpb.Value
	for _, property := range resource.Properties {
		propertyRecord := ResourcePropertyToRecord(property, resource)
		propertyStructList = append(propertyStructList, structpb.NewStructValue(&structpb.Struct{Fields: propertyRecord.Properties}))
	}
	properties["properties"] = structpb.NewListValue(&structpb.ListValue{Values: propertyStructList})

	if resource.Types != nil {
		var lv []*structpb.Value

		for _, subType := range resource.Types {
			lv = append(lv, ResourceTypeToValue(subType))
		}

		properties["types"] = structpb.NewListValue(&structpb.ListValue{Values: lv})
	}

	MapSpecialColumnsToRecord(resource, &properties)

	return &model.Record{
		Id:         resource.Id,
		Properties: properties,
	}
}

func ResourceFromRecord(record *model.Record) *model.Resource {
	if record == nil {
		return nil
	}

	var resource = &model.Resource{
		Id:        record.Id,
		Name:      record.Properties["name"].GetStringValue(),
		Namespace: record.Properties["namespace"].GetStructValue().GetFields()["name"].GetStringValue(),
		Virtual:   record.Properties["virtual"].GetBoolValue(),
		Abstract:  record.Properties["abstract"].GetBoolValue(),
		Immutable: record.Properties["immutable"].GetBoolValue(),
		SourceConfig: &model.ResourceSourceConfig{
			DataSource: record.Properties["dataSource"].GetStructValue().GetFields()["name"].GetStringValue(),
			Entity:     record.Properties["entity"].GetStringValue(),
			Catalog:    record.Properties["catalog"].GetStringValue(),
		},
		Properties: util.ArrayMap(record.Properties["properties"].GetListValue().Values, func(t *structpb.Value) *model.ResourceProperty {
			return ResourcePropertyFromRecord(&model.Record{Properties: t.GetStructValue().Fields})
		}),
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

	if record.Properties["types"] != nil {
		list := record.Properties["types"].GetListValue()

		for _, val := range list.Values {
			resource.Types = append(resource.Types, ResourceTypeFromValue(val))
		}
	}

	if record.Properties["title"] != nil {
		resource.Title = new(string)
		*resource.Title = record.Properties["title"].GetStringValue()
	}

	if record.Properties["description"] != nil {
		resource.Description = new(string)
		*resource.Description = record.Properties["description"].GetStringValue()
	}

	MapSpecialColumnsFromRecord(resource, &record.Properties)

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

func ResourceTypeFromValue(val *structpb.Value) *model.ResourceSubType {
	jData, err := val.MarshalJSON()

	if err != nil {
		panic(err)
	}

	var ri = new(model.ResourceSubType)

	err = json.Unmarshal(jData, ri)

	if err != nil {
		panic(err)
	}

	return ri
}

func ResourceTypeToValue(index *model.ResourceSubType) *structpb.Value {
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
