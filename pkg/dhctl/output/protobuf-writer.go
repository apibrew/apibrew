package output

import (
	"github.com/tislib/data-handler/pkg/batch"
	"github.com/tislib/data-handler/pkg/model"
	"log"
)

type protobufWriter struct {
	batchWriter batch.Writer
	started     bool
}

func (c *protobufWriter) IsBinary() bool {
	return true
}

func (c *protobufWriter) nextBatch() {
	if c.started {
		if err := c.batchWriter.EndBatch(); err != nil {
			log.Fatal(err)
		}
	}

	if err := c.batchWriter.StartBatch(&model.BatchHeader{
		Mode:        model.BatchMode_BATCH_CREATE,
		Annotations: nil,
	}); err != nil {
		log.Fatal(err)
	}

	c.started = true
}

func (c *protobufWriter) WriteResources(resources []*model.Resource) {
	c.nextBatch()

	c.batchWriter.WriteResource(resources...)
}

func (c *protobufWriter) WriteRecords(resource *model.Resource, recordsChan chan *model.Record) {
	c.nextBatch()

	var buf []*model.Record

	for item := range recordsChan {
		buf = append(buf, item)

		if len(buf) == 1000 {
			c.batchWriter.WriteRecord(resource.Namespace, resource.Name, buf...)
			buf = []*model.Record{}
		}
	}

	if len(buf) > 0 {
		c.batchWriter.WriteRecord(resource.Namespace, resource.Name, buf...)
	}
}
