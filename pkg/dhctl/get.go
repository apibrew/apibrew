package dhctl

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tislib/data-handler/pkg/dhctl/flags"
	"github.com/tislib/data-handler/pkg/dhctl/output"
	"io"
	"os"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get - get type",
	Run: func(cmd *cobra.Command, args []string) {
		parseRootFlags(cmd)

		f := getFlag(cmd, "format", true)
		o := getFlag(cmd, "output", false)

		var selection = &flags.SelectedRecordsResult{}

		selectorFlags.Parse(selection, cmd, args)

		var w io.Writer
		if o == "" || o == "-" {
			w = os.Stdout
		} else {
			var err error
			var wf *os.File
			wf, err = os.Create(o)

			if err != nil {
				log.Fatal(err)
			}

			w = wf

			defer func() {
				err = wf.Close()

				if err != nil {
					log.Fatal(err)
				}
			}()
		}

		writer := output.NewOutputWriter(f, w)

		if writer.IsBinary() && o == "" {
			log.Fatal("format is binary but output is not specified")
		}

		if selection.Resources != nil {
			writer.WriteResources(selection.Resources)
		}

		for _, records := range selection.Records {
			writer.WriteRecords(records.Resource, records.Records)
		}

	},
}

func initGetCmd() {
	getCmd.PersistentFlags().StringP("format", "f", "console", "format")
	getCmd.PersistentFlags().StringP("output", "o", "", "output")
	getCmd.PersistentFlags().Int64("limit", 100, "limit")
	getCmd.PersistentFlags().Int64("offset", 0, "offset")
	selectorFlags.Declare(getCmd)
}
