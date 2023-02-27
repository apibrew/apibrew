package dhctl

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create - Create resource from existing table",
	Run: func(cmd *cobra.Command, args []string) {

	},
	PreRun: func(cmd *cobra.Command, args []string) {
		cmd.AddCommand(createResourceCmd)
		cmd.AddCommand(createRecordCmd)
		log.Print("Pre Run")
	},
}

var createResourceCmd = &cobra.Command{
	Use:   "resource",
	Short: "Create resource from existing table",
	Run: func(cmd *cobra.Command, args []string) {

		defineRootFlags(cmd)
		cmd.Flags().String("name", "", "")

		err := cmd.Flags().Parse(args)
		check(err)

		parseRootFlags(cmd)

		initClient(cmd.Context())

		log.Print(cmd.Flags().Args())
	},
}

var createRecordCmd = &cobra.Command{
	Use:   "record",
	Short: "Create record",
	Run: func(cmd *cobra.Command, args []string) {

		defineRootFlags(cmd)
		cmd.Flags().String("name", "", "")

		err := cmd.Flags().Parse(args)
		check(err)

		parseRootFlags(cmd)

		initClient(cmd.Context())

		log.Print(cmd.Flags().Args())
	},
}

func init() {

}
