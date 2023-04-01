package yamlformat

import (
	"encoding/binary"
	"errors"
	log "github.com/sirupsen/logrus"
	"github.com/tislib/data-handler/pkg/formats"
	"github.com/tislib/data-handler/pkg/model"
	"google.golang.org/protobuf/proto"
	"io"
)

type writer struct {
	output io.Writer
	batch  *model.Batch
}

func (w *writer) StartBatch(header *model.BatchHeader) error {
	if w.batch != nil {
		return errors.New("batch is already started")
	}

	w.batch = &model.Batch{
		Header: header,
	}

	return nil
}

func (w *writer) EndBatch() error {
	if w.batch == nil {
		return errors.New("batch is not started")
	}

	if len(w.batch.BatchRecords) == 0 && len(w.batch.Resources) == 0 {
		w.batch = nil
		return nil
	}

	data, err := proto.Marshal(w.batch)

	if err != nil {
		return err
	}

	err = binary.Write(w.output, binary.BigEndian, uint32(len(data)))

	if err != nil {
		return err
	}

	_, err = w.output.Write(data)

	if err != nil {
		return err
	}

	w.batch = nil

	return nil
}

func (w *writer) WriteResource(resource ...*model.Resource) error {
	if w.batch == nil {
		return errors.New("batch is not started")
	}

	w.batch.Resources = append(w.batch.Resources, resource...)

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
