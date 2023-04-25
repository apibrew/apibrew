package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/tislib/apibrew/pkg/apbr"
)

func main() {
	rootCmd := apbr.PrepareRootCmd()

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
