package apbr

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func getFlag(cmd *cobra.Command, commandName string, required bool) string {
	o, err := cmd.Flags().GetString(commandName)

	if err != nil {
		panic(err)
	}

	if o == "" && required {
		log.Fatalf("%s is required but not provided", commandName)
	}

	return o
}
