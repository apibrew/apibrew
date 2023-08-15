package mapping

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/util"
	log "github.com/sirupsen/logrus"
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

	var propertyStructList []*structpb.Value
	for _, property := range resource.Properties {
		propertyRecord := ResourcePropertyToRecord(property, resource)
		propertyStructList = append(propertyStructList, structpb.NewStructValue(&structpb.Struct{Fields: propertyRecord.Properties}))
	}
	properties["properties"] = structpb.NewListValue(&structpb.ListValue{Values: propertyStructList})

	if resource.Types != nil {
		var lv []*structpb.Value

		for _, subType := range resource.Types {
			lv = append(lv, ResourceTypeToValue(resource, subType))
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
		Id:              record.Id,
		Name:            record.Properties["name"].GetStringValue(),
		Namespace:       record.Properties["namespace"].GetStructValue().GetFields()["name"].GetStringValue(),
		Virtual:         record.Properties["virtual"].GetBoolValue(),
		Abstract:        record.Properties["abstract"].GetBoolValue(),
		Immutable:       record.Properties["immutable"].GetBoolValue(),
		CheckReferences: record.Properties["checkReferences"].GetBoolValue(),
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
		log.Print("I am found")
	}

	return ri
}

func ResourceIndexToValue(index *model.ResourceIndex) *structpb.Value {
	return structpb.NewStructValue(&structpb.Struct{
		Fields: map[string]*structpb.Value{
			// structpb.NewValue(convertMap(resource.Annotations, func(v string) interface{} {
			//		return v
			//	}))
		},
	})
}

func ResourceTypeFromValue(val *structpb.Value) *model.ResourceSubType {
	var st = val.GetStructValue()
	rt := &model.ResourceSubType{
		Name: st.GetFields()["name"].GetStringValue(),
		Properties: util.ArrayMap(st.GetFields()["properties"].GetListValue().Values, func(t *structpb.Value) *model.ResourceProperty {
			return ResourcePropertyFromRecord(&model.Record{Properties: t.GetStructValue().Fields})
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
	var propertyStructList []*structpb.Value
	for _, property := range subType.Properties {
		propertyRecord := ResourcePropertyToRecord(property, resource)
		propertyStructList = append(propertyStructList, structpb.NewStructValue(&structpb.Struct{Fields: propertyRecord.Properties}))
	}

	return structpb.NewStructValue(&structpb.Struct{
		Fields: map[string]*structpb.Value{
			"name":        structpb.NewStringValue(subType.Name),
			"title":       structpb.NewStringValue(subType.Title),
			"description": structpb.NewStringValue(subType.Description),
			"properties":  structpb.NewListValue(&structpb.ListValue{Values: propertyStructList}),
		},
	})
}
