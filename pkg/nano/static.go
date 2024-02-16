package nano

import (
	"github.com/apibrew/apibrew/pkg/nano/statik"
	_ "github.com/apibrew/apibrew/pkg/server/rest/docs/statik"
	"github.com/rakyll/statik/fs"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
)

var statikFS http.FileSystem

func init() {
	var err error
	statikFS, err = fs.NewWithNamespace(statik.NanoBuiltin)

	if err != nil {
		log.Fatal(err)
	}
}

func GetBuiltinJs() string {
	entityExistsFile, err := statikFS.Open("/builtin.js")

	if err != nil {
		log.Fatal(err)
	}

	data, err := io.ReadAll(entityExistsFile)

	if err != nil {
		log.Fatal(err)
	}

	return string(data)
}
