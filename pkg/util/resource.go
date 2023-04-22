package util

import (
	"github.com/tislib/apibrew/pkg/model"
	"github.com/tislib/apibrew/pkg/resources"
	"github.com/tislib/apibrew/pkg/service/annotations"
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
		for _, prop := range resources.AuditProperties {
			if propertyNameMap[prop.Name] != nil && propertyNameMap[prop.Name].Type == prop.Type {
				auditPropertyCount++
			}
		}

		if auditPropertyCount == 4 {
			annotations.Enable(resource, annotations.EnableAudit)
		}
	}

	if !annotations.IsEnabled(resource, annotations.DisableVersion) {
		if propertyNameMap[resources.VersionProperty.Name] == nil || propertyNameMap[resources.VersionProperty.Name].Type != resources.VersionProperty.Type {
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
		for _, prop := range resources.AuditProperties {
			if propertyNameMap[prop.Name] != nil {
				exists = true
				break
			}
		}
		if !exists {
			resource.Properties = append(resource.Properties, resources.AuditProperties...)
		}
	}

	if !annotations.IsEnabled(resource, annotations.DisableVersion) && propertyNameMap[resources.VersionProperty.Name] == nil {
		resource.Properties = append(resource.Properties, resources.VersionProperty)
	}

	if !HasResourcePrimaryProp(resource) && propertyNameMap[resources.IdProperty.Name] == nil {
		resource.Properties = append([]*model.ResourceProperty{resources.IdProperty}, resource.Properties...)
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
