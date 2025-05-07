package abs

import (
	"fmt"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/util"
	"google.golang.org/protobuf/types/known/structpb"
	"time"
)

func ValueToRecord(resource *model.Resource, resultExported interface{}) (*model.Record, errors.ServiceError) {
	recordObj, ok := resultExported.(map[string]interface{})

	if !ok {
		return nil, errors.LogicalError.WithDetails(fmt.Sprintf("Cannot accept nano function result: %v", resultExported))
	}

	var record = new(model.Record)
	record.Properties = make(map[string]*structpb.Value)

	var props = util.GetNamedMap(resource.Properties)

	for key, value := range recordObj {
		var prop = props[key]
		if timeValue, ok := value.(time.Time); ok {
			if prop.Type == model.ResourceProperty_TIME {
				value = timeValue.UTC().Format(time.TimeOnly)
			} else if prop.Type == model.ResourceProperty_DATE {
				value = timeValue.UTC().Format(time.DateOnly)
			} else {
				value = timeValue.UTC().Format(time.RFC3339)
			}
		}
		sv, verr := structpb.NewValue(value)

		if verr != nil {
			return nil, errors.LogicalError.WithDetails(verr.Error())
		}

		record.Properties[key] = sv
	}
	return record, nil
}
