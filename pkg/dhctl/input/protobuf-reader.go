package output

import (
	"encoding/binary"
	"github.com/tislib/data-handler/pkg/model"
	"google.golang.org/protobuf/proto"
	"io"
)

type protobufReader struct {
	writer            io.Writer
	hasMessageWritten bool
}

func (c *protobufReader) IsBinary() bool {
	return true
}

func (c *protobufReader) WriteResources(resources []*model.Resource) {
	for _, resource := range resources {
		c.writeMessage(resource)
	}
}

func (c *protobufReader) writeMessage(message proto.Message) {
	body, err := proto.Marshal(message)

	check(err)

	var messageType = string(message.ProtoReflect().Descriptor().Name())

	err = binary.Write(c.writer, binary.BigEndian, int32(len(body)))

	check(err)

	_, err = c.writer.Write([]byte(messageType))

	check(err)

	_, err = c.writer.Write(body)

	check(err)
}

func (c *protobufReader) WriteRecords(resource *model.Resource, total uint32, recordsChan chan *model.Record) {
	for record := range recordsChan {
		c.writeMessage(record)
	}
}
