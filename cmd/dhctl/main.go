package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/tislib/apibrew/pkg/dhctl"
)

func main() {
	rootCmd := dhctl.PrepareRootCmd()

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}