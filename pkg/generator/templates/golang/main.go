package golang

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/util"
	log "github.com/sirupsen/logrus"
	"go/format"
	"os"
)

func GenerateGoResourceCode(pkg string, resources []*model.Resource, actions map[*model.Resource][]*model.Resource, path string, namespace string) error {
	for _, resource := range resources {
		if err := generateResourceCode(pkg, resource, path); err != nil {
			return err
		}

		if err := generaDefCode(pkg, resource, path); err != nil {
			return err
		}

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
		code = []byte(rawCode)

		log.Warn(err)
	}

	resourceFileName := util.PathSlug(resource.Name) + "-mapping.go"

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
		code = []byte(rawCode)

		log.Warn(err)
	}

	resourceFileName := util.PathSlug(resource.Name) + ".go"

	if err := os.MkdirAll(path, 0777); err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(path+"/"+resourceFileName, code, 0777)

	if err != nil {
		return err
	}
	return nil
}

func generaDefCode(pkg string, resource *model.Resource, path string) error {
	rawCode := GenerateResourceDefCode(pkg, resource)

	code, err := format.Source([]byte(rawCode))
	if err != nil {
		code = []byte(rawCode)

		log.Warn(err)
	}

	resourceFileName := util.PathSlug(resource.Name) + "-def" + ".go"

	if err := os.MkdirAll(path, 0777); err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(path+"/"+resourceFileName, code, 0777)

	if err != nil {
		return err
	}
	return nil
}
