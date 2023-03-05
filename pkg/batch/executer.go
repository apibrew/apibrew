package batch

import (
	"context"
	"encoding/binary"
	log "github.com/sirupsen/logrus"
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/stub"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/structpb"
	"io"
	"os"
)

type executor struct {
	params ExecutorParams
}

func (e executor) Restore(ctx context.Context, in *os.File) error {
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

		err = e.processBatch(ctx, batch)

		if err != nil {
			return err
		}
	}
}

func (e executor) processBatch(ctx context.Context, batch *model.Batch) error {
	if batch.Header.Mode == model.BatchMode_BATCH_CREATE {
		resp, err := e.params.ResourceServiceClient.Create(ctx, &stub.CreateResourceRequest{
			Token:          e.params.Token,
			Resources:      batch.Resources,
			DoMigration:    e.params.DoMigration,
			ForceMigration: e.params.ForceMigration,
			Annotations:    batch.Header.Annotations,
		})

		if err != nil {
			log.Error(err)
		} else {
			for _, r := range resp.Resources {
				log.Tracef("Resource created: %s/%s(%s)", r.Namespace, r.Name, r.Id)
			}
		}

		for _, res := range batch.BatchRecords {
			var records []*model.Record

			// override config
			if e.params.OverrideConfig.Namespace != "" {
				res.Namespace = e.params.OverrideConfig.Namespace
			}

			resourceResp, err := e.params.ResourceServiceClient.GetByName(ctx, &stub.GetResourceByNameRequest{
				Token:       e.params.Token,
				Namespace:   res.Namespace,
				Name:        res.Resource,
				Annotations: batch.Header.Annotations,
			})

			if err != nil {
				return err
			}

			resPropertyCount := len(resourceResp.Resource.Properties)
			recordCount := len(res.Values) / resPropertyCount

			for i := 0; i < recordCount; i++ {
				var record = new(model.Record)
				record.Properties = make(map[string]*structpb.Value, resPropertyCount)

				for pi, prop := range resourceResp.Resource.Properties {
					idx := resPropertyCount*i + pi
					record.Properties[prop.Name] = res.Values[idx]
				}

				records = append(records, record)
			}

			resp, err := e.params.RecordServiceClient.Create(ctx, &stub.CreateRecordRequest{
				Token:          e.params.Token,
				Namespace:      res.Namespace,
				Resource:       res.Resource,
				Annotations:    batch.Header.Annotations,
				Records:        records,
				IgnoreIfExists: true,
			})

			if err != nil {
				return err
			}

			for _, r := range resp.Records {
				log.Tracef("Record created: %s/%s(%s)", res.Namespace, res.Resource, r.Id)
			}
		}
	}

	return nil
}

type OverrideConfig struct {
	Namespace  string
	DataSource string
}

type ExecutorParams struct {
	Input                 io.Reader
	ResourceServiceClient stub.ResourceServiceClient
	RecordServiceClient   stub.RecordServiceClient
	OverrideConfig        OverrideConfig
	Token                 string
	DoMigration           bool
	ForceMigration        bool
}

func NewExecutor(params ExecutorParams) Executor {
	return &executor{params: params}
}
