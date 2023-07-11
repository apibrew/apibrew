package golang

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/gosimple/slug"
	log "github.com/sirupsen/logrus"
	"os"
)

func GenerateGoResourceCode(pkg string, resources []*model.Resource, path string, namespace string) error {
	for _, resource := range resources {
		rawCode := GenerateResourceCode(pkg, resource)

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

	for _, resource := range resources {
		rawCode := GenerateResourceMappingCode(pkg, resource, resources)

		code := []byte(rawCode)

		resourceFileName := slug.Make(resource.Name) + "-mapping.go"

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
