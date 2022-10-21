package test

import (
	"data-handler/service/types"
	"testing"
)

func TestGetPropertyType(t *testing.T) {
	allTypes := types.GetAllResourcePropertyTypes()

	for _, resourcePropertyType := range allTypes {
		types.ByResourcePropertyType(resourcePropertyType)
	}

	for _, resourcePropertyType := range allTypes {
		propertyType := types.ByResourcePropertyType(resourcePropertyType)

		val := propertyType.Default()

		propertyType.String(val)
		propertyType.Pointer(false)
		propertyType.Pointer(true)
		propertyType.IsEmpty(val)

		if propertyType.ValidateValue(val) != nil {
			t.Error("Cannot validate default value: " + propertyType.ValidateValue(val).Error())
		}
	}
}
