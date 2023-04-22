package dhctl

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tislib/apibrew/pkg/dhctl/flags"
	"github.com/tislib/apibrew/pkg/dhctl/output"
	"io"
	"os"
	"strconv"
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

		for _, recordProvider := range selection.RecordProviders {
			log.Print("Before begin")
			records := recordProvider()
			log.Println("Begin " + records.Resource.Name + " " + strconv.Itoa(int(records.Total)))
			writer.WriteRecords(records.Resource, records.Total, records.Records)
			log.Println("End2 " + records.Resource.Name)
		}

		log.Println("DONE ALL")

	},
}

func initGetCmd() {
	getCmd.PersistentFlags().StringP("format", "f", "console", "format")
	getCmd.PersistentFlags().StringP("output", "o", "", "output")
	getCmd.PersistentFlags().Int64("limit", 100, "limit")
	getCmd.PersistentFlags().Int64("offset", 0, "offset")
	getCmd.PersistentFlags().Bool("backup", false, "backup")
	selectorFlags.Declare(getCmd)
}
