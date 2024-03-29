package apbr

import (
	"github.com/apibrew/apibrew/pkg/client"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var server = ""

func defineRootFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().String("server", "", "")
	cmd.PersistentFlags().Bool("debug", false, "Enable Debug Logging")
	cmd.PersistentFlags().Bool("verbose", false, "")
}

func parseRootFlags(cmd *cobra.Command) {
	server, _ = cmd.Flags().GetString("server")
	verbose, _ := cmd.Flags().GetBool("verbose")
	debug, _ := cmd.Flags().GetBool("debug")

	if verbose {
		log.SetLevel(log.TraceLevel)
		log.SetReportCaller(true)
	} else if debug {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	var err error
	dhClient, err = client.NewClientWithServerName(server)

	if err != nil {
		log.Fatal(err)
	}

}
