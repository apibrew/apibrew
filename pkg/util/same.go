package util

import (
	"github.com/tislib/data-handler/pkg/model"
)

func IsSameResourceProperty(property1, property2 *model.ResourceProperty) bool {
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

	//if property1.SourceConfig != property2.SourceConfig { @todo fixme
	//	return false
	//}

	if property1.Length != property2.Length {
		return false
	}

	return true
}
