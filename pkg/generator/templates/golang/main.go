package golang

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/gosimple/slug"
	log "github.com/sirupsen/logrus"
	"go/format"
	"os"
)

func GenerateGoResourceCode(pkg string, resources []*model.Resource, path string, namespace string) error {
	for _, resource := range resources {
		if err := generateResourceCode(pkg, resource, path); err != nil {
			return err
		}

		//if err := generaDefCode(pkg, resources, path, resource); err != nil {
		//	return err
		//}

		if err := generaMappingCode(pkg, resources, path, resource); err != nil {
			return err
		}
	}

	return nil
}

func generaMappingCode(pkg string, resources []*model.Resource, path string, resource *model.Resource) error {
	rawCode := GenerateResourceMappingCode(pkg, resource, resources)

	code, err := format.Source([]byte(rawCode))
	if err != nil {
		log.Print(code)

		return err
	}

	resourceFileName := slug.Make(resource.Name) + "-mapping.go"

	if err := os.MkdirAll(path, 0777); err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(path+"/"+resourceFileName, code, 0777)

	if err != nil {
		return err
	}
	return nil
}

func generaDefCode(pkg string, resources []*model.Resource, path string, resource *model.Resource) error {
	rawCode := GenerateResourceDefCode(pkg, resource)

	code, err := format.Source([]byte(rawCode))
	if err != nil {
		log.Print(code)

		return err
	}

	resourceFileName := slug.Make(resource.Name) + "-def.go"

	if err := os.MkdirAll(path, 0777); err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(path+"/"+resourceFileName, code, 0777)

	if err != nil {
		return err
	}
	return nil
}

func generateResourceCode(pkg string, resource *model.Resource, path string) error {
	rawCode := GenerateResourceCode(pkg, resource)

	code, err := format.Source([]byte(rawCode))
	if err != nil {
		log.Print(code)

		return err
	}

	resourceFileName := slug.Make(resource.Name) + ".go"

	if err := os.MkdirAll(path, 0777); err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(path+"/"+resourceFileName, code, 0777)

	if err != nil {
		return err
	}
	return nil
}
