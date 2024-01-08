package util

import (
	"github.com/apibrew/apibrew/pkg/abs"
	"strings"
)

func ParseType(elemType string) abs.ResourceIdentity {
	var identity = abs.ResourceIdentity{}

	if strings.Contains(elemType, "/") {
		identity.Namespace = strings.Split(elemType, "/")[0]
		identity.Name = strings.Split(elemType, "/")[1]
	} else {
		identity.Namespace = "default"
		identity.Name = elemType
	}

	return identity
}
