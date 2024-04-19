package openapi

import (
	_ "github.com/apibrew/apibrew/pkg/docs/openapi/statik"
	"github.com/rakyll/statik/fs"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
)

var statikFS http.FileSystem

var openApiBaseContent string

func init() {
	var err error
	statikFS, err = fs.NewWithNamespace("openapi")

	if err != nil {
		log.Fatal(err)
	}

	openApiBaseContent = getStaticFile("openapi-base.json")
}

func getStaticFile(name string) string {
	entityExistsFile, err := statikFS.Open("/" + name)

	if err != nil {
		log.Fatal(err)
	}

	data, err := io.ReadAll(entityExistsFile)

	if err != nil {
		log.Fatal(err)
	}

	return string(data)
}
