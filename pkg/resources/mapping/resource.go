package mapping

import (
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
	properties["checkReferences"] = structpb.NewBoolValue(resource.CheckReferences)
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

	var propertiesStruct = make(map[string]*structpb.Value)
	for _, property := range resource.Properties {
		propertyRecord := ResourcePropertyToRecord(property, resource)
		propertiesStruct[property.Name] = structpb.NewStructValue(&structpb.Struct{Fields: propertyRecord.Properties})
	}
	properties["properties"] = structpb.NewStructValue(&structpb.Struct{Fields: propertiesStruct})

	if resource.Types != nil {
		var lv []*structpb.Value

		for _, subType := range resource.Types {
			lv = append(lv, ResourceTypeToValue(resource, subType))
		}

		properties["types"] = structpb.NewListValue(&structpb.ListValue{Values: lv})
	}

	MapSpecialColumnsToRecord(resource, &properties)

	if resource.Id != "" {
		properties["id"] = structpb.NewStringValue(resource.Id)
	}

	return &model.Record{
		Properties: properties,
	}
}

func ResourceFromRecord(record *model.Record) *model.Resource {
	if record == nil {
		return nil
	}

	var namespace = record.Properties["namespace"].GetStructValue().GetFields()["name"].GetStringValue()

	var resource = &model.Resource{
		Id:              record.Properties["id"].GetStringValue(),
		Name:            record.Properties["name"].GetStringValue(),
		Namespace:       namespace,
		Virtual:         record.Properties["virtual"].GetBoolValue(),
		Abstract:        record.Properties["abstract"].GetBoolValue(),
		Immutable:       record.Properties["immutable"].GetBoolValue(),
		CheckReferences: record.Properties["checkReferences"].GetBoolValue(),
		SourceConfig: &model.ResourceSourceConfig{
			DataSource: record.Properties["dataSource"].GetStructValue().GetFields()["name"].GetStringValue(),
			Entity:     record.Properties["entity"].GetStringValue(),
			Catalog:    record.Properties["catalog"].GetStringValue(),
		},
		Properties: util.ArrayMap(util.MapToArray(record.Properties["properties"].GetStructValue().Fields), func(t util.MapEntry[*structpb.Value]) *model.ResourceProperty {
			return ResourcePropertyFromRecord(t.Key, &model.Record{Properties: t.Val.GetStructValue().Fields})
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
	var ri = &model.ResourceIndex{
		Properties: nil,
		IndexType:  0,
	}

	st := val.GetStructValue()

	if st.Fields["annotations"] != nil {
		ri.Annotations = convertMap(st.Fields["annotations"].GetStructValue().AsMap(), func(v interface{}) string {
			return v.(string)
		})
	}

	if st.Fields["unique"] != nil {
		ri.Unique = st.Fields["unique"].GetBoolValue()
	}

	if st.Fields["properties"] != nil {
		list := st.Fields["properties"].GetListValue()

		for _, val := range list.Values {
			pVal := val.GetStructValue()
			ri.Properties = append(ri.Properties, &model.ResourceIndexProperty{
				Name:  pVal.Fields["name"].GetStringValue(),
				Order: model.Order(model.Order_value["ORDER_"+pVal.Fields["order"].GetStringValue()]),
			})
		}
	}

	return ri
}

func ResourceIndexToValue(index *model.ResourceIndex) *structpb.Value {
	annotations, err := structpb.NewValue(convertMap(index.Annotations, func(v string) interface{} {
		return v
	}))

	if err != nil {
		panic(err)
	}

	var propertyStructList []*structpb.Value

	for _, property := range index.Properties {
		propertyStructList = append(propertyStructList, structpb.NewStructValue(&structpb.Struct{Fields: map[string]*structpb.Value{
			"name":  structpb.NewStringValue(property.Name),
			"order": structpb.NewStringValue(property.Order.String()[len("ORDER_"):]),
		}}))
	}

	return structpb.NewStructValue(&structpb.Struct{
		Fields: map[string]*structpb.Value{
			"annotations": annotations,
			"unique":      structpb.NewBoolValue(index.Unique),
			"properties":  structpb.NewListValue(&structpb.ListValue{Values: propertyStructList}),
		},
	})
}

func ResourceTypeFromValue(val *structpb.Value) *model.ResourceSubType {
	var st = val.GetStructValue()
	rt := &model.ResourceSubType{
		Name: st.GetFields()["name"].GetStringValue(),
		Properties: util.ArrayMap(util.MapToArray(st.GetFields()["properties"].GetStructValue().Fields), func(t util.MapEntry[*structpb.Value]) *model.ResourceProperty {
			return ResourcePropertyFromRecord(t.Key, &model.Record{Properties: t.Val.GetStructValue().Fields})
		}),
	}

	if st.GetFields()["title"] != nil {
		rt.Title = st.GetFields()["title"].GetStringValue()
	}

	if st.GetFields()["description"] != nil {
		rt.Description = st.GetFields()["description"].GetStringValue()
	}

	return rt
}

func ResourceTypeToValue(resource *model.Resource, subType *model.ResourceSubType) *structpb.Value {
	var propertiesStruct = make(map[string]*structpb.Value)
	for _, property := range subType.Properties {
		propertyRecord := ResourcePropertyToRecord(property, resource)
		propertiesStruct[property.Name] = structpb.NewStructValue(&structpb.Struct{Fields: propertyRecord.Properties})
	}

	return structpb.NewStructValue(&structpb.Struct{
		Fields: map[string]*structpb.Value{
			"name":        structpb.NewStringValue(subType.Name),
			"title":       structpb.NewStringValue(subType.Title),
			"description": structpb.NewStringValue(subType.Description),
			"properties":  structpb.NewStructValue(&structpb.Struct{Fields: propertiesStruct}),
		},
	})
}
