package util

import "github.com/apibrew/apibrew/pkg/util"

func resourcePath(namespaceName string, resourceName string) string {
	if namespaceName == "default" {
		return util.PathSlug(resourceName)
	} else {
		return util.PathSlug(namespaceName) + "-" + util.PathSlug(resourceName)
	}
}
