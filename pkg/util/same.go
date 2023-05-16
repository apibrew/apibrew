package util

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/service/annotations"
)

func IsSameIdentifiedResourceIndex(index1, index2 *model.ResourceIndex) bool {
	return IsSameResourceIndex(index1, index2)
}

func IsSameResourceIndex(index1, index2 *model.ResourceIndex) bool {
	if len(index1.Properties) != len(index2.Properties) {
		return false
	}

	if index1.Unique != index2.Unique {
		return false
	}

	for _, prop1 := range index1.Properties {
		var found = false
		for _, prop2 := range index2.Properties {
			if prop1.Name == prop2.Name {
				found = true
			}
		}

		if !found {
			return false
		}
	}

	//todo check property order and index type

	return true
}

func IsSameResourceProperty(property1, property2 *model.ResourceProperty) bool {
	if property1.Mapping != property2.Mapping {
		return false
	}

	if property1.Name != property2.Name {
		return false
	}

	if property1.Type != property2.Type {
		return false
	}

	if property1.Required != property2.Required {
		return false
	}

	if property1.Unique != property2.Unique {
		return false
	}

	if property1.Type == model.ResourceProperty_STRING && property1.Length != property2.Length {
		return false
	}

	if property1.Type == model.ResourceProperty_REFERENCE {
		if (property1.Reference == nil) != (property2.Reference == nil) {
			return false
		}

		if property1.Reference.ReferencedResource != property2.Reference.ReferencedResource {
			return false
		}

		if property1.Reference.Cascade != property2.Reference.Cascade {
			return false
		}
	}

	if !annotations.IsSame(property1.Annotations, property2.Annotations) {
		return false
	}

	return true
}

func IsSameIdentifiedResourceProperty(property1, property2 *model.ResourceProperty) bool {
	if property1.Mapping == property2.Mapping {
		return true
	}

	if property1.Name == property2.Name {
		return true
	}

	property1MatchKey := annotations.Get(property1, annotations.SourceMatchKey)
	property2MatchKey := annotations.Get(property2, annotations.SourceMatchKey)

	if property1MatchKey != "" && property1MatchKey == property2MatchKey {
		return true
	}

	return false
}
