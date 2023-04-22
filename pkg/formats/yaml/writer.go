package yamlformat

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"github.com/tislib/apibrew/pkg/formats"
	"github.com/tislib/apibrew/pkg/model"
	"io"
)

type writer struct {
	output io.Writer
	batch  *model.Batch
}

func (w *writer) WriteResource(resource ...*model.Resource) error {
	return nil
}

func (w *writer) WriteRecord(namespace string, resourceName string, record ...*model.Record) error {
	log.Infof("Writing records for %s/%s (%d)", namespace, resourceName, len(record))

	if w.batch == nil {
		return errors.New("batch is not started")
	}

	batchPart := &model.BatchRecordsPart{
		Namespace: namespace,
		Resource:  resourceName,
	}

	for _, item := range record {
		batchPart.Values = append(batchPart.Values, item.PropertiesPacked...)
	}

	w.batch.BatchRecords = append(w.batch.BatchRecords, batchPart)

	log.Info("Done")

	return nil
}

func NewWriter(output io.Writer) formats.Writer {
	return &writer{output: output}
}
