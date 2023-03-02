package dhctl

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var server = ""

func defineRootFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().String("server", "", "")
	cmd.PersistentFlags().Bool("verbose", false, "")
}

func parseRootFlags(cmd *cobra.Command) {
	server, _ = cmd.PersistentFlags().GetString("server")
	verbose, _ := cmd.PersistentFlags().GetBool("verbose")

	if verbose {
		log.SetLevel(log.TraceLevel)
		log.SetReportCaller(true)
	}

	log.Print("server:", server)
}
