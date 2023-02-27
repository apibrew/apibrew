package dhctl

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var server = ""

func defineRootFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().String("server", "", "")
}

func parseRootFlags(cmd *cobra.Command) {
	server, _ = cmd.Flags().GetString("server")

	log.Print("server:", server)
}
