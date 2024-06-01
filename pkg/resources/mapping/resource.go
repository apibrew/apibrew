package mapping

import (
	"github.com/apibrew/apibrew/pkg/abs"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/util"
	"google.golang.org/protobuf/types/known/structpb"
)

func ResourceToRecord(resource *model.Resource) abs.RecordLike {
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

		lv = make([]*structpb.Value, 0)

		for _, index := range resource.Indexes {
			lv = append(lv, ResourceIndexToValue(index))
		}

		properties["indexes"] = structpb.NewListValue(&structpb.ListValue{Values: lv})
	}

	var propertiesStruct = make(map[string]*structpb.Value)
	for _, property := range resource.Properties {
		propertyRecord := ResourcePropertyToRecord(property, resource)
		propertiesStruct[property.Name] = structpb.NewStructValue(propertyRecord.ToStruct())
	}
	properties["properties"] = structpb.NewStructValue(&structpb.Struct{Fields: propertiesStruct})

	if resource.Types != nil {
		var lv []*structpb.Value

		lv = make([]*structpb.Value, 0)

		for _, subType := range resource.Types {
			lv = append(lv, ResourceTypeToValue(resource, subType))
		}

		properties["types"] = structpb.NewListValue(&structpb.ListValue{Values: lv})
	}

	MapSpecialColumnsToRecord(resource, &properties)

	if resource.Id != "" {
		properties["id"] = structpb.NewStringValue(resource.Id)
	}

	return abs.NewRecordLikeWithProperties(properties)
}

func ResourceFromRecord(record abs.RecordLike) *model.Resource {
	if record == nil {
		return nil
	}

	var namespace = record.GetStructProperty("namespace").GetStructValue().GetFields()["name"].GetStringValue()

	var resource = &model.Resource{
		Id:              record.GetStructProperty("id").GetStringValue(),
		Name:            record.GetStructProperty("name").GetStringValue(),
		Namespace:       namespace,
		Virtual:         record.GetStructProperty("virtual").GetBoolValue(),
		Abstract:        record.GetStructProperty("abstract").GetBoolValue(),
		Immutable:       record.GetStructProperty("immutable").GetBoolValue(),
		CheckReferences: record.GetStructProperty("checkReferences").GetBoolValue(),
		SourceConfig: &model.ResourceSourceConfig{
			DataSource: record.GetStructProperty("dataSource").GetStructValue().GetFields()["name"].GetStringValue(),
			Entity:     record.GetStructProperty("entity").GetStringValue(),
			Catalog:    record.GetStructProperty("catalog").GetStringValue(),
		},
		Properties: util.ArrayMap(util.MapToArray(record.GetStructProperty("properties").GetStructValue().Fields), func(t util.MapEntry[*structpb.Value]) *model.ResourceProperty {
			return ResourcePropertyFromRecord(t.Key, abs.NewRecordLikeWithProperties(t.Val.GetStructValue().Fields))
		}),
		Annotations: convertMap(record.GetStructProperty("annotations").GetStructValue().AsMap(), func(v interface{}) string {
			return v.(string)
		}),
	}

	if record.HasProperty("indexes") {
		list := record.GetStructProperty("indexes").GetListValue()

		resource.Indexes = make([]*model.ResourceIndex, 0)

		for _, val := range list.Values {
			resource.Indexes = append(resource.Indexes, ResourceIndexFromValue(val))
		}
	}

	if record.HasProperty("types") {
		list := record.GetStructProperty("types").GetListValue()

		resource.Types = make([]*model.ResourceSubType, 0)

		for _, val := range list.Values {
			resource.Types = append(resource.Types, ResourceTypeFromValue(val))
		}
	}

	if record.HasProperty("title") {
		resource.Title = new(string)
		*resource.Title = record.GetStructProperty("title").GetStringValue()
	}

	if record.HasProperty("description") {
		resource.Description = new(string)
		*resource.Description = record.GetStructProperty("description").GetStringValue()
	}

	var properties = record.ToStruct().GetFields()
	MapSpecialColumnsFromRecord(resource, &properties)

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
			return ResourcePropertyFromRecord(t.Key, abs.NewRecordLikeWithProperties(t.Val.GetStructValue().Fields))
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
		propertiesStruct[property.Name] = structpb.NewStructValue(propertyRecord.ToStruct())
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
