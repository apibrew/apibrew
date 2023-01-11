package main

import "github.com/spf13/cobra"

var server = ""

func defineRootFlags() {
	rootCmd.PersistentFlags().String("server", "", "")
}

func parseRootFlags(cmd *cobra.Command) {
	server, _ = cmd.Flags().GetString("server")
}
