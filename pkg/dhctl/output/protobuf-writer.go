package output

import (
	"context"
	"github.com/tislib/data-handler/pkg/batch"
	"github.com/tislib/data-handler/pkg/model"
	"log"
	"time"
)

type protobufWriter struct {
	batchWriter batch.Writer
	started     bool
}

func (c *protobufWriter) IsBinary() bool {
	return true
}

func (c *protobufWriter) nextBatch() {
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

func (c *protobufWriter) WriteRecords(resource *model.Resource, total uint32, recordsChan chan *model.Record) {
	log.Printf("Total records to be written: %d \n", total)
	if err := c.batchWriter.StartBatch(&model.BatchHeader{
		Mode:        model.BatchMode_BATCH_CREATE,
		Annotations: nil,
	}); err != nil {
		log.Fatal(err)
	}

	var buf []*model.Record

	var i int64 = 0
	var prevI int64 = 0
	startTime := time.Now()

	ctx, cancel := context.WithCancel(context.TODO())

	defer func() {
		cancel()
	}()

	go func() {
		for {
			time.Sleep(time.Second)
			select {
			case <-ctx.Done():
				return
			default:

			}

			var speed = i - prevI
			avgSpeed := i / (time.Now().Unix() - startTime.Unix())
			if avgSpeed == 0 {
				avgSpeed = 1
			}
			var remTime = time.Second * time.Duration((int64(total)-i)/avgSpeed)
			log.Printf("%d/%d records written; %d per second; remeaning %v", total, i, speed, remTime)
			prevI = i
		}
	}()

	for item := range recordsChan {
		buf = append(buf, item)
		i++

		if len(buf) == 10000 {
			c.batchWriter.WriteRecord(resource.Namespace, resource.Name, buf...)
			buf = []*model.Record{}
			c.nextBatch()
		}
	}

	if len(buf) > 0 {
		c.batchWriter.WriteRecord(resource.Namespace, resource.Name, buf...)
	}

	if err := c.batchWriter.EndBatch(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Total records written: %d/%d", total, i)
}
