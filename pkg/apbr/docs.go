package apbr

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra/doc"
)

func GenerateDocs() {
	err := doc.GenMarkdownTree(rootCmd, "docs/content/apbr")
	if err != nil {
		log.Fatal(err)
	}
}
