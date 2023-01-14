package main

import (
	"data-handler/model"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func checkError(err *model.Error) {
	if err != nil {
		log.Fatal(err.Message)
	}
}

func getFlag(cmd *cobra.Command, commandName string, required bool) string {
	o, err := cmd.Flags().GetString(commandName)
	check(err)

	if o == "" && required {
		log.Fatalf("%s is required but not provided", commandName)
	}

	return o
}
