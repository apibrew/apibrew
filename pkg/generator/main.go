//go:generate qtc -dir=templates

package generator

import (
	"fmt"
	"github.com/apibrew/apibrew/pkg/generator/templates/golang"
	"github.com/apibrew/apibrew/pkg/generator/typescript"
	"github.com/apibrew/apibrew/pkg/model"
)

func GenerateResourceCodes(platform string, pkg string, resources []*model.Resource, path string, namespace string) error {
	switch platform {
	case "golang":
		return golang.GenerateGoResourceCode(pkg, resources, path, namespace)

	case "typescript":
		return typescript.GenerateResourceCode(typescript.GenerateResourceCodeParams{
			Namespace: namespace,
			Package:   pkg,
			Resources: resources,
			Path:      path,
		})
	default:
		return fmt.Errorf("Unknown platform: " + platform)
	}
}
