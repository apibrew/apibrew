package util

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resources/special"
	"github.com/apibrew/apibrew/pkg/service/annotations"
	"google.golang.org/protobuf/proto"
)

type Named interface {
	GetName() string
}

func GetNamedMap[T Named](items []T) map[string]T {
	var result = make(map[string]T)

	for _, prop := range items {
		result[prop.GetName()] = prop
	}

	return result
}

func GetArrayIndex[T comparable](items []T, item T, comparator func(a, b T) bool) int {
	for i, elem := range items {
		if comparator(elem, item) {
			return i
		}
	}

	return -1
}

func LocatePropertyByName(resource *model.Resource, propertyName string) *model.ResourceProperty {
	for _, property := range resource.Properties {
		if property.Name == propertyName {
			return property
		}
	}

	return nil
}

func HistoryResource(resource *model.Resource) *model.Resource {
	if resource == nil || resource.SourceConfig == nil {
		return resource
	}

	var historyResource = proto.Clone(resource).(*model.Resource)
	historyResource.SourceConfig.Entity += "_h"
	historyResource.Indexes = nil

	annotations.Enable(historyResource, annotations.HistoryResource)

	for _, prop := range historyResource.Properties {
		if prop.Name == "version" {
			prop.Primary = true
		}
	}

	return historyResource
}

func HistoryPlan(plan *model.ResourceMigrationPlan) *model.ResourceMigrationPlan {
	return &model.ResourceMigrationPlan{
		ExistingResource: HistoryResource(plan.ExistingResource),
		CurrentResource:  HistoryResource(plan.CurrentResource),
		Steps:            plan.Steps,
	}
}

func RemarkResource(resource *model.Resource) {
	propertyNameMap := GetNamedMap(resource.Properties)

	if !annotations.IsEnabled(resource, annotations.EnableAudit) {
		auditPropertyCount := 0
		for _, prop := range special.AuditProperties {
			if propertyNameMap[prop.Name] != nil && propertyNameMap[prop.Name].Type == prop.Type {
				auditPropertyCount++
			}
		}

		if auditPropertyCount == 4 {
			annotations.Enable(resource, annotations.EnableAudit)
		}
	}

	if !annotations.IsEnabled(resource, annotations.DisableVersion) {
		if propertyNameMap[special.VersionProperty.Name] == nil || propertyNameMap[special.VersionProperty.Name].Type != special.VersionProperty.Type {
			annotations.Enable(resource, annotations.DisableVersion)
		}
	}
}

func NormalizeResource(resource *model.Resource) {
	propertyNameMap := GetNamedMap(resource.Properties)

	if resource.Annotations == nil {
		resource.Annotations = make(map[string]string)
	}

	if annotations.IsEnabled(resource, annotations.EnableAudit) {
		exists := false
		for _, prop := range special.AuditProperties {
			if propertyNameMap[prop.Name] != nil {
				exists = true
				break
			}
		}
		if !exists {
			resource.Properties = append(resource.Properties, special.AuditProperties...)
		}
	}

	if !annotations.IsEnabled(resource, annotations.DisableVersion) && propertyNameMap[special.VersionProperty.Name] == nil {
		resource.Properties = append(resource.Properties, special.VersionProperty)
	}

	if !HasResourcePrimaryProp(resource) && propertyNameMap[special.IdProperty.Name] == nil {
		resource.Properties = append([]*model.ResourceProperty{special.IdProperty}, resource.Properties...)
	}

	annotations.Enable(resource, annotations.NormalizedResource)
}

func HasResourceSinglePrimaryProp(resource *model.Resource) bool {
	primaryPropCount := 0

	for _, item := range resource.Properties {
		if item.Primary {
			primaryPropCount++
		}
	}

	return primaryPropCount == 1
}

func HasResourcePrimaryProp(resource *model.Resource) bool {
	for _, item := range resource.Properties {
		if item.Primary {
			return true
		}
	}

	return false
}

func GetResourceSinglePrimaryProp(resource *model.Resource) *model.ResourceProperty {
	for _, item := range resource.Properties {
		if item.Primary {
			return item
		}
	}

	return nil
}

func ResourceWalkProperties(resource *model.Resource, callback func(path string, property *model.ResourceProperty)) {
	ResourceWalkPropertiesRecursive(resource, "", resource.Properties, callback)
}

func ResourceWalkPropertiesRecursive(resource *model.Resource, path string, properties []*model.ResourceProperty, callback func(path string, property *model.ResourceProperty)) {
	for _, property := range properties {
		callback(path, property)

		if property.Type == model.ResourceProperty_LIST || property.Type == model.ResourceProperty_MAP {
			ResourceWalkPropertiesRecursive(resource, path+"."+property.Name+"[]", []*model.ResourceProperty{property.Item}, func(path string, property *model.ResourceProperty) {
				callback(path, property)
			})
		} else if property.Type == model.ResourceProperty_STRUCT {
			ResourceWalkPropertiesRecursive(resource, path+"."+property.Name, property.Properties, func(path string, property *model.ResourceProperty) {
				callback(path, property)
			})

			if property.TypeRef != nil {
				var subType *model.ResourceSubType

				for _, subTypeItem := range resource.Types {
					if subTypeItem.Name == *property.TypeRef {
						subType = subTypeItem
						break
					}
				}

				if subType == nil {
					panic("sub type not found: " + *property.TypeRef)
				}

				ResourceWalkPropertiesRecursive(resource, path+"."+property.Name, subType.Properties, func(path string, property *model.ResourceProperty) {
					callback(path, property)
				})
			}
		}
	}
}
