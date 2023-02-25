package dhctl

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tislib/data-handler/pkg/dhctl/output"
	"io"
	"os"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get - get type",
	Run: func(cmd *cobra.Command, args []string) {
		parseRootFlags(cmd)
		initClient(cmd.Context())

		f := getFlag(cmd, "format", true)
		o := getFlag(cmd, "output", false)

		selection := selectData(cmd, args)

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

		if selection.resources != nil {
			writer.WriteResources(selection.resources)
		}

		for _, records := range selection.records {
			writer.WriteRecords(records.resource, records.records)
		}

	},
}

func init() {
	getCmd.PersistentFlags().StringP("format", "f", "console", "format")
	getCmd.PersistentFlags().StringP("output", "o", "", "output")
	initSelectorFlags(getCmd)
}
