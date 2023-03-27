package sqlite

import (
	"github.com/rakyll/statik/fs"
	log "github.com/sirupsen/logrus"
	_ "github.com/tislib/data-handler/pkg/backend/sqlite/sql/statik"
	"io"
	"net/http"
)

var statikFS http.FileSystem

func init() {
	var err error
	statikFS, err = fs.NewWithNamespace("sqlite")

	if err != nil {
		log.Fatal(err)
	}
}

func (p sqliteBackendOptions) GetSql(name string) string {
	entityExistsFile, err := statikFS.Open("/" + name + ".sql")

	if err != nil {
		log.Fatal(err)
	}

	data, err := io.ReadAll(entityExistsFile)

	if err != nil {
		log.Fatal(err)
	}

	return string(data)
}
