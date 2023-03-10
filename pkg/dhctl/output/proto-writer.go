package output

import (
	"github.com/tislib/data-handler/pkg/batch"
	"github.com/tislib/data-handler/pkg/model"
	"log"
)

type protoWriter struct {
	batchWriter batch.Writer
}

func (c *protoWriter) IsBinary() bool {
	return true
}

func (c *protoWriter) nextBatch() {
	if err := c.batchWriter.EndBatch(); err != nil {
		log.Fatal(err)
	}

	if err := c.batchWriter.StartBatch(&model.BatchHeader{
		Mode:        model.BatchMode_BATCH_CREATE,
		Annotations: nil,
	}); err != nil {
		log.Fatal(err)
	}
}

func (c *protoWriter) WriteResources(resources []*model.Resource) {
	if err := c.batchWriter.StartBatch(&model.BatchHeader{
		Mode:        model.BatchMode_BATCH_CREATE,
		Annotations: nil,
	}); err != nil {
		log.Fatal(err)
	}

	if err := c.batchWriter.WriteResource(resources...); err != nil {
		log.Fatal(err)
	}

	if err := c.batchWriter.EndBatch(); err != nil {
		log.Fatal(err)
	}
}

func (c *protoWriter) WriteRecords(resource *model.Resource, total uint32, recordsChan chan *model.Record) {
	log.Fatal("Writing records is not supported to proto file")
}
