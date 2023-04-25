package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/tislib/apibrew/pkg/apbradm"
)

func main() {
	rootCmd := apbradm.PrepareRootCmd()

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
