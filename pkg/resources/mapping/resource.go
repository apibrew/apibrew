package mapping

import (
	"github.com/apibrew/apibrew/pkg/abs"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/util"
	"google.golang.org/protobuf/types/known/structpb"
)

func ResourceToRecord(resource *model.Resource) abs.RecordLike {
	properties := make(map[string]interface{})

	properties["name"] = resource.Name
	if resource.Title != nil {
		properties["title"] = *resource.Title
	}
	if resource.Description != nil {
		properties["description"] = *resource.Description
	}
	properties["namespace"] = util.StructKv("name", resource.Namespace)
	properties["virtual"] = structpb.NewBoolValue(resource.Virtual)
	properties["abstract"] = structpb.NewBoolValue(resource.Abstract)
	properties["immutable"] = structpb.NewBoolValue(resource.Immutable)
	properties["checkReferences"] = structpb.NewBoolValue(resource.CheckReferences)
	if resource.SourceConfig != nil {
		properties["dataSource"] = util.StructKv("name", resource.SourceConfig.DataSource)
		properties["entity"] = resource.SourceConfig.Entity
		properties["catalog"] = resource.SourceConfig.Catalog
	}
	properties["annotations"], _ = structpb.NewValue(convertMap(resource.Annotations, func(v string) interface{} {
		return v
	}))

	if resource.Indexes != nil {
		var lv []interface{}

		lv = make([]interface{}, 0)

		for _, index := range resource.Indexes {
			lv = append(lv, ResourceIndexToValue(index))
		}

		properties["indexes"] = lv
	}

	var propertiesStruct = make(map[string]interface{})
	for _, property := range resource.Properties {
		propertyRecord := ResourcePropertyToRecord(property, resource)
		propertiesStruct[property.Name] = structpb.NewStructValue(propertyRecord.ToStruct())
	}
	properties["properties"] = propertiesStruct

	if resource.Types != nil {
		var lv []interface{}

		lv = make([]interface{}, 0)

		for _, subType := range resource.Types {
			lv = append(lv, ResourceTypeToValue(resource, subType))
		}

		properties["types"] = lv
	}

	MapSpecialColumnsToRecord(resource, &properties)

	if resource.Id != "" {
		properties["id"] = resource.Id
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
		Properties: util.ArrayMap(util.MapToArray(record.GetStructProperty("properties").GetStructValue().Fields), func(t util.MapEntry[interface{}]) *model.ResourceProperty {
			return ResourcePropertyFromRecord(t.Key, abs.NewRecordLikeWithStructProperties(t.Val.(map[string]interface{})))
		}),
		Annotations: convertMap(record.GetStructProperty("annotations").GetStructValue().AsMap(), func(v interface{}) string {
			return v.(string)
		}),
	}

	if record.HasProperty("indexes") {
		list := record.GetProperty("indexes").([]interface{})

		resource.Indexes = make([]*model.ResourceIndex, 0)

		for _, val := range list {
			resource.Indexes = append(resource.Indexes, ResourceIndexFromValue(val.(map[string]interface{})))
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

func ResourceIndexFromValue(val map[string]interface{}) *model.ResourceIndex {
	var ri = &model.ResourceIndex{
		Properties: nil,
		IndexType:  0,
	}

	if val["annotations"] != nil {
		ri.Annotations = convertMap(val["annotations"].(map[string]interface{}), func(v interface{}) string {
			return v.(string)
		})
	}

	if val["unique"] != nil {
		ri.Unique = val["unique"].(bool)
	}

	if val["properties"] != nil {
		list := val["properties"].(bool)

		for _, val := range list {
			pVal := val.GetStructValue()
			ri.Properties = append(ri.Properties, &model.ResourceIndexProperty{
				Name:  pVal.Fields["name"].GetStringValue(),
				Order: model.Order(model.Order_value["ORDER_"+pVal.Fields["order"].GetStringValue()]),
			})
		}
	}

	return ri
}

func ResourceIndexToValue(index *model.ResourceIndex) interface{} {
	annotations, err := structpb.NewValue(convertMap(index.Annotations, func(v string) interface{} {
		return v
	}))

	if err != nil {
		panic(err)
	}

	var propertyStructList []interface{}

	for _, property := range index.Properties {
		propertyStructList = append(propertyStructList, map[string]interface{}{
			"name":  property.Name,
			"order": property.Order.String()[len("ORDER_"):],
		})
	}

	return map[string]interface{}{
		"annotations": annotations,
		"unique":      index.Unique,
		"properties":  propertyStructList,
	}
}

func ResourceTypeFromValue(val interface{}) *model.ResourceSubType {
	var st = val.GetStructValue()
	rt := &model.ResourceSubType{
		Name: st.GetFields()["name"].GetStringValue(),
		Properties: util.ArrayMap(util.MapToArray(st.GetFields()["properties"].GetStructValue().Fields), func(t util.MapEntry[interface{}]) *model.ResourceProperty {
			return ResourcePropertyFromRecord(t.Key, abs.NewRecordLikeWithStructProperties(t.Val.GetStructValue().Fields))
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

func ResourceTypeToValue(resource *model.Resource, subType *model.ResourceSubType) interface{} {
	var propertiesStruct = make(map[string]interface{})
	for _, property := range subType.Properties {
		propertyRecord := ResourcePropertyToRecord(property, resource)
		propertiesStruct[property.Name] = structpb.NewStructValue(propertyRecord.ToStruct())
	}

	return structpb.NewStructValue(&structpb.Struct{
		Fields: map[string]interface{}{
			"name":        subType.Name,
			"title":       subType.Title,
			"description": subType.Description,
			"properties":  structpb.NewStructValue(&structpb.Struct{Fields: propertiesStruct}),
		},
	})
}
