package mapping

import (
	"github.com/apibrew/apibrew/pkg/abs"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/util"
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
	properties["namespace"] = util.Kv("name", resource.Namespace)
	properties["virtual"] = resource.Virtual
	properties["abstract"] = resource.Abstract
	properties["immutable"] = resource.Immutable
	properties["checkReferences"] = resource.CheckReferences
	if resource.SourceConfig != nil {
		properties["dataSource"] = util.Kv("name", resource.SourceConfig.DataSource)
		properties["entity"] = resource.SourceConfig.Entity
		properties["catalog"] = resource.SourceConfig.Catalog
	}
	properties["annotations"] = convertMap(resource.Annotations, func(v string) interface{} {
		return v
	})

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
		propertiesStruct[property.Name] = propertyRecord
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

	if resource.Id != "" {
		properties["id"] = resource.Id
	}

	return abs.NewRecordLikeWithProperties(properties)
}

func ResourceFromRecord(record abs.RecordLike) *model.Resource {
	if record == nil {
		return nil
	}

	var namespace = record.GetPropertyWithDefault("namespace", util.Kv("name", "")).(map[string]interface{})["name"].(string)

	var resource = &model.Resource{
		Id:              record.GetPropertyWithDefault("id", "").(string),
		Name:            record.GetPropertyWithDefault("name", "").(string),
		Namespace:       namespace,
		Virtual:         record.GetPropertyWithDefault("virtual", false).(bool),
		Abstract:        record.GetPropertyWithDefault("abstract", false).(bool),
		Immutable:       record.GetPropertyWithDefault("immutable", false).(bool),
		CheckReferences: record.GetPropertyWithDefault("checkReferences", false).(bool),
		SourceConfig: &model.ResourceSourceConfig{
			DataSource: record.GetPropertyWithDefault("dataSource", util.Kv("name", "")).(map[string]interface{})["name"].(string),
			Entity:     record.GetPropertyWithDefault("entity", "").(string),
			Catalog:    record.GetPropertyWithDefault("catalog", "").(string),
		},
		Properties: util.ArrayMap(util.MapToArray(record.GetPropertyWithDefault("properties", emptyMap()).(map[string]interface{})), func(t util.MapEntry[interface{}]) *model.ResourceProperty {
			return ResourcePropertyFromRecord(t.Key, abs.NewRecordLikeWithProperties(t.Val.(map[string]interface{})))
		}),
		Annotations: convertMap(record.GetPropertyWithDefault("annotations", emptyMap()).(map[string]interface{}), func(v interface{}) string {
			return v.(string)
		}),
	}

	if record.HasProperty("indexes") {
		list := record.GetPropertyWithDefault("indexes", emptyMap()).([]interface{})

		resource.Indexes = make([]*model.ResourceIndex, 0)

		for _, val := range list {
			resource.Indexes = append(resource.Indexes, ResourceIndexFromValue(val.(map[string]interface{})))
		}
	}

	if record.HasProperty("types") {
		list := record.GetProperty("types").([]interface{})

		resource.Types = make([]*model.ResourceSubType, 0)

		for _, val := range list {
			resource.Types = append(resource.Types, ResourceTypeFromValue(val))
		}
	}

	if record.HasProperty("title") {
		resource.Title = new(string)
		*resource.Title = record.GetPropertyWithDefault("title", "").(string)
	}

	if record.HasProperty("description") {
		resource.Description = new(string)
		*resource.Description = record.GetPropertyWithDefault("description", "").(string)
	}

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
		list := val["properties"].(map[string]interface{})

		for _, val := range list {
			pVal := val.(map[string]interface{})
			ri.Properties = append(ri.Properties, &model.ResourceIndexProperty{
				Name:  pVal["name"].(string),
				Order: model.Order(model.Order_value["ORDER_"+pVal["order"].(string)]),
			})
		}
	}

	return ri
}

func ResourceIndexToValue(index *model.ResourceIndex) interface{} {
	var propertyStructList []interface{}

	for _, property := range index.Properties {
		propertyStructList = append(propertyStructList, map[string]interface{}{
			"name":  property.Name,
			"order": property.Order.String()[len("ORDER_"):],
		})
	}

	return map[string]interface{}{
		"annotations": convertMap(index.Annotations, func(v string) interface{} {
			return v
		}),
		"unique":     index.Unique,
		"properties": propertyStructList,
	}
}

func ResourceTypeFromValue(val interface{}) *model.ResourceSubType {
	var st = val.(map[string]interface{})
	rt := &model.ResourceSubType{
		Name: st["name"].(string),
		Properties: util.ArrayMap(util.MapToArray(st["properties"].(map[string]interface{})), func(t util.MapEntry[interface{}]) *model.ResourceProperty {
			return ResourcePropertyFromRecord(t.Key, abs.NewRecordLikeWithProperties(t.Val.(map[string]interface{})))
		}),
	}

	if st["title"] != nil {
		rt.Title = st["title"].(string)
	}

	if st["description"] != nil {
		rt.Description = st["description"].(string)
	}

	return rt
}

func ResourceTypeToValue(resource *model.Resource, subType *model.ResourceSubType) interface{} {
	var properties = make(map[string]interface{})
	for _, property := range subType.Properties {
		propertyRecord := ResourcePropertyToRecord(property, resource)
		properties[property.Name] = propertyRecord
	}

	return map[string]interface{}{
		"name":        subType.Name,
		"title":       subType.Title,
		"description": subType.Description,
		"properties":  properties,
	}
}

func emptyMap() map[string]interface{} {
	return make(map[string]interface{})
}
