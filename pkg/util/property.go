package util

import "github.com/apibrew/apibrew/pkg/model"

func IsFilterableProperty(propertyType model.ResourceProperty_Type) bool {
	if propertyType == model.ResourceProperty_LIST || propertyType == model.ResourceProperty_MAP {
		return false
	}

	if propertyType == model.ResourceProperty_OBJECT || propertyType == model.ResourceProperty_STRUCT {
		return false
	}

	if propertyType == model.ResourceProperty_BYTES {
		return false
	}

	return true
}
