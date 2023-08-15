package helper

import "github.com/apibrew/apibrew/pkg/model"

func IsPropertyOmitted(property *model.ResourceProperty) bool {
	if property.Type == model.ResourceProperty_LIST && property.Item.Type == model.ResourceProperty_REFERENCE && property.Item.BackReference != nil {
		// skip back references as they will be populated on service layer
		return true
	}

	if property.BackReference != nil {
		return true
	}

	return false
}
