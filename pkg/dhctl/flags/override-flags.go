package flags

import (
	"github.com/spf13/cobra"
)

type overrideFlags struct {
}

func (s *overrideFlags) Declare(cmd *cobra.Command) {
	cmd.PersistentFlags().String("override-namespace", "", "Override namespace")
	cmd.PersistentFlags().String("override-data-source", "", "Override data source")
}

func (s *overrideFlags) Parse(result *OverrideConfig, cmd *cobra.Command, args []string) {
	result.Namespace, _ = cmd.PersistentFlags().GetString("override-namespace")
	result.DataSource, _ = cmd.PersistentFlags().GetString("override-data-source")
}

type OverrideConfig struct {
	Namespace  string
	DataSource string
}

func NewOverrideFlags() FlagHelper[*OverrideConfig] {
	return &overrideFlags{}
}
