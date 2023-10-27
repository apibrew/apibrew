package apbr

import (
	"errors"
	"github.com/apibrew/apibrew/pkg/apbr/flags"
	"github.com/apibrew/apibrew/pkg/apbr/output"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"io"
	"os"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get - get resource/record docs: https://apibrew.io/docs/cli#apply",
	RunE: func(cmd *cobra.Command, args []string) error {
		parseRootFlags(cmd)

		f := getFlag(cmd, "format", true)
		o := getFlag(cmd, "output", false)
		forApply, err := cmd.Flags().GetBool("for-apply")

		if err != nil {
			return err
		}

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

		writer := output.NewOutputWriter(f, w, map[string]string{
			"for-apply": boolToString(forApply),
		})

		if writer.IsBinary() && o == "" {
			return errors.New("format is binary but output is not specified")
		}

		if selection.Resources != nil {
			err := writer.WriteResource(selection.Resources...)

			if err != nil {
				return err
			}
		}

		for _, recordProvider := range selection.RecordProviders {
			records := recordProvider()
			err := writer.WriteRecordsChan(records.Resource, records.Total, records.Records)

			if err != nil {
				return err
			}
		}

		return nil
	},
}

func boolToString(apply bool) string {
	if apply {
		return "true"
	} else {
		return "false"
	}
}

func initGetCmd() {
	getCmd.PersistentFlags().StringP("format", "f", "console", "format")
	getCmd.PersistentFlags().StringP("output", "o", "", "output")
	getCmd.PersistentFlags().Int64("limit", 100, "limit")
	getCmd.PersistentFlags().Int64("offset", 0, "offset")
	getCmd.PersistentFlags().Bool("for-apply", false, "Prepare for apply")
	selectorFlags.Declare(getCmd)
}
