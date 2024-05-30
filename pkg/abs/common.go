package abs

import (
	"github.com/apibrew/apibrew/pkg/model"
	"google.golang.org/protobuf/types/known/structpb"
)

type RecordLike interface {
	GetProperties() map[string]*structpb.Value
}

type ResourceLike interface {
	GetProperties() []*model.ResourceProperty
	GetTypes() []*model.ResourceSubType
}

func RecordLikeAsRecord(record RecordLike) *model.Record {
	if record == nil {
		return nil
	}

	var properties = record.GetProperties()

	if properties == nil {
		properties = make(map[string]*structpb.Value)
	}

	return &model.Record{
		Properties: properties,
	}
}

func RecordLikeAsRecords(record []RecordLike) []*model.Record {
	records := make([]*model.Record, 0, len(record))

	for _, r := range record {
		records = append(records, RecordLikeAsRecord(r))
	}

	return records
}

func RecordLikeAsRecords2(record []*model.Record) []RecordLike {
	records := make([]RecordLike, 0, len(record))

	for _, r := range record {
		records = append(records, RecordLikeAsRecord(r))
	}

	return records
}

func UpdateRecordsProperties(record RecordLike, properties map[string]*structpb.Value) {
	for key, value := range properties {
		record.GetProperties()[key] = value
	}
}

func NewRecordLike() RecordLike {
	return &model.Record{
		Properties: make(map[string]*structpb.Value),
	}
}
