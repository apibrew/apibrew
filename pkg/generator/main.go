//go:generate go get -u github.com/valyala/quicktemplate/qtc
//go:generate qtc -dir=templates

package generator

import (
	"fmt"
	"github.com/apibrew/apibrew/pkg/generator/golang"
	"github.com/apibrew/apibrew/pkg/generator/typescript"
	"github.com/apibrew/apibrew/pkg/model"
)

func GenerateResourceCodes(platform string, pkg string, resources []*model.Resource, path string, namespace string) error {
	switch platform {
	case "golang":
		return golang.GenerateGoResourceCode(golang.GenerateResourceCodeParams{
			Namespace: namespace,
			Package:   pkg,
			Resources: resources,
			Path:      path,
		})

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
