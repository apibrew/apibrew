package util

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resources/special"
	"github.com/apibrew/apibrew/pkg/service/annotations"
	"github.com/gosimple/slug"
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
			prop.Annotations = annotations.EnableWith(prop.Annotations, annotations.PrimaryProperty)
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
		if propertyNameMap[special.AuditProperty.Name] != nil && propertyNameMap[special.AuditProperty.Name].Type == special.AuditProperty.Type {
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

	if !annotations.IsEnabled(resource, annotations.DisableVersion) && propertyNameMap[special.VersionProperty.Name] == nil {
		resource.Properties = append(resource.Properties, special.VersionProperty)
	}

	if annotations.IsEnabled(resource, annotations.EnableAudit) {
		if propertyNameMap[special.AuditProperty.Name] == nil {
			resource.Properties = append(resource.Properties, special.AuditProperty)
		}

		var found = false
		for _, subType := range resource.Types {
			if subType.Name == special.AuditDataSubType.Name {
				found = true
			}
		}

		if !found {
			resource.Types = append(resource.Types, special.AuditDataSubType)
		}
	}

	if !HasResourcePrimaryProp(resource) && propertyNameMap[special.IdProperty.Name] == nil {
		resource.Properties = append([]*model.ResourceProperty{special.IdProperty}, resource.Properties...)
	}

	annotations.Enable(resource, annotations.NormalizedResource)
}

func HasResourceSinglePrimaryProp(resource *model.Resource) bool {
	primaryPropCount := 0

	for _, item := range resource.Properties {
		if annotations.IsEnabled(item, annotations.PrimaryProperty) {
			primaryPropCount++
		}
	}

	return primaryPropCount == 1
}

func HasResourcePrimaryProp(resource *model.Resource) bool {
	for _, item := range resource.Properties {
		if annotations.IsEnabled(item, annotations.PrimaryProperty) {
			return true
		}
	}

	return false
}

func GetResourceSinglePrimaryProp(resource *model.Resource) *model.ResourceProperty {
	for _, item := range resource.Properties {
		if annotations.IsEnabled(item, annotations.PrimaryProperty) {
			return item
		}
	}

	return nil
}

func ResourceWalkProperties(resource *model.Resource, callback func(path string, property *model.ResourceProperty)) {
	resourceWalkPropertiesRecursive(resource, "$", resource.Properties, false, callback)

	for _, subType := range resource.Types {
		resourceWalkPropertiesRecursive(resource, "$."+subType.Name, subType.Properties, false, callback)
	}
}

func resourceWalkPropertiesRecursive(resource *model.Resource, path string, properties []*model.ResourceProperty, isCollectionItem bool, callback func(path string, property *model.ResourceProperty)) {
	for _, property := range properties {
		var newName = path
		if !isCollectionItem {
			newName += "." + property.Name
		}
		callback(newName, property)

		if property.Type == model.ResourceProperty_LIST || property.Type == model.ResourceProperty_MAP {
			resourceWalkPropertiesRecursive(resource, newName+"[]", []*model.ResourceProperty{property.Item}, true, callback)
		}
	}
}

func ResourceRestPath(resource *model.Resource) string {
	if annotations.Get(resource, annotations.OpenApiRestPath) != "" {
		return annotations.Get(resource, annotations.OpenApiRestPath)
	} else if resource.Namespace == "default" {
		return slug.Make(resource.Name)
	} else {
		return slug.Make(resource.Namespace + "/" + resource.Name)
	}
}
