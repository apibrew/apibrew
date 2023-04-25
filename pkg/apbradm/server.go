package apbradm

import (
	"github.com/spf13/cobra"
	"github.com/tislib/apibrew/pkg/apbradm/server"
)

var serverCmdInit *string = new(string)
var serverLogLevelInit *string = new(string)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "server - run apibrew cluster server",
	RunE: func(cmd *cobra.Command, args []string) error {

		server := server.Server{
			Init:     *serverCmdInit,
			LogLevel: *serverLogLevelInit,
		}

		server.Run()

		return nil
	},
}

func init() {
	serverCmdInit = serverCmd.PersistentFlags().String("init", "", "Initial Data for configuring system")
	serverLogLevelInit = serverCmd.PersistentFlags().String("log-level", "info", "Debug flag")
}
