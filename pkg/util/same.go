package util

import (
	"github.com/tislib/data-handler/pkg/model"
)

func IsSameIdentifiedResourceProperty(property1, property2 *model.ResourceProperty) bool {
	if property1.Mapping != property2.Mapping {
		return false
	}

	return true
}

func IsSameResourceProperty(property1, property2 *model.ResourceProperty) bool {
	if property1.Mapping != property2.Mapping {
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

	if property1.Type == model.ResourcePropertyType_TYPE_STRING && property1.Length != property2.Length {
		return false
	}

	if property1.Type == model.ResourcePropertyType_TYPE_REFERENCE {
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

	return true
}
