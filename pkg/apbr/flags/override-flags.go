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

func (s *overrideFlags) Parse(result *OverrideConfig, cmd *cobra.Command, args []string) error {
	var err error
	result.Namespace, err = cmd.PersistentFlags().GetString("override-namespace")

	if err != nil {
		return err
	}

	result.DataSource, err = cmd.PersistentFlags().GetString("override-data-source")

	if err != nil {
		return err
	}

	return nil
}

type OverrideConfig struct {
	Namespace  string
	DataSource string
}

func NewOverrideFlags() FlagHelper[*OverrideConfig] {
	return &overrideFlags{}
}
