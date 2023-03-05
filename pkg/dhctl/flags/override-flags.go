package flags

import (
	"github.com/spf13/cobra"
)

type overrideFlags struct {
	overrideNamespace  *string
	overrideDataSource *string
}

func (s *overrideFlags) Declare(cmd *cobra.Command) {
	s.overrideNamespace = cmd.PersistentFlags().String("override-namespace", "", "Override namespace")
	s.overrideDataSource = cmd.PersistentFlags().String("override-data-source", "", "Override data source")
}

func (s *overrideFlags) Parse(result *OverrideConfig, cmd *cobra.Command, args []string) {
	result.Namespace = *s.overrideNamespace
	result.DataSource = *s.overrideDataSource
}

type OverrideConfig struct {
	Namespace  string
	DataSource string
}

func NewOverrideFlags() FlagHelper[*OverrideConfig] {
	return &overrideFlags{}
}
