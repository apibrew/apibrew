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
	if err := c.batchWriter.StartBatch(&model.BatchHeader{
		Mode:        model.BatchMode_BATCH_CREATE,
		Annotations: nil,
	}); err != nil {
		log.Fatal(err)
	}

	c.batchWriter.WriteResource(resources...)
	if err := c.batchWriter.EndBatch(); err != nil {
		log.Fatal(err)
	}
}

func (c *protobufWriter) WriteRecords(resource *model.Resource, recordsChan chan *model.Record) {
	var buf []*model.Record

	var i = 0
	for item := range recordsChan {
		buf = append(buf, item)
		i++

		if len(buf) == 1000 {
			c.nextBatch()
			c.batchWriter.WriteRecord(resource.Namespace, resource.Name, buf...)
			buf = []*model.Record{}
		}
		if i%10000 == 0 {
			log.Printf("%d records written", i)
		}
	}

	if len(buf) > 0 {
		c.batchWriter.WriteRecord(resource.Namespace, resource.Name, buf...)
	}

	if err := c.batchWriter.EndBatch(); err != nil {
		log.Fatal(err)
	}
}
