package rest

import (
	"github.com/rakyll/statik/fs"
	log "github.com/sirupsen/logrus"
	"net/http"
)

var statikFS http.FileSystem

func init() {
	var err error
	statikFS, err = fs.NewWithNamespace("rest")

	if err != nil {
		log.Fatal(err)
	}
}
