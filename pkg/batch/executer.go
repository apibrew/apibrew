package batch

import (
	"encoding/binary"
	log "github.com/sirupsen/logrus"
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/stub"
	"google.golang.org/protobuf/proto"
	"io"
	"os"
)

type executor struct {
	input                 io.Reader
	resourceServiceClient stub.ResourceServiceClient
	recordServiceClient   stub.RecordServiceClient
}

func (e executor) Restore(in *os.File) error {
	for {
		var messageLength uint32

		err := binary.Read(in, binary.BigEndian, &messageLength)

		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}

		var messageData = make([]byte, messageLength)

		_, err = in.Read(messageData)

		if err != nil {
			return err
		}

		var batch = &model.Batch{}

		err = proto.Unmarshal(messageData, batch)

		if err != nil {
			return err
		}

		e.processBatch(batch)
	}
}

func (e executor) processBatch(batch *model.Batch) {
	log.Print(batch)
}

type ExecutorParams struct {
	Input                 io.Reader
	ResourceServiceClient stub.ResourceServiceClient
	RecordServiceClient   stub.RecordServiceClient
}

func NewExecutor(params ExecutorParams) Executor {
	return &executor{input: params.Input, resourceServiceClient: params.ResourceServiceClient, recordServiceClient: params.RecordServiceClient}
}
