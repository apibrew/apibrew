package java

import (
	"github.com/apibrew/apibrew/pkg/model"
	log "github.com/sirupsen/logrus"
	"os"
)

func GenerateResourceCode(pkg string, resources []*model.Resource, path string, namespace string) error {
	for _, resource := range resources {
		if err := generateResourceCode(pkg, resource, path); err != nil {
			return err
		}
	}

	return nil
}

func generateResourceCode(pkg string, resource *model.Resource, path string) error {
	code := GenerateClassCode(pkg, resource)

	resourceFileName := javaClassName(resource.Name) + ".java"

	if err := os.MkdirAll(path, 0777); err != nil {
		log.Fatal(err)
	}

	err := os.WriteFile(path+"/"+resourceFileName, []byte(code), 0777)

	if err != nil {
		return err
	}
	return nil
}
