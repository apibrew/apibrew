package main

import (
	"github.com/apibrew/apibrew/pkg/apbr"
	log "github.com/sirupsen/logrus"
)

func main() {
	rootCmd := apbr.PrepareRootCmd()

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
