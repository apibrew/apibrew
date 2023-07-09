//go:generate go get -u github.com/valyala/quicktemplate/qtc
//go:generate qtc -dir=templates

package generator

import (
	"fmt"
	"github.com/apibrew/apibrew/pkg/generator/templates/golang"
	"github.com/apibrew/apibrew/pkg/generator/typescript"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/gosimple/slug"
	log "github.com/sirupsen/logrus"
	"os"
)

func GenerateResourceCodes(platform string, pkg string, resources []*model.Resource, path string, namespace string) error {
	switch platform {
	case "golang":
		return GenerateGoResourceCode(pkg, resources, path, namespace)

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

func GenerateGoResourceCode(pkg string, resources []*model.Resource, path string, namespace string) error {
	for _, resource := range resources {
		rawCode := golang.GenerateResourceCode(pkg, resource, resources)

		//code, err := format.Source([]byte(rawCode))
		//if err != nil {
		//	log.Print(code)
		//
		//	return err
		//}

		code := []byte(rawCode)

		resourceFileName := slug.Make(resource.Name) + ".go"

		if err := os.MkdirAll(path, 0777); err != nil {
			log.Fatal(err)
		}

		err := os.WriteFile(path+"/"+resourceFileName, code, 0777)

		if err != nil {
			return err
		}
	}

	return nil
}
