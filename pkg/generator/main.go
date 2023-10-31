//go:generate qtc -dir=templates

package generator

import (
	"fmt"
	"github.com/apibrew/apibrew/pkg/generator/templates/golang"
	"github.com/apibrew/apibrew/pkg/generator/templates/java"
	"github.com/apibrew/apibrew/pkg/generator/templates/python"
	"github.com/apibrew/apibrew/pkg/generator/templates/typescript"
	"github.com/apibrew/apibrew/pkg/model"
)

func GenerateResourceCodes(platform string, pkg string, resources []*model.Resource, resourceActions map[*model.Resource][]*model.Resource, path string, namespace string) error {
	switch platform {
	case "golang":
		return golang.GenerateGoResourceCode(pkg, resources, resourceActions, path, namespace)
	case "java":
		return java.GenerateResourceCode(pkg, resources, resourceActions, path)
	case "typescript":
		return typescript.GenerateResourceCode(pkg, resources, resourceActions, path)
	case "python":
		return python.GenerateResourceCode(pkg, resources, resourceActions, path)
	default:
		return fmt.Errorf("Unknown platform: " + platform)
	}
}
