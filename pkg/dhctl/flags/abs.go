package flags

import "github.com/spf13/cobra"

type FlagHelper[T any] interface {
	Declare(cmd *cobra.Command)
	Parse(elem T, cmd *cobra.Command, args []string)
}
