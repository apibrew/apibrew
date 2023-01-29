package test

import (
	"github.com/tislib/data-handler/pkg/types"
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
		packedVal, _ := propertyType.Pack(val)

		propertyType.String(val)
		propertyType.Pointer(false)
		propertyType.Pointer(true)
		propertyType.IsEmpty(val)

		if propertyType.ValidatePackedValue(packedVal) != nil {
			t.Error("Cannot validate default value: " + propertyType.ValidatePackedValue(packedVal).Error())
		}
	}
}
