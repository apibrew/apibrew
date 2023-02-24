package output

import (
	"github.com/tislib/data-handler/pkg/model"
	"log"
	"os"
)

type Writer interface {
	WriteResources(resources []*model.Resource)
	WriteRecords(resource *model.Resource, recordsChan chan *model.Record)
}

func NewOutputWriter(outputFormat string) Writer {
	switch outputFormat {
	case "console":
		return &consoleWriter{
			writer:   os.Stdout,
			describe: false,
		}
	case "describe":
		return &consoleWriter{
			writer:   os.Stdout,
			describe: true,
		}
	case "yaml":
		return &yamlWriter{
			writer: os.Stdout,
		}
	case "yml":
		return &yamlWriter{
			writer: os.Stdout,
		}
	}

	log.Fatal("Writer not found: " + outputFormat)

	return nil
}
