package output

import (
	"data-handler/model"
	"log"
	"os"
)

type Writer interface {
	WriteResources(resources []*model.Resource)
	WriteRecords(resource *model.Resource, records []*model.Record)
	DescribeResource(resource *model.Resource)
}

func NewOutputWriter(outputFormat string) Writer {
	switch outputFormat {
	case "console":
		return &consoleWriter{
			writer: os.Stdout,
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
